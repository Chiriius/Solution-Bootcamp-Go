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
		encodeGenericResponse,
	))
	m.Handle("/user/create", httpTransport.NewServer(
		endpointss.AddUser,
		decodeAddUserRequest,
		encodeGenericResponse,
	))
	m.Handle("/user/edit", httpTransport.NewServer(
		endpointss.UpdateUser,
		decodeModifyRequest,
		encodeGenericResponse,
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

func decodeModifyRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.ModifyUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}
func encodeGenericResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {

	return json.NewEncoder(w).Encode(response)

}

func decodeAddUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}
