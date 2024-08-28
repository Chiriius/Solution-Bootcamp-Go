package endpoints

import (
	"bootcamp_api/api/entities"
	"bootcamp_api/api/services"
	errorss "bootcamp_api/api/utils/errors"
	"context"
	"log"

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
			log.Printf(errorss.ErrorInterfaceDifType.Message)
			return nil, errorss.ErrorInterfaceDifType
		}
		user, err := s.GetUser(req.ID)
		if err != nil {
			log.Println(err.Error())
			return GetUserResponse{}, err
		}
		log.Printf("Obtained user: %s sucessfully", req.ID)
		return GetUserResponse{User: user}, nil

	}
}

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
			log.Println(errorss.ErrorInterfaceDifType)
			return nil, errorss.ErrorInterfaceDifType
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
			log.Println(err)
			return CreateUserResponse{}, err
		}
		log.Printf("Created user with id:%s sucessfully ", serviceUser.ID)
		return CreateUserResponse{Id: serviceUser.ID}, nil

	}

}
