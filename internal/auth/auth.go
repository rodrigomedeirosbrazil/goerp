package auth

import (
	models "goerp/internal/auth/model"
	database "goerp/internal/database"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func Login(c *fiber.Ctx) error {
	emailField := c.FormValue("email")
	passwordField := c.FormValue("password")

	var user models.User
	result := database.DB.First(&user, "email = ?", emailField)

	if result.Error != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	if user.Email != emailField || user.Password != passwordField {
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
