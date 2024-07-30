package model

import (
	"gogo/common"
)

type Accounts struct {
	common.SQLModel
	Owner    string `json:"owner" form:"owner" gorm:"column:owner;"`
	Balance  int64  `json:"balance" form:"balance" gorm:"column:balance;"`
	Currency string `json:"currency" form:"currency" gorm:"column:currency;"`
}

func (Accounts) TableName() string {
	return "accounts"
}
