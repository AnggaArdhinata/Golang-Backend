package interfaces

import (
	"backend-golang/src/database/gorm/models"
	"backend-golang/src/helpers"
)

type VehicleRepo interface {
	FindAll() (*models.Vehicles, error)
	FindByTypeCars(cars string) (*models.Vehicle, error)
	Add(*models.Vehicle) (*models.Vehicle, error)
	Delete(data int) error
	Update(id int, data *models.Vehicle) (*models.Vehicle, error)

}
type VehicleService interface {
	FindAll() (*helpers.Response, error)
	FindByTypeCars(cars string) (*helpers.Response, error)
	Save(*models.Vehicle) (*helpers.Response, error)
	Del(data int) (*helpers.Response, error)
	Updt(id int, data *models.Vehicle) (*helpers.Response, error)

}
