// +build !ignore_autogenerated

/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticsearchDomainPolicy) DeepCopyInto(out *ElasticsearchDomainPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticsearchDomainPolicy.
func (in *ElasticsearchDomainPolicy) DeepCopy() *ElasticsearchDomainPolicy {
	if in == nil {
		return nil
	}
	out := new(ElasticsearchDomainPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ElasticsearchDomainPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticsearchDomainPolicyList) DeepCopyInto(out *ElasticsearchDomainPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ElasticsearchDomainPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticsearchDomainPolicyList.
func (in *ElasticsearchDomainPolicyList) DeepCopy() *ElasticsearchDomainPolicyList {
	if in == nil {
		return nil
	}
	out := new(ElasticsearchDomainPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ElasticsearchDomainPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticsearchDomainPolicyObservation) DeepCopyInto(out *ElasticsearchDomainPolicyObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticsearchDomainPolicyObservation.
func (in *ElasticsearchDomainPolicyObservation) DeepCopy() *ElasticsearchDomainPolicyObservation {
	if in == nil {
		return nil
	}
	out := new(ElasticsearchDomainPolicyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticsearchDomainPolicyParameters) DeepCopyInto(out *ElasticsearchDomainPolicyParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticsearchDomainPolicyParameters.
func (in *ElasticsearchDomainPolicyParameters) DeepCopy() *ElasticsearchDomainPolicyParameters {
	if in == nil {
		return nil
	}
	out := new(ElasticsearchDomainPolicyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticsearchDomainPolicySpec) DeepCopyInto(out *ElasticsearchDomainPolicySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	out.ForProvider = in.ForProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticsearchDomainPolicySpec.
func (in *ElasticsearchDomainPolicySpec) DeepCopy() *ElasticsearchDomainPolicySpec {
	if in == nil {
		return nil
	}
	out := new(ElasticsearchDomainPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticsearchDomainPolicyStatus) DeepCopyInto(out *ElasticsearchDomainPolicyStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticsearchDomainPolicyStatus.
func (in *ElasticsearchDomainPolicyStatus) DeepCopy() *ElasticsearchDomainPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(ElasticsearchDomainPolicyStatus)
	in.DeepCopyInto(out)
	return out
}