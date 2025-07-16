package controllers

import (
	"booking-service/config"
	"booking-service/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func SearchFlights(c *gin.Context) {
	locationFrom := c.Query("from")
	locationTo := c.Query("to")
	minPriceStr := c.Query("minPrice")
	maxPriceStr := c.Query("maxPrice")

	var flights []models.Flight
	query := config.DB

	if locationFrom != "" {
		query = query.Where("location_from ILIKE ?", "%"+locationFrom+"%")
	}
	if locationTo != "" {
		query = query.Where("location_to ILIKE ?", "%"+locationTo+"%")
	}
	if minPriceStr != "" {
		if minPrice, err := strconv.Atoi(minPriceStr); err == nil {
			query = query.Where("price >= ?", minPrice)
		}
	}
	if maxPriceStr != "" {
		if maxPrice, err := strconv.Atoi(maxPriceStr); err == nil {
			query = query.Where("price <= ?", maxPrice)
		}
	}

	if err := query.Find(&flights).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search flights"})
		return
	}

	if len(flights) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No flights found for the given filters"})
		return
	}

	c.JSON(http.StatusOK, flights)
}
