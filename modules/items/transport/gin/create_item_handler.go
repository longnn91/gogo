package ginitem

import (
	"gogo/common"
	"gogo/modules/items/biz"
	"gogo/modules/items/model"
	"gogo/modules/items/storage"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateItem(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data model.TodoItemsCreation
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		store := storage.NewSQLStore(db)

		business := biz.NewCreateItemBiz(store)

		if err := business.CreateItem(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.Id))
	}
}
