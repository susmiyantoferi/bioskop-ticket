package entity

import (
	"time"

	"github.com/google/uuid"
)

type Seat struct {
	ID         uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	StudioID   uuid.UUID  `gorm:"notnull;index:idx_studio_id" json:"studio_id"`
	Studio     Studio     `gorm:"foreignKey:StudioID;references:ID;constraint:OnDelete:RESTRICT;" json:"studio"`
	SeatRow    string     `gorm:"notnull" json:"seat_row"`
	SeatNumber int        `gorm:"notnull" json:"seat_number"`
	CreatedAt  time.Time  `gorm:"type:timestamptz;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  *time.Time `gorm:"type:timestamptz;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type SeatLock struct {
	ID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`

	UserID uuid.UUID `gorm:"notnull;index"`
	User   User      `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`

	ScheduleID uuid.UUID `gorm:"notnull;uniqueIndex"`
	Schedule   Schedule  `gorm:"foreignKey:ScheduleID;references:ID;constraint:OnDelete:CASCADE"`

	SeatID uuid.UUID `gorm:"notnull;uniqueIndex"`
	Seat   Seat      `gorm:"foreignKey:SeatID;references:ID;constraint:OnDelete:CASCADE"`

	Status string `gorm:"type:varchar(30);notnull"`

	ExpiredAt time.Time `gorm:"notnull"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
