package main

import (
	"github.com/gofiber/fiber/v2"
	Auth "github.com/hktrib/SlackPlus/routes"
)

type CustomError struct {
	Code    int
	Message string
}

var user struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func main() {
	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		AppName:       "Orbit",
	})

	Auth.Auth(app)

	// app.Post("/Register", func(c *fiber.Ctx) error {
	// 	if c.Method() != fiber.MethodPost {
	// 		c.Status(fiber.StatusMethodNotAllowed)
	// 		return fiber.ErrNotAcceptable
	// 	}

	// 	if err := c.BodyParser(&user); err != nil {
	// 		c.Status(fiber.StatusBadRequest)
	// 		return fiber.ErrBadRequest
	// 	}

	// 	fmt.Println("Creating User:", user)

	// 	c.JSON(user)
	// })

	app.Listen(":5173")

}
