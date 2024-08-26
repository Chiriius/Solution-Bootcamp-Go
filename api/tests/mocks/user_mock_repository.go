package mocks

import (
	"bootcamp_api/api/models"

	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (m *UserRepositoryMock) GetUserById(id string) (models.User, error) {
	r := m.Called(id)
	return r.Get(0).(models.User), r.Error(0)
}

func (m *UserRepositoryMock)AddUser( user models.User) error{

	r := m.Called(user)
    return r.Error(1)
}