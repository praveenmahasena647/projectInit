package handlers

import (
	"fmt"

	routing "github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
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
	var r = routing.New()

	r.Get("/app", func(ctx *routing.Context) error {
		fmt.Fprintln(ctx, map[string]string{
			"greetings": "hello",
		})
		return nil
	})

	return fasthttp.ListenAndServe(s.listenAddr, r.HandleRequest)
}
