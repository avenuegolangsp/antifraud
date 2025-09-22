package repository

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

type Location struct {
	City      string  `json:"city"`
	Country   string  `json:"country"`
	Frequency float64 `json:"frequency"`
}

type FraudHistory struct {
	TotalTransactions      int     `json:"total_transactions"`
	SuspiciousTransactions int     `json:"suspicious_transactions"`
	BlockedTransactions    int     `json:"blocked_transactions"`
	AvgTransactionAmount   float64 `json:"avg_transaction_amount"`
	MaxTransactionAmount   float64 `json:"max_transaction_amount"`
	LastTransactionDate    string  `json:"last_transaction_date"`
}

type BehaviorPatterns struct {
	TypicalTransactionHours []int      `json:"typical_transaction_hours"`
	TypicalLocations        []Location `json:"typical_locations"`
	TypicalAmounts          []int      `json:"typical_amounts"`
	PreferredMerchants      []string   `json:"preferred_merchants"`
}

type TrustedDevice struct {
	DeviceID   string  `json:"device_id"`
	Platform   string  `json:"platform"`
	TrustScore float64 `json:"trust_score"`
	FirstSeen  string  `json:"first_seen"`
	LastSeen   string  `json:"last_seen"`
}

type User struct {
	ID               string           `json:"id"`
	Name             string           `json:"name"`
	Profile          string           `json:"profile"`
	Email            string           `json:"email"`
	Phone            string           `json:"phone"`
	CreatedAt        string           `json:"created_at"`
	Status           string           `json:"status"`
	FraudHistory     FraudHistory     `json:"fraud_history"`
	BehaviorPatterns BehaviorPatterns `json:"behavior_patterns"`
	TrustedDevices   []TrustedDevice  `json:"trusted_devices"`
	AnomalusAmount   AnomalusAmount   `json:"anomalous_amount"`
}

type AnomalusAmount struct{
	MultiplierAlert float32 `json:"multiplier_alert"`
	MultiplierBlock float32 `json:"multiplier_block"`
}

func GetUserList() []User {

	var Users []User

	files := []string{
		"antifraud/data/users01.json",
		"antifraud/data/user02.json",
		"antifraud/data/users03.json",
	}

	for _, file := range files {
		path := filepath.Join(file)
		data, err := os.ReadFile(path)
		if err != nil {
			log.Printf("erro ao ler arquivo %s: %v", file, err)
			continue
		}

		// estrutura auxiliar para casar com o JSON
		var wrapper struct {
			Users []User `json:"users"`
		}

		if err := json.Unmarshal(data, &wrapper); err != nil {
			log.Printf("erro ao parsear %s: %v", file, err)
			continue
		}

		Users = append(Users, wrapper.Users...)
	}

	return Users
}
