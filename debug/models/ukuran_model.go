package models

import (
	"time"

	"github.com/google/uuid"
)

type Measurement struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	CustomerID uuid.UUID `json:"customer_id"`
	Customer   Customer  `gorm:"foreignKey:CustomerID" json:"-"`

	Type     string  `json:"type"`
	Chest    float32 `json:"chest"`
	Waist    float32 `json:"waist"`
	Hips     float32 `json:"hips"`
	Shoulder float32 `json:"shoulder"`
	Length   float32 `json:"length"`
	Sleeve   float32 `json:"sleeve"`
	Note     string  `json:"note"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
