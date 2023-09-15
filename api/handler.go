package api

import "net/http"

func (s *server) start(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("To be implemented"))
}
