package rules

//only to stop error
type Rule1_ImpossibleTravel2 struct {
}

func (r *Rule1_ImpossibleTravel) Apply2(input RuleInput) RuleResult {
	return RuleResult{Score: 0}
}
