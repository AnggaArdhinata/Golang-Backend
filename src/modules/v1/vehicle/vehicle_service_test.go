package vehicle

import (
	"backend-golang/src/database/gorm/models"
	"backend-golang/src/modules/v1/vehicle/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var vehicleMock = models.Vehicle{

	Vehicleid: 1,
}

func TestVehicleId(t *testing.T) {
	var repos = mocks.VehicleMock{Mock: mock.Mock{}}
	var service = vehicles_services{&repos}

	repos.Mock.On("FindByVehiceleId", 1).Return(&vehicleMock, nil)
	data, err := service.FindById(1)

	var ekspektasi = uint(1)
	orders := data.Data.(*models.Vehicle)
	assert.Equal(t, ekspektasi, orders.Vehicleid, "Expect id = 1")
	assert.Nil(t, err)
}
