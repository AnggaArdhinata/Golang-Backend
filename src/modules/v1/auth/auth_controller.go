package auth

import (
	"backend-golang/src/database/gorm/models"
	"backend-golang/src/interfaces"
	"encoding/json"
	"net/http"
)

type auth_ctrl struct {
	rep interfaces.Auth_Service
}

func NewCtrl(rep interfaces.Auth_Service) *auth_ctrl {
	return &auth_ctrl{rep}
}
func (re *auth_ctrl) SignIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var data models.User

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	
	result, err := re.rep.Login(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	result.Send(w)
}