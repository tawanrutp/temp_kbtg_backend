package main

import (
	"log"
	"temp_kbtg_backend/database"
	"temp_kbtg_backend/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// Initialize database
	database.InitDatabase()

	// Create a new Fiber instance
	app := fiber.New(fiber.Config{
		AppName: "KBTG Backend API",
	})

	// Add middleware
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

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
			"endpoints": fiber.Map{
				"customers": "/api/v1/customers",
				"orders":    "/api/v1/orders",
			},
		})
	})

	// Setup API routes
	routes.SetupRoutes(app)

	// Start server on port 3000
	log.Printf("Server starting on http://localhost:3000")
	log.Printf("API endpoints available at http://localhost:3000/api/v1")
	if err := app.Listen(":3000"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
