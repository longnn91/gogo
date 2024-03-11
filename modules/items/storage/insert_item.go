package storage

import (
	"context"
	"gogo/modules/items/model"
)

func (s *sqlStore) CreateItem(ctx context.Context, data *model.TodoItemsCreation) error {
	if err := s.db.Create(&data); err.Error != nil {
		return err.Error
	}

	return nil
}
