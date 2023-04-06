package auth

import (
	"api_golang/internal/users"
	"api_golang/pkg/validation"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	AuthService AuthServiceInterface
}

func NewAuthController(service AuthServiceInterface) AuthController {
	return AuthController{
		AuthService: service,
	}
}

func (ac *AuthController) Login(c *gin.Context) {
	var input users.LoginRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validation.Errors(err)})
		return
	}
	token, err := ac.AuthService.Login(input.Email, input.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}
