package auth

import (
	"errors"
	"os"
	"strings"
	"time"

	models "goerp/internal/auth/model"
	database "goerp/internal/database"
	bcrypt "goerp/internal/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

func Login(c *fiber.Ctx) error {
	emailField := c.FormValue("email")
	passwordField := c.FormValue("password")

	var user models.User
	result := database.DB.First(&user, "email = ?", emailField)

	if result.Error != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	if user.Email != emailField {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	if !bcrypt.CheckPasswordHash(passwordField, user.Password) {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	// Create the Claims
	claims := jwt.MapClaims{
		"name": user.Name,
		"exp":  time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	jwt_key := os.Getenv("JWT_KEY")

	t, err := token.SignedString([]byte(jwt_key))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": t})
}

func Signup(c *fiber.Ctx) error {
	db := database.DB
	user := new(models.User)

	// Store the body in the user and return error if encountered
	err := c.BodyParser(user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}

	err = db.Create(&user).Error

	if err != nil && (errors.Is(err, gorm.ErrDuplicatedKey) || strings.Contains(err.Error(), "UNIQUE constraint failed")) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Email in use."})
	}

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create user", "data": err})
	}

	// Return the created user
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "User has created", "data": user})
}
