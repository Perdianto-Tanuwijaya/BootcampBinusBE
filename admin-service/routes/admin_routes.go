package routes

import (
	"admin-service/controllers"
	// "admin-service/middleware"

	"github.com/gin-gonic/gin"
)

func AdminRoutes(r *gin.Engine) {
	admin := r.Group("/admin")
	// admin.Use(middleware.JWTAuthMiddleware(), middleware.CheckIsAdmin())
	{
		inventory := admin.Group("/inventory") 
		{
			inventory.POST("/flight/new", controllers.AddFlight)
			inventory.POST("/flight/edit/:flight_id", controllers.EditFlight)
			inventory.POST("/hotel/new", controllers.AddHotel)
			inventory.POST("/hotel/edit/:hotel_id", controllers.EditHotel)
		}
		admin.GET("/analytics", controllers.GetAnalytics)
	}
}