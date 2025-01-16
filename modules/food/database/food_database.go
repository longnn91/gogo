package database

import (
	"context"
	"gogo/common"
	"gogo/modules/food/model"
)

func (s *sqlStore) CreateFood(ctx context.Context, data *model.FoodCreation) error {
	if err := s.db.Create(&data); err.Error != nil {
		// fmt.Print("Error: User: ", common.ErrDB(err.Error))
		return common.ErrDB(err.Error)
	}

	return nil
}
