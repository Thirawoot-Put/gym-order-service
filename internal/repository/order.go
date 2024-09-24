package repository

import (
	"github.com/Thirawoot-Put/event-ticketing/payment-service/internal/domain"
	"gorm.io/gorm"
)

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *orderRepository {
	return &orderRepository{db: db}
}

func (r *orderRepository) Create(order *domain.OrderRepository) error {
	return r.db.Create(order).Error
}

func (r *orderRepository) Update(order *domain.OrderRepository) error {
	return r.db.Save(order).Error
}

func (r *orderRepository) GetByUserId(userId uint) ([]domain.OrderEntity, error) {
	var orders []domain.OrderEntity
	err := r.db.Find(&orders, r.db.Where("userId = ?", userId)).Error
	return orders, err
}
