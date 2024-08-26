package endpoints

import (
	"bootcamp_api/api/models"
	"bootcamp_api/api/services"
	"context"
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
		configureMock   func(*serviceMock, models.User)
		endpointRequest interface{}
		mockResponse    models.User
		expectedOutput  GetUserResponse
		expetedError    error
	}{
		{
			testName: "test  get user",
			mock:     &serviceMock{},
			mockResponse: models.User{
				ID: "3",
			},
			configureMock: func(m *serviceMock, mockResponse models.User) {
				m.On("GetUser", mock.Anything).Return(mockResponse, nil)
			},
			expectedOutput: GetUserResponse{User: models.User{
				ID: "3",
			}},
			expetedError:    nil,
			endpointRequest: GetUserRequest{ID: "3"},
		},
	}
	for _, tt := range testScenarios {
		t.Run(tt.testName, func(t *testing.T) {

			// Prepare
			tt.endpoint = MakeGetUserEndpoint
			if tt.configureMock != nil {
				tt.configureMock(tt.mock, tt.mockResponse)
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
