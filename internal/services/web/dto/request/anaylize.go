package request

import (
	"time"
)

type Analyze struct {
	UserID    string          `json:"user_id" validate:"required"`
	Amount    float64         `json:"amount" validate:"required,gt=0"`
	Type      string          `json:"type" validate:"required,oneof=pix credito debito"`
	Direction string          `json:"direction" validate:"required,oneof=credito debito"`
	Location  AnalyzeLocation `json:"location" validate:"required,dive"`
	Timestamp time.Time       `json:"timestamp" validate:"required"`
}

type AnalyzeLocation struct {
	Country   string  `json:"country" validate:"required,len=2"`
	City      string  `json:"city" validate:"required"`
	Latitude  float64 `json:"latitude" validate:"required"`
	Longitude float64 `json:"longitude" validate:"required"`
}
