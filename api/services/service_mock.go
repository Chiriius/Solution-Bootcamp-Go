package services

import (
	"bootcamp_api/api/entities"

	"github.com/stretchr/testify/mock"
)

type userServiceMock struct {
	mock.Mock
}

func (m *userServiceMock) GetUser(id string) (entities.User, error) {
	r := m.Called(id)
	return r.Get(0).(entities.User), r.Error(1)
}

func (m *userServiceMock) AddUser(user entities.User) (entities.User, error) {
	r := m.Called(user)
	return r.Get(0).(entities.User), r.Error(1)
}

func (m *userServiceMock) UpdateUser(user entities.User) (entities.User, error) {
	r := m.Called(user)
	return r.Get(0).(entities.User), r.Error(1)

}
