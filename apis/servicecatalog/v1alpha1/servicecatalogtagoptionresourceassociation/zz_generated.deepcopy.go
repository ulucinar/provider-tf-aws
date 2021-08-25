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
func (in *ServicecatalogTagOptionResourceAssociation) DeepCopyInto(out *ServicecatalogTagOptionResourceAssociation) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServicecatalogTagOptionResourceAssociation.
func (in *ServicecatalogTagOptionResourceAssociation) DeepCopy() *ServicecatalogTagOptionResourceAssociation {
	if in == nil {
		return nil
	}
	out := new(ServicecatalogTagOptionResourceAssociation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServicecatalogTagOptionResourceAssociation) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServicecatalogTagOptionResourceAssociationList) DeepCopyInto(out *ServicecatalogTagOptionResourceAssociationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ServicecatalogTagOptionResourceAssociation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServicecatalogTagOptionResourceAssociationList.
func (in *ServicecatalogTagOptionResourceAssociationList) DeepCopy() *ServicecatalogTagOptionResourceAssociationList {
	if in == nil {
		return nil
	}
	out := new(ServicecatalogTagOptionResourceAssociationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServicecatalogTagOptionResourceAssociationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServicecatalogTagOptionResourceAssociationObservation) DeepCopyInto(out *ServicecatalogTagOptionResourceAssociationObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServicecatalogTagOptionResourceAssociationObservation.
func (in *ServicecatalogTagOptionResourceAssociationObservation) DeepCopy() *ServicecatalogTagOptionResourceAssociationObservation {
	if in == nil {
		return nil
	}
	out := new(ServicecatalogTagOptionResourceAssociationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServicecatalogTagOptionResourceAssociationParameters) DeepCopyInto(out *ServicecatalogTagOptionResourceAssociationParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServicecatalogTagOptionResourceAssociationParameters.
func (in *ServicecatalogTagOptionResourceAssociationParameters) DeepCopy() *ServicecatalogTagOptionResourceAssociationParameters {
	if in == nil {
		return nil
	}
	out := new(ServicecatalogTagOptionResourceAssociationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServicecatalogTagOptionResourceAssociationSpec) DeepCopyInto(out *ServicecatalogTagOptionResourceAssociationSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	out.ForProvider = in.ForProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServicecatalogTagOptionResourceAssociationSpec.
func (in *ServicecatalogTagOptionResourceAssociationSpec) DeepCopy() *ServicecatalogTagOptionResourceAssociationSpec {
	if in == nil {
		return nil
	}
	out := new(ServicecatalogTagOptionResourceAssociationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServicecatalogTagOptionResourceAssociationStatus) DeepCopyInto(out *ServicecatalogTagOptionResourceAssociationStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServicecatalogTagOptionResourceAssociationStatus.
func (in *ServicecatalogTagOptionResourceAssociationStatus) DeepCopy() *ServicecatalogTagOptionResourceAssociationStatus {
	if in == nil {
		return nil
	}
	out := new(ServicecatalogTagOptionResourceAssociationStatus)
	in.DeepCopyInto(out)
	return out
}