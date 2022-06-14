package auth

import (
	"backend-golang/src/database/gorm/models"
	"backend-golang/src/helpers"
	"backend-golang/src/interfaces"
	
)

type token_response struct {
	Tokens string `json:"token`
}

type auth_service struct {
	rep interfaces.UserRepo
}

func NewService(rep interfaces.UserRepo) *auth_service {
	return &auth_service{rep}
}

func (r *auth_service) Login(body models.User) (*helpers.Response, error) {
	user, err := r.rep.FindByUsername(body.Username)


	if err != nil {
		return nil, err
	}
	if !helpers.CheckPassword(user.Password, body.Password) {
		return helpers.New("Password salah", 401, true), nil
	}

	token := helpers.NewToken(body.Username)
	theToken, err := token.Create()
	if err != nil {
		return nil, err
	}
	res := helpers.New(token_response{Tokens: theToken}, 200, false)

	return res, nil

}
