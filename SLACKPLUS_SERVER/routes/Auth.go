package Auth

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Auth(app *fiber.App) {

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

	fmt.Println(c.JSON(user))
	return fiber.ErrForbidden
}
