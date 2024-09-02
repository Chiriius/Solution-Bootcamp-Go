package server

import (
	"bootcamp_api/api/endpoints"
	"bootcamp_api/api/repository/mysql"
	"bootcamp_api/api/services"
	adapters "bootcamp_api/api/transports/http"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Server struct {
	db  *sqlx.DB
	mux *http.ServeMux
}

func New() (*Server, error) {
	db, err := sqlx.Connect("mysql", "root:12345678@tcp(127.0.0.1:3306)/bootcampgo")
	if err != nil {
		return nil, err
	}

	userRepository := mysql.NewMySQLUserRepository(db)
	userService := services.NewUserService(userRepository)
	userEndpoints := endpoints.MakeServerEndpoints(userService)
	httpHandler := adapters.NewHTTPHandler(userEndpoints)

	mux := http.NewServeMux()
	mux.Handle("/", httpHandler)

	return &Server{
		db:  db,
		mux: mux,
	}, nil
}

func (s *Server) Start(addr string) error {
	return http.ListenAndServe(addr, s.mux)
}

func (s *Server) Close() error {
	return s.db.Close()
}
