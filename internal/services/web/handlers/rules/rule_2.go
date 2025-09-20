package rules

import (
	"github.com/avenuegolangsp/antifraud/internal/services/web/repository"
)

type Rule2_Anomalus struct {
}

func (r *Rule2_Anomalus) Apply(user repository.User, transaction  repository.Transaction) RuleResult {
	alertMultiplier := input.AnomalusAmount.MultiplierAlert
	blockMultiplier := input.AnomalusAmount.MultiplierBlock
	FraudHistory := input.AnomalusAmount.FraudHistory


	
	
	return RuleResult{Score: 0}
}