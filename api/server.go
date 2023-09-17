package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

type server struct {
	router     *mux.Router
	db         *Database
	httpClient *http.Client
}

func NewServer(db *Database) (*server, error) {

	s := &server{}

	s.router = mux.NewRouter()
	s.db = db
	s.httpClient = &http.Client{}

	return s, nil

}
