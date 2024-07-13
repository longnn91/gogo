package biz

import (
	"context"
	"gogo/modules/auth/model"
)

type CreateUserStorage interface {
	CreateUser(ctx context.Context, data *model.UsersCreation) error
}

type createUserBiz struct {
	store CreateUserStorage
}

func NewcreateUserBiz(store CreateUserStorage) *createUserBiz {
	return &createUserBiz{store: store}
}

func (biz *createUserBiz) CreateUser(ctx context.Context, data *model.UsersCreation) error {

	if err := biz.store.CreateUser(ctx, data); err != nil {
		return err
	}

	return nil
}
