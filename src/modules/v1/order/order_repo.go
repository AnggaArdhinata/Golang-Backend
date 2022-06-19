package order

import (
	"backend-golang/src/database/gorm/models"
	"errors"

	"gorm.io/gorm"
)

type orders_repo struct {
	db *gorm.DB
}

func NewRepo(grm *gorm.DB) *orders_repo {
	return &orders_repo{grm}
}
func (r *orders_repo) FindAll() (*models.Orders, error) {
	var orders models.Orders

	result := r.db.Find(&orders)

	if result.Error != nil {
		return nil, errors.New("gagal mengambil data")
	}
	return &orders, nil
}
func (r *orders_repo) FindByUserId(id int) (*models.Order, error) {
	var orders models.Order

	result := r.db.Preload("Vehicle").Where("userid = ?").Find(&orders)

	if result.Error != nil {
		return nil, errors.New("gagal mengambil data")
	}
	return &orders, nil
}
func (r *orders_repo) Add(data *models.Order) (*models.Order, error) {
	result := r.db.Create(data)

	if result.Error != nil {
		return nil, errors.New("gagal menyimpan data")
	}
	return data, nil
}