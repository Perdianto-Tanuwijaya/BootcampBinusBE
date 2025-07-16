package routes

import (
	"booking-service/controllers"

	"github.com/gin-gonic/gin"
)

func FlightRoutes(r *gin.Engine) {
	flight := r.Group("/flights")
	flight.GET("/search", controllers.SearchFlights)
}
