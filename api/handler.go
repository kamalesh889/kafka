package api

import (
	"encoding/json"
	"net/http"
)

func (s *server) start(w http.ResponseWriter, r *http.Request) {

}

func (s *server) createUser(w http.ResponseWriter, r *http.Request) {

	var user UserRequest
	var userresp UserResponse

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userid, err := s.service.CreateUser(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	userresp = UserResponse{
		Name: user.Name,
		Id:   userid,
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

	err = s.service.CreateProduct(&product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
