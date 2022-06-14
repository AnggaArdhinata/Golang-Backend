package routers

import (
	database "backend-golang/src/database/gorm"

	"backend-golang/src/modules/v1/auth"
	"backend-golang/src/modules/v1/users"
	"backend-golang/src/modules/v1/vehicle"
	"backend-golang/src/modules/v1/order"

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
	order.New(mainRoute, db)
	auth.New(mainRoute, db)
	

	
	return mainRoute, nil
}

