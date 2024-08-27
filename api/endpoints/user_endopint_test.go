package endpoints

import (
	"bootcamp_api/api/entities"
	"bootcamp_api/api/services"
	"context"
	"errors"
	"testing"

	"github.com/go-kit/kit/endpoint"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetUser(t *testing.T) {

	testScenarios := []struct {
		testName        string
		endpoint        func(services.UserService) endpoint.Endpoint
		service         services.UserService
		mock            *serviceMock
		configureMock   func(*serviceMock, entities.User, error)
		endpointRequest interface{}
		mockResponse    entities.User
		mockError       error
		expectedOutput  GetUserResponse
		expetedError    error
	}{
		{
			testName: "test  get user",
			mock:     &serviceMock{},
			mockResponse: entities.User{
				ID: "3",
			},
			configureMock: func(m *serviceMock, mockResponse entities.User, mockError error) {
				m.On("GetUser", mock.Anything).Return(mockResponse, mockError)
			},
			expectedOutput: GetUserResponse{User: entities.User{
				ID: "3",
			}},
			expetedError:    nil,
			endpointRequest: GetUserRequest{ID: "3"},
		},
		{
			testName: "test  get user with error",
			mock:     &serviceMock{},
			mockResponse: entities.User{
				ID: "3",
			},
			configureMock: func(m *serviceMock, mockResponse entities.User, mockError error) {
				m.On("GetUser", mock.Anything).Return(mockResponse, mockError)
			},
			expectedOutput: GetUserResponse{},
			mockError:       errors.New("Service error"),
			expetedError:    errors.New("Service error"),
			endpointRequest: GetUserRequest{ID: "3"},
		},
	}
	for _, tt := range testScenarios {
		t.Run(tt.testName, func(t *testing.T) {

			// Prepare
			tt.endpoint = MakeGetUserEndpoint
			if tt.configureMock != nil {
				tt.configureMock(tt.mock, tt.mockResponse, tt.mockError)
			}
			ctx := context.TODO()

			//Nothing to prepare in this case

			// Act
			result, err := tt.endpoint(tt.mock)(ctx, tt.endpointRequest)

			// Assert
			assert.Equal(t, tt.expetedError, err)
			assert.Equal(t, tt.expectedOutput, result)
		})
	}

}

func TestAddUser(t *testing.T) {

	testScenarios := []struct {
		testName        string
		endpoint        func(services.UserService) endpoint.Endpoint
		service         services.UserService
		mock            *serviceMock
		configureMock   func(*serviceMock, entities.User)
		endpointRequest interface{}
		mockResponse    entities.User
		expectedOutput  CreateUserResponse
		expectedError   error
	}{
		{
			testName: "test add user",
			mock:     &serviceMock{},
			mockResponse: entities.User{
				ID: "5",
			},
			configureMock: func(m *serviceMock, mockResponse entities.User) {
				m.On("AddUser", mock.Anything).Return(mockResponse, nil)
			},
			expectedOutput: CreateUserResponse{
				Id: "5",
			},
			expectedError:   nil,
			endpointRequest: CreateUserRequest{Password: "12345678", Age: "20", Information: "add", Parents: "idk", Email: "alexer@gmail.com", Name: "Alexer Maestre"},
		},
	}

	for _, tt := range testScenarios {
		t.Run(tt.testName, func(t *testing.T) {

			// Prepare
			tt.endpoint = MakeAddUserEndpoint
			if tt.configureMock != nil {
				tt.configureMock(tt.mock, tt.mockResponse)
			}
			ctx := context.TODO()

			// Act
			result, err := tt.endpoint(tt.mock)(ctx, tt.endpointRequest)

			// Assert
			assert.Equal(t, tt.expectedError, err)
			assert.Equal(t, tt.expectedOutput, result)
		})
	}
}
