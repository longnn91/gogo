package service

import model "gogo/modules/transaction/model"

type AccountActions interface {
	GetAccountById(accountId int) (*model.Account, error)
	UpdateAccount(data *model.AccountUpdate, id int) error
}

type accountService struct {
	modelActions AccountActions
}

func (service *accountService) GetAccountById(accountId int) (*model.Account, error) {
	data, err := service.modelActions.GetAccountById(accountId)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (service *accountService) UpdateAccount(data *model.AccountUpdate, id int) error {
	if err := service.modelActions.UpdateAccount(data, id); err != nil {
		return err
	}

	return nil
}

func GetAccountService(modelActions AccountActions) *accountService {
	return &accountService{modelActions: modelActions}
}
