package users

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Constants
const ROLE_ADMIN = "admin"
const ROLE_USER = "user"
const STATUS_ACTIVE = "active"

type User struct {
	gorm.Model  `json:"-"`
	Uuid        uuid.UUID `json:"uuid" gorm:"size:36;uniqueIndex;not null"`
	Full_name   string    `json:"full_name" gorm:"size:255;not null"`
	Email       string    `json:"email" gorm:"size:255;uniqueIndex;not null"`
	Password    string    `json:"-" gorm:"size:255;not null"`
	Role        string    `json:"role" gorm:"size:255;not null;default:'user'"`
	Token       string    `json:"-" gorm:"size:255;uniqueIndex;not null"`
	TokenExpire time.Time `json:"-" gorm:"type:timestamp;not null"`
	Status      string    `json:"status" gorm:"size:255;not null;default:'active'"`
}

func (u *User) GenerateRefreshToken() {
	u.Token = uuid.New().String()
	u.TokenExpire = time.Now().Add(time.Hour * 24 * 7)
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.Uuid = uuid.New()
	u.GenerateRefreshToken()
	return
}
