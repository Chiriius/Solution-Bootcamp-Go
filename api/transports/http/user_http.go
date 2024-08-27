package adapters

import (
	"bootcamp_api/api/endpoints"
	"context"
	"encoding/json"
	
	"net/http"

	httpTransport "github.com/go-kit/kit/transport/http"
)

func NewHTTPHandler(endpointss endpoints.Endpoints) http.Handler {

	m := http.NewServeMux()
	m.Handle("/user", httpTransport.NewServer(
		endpointss.GetUser,
		decodeGerUserRequest,
		encodeGetUserResponse,
	))
	m.Handle("/user/create", httpTransport.NewServer(
		endpointss.AddUser,
		decodeAddUserRequest,
		encodeGetUserResponse,
	))
	return m
}

func decodeGerUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.GetUserRequest
	if err := r.ParseForm(); err != nil {
		return nil, err
	}
	req.ID = r.FormValue("id")
	return req, nil
}

func encodeGetUserResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {

	// var req endpoints.GetUserResponse
	// if req != response {
	// 	return errors.New("Diferent type")
	// }
	//Agregar codigo de respuesta
	//Agregar validaciones
	return json.NewEncoder(w).Encode(response)

}

func decodeAddUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}
