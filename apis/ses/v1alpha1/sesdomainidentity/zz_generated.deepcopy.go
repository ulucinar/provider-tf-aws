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
func (in *SesDomainIdentity) DeepCopyInto(out *SesDomainIdentity) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SesDomainIdentity.
func (in *SesDomainIdentity) DeepCopy() *SesDomainIdentity {
	if in == nil {
		return nil
	}
	out := new(SesDomainIdentity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SesDomainIdentity) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SesDomainIdentityList) DeepCopyInto(out *SesDomainIdentityList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SesDomainIdentity, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SesDomainIdentityList.
func (in *SesDomainIdentityList) DeepCopy() *SesDomainIdentityList {
	if in == nil {
		return nil
	}
	out := new(SesDomainIdentityList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SesDomainIdentityList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SesDomainIdentityObservation) DeepCopyInto(out *SesDomainIdentityObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SesDomainIdentityObservation.
func (in *SesDomainIdentityObservation) DeepCopy() *SesDomainIdentityObservation {
	if in == nil {
		return nil
	}
	out := new(SesDomainIdentityObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SesDomainIdentityParameters) DeepCopyInto(out *SesDomainIdentityParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SesDomainIdentityParameters.
func (in *SesDomainIdentityParameters) DeepCopy() *SesDomainIdentityParameters {
	if in == nil {
		return nil
	}
	out := new(SesDomainIdentityParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SesDomainIdentitySpec) DeepCopyInto(out *SesDomainIdentitySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	out.ForProvider = in.ForProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SesDomainIdentitySpec.
func (in *SesDomainIdentitySpec) DeepCopy() *SesDomainIdentitySpec {
	if in == nil {
		return nil
	}
	out := new(SesDomainIdentitySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SesDomainIdentityStatus) DeepCopyInto(out *SesDomainIdentityStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SesDomainIdentityStatus.
func (in *SesDomainIdentityStatus) DeepCopy() *SesDomainIdentityStatus {
	if in == nil {
		return nil
	}
	out := new(SesDomainIdentityStatus)
	in.DeepCopyInto(out)
	return out
}