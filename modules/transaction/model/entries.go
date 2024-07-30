package model

import (
	"gogo/common"
)

type Entries struct {
	common.SQLModel
	AccountId int   `json:"account_id" form:"account_id" gorm:"column:account_id;"`
	Amount    int64 `json:"amount" form:"amount" gorm:"column:amount;"`
}

func (Entries) TableName() string {
	return "entries"
}

type EntriesCreation struct {
	AccountId int   `json:"account_id" form:"account_id" gorm:"column:account_id;"`
	Amount    int64 `json:"amount" form:"amount" gorm:"column:amount;"`
}

func (EntriesCreation) TableName() string {
	return Entries{}.TableName()
}
