package main

import (
	"admin-service/config"
	"admin-service/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	config.ConnectDatabase()
	routes.AdminRoutes(r)
	r.Run(":8080")
}