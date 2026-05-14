package entity

import (
	"time"

	"github.com/google/uuid"
)

type Region struct {
	ID        uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Name      string     `gorm:"size:255;notnull" json:"name"`
	CreatedAt time.Time  `gorm:"type:timestamptz;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt *time.Time `gorm:"type:timestamptz;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type City struct {
	ID        uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	RegionID  uuid.UUID  `gorm:"notnull;index:idx_region_id" json:"region_id"`
	Region    Region     `gorm:"foreignKey:RegionID;references:ID;constraint:OnDelete:RESTRICT;" json:"region"`
	Name      string     `gorm:"size:255;notnull" json:"name"`
	CreatedAt time.Time  `gorm:"type:timestamptz;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt *time.Time `gorm:"type:timestamptz;default:CURRENT_TIMESTAMP" json:"updated_at"`
}
