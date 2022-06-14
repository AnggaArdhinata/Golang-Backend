package vehicle

import (
	"backend-golang/src/database/gorm/models"
	"backend-golang/src/helpers"
	"backend-golang/src/interfaces"
	"fmt"
)

type vehicles_services struct {
	rep interfaces.VehicleRepo 
}

func NewService(rep interfaces.VehicleRepo) *vehicles_services {
	return &vehicles_services{rep}
}

func (re *vehicles_services) FindAll() (*helpers.Response, error) {
	data, err := re.rep.FindAll()
	if err !=nil {
		return nil, err
	}

	response := helpers.New(data, 200, false)
	return response, nil
}
func (re *vehicles_services) FindByTypeCars(cars string) (*helpers.Response, error) {
	data, err := re.rep.FindByTypeCars(cars)
	if err !=nil {
		return nil, err
	}

	response := helpers.New(data, 200, false)
	return response, nil
}
func (re *vehicles_services) Save(vhcl *models.Vehicle) (*helpers.Response, error) {
	result, err := re.rep.Add(vhcl)
	if err !=nil {
		return nil, err
	}
	response := helpers.New(result, 200, false)
	return response, nil

}
func (re *vehicles_services) Del(data int) (*helpers.Response, error) {
	err := re.rep.Delete(data)

	if err !=nil {
		return nil, err
	}

	result := helpers.New("Berhasil Delete Data", 200, false)

	return result, nil
}
func (re *vehicles_services) Updt(id int, data *models.Vehicle) (*helpers.Response, error) {
	fmt.Println(data)
	data, err := re.rep.Update(id, data)
	if err !=nil {
		return nil, err
	}
	result := helpers.New("Berhasil Update Data", 200, false)

	return result, nil


}