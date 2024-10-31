package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
	"telegram_todo_bot/client"
	"telegram_todo_bot/config"
	"telegram_todo_bot/handlers"
)

func main() {
	// Init new fiber app
	app := fiber.New()

	// Init cors middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
	}))

	// Init telegram bot client
	bot := client.Init()
	handlers.Init(bot)

	// Start app
	log.Fatal(app.Listen(":" + config.Config("PORT")))
}
