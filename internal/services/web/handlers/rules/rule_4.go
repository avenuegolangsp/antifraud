package rules

import (
	"time"
)

// Rule4_TransactionVelocity verifica se há um número excessivo de transações
// em um curto período de tempo, o que pode indicar fraude.
type Rule4_TransactionVelocity struct {
	// Threshold é o limite de transações permitidas por um período
	Threshold int
	// Window é o período de tempo a ser verificado (ex: 5 minutos)
	Window time.Duration
}

// Apply aplica a regra de velocidade de transação.
// Note: Esta é uma implementação simplificada. Na vida real, você precisaria
// de um banco de dados ou um cache (como Redis) para rastrear o histórico
// de transações recentes do usuário.
func (r *Rule4_TransactionVelocity) Apply(input RuleInput) RuleResult {
	// Lógica para buscar transações recentes do usuário e contá-las.
	// Exemplo:
	// transactions := getRecentTransactions(input.CustomerID, r.Window)
	// if len(transactions) > r.Threshold {
	//     return RuleResult{Score: 80, Details: "Velocidade de transação excedida."}
	// }

	// Implementação de exemplo (simples)
	// Se a regra detecta um possível problema, atribui um score alto.
	// Por enquanto, retorna 0.
	if input.TransactionVelocity > r.Threshold {
		return RuleResult{Score: 80, Details: "Velocidade de transação excedida."}
	}

	return RuleResult{Score: 0}
}
