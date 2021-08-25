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
func (in *DbProxyTarget) DeepCopyInto(out *DbProxyTarget) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DbProxyTarget.
func (in *DbProxyTarget) DeepCopy() *DbProxyTarget {
	if in == nil {
		return nil
	}
	out := new(DbProxyTarget)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DbProxyTarget) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DbProxyTargetList) DeepCopyInto(out *DbProxyTargetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DbProxyTarget, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DbProxyTargetList.
func (in *DbProxyTargetList) DeepCopy() *DbProxyTargetList {
	if in == nil {
		return nil
	}
	out := new(DbProxyTargetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DbProxyTargetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DbProxyTargetObservation) DeepCopyInto(out *DbProxyTargetObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DbProxyTargetObservation.
func (in *DbProxyTargetObservation) DeepCopy() *DbProxyTargetObservation {
	if in == nil {
		return nil
	}
	out := new(DbProxyTargetObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DbProxyTargetParameters) DeepCopyInto(out *DbProxyTargetParameters) {
	*out = *in
	if in.DbClusterIdentifier != nil {
		in, out := &in.DbClusterIdentifier, &out.DbClusterIdentifier
		*out = new(string)
		**out = **in
	}
	if in.DbInstanceIdentifier != nil {
		in, out := &in.DbInstanceIdentifier, &out.DbInstanceIdentifier
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DbProxyTargetParameters.
func (in *DbProxyTargetParameters) DeepCopy() *DbProxyTargetParameters {
	if in == nil {
		return nil
	}
	out := new(DbProxyTargetParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DbProxyTargetSpec) DeepCopyInto(out *DbProxyTargetSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DbProxyTargetSpec.
func (in *DbProxyTargetSpec) DeepCopy() *DbProxyTargetSpec {
	if in == nil {
		return nil
	}
	out := new(DbProxyTargetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DbProxyTargetStatus) DeepCopyInto(out *DbProxyTargetStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DbProxyTargetStatus.
func (in *DbProxyTargetStatus) DeepCopy() *DbProxyTargetStatus {
	if in == nil {
		return nil
	}
	out := new(DbProxyTargetStatus)
	in.DeepCopyInto(out)
	return out
}