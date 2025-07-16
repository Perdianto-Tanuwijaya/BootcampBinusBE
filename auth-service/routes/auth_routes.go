package routes

import (
	"auth-service/controllers"
	"auth-service/middleware"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine) {
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	auth := r.Group("/user")
	auth.Use(middleware.JWTAuthMiddleware())
	{
		auth.PUT("/profile/edit", controllers.UpdateProfile)
	}
}
func AdminRoutes(r *gin.Engine) {
	admin := r.Group("/admin")
	admin.Use(middleware.JWTAuthMiddleware(), middleware.CheckIsAdmin())
	{
		admin.GET("/dashboard", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "Welcome, admin!"})
		})
	}
}
