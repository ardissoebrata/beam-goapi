package users

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model `json:"-"`
	Uuid       uuid.UUID `json:"uuid" gorm:"size:36;uniqueIndex;not null"`
	Full_name  string    `json:"full_name" gorm:"size:255;not null"`
	Email      string    `json:"email" gorm:"size:255;uniqueIndex;not null"`
	Password   string    `json:"-" gorm:"size:255;not null"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.Uuid = uuid.New()
	return
}
