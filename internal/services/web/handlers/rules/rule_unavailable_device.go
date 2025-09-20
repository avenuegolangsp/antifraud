package rules

import "github.com/avenuegolangsp/antifraud/internal/services/web/repository"

type Rule1_UnavalaibleDevice struct {
}

const SCORE_MAX = 70

func (r *Rule1_UnavalaibleDevice) Apply(input RuleInput) RuleResult {
	users := repository.GetUserList()
	var currentUser repository.User

	for _,user := range users {
		if user.ID == input.User.ID {
			currentUser = user
			break
		}
	}

	isTrustedDevice := false

	for _, device := range currentUser.TrustedDevices {
		if device.DeviceID == input.DeviceID {
			isTrustedDevice = true
		}
	}

	score := 0

	if (!isTrustedDevice) {
		score = 1 * SCORE_MAX
	}

	return RuleResult{Score: score}
}
