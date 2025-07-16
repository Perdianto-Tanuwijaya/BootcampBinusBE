package seed

import (
	"booking-service/config"
	"booking-service/models"
	"time"
)

func Load() {
	// Seed Hotel
	config.DB.Create(&models.Hotel{
		HotelID:      "hotel-001",
		Name:         "Grand Indonesia Hotel",
		Price:        500000,
		Availability: 10,
		Location:     "Jakarta",
	})

	// Seed Flight
	config.DB.Create(&models.Flight{
		FlightID:      "flight-001",
		Name:          "Garuda Indonesia",
		FlightNumber:  "GA123",
		Price:         750000,
		DepartureTime: time.Now().Add(2 * time.Hour),
		ArrivalTime:   time.Now().Add(5 * time.Hour),
		Availability:  20,
		LocationFrom:  "Jakarta",
		LocationTo:    "Bali",
	})
}
