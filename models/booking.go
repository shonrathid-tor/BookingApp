package models

import "gorm.io/gorm"

type Booking struct {
	gorm.Model
	UserID  uint   `json:"user_id"`
	Service string `json:"service"`
	Date    string `json:"date"`

	Time   string `json:"time"`
	Status string `json:"status"`
}
