package database

import (
	"gogo/common"
	"gogo/modules/transaction/model"
)

func (s *sqlStore) CreateEntries(data *model.EntriesCreation) error {
	if err := s.db.Create(&data); err.Error != nil {
		// fmt.Print("Error: User: ", common.ErrDB(err.Error))
		return common.ErrDB(err.Error)
	}

	return nil
}
