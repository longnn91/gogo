package model

import (
	"gogo/common"
)

type Transfers struct {
	common.SQLModel
	FromAccountId int   `json:"from_account_id" form:"from_account_id" gorm:"column:from_account_id;"`
	ToAccountId   int   `json:"to_account_id" form:"to_account_id" gorm:"column:to_account_id;"`
	Amount        int64 `json:"amount" form:"amount" gorm:"column:amount;"`
}

func (Transfers) TableName() string {
	return "transfers"
}

type TransfersCreation struct {
	FromAccountId int   `json:"from_account_id" form:"from_account_id" gorm:"column:from_account_id;"`
	ToAccountId   int   `json:"to_account_id" form:"to_account_id" gorm:"column:to_account_id;"`
	Amount        int64 `json:"amount" form:"amount" gorm:"column:amount;"`
}

func (TransfersCreation) TableName() string {
	return Transfers{}.TableName()
}
