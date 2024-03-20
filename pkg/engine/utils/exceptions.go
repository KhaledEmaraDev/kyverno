package utils

import (
	"github.com/go-logr/logr"
	kyvernov2beta1 "github.com/kyverno/kyverno/api/kyverno/v2beta1"
	engineapi "github.com/kyverno/kyverno/pkg/engine/api"
	"github.com/kyverno/kyverno/pkg/utils/conditions"
	matched "github.com/kyverno/kyverno/pkg/utils/match"
)

// MatchesException takes a list of exceptions and checks if there is an exception applies to the incoming resource.
// It returns the matched policy exception.
func MatchesException(
	polexs []*kyvernov2beta1.PolicyException,
	policyContext engineapi.PolicyContext,
	logger logr.Logger,
) *kyvernov2beta1.PolicyException {
	gvk, subresource := policyContext.ResourceKind()
	resource := policyContext.NewResource()
	if resource.Object == nil {
		resource = policyContext.OldResource()
	}

	for i := range polexs {
		err := matched.CheckMatchesResources(
			resource,
			polexs[i].Spec.Match,
			policyContext.NamespaceLabels(),
			policyContext.AdmissionInfo(),
			gvk,
			subresource,
		)
		// if there is an error it means a no-match
		if err != nil {
			continue
		}

		if polexs[i].Spec.Conditions != nil {
			passed, err := conditions.CheckAnyAllConditions(logger, policyContext.JSONContext(), *polexs[i].Spec.Conditions)
			if err != nil {
				logger.Error(err, "failed to evaluate conditions")
				continue
			}
			if !passed {
				continue
			}
		}

		return polexs[i]
	}

	return nil
}
