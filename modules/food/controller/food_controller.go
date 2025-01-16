package transaction

import (
	"gogo/common"
	"gogo/modules/food/database"
	"gogo/modules/food/model"
	"gogo/modules/food/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateFood(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data model.FoodCreation
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		store := database.NewSQLStore(db)

		business := service.GetFoodService(store)

		//Create transfer
		if err := business.CreateFood(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
