package models

import (
	bcrypt "goerp/internal/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:uuid;"`
	Name     string    `json:"name" binding:"required"`
	Email    string    `json:"email" binding:"required" gorm:"unique;not null"`
	Password string    `json:"password" binding:"required"`
}

type Users struct {
	Users []User `json:"users"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.ID = uuid.New()

	user.Password, err = bcrypt.HashPassword(user.Password)
	if err != nil {
		return err
	}

	return
}
