package users

import (
	"backend-golang/src/database/gorm/models"
	"backend-golang/src/interfaces"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type users_ctrl struct {
	repo interfaces.UserService
}

func NewCtrl(rep interfaces.UserService) *users_ctrl {
	return &users_ctrl{rep}
}

func (rep *users_ctrl) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data, err := rep.repo.FindAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	data.Send(w)
	
}
func (rep *users_ctrl) GetById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	iduser, err := strconv.Atoi(vars["userid"])

	result, err := rep.repo.FindById(iduser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	result.Send(w)
	
}

func (rep *users_ctrl) AddData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var data models.User

	json.NewDecoder(r.Body).Decode(&data)
	result, err := rep.repo.Save(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	result.Send(w)
}

func (rep *users_ctrl) DeleteData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	iduser, err := strconv.Atoi(vars["userid"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	result, err := rep.repo.Del(iduser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	result.Send(w)

}
func (rep *users_ctrl) UpdateData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var dataId = r.URL.Query()
	var data models.User

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
