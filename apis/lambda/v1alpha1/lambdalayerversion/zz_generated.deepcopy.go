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
func (in *LambdaLayerVersion) DeepCopyInto(out *LambdaLayerVersion) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LambdaLayerVersion.
func (in *LambdaLayerVersion) DeepCopy() *LambdaLayerVersion {
	if in == nil {
		return nil
	}
	out := new(LambdaLayerVersion)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LambdaLayerVersion) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LambdaLayerVersionList) DeepCopyInto(out *LambdaLayerVersionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LambdaLayerVersion, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LambdaLayerVersionList.
func (in *LambdaLayerVersionList) DeepCopy() *LambdaLayerVersionList {
	if in == nil {
		return nil
	}
	out := new(LambdaLayerVersionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LambdaLayerVersionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LambdaLayerVersionObservation) DeepCopyInto(out *LambdaLayerVersionObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LambdaLayerVersionObservation.
func (in *LambdaLayerVersionObservation) DeepCopy() *LambdaLayerVersionObservation {
	if in == nil {
		return nil
	}
	out := new(LambdaLayerVersionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LambdaLayerVersionParameters) DeepCopyInto(out *LambdaLayerVersionParameters) {
	*out = *in
	if in.CompatibleRuntimes != nil {
		in, out := &in.CompatibleRuntimes, &out.CompatibleRuntimes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Filename != nil {
		in, out := &in.Filename, &out.Filename
		*out = new(string)
		**out = **in
	}
	if in.LicenseInfo != nil {
		in, out := &in.LicenseInfo, &out.LicenseInfo
		*out = new(string)
		**out = **in
	}
	if in.S3Bucket != nil {
		in, out := &in.S3Bucket, &out.S3Bucket
		*out = new(string)
		**out = **in
	}
	if in.S3Key != nil {
		in, out := &in.S3Key, &out.S3Key
		*out = new(string)
		**out = **in
	}
	if in.S3ObjectVersion != nil {
		in, out := &in.S3ObjectVersion, &out.S3ObjectVersion
		*out = new(string)
		**out = **in
	}
	if in.SourceCodeHash != nil {
		in, out := &in.SourceCodeHash, &out.SourceCodeHash
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LambdaLayerVersionParameters.
func (in *LambdaLayerVersionParameters) DeepCopy() *LambdaLayerVersionParameters {
	if in == nil {
		return nil
	}
	out := new(LambdaLayerVersionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LambdaLayerVersionSpec) DeepCopyInto(out *LambdaLayerVersionSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LambdaLayerVersionSpec.
func (in *LambdaLayerVersionSpec) DeepCopy() *LambdaLayerVersionSpec {
	if in == nil {
		return nil
	}
	out := new(LambdaLayerVersionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LambdaLayerVersionStatus) DeepCopyInto(out *LambdaLayerVersionStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LambdaLayerVersionStatus.
func (in *LambdaLayerVersionStatus) DeepCopy() *LambdaLayerVersionStatus {
	if in == nil {
		return nil
	}
	out := new(LambdaLayerVersionStatus)
	in.DeepCopyInto(out)
	return out
}