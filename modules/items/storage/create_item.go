package storage

import (
	"context"
	"gogo/common"
	"gogo/modules/items/model"
)

func (s *sqlStore) CreateItem(ctx context.Context, data *model.TodoItemsCreation) error {
	if err := s.db.Create(&data); err.Error != nil {
		return common.ErrDB(err.Error)
	}

	return nil
}
