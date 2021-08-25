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
func (in *RevocationRecordObservation) DeepCopyInto(out *RevocationRecordObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RevocationRecordObservation.
func (in *RevocationRecordObservation) DeepCopy() *RevocationRecordObservation {
	if in == nil {
		return nil
	}
	out := new(RevocationRecordObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RevocationRecordParameters) DeepCopyInto(out *RevocationRecordParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RevocationRecordParameters.
func (in *RevocationRecordParameters) DeepCopy() *RevocationRecordParameters {
	if in == nil {
		return nil
	}
	out := new(RevocationRecordParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SignatureValidityPeriodObservation) DeepCopyInto(out *SignatureValidityPeriodObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SignatureValidityPeriodObservation.
func (in *SignatureValidityPeriodObservation) DeepCopy() *SignatureValidityPeriodObservation {
	if in == nil {
		return nil
	}
	out := new(SignatureValidityPeriodObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SignatureValidityPeriodParameters) DeepCopyInto(out *SignatureValidityPeriodParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SignatureValidityPeriodParameters.
func (in *SignatureValidityPeriodParameters) DeepCopy() *SignatureValidityPeriodParameters {
	if in == nil {
		return nil
	}
	out := new(SignatureValidityPeriodParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SignerSigningProfile) DeepCopyInto(out *SignerSigningProfile) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SignerSigningProfile.
func (in *SignerSigningProfile) DeepCopy() *SignerSigningProfile {
	if in == nil {
		return nil
	}
	out := new(SignerSigningProfile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SignerSigningProfile) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SignerSigningProfileList) DeepCopyInto(out *SignerSigningProfileList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SignerSigningProfile, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SignerSigningProfileList.
func (in *SignerSigningProfileList) DeepCopy() *SignerSigningProfileList {
	if in == nil {
		return nil
	}
	out := new(SignerSigningProfileList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SignerSigningProfileList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SignerSigningProfileObservation) DeepCopyInto(out *SignerSigningProfileObservation) {
	*out = *in
	if in.RevocationRecord != nil {
		in, out := &in.RevocationRecord, &out.RevocationRecord
		*out = make([]RevocationRecordObservation, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SignerSigningProfileObservation.
func (in *SignerSigningProfileObservation) DeepCopy() *SignerSigningProfileObservation {
	if in == nil {
		return nil
	}
	out := new(SignerSigningProfileObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SignerSigningProfileParameters) DeepCopyInto(out *SignerSigningProfileParameters) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NamePrefix != nil {
		in, out := &in.NamePrefix, &out.NamePrefix
		*out = new(string)
		**out = **in
	}
	if in.SignatureValidityPeriod != nil {
		in, out := &in.SignatureValidityPeriod, &out.SignatureValidityPeriod
		*out = make([]SignatureValidityPeriodParameters, len(*in))
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SignerSigningProfileParameters.
func (in *SignerSigningProfileParameters) DeepCopy() *SignerSigningProfileParameters {
	if in == nil {
		return nil
	}
	out := new(SignerSigningProfileParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SignerSigningProfileSpec) DeepCopyInto(out *SignerSigningProfileSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SignerSigningProfileSpec.
func (in *SignerSigningProfileSpec) DeepCopy() *SignerSigningProfileSpec {
	if in == nil {
		return nil
	}
	out := new(SignerSigningProfileSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SignerSigningProfileStatus) DeepCopyInto(out *SignerSigningProfileStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SignerSigningProfileStatus.
func (in *SignerSigningProfileStatus) DeepCopy() *SignerSigningProfileStatus {
	if in == nil {
		return nil
	}
	out := new(SignerSigningProfileStatus)
	in.DeepCopyInto(out)
	return out
}