package services

import (
    "bootcamp_api/api/models"
    "bootcamp_api/api/repository/mysql"
    "github.com/google/uuid"
)

type UserService interface{
    GetUser(id string) (models.User, error)
    AddUser(user models.User) error
}


type userService struct{
    repository mysql.UserRepository
}

func NewUserService(repo mysql.UserRepository) *userService{
    return &userService{repository: repo}
}

func (s *userService) GetUser(id string) (models.User, error) { // reemplazar el model.user duplicando la struct
     return s.repository.GetUserById(id)

}

func (s *userService) AddUser(user models.User) error{ // reemplazar el model.user duplicando la struct
    
    user.ID = uuid.NewString()
    return s.repository.AddUser(user)
}