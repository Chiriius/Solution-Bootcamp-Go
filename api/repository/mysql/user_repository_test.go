package mysql

import (
	"bootcamp_api/api/entities"
	errorss "bootcamp_api/api/utils/errors"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestGetUserById(t *testing.T) {
	testScenarios := []struct {
		testName       string
		mockResponse   entities.User
		mockError      error
		configureMock  func(sqlmock.Sqlmock)
		expectedOutput entities.User
		expectedError  error
	}{
		{
			testName: "TestGetUser",
			mockResponse: entities.User{
				ID:          "1",
				Name:        "Sebas",
				Email:       "sebas@gmail.com",
				Password:    "sds",
				Age:         "34",
				Information: "asda",
				Parents:     "sds",
			},
			mockError: nil,
			configureMock: func(m sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "name", "email", "password", "age", "information", "parents"}).
					AddRow("1", "Sebas", "sebas@gmail.com", "sds", "34", "asda", "sds")
				(m).ExpectQuery(`SELECT \* FROM users WHERE id=\?`).
					WithArgs("1").
					WillReturnRows(rows)
			},
			expectedOutput: entities.User{
				ID:          "1",
				Name:        "Sebas",
				Email:       "sebas@gmail.com",
				Password:    "sds",
				Age:         "34",
				Information: "asda",
				Parents:     "sds",
			},
			expectedError: nil,
		},
		{
			testName:     "TestGetUser with error",
			mockResponse: entities.User{},
			mockError:    errorss.ErrorUserNotFound,
			configureMock: func(m sqlmock.Sqlmock) {
				(m).ExpectQuery(`SELECT \* FROM users WHERE id=\?`).
					WithArgs("1").
					WillReturnError(sql.ErrNoRows)
			},
			expectedOutput: entities.User{},
			expectedError:  errorss.ErrorUserNotFound,
		},
	}

	for _, tt := range testScenarios {
		t.Run(tt.testName, func(t *testing.T) {

			// Prepare
			db, mock, err := sqlmock.New()
			assert.NoError(t, err)
			defer db.Close()
			sqlxDB := sqlx.NewDb(db, "mysql")
			repo := NewMySQLUserRepository(sqlxDB, nil)

			if tt.configureMock != nil {
				tt.configureMock(mock)
			}

			// Act
			result, err := repo.GetUser("1")

			// Assert
			assert.Equal(t, tt.expectedError, err)
			assert.Equal(t, tt.expectedOutput, result)
		})
	}
}

func TestAddUser(t *testing.T) {
	testScenarios := []struct {
		testName       string
		inputUser      entities.User
		mockResponse   entities.User
		mockError      error
		configureMock  func(sqlmock.Sqlmock)
		expectedOutput entities.User
		expectedError  error
	}{
		{
			testName: "TestAddUser_Success",
			inputUser: entities.User{
				ID:          "1",
				Name:        "Sebas",
				Email:       "sebas@gmail.com",
				Password:    "sds",
				Age:         "34",
				Information: "asda",
				Parents:     "sds",
			},
			mockResponse: entities.User{
				ID:          "1",
				Name:        "Sebas",
				Email:       "sebas@gmail.com",
				Password:    "sds",
				Age:         "34",
				Information: "asda",
				Parents:     "sds",
			},
			mockError: nil,
			configureMock: func(m sqlmock.Sqlmock) {
				m.ExpectExec(`^INSERT INTO users \(id, password, age, information, parents, email , name \) VALUES\( \?, \?, \?, \?, \?, \?, \?\)$`).
					WithArgs("1", "sds", "34", "asda", "sds", "sebas@gmail.com", "Sebas").
					WillReturnResult(sqlmock.NewResult(1, 1)) // Assuming the insert affects one row
			},
			expectedOutput: entities.User{
				ID:          "1",
				Name:        "Sebas",
				Email:       "sebas@gmail.com",
				Password:    "sds",
				Age:         "34",
				Information: "asda",
				Parents:     "sds",
			},
			expectedError: nil,
		},
		// {
		// 	testName: "TestAddUser with error",
		// 	inputUser: entities.User{
		// 		ID:          "1",
		// 		Name:        "Sebas",
		// 		Email:       "sebas@gmail.com",
		// 		Password:    "sds",
		// 		Age:         "34",
		// 		Information: "asda",
		// 		Parents:     "sds",
		// 	},
		// 	mockError: errors.New("Codigo:505 Message:Server error"),
		// 	configureMock: func(m sqlmock.Sqlmock) {
		// 		m.ExpectExec(`^INSERT INTO users \(id, password, age, information, parents, email , name \) VALUES\( \?, \?, \?, \?, \?, \?, \?\)$`).
		// 			WithArgs("1", "sds", "34", "asda", "sds", "sebas@gmail.com", "Sebas").
		// 			WillReturnError(errors.New("Codigo:505 Message:Server error"))
		// 	},
		// 	expectedOutput: entities.User{},
		// 	expectedError:  errors.New("Codigo:505 Message:Server error"),
		// },
	}

	for _, tt := range testScenarios {
		t.Run(tt.testName, func(t *testing.T) {

			// Prepare
			db, mock, err := sqlmock.New()
			assert.NoError(t, err)
			defer db.Close()
			sqlxDB := sqlx.NewDb(db, "mysql")
			repo := NewMySQLUserRepository(sqlxDB, nil)

			if tt.configureMock != nil {
				tt.configureMock(mock)
			}

			// Act
			result, err := repo.AddUser(tt.inputUser)

			// Assert
			assert.Equal(t, tt.expectedError, err)
			assert.Equal(t, tt.expectedOutput, result)
		})
	}
}

func TestUpdateUser(t *testing.T) {
	testScenarios := []struct {
		testName       string
		mockResponse   entities.User
		mockError      error
		configureMock  func(sqlmock.Sqlmock)
		expectedOutput entities.User
		expectedError  error
	}{
		{
			testName: "TestUpdateUser",
			mockResponse: entities.User{
				ID:          "1",
				Name:        "Sebas",
				Email:       "sebas@gmail.com",
				Password:    "sds",
				Age:         "34",
				Information: "asda",
				Parents:     "sds",
			},
			mockError: nil,
			configureMock: func(m sqlmock.Sqlmock) {
				m.ExpectExec(`^UPDATE users SET password = \?, age = \?, information = \?, parents = \?, email = \?, name = \? WHERE id = \?$`).
					WithArgs("sds", "34", "asda", "sds", "sebas@gmail.com", "Sebas", "1").
					WillReturnResult(sqlmock.NewResult(1, 1)) // Assuming the update affects one row
			},
			expectedOutput: entities.User{
				ID:          "1",
				Name:        "Sebas",
				Email:       "sebas@gmail.com",
				Password:    "sds",
				Age:         "34",
				Information: "asda",
				Parents:     "sds",
			},
			expectedError: nil,
		},
	}

	for _, tt := range testScenarios {
		t.Run(tt.testName, func(t *testing.T) {

			// Prepare
			db, mock, err := sqlmock.New()
			assert.NoError(t, err)
			defer db.Close()
			sqlxDB := sqlx.NewDb(db, "mysql")
			repo := NewMySQLUserRepository(sqlxDB, nil)

			if tt.configureMock != nil {
				tt.configureMock(mock)
			}

			// Act
			result, err := repo.UpdateUser(tt.mockResponse)

			// Assert
			assert.Equal(t, tt.expectedError, err)
			assert.Equal(t, tt.expectedOutput, result)
		})
	}
}

// FORMA QUE NO SIRVE YA QUE NO TESTEA LOS QUERYS
func TestGetUserByIdd(t *testing.T) {
	testScenarios := []struct {
		testName       string
		mock           *userRepositoryMock
		mockResponse   entities.User
		mockError      error
		configureMock  func(*userRepositoryMock, entities.User, error)
		expectedOutput entities.User
		expectedError  error
	}{
		{
			testName: "TestGetUser",
			mock:     &userRepositoryMock{},
			mockResponse: entities.User{
				ID:          "1",
				Name:        "Sebas",
				Email:       "sebas@gmail.com",
				Password:    "sds",
				Age:         "34",
				Information: "asda",
				Parents:     "sds",
			},
			mockError: nil,
			configureMock: func(m *userRepositoryMock, mockResponse entities.User, mockError error) {
				m.On("GetUser", "1").Return(mockResponse, mockError)
			},
			expectedOutput: entities.User{
				ID:          "1",
				Name:        "Sebas",
				Email:       "sebas@gmail.com",
				Password:    "sds",
				Age:         "34",
				Information: "asda",
				Parents:     "sds",
			},
			expectedError: nil,
		},
		{
			testName:     "TestGetUser with error",
			mock:         &userRepositoryMock{},
			mockResponse: entities.User{},
			mockError:    errorss.ErrorUserNotFound,
			configureMock: func(m *userRepositoryMock, mockResponse entities.User, mockError error) {
				m.On("GetUser", "1").Return(mockResponse, mockError)
			},
			expectedOutput: entities.User{},
			expectedError:  errorss.ErrorUserNotFound,
		},
	}

	for _, tt := range testScenarios {
		t.Run(tt.testName, func(t *testing.T) {

			// Prepare
			if tt.configureMock != nil {
				tt.configureMock(tt.mock, tt.mockResponse, tt.mockError)
			}

			// Act
			result, err := UserRepository.GetUser(tt.mock, "1")

			// Assert
			assert.Equal(t, tt.expectedError, err)
			assert.Equal(t, tt.expectedOutput, result)
		})
	}
}
