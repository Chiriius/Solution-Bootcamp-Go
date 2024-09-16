package mysql

import (
	"bootcamp_api/api/entities"

	"github.com/stretchr/testify/mock"
)

type userRepositoryMock struct {
	mock.Mock
}

func (m *userRepositoryMock) GetUser(id string) (entities.User, error) {
	r := m.Called(id)
	return r.Get(0).(entities.User), r.Error(1)
}

func (m *userRepositoryMock) AddUser(user entities.User) (entities.User, error) {
	r := m.Called(user)
	return r.Get(0).(entities.User), r.Error(1)
}

func (m *userRepositoryMock) UpdateUser(user entities.User) (entities.User, error) {
	r := m.Called(user)
	return r.Get(0).(entities.User), r.Error(1)
}
