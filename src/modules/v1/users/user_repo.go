package users

import (
	"backend-golang/src/database/gorm/models"
	"errors"

	"gorm.io/gorm"
)

type users_repo struct {
	db *gorm.DB
}

func NewRepo(grm *gorm.DB) *users_repo {
	return &users_repo{grm}
}
func (r *users_repo) FindAll() (*models.Users, error) {
	var users models.Users

	result := r.db.Find(&users)

	if result.Error != nil {
		return nil, errors.New("Gagal mengambil data")
	}
	return &users, nil
}
func (r *users_repo) FindByUserId(id int) (*models.User, error) {
	var users models.User

	result := r.db.Preload("Order").First(&users)

	if result.Error != nil {
		return nil, errors.New("Gagal mengambil data")
	}
	return &users, nil
}

func (r *users_repo) FindByUsername(username string) (*models.User, error) {
	var data models.User

	result := r.db.First(&data, "username = ?", username)

	if result.Error != nil {
		return nil, errors.New("Gagal mengambil data")
	}
	return &data, nil
}
func (r *users_repo) Add(data *models.User) (*models.User, error) {
	result := r.db.Create(data)

	if result.Error != nil {
		return nil, errors.New("Gagal menyimpan data")
	}
	return data, nil
}
func (r *users_repo) Delete(data int) error {
	result := r.db.Delete(models.User{}, data)

	if result.Error != nil {
		return errors.New("Gagal menghapus data")
	}
	return nil

}
func (r *users_repo) Update(id int, data *models.User) (*models.User, error){
	var users models.User
	result := r.db.Model(&users).Where("userid = ?", id).Updates(data)
	if result.Error !=nil {
		return nil, errors.New("Gagal Mengupdate data")

	}
	return data, nil
}