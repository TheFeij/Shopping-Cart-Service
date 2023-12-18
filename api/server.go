package api

import (
	"Shoping-Cart-Service/api/handlers"
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

	handler := handlers.NewBasketsHandler(server.db)
	server.router.GET("/baskets", handler.GetBasketsList)
	server.router.GET("/basket/:id", handler.GetBasket)
	server.router.POST("/basket", handler.CreateNewBasket)
	server.router.PATCH("/basket/:id", handler.UpdateBasket)
	server.router.DELETE("/basket/:id", handler.DeleteBasket)

	return server
}

func (s *Server) Start(address string) error {
	return s.router.Run(address)
}
