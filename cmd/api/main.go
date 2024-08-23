package main

import (
	"fmt"
	"os"

	"goerp/internal/auth"
	models "goerp/internal/database"

	"github.com/gofiber/fiber/v2"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	models.ConnectDatabase()
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to GOERP!")
	})

	app.Post("/signup", auth.Signup)
	app.Post("/login", auth.Login)

	jwt_key := os.Getenv("JWT_KEY")

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(jwt_key)},
	}))

	port := os.Getenv("APP_PORT")
	app.Listen(fmt.Sprintf(":%s", port))
}
