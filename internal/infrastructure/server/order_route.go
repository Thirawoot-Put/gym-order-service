package server

import (
	"github.com/Thirawoot-Put/event-ticketing/order-service/internal/handler/api"
	"github.com/Thirawoot-Put/event-ticketing/order-service/internal/repository"
	"github.com/Thirawoot-Put/event-ticketing/order-service/internal/service"
	"gorm.io/gorm"
)

func (s *Server) initOrderRoute(db *gorm.DB) {
	orderRoute := s.app.Group("/api/v1/orders")

	orderReposetory := repository.NewOrderRepository(db)
	ervice := service.NewOrderService(orderReposetory)
	orderHandler := api.NewOrderHandler(ervice)

	orderRoute.POST("", orderHandler.CreateOrder)
	orderRoute.PATCH("", orderHandler.UpdateOrder)
	orderRoute.GET("/users/:id", orderHandler.GetOrderByUserId)

}
