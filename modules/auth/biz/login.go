package biz

import (
	"context"
	"gogo/modules/auth/model"
	// Add the missing import statement for the "store" package
)

type GetUserStorage interface {
	GetUser(ctx context.Context, id int) (*model.Users, error)
}

type getUserBiz struct {
	store GetUserStorage
}

func NewGetUserBiz(store GetUserStorage) *getUserBiz {
	return &getUserBiz{store: store}
}

func (biz *getUserBiz) GetUserById(ctx context.Context, id int) (*model.Users, error) {
	data, err := biz.store.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}

	return data, nil
}
