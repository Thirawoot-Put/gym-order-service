package domain

import "time"

type Orders struct {
	ID uint `gorm:"type:uuid;autoIncrement;primaryKey;column:id"`

	ChargeID  *string `gorm:"type:varchar(255);unique;column:charge_id"`
	InvoiceNo *string `gorm:"type:varchar(255);unique;column:invoice_no"`

	Vat         *float64 `gorm:"type:decimal(5,2);column:vat"`
	Service     *float64 `gorm:"type:decimal(5,2);column:service"`
	Amount      float64  `gorm:"type:decimal(10,2);default:0;column:amount"`
	TotalAmount float64  `gorm:"type:decimal(10,2);default:0;column:total_amount"`

	OrderStatus   *string `gorm:"type:varchar(50);column:order_status"`
	PaymentMethod *string `gorm:"type:string(50);column:payment_method"`

	SeatLabel *string `gorm:"type:varchar(255);column:seat_label"`

	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

// type OrderRepository interface {
// 	Create(order *OrderEntity) error
// 	Update(order *OrderEntity) error
// 	GetByUserId(userId uint) ([]OrderEntity, error)
// }
