package middleware

import (
	"api_golang/internal/app"
	"api_golang/internal/users"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func Jwt(db *gorm.DB) gin.HandlerFunc {
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
			return
		}

		// Check if the token is valid
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// Get the user's UUID from the token's claims
			uuid, err := uuid.Parse(claims["sub"].(string))
			if err != nil {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"error": "invalid token",
				})
				return
			}

			// Check if token is not expired
			if claims["exp"].(float64) < float64(time.Now().Unix()) {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"error": "token expired",
				})
				return
			}

			// Check if refresh token is still valid
			user, err := users.NewUserRepository(db).GetByUuid(uuid)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"error": "invalid token",
				})
				return
			}
			if user.TokenExpire.Before(time.Now()) {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"error": "refresh token expired",
				})
				return
			}

			// Set the user's ID as a variable in the context
			c.Set("userUuid", uuid.String())
			c.Next()
		} else {
			// Token is invalid
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "invalid token",
			})
		}
	}
}
