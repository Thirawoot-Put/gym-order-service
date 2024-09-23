package payment

import "github.com/google/uuid"

type Order struct {
	OrderID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey;column:order_id"`

	OrderCode *string `gorm:"type:varchar(255);column:order_code"`
	ChargeID  *string `gorm:"type:varchar(255);unique;column:charge_id"`
	InvoiceNo *string `gorm:"type:varchar(255);unique;column:invoice_no"`

	Amount      float64 `gorm:"type:decimal(10,2);column:amount;default:0"`
	TotalAmount float64 `gorm:"type:decimal(10,2);column:total_amount;default:0"`

	OrderStatus   *string `gorm:"type:varchar(50);column:order_status"`
	PaymentMethod *string `gorm:"type:string(50);column:payment_method"`
	ReservedCount int     `gorm:"default:0;column:reserved_count"`
	SeatLabel     *string `gorm:"type:varchar(255);column:seat_label"`

	Vat *float64 `gorm:"type:decimal(5,2);column:vat"`
}
