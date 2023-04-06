package users

import "gorm.io/gorm"

type UserRepositoryInterface interface {
	GetAll() ([]User, error)
	GetByID(id string) (User, error)
	GetByEmail(email string) (User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
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
