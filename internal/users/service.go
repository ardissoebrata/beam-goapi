package users

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserServiceInterface interface {
	GetAll() ([]User, error)
	GetByID(id string) (User, error)
	Login(req *LoginRequest) (User, error)
}

type UserService struct {
	UserRepository UserRepositoryInterface
}

func NewUserService(repo UserRepositoryInterface) *UserService {
	return &UserService{
		UserRepository: repo,
	}
}

func (us *UserService) GetAll() ([]User, error) {
	return us.UserRepository.GetAll()
}

func (us *UserService) GetByID(id string) (User, error) {
	return us.UserRepository.GetByID(id)
}

func (us *UserService) Login(req *LoginRequest) (User, error) {
	user, err := us.UserRepository.GetByEmail(req.Email)
	if err != nil {
		return User{}, err
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)) != nil {
		return User{}, gorm.ErrRecordNotFound
	}

	return user, nil
}
