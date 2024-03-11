package main

import (
	"fmt"
	"gogo/common"
	"gogo/modules/items/model"
	ginitem "gogo/modules/items/transport/gin"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("Hello, World!  33331")

	//Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	//Connect to database
	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to database", err)
	}

	fmt.Println(db)

	//Config API use gin
	r := gin.Default()

	//CURD: Create, Read, Update, Delete
	//POST v1/items (Create a new item)
	//GET v1/items (Read all items)
	//GET v1/items/:id (Read a item)
	//(PUT | PATCH) v1/items/:id (Update a item)
	//DELETE v1/items/:id (Delete a item)

	v1 := r.Group("/v1")
	{
		v1.POST("/items", ginitem.CreateItem(db))
		v1.GET("/items/:id", ginitem.GetItem(db))
		v1.PUT("/items/:id", ginitem.UpdateItem(db))
		v1.PATCH("/items/:id", ginitem.UpdateItem(db))
		v1.DELETE("/items/:id", ginitem.DeleteItem(db))
		v1.GET("/items", GetItems(db))
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	server_err := r.Run(":" + os.Getenv("PORT")) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if server_err != nil {
		log.Fatal("Error loading .env file", server_err)
		panic(server_err)
	}
}

// func CreateItem(db *gorm.DB) func(*gin.Context) {
// 	return func(c *gin.Context) {
// 		var data model.TodoItemsCreation
// 		if err := c.ShouldBind(&data); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}

// 		fmt.Println(data)

// 		if err := db.Create(&data); err.Error != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error})
// 			return
// 		}

// 		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.Id))
// 	}
// }

// func GetItem(db *gorm.DB) func(*gin.Context) {
// 	return func(c *gin.Context) {
// 		var data model.TodoItem

// 		fmt.Println(data)

// 		if err := db.First(&data, c.Param("id")).Error; err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}

// 		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
// 	}
// }

// func UpdateItem(db *gorm.DB) func(*gin.Context) {
// 	return func(c *gin.Context) {
// 		var data model.TodoItemsUpdate
// 		var id = c.Param("id")
// 		if err := c.ShouldBind(&data); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}

// 		if err := db.Where("id = ?", id).Updates(&data).Error; err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}

// 		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
// 	}
// }

// func DeleteItem(db *gorm.DB) func(*gin.Context) {
// 	return func(c *gin.Context) {
// 		if err := db.Table(model.TodoItem{}.TableName()).Where("id = ?", c.Param("id")).Updates(map[string]interface{}{
// 			"status": "Deleted",
// 		}).Error; err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}

// 		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
// 	}
// }

func GetItems(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data []model.TodoItem
		var query common.Paging

		if err := c.ShouldBind(&query); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		query.Process()

		db = db.Where("status <> ?", "Deleted")

		if err := db.Table(model.TodoItem{}.TableName()).Count(&query.Total).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := db.Order("id desc").
			Offset((query.Page - 1) * query.Limit).
			Limit(query.Limit).Find(&data).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(data, query, nil))
	}
}
