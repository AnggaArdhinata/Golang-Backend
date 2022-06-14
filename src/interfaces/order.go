package interfaces

import (
	"backend-golang/src/database/gorm/models"
	"backend-golang/src/helpers"
)

type OrderRepo interface {
	FindAll() (*models.Orders, error)
	FindByUserId(id int) (*models.Order, error)
	Add(*models.Order) (*models.Order, error)
	// Delete(data int) error
	
}
type OrderService interface {
	FindAll() (*helpers.Response, error)
	FindById(id int) (*helpers.Response, error)
	Save(*models.Order) (*helpers.Response, error)
	// Del(data int) (*helpers.Response, error)
	
}