package endpoints

import (
	"bootcamp_api/api/entities"
	"bootcamp_api/api/services"

	"github.com/stretchr/testify/mock"
)

type serviceMock struct {
	mock.Mock
}

func (m *serviceMock) GetUser(id string) (entities.User, error) {
	r := m.Called(id)
	return r.Get(0).(entities.User), r.Error(1)

}

func (m *serviceMock) AddUser(user entities.User) (entities.User, error) { // reemplazar el model.user duplicando la struct
	r := m.Called(user)
	return r.Get(0).(entities.User), r.Error(1)
}

func (m *serviceMock) MakeServerEndpoints(s services.UserService) Endpoints {
	r:= m.Called(s)
	return r.Get(0).(Endpoints)
}
func (m *serviceMock) ModifyUser(user entities.User) (entities.User, error) {
	r:= m.Called(user)
	return r.Get(0).(entities.User),r.Error(1)
}
