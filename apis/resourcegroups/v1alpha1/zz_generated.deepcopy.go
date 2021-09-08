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
func (in *ResourceQueryObservation) DeepCopyInto(out *ResourceQueryObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceQueryObservation.
func (in *ResourceQueryObservation) DeepCopy() *ResourceQueryObservation {
	if in == nil {
		return nil
	}
	out := new(ResourceQueryObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceQueryParameters) DeepCopyInto(out *ResourceQueryParameters) {
	*out = *in
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceQueryParameters.
func (in *ResourceQueryParameters) DeepCopy() *ResourceQueryParameters {
	if in == nil {
		return nil
	}
	out := new(ResourceQueryParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourcegroupsGroup) DeepCopyInto(out *ResourcegroupsGroup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourcegroupsGroup.
func (in *ResourcegroupsGroup) DeepCopy() *ResourcegroupsGroup {
	if in == nil {
		return nil
	}
	out := new(ResourcegroupsGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ResourcegroupsGroup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourcegroupsGroupList) DeepCopyInto(out *ResourcegroupsGroupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ResourcegroupsGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourcegroupsGroupList.
func (in *ResourcegroupsGroupList) DeepCopy() *ResourcegroupsGroupList {
	if in == nil {
		return nil
	}
	out := new(ResourcegroupsGroupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ResourcegroupsGroupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourcegroupsGroupObservation) DeepCopyInto(out *ResourcegroupsGroupObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourcegroupsGroupObservation.
func (in *ResourcegroupsGroupObservation) DeepCopy() *ResourcegroupsGroupObservation {
	if in == nil {
		return nil
	}
	out := new(ResourcegroupsGroupObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourcegroupsGroupParameters) DeepCopyInto(out *ResourcegroupsGroupParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.ResourceQuery != nil {
		in, out := &in.ResourceQuery, &out.ResourceQuery
		*out = make([]ResourceQueryParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
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
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourcegroupsGroupParameters.
func (in *ResourcegroupsGroupParameters) DeepCopy() *ResourcegroupsGroupParameters {
	if in == nil {
		return nil
	}
	out := new(ResourcegroupsGroupParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourcegroupsGroupSpec) DeepCopyInto(out *ResourcegroupsGroupSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourcegroupsGroupSpec.
func (in *ResourcegroupsGroupSpec) DeepCopy() *ResourcegroupsGroupSpec {
	if in == nil {
		return nil
	}
	out := new(ResourcegroupsGroupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourcegroupsGroupStatus) DeepCopyInto(out *ResourcegroupsGroupStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourcegroupsGroupStatus.
func (in *ResourcegroupsGroupStatus) DeepCopy() *ResourcegroupsGroupStatus {
	if in == nil {
		return nil
	}
	out := new(ResourcegroupsGroupStatus)
	in.DeepCopyInto(out)
	return out
}