package middleware

import (
	"api_golang/internal/app"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the JWT from the request header
		tokenString := c.GetHeader("Authorization")

		// Parse the token and verify its signature
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Check that the signing method is what we expect
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			// Return the secret that was used to sign the token
			return []byte(app.JWT_KEY), nil
		})

		if err != nil {
			// Token is invalid
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "invalid token",
			})
		}

		// Check if the token is valid
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// Set the user's ID as a variable in the context
			c.Set("userId", claims["userId"])
			c.Next()
		} else {
			// Token is invalid
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "invalid token",
			})
		}
	}
}
