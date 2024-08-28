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
			expectedOutput:  GetUserResponse{},
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
		mockError       error
		configureMock   func(*serviceMock, entities.User, error)
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
			configureMock: func(m *serviceMock, mockResponse entities.User, mockError error) {
				m.On("AddUser", mock.Anything).Return(mockResponse, mockError)
			},
			expectedOutput: CreateUserResponse{
				Id: "5",
			},
			expectedError:   nil,
			endpointRequest: CreateUserRequest{Password: "12345678", Age: "20", Information: "add", Parents: "idk", Email: "alexer@gmail.com", Name: "Alexer Maestre"},
		},
		{
			testName: "test add user with error",
			mock:     &serviceMock{},
			mockResponse: entities.User{
				ID: "5",
			},
			configureMock: func(m *serviceMock, mockResponse entities.User, mockError error) {
				m.On("AddUser", mock.Anything).Return(mockResponse, mockError)
			},
			expectedOutput:  CreateUserResponse{},
			mockError:       errors.New("Server Error"),
			expectedError:   errors.New("Server Error"),
			endpointRequest: CreateUserRequest{Password: "12345678", Age: "20", Information: "add", Parents: "idk", Email: "alexer@gmail.com", Name: "Alexer Maestre"},
		},
		// {
		// 	testName:        "test add user with wrong request type",
		// 	mock:            &serviceMock{},
		// 	configureMock: func(m *serviceMock, mockResponse entities.User, mockError error) {
		// 		m.On("AddUser", mock.Anything).Return(mockResponse, mockError)
		// 	},
		// 	expectedOutput:  CreateUserResponse{},
		// 	mockError:       errorss.ErrorInterfaceDifType,
		// 	expectedError:   errorss.ErrorInterfaceDifType,
		// 	endpointRequest: "invalid request",
		// },
		{
			testName:     "test add user with missing fields",
			mock:         &serviceMock{},
			mockResponse: entities.User{},
			configureMock: func(m *serviceMock, mockResponse entities.User, mockError error) {
				m.On("AddUser", mock.Anything).Return(mockResponse, mockError)
			},
			expectedOutput:  CreateUserResponse{},
			mockError:       errors.New("some validation error"),
			expectedError:   errors.New("some validation error"),
			endpointRequest: CreateUserRequest{Password: "", Age: "", Information: "", Parents: "", Email: "", Name: ""},
		},
	}

	for _, tt := range testScenarios {
		t.Run(tt.testName, func(t *testing.T) {

			// Prepare
			tt.endpoint = MakeAddUserEndpoint
			if tt.configureMock != nil {
				tt.configureMock(tt.mock, tt.mockResponse, tt.mockError)
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

func TestModifyUser(t *testing.T) {

	testScenarios := []struct {
		testName        string
		service         services.UserService
		endpoint        func(services.UserService) endpoint.Endpoint
		mock            *serviceMock
		mockError       error
		mockResponse    entities.User
		configureMock   func(*serviceMock, entities.User, error)
		expectedOutput  ModifyUserResponse
		expectedError   error
		endpointRequest interface{}
	}{
		{
			testName: "test modify user",
			mock:     &serviceMock{},
			mockResponse: entities.User{
				ID: "222",
			},
			configureMock: func(m *serviceMock, mockResponse entities.User, mockError error) {
				m.On("GetUser", mock.Anything).Return(mockResponse, mockError)
				m.On("ModifyUser", mock.Anything).Return(mockResponse, mockError)
			},
			expectedOutput: ModifyUserResponse{
				Id: "222",
			},
			expectedError:   nil,
			endpointRequest: ModifyUserRequest{Id: "222", Password: "12345678", Age: "20", Information: "add", Parents: "idk", Email: "alexer@gmail.com", Name: "Alexer Maestre"},
		},
	}

	for _, tt := range testScenarios {
		t.Run(tt.testName, func(t *testing.T) {
			
			// Prepare
			tt.endpoint = MakeModifyUserEndpoint
			if tt.configureMock != nil {
				tt.configureMock(tt.mock, tt.mockResponse, tt.mockError)
			}

			// Act
			ctx := context.TODO()
			result, err := tt.endpoint(tt.mock)(ctx, tt.endpointRequest)

			// Assert
			assert.Equal(t, tt.expectedError, err)
			assert.Equal(t, tt.expectedOutput, result)
		})
	}
}

func TestMakeServerEndpoints(t *testing.T) {

	testScenarios := []struct {
		testName       string
		service        services.UserService
		mock           *serviceMock
		mockResponse   Endpoints
		configureMock  func(*serviceMock, Endpoints)
		expectedOutput Endpoints
	}{
		{
			testName: "test MakeServerEndpoints",
			mock:     &serviceMock{},
			configureMock: func(m *serviceMock, mockResponse Endpoints) {
				m.On("ModifyUser", mock.Anything).Return(mockResponse)
			},
			expectedOutput: Endpoints{
				GetUser: MakeGetUserEndpoint(&serviceMock{}),
				AddUser: MakeAddUserEndpoint(&serviceMock{}),
			},
		},
	}

	for _, tt := range testScenarios {

		// Prepare
		t.Run(tt.testName, func(t *testing.T) {
			if tt.configureMock != nil {
				tt.configureMock(tt.mock, tt.mockResponse)
			}

			// Act
			result := MakeServerEndpoints(tt.mock)

			// Assert
			assert.Equal(t, tt.expectedOutput, result)
		})
	}
}

// func TestMakeServerEndpoints(t *testing.T) {

// 	testScenarios := []struct {
// 		testName       string
// 		service        services.UserService
// 		mock           *serviceMock
// 		configureMock  func(*serviceMock)
// 		expectedOutput Endpoints
// 	}{
// 		{
// 			testName: "test MakeServerEndpoints",
// 			mock:     &serviceMock{},
// 			configureMock: func(m *serviceMock) {
// 				m.On("GetUser", mock.Anything).Return(entities.User{ID: "3"}, nil)
// 				m.On("AddUser", mock.Anything).Return(entities.User{ID: "5"}, nil)
// 			},
// 			expectedOutput: Endpoints{
// 				GetUser: MakeGetUserEndpoint(&serviceMock{}),
// 				AddUser: MakeAddUserEndpoint(&serviceMock{}),
// 			},
// 		},
// 	}

// 	for _, tt := range testScenarios {

// 		// Prepare
// 		t.Run(tt.testName, func(t *testing.T) {
// 			if tt.configureMock != nil {
// 				tt.configureMock(tt.mock)
// 			}

// 			// Act
// 			result := MakeServerEndpoints(tt.mock)

// 			// Assert
// 			assert.NotNil(t, result.GetUser)
// 			assert.NotNil(t, result.AddUser)
// 		})
// 	}
// }
