package storage

import (
	"context"
	"fmt"
	"gogo/common"
	"gogo/modules/items/model"
)

func (s *sqlStore) CreateItem(ctx context.Context, data *model.TodoItemsCreation) error {
	if err := s.db.Create(&data); err.Error != nil {
		fmt.Print("Error: shenL: ", common.ErrDB(err.Error))
		return common.ErrDB(err.Error)
	}

	return nil
}
