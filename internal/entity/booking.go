package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Booking struct {
	ID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`

	UserID uuid.UUID `gorm:"notnull;index:idx_user_id" json:"user_id"`
	User   User      `gorm:"foreignKey:UserId;references:ID;constraint:OnDelete:RESTRICT;" json:"user"`

	ScheduleID uuid.UUID `gorm:"notnull;index:idx_schedule_id" json:"schedule_id"`
	Schedule   Schedule  `gorm:"foreignKey:ScheduleID;references:ID;constraint:OnDelete:RESTRICT;" json:"schedule"`

	BookingCode string          `gorm:"notnull;uniqueIndex" json:"booking_code"`
	TotalAmount decimal.Decimal `gorm:"type:decimal(12,2);notnull"`
	Status      string          `gorm:"type:varchar(50);notnull" json:"status"`

	BookingSeats []BookingSeat `gorm:"foreignKey:BookingID"`

	Payment *Payment `gorm:"foreignKey:BookingID"`

	Refund *Refund `gorm:"foreignKey:BookingID"`

	CreatedAt time.Time  `gorm:"type:timestamptz;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt *time.Time `gorm:"type:timestamptz;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type BookingSeat struct {
	ID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`

	BookingID uuid.UUID `gorm:"notnull;index:idx_booking_id" json:"booking_id"`
	Booking   Booking   `gorm:"foreignKey:BookingID;references:ID;constraint:OnDelete:RESTRICT;" json:"booking"`

	ScheduleID uuid.UUID `gorm:"notnull;index:uniqueIndex" json:"schedule_id"`
	Schedule   Schedule  `gorm:"foreignKey:ScheduleID;references:ID;constraint:OnDelete:RESTRICT;" json:"schedule"`

	SeatID uuid.UUID `gorm:"notnull;index:uniqueIndex" json:"seat_id"`
	Seat   Seat      `gorm:"foreignKey:SeatID;references:ID;constraint:OnDelete:RESTRICT;" json:"seat"`

	CreatedAt time.Time  `gorm:"type:timestamptz;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt *time.Time `gorm:"type:timestamptz;default:CURRENT_TIMESTAMP" json:"updated_at"`
}
