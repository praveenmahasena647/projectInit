package handlers

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
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
	var ch = chi.NewRouter()

	ch.Get("/app", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, map[string]string{
			"greetings": "hello",
		})
	})

	return http.ListenAndServe(s.listenAddr, ch)
}
