package service

import (
	"context"
	model "gogo/modules/transaction/model"
)

type TransferActions interface {
	CreateTransfer(ctx context.Context, data *model.TransfersCreation) error
}

type transferService struct {
	modelActions TransferActions
}

func (service *transferService) CreateTransfer(ctx context.Context, data *model.TransfersCreation) error {
	if err := service.modelActions.CreateTransfer(ctx, data); err != nil {
		return err
	}

	return nil
}

func GetTransferService(modelActions TransferActions) *transferService {
	return &transferService{modelActions: modelActions}
}
