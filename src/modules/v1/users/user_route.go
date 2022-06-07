package users

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func New(rt *mux.Router, db *gorm.DB) {
	route := rt.PathPrefix("/user/").Subrouter()

	repo := NewRepo(db)
	ctrl := NewCtrl(repo)

	route.HandleFunc("/", ctrl.GetAll).Methods("GET")
	route.HandleFunc("/", ctrl.AddData).Methods("POST")
	route.HandleFunc("/update", ctrl.UpdateData).Methods("PUT")
	route.HandleFunc("/{id}", ctrl.DeleteData).Methods("DELETE")


}
