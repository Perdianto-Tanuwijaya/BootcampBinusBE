package main

import (
	"payment-service/config"
	"payment-service/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	config.ConnectDatabase()
	routes.PaymentRoutes(r)
	r.Run(":8082")
}
