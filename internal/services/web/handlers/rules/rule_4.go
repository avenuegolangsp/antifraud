package rules

import (
	"sync"
	"time"
)

type TransactionRecord struct {
	UserID    string
	Amount    float64
	Timestamp time.Time
}

type VelocityRules struct {
	userTransactions map[string][]TransactionRecord
	mutex            sync.RWMutex
}

func (vr *VelocityRules) AddTransaction(userID string, amount float64, timestamp time.Time) {
	vr.mutex.Lock()
	defer vr.mutex.Unlock()
	
	transaction := TransactionRecord{
		UserID:    userID,
		Amount:    amount,
		Timestamp: timestamp,
	}
	
	vr.userTransactions[userID] = append(vr.userTransactions[userID], transaction)
}

func (vr *VelocityRules) GetRecentTransactions(userID string, window time.Duration) []TransactionRecord {
	vr.mutex.RLock()
	defer vr.mutex.RUnlock()
	
	transactions := vr.userTransactions[userID]
	if transactions == nil {
		return []TransactionRecord{}
	}
	
	cutoffTime := time.Now().Add(-window)
	var recentTransactions []TransactionRecord
	
	for _, transaction := range transactions {
		if transaction.Timestamp.After(cutoffTime) {
			recentTransactions = append(recentTransactions, transaction)
		}
	}
	
	return recentTransactions
}

func (vr *VelocityRules) CleanupOldTransactions(maxAge time.Duration) {
	vr.mutex.Lock()
	defer vr.mutex.Unlock()
	
	cutoffTime := time.Now().Add(-maxAge)
	
	for userID, transactions := range vr.userTransactions {
		var recentTransactions []TransactionRecord
		for _, transaction := range transactions {
			if transaction.Timestamp.After(cutoffTime) {
				recentTransactions = append(recentTransactions, transaction)
			}
		}
		
		if len(recentTransactions) == 0 {
			delete(vr.userTransactions, userID)
		} else {
			vr.userTransactions[userID] = recentTransactions
		}
	}
}

type Rule4_TransactionVelocity struct {
	Threshold int
	Window time.Duration
	store *VelocityRules
}

func NewRule4_TransactionVelocity(threshold int, window time.Duration) *Rule4_TransactionVelocity {
	return &Rule4_TransactionVelocity{
		Threshold: threshold,
		Window:    window,
		store: &VelocityRules{
			userTransactions: make(map[string][]TransactionRecord),
		},
	}
}

func (r *Rule4_TransactionVelocity) Apply(input RuleInput) RuleResult {
	r.store.AddTransaction(
		input.User.ID, 
		input.Transaction.Amount, 
		input.Transaction.TransactionDate,
	)
	
	recentTransactions := r.store.GetRecentTransactions(input.User.ID, r.Window)
	
	transactionCount := len(recentTransactions)
	if transactionCount > r.Threshold {
		score := 80
		if transactionCount > r.Threshold*2 {
			score = 85
		}
		
		return RuleResult{Score: score}
	}
	
	if transactionCount%10 == 0 { 
		go r.store.CleanupOldTransactions(time.Hour)
	}
	
	return RuleResult{Score: 0}
}
