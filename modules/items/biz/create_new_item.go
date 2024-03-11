package biz

import (
	"context"
	"gogo/modules/items/model"
	"strings"
)

type CreateItemStorage interface {
	CreateItem(ctx context.Context, data *model.TodoItemsCreation) error
}

type createItemBiz struct {
	store CreateItemStorage
}

func NewCreateItemBiz(store CreateItemStorage) *createItemBiz {
	return &createItemBiz{store: store}
}

func (biz *createItemBiz) CreateItem(ctx context.Context, data *model.TodoItemsCreation) error {
	title := strings.TrimSpace(data.Title)

	if title == "" {
		return model.ErrorTitleIsBlank
	}

	if err := biz.store.CreateItem(ctx, data); err != nil {
		return err
	}

	return nil
}
