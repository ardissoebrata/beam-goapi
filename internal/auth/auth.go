package auth

import (
	"api_golang/internal/users"
	"api_golang/pkg/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterPackage(db *gorm.DB, route *gin.Engine) {
	authRepository := users.NewUserRepository(db)
	authService := NewAuthService(authRepository)
	authController := NewAuthController(authService)

	routes := route.Group("/v1/auth")
	routes.POST("/login", authController.Login)
	routes.POST("/access-token", middleware.Jwt(db), authController.GenerateAccessToken)
	routes.POST("/logout", middleware.Jwt(db), authController.Logout)
}
