package routes

import (
	"temp_kbtg_backend/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// API v1 group
	api := app.Group("/api/v1")

	// Customer routes
	customers := api.Group("/customers")
	customers.Get("/", handlers.GetCustomers)
	customers.Get("/:id", handlers.GetCustomer)
	customers.Post("/", handlers.CreateCustomer)
	customers.Put("/:id", handlers.UpdateCustomer)
	customers.Delete("/:id", handlers.DeleteCustomer)

	// Order routes
	orders := api.Group("/orders")
	orders.Get("/", handlers.GetOrders)
	orders.Get("/:id", handlers.GetOrder)
	orders.Post("/", handlers.CreateOrder)
	orders.Put("/:id", handlers.UpdateOrder)
	orders.Delete("/:id", handlers.DeleteOrder)
}
