package service

import (
	"context"
	model "gogo/modules/transaction/model"
)

type CreateTransferActions interface {
	CreateTransferData(ctx context.Context, data *model.TransfersCreation) error
}

type createTransferService struct {
	modelActions CreateTransferActions
}

func (service *createTransferService) CreateTransfer(ctx context.Context, data *model.TransfersCreation) error {
	if err := service.modelActions.CreateTransferData(ctx, data); err != nil {
		return err
	}

	return nil
}

func GetCreateTransferService(modelActions CreateTransferActions) *createTransferService {
	return &createTransferService{modelActions: modelActions}
}
