package repository

import (
	"github.com/Thirawoot-Put/event-ticketing/payment-service/internal/domain"
	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) Create(order *domain.OrderEntity) error {
	return r.db.Create(order).Error
}

func (r *OrderRepository) Update(order *domain.OrderEntity) error {
	return r.db.Model(&order).Updates(order).Error
	// return r.db.Save(order).Error
}

func (r *OrderRepository) GetByUserId(userId uint) ([]domain.OrderEntity, error) {
	var orders []domain.OrderEntity
	err := r.db.Find(&orders, r.db.Where("userId = ?", userId)).Error
	return orders, err
}
