package adapter

import (
	"bootcamp_api/api/endpoints"
	"bootcamp_api/api/transports/grpc/pb"
	"context"

	gt "github.com/go-kit/kit/transport/grpc"
)

type gRPCServer struct {
	getUser gt.Handler
	addUser gt.Handler
}

func NewGRPCServer(endpoints endpoints.Endpoints) pb.UserServiceServer {
	return &gRPCServer{
		getUser: gt.NewServer(
			endpoints.GetUser,
			decodeUserRequest,
			encodeUserResponse,
		),
	}
}

func (s *gRPCServer) mustEmbedUnimplementedUserServiceServer() {}

func (s *gRPCServer) AddUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	_, resp, err := s.addUser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.CreateUserResponse), nil
}

func (s *gRPCServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	_, resp, err := s.getUser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.GetUserResponse), nil
}

func decodeUserRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.GetUserRequest)
	return endpoints.GetUserRequest{ID: req.Id}, nil
}

func encodeUserResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(endpoints.GetUserResponse)
	user := &pb.User{Id: res.User.ID, Password: res.User.Password, Age: res.User.Age, Email: res.User.Email, Information: res.User.Information, Name: res.User.Name, Parents: res.User.Parents}
	return &pb.GetUserResponse{User: user}, nil
}
