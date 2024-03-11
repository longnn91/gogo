package biz

import (
	"context"
	"gogo/modules/items/model"
	// Add the missing import statement for the "store" package
)

type UpdateItemStorage interface {
	GetItem(ctx context.Context, id int) (*model.TodoItem, error)
	UpdateItem(ctx context.Context, id int, dataUpdate *model.TodoItemsUpdate) error
}

type updateItemBiz struct {
	store UpdateItemStorage
}

func NewUpdateItemBiz(store UpdateItemStorage) *updateItemBiz {
	return &updateItemBiz{store: store}
}

func (biz *updateItemBiz) UpdateItemById(ctx context.Context, id int, dataUpdate *model.TodoItemsUpdate) error {
	data, err := biz.store.GetItem(ctx, id)
	if err != nil {
		return err
	}

	if data == nil {
		return model.ErrorItemNotFound
	}

	if err := biz.store.UpdateItem(ctx, id, dataUpdate); err != nil {
		return err
	}

	return nil
}
