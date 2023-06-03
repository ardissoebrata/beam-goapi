package auth

import (
	"api_golang/pkg/validation"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AuthController struct {
	AuthService AuthServiceInterface
}

func NewAuthController(service AuthServiceInterface) AuthController {
	return AuthController{
		AuthService: service,
	}
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (ac *AuthController) Login(c *gin.Context) {
	var input LoginRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validation.Errors(err)})
		return
	}
	user, err := ac.AuthService.Login(input.Email, input.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, expire, err := ac.AuthService.GenerateAccessToken(user.Token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"user":                 user,
		"refresh_token":        user.Token,
		"refresh_token_expire": user.TokenExpire,
		"access_token":         token,
		"access_token_expire":  expire,
	})
}

type GenerateAccessTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

func (ac *AuthController) GenerateAccessToken(c *gin.Context) {
	var input GenerateAccessTokenRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validation.Errors(err)})
		return
	}
	token, expire, err := ac.AuthService.GenerateAccessToken(input.RefreshToken)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"access_token":        token,
		"access_token_expire": expire,
	})
}

func (ac *AuthController) Logout(c *gin.Context) {
	fmt.Println("logout")
	uuid, err := uuid.Parse(c.GetString("userUuid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ac.AuthService.Logout(uuid)
	c.JSON(http.StatusOK, gin.H{
		"message": "logged out",
	})
}
