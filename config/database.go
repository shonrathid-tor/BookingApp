package config

import (
	"fmt"
	"log"

	"booking_app/models" // Import your models

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=localhost user=postgres password=1234 dbname=booking port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database!", err)
	}

	// Run AutoMigrate to create tables
	err = database.AutoMigrate(&models.User{}, &models.Booking{})
	if err != nil {
		log.Fatal("Failed to migrate database!", err)
	}

	DB = database
	fmt.Println("Database connected and migrated successfully!")
}
