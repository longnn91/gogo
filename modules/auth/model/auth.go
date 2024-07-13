package model

import (
	"errors"
	"gogo/common"
)

var (
	ErrorTitleIsBlank = errors.New("title is required")
	ErrorItemNotFound = errors.New("item not found")
)

type Users struct {
	common.SQLModel
	Username string `json:"username" gorm:"column:username;"`
	Password string `json:"password" gorm:"column:password;"`
}

func (Users) TableName() string {
	return "users"
}

type UsersCreation struct {
	Id       int    `json:"_" gorm:"column:id;"`
	Username string `json:"username" gorm:"column:username;"`
	Password string `json:"password" gorm:"column:password;"`
}

func (UsersCreation) TableName() string {
	return Users{}.TableName()
}

type UsersUpdate struct {
	Username string `json:"username" gorm:"column:username;"`
	Password string `json:"password" gorm:"column:password;"`
}

func (UsersUpdate) TableName() string {
	return Users{}.TableName()
}
