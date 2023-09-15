package api

import "github.com/gorilla/mux"

func Router(s *server) *mux.Router {

	r := s.router

	r.HandleFunc("/start", s.start).Methods("GET")

	return r

}
