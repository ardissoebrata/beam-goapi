package auth

import (
	"api_golang/internal/app"
	"api_golang/internal/users"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthServiceInterface interface {
	Login(email, password string) (*users.User, error)
	GenerateAccessToken(refreshToken string) (string, time.Time, error)
	Logout(uuid uuid.UUID) error
}

type AuthService struct {
	UserRepository users.UserRepositoryInterface
}

func NewAuthService(repo users.UserRepositoryInterface) *AuthService {
	return &AuthService{
		UserRepository: repo,
	}
}

func (as *AuthService) Login(email, password string) (*users.User, error) {
	var user users.User
	user, err := as.UserRepository.GetByEmail(email)
	if err != nil {
		return nil, err
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return nil, err
	}

	user.GenerateRefreshToken()
	err = as.UserRepository.Save(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (as *AuthService) GenerateAccessToken(refreshToken string) (string, time.Time, error) {
	var user users.User
	user, err := as.UserRepository.GetByRefreshToken(refreshToken)
	if err != nil {
		return "", time.Time{}, err
	}

	var timeNow = time.Now()
	var timeExpire = timeNow.Add(time.Hour * 24)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.Uuid,
		"exp": timeExpire.Unix(),
		"iat": timeNow.Unix(),
	})

	tokenString, err := token.SignedString([]byte(app.JWT_KEY))
	if err != nil {
		return "", time.Time{}, err
	}

	return tokenString, timeExpire, nil
}

func (as *AuthService) Logout(uuid uuid.UUID) error {
	var user users.User
	user, err := as.UserRepository.GetByUuid(uuid)
	if err != nil {
		return err
	}

	user.Token = ""
	user.TokenExpire = time.Time{}

	err = as.UserRepository.Save(&user)
	if err != nil {
		return err
	}
	return nil
}
