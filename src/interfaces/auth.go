package interfaces

import (
	"backend-golang/src/database/gorm/models"
	"backend-golang/src/helpers"
)

type Auth_Service interface {
	Login(body models.User) (*helpers.Response, error)
}