package endpoints

import (
	"bootcamp_api/api/entities"
	"bootcamp_api/api/services"
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
)

type GetUserRequest struct {
	ID string
}

type GetUserResponse struct {
	User entities.User
	Err  string `json:"error,omitempty"`
}

type Endpoints struct {
	GetUser endpoint.Endpoint
	AddUser endpoint.Endpoint
}

func MakeServerEndpoints(s services.UserService) Endpoints {
	return Endpoints{
		GetUser: MakeGetUserEndpoint(s),
		AddUser: MakeAddUserEndpoint(s),
	}
}

func MakeGetUserEndpoint(s services.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		var req GetUserRequest
		var ok bool = false
		if req, ok = request.(GetUserRequest); !ok {
			return nil, errors.New("sds") //
		}
		user, err := s.GetUser(req.ID)
		if err != nil {
			return GetUserResponse{User: user, Err: err.Error()}, nil
		}
		return GetUserResponse{User: user}, nil
	}
}

// ESTUDIAR CLOSURE EN GO
// CASTING
// Crear carpeta entity con la entity user
type CreateUserRequest struct {
	Password    string
	Age         string
	Information string
	Parents     string
	Email       string
	Name        string
}

type CreateUserResponse struct {
	Id  string `json:"id,omitempty"`
	Err string `json:"error,omitempty"`
}

func MakeAddUserEndpoint(s services.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var req CreateUserRequest
		var ok bool = false
		if req, ok = request.(CreateUserRequest); !ok {
			return nil, errors.New("ss")
		}
		user := entities.User{
			Password:    req.Password,
			Age:         req.Age,
			Information: req.Information,
			Parents:     req.Parents,
			Email:       req.Email,
			Name:        req.Name,
		}
		serviceUser, err := s.AddUser(user)
		if err != nil {
			return CreateUserResponse{Err: err.Error()}, nil
		}
		return CreateUserResponse{Id: serviceUser.ID}, nil

	}

}
