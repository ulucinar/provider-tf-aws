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
func (in *GlueUserDefinedFunction) DeepCopyInto(out *GlueUserDefinedFunction) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlueUserDefinedFunction.
func (in *GlueUserDefinedFunction) DeepCopy() *GlueUserDefinedFunction {
	if in == nil {
		return nil
	}
	out := new(GlueUserDefinedFunction)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GlueUserDefinedFunction) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlueUserDefinedFunctionList) DeepCopyInto(out *GlueUserDefinedFunctionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GlueUserDefinedFunction, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlueUserDefinedFunctionList.
func (in *GlueUserDefinedFunctionList) DeepCopy() *GlueUserDefinedFunctionList {
	if in == nil {
		return nil
	}
	out := new(GlueUserDefinedFunctionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GlueUserDefinedFunctionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlueUserDefinedFunctionObservation) DeepCopyInto(out *GlueUserDefinedFunctionObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlueUserDefinedFunctionObservation.
func (in *GlueUserDefinedFunctionObservation) DeepCopy() *GlueUserDefinedFunctionObservation {
	if in == nil {
		return nil
	}
	out := new(GlueUserDefinedFunctionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlueUserDefinedFunctionParameters) DeepCopyInto(out *GlueUserDefinedFunctionParameters) {
	*out = *in
	if in.CatalogId != nil {
		in, out := &in.CatalogId, &out.CatalogId
		*out = new(string)
		**out = **in
	}
	if in.ResourceUris != nil {
		in, out := &in.ResourceUris, &out.ResourceUris
		*out = make([]ResourceUrisParameters, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlueUserDefinedFunctionParameters.
func (in *GlueUserDefinedFunctionParameters) DeepCopy() *GlueUserDefinedFunctionParameters {
	if in == nil {
		return nil
	}
	out := new(GlueUserDefinedFunctionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlueUserDefinedFunctionSpec) DeepCopyInto(out *GlueUserDefinedFunctionSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlueUserDefinedFunctionSpec.
func (in *GlueUserDefinedFunctionSpec) DeepCopy() *GlueUserDefinedFunctionSpec {
	if in == nil {
		return nil
	}
	out := new(GlueUserDefinedFunctionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlueUserDefinedFunctionStatus) DeepCopyInto(out *GlueUserDefinedFunctionStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlueUserDefinedFunctionStatus.
func (in *GlueUserDefinedFunctionStatus) DeepCopy() *GlueUserDefinedFunctionStatus {
	if in == nil {
		return nil
	}
	out := new(GlueUserDefinedFunctionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceUrisObservation) DeepCopyInto(out *ResourceUrisObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceUrisObservation.
func (in *ResourceUrisObservation) DeepCopy() *ResourceUrisObservation {
	if in == nil {
		return nil
	}
	out := new(ResourceUrisObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceUrisParameters) DeepCopyInto(out *ResourceUrisParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceUrisParameters.
func (in *ResourceUrisParameters) DeepCopy() *ResourceUrisParameters {
	if in == nil {
		return nil
	}
	out := new(ResourceUrisParameters)
	in.DeepCopyInto(out)
	return out
}