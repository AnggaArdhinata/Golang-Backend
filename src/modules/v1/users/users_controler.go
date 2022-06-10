package users

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"backend-golang/src/database/gorm/models"

	"github.com/gorilla/mux"
)

type user_ctrl struct {
	repo *user_repo
}

func NewCtrl(rep *user_repo) *user_ctrl {
	return &user_ctrl{rep}
}
func (rep *user_ctrl) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data, err := rep.repo.FindAll()
	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	json.NewEncoder(w).Encode(data)
}
func (rep *user_ctrl) AddData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var data models.User
	json.NewDecoder(r.Body).Decode(&data)

	resultusr, err := rep.repo.Add(&data)
	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	json.NewEncoder(w).Encode(&resultusr)
}
func (rep *user_ctrl) UpdateData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var dataId = r.URL.Query()
	var data models.User

	json.NewDecoder(r.Body).Decode(&data)

	id, err := strconv.Atoi(dataId["id"][0])
	if err != nil {
		log.Fatal("Terjadi Kesalahan", err)
	}

	result, err := rep.repo.UpdateUsr(&id, &data)
	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	json.NewEncoder(w).Encode(&result)
}

func (rep *user_ctrl) DeleteData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var data = mux.Vars(r)
	id, err := strconv.Atoi(data["id"])

	result, err := rep.repo.DeleteUsr(&id)
	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	json.NewEncoder(w).Encode(&result)
}