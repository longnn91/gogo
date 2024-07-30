package storage

import (
	"context"
	"gogo/common"
	"gogo/modules/transaction/model"
)

func (s *sqlStore) CreateTransfer(ctx context.Context, data *model.TransfersCreation) error {
	if err := s.db.Create(&data); err.Error != nil {
		// fmt.Print("Error: User: ", common.ErrDB(err.Error))
		return common.ErrDB(err.Error)
	}

	return nil
}
