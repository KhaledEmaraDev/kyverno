package api

// PolicyResponse policy application response
type PolicyResponse struct {
	// stats contains policy statistics
	stats PolicyStats
	// Rules contains policy rules responses
	Rules []RuleResponse
}

func (pr *PolicyResponse) Add(stats ExecutionStats, responses ...RuleResponse) {
	for _, response := range responses {
		pr.Rules = append(pr.Rules, response.WithStats(stats))
		status := response.Status()
		if status == RuleStatusPass || status == RuleStatusFail {
			pr.stats.rulesAppliedCount++
		} else if status == RuleStatusError {
			pr.stats.rulesErrorCount++
		}
	}
}

func NewPolicyResponse() PolicyResponse {
	return PolicyResponse{}
}

func (pr *PolicyResponse) Patches() [][]byte {
	var patches [][]byte
	for _, rule := range pr.Rules {
		if rule.HasMutatePatch() {
			patches = append(patches, rule.MutatePatches()...)
		}
	}
	return patches
}

func (pr *PolicyResponse) Stats() PolicyStats {
	return pr.stats
}

func (pr *PolicyResponse) RulesAppliedCount() int {
	return pr.stats.RulesAppliedCount()
}

func (pr *PolicyResponse) RulesErrorCount() int {
	return pr.stats.RulesErrorCount()
}

func (pr *PolicyResponse) HasPatches() bool {
	for _, rule := range pr.Rules {
		if rule.HasMutatePatch() {
			return true
		}
	}
	return false
}
