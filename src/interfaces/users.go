package interfaces

import (
	"backend-golang/src/database/gorm/models"
	"backend-golang/src/helpers"
)

type UserRepo interface {
	FindAll() (*models.Users, error)
	FindByUserId(id int) (*models.User, error)
	FindByUsername(username string) (*models.User, error)
	Add(*models.User) (*models.User, error)
	Delete(data int) error
	Update(id int, data *models.User) (*models.User, error)

}
type UserService interface {
	FindAll() (*helpers.Response, error)
	FindById(id int) (*helpers.Response, error)
	FindByUsername(username string) (*helpers.Response, error)
	Save(*models.User) (*helpers.Response, error)
	Del(data int) (*helpers.Response, error)
	Updt(id int, data *models.User) (*helpers.Response, error)

}
