package order

import (
	"backend-golang/src/database/gorm/models"
	"backend-golang/src/interfaces"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type order_ctrl struct {
	repo interfaces.OrderService
}

func NewCtrl(rep interfaces.OrderService) *order_ctrl {
	return &order_ctrl{rep}
}

func (rep *order_ctrl) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data, err := rep.repo.FindAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	data.Send(w)
	
}
func (rep *order_ctrl) GetById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	iduser, err := strconv.Atoi(vars["userid"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}


	result, err := rep.repo.FindById(iduser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	result.Send(w)
	
}

func (rep *order_ctrl) AddData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var data models.Order

	json.NewDecoder(r.Body).Decode(&data)
	result, err := rep.repo.Save(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	result.Send(w)
}