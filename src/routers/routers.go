package routers

import (
	database "backend-golang/src/database/gorm"
	"backend-golang/src/modules/v1/users"
	"backend-golang/src/modules/v1/vehicle"

	"github.com/gorilla/mux"
)

func New() (*mux.Router, error) {
	mainRoute := mux.NewRouter()
	db, err := database.New()
	if err != nil {
		return nil, err
	}

	users.New(mainRoute, db)
	vehicle.New(mainRoute, db)

	
	return mainRoute, nil
}

