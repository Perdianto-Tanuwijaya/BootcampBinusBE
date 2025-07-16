package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"payment-service/config"
	"payment-service/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PaymentInput struct {
	BookingID     string `json:"bookingId" binding:"required"`
	PaymentMethod string `json:"paymentMethod" binding:"required"`
	Amount        int    `json:"amount" binding:"required"`
}

type BookingResponse struct {
	BookingID   string `json:"bookingId"`
	UserID      string `json:"userId"`
	ItemID      string `json:"itemId"`
	BookingType string `json:"bookingType"`
	Status      string `json:"status"`
	Amount      int    `json:"amount"`
}

func CreatePayment(c *gin.Context) {
	var input PaymentInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Ambil data booking dari booking-service
	bookingServiceURL := "http://localhost:8081/booking/" + input.BookingID
	req, err := http.NewRequest(http.MethodGet, bookingServiceURL, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request to booking-service"})
		return
	}
	req.Header.Set("Authorization", c.GetHeader("Authorization"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != http.StatusOK {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to retrieve booking"})
		return
	}
	defer resp.Body.Close()

	var booking BookingResponse
	if err := json.NewDecoder(resp.Body).Decode(&booking); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse booking data"})
		return
	}

	// Validasi jumlah pembayaran
	if booking.Amount != input.Amount {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Payment amount does not match booking amount"})
		return
	}

	// Simpan payment
	payment := models.Payment{
		PaymentID:     uuid.New().String(),
		BookingID:     input.BookingID,
		PaymentMethod: input.PaymentMethod,
		Amount:        input.Amount,
		Status:        "success",
	}

	if err := config.DB.Create(&payment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process payment"})
		return
	}

	// Update status booking ke confirmed
	confirmURL := "http://localhost:8081/booking/" + input.BookingID + "/confirm"
	confirmReq, _ := http.NewRequest(http.MethodPut, confirmURL, nil)
	confirmReq.Header.Set("Authorization", c.GetHeader("Authorization"))
	confirmResp, err := client.Do(confirmReq)
	if err != nil {
		fmt.Println("Error updating booking:", err)
	} else {
		defer confirmResp.Body.Close()
		fmt.Println("Booking confirmed:", confirmResp.Status)
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":    "Payment successful",
		"payment_id": payment.PaymentID,
	})
}
