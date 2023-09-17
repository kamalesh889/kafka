package api

import (
	"encoding/json"
	"net/http"
)

func (s *server) start(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("To be implemented"))
}

func (s *server) createUser(w http.ResponseWriter, r *http.Request) {

	var user UserRequest
	var userresp UserResponse

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userresp = UserResponse{
		Name: user.Name,
		Id:   1,
	}

	json.NewEncoder(w).Encode(userresp)
	w.WriteHeader(http.StatusCreated)

}

func (s *server) createProduct(w http.ResponseWriter, r *http.Request) {

	var product ProductRequest

	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
