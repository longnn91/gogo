package biz

import (
	"context"
	"gogo/modules/items/model"
	// Add the missing import statement for the "store" package
)

type GetItemStorage interface {
	GetItem(ctx context.Context, id int) (*model.TodoItem, error)
}

type getItemBiz struct {
	store GetItemStorage
}

func NewGetItemBiz(store GetItemStorage) *getItemBiz {
	return &getItemBiz{store: store}
}

func (biz *getItemBiz) GetItemById(ctx context.Context, id int) (*model.TodoItem, error) {
	data, err := biz.store.GetItem(ctx, id)
	if err != nil {
		return nil, err
	}

	return data, nil
}
