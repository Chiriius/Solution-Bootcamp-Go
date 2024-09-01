package services

import (
	"bootcamp_api/api/entities"
	"bootcamp_api/api/repository/mysql"

	"github.com/google/uuid"
)

type UserService interface {
	GetUser(id string) (entities.User, error)
	AddUser(user entities.User) (entities.User, error)
	UpdateUser(user entities.User) (entities.User, error)
}

type userService struct {
	repository mysql.UserRepository
}

func NewUserService(repo mysql.UserRepository) *userService {
	return &userService{repository: repo}
}

func (s *userService) GetUser(id string) (entities.User, error) { // reemplazar el model.user duplicando la struct
	return s.repository.GetUser(id)
}

func (s *userService) AddUser(user entities.User) (entities.User, error) { // reemplazar el model.user duplicando la struct

	user.ID = uuid.NewString()
	return s.repository.AddUser(user)
}

func (s *userService) UpdateUser(user entities.User) (entities.User, error) {

	return s.repository.UpdateUser(user)
}
