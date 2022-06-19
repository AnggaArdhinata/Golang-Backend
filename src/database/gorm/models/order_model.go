package models

import (
	"time"
)

type Order struct {
	Orderid    uint      `gorm:"primaryKey" json:"oder_id"`
	Userid     int       `json:"user_id"`
	Vehiclesid int       `json:"Vehicleid"`
	Vehicles   Vehicle   `gorm:"foreignKey:Vehiclesid;references:Vehicleid"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type Orders []Order
