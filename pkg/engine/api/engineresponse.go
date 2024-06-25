package api

import (
	"fmt"

	"github.com/go-logr/logr"
	kyvernov1 "github.com/kyverno/kyverno/api/kyverno/v1"
	"github.com/kyverno/kyverno/ext/wildcard"
	"github.com/kyverno/kyverno/pkg/engine/mutate/patch"
	datautils "github.com/kyverno/kyverno/pkg/utils/data"
	jsonutils "github.com/kyverno/kyverno/pkg/utils/json"
	kubeutils "github.com/kyverno/kyverno/pkg/utils/kube"
	utils "github.com/kyverno/kyverno/pkg/utils/match"
	"gomodules.xyz/jsonpatch/v2"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

// ResourceChange represents a change to be applied to a resource
type ResourceChange struct {
	// Patches is the JSON patch representing the change
	Patches [][]byte
	// PatchedResource is the resource after applying the change
	PatchedResource unstructured.Unstructured
}

func (rc ResourceChange) IsPatch() bool {
	return len(rc.Patches) != 0
}

func (rc ResourceChange) GetPatches(base *unstructured.Unstructured) []jsonpatch.Operation {
	if rc.IsPatch() {
		return []jsonpatch.Operation{}
	}

	originalBytes, err := base.MarshalJSON()
	if err != nil {
		return nil
	}
	patchedBytes, err := rc.PatchedResource.MarshalJSON()
	if err != nil {
		return nil
	}

	patches, err := jsonpatch.CreatePatch(originalBytes, patchedBytes)
	if err != nil {
		return nil
	}

	return patches
}

func (rc ResourceChange) GetPatchBytes(base *unstructured.Unstructured) [][]byte {
	if rc.IsPatch() {
		return rc.Patches
	}

	patches := rc.GetPatches(base)

	return patch.ConvertPatches(patches...)
}

func (rc ResourceChange) GetPatchedResourceBytes(logger logr.Logger, resoruceRaw []byte) ([]byte, error) {
	if !rc.IsPatch() {
		return rc.PatchedResource.MarshalJSON()
	}

	combinedPatch := jsonutils.JoinPatches(rc.Patches...)
	patchedResource, err := patch.ProcessPatchJSON6902(logger, resoruceRaw, combinedPatch)
	if err != nil {
		return nil, err
	}

	return patchedResource, nil
}

func (rc ResourceChange) GetPatchedResource(logger logr.Logger, resoruceRaw []byte) (*unstructured.Unstructured, error) {
	if !rc.IsPatch() {
		return &rc.PatchedResource, nil
	}

	combinedPatch := jsonutils.JoinPatches(rc.Patches...)
	patched, err := patch.ProcessPatchJSON6902(logger, resoruceRaw, combinedPatch)
	if err != nil {
		return nil, err
	}

	return kubeutils.BytesToUnstructured(patched)
}

func (rc ResourceChange) GetResourceSpec(resoruce unstructured.Unstructured) ResourceSpec {
	if !rc.IsPatch() {
		return ResourceSpec{
			Kind:       rc.PatchedResource.GetKind(),
			APIVersion: rc.PatchedResource.GetAPIVersion(),
			Namespace:  rc.PatchedResource.GetNamespace(),
			Name:       rc.PatchedResource.GetName(),
			UID:        string(rc.PatchedResource.GetUID()),
		}
	}

	// TODO: handle patch case
	return ResourceSpec{
		Kind:       resoruce.GetKind(),
		APIVersion: resoruce.GetAPIVersion(),
		Namespace:  resoruce.GetNamespace(),
		Name:       resoruce.GetName(),
		UID:        string(resoruce.GetUID()),
	}
}

// EngineResponse engine response to the action
type EngineResponse struct {
	// Resource is the original resource
	Resource unstructured.Unstructured
	// Policy is the original policy
	policy GenericPolicy
	// namespaceLabels given by policy context
	namespaceLabels map[string]string
	// Change represents the change to be applied to the resource
	Change ResourceChange
	// PolicyResponse contains the engine policy response
	PolicyResponse PolicyResponse
	// stats contains engine statistics
	stats ExecutionStats
}

func resource(policyContext PolicyContext) unstructured.Unstructured {
	resource := policyContext.NewResource()
	if resource.Object == nil {
		resource = policyContext.OldResource()
	}
	return resource
}

func NewEngineResponseFromPolicyContext(policyContext PolicyContext) EngineResponse {
	return NewEngineResponse(
		resource(policyContext),
		NewKyvernoPolicy(policyContext.Policy()),
		policyContext.NamespaceLabels(),
	)
}

func NewEngineResponse(
	resource unstructured.Unstructured,
	policy GenericPolicy,
	namespaceLabels map[string]string,
) EngineResponse {
	return EngineResponse{
		Resource:        resource,
		policy:          policy,
		namespaceLabels: namespaceLabels,
		Change: ResourceChange{
			PatchedResource: resource,
			Patches:         [][]byte{},
		},
	}
}

func (er EngineResponse) WithPolicy(policy GenericPolicy) EngineResponse {
	er.policy = policy
	return er
}

func (er EngineResponse) WithPolicyResponse(policyResponse PolicyResponse) EngineResponse {
	er.PolicyResponse = policyResponse
	return er
}

func (r EngineResponse) WithStats(stats ExecutionStats) EngineResponse {
	r.stats = stats
	return r
}

func (er EngineResponse) WithPatchedResource(patchedResource unstructured.Unstructured) EngineResponse {
	er.Change.PatchedResource = patchedResource
	return er
}

func (er EngineResponse) WithPatches(patches [][]byte) EngineResponse {
	er.Change.Patches = patches
	return er
}

func (er EngineResponse) WithNamespaceLabels(namespaceLabels map[string]string) EngineResponse {
	er.namespaceLabels = namespaceLabels
	return er
}

func (er *EngineResponse) NamespaceLabels() map[string]string {
	return er.namespaceLabels
}

func (er *EngineResponse) Policy() GenericPolicy {
	return er.policy
}

// IsOneOf checks if any rule has status in a given list
func (er EngineResponse) IsOneOf(status ...RuleStatus) bool {
	for _, r := range er.PolicyResponse.Rules {
		if r.HasStatus(status...) {
			return true
		}
	}
	return false
}

// IsSuccessful checks if any rule has failed or produced an error during execution
func (er EngineResponse) IsSuccessful() bool {
	return !er.IsOneOf(RuleStatusFail, RuleStatusError)
}

// IsSkipped checks if any rule has skipped resource or not.
func (er EngineResponse) IsSkipped() bool {
	return er.IsOneOf(RuleStatusSkip)
}

// IsFailed checks if any rule created a policy violation
func (er EngineResponse) IsFailed() bool {
	return er.IsOneOf(RuleStatusFail)
}

// IsError checks if any rule resulted in a processing error
func (er EngineResponse) IsError() bool {
	return er.IsOneOf(RuleStatusError)
}

// IsEmpty checks if any rule results are present
func (er EngineResponse) IsEmpty() bool {
	return len(er.PolicyResponse.Rules) == 0
}

// isNil checks if rule is an empty rule
func (er EngineResponse) IsNil() bool {
	return datautils.DeepEqual(er, EngineResponse{})
}

func (er EngineResponse) GetPatchBytes() [][]byte {
	return er.Change.GetPatchBytes(&er.Resource)
}

// GetPatches returns all the patches joined
func (er EngineResponse) GetPatches() []jsonpatch.JsonPatchOperation {
	return er.Change.GetPatches(&er.Resource)
}

// GetFailedRules returns failed rules
func (er EngineResponse) GetFailedRules() []string {
	return er.getRules(func(rule RuleResponse) bool { return rule.HasStatus(RuleStatusFail, RuleStatusError) })
}

// GetFailedRulesWithErrors returns failed rules with corresponding error messages
func (er EngineResponse) GetFailedRulesWithErrors() []string {
	return er.getRulesWithErrors(func(rule RuleResponse) bool { return rule.HasStatus(RuleStatusFail, RuleStatusError) })
}

// GetSuccessRules returns success rules
func (er EngineResponse) GetSuccessRules() []string {
	return er.getRules(func(rule RuleResponse) bool { return rule.HasStatus(RuleStatusPass) })
}

// GetResourceSpec returns resourceSpec of er
func (er EngineResponse) GetResourceSpec() ResourceSpec {
	return er.Change.GetResourceSpec(er.Resource)
}

func (er EngineResponse) getRules(predicate func(RuleResponse) bool) []string {
	var rules []string
	for _, r := range er.PolicyResponse.Rules {
		if predicate(r) {
			rules = append(rules, r.Name())
		}
	}
	return rules
}

func (er EngineResponse) getRulesWithErrors(predicate func(RuleResponse) bool) []string {
	var rules []string
	for _, r := range er.PolicyResponse.Rules {
		if predicate(r) {
			rules = append(rules, fmt.Sprintf("%s: %s", r.Name(), r.Message()))
		}
	}
	return rules
}

// If the policy is of type ValidatingAdmissionPolicy, an empty string is returned.
func (er EngineResponse) GetValidationFailureAction() kyvernov1.ValidationFailureAction {
	pol := er.Policy()
	if polType := pol.GetType(); polType == ValidatingAdmissionPolicyType {
		return ""
	}
	spec := pol.AsKyvernoPolicy().GetSpec()
	for _, v := range spec.ValidationFailureActionOverrides {
		if !v.Action.IsValid() {
			continue
		}
		if v.Namespaces == nil {
			hasPass, err := utils.CheckSelector(v.NamespaceSelector, er.namespaceLabels)
			if err == nil && hasPass {
				return v.Action
			}
		}
		for _, ns := range v.Namespaces {
			// TODO: handle new namespace
			if wildcard.Match(ns, er.Resource.GetNamespace()) {
				if v.NamespaceSelector == nil {
					return v.Action
				}
				hasPass, err := utils.CheckSelector(v.NamespaceSelector, er.namespaceLabels)
				if err == nil && hasPass {
					return v.Action
				}
			}
		}
	}
	return spec.ValidationFailureAction
}
