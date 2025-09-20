package rules

type Rule1_ImpossibleTravel struct {
}

func (r *Rule1_ImpossibleTravel) Apply(input RuleInput) RuleResult {
	return RuleResult{Score: 0}
}
