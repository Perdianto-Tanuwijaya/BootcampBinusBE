package controllers

import (
	"booking-service/config"
	"booking-service/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type BookingInput struct {
	ItemID      string `json:"itemId" binding:"required"`
	BookingType string `json:"bookingType" binding:"required,oneof=flight hotel"`
}

func CreateBooking(c *gin.Context) {
	var input struct {
		ItemID      string `json:"itemId" binding:"required"`
		BookingType string `json:"bookingType" binding:"required,oneof=flight hotel"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId := c.MustGet("userId").(string)

	var price int
	if input.BookingType == "hotel" {
		// Cek dari tabel hotel
		var hotel models.Hotel
		if err := config.DB.First(&hotel, "hotel_id = ?", input.ItemID).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Hotel not found"})
			return
		}
		price = hotel.Price
	} else {
		var flight models.Flight
		if err := config.DB.First(&flight, "flight_id = ?", input.ItemID).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Flight not found"})
			return
		}
		price = flight.Price
	}

	booking := models.Booking{
		BookingID:   uuid.New().String(),
		UserID:      userId,
		ItemID:      input.ItemID,
		BookingType: input.BookingType,
		Status:      "pending",
		Amount:      price,
		ExpiresAt:   time.Now().Add(15 * time.Minute),
		CreatedAt:   time.Now(),
	}

	if err := config.DB.Create(&booking).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create booking"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"bookingId": booking.BookingID, "amount": booking.Amount, "expiresAt": booking.ExpiresAt})
}

func ConfirmBooking(c *gin.Context) {
	bookingID := c.Param("bookingId")

	var booking models.Booking
	if err := config.DB.First(&booking, "booking_id = ?", bookingID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Booking not found"})
		return
	}

	booking.Status = "confirmed"

	if err := config.DB.Save(&booking).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update booking status"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Booking confirmed"})
}

func GetBookingByID(c *gin.Context) {
	bookingID := c.Param("bookingId")
	var booking models.Booking

	if err := config.DB.Where("booking_id = ?", bookingID).First(&booking).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Booking not found"})
		return
	}

	c.JSON(http.StatusOK, booking)
}

func GetBookingStatusByID(c *gin.Context) {
	bookingId := c.Param("bookingId")

	var booking models.Booking
	if err := config.DB.Where("booking_id = ?", bookingId).First(&booking).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Booking not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"bookingId": booking.BookingID,
		"status":    booking.Status,
	})
}
