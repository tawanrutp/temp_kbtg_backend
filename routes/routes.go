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

	// User routes
	users := api.Group("/users")
	users.Get("/", handlers.GetUsers)
	users.Get("/:id", handlers.GetUser)
	users.Post("/", handlers.CreateUser)
	users.Put("/:id", handlers.UpdateUser)
	users.Delete("/:id", handlers.DeleteUser)
	users.Get("/:id/balance", handlers.GetUserBalance)

	// Transfer routes
	transfers := api.Group("/transfers")
	transfers.Get("/", handlers.GetTransfers)
	transfers.Get("/:id", handlers.GetTransfer)
	transfers.Post("/", handlers.CreateTransfer)
	transfers.Delete("/:id", handlers.CancelTransfer)

	// Point Ledger routes
	api.Get("/users/:user_id/ledger", handlers.GetUserLedger)
}
