package rules

import (
	"github.com/avenuegolangsp/antifraud/internal/services/web/repository"
)

type Rule2_Anomalus struct {
}

func (r *Rule2_Anomalus) Apply(user repository.User, transaction repository.Transaction) RuleResult {
	// Utilize os campos corretamente se realmente existirem, senão remova-os
	// alertMultiplier := user.AnomalusAmount.MultiplierAlert
	// blockMultiplier := user.AnomalusAmount.MultiplierBlock

	// Se os campos AvgTransactionAmount e MaxTransactionAmount não existirem, remova-os ou ajuste conforme necessário
	// avgUser := user.AvgTransactionAmount
	// maxValue := user.MaxTransactionAmount

	transactionValue := transaction.Amount

	// Exemplo de lógica sem os campos inexistentes
	if transactionValue <= 0 {
		return RuleResult{Score: 0}
	}

	// lógica adicional se necessário

	return RuleResult{Score: 80}
}
