package biz

import (
	"context"
	"gogo/modules/auth/model"
	"log"
	"os"
	"time"

	"github.com/alexedwards/argon2id"
	"github.com/golang-jwt/jwt/v5"
)

func createToken(username string) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": username,                              // Subject (user identifier)
		"aud": "user",                                // Audience (user role)
		"exp": time.Now().Add(time.Hour * 60).Unix(), // Expiration time
		"iat": time.Now().Unix(),                     // Issued at
	})

	secretKey := os.Getenv("SECRET_KEY")
	jwtKey := []byte(secretKey)
	tokenString, err := claims.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (p *Password) Matches(plaintextPassword string) (bool, error) {
	match, err := argon2id.ComparePasswordAndHash(plaintextPassword, p.hash)
	if err != nil {
		log.Fatal(err)
	}

	return match, nil
}

type GetUserStorage interface {
	GetUser(ctx context.Context, username string) (*model.Users, error)
	GetUserByUsername(ctx context.Context, userData *model.UserLogin) (*model.Users, error)
}

type getUserBiz struct {
	store GetUserStorage
}

func NewGetUserBiz(store GetUserStorage) *getUserBiz {
	return &getUserBiz{store: store}
}

func (biz *getUserBiz) GetUserByUsername(ctx context.Context, username string) (*model.Users, error) {
	data, err := biz.store.GetUser(ctx, username)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (biz *getUserBiz) GetTokenByLogin(ctx context.Context, userData *model.UserLogin) (string, error) {

	data, err := biz.store.GetUserByUsername(ctx, userData)
	if err != nil {
		return "", err
	}

	password := Password{hash: data.Password}

	match, err := password.Matches(userData.Password)

	if err != nil {
		return "", err
	}

	if token, err := createToken(userData.Username); match && err == nil {
		return token, nil
	} else {
		return "", err
	}
}
