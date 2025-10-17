package database

import (
	"log"
	"temp_kbtg_backend/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDatabase() {
	var err error
	
	// Connect to SQLite database
	DB, err = gorm.Open(sqlite.Open("kbtg.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	log.Println("Database connected successfully")

	// Auto migrate all models
	err = DB.AutoMigrate(
		&models.Customer{},
		&models.DeliveryAddress{},
		&models.Order{},
		&models.LineItem{},
		&models.User{},
		&models.Transfer{},
		&models.PointLedger{},
	)

	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	log.Println("Database migration completed")
}
