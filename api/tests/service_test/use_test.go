package service_test

import (
	"bootcamp_api/api/entities"

	"bootcamp_api/api/services"
	"bootcamp_api/api/tests/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUser(t *testing.T) {
	userRepositoryMock := new(mocks.UserRepositoryMock)

	userService := services.NewUserService(userRepositoryMock)

	testUserID := "123"
	userr := entities.User{
		ID:          testUserID,
		Password:    "secret",
		Age:         "25",
		Information: "Juancho",
		Parents:     "Maestre",
		Email:       "Juancho@gamil.com",
		Name:        "Juancho Roy",
	}

	userRepositoryMock.On("GetUserById", testUserID).Return(userr, nil)

	user, err := userService.GetUser(testUserID)

	assert.Nil(t, err)
	assert.Equal(t, userr, user)

	userRepositoryMock.AssertCalled(t, "GetUserById", testUserID)
}
