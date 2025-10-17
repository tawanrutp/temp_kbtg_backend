package handlers

import (
	"temp_kbtg_backend/database"
	"temp_kbtg_backend/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

// Get all orders
func GetOrders(c *fiber.Ctx) error {
	var orders []models.Order
	
	if err := database.DB.Find(&orders).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to fetch orders",
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    orders,
	})
}

// Get order by ID
func GetOrder(c *fiber.Ctx) error {
	id := c.Params("id")
	var order models.Order

	if err := database.DB.First(&order, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "Order not found",
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    order,
	})
}

// Create new order
func CreateOrder(c *fiber.Ctx) error {
	order := new(models.Order)

	if err := c.BodyParser(order); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Set order date to now if not provided
	if order.OrderDate.IsZero() {
		order.OrderDate = time.Now()
	}

	if err := database.DB.Create(&order).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to create order",
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"success": true,
		"data":    order,
	})
}

// Update order
func UpdateOrder(c *fiber.Ctx) error {
	id := c.Params("id")
	var order models.Order

	if err := database.DB.First(&order, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "Order not found",
		})
	}

	if err := c.BodyParser(&order); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := database.DB.Save(&order).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to update order",
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    order,
	})
}

// Delete order
func DeleteOrder(c *fiber.Ctx) error {
	id := c.Params("id")
	var order models.Order

	if err := database.DB.First(&order, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "Order not found",
		})
	}

	if err := database.DB.Delete(&order).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to delete order",
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Order deleted successfully",
	})
}
