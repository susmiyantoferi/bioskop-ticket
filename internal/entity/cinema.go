package entity

import (
	"time"

	"github.com/google/uuid"
)

type Cinema struct {
	ID        uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	CityID    uuid.UUID  `gorm:"notnull;index:idx_cty_id" json:"city_id"`
	City      City       `gorm:"foreignKey:CityID;references:ID;constraint:OnDelete:RESTRICT;" json:"city"`
	Name      string     `gorm:"size:255;notnull" json:"name"`

	Studios []Studio `gorm:"foreignKey:CinemaID" json:"studios"`

	CreatedAt time.Time  `gorm:"type:timestamptz;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt *time.Time `gorm:"type:timestamptz;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type Studio struct {
	ID        uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	CinemaID  uuid.UUID  `gorm:"notnull;index:idx_cinema_id" json:"cinema_id"`
	Cinema    Cinema     `gorm:"foreignKey:CinemaID;references:ID;constraint:OnDelete:RESTRICT;" json:"cinema"`
	Name      string     `gorm:"size:255;notnull" json:"name"`
	Capacity  int        `gorm:"notnull" json:"capacity"`

	Seats []Seat `gorm:"foreignKey:StudioID" json:"seats"`

	CreatedAt time.Time  `gorm:"type:timestamptz;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt *time.Time `gorm:"type:timestamptz;default:CURRENT_TIMESTAMP" json:"updated_at"`
}
