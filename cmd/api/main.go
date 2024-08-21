package main

import (
	"goerp/internal/auth"
	models "goerp/internal/database"

	"github.com/gofiber/fiber/v2"

	jwtware "github.com/gofiber/contrib/jwt"
)

func main() {
	models.ConnectDatabase()
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to GOERP!")
	})

	app.Post("/login", auth.Login)

	// JWT Middleware
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte("secret")},
	}))

	app.Listen(":3000")
}
