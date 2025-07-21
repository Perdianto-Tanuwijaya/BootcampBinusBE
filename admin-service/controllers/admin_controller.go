package controllers

import (
	"admin-service/config"
	"admin-service/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type FlightInput struct {
	Name          string    `json:"name" binding:"required"`
	Price         int       `json:"price" binding:"required"`
	DepartureTime time.Time `json:"departureTime" binding:"required"`
	ArrivalTime   time.Time `json:"arrivalTime" binding:"required"`
	Availability  int       `json:"availability" binding:"required"`
	LocationFrom  string    `json:"locationFrom" binding:"required"`
	LocationTo    string    `json:"locationTo" binding:"required"`
}


type HotelInput struct {
	Name         string `json:"name" binding:"required"`
	Price        int    `json:"price" binding:"required"`
	Availability int    `json:"availability" binding:"required"`
	Location     string `json:"location" binding:"required"`
}


type FlightBooking struct {
	ID     string
	Status string
}

type HotelBooking struct {
	ID     string
	Status string
}

type User struct {
	ID string
}


func AddFlight(c *gin.Context) {
	var input FlightInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	

	flight := models.Flight {
		FlightID: "flight-" + uuid.New().String(),
		Name: input.Name,
		Price: input.Price,
		DepartureTime: input.DepartureTime,
		ArrivalTime: input.ArrivalTime,
		Availability: input.Availability,
		LocationFrom: input.LocationFrom,
		LocationTo: input.LocationTo,
	}

	if err := config.DB.Create(&flight).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add flight"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Flight added successfully"})
}


func EditFlight(c *gin.Context) {
	flightID := c.Param("flight_id")

	var input FlightInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//check exist
	var flight models.Flight
	if err := config.DB.First(&flight, "flight_id = ?", flightID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Flight not found"})
		return
	}

	flight.Name = input.Name
	flight.Price = input.Price
	flight.DepartureTime = input.DepartureTime
	flight.ArrivalTime = input.ArrivalTime
	flight.Availability = input.Availability
	flight.LocationFrom = input.LocationFrom
	flight.LocationTo = input.LocationTo

	if err := config.DB.Save(&flight).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update flight"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Flight updated successfully"})
}



func AddHotel(c *gin.Context) {
	var input HotelInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	

	hotel := models.Hotel {
		HotelID: "hotel-" + uuid.New().String(),
		Name: input.Name,
		Price: input.Price,
		Availability: input.Availability,
		Location: input.Location,
	}

	if err := config.DB.Create(&hotel).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add hotel"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Hotel added successfully"})
}


func EditHotel(c *gin.Context) {
	hotelID := c.Param("hotel_id")

	var input HotelInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//check exist
	var hotel models.Hotel
	if err := config.DB.First(&hotel, "hotel_id = ?", hotelID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Hotel not found"})
		return
	}

	hotel.Name = input.Name
	hotel.Price = input.Price
	hotel.Availability = input.Availability
	hotel.Location = input.Location

	if err := config.DB.Save(&hotel).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update hotel"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Hotel updated successfully"})
}

func GetAnalytics (c *gin.Context) {
	var response models.AnalyticsResponse

	// Flight Booking counts
	var totalFlight, canceledFlight, pendingFlight int64
	config.DB.Model(&FlightBooking{}).Count(&totalFlight)
	config.DB.Model(&FlightBooking{}).Where("status = ?", "canceled").Count(&canceledFlight)
	config.DB.Model(&FlightBooking{}).Where("status = ?", "pending").Count(&pendingFlight)

	response.Booking.Flight.Total = int(totalFlight)
	response.Booking.Flight.Status.Canceled = int(canceledFlight)
	response.Booking.Flight.Status.Pending = int(pendingFlight)

	// Hotel Booking counts
	var totalHotel, canceledHotel, pendingHotel int64
	config.DB.Model(&HotelBooking{}).Count(&totalHotel)
	config.DB.Model(&HotelBooking{}).Where("status = ?", "canceled").Count(&canceledHotel)
	config.DB.Model(&HotelBooking{}).Where("status = ?", "pending").Count(&pendingHotel)

	response.Booking.Hotel.Total = int(totalHotel)
	response.Booking.Hotel.Status.Canceled = int(canceledHotel)
	response.Booking.Hotel.Status.Pending = int(pendingHotel)

	// User count
	var totalUsers int64
	config.DB.Model(&models.User{}).Count(&totalUsers)
	response.User.Total = int(totalUsers)

	c.JSON(http.StatusOK, response)
}