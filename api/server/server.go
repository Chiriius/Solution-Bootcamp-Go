package server

import (
	"bootcamp_api/api/endpoints"
	"bootcamp_api/api/repository/mysql"
	"bootcamp_api/api/services"
	adapter "bootcamp_api/api/transports/grpc"
	"bootcamp_api/api/transports/grpc/pb"
	adapters "bootcamp_api/api/transports/http"

	"log"
	"net"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type Server struct {
	db       *sqlx.DB
	httpMux  *http.ServeMux
	grpcSrv  *grpc.Server
	grpcAddr string
	httpAddr string
	logger   logrus.FieldLogger
}

func New(logger logrus.FieldLogger, httpAddr, grpcAddr, dbUrl string) (*Server, error) {

	db, err := sqlx.Connect("mysql", dbUrl)
	if err != nil {
		logger.Panic("Layer:Server, Error al conectar la db:", err, " esta es la url:", dbUrl)
		return nil, err
	}

	userRepository := mysql.NewMySQLUserRepository(db, logger)
	userService := services.NewUserService(userRepository, logger)
	userEndpoints := endpoints.MakeServerEndpoints(userService)
	httpHandler := adapters.NewHTTPHandler(userEndpoints)
	grpcServerr := adapter.NewGRPCServer(userEndpoints)

	httpMux := http.NewServeMux()
	httpMux.Handle("/", httpHandler)

	baseServer := grpc.NewServer()
	pb.RegisterUserServiceServer(baseServer, grpcServerr)

	return &Server{
		db:       db,
		httpMux:  httpMux,
		grpcSrv:  baseServer,
		grpcAddr: grpcAddr,
		httpAddr: httpAddr,
	}, nil
}

func (s *Server) Start() error {
	go func() {
		log.Printf("HTTP server listening on %s", s.httpAddr)
		if err := http.ListenAndServe(s.httpAddr, s.httpMux); err != nil {
			log.Fatalf("HTTP server failed: %v", err)
		}
	}()

	lis, err := net.Listen("tcp", s.grpcAddr)
	if err != nil {
		return err
	}
	log.Printf("gRPC server listening on %s", s.grpcAddr)
	return s.grpcSrv.Serve(lis)
}

func (s *Server) Close() error {
	s.grpcSrv.GracefulStop()
	return s.db.Close()
}
