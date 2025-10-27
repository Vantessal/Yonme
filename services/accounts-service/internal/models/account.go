package models

import (
	"github.com/google/uuid"
	"time"

	
)

type Account struct {
	ID       uuid.UUID `gorm:"primaryKey" json:"id"`
	Email    string `gorm:"uniqueIndex;not null" json:"email"`
	Username string `json:"username"`
	Password string `json:"-"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}