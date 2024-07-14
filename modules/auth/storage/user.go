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

func (s *sqlStore) GetItem(ctx context.Context, id int) (*model.Users, error) {
	var data model.Users
	if err := s.db.First(&data, id).Error; err != nil {
		return nil, err
	}

	return &data, nil
}
