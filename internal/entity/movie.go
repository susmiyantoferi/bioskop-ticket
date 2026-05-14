package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Movie struct {
	ID              uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Name            string    `gorm:"size:255;notnull" json:"name"`
	DurationMinutes int       `gorm:"notnull" json:"duration_minutes"`
	Description     string    `gorm:"notnull" json:"description"`

	Schedules []Schedule `gorm:"foreignKey:MovieID" json:"schedules"`

	CreatedAt time.Time  `gorm:"type:timestamptz;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt *time.Time `gorm:"type:timestamptz;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type Schedule struct {
	ID        uuid.UUID       `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	MovieID   uuid.UUID       `gorm:"notnull;index:idx_movie_id" json:"movie_id"`
	Movie     Movie           `gorm:"foreignKey:MovieID;references:ID;constraint:OnDelete:RESTRICT;" json:"movie"`
	StudioID  uuid.UUID       `gorm:"notnull;index:idx_studio_id" json:"studio_id"`
	Studio    Studio          `gorm:"foreignKey:StudioID;references:ID;constraint:OnDelete:RESTRICT;" json:"studio"`
	ShowTime  time.Time       `gorm:"notnull" json:"show_time"`
	EndTime   time.Time       `gorm:"notnull" json:"end_time"`
	Price     decimal.Decimal `gorm:"type:decimal(12,2);notnull" json:"decimal"`
	CreatedAt time.Time       `gorm:"type:timestamptz;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt *time.Time      `gorm:"type:timestamptz;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt *gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
