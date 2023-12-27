package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
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
	var e = echo.New()
	e.GET("/app", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"greetings": "hello",
		})
	})
	return e.Run(s.listenAddr)
}
