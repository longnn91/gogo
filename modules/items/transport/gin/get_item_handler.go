package ginitem

import (
	"gogo/common"
	"gogo/modules/items/storage"
	"net/http"
	"strconv"

	"gogo/modules/items/biz"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetItem(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		store := storage.NewSQLStore(db)
		business := biz.NewGetItemBiz(store)

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		data, err := business.GetItemById(c.Request.Context(), id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))

	}
}
