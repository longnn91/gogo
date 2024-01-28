package main

import (
	"fmt"
	"gogo/config"
	"gogo/helper"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Hello, World!  33331")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	db, err := config.ConnectDB()
	helper.ErrorPanic(err)

	defer db.Disconnect()

	r := gin.Default()
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
