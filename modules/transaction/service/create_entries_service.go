package service

import (
	model "gogo/modules/transaction/model"
)

type CreateEntriesActions interface {
	CreateEntriesData(data *model.EntriesCreation) error
}

type createEntriesService struct {
	modelActions CreateEntriesActions
}

func (service *createEntriesService) CreateEntries(data *model.EntriesCreation) error {
	if err := service.modelActions.CreateEntriesData(data); err != nil {
		return err
	}

	return nil
}

func GetCreateEntriesService(modelActions CreateEntriesActions) *createEntriesService {
	return &createEntriesService{modelActions: modelActions}
}
