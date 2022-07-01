package mocks

import (
	"backend-golang/src/database/gorm/models"

	"github.com/stretchr/testify/mock"
)

type VehicleMock struct {
	Mock mock.Mock
}

func (vr *VehicleMock) FindAll() (*models.Vehicles, error)

func (vr *VehicleMock) FindByVehiceleId(id int) (*models.Vehicle, error) {
	args := vr.Mock.Called(id)
	return args.Get(0).(*models.Vehicle), args.Error(1)
}

func (vr *VehicleMock) Add(data *models.Vehicle) (*models.Vehicle, error){
args := vr.Mock.Called(data)
	return args.Get(0).(*models.Vehicle), args.Error(1)
}

func (vr *VehicleMock) Delete(data int) error
func (vr *VehicleMock) Update(id int, data *models.Vehicle) (*models.Vehicle, error)