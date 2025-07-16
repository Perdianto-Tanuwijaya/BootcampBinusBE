package main

import (
	"auth-service/config"
	"auth-service/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	config.ConnectDatabase()
	routes.AuthRoutes(r)
	routes.AdminRoutes(r)
	r.Run(":8080")
}
