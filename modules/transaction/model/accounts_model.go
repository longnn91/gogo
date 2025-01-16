package model

import (
	"gogo/common"
)

type Account struct {
	common.SQLModel
	Owner    string `json:"owner" form:"owner" gorm:"column:owner;"`
	Balance  int64  `json:"balance" form:"balance" gorm:"column:balance;"`
	Currency string `json:"currency" form:"currency" gorm:"column:currency;"`
}

func (Account) TableName() string {
	return "accounts"
}

type AccountUpdate struct {
	Balance int64 `json:"balance" form:"balance" gorm:"column:balance;"`
}

func (AccountUpdate) TableName() string {
	return Account{}.TableName()
}
