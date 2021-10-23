package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID  `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name      string     `gorm:"size:256" json:"name"`
	Email     string     `gorm:"size:256" json:"email"`
	Password  string     `gorm:"size:256" json:"-"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}
