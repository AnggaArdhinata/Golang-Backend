package models

import (
	"time"
)

type Vehicle struct {
	Vehicleid    uint   `gorm:"primaryKey" json:"vehicle_id"`
	Cars  		string  `json:"car"`
	Motorbike  	string  `json:"motorbike"`
	Bike     	string  `json:"bike"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Vehicles []Vehicle
