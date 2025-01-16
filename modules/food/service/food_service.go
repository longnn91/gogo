package service

import (
	"context"
	"gogo/modules/food/model"
)

type FoodActions interface {
	CreateFood(ctx context.Context, data *model.FoodCreation) error
}

type foodService struct {
	modelActions FoodActions
}

func (service *foodService) CreateFood(ctx context.Context, data *model.FoodCreation) error {
	if err := service.modelActions.CreateFood(ctx, data); err != nil {
		return err
	}

	return nil
}

func GetFoodService(modelActions FoodActions) *foodService {
	return &foodService{modelActions: modelActions}
}
