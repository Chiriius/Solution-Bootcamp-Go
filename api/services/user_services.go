package services

import (
	"bootcamp_api/api/entities"
	"bootcamp_api/api/repository/mysql"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type UserService interface {
	GetUser(id string) (entities.User, error)
	AddUser(user entities.User) (entities.User, error)
	UpdateUser(user entities.User) (entities.User, error)
}

type userService struct {
	repository mysql.UserRepository
	logger     logrus.FieldLogger
}

func NewUserService(repo mysql.UserRepository, logger logrus.FieldLogger) *userService {
	return &userService{
		repository: repo,
		logger:     logger,
	}
}

func (s *userService) GetUser(id string) (entities.User, error) {
	return s.repository.GetUser(id)
}

func (s *userService) AddUser(user entities.User) (entities.User, error) {

	user.ID = uuid.NewString()

	return s.repository.AddUser(user)
}

func (s *userService) UpdateUser(user entities.User) (entities.User, error) {

	return s.repository.UpdateUser(user)
}
