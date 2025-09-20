package repository

import (
	"time"
)

type Transaction struct {
	TransactionID   string  `gorm:"primaryKey;size:100"`
	UserID          string  `gorm:"size:100;not null"`
	Amount          float64 `gorm:"type:decimal(15,2)"`
	TransactionDate time.Time
	LocationCity    string                `gorm:"size:100"`
	LocationCountry string                `gorm:"size:50"`
	Analysis        []TransactionAnalysis `gorm:"foreignKey:TransactionID"`
}

type TransactionAnalysis struct {
	ID            uint   `gorm:"primaryKey"`
	TransactionID string `gorm:"size:100;not null"`
	RiskScore     int
	RiskLevel     string `gorm:"size:20"`
	Approved      bool
	Alerts        []TransactionAlert `gorm:"foreignKey:AnalysisID"`
}

type TransactionAlert struct {
	ID         uint `gorm:"primaryKey"`
	AnalysisID uint
	Message    string
	Type       string
	Priority   string
}
