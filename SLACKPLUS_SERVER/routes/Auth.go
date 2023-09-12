package Auth

import (
	"database/sql"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	Username string `json:"username"`
	Email    string `json:"emailAddress"`
	Password string `json:"password"`
}

func Auth(app *fiber.App, db *sql.DB) {

	app.Post("/register", registerHandler)
	// fmt.Println("Hare Krsna Aksara, we auth'ing now")
}

func registerHandler(c *fiber.Ctx) error {

	user := User{}
	fmt.Println("Entered Post method")

	if c.Method() != fiber.MethodPost {
		c.Status(fiber.StatusMethodNotAllowed)
		return fiber.ErrNotAcceptable
	}

	if err := c.BodyParser(&user); err != nil {
		c.Status(fiber.StatusBadRequest)
		return fiber.ErrBadRequest
	}

	fmt.Println("Creating User:", user)
	return c.SendString("Everything Good here")
}
