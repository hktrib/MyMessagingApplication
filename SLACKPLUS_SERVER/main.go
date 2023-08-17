package main

import (
	"github.com/gofiber/fiber/v2"
	Auth "github.com/hktrib/SlackPlus/routes"
)

type CustomError struct {
	Code    int
	Message string
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
	//

	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.SendString("Hare Krsna....Aksara Nitai Dasa speaking here.")
	// })
	app.Listen(":8080")

}
