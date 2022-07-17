package models

import (
	"time"
)

type Vehicle struct {
	Vehicleid 	uint    `gorm:"primaryKey" json:"vehicleid"`
	Name      	string  `json:"name"`
	Location 	string  `json:"location"`
	Type     	string  `json:"type"`
	Desc      	string  `json:"desc"`
	Price     	string  `json:"price"`
	Image		string	`json:"image"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Vehicles []Vehicle
