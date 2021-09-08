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
func (in *AbortIncompleteMultipartUploadObservation) DeepCopyInto(out *AbortIncompleteMultipartUploadObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AbortIncompleteMultipartUploadObservation.
func (in *AbortIncompleteMultipartUploadObservation) DeepCopy() *AbortIncompleteMultipartUploadObservation {
	if in == nil {
		return nil
	}
	out := new(AbortIncompleteMultipartUploadObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AbortIncompleteMultipartUploadParameters) DeepCopyInto(out *AbortIncompleteMultipartUploadParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AbortIncompleteMultipartUploadParameters.
func (in *AbortIncompleteMultipartUploadParameters) DeepCopy() *AbortIncompleteMultipartUploadParameters {
	if in == nil {
		return nil
	}
	out := new(AbortIncompleteMultipartUploadParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExpirationObservation) DeepCopyInto(out *ExpirationObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExpirationObservation.
func (in *ExpirationObservation) DeepCopy() *ExpirationObservation {
	if in == nil {
		return nil
	}
	out := new(ExpirationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExpirationParameters) DeepCopyInto(out *ExpirationParameters) {
	*out = *in
	if in.Date != nil {
		in, out := &in.Date, &out.Date
		*out = new(string)
		**out = **in
	}
	if in.Days != nil {
		in, out := &in.Days, &out.Days
		*out = new(int64)
		**out = **in
	}
	if in.ExpiredObjectDeleteMarker != nil {
		in, out := &in.ExpiredObjectDeleteMarker, &out.ExpiredObjectDeleteMarker
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExpirationParameters.
func (in *ExpirationParameters) DeepCopy() *ExpirationParameters {
	if in == nil {
		return nil
	}
	out := new(ExpirationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FilterObservation) DeepCopyInto(out *FilterObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FilterObservation.
func (in *FilterObservation) DeepCopy() *FilterObservation {
	if in == nil {
		return nil
	}
	out := new(FilterObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FilterParameters) DeepCopyInto(out *FilterParameters) {
	*out = *in
	if in.Prefix != nil {
		in, out := &in.Prefix, &out.Prefix
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
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FilterParameters.
func (in *FilterParameters) DeepCopy() *FilterParameters {
	if in == nil {
		return nil
	}
	out := new(FilterParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuleObservation) DeepCopyInto(out *RuleObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuleObservation.
func (in *RuleObservation) DeepCopy() *RuleObservation {
	if in == nil {
		return nil
	}
	out := new(RuleObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuleParameters) DeepCopyInto(out *RuleParameters) {
	*out = *in
	if in.AbortIncompleteMultipartUpload != nil {
		in, out := &in.AbortIncompleteMultipartUpload, &out.AbortIncompleteMultipartUpload
		*out = make([]AbortIncompleteMultipartUploadParameters, len(*in))
		copy(*out, *in)
	}
	if in.Expiration != nil {
		in, out := &in.Expiration, &out.Expiration
		*out = make([]ExpirationParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Filter != nil {
		in, out := &in.Filter, &out.Filter
		*out = make([]FilterParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuleParameters.
func (in *RuleParameters) DeepCopy() *RuleParameters {
	if in == nil {
		return nil
	}
	out := new(RuleParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3ControlBucket) DeepCopyInto(out *S3ControlBucket) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3ControlBucket.
func (in *S3ControlBucket) DeepCopy() *S3ControlBucket {
	if in == nil {
		return nil
	}
	out := new(S3ControlBucket)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *S3ControlBucket) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3ControlBucketLifecycleConfiguration) DeepCopyInto(out *S3ControlBucketLifecycleConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3ControlBucketLifecycleConfiguration.
func (in *S3ControlBucketLifecycleConfiguration) DeepCopy() *S3ControlBucketLifecycleConfiguration {
	if in == nil {
		return nil
	}
	out := new(S3ControlBucketLifecycleConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *S3ControlBucketLifecycleConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3ControlBucketLifecycleConfigurationList) DeepCopyInto(out *S3ControlBucketLifecycleConfigurationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]S3ControlBucketLifecycleConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3ControlBucketLifecycleConfigurationList.
func (in *S3ControlBucketLifecycleConfigurationList) DeepCopy() *S3ControlBucketLifecycleConfigurationList {
	if in == nil {
		return nil
	}
	out := new(S3ControlBucketLifecycleConfigurationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *S3ControlBucketLifecycleConfigurationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3ControlBucketLifecycleConfigurationObservation) DeepCopyInto(out *S3ControlBucketLifecycleConfigurationObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3ControlBucketLifecycleConfigurationObservation.
func (in *S3ControlBucketLifecycleConfigurationObservation) DeepCopy() *S3ControlBucketLifecycleConfigurationObservation {
	if in == nil {
		return nil
	}
	out := new(S3ControlBucketLifecycleConfigurationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3ControlBucketLifecycleConfigurationParameters) DeepCopyInto(out *S3ControlBucketLifecycleConfigurationParameters) {
	*out = *in
	if in.Rule != nil {
		in, out := &in.Rule, &out.Rule
		*out = make([]RuleParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3ControlBucketLifecycleConfigurationParameters.
func (in *S3ControlBucketLifecycleConfigurationParameters) DeepCopy() *S3ControlBucketLifecycleConfigurationParameters {
	if in == nil {
		return nil
	}
	out := new(S3ControlBucketLifecycleConfigurationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3ControlBucketLifecycleConfigurationSpec) DeepCopyInto(out *S3ControlBucketLifecycleConfigurationSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3ControlBucketLifecycleConfigurationSpec.
func (in *S3ControlBucketLifecycleConfigurationSpec) DeepCopy() *S3ControlBucketLifecycleConfigurationSpec {
	if in == nil {
		return nil
	}
	out := new(S3ControlBucketLifecycleConfigurationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3ControlBucketLifecycleConfigurationStatus) DeepCopyInto(out *S3ControlBucketLifecycleConfigurationStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3ControlBucketLifecycleConfigurationStatus.
func (in *S3ControlBucketLifecycleConfigurationStatus) DeepCopy() *S3ControlBucketLifecycleConfigurationStatus {
	if in == nil {
		return nil
	}
	out := new(S3ControlBucketLifecycleConfigurationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3ControlBucketList) DeepCopyInto(out *S3ControlBucketList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]S3ControlBucket, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3ControlBucketList.
func (in *S3ControlBucketList) DeepCopy() *S3ControlBucketList {
	if in == nil {
		return nil
	}
	out := new(S3ControlBucketList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *S3ControlBucketList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3ControlBucketObservation) DeepCopyInto(out *S3ControlBucketObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3ControlBucketObservation.
func (in *S3ControlBucketObservation) DeepCopy() *S3ControlBucketObservation {
	if in == nil {
		return nil
	}
	out := new(S3ControlBucketObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3ControlBucketParameters) DeepCopyInto(out *S3ControlBucketParameters) {
	*out = *in
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3ControlBucketParameters.
func (in *S3ControlBucketParameters) DeepCopy() *S3ControlBucketParameters {
	if in == nil {
		return nil
	}
	out := new(S3ControlBucketParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3ControlBucketPolicy) DeepCopyInto(out *S3ControlBucketPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3ControlBucketPolicy.
func (in *S3ControlBucketPolicy) DeepCopy() *S3ControlBucketPolicy {
	if in == nil {
		return nil
	}
	out := new(S3ControlBucketPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *S3ControlBucketPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3ControlBucketPolicyList) DeepCopyInto(out *S3ControlBucketPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]S3ControlBucketPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3ControlBucketPolicyList.
func (in *S3ControlBucketPolicyList) DeepCopy() *S3ControlBucketPolicyList {
	if in == nil {
		return nil
	}
	out := new(S3ControlBucketPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *S3ControlBucketPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3ControlBucketPolicyObservation) DeepCopyInto(out *S3ControlBucketPolicyObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3ControlBucketPolicyObservation.
func (in *S3ControlBucketPolicyObservation) DeepCopy() *S3ControlBucketPolicyObservation {
	if in == nil {
		return nil
	}
	out := new(S3ControlBucketPolicyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3ControlBucketPolicyParameters) DeepCopyInto(out *S3ControlBucketPolicyParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3ControlBucketPolicyParameters.
func (in *S3ControlBucketPolicyParameters) DeepCopy() *S3ControlBucketPolicyParameters {
	if in == nil {
		return nil
	}
	out := new(S3ControlBucketPolicyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3ControlBucketPolicySpec) DeepCopyInto(out *S3ControlBucketPolicySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	out.ForProvider = in.ForProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3ControlBucketPolicySpec.
func (in *S3ControlBucketPolicySpec) DeepCopy() *S3ControlBucketPolicySpec {
	if in == nil {
		return nil
	}
	out := new(S3ControlBucketPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3ControlBucketPolicyStatus) DeepCopyInto(out *S3ControlBucketPolicyStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3ControlBucketPolicyStatus.
func (in *S3ControlBucketPolicyStatus) DeepCopy() *S3ControlBucketPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(S3ControlBucketPolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3ControlBucketSpec) DeepCopyInto(out *S3ControlBucketSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3ControlBucketSpec.
func (in *S3ControlBucketSpec) DeepCopy() *S3ControlBucketSpec {
	if in == nil {
		return nil
	}
	out := new(S3ControlBucketSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3ControlBucketStatus) DeepCopyInto(out *S3ControlBucketStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3ControlBucketStatus.
func (in *S3ControlBucketStatus) DeepCopy() *S3ControlBucketStatus {
	if in == nil {
		return nil
	}
	out := new(S3ControlBucketStatus)
	in.DeepCopyInto(out)
	return out
}