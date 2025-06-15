package dto

type MeasurementResponse struct {
	ID        uint    `json:"id"`
	Name      string  `json:"name"`
	Phone     string  `json:"phone"`
	Type      string  `json:"type"`
	Chest     float32 `json:"chest"`
	Waist     float32 `json:"waist"`
	Hips      float32 `json:"hips"`
	Shoulder  float32 `json:"shoulder"`
	Length    float32 `json:"length"`
	Sleeve    float32 `json:"sleeve"`
	Note      string  `json:"note"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}
