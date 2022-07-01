package vehicle

import (
	"backend-golang/src/database/gorm/models"
	"errors"

	"gorm.io/gorm"
)

type vehicles_repo struct {
	db *gorm.DB
}

func NewRepo(grm *gorm.DB) *vehicles_repo {
	return &vehicles_repo{grm}
}
func (r *vehicles_repo) FindAll() (*models.Vehicles, error) {
	var vehicles models.Vehicles

	result := r.db.Find(&vehicles)

	if result.Error != nil {
		return nil, errors.New("gagal mengambil data")
	}
	return &vehicles, nil
}
func (r *vehicles_repo) FindByVehiceleId(id int) (*models.Vehicle, error) {
	var data models.Vehicle

	result := r.db.First(&data, "cars = ?", id)

	if result.Error != nil {
		return nil, errors.New("gagal mengambil data")
	}
	return &data, nil
}
func (r *vehicles_repo) Add(data *models.Vehicle) (*models.Vehicle, error) {
	result := r.db.Create(data)

	if result.Error != nil {
		return nil, errors.New("gagal menyimpan data")
	}
	return data, nil
}
func (r *vehicles_repo) Delete(data int) error {
	result := r.db.Delete(models.Vehicle{}, data)

	if result.Error != nil {
		return errors.New("gagal menghapus data")
	}
	return nil

}
func (r *vehicles_repo) Update(id int, data *models.Vehicle) (*models.Vehicle, error){
	var vehicles models.Vehicle
	result := r.db.Model(&vehicles).Where("vehicleid = ?", id).Updates(data)
	if result.Error !=nil {
		return nil, errors.New("gagal Mengupdate data")

	}
	return data, nil
}