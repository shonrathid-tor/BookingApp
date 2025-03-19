package routes

import (
	"booking_app/controllers"

	"github.com/gin-gonic/gin"
)

func BookingRoutes(router *gin.Engine) {
	router.POST("/bookings", controllers.CreateBooking)
	router.GET("/bookings", controllers.GetBookings)
}
