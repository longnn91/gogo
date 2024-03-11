package storage

import (
	"context"
	"gogo/modules/items/model"
)

func (s *sqlStore) UpdateItem(ctx context.Context, id int, dataUpdate *model.TodoItemsUpdate) error {
	if err := s.db.Where("id = ?", id).Updates(dataUpdate).Error; err != nil {
		return err
	}

	return nil
}
