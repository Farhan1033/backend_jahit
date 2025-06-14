package models

import "time"

type Measurement struct {
	IdCustomer   uint      `gorm:"primaryKey;autoIncrement" json:"id_customer"`
	CustomerName string    `json:"customer_name"`
	Gender       string    `json:"gender"`
	Chest        float64   `json:"chest"`
	Waist        float64   `json:"waist"`
	Hip          float64   `json:"hip"`
	ArmLength    float64   `json:"arm_length"`
	ShirtLength  float64   `json:"shirt_length"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
