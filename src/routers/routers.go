package routers

import (
	"backend-golang-try/src/configs/database"
	"backend-golang-try/src/modules/v1/users"
	"backend-golang-try/src/modules/v1/vehicle"

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

