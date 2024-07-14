package middleware

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func verifyToken(tokenString string) bool {
	// Parse the token with the secret key

	secretKey := os.Getenv("SECRET_KEY")
	jwtKey := []byte(secretKey)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	// Check for verification errors
	if err != nil {
		return false
	}

	// Check if the token is valid
	if !token.Valid {
		return false
	}

	// Return the verified token
	return true
}

func AuthMiddleware(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")

	if authHeader == "" {
		c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	byPass := verifyToken(authHeader)

	if !byPass {
		c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	c.Next()
}
