package auth

import (
	"fmt"
	"gogo/common"
	"gogo/modules/auth/biz"
	"gogo/modules/auth/model"
	"gogo/modules/auth/storage"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data model.UsersCreation
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		store := storage.NewSQLStore(db)

		business := biz.NewcreateUserBiz(store)

		if err := business.CreateUser(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}

func Login(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data model.UsersCreation
		var paging common.Paging
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// store := storage.NewSQLStore(db)

		// business := biz.NewcreateUserBiz(store)

		// if err := business.CreateUser(c.Request.Context(), &data); err != nil {
		// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		// 	return
		// }

		fmt.Println("paging", paging)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
