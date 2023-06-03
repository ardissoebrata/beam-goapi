package users

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	Save(user *User) error
	GetAll() ([]User, error)
	GetByID(id string) (User, error)
	GetByEmail(email string) (User, error)
	GetByUuid(uuid uuid.UUID) (User, error)
	GetByRefreshToken(token string) (User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) Save(user *User) error {
	err := ur.db.Save(user).Error
	return err
}

func (ur *UserRepository) GetAll() ([]User, error) {
	var users []User
	err := ur.db.Find(&users).Error
	return users, err
}

func (ur *UserRepository) GetByID(id string) (User, error) {
	var user User
	err := ur.db.First(&user, id).Error
	return user, err
}

func (ur *UserRepository) GetByEmail(email string) (User, error) {
	var user User
	err := ur.db.Where("email = ?", email).First(&user).Error
	return user, err
}

func (ur *UserRepository) GetByUuid(uuid uuid.UUID) (User, error) {
	var user User
	err := ur.db.Where("uuid = ?", uuid).First(&user).Error
	return user, err
}

func (ur *UserRepository) GetByRefreshToken(token string) (User, error) {
	var user User
	err := ur.db.Where("token = ?", token).Where("token_expire >= ?", time.Now()).First(&user).Error
	return user, err
}
