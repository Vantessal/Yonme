package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Account struct {
	ID 		uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Email    string `gorm:"uniqueIndex;not null" json:"email" binding:"required,email"`
	Username string `json:"username" binding:"required"`
	Password string `json:"-"`
	CreatedAt time.Time `gorm:"<-:create"`
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:" index;<-:create"`
}