package main

import (
	"fablab-projects/controllers"
	"fablab-projects/database"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		c.Send([]byte("Welcome to the vjbg_fablab projects API"));

		return nil
	})

	app.Get("/login", controllers.Login)
	app.Post("/register", controllers.Register)
}

func main() {
	app := fiber.New()
	database.ConnectDB()

	app.Use(logger.New())
	setupRoutes(app)

	log.Fatal(app.Listen(":"+os.Getenv("PORT")))
}