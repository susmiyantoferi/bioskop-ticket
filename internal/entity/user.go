package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID       `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Name      string          `gorm:"size:255;notnull" json:"name"`
	Email     string          `gorm:"uniqueIndex:user_deleted_at,where:deleted_at IS NULL;size:100;notnull" json:"email"`
	Password  string          `gorm:"size:255;notnull" json:"password"`
	Role      UserRole        `gorm:"default:'CUSTOMER';notnull" json:"role"`

	Bookings []Booking `gorm:"foreignKey:UserID" json:"bookings"`

	CreatedAt time.Time       `gorm:"type:timestamptz;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt *time.Time      `gorm:"type:timestamptz;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt *gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type UserRole string

const (
	Customer UserRole = "CUSTOMER"
	Admin    UserRole = "ADMIN"
)
