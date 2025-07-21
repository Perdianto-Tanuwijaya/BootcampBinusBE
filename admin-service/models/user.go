package models

type User struct {
	UserID   string `gorm:"primaryKey"`
	Name     string
	Email    string `gorm:"unique"`
	Password string
	IsAdmin  bool
}
