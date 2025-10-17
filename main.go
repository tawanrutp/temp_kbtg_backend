package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// Create a new Fiber instance
	app := fiber.New(fiber.Config{
		AppName: "KBTG Backend API",
	})

	// Add logger middleware
	app.Use(logger.New())

	// Hello World route
	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Hello World",
			"status":  "success",
		})
	})

	// Root route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Welcome to KBTG Backend API",
			"version": "1.0.0",
		})
	})

	// Start server on port 3000
	log.Printf("Server starting on http://localhost:3000")
	if err := app.Listen(":3000"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
