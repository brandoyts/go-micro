package entities

import (
	passwordHashing "microservices/pkg/password-hashing"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Uuid       string `gorm:"primaryKey"`
	Name       string `gorm:"not null"`
	Username   string `gorm:"unique;not null"`
	Password   string `gorm:"not null"`
	IsVerified bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

// BeforeCreate is automatically called by GORM
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	hashedPassword, hashErr := passwordHashing.HashPassword(u.Password)
	if hashErr != nil {
		return hashErr
	}

	u.Uuid = uuid.NewString()
	u.Password = hashedPassword

	return
}
