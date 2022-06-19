package order

import (
	"backend-golang/src/database/gorm/models"
	"backend-golang/src/modules/v1/order/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var modelMock = models.Order{

	Orderid: 1,
}

func TestFindById(t *testing.T) {
	var repos = mocks.RepoMock{Mock: mock.Mock{}}
	var service = orders_service{&repos}

	repos.Mock.On("FindByUserId", 1).Return(&modelMock, nil)
	data, err := service.FindById(1)

	var ekspektasi = uint(1)
	orders := data.Data.(*models.Order)
	assert.Equal(t, ekspektasi, orders.Orderid, "Expect id = 1")
	assert.Nil(t, err)
}
