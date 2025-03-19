package controllers

import (
	"booking_app/config"
	"booking_app/models"

	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateBooking(c *gin.Context) {
	var booking models.Booking
	if err := c.ShouldBindJSON(&booking); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "BAD request " })
		return
	}

	config.DB.Create(&booking)
	c.JSON(http.StatusOK, gin.H{"message": "Booking created successfully!"})
}

func GetBookings(c *gin.Context) {
	var bookings []models.Booking
	config.DB.Find(&bookings)
	c.JSON(http.StatusOK, bookings)
}
