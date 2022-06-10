package vehicle

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"backend-golang/src/database/gorm/models"

	"github.com/gorilla/mux"
)

type vehicle_ctrl struct {
	repo *vehicle_repo
}

func NewCtrl(rep *vehicle_repo) *vehicle_ctrl {
	return &vehicle_ctrl{rep}
}
func (rep *vehicle_ctrl) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data, err := rep.repo.FindAll()
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	json.NewEncoder(w).Encode(data)
}
func (rep *vehicle_ctrl) AddData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var data models.Vehicle
	json.NewDecoder(r.Body).Decode(&data)

	resultusr, err := rep.repo.Add(&data)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	json.NewEncoder(w).Encode(&resultusr)
}
func (rep *vehicle_ctrl) UpdateData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var dataId = r.URL.Query()
	var data models.Vehicle

	json.NewDecoder(r.Body).Decode(&data)

	id, err := strconv.Atoi(dataId["id"][0])
	if err != nil {
		log.Fatalf("Terjadi Kesalahan", err)
	}

	result, err := rep.repo.UpdateVhc(&id, &data)
	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	json.NewEncoder(w).Encode(&result)
}

func (rep *vehicle_ctrl) DeleteData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	var data = mux.Vars(r)
	id, err := strconv.Atoi(data["id"])

	if err != nil {
		log.Fatalf("Gagal Delete Data", err)
	}

	result, err := rep.repo.DeleteVhc(&id)
	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	json.NewEncoder(w).Encode(&result)
}