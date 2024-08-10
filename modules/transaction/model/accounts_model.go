package model

import (
	"gogo/common"
)

type Accounts1 struct {
	common.SQLModel
	Owner    string `json:"owner" form:"owner" gorm:"column:owner;"`
	Balance  int64  `json:"balance" form:"balance" gorm:"column:balance;"`
	Currency string `json:"currency" form:"currency" gorm:"column:currency;"`
}

func (Accounts1) TableName() string {
	return "accounts"
}
