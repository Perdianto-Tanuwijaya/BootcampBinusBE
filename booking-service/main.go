package main

import (
	"booking-service/config"
	"booking-service/models"
	"booking-service/routes"
	"booking-service/seed"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	config.ConnectDatabase()
	seed.Load()

	routes.BookingRoutes(r)
	routes.HotelRoutes(r)
	routes.FlightRoutes(r)

	// Auto-cancel expired bookings
	go func() {
		for {
			time.Sleep(1 * time.Minute)

			result := config.DB.Model(&models.Booking{}).
				Where("status = ? AND expires_at < ?", "pending", time.Now()).
				Update("status", "cancelled")

			if result.Error != nil {
				fmt.Println("Error auto-cancelling:", result.Error)
			} else if result.RowsAffected > 0 {
				fmt.Printf("Auto-cancelled %d expired bookings\n", result.RowsAffected)
			}
		}
	}()

	r.Run(":8081")
}
