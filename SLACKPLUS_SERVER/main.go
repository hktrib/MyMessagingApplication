package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/hktrib/SlackPlus/initializers"
	Auth "github.com/hktrib/SlackPlus/routes"
)

func init() {
	initializers.LoadEnvVariables()
}

func main() {

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("HOST"), os.Getenv("DP_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		AppName:       "Orbit",
	})
	app.Use(cors.New())

	Auth.Auth(app, db)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// app.Post("/Register", func(c *fiber.Ctx) error {
	//

	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.SendString("Hare Krsna....Aksara Nitai Dasa speaking here.")
	// // })
	app.Listen(fmt.Sprintf(":%s", os.Getenv("SERVER_PORT")))

}
