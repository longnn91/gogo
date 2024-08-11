package database

import (
	"gogo/modules/transaction/model"
)

func (s *sqlStore) UpdateAccount(data *model.AccountUpdate, id int) error {
	if err := s.db.Where("id = ?", id).Updates(data).Error; err != nil {
		return err
	}

	return nil
}

func (s *sqlStore) GetAccountById(accountId int) (*model.Account, error) {
	var data model.Account
	if err := s.db.First(&data, accountId).Error; err != nil {
		return nil, err
	}
	return &data, nil
}
