package vehicle

import (
	"backend-golang/src/database/gorm/models"
	"backend-golang/src/interfaces"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type vehicles_ctrl struct {
	repo interfaces.VehicleService
}

func NewCtrl(rep interfaces.VehicleService) *vehicles_ctrl {
	return &vehicles_ctrl {rep}
}

func (rep *vehicles_ctrl) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data, err := rep.repo.FindAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	data.Send(w)
	
}

func (rep *vehicles_ctrl) AddData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var data models.Vehicle

	json.NewDecoder(r.Body).Decode(&data)
	result, err := rep.repo.Save(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	result.Send(w)
}

func (rep *vehicles_ctrl) DeleteData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	idVehicle, err := strconv.Atoi(vars["vehicleid"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	result, err := rep.repo.Del(idVehicle)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	result.Send(w)

}
func (rep *vehicles_ctrl) UpdateData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var dataId = r.URL.Query()
	var data models.Vehicle

	json.NewDecoder(r.Body).Decode(&data)

	id, err := strconv.Atoi(dataId["id"][0])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	result, err := rep.repo.Updt(id, &data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	result.Send(w)

}