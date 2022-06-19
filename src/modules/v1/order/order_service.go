package order

import (
	"backend-golang/src/database/gorm/models"
	"backend-golang/src/helpers"
	"backend-golang/src/interfaces"
)

type orders_service struct {
	rep interfaces.OrderRepo 
}

func NewService(rep interfaces.OrderRepo) *orders_service {
	return &orders_service{rep}
}

// func (re *orders_service) FindAll() (*helpers.Response, error) {
// 	data, err := re.rep.FindAll()
// 	if err !=nil {
// 		return nil, err
// 	}

// 	response := helpers.New(data, 200, false)
// 	return response, nil
// }
func (re *orders_service) FindById(id int) (*helpers.Response, error) {
	data, err := re.rep.FindByUserId(id)
	if err !=nil {
		return nil, err
	}

	response := helpers.New(data, 200, false)
	return response, nil
}
func (re *orders_service) Save(ordr *models.Order) (*helpers.Response, error) {
	result, err := re.rep.Add(ordr)
	if err !=nil {
		return nil, err
	}
	response := helpers.New(result, 200, false)
	return response, nil
}