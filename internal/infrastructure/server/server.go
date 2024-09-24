package server

import (
	"fmt"
	"net/http"

	"github.com/Thirawoot-Put/event-ticketing/payment-service/internal/handler/api"
	"github.com/Thirawoot-Put/event-ticketing/payment-service/internal/infrastructure/database"
	"github.com/Thirawoot-Put/event-ticketing/payment-service/internal/repository"
	"github.com/Thirawoot-Put/event-ticketing/payment-service/internal/service"
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
	db := database.Connecting()

	orderReposetory := repository.NewOrderRepository(db)
	orderService := service.NewOrderService(orderReposetory)
	orderHandler := api.NewOrderHandler(orderService)

	s.app.GET("/health-check", healthCheck)

	// order route
	s.app.POST("/orders", orderHandler.CreateOrder)
	s.app.PATCH("/orders", orderHandler.UpdateOrder)
	s.app.GET("/orders/users/:id", orderHandler.GetOrderByUserId)

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
