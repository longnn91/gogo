package main

import (
	"fmt"
	item "gogo/modules/items/transport"
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

	//Handle authentication with JWT
	// secretKey := os.Getenv("SECRET_KEY")

	// func createToken(username string) (string, error) {
	// 	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	// 		"sub": username,
	// 		"iss": "gogo",
	// 		"aud": "user",
	// 		"exp": time.Now().Add(time.Hour).Unix(),
	// 		"iat": time.Now().Unix(),
	// 	})

	// 	tokenString, err := claims.SignedString(secretKey)
	// 	if err != nil {
	// 		return "", err
	// 	}

	// 	fmt.Printf("Token claims added: %+v\n", claims)
	// 	return tokenString, nil
	// }

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		log.Fatal("Error connecting to database", err)
	}

	//Config API use gin
	r := gin.Default()
	v1 := r.Group("/v1")
	itemRouter := v1.Group("/items")
	{
		itemRouter.POST("/", item.CreateItem(db))
		itemRouter.GET("/", item.ListItem(db))
		itemRouter.GET("/:id", item.GetItem(db))
		itemRouter.PUT("/:id", item.UpdateItem(db))
		itemRouter.PATCH("/:id", item.UpdateItem(db))
		itemRouter.DELETE("/:id", item.DeleteItem(db))
	}

	// authRouter := v1.Group("/auth")
	// {
	// 	authRouter.POST("/login", auth.Login(db))
	// }

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
