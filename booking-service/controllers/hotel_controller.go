package controllers

import (
	"booking-service/config"
	"booking-service/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func SearchHotels(c *gin.Context) {
	location := c.Query("location")
	minPriceStr := c.Query("minPrice")
	maxPriceStr := c.Query("maxPrice")

	var hotels []models.Hotel
	query := config.DB

	if location != "" {
		query = query.Where("location ILIKE ?", "%"+location+"%")
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

	if err := query.Find(&hotels).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search hotels"})
		return
	}

	if len(hotels) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No hotels found for the given filters"})
		return
	}

	c.JSON(http.StatusOK, hotels)
}
