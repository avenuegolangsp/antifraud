package rules

import "github.com/avenuegolangsp/antifraud/internal/services/web/repository"

// TYPES!!@!!
type AnalyzeResponse struct {
	TransactionID string   `json:"transaction_id"`
	RiskScore     int      `json:"risk_score"`
	RiskLevel     string   `json:"risk_level"`
	Approved      bool     `json:"approved"`
	Alerts        []string `json:"alerts"`
}

type AnalyzeRequest struct {
	UserID    string          `json:"user_id"`
	Amount    float64         `json:"amount"`
	Type      string          `json:"type"`
	Direction string          `json:"direction"`
	Location  AnalyzeLocation `json:"location"`
	Timestamp string          `json:"timestamp"`
}

type AnalyzeLocation struct {
	Country   string  `json:"country"`
	City      string  `json:"city"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type RuleInput struct {
	User        repository.User
	Transaction repository.Transaction
	Location    AnalyzeLocation
}

type RuleResult struct {
	Score int
}

type RuleEntity interface {
	Apply(input RuleInput) RuleResult
}

// implementation above

//go:generate mockery --name IRuleManager --output ./mocks
type IRuleManager interface {
	AnalyzeTransaction(req AnalyzeRequest) (AnalyzeResponse, error)
}

type ruleManager struct {
	rules []RuleEntity
}

func NewRuleManager() IRuleManager {
	return &ruleManager{
		rules: []RuleEntity{
			&Rule1_ImpossibleTravel{},
		},
	}
}

func (rm *ruleManager) AnalyzeTransaction(req AnalyzeRequest) (AnalyzeResponse, error) {
	return AnalyzeResponse{}, nil
}
