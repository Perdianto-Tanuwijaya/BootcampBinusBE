package config

import (
	"booking-service/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "host=localhost user=postgres password=perdianto26 dbname=travelokaclone port=5432 sslmode=disable"
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Migration
	err = db.AutoMigrate(
		&models.Booking{},
		&models.Hotel{},
		&models.Flight{},
	)
	if err != nil {
		log.Fatal("Migration failed:", err)
	}

	DB = db
}
