package server

import (
	"fmt"
	"net/http"

	"github.com/Thirawoot-Put/event-ticketing/payment-service/internal/infrastructure/database"
	"github.com/gin-gonic/gin"
)

type Server struct {
	app *gin.Engine
}

func AppServer() *Server {
	ginApp := gin.Default()

	return &Server{
		app: ginApp,
	}
}

func (s *Server) Start(port string) {
	database.Connecting()

	s.app.GET("/health-check", healthCheck)

	url := fmt.Sprintf(":%s", port)
	s.HTTPListenPort(url)
}

func (s *Server) HTTPListenPort(port string) {
	err := s.app.Run()
	if err != nil {
		panic("Failed to start server")
	}
}

func healthCheck(ctx *gin.Context) {
	m := map[string]string{"message": "App healthy"}
	ctx.JSON(http.StatusOK, m)
}
