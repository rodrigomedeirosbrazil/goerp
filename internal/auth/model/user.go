package models

import (
	bcrypt "goerp/internal/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:uuid;"`
	Name     string    `json:"name" binding:"required" validate:"required,min=3"`
	Email    string    `json:"email" binding:"required" validate:"required,email" gorm:"unique;not null"`
	Password string    `json:"password" binding:"required" validate:"required,min=6"`
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
