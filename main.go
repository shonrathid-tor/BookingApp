package main

import (
	"booking_app/config"
	"booking_app/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//connect DB

	config.ConnectDatabase()

	//routes

	routes.UserRoutes(r)
	routes.BookingRoutes(r)

	r.Run(":8080")

}
