package handlers

import (
	"fmt"
	"temp_kbtg_backend/database"
	"temp_kbtg_backend/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// CreateTransferRequest represents the request body for creating a transfer
type CreateTransferRequest struct {
	FromUserID     uint   `json:"from_user_id"`
	ToUserID       uint   `json:"to_user_id"`
	Amount         int    `json:"amount"`
	Note           string `json:"note"`
	IdempotencyKey string `json:"idempotency_key"`
}

// GetTransfers returns all transfers with optional filtering
func GetTransfers(c *fiber.Ctx) error {
	var transfers []models.Transfer

	query := database.DB.Preload("FromUser").Preload("ToUser")

	// Optional filters
	if userID := c.Query("user_id"); userID != "" {
		query = query.Where("from_user_id = ? OR to_user_id = ?", userID, userID)
	}
	if status := c.Query("status"); status != "" {
		query = query.Where("status = ?", status)
	}

	if err := query.Order("created_at DESC").Find(&transfers).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to fetch transfers",
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    transfers,
	})
}

// GetTransfer returns a single transfer by idempotency_key
func GetTransfer(c *fiber.Ctx) error {
	idempotencyKey := c.Params("id") // Using idempotency_key as ID
	var transfer models.Transfer

	if err := database.DB.Preload("FromUser").Preload("ToUser").
		Where("idempotency_key = ?", idempotencyKey).
		First(&transfer).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "Transfer not found",
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    transfer,
	})
}

// CreateTransfer creates a new point transfer
func CreateTransfer(c *fiber.Ctx) error {
	req := new(CreateTransferRequest)

	if err := c.BodyParser(req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Validate required fields
	if req.FromUserID == 0 || req.ToUserID == 0 || req.Amount <= 0 || req.IdempotencyKey == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "Missing required fields or invalid amount",
		})
	}

	// Check if users are the same
	if req.FromUserID == req.ToUserID {
		return c.Status(400).JSON(fiber.Map{
			"error": "Cannot transfer to the same user",
		})
	}

	// Check for duplicate idempotency key (idempotent request)
	var existingTransfer models.Transfer
	if err := database.DB.Where("idempotency_key = ?", req.IdempotencyKey).First(&existingTransfer).Error; err == nil {
		// Return existing transfer
		return c.Status(200).JSON(fiber.Map{
			"success": true,
			"data":    existingTransfer,
			"message": "Transfer already exists (idempotent)",
		})
	}

	// Start transaction
	tx := database.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Verify both users exist
	var fromUser, toUser models.User
	if err := tx.First(&fromUser, req.FromUserID).Error; err != nil {
		tx.Rollback()
		return c.Status(404).JSON(fiber.Map{
			"error": "From user not found",
		})
	}
	if err := tx.First(&toUser, req.ToUserID).Error; err != nil {
		tx.Rollback()
		return c.Status(404).JSON(fiber.Map{
			"error": "To user not found",
		})
	}

	// Check if from_user has sufficient balance
	if fromUser.Balance < req.Amount {
		tx.Rollback()
		return c.Status(400).JSON(fiber.Map{
			"error": "Insufficient balance",
		})
	}

	// Create transfer record with status "processing"
	transfer := models.Transfer{
		FromUserID:     req.FromUserID,
		ToUserID:       req.ToUserID,
		Amount:         req.Amount,
		Status:         "processing",
		Note:           req.Note,
		IdempotencyKey: req.IdempotencyKey,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	if err := tx.Create(&transfer).Error; err != nil {
		tx.Rollback()
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to create transfer",
		})
	}

	// Deduct from sender
	fromUser.Balance -= req.Amount
	if err := tx.Save(&fromUser).Error; err != nil {
		tx.Rollback()
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to update sender balance",
		})
	}

	// Add to receiver
	toUser.Balance += req.Amount
	if err := tx.Save(&toUser).Error; err != nil {
		tx.Rollback()
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to update receiver balance",
		})
	}

	// Create ledger entries
	// 1. Deduct from sender
	ledgerOut := models.PointLedger{
		UserID:       req.FromUserID,
		Change:       -req.Amount,
		BalanceAfter: fromUser.Balance,
		EventType:    "transfer_out",
		TransferID:   &transfer.ID,
		Reference:    req.IdempotencyKey,
		CreatedAt:    time.Now(),
	}
	if err := tx.Create(&ledgerOut).Error; err != nil {
		tx.Rollback()
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to create ledger entry for sender",
		})
	}

	// 2. Add to receiver
	ledgerIn := models.PointLedger{
		UserID:       req.ToUserID,
		Change:       req.Amount,
		BalanceAfter: toUser.Balance,
		EventType:    "transfer_in",
		TransferID:   &transfer.ID,
		Reference:    req.IdempotencyKey,
		CreatedAt:    time.Now(),
	}
	if err := tx.Create(&ledgerIn).Error; err != nil {
		tx.Rollback()
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to create ledger entry for receiver",
		})
	}

	// Update transfer status to "completed"
	completedAt := time.Now()
	transfer.Status = "completed"
	transfer.CompletedAt = &completedAt
	transfer.UpdatedAt = time.Now()
	if err := tx.Save(&transfer).Error; err != nil {
		tx.Rollback()
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to complete transfer",
		})
	}

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to commit transaction",
		})
	}

	// Load relations for response
	database.DB.Preload("FromUser").Preload("ToUser").First(&transfer, transfer.ID)

	return c.Status(201).JSON(fiber.Map{
		"success": true,
		"data":    transfer,
	})
}

// CancelTransfer cancels a pending or processing transfer
func CancelTransfer(c *fiber.Ctx) error {
	idempotencyKey := c.Params("id")
	var transfer models.Transfer

	// Start transaction
	tx := database.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Where("idempotency_key = ?", idempotencyKey).First(&transfer).Error; err != nil {
		tx.Rollback()
		return c.Status(404).JSON(fiber.Map{
			"error": "Transfer not found",
		})
	}

	// Only allow cancellation of pending or processing transfers
	if transfer.Status != "pending" && transfer.Status != "processing" {
		tx.Rollback()
		return c.Status(400).JSON(fiber.Map{
			"error": fmt.Sprintf("Cannot cancel transfer with status: %s", transfer.Status),
		})
	}

	// If transfer was processing, need to reverse the points
	if transfer.Status == "processing" {
		var fromUser, toUser models.User
		if err := tx.First(&fromUser, transfer.FromUserID).Error; err != nil {
			tx.Rollback()
			return c.Status(500).JSON(fiber.Map{
				"error": "Failed to find sender",
			})
		}
		if err := tx.First(&toUser, transfer.ToUserID).Error; err != nil {
			tx.Rollback()
			return c.Status(500).JSON(fiber.Map{
				"error": "Failed to find receiver",
			})
		}

		// Reverse the transfer
		fromUser.Balance += transfer.Amount
		toUser.Balance -= transfer.Amount

		if err := tx.Save(&fromUser).Error; err != nil {
			tx.Rollback()
			return c.Status(500).JSON(fiber.Map{
				"error": "Failed to update sender balance",
			})
		}
		if err := tx.Save(&toUser).Error; err != nil {
			tx.Rollback()
			return c.Status(500).JSON(fiber.Map{
				"error": "Failed to update receiver balance",
			})
		}
	}

	// Update transfer status
	transfer.Status = "cancelled"
	transfer.UpdatedAt = time.Now()
	if err := tx.Save(&transfer).Error; err != nil {
		tx.Rollback()
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to cancel transfer",
		})
	}

	if err := tx.Commit().Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to commit transaction",
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    transfer,
		"message": "Transfer cancelled successfully",
	})
}

// GetUserLedger returns the point ledger for a user
func GetUserLedger(c *fiber.Ctx) error {
	userID := c.Params("user_id")
	var ledgers []models.PointLedger

	query := database.DB.Preload("Transfer").Where("user_id = ?", userID)

	// Optional filters
	if eventType := c.Query("event_type"); eventType != "" {
		query = query.Where("event_type = ?", eventType)
	}

	if err := query.Order("created_at DESC").Find(&ledgers).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to fetch ledger",
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    ledgers,
	})
}
