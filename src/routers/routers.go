package routers

import (
	database "backend-golang/src/database/gorm"
	"net/http"

	"backend-golang/src/modules/v1/auth"
	"backend-golang/src/modules/v1/order"
	"backend-golang/src/modules/v1/users"
	"backend-golang/src/modules/v1/vehicle"

	"github.com/gorilla/mux"
	"github.com/newrelic/go-agent/v3/integrations/nrgorilla"
	"github.com/newrelic/go-agent/v3/newrelic"
)

func New() (*mux.Router, error) {
	mainRoute := mux.NewRouter()
	nRelic, err := newrelic.NewApplication(
		newrelic.ConfigAppName("backend_golang"),
		newrelic.ConfigLicense("6724dbccbb20efd4929c41e65dc6ef6bb65fNRAL"),
		newrelic.ConfigDistributedTracerEnabled(true),
	)
	if err != nil {
		return nil, err
	}
	mainRoute.Use(nrgorilla.Middleware(nRelic))

	db, err := database.New()
	if err != nil {
		return nil, err
	}
	
	mainRoute.HandleFunc(newrelic.WrapHandleFunc(nRelic, "/relic", relicHandler)).Methods("GET")
	mainRoute.HandleFunc("/", simpleHandler).Methods("GET")

	users.New(mainRoute, db)
	vehicle.New(mainRoute, db)
	order.New(mainRoute, db)
	auth.New(mainRoute, db)
	

	
	return mainRoute, nil
}

func relicHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Relic"))
}

func simpleHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to Vehicle-Rental-App-Backend"))
}