package models

import "time"

type Customer struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"size:100;not null" json:"name"`
	Email     string    `gorm:"size:100;unique;not null" json:"email"`
	Phone     string    `gorm:"size:20" json:"phone"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type DeliveryAddress struct {
	ID         uint   `gorm:"primaryKey" json:"id"`
	CustomerID uint   `gorm:"not null" json:"customer_id"`
	Address    string `gorm:"size:255;not null" json:"address"`
	City       string `gorm:"size:100" json:"city"`
	PostalCode string `gorm:"size:20" json:"postal_code"`
	Country    string `gorm:"size:100" json:"country"`
	IsDefault  bool   `gorm:"default:false" json:"is_default"`
}

type Order struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	CustomerID uint      `gorm:"not null" json:"customer_id"`
	OrderDate  time.Time `gorm:"not null" json:"order_date"`
	Status     string    `gorm:"size:50;default:'pending'" json:"status"`
	TotalPrice float64   `gorm:"type:decimal(10,2);default:0" json:"total_price"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type LineItem struct {
	ID          uint    `gorm:"primaryKey" json:"id"`
	OrderID     uint    `gorm:"not null" json:"order_id"`
	ProductName string  `gorm:"size:100;not null" json:"product_name"`
	Quantity    int     `gorm:"not null" json:"quantity"`
	UnitPrice   float64 `gorm:"type:decimal(10,2);not null" json:"unit_price"`
	TotalPrice  float64 `gorm:"type:decimal(10,2)" json:"total_price"`
}
