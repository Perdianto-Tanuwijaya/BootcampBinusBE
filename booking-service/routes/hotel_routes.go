package routes

import (
	"booking-service/controllers"

	"github.com/gin-gonic/gin"
)

func HotelRoutes(r *gin.Engine) {
	hotel := r.Group("/hotels")
	hotel.GET("/search", controllers.SearchHotels)
}
