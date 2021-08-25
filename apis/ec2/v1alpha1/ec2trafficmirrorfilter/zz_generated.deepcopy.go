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
func (in *Ec2TrafficMirrorFilter) DeepCopyInto(out *Ec2TrafficMirrorFilter) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Ec2TrafficMirrorFilter.
func (in *Ec2TrafficMirrorFilter) DeepCopy() *Ec2TrafficMirrorFilter {
	if in == nil {
		return nil
	}
	out := new(Ec2TrafficMirrorFilter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Ec2TrafficMirrorFilter) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Ec2TrafficMirrorFilterList) DeepCopyInto(out *Ec2TrafficMirrorFilterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Ec2TrafficMirrorFilter, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Ec2TrafficMirrorFilterList.
func (in *Ec2TrafficMirrorFilterList) DeepCopy() *Ec2TrafficMirrorFilterList {
	if in == nil {
		return nil
	}
	out := new(Ec2TrafficMirrorFilterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Ec2TrafficMirrorFilterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Ec2TrafficMirrorFilterObservation) DeepCopyInto(out *Ec2TrafficMirrorFilterObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Ec2TrafficMirrorFilterObservation.
func (in *Ec2TrafficMirrorFilterObservation) DeepCopy() *Ec2TrafficMirrorFilterObservation {
	if in == nil {
		return nil
	}
	out := new(Ec2TrafficMirrorFilterObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Ec2TrafficMirrorFilterParameters) DeepCopyInto(out *Ec2TrafficMirrorFilterParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.NetworkServices != nil {
		in, out := &in.NetworkServices, &out.NetworkServices
		*out = make([]string, len(*in))
		copy(*out, *in)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Ec2TrafficMirrorFilterParameters.
func (in *Ec2TrafficMirrorFilterParameters) DeepCopy() *Ec2TrafficMirrorFilterParameters {
	if in == nil {
		return nil
	}
	out := new(Ec2TrafficMirrorFilterParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Ec2TrafficMirrorFilterSpec) DeepCopyInto(out *Ec2TrafficMirrorFilterSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Ec2TrafficMirrorFilterSpec.
func (in *Ec2TrafficMirrorFilterSpec) DeepCopy() *Ec2TrafficMirrorFilterSpec {
	if in == nil {
		return nil
	}
	out := new(Ec2TrafficMirrorFilterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Ec2TrafficMirrorFilterStatus) DeepCopyInto(out *Ec2TrafficMirrorFilterStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Ec2TrafficMirrorFilterStatus.
func (in *Ec2TrafficMirrorFilterStatus) DeepCopy() *Ec2TrafficMirrorFilterStatus {
	if in == nil {
		return nil
	}
	out := new(Ec2TrafficMirrorFilterStatus)
	in.DeepCopyInto(out)
	return out
}