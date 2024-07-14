package biz

import (
	"context"
	"gogo/modules/auth/model"

	"github.com/alexedwards/argon2id"
)

type Password struct {
	plaintext *string
	hash      string
}

func (p *Password) Set(plaintextPassword string) error {
	hash, err := argon2id.CreateHash(plaintextPassword, argon2id.DefaultParams)
	if err != nil {
		return err
	}
	p.plaintext = &plaintextPassword
	p.hash = hash
	return nil
}

type CreateUserStorage interface {
	CreateUser(ctx context.Context, data *model.UsersCreation) error
	GetUserByUsername(ctx context.Context, userData *model.UserLogin) (*model.Users, error)
}

type createUserBiz struct {
	store CreateUserStorage
}

func NewcreateUserBiz(store CreateUserStorage) *createUserBiz {
	return &createUserBiz{store: store}
}

func (biz *createUserBiz) CreateUser(ctx context.Context, data *model.UsersCreation) error {

	//Check if user already exists
	userData := &model.UserLogin{
		Username: data.Username,
	}

	_, err := biz.store.GetUserByUsername(ctx, userData)

	if err == nil {
		return model.ErrorUserNameIsExist
	}

	//Handle to save user with hashed password
	var hashPassword Password
	if err := hashPassword.Set(data.Password); err != nil {
		return err
	}

	data.Password = hashPassword.hash

	if err := biz.store.CreateUser(ctx, data); err != nil {
		return err
	}

	return nil
}
