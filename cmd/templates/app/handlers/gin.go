package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	return r.Run(s.listenAddr)
}
