package api

import "github.com/gorilla/mux"

func Router(s *server) *mux.Router {

	r := s.router

	r.HandleFunc("/create-user", s.CreateUserHandler).Methods("POST")
	r.HandleFunc("/create-product", s.CreateProductHandler).Methods("POST")

	return r

}
