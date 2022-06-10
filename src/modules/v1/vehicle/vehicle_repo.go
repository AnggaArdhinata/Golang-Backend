package vehicle

import (
	"backend-golang/src/database/gorm/models"
	"backend-golang/src/helpers"
	"errors"

	"gorm.io/gorm"
)

var vehicles models.Vehicles
var response helpers.Response

type vehicle_repo struct {
	db *gorm.DB
}

func NewRepo(grm *gorm.DB) *vehicle_repo {
	return &vehicle_repo{grm}
}

func (r *vehicle_repo) FindAll() (*models.Vehicles, error) {
	var vehicles models.Vehicles
	
	result := r.db.Find(&vehicles)

	if result.Error != nil {
		return nil, errors.New("Gagal mengambil data")
	}
	return &vehicles, nil
}
func (r *vehicle_repo) Add(data *models.Vehicle) (*models.Vehicle, error) {

	result := r.db.Create(data)

	if result.Error != nil {
		return nil, errors.New("Gagal menambah data")
	}
	return data, nil
}
func (r *vehicle_repo) UpdateVhc(id *int, data *models.Vehicle) (*helpers.Response, error) {

	result := r.db.Model(&models.Vehicle{}).Where("vehicleid = ?", &id).Updates(&models.Vehicle{Cars: data.Cars, Motorbike: data.Motorbike, Bike: data.Bike})

	if result.Error != nil {
		res := response.ResponseJSON(400, vehicles)
		return res, nil
	}

	getData := r.db.First(&vehicles, &id)
	if getData.RowsAffected < 1 {
		res := response.ResponseJSON(404, vehicles)
		return res, nil
	}

	res := response.ResponseJSON(201, vehicles)
	return res, nil
}
func (r *vehicle_repo) DeleteVhc(data *int) (*helpers.Response, error) {
	r.db.First(&vehicles, &data)
	result := r.db.Delete(&models.Vehicles{}, &data)
	if result.Error != nil {
		res := response.ResponseJSON(400, vehicles)
		return res, nil
	}

	res := response.ResponseJSON(204, vehicles)
	return res, nil
	
} 
