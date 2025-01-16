package service

import (
	model "gogo/modules/transaction/model"
)

type EntriesActions interface {
	CreateEntries(data *model.EntriesCreation) error
}

type entriesService struct {
	modelActions EntriesActions
}

func (service *entriesService) CreateEntries(data *model.EntriesCreation) error {
	if err := service.modelActions.CreateEntries(data); err != nil {
		return err
	}

	return nil
}

func GetEntriesService(modelActions EntriesActions) *entriesService {
	return &entriesService{modelActions: modelActions}
}
