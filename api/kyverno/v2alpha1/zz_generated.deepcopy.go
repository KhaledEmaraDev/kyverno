//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by deepcopy-gen. DO NOT EDIT.

package v2alpha1

import (
	v1 "k8s.io/api/admissionregistration/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutogenRule) DeepCopyInto(out *AutogenRule) {
	*out = *in
	if in.MatchConstraints != nil {
		in, out := &in.MatchConstraints, &out.MatchConstraints
		*out = new(v1.MatchResources)
		(*in).DeepCopyInto(*out)
	}
	if in.MatchConditions != nil {
		in, out := &in.MatchConditions, &out.MatchConditions
		*out = make([]v1.MatchCondition, len(*in))
		copy(*out, *in)
	}
	if in.Validations != nil {
		in, out := &in.Validations, &out.Validations
		*out = make([]v1.Validation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AuditAnnotation != nil {
		in, out := &in.AuditAnnotation, &out.AuditAnnotation
		*out = make([]v1.AuditAnnotation, len(*in))
		copy(*out, *in)
	}
	if in.Variables != nil {
		in, out := &in.Variables, &out.Variables
		*out = make([]v1.Variable, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutogenRule.
func (in *AutogenRule) DeepCopy() *AutogenRule {
	if in == nil {
		return nil
	}
	out := new(AutogenRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutogenStatus) DeepCopyInto(out *AutogenStatus) {
	*out = *in
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]AutogenRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutogenStatus.
func (in *AutogenStatus) DeepCopy() *AutogenStatus {
	if in == nil {
		return nil
	}
	out := new(AutogenStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CELPolicyException) DeepCopyInto(out *CELPolicyException) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CELPolicyException.
func (in *CELPolicyException) DeepCopy() *CELPolicyException {
	if in == nil {
		return nil
	}
	out := new(CELPolicyException)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CELPolicyException) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CELPolicyExceptionList) DeepCopyInto(out *CELPolicyExceptionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CELPolicyException, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CELPolicyExceptionList.
func (in *CELPolicyExceptionList) DeepCopy() *CELPolicyExceptionList {
	if in == nil {
		return nil
	}
	out := new(CELPolicyExceptionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CELPolicyExceptionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CELPolicyExceptionSpec) DeepCopyInto(out *CELPolicyExceptionSpec) {
	*out = *in
	if in.PolicyRefs != nil {
		in, out := &in.PolicyRefs, &out.PolicyRefs
		*out = make([]PolicyRef, len(*in))
		copy(*out, *in)
	}
	if in.MatchConditions != nil {
		in, out := &in.MatchConditions, &out.MatchConditions
		*out = make([]v1.MatchCondition, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CELPolicyExceptionSpec.
func (in *CELPolicyExceptionSpec) DeepCopy() *CELPolicyExceptionSpec {
	if in == nil {
		return nil
	}
	out := new(CELPolicyExceptionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalAPICall) DeepCopyInto(out *ExternalAPICall) {
	*out = *in
	in.APICall.DeepCopyInto(&out.APICall)
	if in.RefreshInterval != nil {
		in, out := &in.RefreshInterval, &out.RefreshInterval
		*out = new(metav1.Duration)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalAPICall.
func (in *ExternalAPICall) DeepCopy() *ExternalAPICall {
	if in == nil {
		return nil
	}
	out := new(ExternalAPICall)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlobalContextEntry) DeepCopyInto(out *GlobalContextEntry) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlobalContextEntry.
func (in *GlobalContextEntry) DeepCopy() *GlobalContextEntry {
	if in == nil {
		return nil
	}
	out := new(GlobalContextEntry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GlobalContextEntry) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlobalContextEntryList) DeepCopyInto(out *GlobalContextEntryList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GlobalContextEntry, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlobalContextEntryList.
func (in *GlobalContextEntryList) DeepCopy() *GlobalContextEntryList {
	if in == nil {
		return nil
	}
	out := new(GlobalContextEntryList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GlobalContextEntryList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlobalContextEntryProjection) DeepCopyInto(out *GlobalContextEntryProjection) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlobalContextEntryProjection.
func (in *GlobalContextEntryProjection) DeepCopy() *GlobalContextEntryProjection {
	if in == nil {
		return nil
	}
	out := new(GlobalContextEntryProjection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlobalContextEntrySpec) DeepCopyInto(out *GlobalContextEntrySpec) {
	*out = *in
	if in.KubernetesResource != nil {
		in, out := &in.KubernetesResource, &out.KubernetesResource
		*out = new(KubernetesResource)
		**out = **in
	}
	if in.APICall != nil {
		in, out := &in.APICall, &out.APICall
		*out = new(ExternalAPICall)
		(*in).DeepCopyInto(*out)
	}
	if in.Projections != nil {
		in, out := &in.Projections, &out.Projections
		*out = make([]GlobalContextEntryProjection, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlobalContextEntrySpec.
func (in *GlobalContextEntrySpec) DeepCopy() *GlobalContextEntrySpec {
	if in == nil {
		return nil
	}
	out := new(GlobalContextEntrySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlobalContextEntryStatus) DeepCopyInto(out *GlobalContextEntryStatus) {
	*out = *in
	if in.Ready != nil {
		in, out := &in.Ready, &out.Ready
		*out = new(bool)
		**out = **in
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.LastRefreshTime.DeepCopyInto(&out.LastRefreshTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlobalContextEntryStatus.
func (in *GlobalContextEntryStatus) DeepCopy() *GlobalContextEntryStatus {
	if in == nil {
		return nil
	}
	out := new(GlobalContextEntryStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubernetesResource) DeepCopyInto(out *KubernetesResource) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubernetesResource.
func (in *KubernetesResource) DeepCopy() *KubernetesResource {
	if in == nil {
		return nil
	}
	out := new(KubernetesResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyRef) DeepCopyInto(out *PolicyRef) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyRef.
func (in *PolicyRef) DeepCopy() *PolicyRef {
	if in == nil {
		return nil
	}
	out := new(PolicyRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyStatus) DeepCopyInto(out *PolicyStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Autogen.DeepCopyInto(&out.Autogen)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyStatus.
func (in *PolicyStatus) DeepCopy() *PolicyStatus {
	if in == nil {
		return nil
	}
	out := new(PolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ValidatingPolicy) DeepCopyInto(out *ValidatingPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ValidatingPolicy.
func (in *ValidatingPolicy) DeepCopy() *ValidatingPolicy {
	if in == nil {
		return nil
	}
	out := new(ValidatingPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ValidatingPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ValidatingPolicyList) DeepCopyInto(out *ValidatingPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ValidatingPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ValidatingPolicyList.
func (in *ValidatingPolicyList) DeepCopy() *ValidatingPolicyList {
	if in == nil {
		return nil
	}
	out := new(ValidatingPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ValidatingPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ValidatingPolicySpec) DeepCopyInto(out *ValidatingPolicySpec) {
	*out = *in
	in.ValidatingAdmissionPolicySpec.DeepCopyInto(&out.ValidatingAdmissionPolicySpec)
	if in.ValidationAction != nil {
		in, out := &in.ValidationAction, &out.ValidationAction
		*out = make([]v1.ValidationAction, len(*in))
		copy(*out, *in)
	}
	if in.WebhookConfiguration != nil {
		in, out := &in.WebhookConfiguration, &out.WebhookConfiguration
		*out = new(WebhookConfiguration)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ValidatingPolicySpec.
func (in *ValidatingPolicySpec) DeepCopy() *ValidatingPolicySpec {
	if in == nil {
		return nil
	}
	out := new(ValidatingPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebhookConfiguration) DeepCopyInto(out *WebhookConfiguration) {
	*out = *in
	if in.TimeoutSeconds != nil {
		in, out := &in.TimeoutSeconds, &out.TimeoutSeconds
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebhookConfiguration.
func (in *WebhookConfiguration) DeepCopy() *WebhookConfiguration {
	if in == nil {
		return nil
	}
	out := new(WebhookConfiguration)
	in.DeepCopyInto(out)
	return out
}
