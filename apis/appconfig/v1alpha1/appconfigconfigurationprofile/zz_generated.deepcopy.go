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
func (in *AppconfigConfigurationProfile) DeepCopyInto(out *AppconfigConfigurationProfile) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppconfigConfigurationProfile.
func (in *AppconfigConfigurationProfile) DeepCopy() *AppconfigConfigurationProfile {
	if in == nil {
		return nil
	}
	out := new(AppconfigConfigurationProfile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AppconfigConfigurationProfile) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppconfigConfigurationProfileList) DeepCopyInto(out *AppconfigConfigurationProfileList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AppconfigConfigurationProfile, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppconfigConfigurationProfileList.
func (in *AppconfigConfigurationProfileList) DeepCopy() *AppconfigConfigurationProfileList {
	if in == nil {
		return nil
	}
	out := new(AppconfigConfigurationProfileList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AppconfigConfigurationProfileList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppconfigConfigurationProfileObservation) DeepCopyInto(out *AppconfigConfigurationProfileObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppconfigConfigurationProfileObservation.
func (in *AppconfigConfigurationProfileObservation) DeepCopy() *AppconfigConfigurationProfileObservation {
	if in == nil {
		return nil
	}
	out := new(AppconfigConfigurationProfileObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppconfigConfigurationProfileParameters) DeepCopyInto(out *AppconfigConfigurationProfileParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.RetrievalRoleArn != nil {
		in, out := &in.RetrievalRoleArn, &out.RetrievalRoleArn
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.TagsAll != nil {
		in, out := &in.TagsAll, &out.TagsAll
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Validator != nil {
		in, out := &in.Validator, &out.Validator
		*out = make([]ValidatorParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppconfigConfigurationProfileParameters.
func (in *AppconfigConfigurationProfileParameters) DeepCopy() *AppconfigConfigurationProfileParameters {
	if in == nil {
		return nil
	}
	out := new(AppconfigConfigurationProfileParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppconfigConfigurationProfileSpec) DeepCopyInto(out *AppconfigConfigurationProfileSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppconfigConfigurationProfileSpec.
func (in *AppconfigConfigurationProfileSpec) DeepCopy() *AppconfigConfigurationProfileSpec {
	if in == nil {
		return nil
	}
	out := new(AppconfigConfigurationProfileSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppconfigConfigurationProfileStatus) DeepCopyInto(out *AppconfigConfigurationProfileStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppconfigConfigurationProfileStatus.
func (in *AppconfigConfigurationProfileStatus) DeepCopy() *AppconfigConfigurationProfileStatus {
	if in == nil {
		return nil
	}
	out := new(AppconfigConfigurationProfileStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ValidatorObservation) DeepCopyInto(out *ValidatorObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ValidatorObservation.
func (in *ValidatorObservation) DeepCopy() *ValidatorObservation {
	if in == nil {
		return nil
	}
	out := new(ValidatorObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ValidatorParameters) DeepCopyInto(out *ValidatorParameters) {
	*out = *in
	if in.Content != nil {
		in, out := &in.Content, &out.Content
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ValidatorParameters.
func (in *ValidatorParameters) DeepCopy() *ValidatorParameters {
	if in == nil {
		return nil
	}
	out := new(ValidatorParameters)
	in.DeepCopyInto(out)
	return out
}