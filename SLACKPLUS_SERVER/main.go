package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	db "github.com/hktrib/SlackPlus/db/sqlc"
	handler "github.com/hktrib/SlackPlus/routes"
	"github.com/hktrib/SlackPlus/util"
)

func setupRoutes(app *fiber.App, handlers *handler.Handlers) {

	app.Post("/register", handlers.RegisterHandler)

}

func main() {

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	handlers := handler.NewHandlers(store)

	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		AppName:       "Orbit",
	})
	app.Use(cors.New())

	setupRoutes(app, handlers)

	app.Listen(fmt.Sprintf(":%s", config.ServerAddress))

}
