package biz

import (
	"context"
	"gogo/modules/transaction/model"
)

type CreateTransferStorage interface {
	CreateTransfer(ctx context.Context, data *model.TransfersCreation) error
}

type createTransferBiz struct {
	store CreateTransferStorage
}

func (biz *createTransferBiz) CreateTransfer(ctx context.Context, data *model.TransfersCreation) error {
	if err := biz.store.CreateTransfer(ctx, data); err != nil {
		return err
	}

	return nil
}

func NewCreateTransferBiz(store CreateTransferStorage) *createTransferBiz {
	return &createTransferBiz{store: store}
}
