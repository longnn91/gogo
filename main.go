package main

import (
	"gogo/modules/auth/middleware"
	authModal "gogo/modules/auth/model"
	auth "gogo/modules/auth/transport"
	item "gogo/modules/items/transport"
	transaction "gogo/modules/transaction/controller"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	//Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	//Connect to database
	dsn := os.Getenv("DATABASE_URL")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		log.Fatal("Error connecting to database", err)
	}

	//Migrate the schema

	db.AutoMigrate(&authModal.Users{})
	// db.Migrator().CreateTable(&authModal.Users{})
	// db.Migrator().DropTable(&authModal.Users{})

	//Config API use gin
	r := gin.Default()
	v1 := r.Group("/v1")
	transactionRouter := v1.Group("/transaction").Use(middleware.AuthMiddleware)
	{
		transactionRouter.POST("/", transaction.CreateTransfer(db))
	}

	itemRouter := v1.Group("/items").Use(middleware.AuthMiddleware)
	{
		itemRouter.POST("/", item.CreateItem(db))
		itemRouter.GET("/", item.ListItem(db))
		itemRouter.GET("/:id", item.GetItem(db))
		itemRouter.PUT("/:id", item.UpdateItem(db))
		itemRouter.PATCH("/:id", item.UpdateItem(db))
		itemRouter.DELETE("/:id", item.DeleteItem(db))
	}

	authRouter := v1.Group("/auth")
	{
		authRouter.POST("/login", auth.Login(db))
		authRouter.POST("/register", auth.Register(db))
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
