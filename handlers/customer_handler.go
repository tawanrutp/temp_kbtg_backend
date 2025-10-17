package handlers

import (
	"temp_kbtg_backend/database"
	"temp_kbtg_backend/models"

	"github.com/gofiber/fiber/v2"
)

// Get all customers
func GetCustomers(c *fiber.Ctx) error {
	var customers []models.Customer
	
	if err := database.DB.Find(&customers).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to fetch customers",
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    customers,
	})
}

// Get customer by ID
func GetCustomer(c *fiber.Ctx) error {
	id := c.Params("id")
	var customer models.Customer

	if err := database.DB.First(&customer, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "Customer not found",
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    customer,
	})
}

// Create new customer
func CreateCustomer(c *fiber.Ctx) error {
	customer := new(models.Customer)

	if err := c.BodyParser(customer); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := database.DB.Create(&customer).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to create customer",
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"success": true,
		"data":    customer,
	})
}

// Update customer
func UpdateCustomer(c *fiber.Ctx) error {
	id := c.Params("id")
	var customer models.Customer

	if err := database.DB.First(&customer, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "Customer not found",
		})
	}

	if err := c.BodyParser(&customer); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := database.DB.Save(&customer).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to update customer",
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    customer,
	})
}

// Delete customer
func DeleteCustomer(c *fiber.Ctx) error {
	id := c.Params("id")
	var customer models.Customer

	if err := database.DB.First(&customer, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "Customer not found",
		})
	}

	if err := database.DB.Delete(&customer).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to delete customer",
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Customer deleted successfully",
	})
}
