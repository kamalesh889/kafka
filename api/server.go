package api

import "github.com/gorilla/mux"

type server struct {
	router *mux.Router
}

func NewServer() (*server, error) {

	s := &server{}

	s.router = mux.NewRouter()

	return s, nil

}
