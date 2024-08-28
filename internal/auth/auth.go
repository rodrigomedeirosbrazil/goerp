package auth

import (
	"errors"
	"strings"

	model "goerp/internal/auth/model"
	repository "goerp/internal/auth/repository"
	bcrypt "goerp/internal/utils/bcrypt"
	jwt "goerp/internal/utils/jwt"
	validator "goerp/internal/utils/validator"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Login(c *fiber.Ctx) error {
	emailField := c.FormValue("email")
	passwordField := c.FormValue("password")

	if (emailField == "") || (passwordField == "") {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	user, err := repository.GetUserByEmail(emailField)

	if err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	if user.Email != emailField {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	if !bcrypt.CheckPasswordHash(passwordField, user.Password) {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	token, err := jwt.CreateToken()
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": token})
}

func Signup(c *fiber.Ctx) error {
	user := new(model.User)

	// Store the body in the user and return error if encountered
	err := c.BodyParser(user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}

	err = validator.ValidateStruct(user)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.
			Map{
			"status":  "error",
			"message": "Something's wrong with your input",
			"data":    validator.ToErrResponse(err).Errors,
		})
	}

	user, err = repository.CreateUser(user.Name, user.Email, user.Password)

	if err != nil && (errors.Is(err, gorm.ErrDuplicatedKey) || strings.Contains(err.Error(), "UNIQUE constraint failed")) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Email in use."})
	}

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.
			Map{
			"status":  "error",
			"message": "Could not create user",
			"data":    err,
		})
	}

	// Return the created user
	return c.Status(fiber.StatusCreated).JSON(fiber.
		Map{
		"status":  "success",
		"message": "User has created",
		"data":    user,
	})
}
