package Handler

import (
	"crypto/sha256"
	"fmt"

	"github.com/gofiber/fiber/v2"
	db "github.com/hktrib/SlackPlus/db/sqlc"
)

type Handlers struct {
	Store *db.Store
}

func NewHandlers(store *db.Store) *Handlers {
	return &Handlers{
		Store: store,
	}
}

type User struct {
	Username string `json:"username"`
	Email    string `json:"emailAddress"`
	Password string `json:"password"`
}

func (h *Handlers) RegisterHandler(c *fiber.Ctx) error {
	userReq := User{}
	fmt.Println("Entered Post method")

	// Preliminary Checks
	if c.Method() != fiber.MethodPost {
		c.Status(fiber.StatusMethodNotAllowed)
		return fiber.ErrNotAcceptable
	}

	if err := c.BodyParser(&userReq); err != nil {
		c.Status(fiber.StatusBadRequest)
		return fiber.ErrBadRequest
	}

	encrypter := sha256.New()

	encrypter.Write([]byte(userReq.Password))

	encryptedBS := encrypter.Sum(nil)

	user, err := h.Store.CreateUser(c.Context(), db.CreateUserParams{userReq.Username, fmt.Sprintf("%x", encryptedBS), userReq.Email})

	fmt.Println(fmt.Sprintf("%x", encryptedBS))

	if err != nil {
		return fiber.ErrInternalServerError
	}

	fmt.Println("Creating User:", user)
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User Created Successfully",
	})
}
