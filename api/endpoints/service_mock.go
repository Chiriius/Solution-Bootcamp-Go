package endpoints

import (
	"bootcamp_api/api/models"

	"github.com/stretchr/testify/mock"
)

type serviceMock struct {
	mock.Mock
}

func (m *serviceMock) GetUser(id string) (models.User, error) {
	r := m.Called(id)
	return r.Get(0).(models.User), r.Error(1)

}

func (m *serviceMock) AddUser(user models.User) error { // reemplazar el model.user duplicando la struct

	r := m.Called(user)
	return r.Error(0)
}
