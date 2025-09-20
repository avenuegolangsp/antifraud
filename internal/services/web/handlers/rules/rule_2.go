package rules

import (
	"github.com/avenuegolangsp/antifraud/internal/services/web/repository"
)

type Rule2_Anomalus struct {
}

func (r *Rule2_Anomalus) Apply(user repository.User, transaction  repository.Transaction) RuleResult {
	alertMultiplier := user.AnomalusAmount.MultiplierAlert
	blockMultiplier := user.AnomalusAmount.MultiplierBlock

	avgUser := user.AvgTransactionAmount
	maxValue := user.MaxTransactionAmount

	FraudHistory := user.AnomalusAmount.FraudHistory


	transactionValue := transaction.Amount

	if transactionValue <= avgUser {
		return RuleResult{Score: 0}
	}

	if transactionValue < avgUser && transactionValue >=maxValue{

	}

	return RuleResult{Score: 80}
}