package auth

import (
	"api_golang/internal/users"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterPackage(db *gorm.DB, route *gin.Engine) {
	userRepository := users.NewUserRepository(db)
	authService := NewAuthService(userRepository)
	authController := NewAuthController(authService)

	routes := route.Group("/v1/auth")
	routes.POST("/login", authController.Login)
}
