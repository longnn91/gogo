package biz

import (
	"context"
	"gogo/modules/items/model"
)

type DeleteItemStorage interface {
	GetItem(ctx context.Context, id int) (*model.TodoItem, error)
	DeleteItem(ctx context.Context, id int) error
}

type deleteItemBiz struct {
	store DeleteItemStorage
}

func NewDeleteItemBiz(store DeleteItemStorage) *deleteItemBiz {
	return &deleteItemBiz{store: store}
}

func (biz *deleteItemBiz) DeleteItemById(ctx context.Context, id int) error {
	data, err := biz.store.GetItem(ctx, id)
	if err != nil {
		return err
	}

	if data == nil {
		return model.ErrorItemNotFound
	}

	if err := biz.store.DeleteItem(ctx, id); err != nil {
		return err
	}

	return nil
}
