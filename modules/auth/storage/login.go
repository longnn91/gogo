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
