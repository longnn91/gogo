package auth

import (
	"gogo/common"
	"gogo/modules/transaction/biz"
	"gogo/modules/transaction/model"
	"gogo/modules/transaction/storage"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateTransfer(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data model.TransfersCreation
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		store := storage.NewSQLStore(db)

		business := biz.NewCreateTransferBiz(store)

		if err := business.CreateTransfer(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
