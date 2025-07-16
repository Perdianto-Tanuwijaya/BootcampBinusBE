package models

import "time"

type Booking struct {
	BookingID   string `gorm:"primaryKey"`
	UserID      string
	ItemID      string
	BookingType string
	Status      string
	Amount      int
	ExpiresAt   time.Time
	CreatedAt   time.Time
}
