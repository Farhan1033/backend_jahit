package models

import (
	"time"

	"github.com/google/uuid"
)

type Customer struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(100);not null" json:"name"`
	Phone     string    `gorm:"type:varchar(15);not null" json:"phone"`
	CreatedAt time.Time `json:"created_at"`

	Measurement []Measurement `json:"measurement"`
}
