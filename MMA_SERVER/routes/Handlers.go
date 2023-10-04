package handler

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	db "github.com/hktrib/MyMessagingApplication/db/sqlc"
	"github.com/hktrib/MyMessagingApplication/util"
	emailverification "github.com/hktrib/MyMessagingApplication/util/EmailVerification"
)

type Handlers struct {
	Store  *db.Store
	Config *util.Config
}

func NewHandlers(store *db.Store, config *util.Config) *Handlers {
	return &Handlers{
		Store:  store,
		Config: config,
	}
}

type User struct {
	Username string `json:"username"`
	Email    string `json:"emailAddress"`
	Password string `json:"password"`
}

// PriorRegistrationCheck -> returns Error
// Purpose: For live checks of whether a username is valid
// Status: Under development
// Issues: Queries hanging if username is typed quickly because of multiple in succession
// (perhaps requests are dropped because function isn't being reached)
func (h *Handlers) PriorRegistrationCheck(c *fiber.Ctx) error {
	fmt.Println("Entered Check")
	start := time.Now()
	if emailValue := c.Query("email"); emailValue != "" {
		if exists, _ := h.Store.SearchUserByEmail(c.Context(), emailValue); exists == false {
			end := time.Now()
			fmt.Printf("%v Email-Val, Exists = %v, Elapsed: %v\n", emailValue, exists, end.Sub(start))
			return c.Status(fiber.StatusContinue).JSON(fiber.Map{
				"isUsed": true,
			})
		} else {
			end := time.Now()
			fmt.Printf("%v Email-Val, Exists = %v, Elapsed: %v\n", emailValue, exists, end.Sub(start))
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{
				"isUsed": false,
			})
		}
	}

	if usernameValue := c.Query("username"); usernameValue != "" {
		exists, err := h.Store.SearchUserByUsername(c.Context(), usernameValue)
		if err != nil {
			return c.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{
				"message": "Error: Internal Server failuire",
			})
		}
		if !exists {
			end := time.Now()
			fmt.Printf("%v Username-val, Exists = %v, Elapsed: %v\n", usernameValue, exists, end.Sub(start))
			return c.Status(fiber.StatusContinue).JSON(fiber.Map{
				"isAvailable": true,
			})
		} else {
			end := time.Now()
			fmt.Printf("%v Username-val, Exists = %v, Elapsed: %v\n", usernameValue, exists, end.Sub(start))
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{
				"isAvailable": false,
			})
		}
	}

	end := time.Now()
	fmt.Printf("Error: Misdirected Get Request.....Elapsed: %v\n", end.Sub(start))
	return c.Status(fiber.StatusMisdirectedRequest).JSON(fiber.Map{
		"message": "Error: Misdirected Get Request",
	})
}

func (h *Handlers) RegisterUser(c *fiber.Ctx) error {
	potential_user := User{}
	fmt.Println("Entered Post method")

	// Preliminary Checks
	if c.Method() != fiber.MethodPost {
		c.Status(fiber.StatusMethodNotAllowed)
		return fiber.ErrNotAcceptable
	}

	if err := c.BodyParser(&potential_user); err != nil {
		c.Status(fiber.StatusBadRequest)
		return fiber.ErrBadRequest
	}

	// Checking for prior existence of user
	start := time.Now()
	email_alr_exists, err := h.Store.SearchUserByEmail(c.Context(), potential_user.Email)
	if err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{
			"message": fmt.Sprintf("Error: Internal Server failuire: %v", err),
		})
	}

	if email_alr_exists {
		end := time.Now()
		fmt.Printf("%v Email-Val, Exists = %v, Elapsed: %v\n", potential_user.Email, email_alr_exists, end.Sub(start))
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"isUsed": false,
		})
	} else {
		end := time.Now()
		fmt.Printf("%v Email-Val, Exists = %v, Elapsed: %v\n", potential_user.Email, email_alr_exists, end.Sub(start))
	}
	username_alr_exists, err := h.Store.SearchUserByUsername(c.Context(), potential_user.Username)
	if err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{
			"message": "Error: Internal Server failuire",
		})
	}
	if username_alr_exists {
		end := time.Now()
		fmt.Printf("%v Username-val, Exists = %v, Elapsed: %v\n", potential_user, username_alr_exists, end.Sub(start))
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"isAvailable": false,
		})
	} else {
		end := time.Now()
		fmt.Printf("%v Username-val, Exists = %v, Elapsed: %v\n", potential_user.Username, username_alr_exists, end.Sub(start))
	}

	// Encryption
	potential_user.Password = util.Encrypt(potential_user.Password)

	user, err := h.Store.CreateUser(c.Context(), db.CreateUserParams{potential_user.Username, potential_user.Password, potential_user.Email})
	if err != nil {
		return fiber.ErrInternalServerError
	}

	// Creating Verification Email record in database
	ve_record, err := h.Store.CreateVerifyEmailsRecord(c.Context(), db.CreateVerifyEmailsRecordParams{user.Username, user.Email, util.RandomString()})
	if err != nil {
		return fiber.ErrInternalServerError
	}

	fmt.Println("Created User:", user)
	fmt.Println("Created VE_Record", ve_record)

	// Send verification email using SMTP
	emailverification.SendMail(h.Config, &user.Email, &ve_record.Username, &ve_record.SecretCode)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User Created Successfully",
	})
}
