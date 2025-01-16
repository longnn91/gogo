package model

import (
	"gogo/common"
)

type Food struct {
	common.SQLModel
	Name       string `json:"name" form:"name" gorm:"column:name;"`
	Price      int64  `json:"price" form:"price" gorm:"column:price;"`
	FoodImage  string `json:"food_image" form:"food_image" gorm:"column:food_image;"`
	MenuId     int    `json:"menu_id" form:"menu_id" gorm:"column:menu_id;"`
	CategoryId int    `json:"category_id" form:"category_id" gorm:"column:category_id;"`
}

func (Food) TableName() string {
	return "foods"
}

type FoodCreation struct {
	Name       string `json:"name" form:"name" gorm:"column:name;"`
	Price      int64  `json:"price" form:"price" gorm:"column:price;"`
	FoodImage  string `json:"food_image" form:"food_image" gorm:"column:food_image;"`
	MenuId     int    `json:"menu_id" form:"menu_id" gorm:"column:menu_id;"`
	CategoryId int    `json:"category_id" form:"category_id" gorm:"column:category_id;"`
}

func (FoodCreation) TableName() string {
	return Food{}.TableName()
}

type FoodUpdate struct {
	Name      string `json:"name" form:"name" gorm:"column:name;"`
	Price     int64  `json:"price" form:"price" gorm:"column:price;"`
	FoodImage string `json:"food_image" form:"food_image" gorm:"column:food_image;"`
}

func (FoodUpdate) TableName() string {
	return Food{}.TableName()
}
