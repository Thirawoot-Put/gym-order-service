package service

import (
	"fmt"

	"github.com/Thirawoot-Put/event-ticketing/payment-service/internal/domain"
	"github.com/Thirawoot-Put/event-ticketing/payment-service/internal/repository"
)

type OrderService struct {
	repo *repository.OrderRepository
}

func NewOrderService(repo *repository.OrderRepository) *OrderService {
	return &OrderService{repo: repo}
}

// may be can receive dto
func (s *OrderService) CreateOrder(order *domain.Orders) error {
	// add business logic here
	return s.repo.Create(order)
}

func (s *OrderService) UpdateOrder(order *domain.Orders) error {
	// add business logic here
	return s.repo.Update(order)
}

func (s *OrderService) GetOrderByUserId(userId uint) ([]domain.Orders, error) {
	// add business logic here
	orders, err := s.repo.GetByUserId(userId)
	if err != nil {
		errMsg := fmt.Errorf("Error to get orders of %d, error: %w", userId, err)
		return nil, errMsg
	}

	return orders, nil
}
