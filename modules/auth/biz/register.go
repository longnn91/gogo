package biz

import (
	"context"
	"fmt"
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
}

type createUserBiz struct {
	store CreateUserStorage
}

func NewcreateUserBiz(store CreateUserStorage) *createUserBiz {
	return &createUserBiz{store: store}
}

func (biz *createUserBiz) CreateUser(ctx context.Context, data *model.UsersCreation) error {

	var hashPassword Password
	if err := hashPassword.Set(data.Password); err != nil {
		return err
	}

	// Set the hashed password
	data.Password = hashPassword.hash

	// Print the hashed password for debugging
	fmt.Println("Hashed Password:", hashPassword.hash)

	// Save the user to the database
	if err := biz.store.CreateUser(ctx, data); err != nil {
		return err
	}

	return nil
}
