package model

import (
	"errors"
	"gogo/common"
)

var (
	ErrorTitleIsBlank    = errors.New("Title is required")
	ErrorItemNotFound    = errors.New("Item not found")
	ErrorUserNameIsExist = errors.New("Username is exist")
)

type Users struct {
	common.SQLModel
	Username    string `json:"username" form:"username" gorm:"column:username;not null;unique"`
	Password    string `json:"password" form:"password" gorm:"column:password;not null"`
	FirstName   string `json:"first_name" form:"first_name" gorm:"column:first_name"`
	LastName    string `json:"last_name" form:"last_name" gorm:"column:last_name"`
	PhoneNumber string `json:"phone_number" form:"phone_number" gorm:"column:phone_number"`
}

func (Users) TableName() string {
	return "users"
}

type UsersCreation struct {
	Username    string `json:"username" form:"username" gorm:"column:username;"`
	Password    string `json:"password" form:"password" gorm:"column:password;"`
	FirstName   string `json:"first_name" form:"first_name" gorm:"column:first_name;"`
	LastName    string `json:"last_name" form:"last_name" gorm:"column:last_name;"`
	PhoneNumber string `json:"phone_number" form:"phone_number" gorm:"column:phone_number;"`
}

func (UsersCreation) TableName() string {
	return Users{}.TableName()
}

type UsersUpdate struct {
	Id          int    `json:"_" gorm:"column:id;"`
	Username    string `json:"username" form:"username" gorm:"column:username;"`
	Password    string `json:"password" form:"password" gorm:"column:password;"`
	FirstName   string `json:"first_name" form:"first_name" gorm:"column:first_name;"`
	LastName    string `json:"last_name" form:"last_name" gorm:"column:last_name;"`
	PhoneNumber string `json:"phone_number" form:"phone_number" gorm:"column:phone_number;"`
}

func (UsersUpdate) TableName() string {
	return Users{}.TableName()
}

var (
	ErrDuplicateEmail = errors.New("duplicate email")
)

type UserLogin struct {
	Username string `json:"username" form:"username" gorm:"column:username;not null;unique"`
	Password string `json:"password" form:"password" gorm:"column:password;not null"`
}

func (UserLogin) TableName() string {
	return Users{}.TableName()
}
