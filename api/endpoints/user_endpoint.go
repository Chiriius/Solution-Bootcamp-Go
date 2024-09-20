package endpoints

import (
	"bootcamp_api/api/entities"
	"bootcamp_api/api/services"
	errorss "bootcamp_api/api/utils/errors"
	"context"
	"fmt"
	"log"

	"github.com/go-kit/kit/endpoint"
	"github.com/sirupsen/logrus"
)

type GetUserRequest struct {
	ID string
}

type GetUserResponse struct {
	User entities.User
	Err  string `json:"error,omitempty"`
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

type ModifyUserRequest struct {
	Id          string
	Password    string
	Age         string
	Information string
	Parents     string
	Email       string
	Name        string
}

type ModifyUserResponse struct {
	Id  string `json:"id,omitempty"`
	Err string `json:"error,omitempty"`
}

type Endpoints struct {
	GetUser    endpoint.Endpoint
	AddUser    endpoint.Endpoint
	UpdateUser endpoint.Endpoint
	logger     logrus.FieldLogger
}

func MakeServerEndpoints(s services.UserService) Endpoints {
	return Endpoints{
		GetUser:    MakeGetUserEndpoint(s),
		AddUser:    MakeAddUserEndpoint(s),
		UpdateUser: MakeUpdateUserEndpoint(s),
	}
}

func MakeGetUserEndpoint(s services.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		var req GetUserRequest
		var ok bool = false
		if req, ok = request.(GetUserRequest); !ok {
			//.Errorln(`Layer:user_endpoint Method:MakeGetUser,`, err)
			return nil, errorss.ErrorInterfaceDifType
		}
		user, err := s.GetUser(req.ID)
		if err != nil {
			//l.Errorln(`Layer:user_endpoint Method:MakeGetUser,`, err)
			return GetUserResponse{}, err
		}
		//l.Infoln(`Layer:user_endpoint Method:MakeGetUser, user:`, user)
		return GetUserResponse{User: user}, nil

	}
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

func MakeUpdateUserEndpoint(s services.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		var req ModifyUserRequest
		var ok bool
		if req, ok = request.(ModifyUserRequest); !ok {
			log.Println(errorss.ErrorInterfaceDifType)
			return nil, errorss.ErrorInterfaceDifType
		}

		userG, err := s.GetUser(req.Id)
		fmt.Println(userG)
		if err != nil {
			log.Println(err.Error())
			return ModifyUserResponse{}, err
		}

		user := entities.User{
			ID:          req.Id,
			Password:    req.Password,
			Age:         req.Age,
			Information: req.Information,
			Parents:     req.Parents,
			Email:       req.Email,
			Name:        req.Name,
		}

		serviceUser, err := s.UpdateUser(user)
		if err != nil {
			log.Println(err)
			return ModifyUserResponse{}, err
		}
		log.Printf("Modified user with id:%s sucessfully ", serviceUser.ID)
		return ModifyUserResponse{Id: serviceUser.ID}, nil

	}
}
