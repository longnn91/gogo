package main

import (
	"fmt"
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

	//Config API use gin
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.POST("/items", ginitem.CreateItem(db))
		v1.GET("/items/:id", ginitem.GetItem(db))
		v1.PUT("/items/:id", ginitem.UpdateItem(db))
		v1.PATCH("/items/:id", ginitem.UpdateItem(db))
		v1.DELETE("/items/:id", ginitem.DeleteItem(db))
		v1.GET("/items", ginitem.ListItem(db))
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
