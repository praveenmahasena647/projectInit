package handlers

import (
	"fmt"
	"net/http"
)

type Server struct {
	listAddr string
}

func New(port string) *Server {
	return &Server{
		listAddr: port,
	}
}

func (s *Server) Init() error {
	http.HandleFunc("/app", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, map[string]string{
			"greetings": "hello",
		})
	})
	return http.ListenAndServe(s.listAddr, nil)
}
