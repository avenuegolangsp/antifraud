package repository

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
}
