package storage

import (
	"context"
	"gogo/modules/items/model"
)

func (s *sqlStore) GetItem(ctx context.Context, id int) (*model.TodoItem, error) {
	var data model.TodoItem
	if err := s.db.First(&data, id).Error; err != nil {
		return nil, err
	}

	return &data, nil
}
