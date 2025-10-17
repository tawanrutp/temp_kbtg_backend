package models

import "time"

// User represents a user in the system who can send/receive points
type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"size:100;not null" json:"name"`
	Email     string    `gorm:"size:100;unique;not null" json:"email"`
	Balance   int       `gorm:"default:0;not null" json:"balance"` // Current point balance
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Transfer represents a point transfer between users
type Transfer struct {
	ID             uint       `gorm:"primaryKey" json:"id"`
	FromUserID     uint       `gorm:"not null;index:idx_transfers_from" json:"from_user_id"`
	ToUserID       uint       `gorm:"not null;index:idx_transfers_to" json:"to_user_id"`
	Amount         int        `gorm:"not null;check:amount > 0" json:"amount"`
	Status         string     `gorm:"size:20;not null;check:status IN ('pending','processing','completed','failed','cancelled','reversed')" json:"status"`
	Note           string     `gorm:"type:text" json:"note,omitempty"`
	IdempotencyKey string     `gorm:"size:255;not null;uniqueIndex" json:"idempotency_key"` // Used as ID in GET /transfers/{id}
	CreatedAt      time.Time  `gorm:"not null;index:idx_transfers_created" json:"created_at"`
	UpdatedAt      time.Time  `gorm:"not null" json:"updated_at"`
	CompletedAt    *time.Time `json:"completed_at,omitempty"`
	FailReason     string     `gorm:"type:text" json:"fail_reason,omitempty"`
	
	// Relations
	FromUser User `gorm:"foreignKey:FromUserID" json:"from_user,omitempty"`
	ToUser   User `gorm:"foreignKey:ToUserID" json:"to_user,omitempty"`
}

// PointLedger represents a point transaction log
type PointLedger struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	UserID       uint      `gorm:"not null;index:idx_ledger_user" json:"user_id"`
	Change       int       `gorm:"not null" json:"change"` // +receive / -send
	BalanceAfter int       `gorm:"not null" json:"balance_after"`
	EventType    string    `gorm:"size:20;not null;check:event_type IN ('transfer_out','transfer_in','adjust','earn','redeem')" json:"event_type"`
	TransferID   *uint     `gorm:"index:idx_ledger_transfer" json:"transfer_id,omitempty"` // Reference to transfers.id (internal ID)
	Reference    string    `gorm:"size:255" json:"reference,omitempty"`
	Metadata     string    `gorm:"type:text" json:"metadata,omitempty"` // JSON text
	CreatedAt    time.Time `gorm:"not null;index:idx_ledger_created" json:"created_at"`
	
	// Relations
	User     User      `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Transfer *Transfer `gorm:"foreignKey:TransferID" json:"transfer,omitempty"`
}
