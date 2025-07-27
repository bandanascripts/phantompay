package structs

import "time"

type PhantomPayWallet struct {
	WalletID  string    `gorm:"primaryKey;type:varchar(100)" json:"wallet_id"`
	UserID    string    `gorm:"uniqueIndex;type:varchar(100);not null" json:"user_id"` // One wallet per user
	Balance   float64   `gorm:"not null;default:0" json:"balance"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
