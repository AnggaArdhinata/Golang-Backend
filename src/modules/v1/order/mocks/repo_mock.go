package mocks

import (
	"backend-golang/src/database/gorm/models"

	"github.com/stretchr/testify/mock"
)

type RepoMock struct {
	Mock mock.Mock
}

// func (pr *RepoMock)FindAll() (*models.Orders, error) {
// 	args := pr.Mock.
// 	return mock.Arguments.Get(0), args
// }


func (pr *RepoMock) FindByUserId(id int) (*models.Order, error) {
	args := pr.Mock.Called(id)
	return args.Get(0).(*models.Order), args.Error(1)
}

func (pr *RepoMock) Add(data *models.Order) (*models.Order, error){
	args := pr.Mock.Called(data)
	return args.Get(0).(*models.Order), args.Error(1)
}