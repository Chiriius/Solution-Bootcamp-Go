package endpoints

import (
	"bootcamp_api/api/entities"

	"github.com/stretchr/testify/mock"
)

type serviceMock struct {
	mock.Mock
}

func (m *serviceMock) GetUser(id string) (entities.User, error) {
	r := m.Called(id)
	return r.Get(0).(entities.User), r.Error(1)

}

func (m *serviceMock) AddUser(user entities.User) (entities.User,error) { // reemplazar el model.user duplicando la struct

	r := m.Called(user)
	return user,r.Error(0)
}
