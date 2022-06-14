package users

import (
	"backend-golang/src/database/gorm/models"
	"backend-golang/src/helpers"
	"backend-golang/src/interfaces"
	"fmt"
)

type users_service struct {
	rep interfaces.UserRepo 
}

func NewService(rep interfaces.UserRepo) *users_service {
	return &users_service{rep}
}

func (re *users_service) FindAll() (*helpers.Response, error) {
	data, err := re.rep.FindAll()
	if err !=nil {
		return nil, err
	}

	response := helpers.New(data, 200, false)
	return response, nil
}
func (re *users_service) FindById(id int) (*helpers.Response, error) {
	data, err := re.rep.FindByUserId(id)
	if err !=nil {
		return nil, err
	}

	response := helpers.New(data, 200, false)
	return response, nil
}

func (re *users_service) FindByUsername(username string) (*helpers.Response, error) {
	data, err := re.rep.FindByUsername(username)
	if err !=nil {
		return nil, err
	}

	response := helpers.New(data, 200, false)
	return response, nil
}
func (re *users_service) Save(usr *models.User) (*helpers.Response, error) {
	hassPassword, err := helpers.HashPassword(usr.Password)
	if err !=nil {
		return nil, err
	}

	usr.Password = hassPassword
	data, err := re.rep.Add(usr)
	if err !=nil {
		return nil, err
	}

	response := helpers.New(data, 200, false)
	return response, nil
}
func (re *users_service) Del(data int) (*helpers.Response, error) {
	err := re.rep.Delete(data)

	if err !=nil {
		return nil, err
	}

	result := helpers.New("Berhasil Delete Data", 200, false)

	return result, nil
}
func (re *users_service) Updt(id int, data *models.User) (*helpers.Response, error) {
	fmt.Println(data)
	data, err := re.rep.Update(id, data)
	if err !=nil {
		return nil, err
	}
	result := helpers.New("Berhasil Update Data", 200, false)

	return result, nil


}
	
