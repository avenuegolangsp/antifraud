package rules

import (
	"errors"
	"time"
)

// RuleSuspiciousHours implementa a detecção de transações em horários suspeitos
// baseada nos padrões típicos de horário do usuário
type RuleSuspiciousHours struct{}

// Apply analisa se o horário da transação está fora da janela segura do usuário
func (r *RuleSuspiciousHours) Apply(input RuleInput) (*RuleResult, error) {
	// Score máximo para horário suspeito (conforme especificação)
	const maxScore = 60

	// Verifica se temos dados do usuário
	user, ok := input.User.(map[string]interface{})
	if !ok {
		return nil, errors.New("user not found")
	}

	// Extrai os horários típicos do usuário
	behaviorPatterns, ok := user["behavior_patterns"].(map[string]interface{})
	if !ok {
		return nil, errors.New("user behavior patterns not found")
	}

	typicalHours, ok := behaviorPatterns["typical_transaction_hours"].([]interface{})
	if !ok || len(typicalHours) == 0 {
		return nil, errors.New("user typical transaction hours not found")
	}

	// Converte os horários típicos para slice de int
	var typicalHoursInt []int
	for _, hour := range typicalHours {
		if h, ok := hour.(float64); ok {
			typicalHoursInt = append(typicalHoursInt, int(h))
		}
	}

	// Parse do timestamp da transação
	transactionTime, err := time.Parse(time.RFC3339, input.Transaction.(map[string]interface{})["timestamp"].(string))
	if err != nil {
		return nil, err
	}

	// Calcula a janela de horários seguros
	safeWindow := r.calculateSafeWindow(typicalHoursInt)

	// Verifica se o horário está fora da janela segura
	if r.isOutsideSafeWindow(transactionTime, safeWindow) {
		// Calcula a distância temporal e aplica o multiplicador
		distance := r.calculateTimeDistance(transactionTime, safeWindow)
		multiplier := r.getRiskMultiplier(distance)

		score := int(float64(maxScore) * multiplier)
		if score > maxScore {
			score = maxScore
		}

		return &RuleResult{Score: score}, nil
	}

	return &RuleResult{Score: 0}, nil
}

// calculateSafeWindow calcula a janela de horários seguros baseada nos horários típicos
func (r *RuleSuspiciousHours) calculateSafeWindow(typicalHours []int) []int {
	if len(typicalHours) == 0 {
		return []int{}
	}

	// Encontra o menor e maior horário
	minHour := typicalHours[0]
	maxHour := typicalHours[0]

	for _, hour := range typicalHours {
		if hour < minHour {
			minHour = hour
		}
		if hour > maxHour {
			maxHour = hour
		}
	}

	// Cria a janela segura com uma margem de 1 hora antes e depois
	// para considerar horários próximos como seguros
	safeStart := minHour - 1
	safeEnd := maxHour + 1

	// Ajusta para o formato 24h (0-23)
	if safeStart < 0 {
		safeStart = 0
	}
	if safeEnd > 23 {
		safeEnd = 23
	}

	// Cria a janela de horários seguros
	var safeWindow []int
	for hour := safeStart; hour <= safeEnd; hour++ {
		safeWindow = append(safeWindow, hour)
	}

	return safeWindow
}

// isOutsideSafeWindow verifica se o horário da transação está fora da janela segura
func (r *RuleSuspiciousHours) isOutsideSafeWindow(transactionTime time.Time, safeWindow []int) bool {
	transactionHour := transactionTime.Hour()

	for _, safeHour := range safeWindow {
		if transactionHour == safeHour {
			return false
		}
	}

	return true
}

// calculateTimeDistance calcula a distância temporal em horas até a janela segura mais próxima
func (r *RuleSuspiciousHours) calculateTimeDistance(transactionTime time.Time, safeWindow []int) int {
	transactionHour := transactionTime.Hour()

	if len(safeWindow) == 0 {
		return 24 // Distância máxima se não houver janela segura
	}

	minDistance := 24 // Inicializa com distância máxima possível

	for _, safeHour := range safeWindow {
		distance := r.abs(transactionHour - safeHour)
		if distance < minDistance {
			minDistance = distance
		}
	}

	return minDistance
}

// getRiskMultiplier retorna o multiplicador de risco baseado na distância temporal
func (r *RuleSuspiciousHours) getRiskMultiplier(distanceHours int) float64 {
	switch {
	case distanceHours <= 0:
		return 0.0 // Dentro da janela segura
	case distanceHours <= 1:
		return 0.4 // 30 minutos a 1 hora
	case distanceHours <= 2:
		return 0.5 // 1 a 2 horas
	case distanceHours <= 3:
		return 0.7 // 2 a 3 horas
	default:
		return 1.0 // 3 ou mais horas
	}
}

// abs retorna o valor absoluto de um inteiro
func (r *RuleSuspiciousHours) abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
