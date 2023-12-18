package api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type Server struct {
	router *gin.Engine
	db     *gorm.DB
}

func NewServer(db *gorm.DB) *Server {
	server := &Server{
		router: gin.Default(),
		db:     db,
	}

	server.router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "Welcome to our shop!")
	})

	return server
}

func (s *Server) Start(address string) error {
	return s.router.Run(address)
}
