package domain

type OrderEntity struct {
	ID uint `gorm:"type:uuid;autoIncrement;primaryKey"`

	ChargeID  *string `gorm:"type:varchar(255);unique"`
	InvoiceNo *string `gorm:"type:varchar(255);unique"`

	Vat         *float64 `gorm:"type:decimal(5,2)"`
	Amount      float64  `gorm:"type:decimal(10,2);default:0"`
	TotalAmount float64  `gorm:"type:decimal(10,2);default:0"`

	OrderStatus   *string `gorm:"type:varchar(50)"`
	PaymentMethod *string `gorm:"type:string(50)"`

	SeatLabel *string `gorm:"type:varchar(255)"`
}

type OrderRepository interface {
	Create(order *OrderEntity) error
	Update(order *OrderEntity) error
	GetByUserId(userId uint) ([]OrderEntity, error)
}
