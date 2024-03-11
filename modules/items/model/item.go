package model

import (
	"errors"
	"gogo/common"
)

var (
	ErrorTitleIsBlank = errors.New("title is required")
	ErrorItemNotFound = errors.New("item not found")
)

type TodoItem struct {
	common.SQLModel
	Title       string `json:"title" gorm:"column:title;"`
	Description string `json:"description" gorm:"column:description;"`
	Status      string `json:"status" gorm:"column:status;"`
}

func (TodoItem) TableName() string {
	return "todo_items"
}

type TodoItemsCreation struct {
	Id          int    `json:"_" gorm:"column:id;"`
	Title       string `json:"title" gorm:"column:title;"`
	Description string `json:"description" gorm:"column:description;"`
	Status      string `json:"status" gorm:"column:status;"`
}

func (TodoItemsCreation) TableName() string {
	return TodoItem{}.TableName()
}

type TodoItemsUpdate struct {
	Title       *string `json:"title" gorm:"column:title;"`
	Description *string `json:"description" gorm:"column:description;"`
	Status      *string `json:"status" gorm:"column:status;"`
}

func (TodoItemsUpdate) TableName() string {
	return TodoItem{}.TableName()
}
