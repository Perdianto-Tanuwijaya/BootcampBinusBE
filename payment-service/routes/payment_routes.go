package routes

import (
	"payment-service/controllers"
	"payment-service/middleware"

	"github.com/gin-gonic/gin"
)

func PaymentRoutes(r *gin.Engine) {
	pay := r.Group("/payment")
	pay.Use(middleware.JWTAuthMiddleware())
	{
		pay.POST("/", controllers.CreatePayment)
	}
}
