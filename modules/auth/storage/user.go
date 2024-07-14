package storage

import (
	"context"
	"fmt"
	"gogo/common"
	"gogo/modules/auth/model"
)

func (s *sqlStore) CreateUser(ctx context.Context, data *model.UsersCreation) error {
	if err := s.db.Create(&data); err.Error != nil {
		fmt.Print("Error: User: ", common.ErrDB(err.Error))
		return common.ErrDB(err.Error)
	}

	return nil
}

func (s *sqlStore) GetUser(ctx context.Context, username string) (*model.Users, error) {
	var data model.Users
	if err := s.db.First(&data, username).Error; err != nil {
		return nil, err
	}

	return &data, nil
}

func (s *sqlStore) GetTokenByLogin(ctx context.Context, userData *model.UserLogin) (*model.Users, error) {
	var data model.Users
	username := userData.Username
	if err := s.db.Where("username = ?", username).First(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}
