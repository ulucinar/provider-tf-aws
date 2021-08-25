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
func (in *GrantObservation) DeepCopyInto(out *GrantObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrantObservation.
func (in *GrantObservation) DeepCopy() *GrantObservation {
	if in == nil {
		return nil
	}
	out := new(GrantObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrantParameters) DeepCopyInto(out *GrantParameters) {
	*out = *in
	if in.Email != nil {
		in, out := &in.Email, &out.Email
		*out = new(string)
		**out = **in
	}
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
	if in.Permissions != nil {
		in, out := &in.Permissions, &out.Permissions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Uri != nil {
		in, out := &in.Uri, &out.Uri
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrantParameters.
func (in *GrantParameters) DeepCopy() *GrantParameters {
	if in == nil {
		return nil
	}
	out := new(GrantParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3ObjectCopy) DeepCopyInto(out *S3ObjectCopy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3ObjectCopy.
func (in *S3ObjectCopy) DeepCopy() *S3ObjectCopy {
	if in == nil {
		return nil
	}
	out := new(S3ObjectCopy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *S3ObjectCopy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3ObjectCopyList) DeepCopyInto(out *S3ObjectCopyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]S3ObjectCopy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3ObjectCopyList.
func (in *S3ObjectCopyList) DeepCopy() *S3ObjectCopyList {
	if in == nil {
		return nil
	}
	out := new(S3ObjectCopyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *S3ObjectCopyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3ObjectCopyObservation) DeepCopyInto(out *S3ObjectCopyObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3ObjectCopyObservation.
func (in *S3ObjectCopyObservation) DeepCopy() *S3ObjectCopyObservation {
	if in == nil {
		return nil
	}
	out := new(S3ObjectCopyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3ObjectCopyParameters) DeepCopyInto(out *S3ObjectCopyParameters) {
	*out = *in
	if in.Acl != nil {
		in, out := &in.Acl, &out.Acl
		*out = new(string)
		**out = **in
	}
	if in.BucketKeyEnabled != nil {
		in, out := &in.BucketKeyEnabled, &out.BucketKeyEnabled
		*out = new(bool)
		**out = **in
	}
	if in.CacheControl != nil {
		in, out := &in.CacheControl, &out.CacheControl
		*out = new(string)
		**out = **in
	}
	if in.ContentDisposition != nil {
		in, out := &in.ContentDisposition, &out.ContentDisposition
		*out = new(string)
		**out = **in
	}
	if in.ContentEncoding != nil {
		in, out := &in.ContentEncoding, &out.ContentEncoding
		*out = new(string)
		**out = **in
	}
	if in.ContentLanguage != nil {
		in, out := &in.ContentLanguage, &out.ContentLanguage
		*out = new(string)
		**out = **in
	}
	if in.ContentType != nil {
		in, out := &in.ContentType, &out.ContentType
		*out = new(string)
		**out = **in
	}
	if in.CopyIfMatch != nil {
		in, out := &in.CopyIfMatch, &out.CopyIfMatch
		*out = new(string)
		**out = **in
	}
	if in.CopyIfModifiedSince != nil {
		in, out := &in.CopyIfModifiedSince, &out.CopyIfModifiedSince
		*out = new(string)
		**out = **in
	}
	if in.CopyIfNoneMatch != nil {
		in, out := &in.CopyIfNoneMatch, &out.CopyIfNoneMatch
		*out = new(string)
		**out = **in
	}
	if in.CopyIfUnmodifiedSince != nil {
		in, out := &in.CopyIfUnmodifiedSince, &out.CopyIfUnmodifiedSince
		*out = new(string)
		**out = **in
	}
	if in.CustomerAlgorithm != nil {
		in, out := &in.CustomerAlgorithm, &out.CustomerAlgorithm
		*out = new(string)
		**out = **in
	}
	if in.CustomerKey != nil {
		in, out := &in.CustomerKey, &out.CustomerKey
		*out = new(string)
		**out = **in
	}
	if in.CustomerKeyMd5 != nil {
		in, out := &in.CustomerKeyMd5, &out.CustomerKeyMd5
		*out = new(string)
		**out = **in
	}
	if in.ExpectedBucketOwner != nil {
		in, out := &in.ExpectedBucketOwner, &out.ExpectedBucketOwner
		*out = new(string)
		**out = **in
	}
	if in.ExpectedSourceBucketOwner != nil {
		in, out := &in.ExpectedSourceBucketOwner, &out.ExpectedSourceBucketOwner
		*out = new(string)
		**out = **in
	}
	if in.Expires != nil {
		in, out := &in.Expires, &out.Expires
		*out = new(string)
		**out = **in
	}
	if in.ForceDestroy != nil {
		in, out := &in.ForceDestroy, &out.ForceDestroy
		*out = new(bool)
		**out = **in
	}
	if in.Grant != nil {
		in, out := &in.Grant, &out.Grant
		*out = make([]GrantParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.KmsEncryptionContext != nil {
		in, out := &in.KmsEncryptionContext, &out.KmsEncryptionContext
		*out = new(string)
		**out = **in
	}
	if in.KmsKeyId != nil {
		in, out := &in.KmsKeyId, &out.KmsKeyId
		*out = new(string)
		**out = **in
	}
	if in.Metadata != nil {
		in, out := &in.Metadata, &out.Metadata
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.MetadataDirective != nil {
		in, out := &in.MetadataDirective, &out.MetadataDirective
		*out = new(string)
		**out = **in
	}
	if in.ObjectLockLegalHoldStatus != nil {
		in, out := &in.ObjectLockLegalHoldStatus, &out.ObjectLockLegalHoldStatus
		*out = new(string)
		**out = **in
	}
	if in.ObjectLockMode != nil {
		in, out := &in.ObjectLockMode, &out.ObjectLockMode
		*out = new(string)
		**out = **in
	}
	if in.ObjectLockRetainUntilDate != nil {
		in, out := &in.ObjectLockRetainUntilDate, &out.ObjectLockRetainUntilDate
		*out = new(string)
		**out = **in
	}
	if in.RequestPayer != nil {
		in, out := &in.RequestPayer, &out.RequestPayer
		*out = new(string)
		**out = **in
	}
	if in.ServerSideEncryption != nil {
		in, out := &in.ServerSideEncryption, &out.ServerSideEncryption
		*out = new(string)
		**out = **in
	}
	if in.SourceCustomerAlgorithm != nil {
		in, out := &in.SourceCustomerAlgorithm, &out.SourceCustomerAlgorithm
		*out = new(string)
		**out = **in
	}
	if in.SourceCustomerKey != nil {
		in, out := &in.SourceCustomerKey, &out.SourceCustomerKey
		*out = new(string)
		**out = **in
	}
	if in.SourceCustomerKeyMd5 != nil {
		in, out := &in.SourceCustomerKeyMd5, &out.SourceCustomerKeyMd5
		*out = new(string)
		**out = **in
	}
	if in.StorageClass != nil {
		in, out := &in.StorageClass, &out.StorageClass
		*out = new(string)
		**out = **in
	}
	if in.TaggingDirective != nil {
		in, out := &in.TaggingDirective, &out.TaggingDirective
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
	if in.WebsiteRedirect != nil {
		in, out := &in.WebsiteRedirect, &out.WebsiteRedirect
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3ObjectCopyParameters.
func (in *S3ObjectCopyParameters) DeepCopy() *S3ObjectCopyParameters {
	if in == nil {
		return nil
	}
	out := new(S3ObjectCopyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3ObjectCopySpec) DeepCopyInto(out *S3ObjectCopySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3ObjectCopySpec.
func (in *S3ObjectCopySpec) DeepCopy() *S3ObjectCopySpec {
	if in == nil {
		return nil
	}
	out := new(S3ObjectCopySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3ObjectCopyStatus) DeepCopyInto(out *S3ObjectCopyStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3ObjectCopyStatus.
func (in *S3ObjectCopyStatus) DeepCopy() *S3ObjectCopyStatus {
	if in == nil {
		return nil
	}
	out := new(S3ObjectCopyStatus)
	in.DeepCopyInto(out)
	return out
}