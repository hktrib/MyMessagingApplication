package Auth

import "github.com/gofiber/fiber/v2"

func Auth(app *fiber.App) {

	app.Post("/Register", registerHandler)
}

func registerHandler(c *fiber.Ctx) error {
	return fiber.ErrForbidden
}
