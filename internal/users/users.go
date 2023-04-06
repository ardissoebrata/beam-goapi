package users

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterPackage(db *gorm.DB, route *gin.Engine) {
	userRepository := NewUserRepository(db)
	userService := NewUserService(userRepository)
	userController := NewUserController(userService)

	route.GET("/v1/users", userController.GetAll)
	route.GET("/v1/users/:id", userController.GetByID)
	route.POST("/v1/users/login", userController.Login)
}
