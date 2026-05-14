package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Payment struct {
	ID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`

	BookingID uuid.UUID `gorm:"notnull;index:idx_booking_id" json:"booking_id"`
	Booking   Booking   `gorm:"foreignKey:BookingID;references:ID;constraint:OnDelete:RESTRICT;" json:"booking"`

	PaymentMethod string          `gorm:"type:varchar(50);notnull" json:"payment_method"`
	FinalAmount   decimal.Decimal `gorm:"type:decimal(20,2)" json:"final_amount"`
	Status        string          `gorm:"type:varchar(50);notnull" json:"status"`
	PaidAt        time.Time       `gorm:"notnull" json:"paid_at"`
}

type Refund struct {
	ID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`

	BookingID uuid.UUID `gorm:"notnull;index:idx_booking_id" json:"booking_id"`
	Booking   Booking   `gorm:"foreignKey:BookingID;references:ID;constraint:OnDelete:RESTRICT;" json:"booking"`

	Amount     decimal.Decimal `gorm:"type:decimal(20,2)" json:"amount"`
	Status     string          `gorm:"type:varchar(50);notnull" json:"status"`
	RefundedAt time.Time       `gorm:"notnull" json:"refunded_at"`
}
