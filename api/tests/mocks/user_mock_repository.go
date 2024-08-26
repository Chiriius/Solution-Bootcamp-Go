package mocks

import (
	"bootcamp_api/api/entities"

	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (m *UserRepositoryMock) GetUserById(id string) (entities.User, error) {
	r := m.Called(id)
	return r.Get(0).(entities.User), r.Error(0)
}

func (m *UserRepositoryMock) AddUser(user entities.User) error {

	r := m.Called(user)
	return r.Error(1)
}
