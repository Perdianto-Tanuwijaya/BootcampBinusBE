package routes

import (
	"booking-service/controllers"
	"booking-service/middleware"

	"github.com/gin-gonic/gin"
)

func BookingRoutes(r *gin.Engine) {
	booking := r.Group("/booking")
	booking.Use(middleware.JWTAuthMiddleware())
	{
		booking.POST("/", controllers.CreateBooking)
		// booking.PUT("/:bookingId/confirm", controllers.ConfirmBooking)
		booking.GET("/:bookingId", controllers.GetBookingByID)
		booking.GET("/:bookingId/status", controllers.GetBookingStatusByID)
	}
}
