package models

type Payment struct {
	PaymentID     string `gorm:"primaryKey"`
	BookingID     string
	PaymentMethod string
	Amount        int
	Status        string
}
