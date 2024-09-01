package services

import (
	"bootcamp_api/api/entities"
	"bootcamp_api/api/repository/mysql"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetUserService(t *testing.T) {

	testScenarios := []struct {
		testName       string
		mock           *userServiceMock
		mockResponse   entities.User
		mockError      error
		configureMock  func(*userServiceMock, entities.User, error)
		expectedOutput entities.User
		expectedError  error
	}{
		{
			testName: "TestGetUserService",
			mock:     &userServiceMock{},
			mockResponse: entities.User{
				ID: "3",
			},
			mockError: nil,
			configureMock: func(m *userServiceMock, mockResponse entities.User, mockError error) {
				m.On("GetUser", "3").Return(mockResponse, mockError)
			},
			expectedOutput: entities.User{
				ID: "3",
			},
			expectedError: nil,
		},
	}

	for _, tt := range testScenarios {
		t.Run(tt.testName, func(t *testing.T) {

			// Prepare
			if tt.configureMock != nil {
				tt.configureMock(tt.mock, tt.mockResponse, tt.mockError)
			}

			service := &userService{
				repository: tt.mock,
			}

			// Act
			result, err := service.GetUser(tt.mockResponse.ID)

			// Assert
			assert.Equal(t, tt.expectedError, err)
			assert.Equal(t, tt.expectedOutput, result)
		})
	}

}

func TestAddUserService(t *testing.T) {

	testScenarios := []struct {
		testName       string
		mock           *userServiceMock
		mockResponse   entities.User
		mockError      error
		configureMock  func(*userServiceMock, entities.User, error)
		expectedOutput entities.User
		expectedError  error
	}{
		{
			testName:     "TestAddUserService",
			mock:         &userServiceMock{},
			mockResponse: entities.User{Password: "12345678", Age: "20", Information: "add", Parents: "idk", Email: "alexer@gmail.com", Name: "Alexer Maestre"},
			mockError:    nil,
			configureMock: func(m *userServiceMock, mockResponse entities.User, mockError error) {
				m.On("AddUser", mock.Anything).Return(mockResponse, mockError)
			},
			expectedOutput: entities.User{Password: "12345678", Age: "20", Information: "add", Parents: "idk", Email: "alexer@gmail.com", Name: "Alexer Maestre"},
			expectedError:  nil,
		},
	}

	for _, tt := range testScenarios {
		t.Run(tt.testName, func(t *testing.T) {

			// Prepare
			if tt.configureMock != nil {
				tt.configureMock(tt.mock, tt.mockResponse, tt.mockError)
			}

			tt.expectedOutput.ID = tt.mockResponse.ID

			service := &userService{
				repository: tt.mock,
			}

			// Act
			result, err := service.AddUser(tt.mockResponse)

			// Assert
			assert.Equal(t, tt.expectedError, err)
			assert.Equal(t, tt.expectedOutput, result)
		})
	}

}

func TestUpdateUserService(t *testing.T) {

	testScenarios := []struct {
		testName       string
		mock           *userServiceMock
		mockResponse   entities.User
		mockError      error
		configureMock  func(*userServiceMock, entities.User, error)
		expectedOutput entities.User
		expectedError  error
	}{
		{
			testName:     "TestModifyUserService",
			mock:         &userServiceMock{},
			mockResponse: entities.User{ID: "3", Password: "12345678", Age: "20", Information: "add", Parents: "idk", Email: "alexer@gmail.com", Name: "Alexer Maestre"},
			mockError:    nil,
			configureMock: func(m *userServiceMock, mockResponse entities.User, mockError error) {
				m.On("UpdateUser", mock.Anything).Return(mockResponse, mockError)
			},
			expectedOutput: entities.User{ID: "3", Password: "12345678", Age: "20", Information: "add", Parents: "idk", Email: "alexer@gmail.com", Name: "Alexer Maestre"},
			expectedError:  nil,
		},
	}

	for _, tt := range testScenarios {
		t.Run(tt.testName, func(t *testing.T) {

			// Prepare
			if tt.configureMock != nil {
				tt.configureMock(tt.mock, tt.mockResponse, tt.mockError)
			}

			service := &userService{
				repository: tt.mock,
			}

			// Act
			result, err := service.UpdateUser(tt.mockResponse)

			// Assert
			assert.Equal(t, tt.expectedError, err)
			assert.Equal(t, tt.expectedOutput, result)
		})
	}

}

func TestNewUserService(t *testing.T) {

	sm := &userServiceMock{}
	testScenarios := []struct {
		testName       string
		mockRepo       mysql.UserRepository
		expectedOutput *userService
	}{
		{
			testName: "test MakeServerEndpoints",
			mockRepo: sm,
			expectedOutput: &userService{
				repository: sm,
			},
		},
	}

	for _, tt := range testScenarios {

		// Prepare
		t.Run(tt.testName, func(t *testing.T) {

			// Act
			result := NewUserService(tt.mockRepo)

			// Assert
			assert.NotNil(t, result)
			assert.Equal(t, tt.expectedOutput.repository, result.repository)
		})
	}
}
