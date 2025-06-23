package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Booking struct {
	gorm.Model
	Uuid      string `gorm:"primaryKey"`
	UserUuid  string `gorm:"not null"`
	Schedule  time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

// BeforeCreate is automatically called by GORM
func (b *Booking) BeforeCreate(tx *gorm.DB) (err error) {
	b.Uuid = uuid.NewString()
	return
}
