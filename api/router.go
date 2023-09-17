package api

import "github.com/gorilla/mux"

func Router(s *server) *mux.Router {

	r := s.router

	r.HandleFunc("/start", s.start).Methods("GET")
	r.HandleFunc("/create-user", s.createUser).Methods("POST")
	r.HandleFunc("/crate-product", s.createProduct).Methods("POST")

	return r

}
