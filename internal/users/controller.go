package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService UserServiceInterface
}

func NewUserController(service UserServiceInterface) UserController {
	return UserController{
		UserService: service,
	}
}

func (uc *UserController) GetAll(c *gin.Context) {
	users, err := uc.UserService.GetAll()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": users})
}

func (uc *UserController) GetByID(c *gin.Context) {
	id := c.Param("id")

	user, err := uc.UserService.GetByID(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}
