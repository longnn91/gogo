package storage

import (
	"context"
	"gogo/modules/items/model"
)

func (s *sqlStore) DeleteItem(ctx context.Context, id int) error {
	if err := s.db.Table(model.TodoItem{}.TableName()).Where("id = ?", id).Updates(map[string]interface{}{
		"status": "Deleted",
	}).Error; err != nil {
		return err
	}

	return nil
}
