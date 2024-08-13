package api

import (
	"fmt"

	kyvernov2 "github.com/kyverno/kyverno/api/kyverno/v2"
	pssutils "github.com/kyverno/kyverno/pkg/pss/utils"
	"k8s.io/api/admissionregistration/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/pod-security-admission/api"
)

// PodSecurityChecks details about pod securty checks
type PodSecurityChecks struct {
	// Level is the pod security level
	Level api.Level
	// Version is the pod security version
	Version string
	// Checks contains check result details
	Checks []pssutils.PSSCheckResult
}

// RuleResponse details for each rule application
type RuleResponse struct {
	// name is the rule name specified in policy
	name string
	// ruleType is the rule type (Mutation,Generation,Validation) for Kyverno Policy
	ruleType RuleType
	// message is the message response from the rule application
	message string
	// status rule status
	status RuleStatus
	// stats contains rule statistics
	stats ExecutionStats
	// generatedResource is the generated by the generate rules of a policy
	generatedResource unstructured.Unstructured
	// patchedTarget is the patched resource for mutate.targets
	patchedTarget *unstructured.Unstructured
	// patchedTargetParentResourceGVR is the GVR of the parent resource of the PatchedTarget. This is only populated when PatchedTarget is a subresource.
	patchedTargetParentResourceGVR metav1.GroupVersionResource
	// patchedTargetSubresourceName is the name of the subresource which is patched, empty if the resource patched is not a subresource.
	patchedTargetSubresourceName string
	// podSecurityChecks contains pod security checks (only if this is a pod security rule)
	podSecurityChecks *PodSecurityChecks
	// exceptions are the exceptions applied (if any)
	exceptions []kyvernov2.PolicyException
	// binding is the validatingadmissionpolicybinding (if any)
	binding *v1alpha1.ValidatingAdmissionPolicyBinding
	// emitWarning enable passing rule message as warning to api server warning header
	emitWarning bool
}

func NewRuleResponse(name string, ruleType RuleType, msg string, status RuleStatus) *RuleResponse {
	emitWarn := false
	if status == RuleStatusError || status == RuleStatusFail || status == RuleStatusWarn {
		emitWarn = true
	}
	return &RuleResponse{
		name:        name,
		ruleType:    ruleType,
		message:     msg,
		status:      status,
		emitWarning: emitWarn,
	}
}

func RuleError(name string, ruleType RuleType, msg string, err error) *RuleResponse {
	if err != nil {
		return NewRuleResponse(name, ruleType, fmt.Sprintf("%s: %s", msg, err.Error()), RuleStatusError)
	}
	return NewRuleResponse(name, ruleType, msg, RuleStatusError)
}

func RuleSkip(name string, ruleType RuleType, msg string) *RuleResponse {
	return NewRuleResponse(name, ruleType, msg, RuleStatusSkip)
}

func RuleWarn(name string, ruleType RuleType, msg string) *RuleResponse {
	return NewRuleResponse(name, ruleType, msg, RuleStatusWarn)
}

func RulePass(name string, ruleType RuleType, msg string) *RuleResponse {
	return NewRuleResponse(name, ruleType, msg, RuleStatusPass)
}

func RuleFail(name string, ruleType RuleType, msg string) *RuleResponse {
	return NewRuleResponse(name, ruleType, msg, RuleStatusFail)
}

func (r RuleResponse) WithExceptions(exceptions []kyvernov2.PolicyException) *RuleResponse {
	r.exceptions = exceptions
	return &r
}

func (r RuleResponse) WithBinding(binding *v1alpha1.ValidatingAdmissionPolicyBinding) *RuleResponse {
	r.binding = binding
	return &r
}

func (r RuleResponse) WithPodSecurityChecks(checks PodSecurityChecks) *RuleResponse {
	r.podSecurityChecks = &checks
	return &r
}

func (r RuleResponse) WithPatchedTarget(patchedTarget *unstructured.Unstructured, gvr metav1.GroupVersionResource, subresource string) *RuleResponse {
	r.patchedTarget = patchedTarget
	r.patchedTargetParentResourceGVR = gvr
	r.patchedTargetSubresourceName = subresource
	return &r
}

func (r RuleResponse) WithGeneratedResource(resource unstructured.Unstructured) *RuleResponse {
	r.generatedResource = resource
	return &r
}

func (r RuleResponse) WithStats(stats ExecutionStats) RuleResponse {
	r.stats = stats
	return r
}

func (r RuleResponse) WithEmitWarning(emitWarning bool) *RuleResponse {
	r.emitWarning = emitWarning
	return &r
}

func (r *RuleResponse) Stats() ExecutionStats {
	return r.stats
}

func (r *RuleResponse) Exceptions() []kyvernov2.PolicyException {
	return r.exceptions
}

func (r *RuleResponse) ValidatingAdmissionPolicyBinding() *v1alpha1.ValidatingAdmissionPolicyBinding {
	return r.binding
}

func (r *RuleResponse) IsException() bool {
	return len(r.exceptions) > 0
}

func (r *RuleResponse) PodSecurityChecks() *PodSecurityChecks {
	return r.podSecurityChecks
}

func (r *RuleResponse) PatchedTarget() (*unstructured.Unstructured, metav1.GroupVersionResource, string) {
	return r.patchedTarget, r.patchedTargetParentResourceGVR, r.patchedTargetSubresourceName
}

func (r *RuleResponse) GeneratedResource() unstructured.Unstructured {
	return r.generatedResource
}

func (r *RuleResponse) Message() string {
	return r.message
}

func (r *RuleResponse) Name() string {
	return r.name
}

func (r *RuleResponse) RuleType() RuleType {
	return r.ruleType
}

func (r *RuleResponse) Status() RuleStatus {
	return r.status
}

func (r *RuleResponse) EmitWarning() bool {
	return r.emitWarning
}

// HasStatus checks if rule status is in a given list
func (r *RuleResponse) HasStatus(status ...RuleStatus) bool {
	for _, s := range status {
		if r.status == s {
			return true
		}
	}
	return false
}

// String implements Stringer interface
func (r *RuleResponse) String() string {
	return fmt.Sprintf("rule %s (%s): %v", r.name, r.ruleType, r.message)
}
