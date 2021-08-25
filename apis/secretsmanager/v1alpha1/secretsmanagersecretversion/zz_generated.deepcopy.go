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
func (in *SecretsmanagerSecretVersion) DeepCopyInto(out *SecretsmanagerSecretVersion) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretsmanagerSecretVersion.
func (in *SecretsmanagerSecretVersion) DeepCopy() *SecretsmanagerSecretVersion {
	if in == nil {
		return nil
	}
	out := new(SecretsmanagerSecretVersion)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SecretsmanagerSecretVersion) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretsmanagerSecretVersionList) DeepCopyInto(out *SecretsmanagerSecretVersionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SecretsmanagerSecretVersion, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretsmanagerSecretVersionList.
func (in *SecretsmanagerSecretVersionList) DeepCopy() *SecretsmanagerSecretVersionList {
	if in == nil {
		return nil
	}
	out := new(SecretsmanagerSecretVersionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SecretsmanagerSecretVersionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretsmanagerSecretVersionObservation) DeepCopyInto(out *SecretsmanagerSecretVersionObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretsmanagerSecretVersionObservation.
func (in *SecretsmanagerSecretVersionObservation) DeepCopy() *SecretsmanagerSecretVersionObservation {
	if in == nil {
		return nil
	}
	out := new(SecretsmanagerSecretVersionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretsmanagerSecretVersionParameters) DeepCopyInto(out *SecretsmanagerSecretVersionParameters) {
	*out = *in
	if in.SecretBinary != nil {
		in, out := &in.SecretBinary, &out.SecretBinary
		*out = new(string)
		**out = **in
	}
	if in.SecretString != nil {
		in, out := &in.SecretString, &out.SecretString
		*out = new(string)
		**out = **in
	}
	if in.VersionStages != nil {
		in, out := &in.VersionStages, &out.VersionStages
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretsmanagerSecretVersionParameters.
func (in *SecretsmanagerSecretVersionParameters) DeepCopy() *SecretsmanagerSecretVersionParameters {
	if in == nil {
		return nil
	}
	out := new(SecretsmanagerSecretVersionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretsmanagerSecretVersionSpec) DeepCopyInto(out *SecretsmanagerSecretVersionSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretsmanagerSecretVersionSpec.
func (in *SecretsmanagerSecretVersionSpec) DeepCopy() *SecretsmanagerSecretVersionSpec {
	if in == nil {
		return nil
	}
	out := new(SecretsmanagerSecretVersionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretsmanagerSecretVersionStatus) DeepCopyInto(out *SecretsmanagerSecretVersionStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretsmanagerSecretVersionStatus.
func (in *SecretsmanagerSecretVersionStatus) DeepCopy() *SecretsmanagerSecretVersionStatus {
	if in == nil {
		return nil
	}
	out := new(SecretsmanagerSecretVersionStatus)
	in.DeepCopyInto(out)
	return out
}