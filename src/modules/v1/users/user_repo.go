package users

import (
	"backend-golang-try/src/helpers"
	"errors"
	"gorm.io/gorm"
)

var users Users
var response helpers.Response

type user_repo struct {
	db *gorm.DB
}

func NewRepo(grm *gorm.DB) *user_repo {
	return &user_repo{grm}
}

func (r *user_repo) FindAll() (*Users, error) {
	var users Users

	result := r.db.Find(&users)

	if result.Error != nil {
	 	response.ResponseJSON(400, users)
		return &users, nil
	}
	response.ResponseJSON(204, users)
	return &users, nil

}

func (r *user_repo) Add(data *User) (*User, error) {

	result := r.db.Create(data)

	if result.Error != nil {
		return nil, errors.New("Gagal menambah data")
	}
	return data, nil
}
func (r *user_repo) UpdateUsr(id *int, data *User) (*helpers.Response, error) {

	result := r.db.Model(&User{}).Where("userid = ?", &id).Updates(&User{Username: data.Username, Password: data.Password})

	if result.Error != nil {
		res := response.ResponseJSON(400, users)
		return res, nil
	}

	getData := r.db.First(&users, &id)
	if getData.RowsAffected < 1 {
		res := response.ResponseJSON(404, users)
		return res, nil
	}

	res := response.ResponseJSON(201, users)
	return res, nil
}
func (r *user_repo) DeleteUsr(data *int) (*helpers.Response, error) {
	r.db.First(&users, &data)
	result := r.db.Delete(&User{}, &data)
	if result.Error != nil {
		res := response.ResponseJSON(400, users)
		return res, nil
	}

	res := response.ResponseJSON(204, users)
	return res, nil
	
}

