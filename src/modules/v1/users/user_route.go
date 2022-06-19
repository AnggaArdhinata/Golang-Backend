package users

import (
	"backend-golang/src/middleware"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func New(rt *mux.Router, db *gorm.DB) {
	route := rt.PathPrefix("/users/").Subrouter()

	repo := NewRepo(db)
	svc := NewService(repo)
	ctrl := NewCtrl(svc)

	route.HandleFunc("/", middleware.Do(ctrl.GetAll, middleware.CheckAuth)).Methods("GET")
	route.HandleFunc("/{userid}", ctrl.GetById).Methods("GET")
	route.HandleFunc("/", ctrl.AddData).Methods("POST")
	route.HandleFunc("/update", middleware.Do(ctrl.UpdateData, middleware.CheckAuth)).Methods("PUT")
	route.HandleFunc("/{userid}", middleware.Do(ctrl.DeleteData, middleware.CorsMid)).Methods("DELETE")

}
