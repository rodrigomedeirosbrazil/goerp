package repository

import (
	model "goerp/internal/auth/model"
	database "goerp/internal/database"
)

func CreateUser(name string, email string, password string) (*model.User, error) {
	user := model.User{
		Name:     name,
		Email:    email,
		Password: password,
	}

	if err := database.DB.Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	result := database.DB.First(&user, "email = ?", email)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
