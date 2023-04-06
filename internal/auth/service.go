package auth

import (
	"api_golang/internal/app"
	"api_golang/internal/users"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type AuthServiceInterface interface {
	Login(email, password string) (string, error)
}

type AuthService struct {
	UserRepository users.UserRepositoryInterface
}

func NewAuthService(repo users.UserRepositoryInterface) *AuthService {
	return &AuthService{
		UserRepository: repo,
	}
}

func (as *AuthService) Login(email, password string) (string, error) {
	var user users.User
	user, err := as.UserRepository.GetByEmail(email)
	if err != nil {
		return "", err
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uuid":  user.Uuid,
		"name":  user.Full_name,
		"email": email,
	})

	tokenString, err := token.SignedString([]byte(app.JWT_KEY))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
