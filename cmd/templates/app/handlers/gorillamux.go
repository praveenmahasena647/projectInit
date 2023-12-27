package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	listenAddr string
}

func New(port string) *Server {
	return &Server{
		listenAddr: port,
	}
}

func (s *Server) Init() error {
	var m = mux.NewRouter()

	m.HandleFunc("/app", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, map[string]string{
			"greetings": "hello",
		})
	})

	return http.ListenAndServe(s.listenAddr, m)
}
