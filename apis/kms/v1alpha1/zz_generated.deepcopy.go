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
func (in *ConstraintsObservation) DeepCopyInto(out *ConstraintsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConstraintsObservation.
func (in *ConstraintsObservation) DeepCopy() *ConstraintsObservation {
	if in == nil {
		return nil
	}
	out := new(ConstraintsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConstraintsParameters) DeepCopyInto(out *ConstraintsParameters) {
	*out = *in
	if in.EncryptionContextEquals != nil {
		in, out := &in.EncryptionContextEquals, &out.EncryptionContextEquals
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.EncryptionContextSubset != nil {
		in, out := &in.EncryptionContextSubset, &out.EncryptionContextSubset
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConstraintsParameters.
func (in *ConstraintsParameters) DeepCopy() *ConstraintsParameters {
	if in == nil {
		return nil
	}
	out := new(ConstraintsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KmsAlias) DeepCopyInto(out *KmsAlias) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KmsAlias.
func (in *KmsAlias) DeepCopy() *KmsAlias {
	if in == nil {
		return nil
	}
	out := new(KmsAlias)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KmsAlias) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KmsAliasList) DeepCopyInto(out *KmsAliasList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KmsAlias, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KmsAliasList.
func (in *KmsAliasList) DeepCopy() *KmsAliasList {
	if in == nil {
		return nil
	}
	out := new(KmsAliasList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KmsAliasList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KmsAliasObservation) DeepCopyInto(out *KmsAliasObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KmsAliasObservation.
func (in *KmsAliasObservation) DeepCopy() *KmsAliasObservation {
	if in == nil {
		return nil
	}
	out := new(KmsAliasObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KmsAliasParameters) DeepCopyInto(out *KmsAliasParameters) {
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
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KmsAliasParameters.
func (in *KmsAliasParameters) DeepCopy() *KmsAliasParameters {
	if in == nil {
		return nil
	}
	out := new(KmsAliasParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KmsAliasSpec) DeepCopyInto(out *KmsAliasSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KmsAliasSpec.
func (in *KmsAliasSpec) DeepCopy() *KmsAliasSpec {
	if in == nil {
		return nil
	}
	out := new(KmsAliasSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KmsAliasStatus) DeepCopyInto(out *KmsAliasStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KmsAliasStatus.
func (in *KmsAliasStatus) DeepCopy() *KmsAliasStatus {
	if in == nil {
		return nil
	}
	out := new(KmsAliasStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KmsCiphertext) DeepCopyInto(out *KmsCiphertext) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KmsCiphertext.
func (in *KmsCiphertext) DeepCopy() *KmsCiphertext {
	if in == nil {
		return nil
	}
	out := new(KmsCiphertext)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KmsCiphertext) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KmsCiphertextList) DeepCopyInto(out *KmsCiphertextList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KmsCiphertext, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KmsCiphertextList.
func (in *KmsCiphertextList) DeepCopy() *KmsCiphertextList {
	if in == nil {
		return nil
	}
	out := new(KmsCiphertextList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KmsCiphertextList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KmsCiphertextObservation) DeepCopyInto(out *KmsCiphertextObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KmsCiphertextObservation.
func (in *KmsCiphertextObservation) DeepCopy() *KmsCiphertextObservation {
	if in == nil {
		return nil
	}
	out := new(KmsCiphertextObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KmsCiphertextParameters) DeepCopyInto(out *KmsCiphertextParameters) {
	*out = *in
	if in.Context != nil {
		in, out := &in.Context, &out.Context
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KmsCiphertextParameters.
func (in *KmsCiphertextParameters) DeepCopy() *KmsCiphertextParameters {
	if in == nil {
		return nil
	}
	out := new(KmsCiphertextParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KmsCiphertextSpec) DeepCopyInto(out *KmsCiphertextSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KmsCiphertextSpec.
func (in *KmsCiphertextSpec) DeepCopy() *KmsCiphertextSpec {
	if in == nil {
		return nil
	}
	out := new(KmsCiphertextSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KmsCiphertextStatus) DeepCopyInto(out *KmsCiphertextStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KmsCiphertextStatus.
func (in *KmsCiphertextStatus) DeepCopy() *KmsCiphertextStatus {
	if in == nil {
		return nil
	}
	out := new(KmsCiphertextStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KmsExternalKey) DeepCopyInto(out *KmsExternalKey) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KmsExternalKey.
func (in *KmsExternalKey) DeepCopy() *KmsExternalKey {
	if in == nil {
		return nil
	}
	out := new(KmsExternalKey)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KmsExternalKey) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KmsExternalKeyList) DeepCopyInto(out *KmsExternalKeyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KmsExternalKey, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KmsExternalKeyList.
func (in *KmsExternalKeyList) DeepCopy() *KmsExternalKeyList {
	if in == nil {
		return nil
	}
	out := new(KmsExternalKeyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KmsExternalKeyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KmsExternalKeyObservation) DeepCopyInto(out *KmsExternalKeyObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KmsExternalKeyObservation.
func (in *KmsExternalKeyObservation) DeepCopy() *KmsExternalKeyObservation {
	if in == nil {
		return nil
	}
	out := new(KmsExternalKeyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KmsExternalKeyParameters) DeepCopyInto(out *KmsExternalKeyParameters) {
	*out = *in
	if in.BypassPolicyLockoutSafetyCheck != nil {
		in, out := &in.BypassPolicyLockoutSafetyCheck, &out.BypassPolicyLockoutSafetyCheck
		*out = new(bool)
		**out = **in
	}
	if in.DeletionWindowInDays != nil {
		in, out := &in.DeletionWindowInDays, &out.DeletionWindowInDays
		*out = new(int64)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.KeyMaterialBase64 != nil {
		in, out := &in.KeyMaterialBase64, &out.KeyMaterialBase64
		*out = new(string)
		**out = **in
	}
	if in.Policy != nil {
		in, out := &in.Policy, &out.Policy
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
	if in.ValidTo != nil {
		in, out := &in.ValidTo, &out.ValidTo
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KmsExternalKeyParameters.
func (in *KmsExternalKeyParameters) DeepCopy() *KmsExternalKeyParameters {
	if in == nil {
		return nil
	}
	out := new(KmsExternalKeyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KmsExternalKeySpec) DeepCopyInto(out *KmsExternalKeySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KmsExternalKeySpec.
func (in *KmsExternalKeySpec) DeepCopy() *KmsExternalKeySpec {
	if in == nil {
		return nil
	}
	out := new(KmsExternalKeySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KmsExternalKeyStatus) DeepCopyInto(out *KmsExternalKeyStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KmsExternalKeyStatus.
func (in *KmsExternalKeyStatus) DeepCopy() *KmsExternalKeyStatus {
	if in == nil {
		return nil
	}
	out := new(KmsExternalKeyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KmsGrant) DeepCopyInto(out *KmsGrant) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KmsGrant.
func (in *KmsGrant) DeepCopy() *KmsGrant {
	if in == nil {
		return nil
	}
	out := new(KmsGrant)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KmsGrant) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KmsGrantList) DeepCopyInto(out *KmsGrantList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KmsGrant, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KmsGrantList.
func (in *KmsGrantList) DeepCopy() *KmsGrantList {
	if in == nil {
		return nil
	}
	out := new(KmsGrantList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KmsGrantList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KmsGrantObservation) DeepCopyInto(out *KmsGrantObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KmsGrantObservation.
func (in *KmsGrantObservation) DeepCopy() *KmsGrantObservation {
	if in == nil {
		return nil
	}
	out := new(KmsGrantObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KmsGrantParameters) DeepCopyInto(out *KmsGrantParameters) {
	*out = *in
	if in.Constraints != nil {
		in, out := &in.Constraints, &out.Constraints
		*out = make([]ConstraintsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.GrantCreationTokens != nil {
		in, out := &in.GrantCreationTokens, &out.GrantCreationTokens
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Operations != nil {
		in, out := &in.Operations, &out.Operations
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.RetireOnDelete != nil {
		in, out := &in.RetireOnDelete, &out.RetireOnDelete
		*out = new(bool)
		**out = **in
	}
	if in.RetiringPrincipal != nil {
		in, out := &in.RetiringPrincipal, &out.RetiringPrincipal
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KmsGrantParameters.
func (in *KmsGrantParameters) DeepCopy() *KmsGrantParameters {
	if in == nil {
		return nil
	}
	out := new(KmsGrantParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KmsGrantSpec) DeepCopyInto(out *KmsGrantSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KmsGrantSpec.
func (in *KmsGrantSpec) DeepCopy() *KmsGrantSpec {
	if in == nil {
		return nil
	}
	out := new(KmsGrantSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KmsGrantStatus) DeepCopyInto(out *KmsGrantStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KmsGrantStatus.
func (in *KmsGrantStatus) DeepCopy() *KmsGrantStatus {
	if in == nil {
		return nil
	}
	out := new(KmsGrantStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KmsKey) DeepCopyInto(out *KmsKey) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KmsKey.
func (in *KmsKey) DeepCopy() *KmsKey {
	if in == nil {
		return nil
	}
	out := new(KmsKey)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KmsKey) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KmsKeyList) DeepCopyInto(out *KmsKeyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KmsKey, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KmsKeyList.
func (in *KmsKeyList) DeepCopy() *KmsKeyList {
	if in == nil {
		return nil
	}
	out := new(KmsKeyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KmsKeyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KmsKeyObservation) DeepCopyInto(out *KmsKeyObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KmsKeyObservation.
func (in *KmsKeyObservation) DeepCopy() *KmsKeyObservation {
	if in == nil {
		return nil
	}
	out := new(KmsKeyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KmsKeyParameters) DeepCopyInto(out *KmsKeyParameters) {
	*out = *in
	if in.BypassPolicyLockoutSafetyCheck != nil {
		in, out := &in.BypassPolicyLockoutSafetyCheck, &out.BypassPolicyLockoutSafetyCheck
		*out = new(bool)
		**out = **in
	}
	if in.CustomerMasterKeySpec != nil {
		in, out := &in.CustomerMasterKeySpec, &out.CustomerMasterKeySpec
		*out = new(string)
		**out = **in
	}
	if in.DeletionWindowInDays != nil {
		in, out := &in.DeletionWindowInDays, &out.DeletionWindowInDays
		*out = new(int64)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.EnableKeyRotation != nil {
		in, out := &in.EnableKeyRotation, &out.EnableKeyRotation
		*out = new(bool)
		**out = **in
	}
	if in.IsEnabled != nil {
		in, out := &in.IsEnabled, &out.IsEnabled
		*out = new(bool)
		**out = **in
	}
	if in.KeyUsage != nil {
		in, out := &in.KeyUsage, &out.KeyUsage
		*out = new(string)
		**out = **in
	}
	if in.Policy != nil {
		in, out := &in.Policy, &out.Policy
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
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KmsKeyParameters.
func (in *KmsKeyParameters) DeepCopy() *KmsKeyParameters {
	if in == nil {
		return nil
	}
	out := new(KmsKeyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KmsKeySpec) DeepCopyInto(out *KmsKeySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KmsKeySpec.
func (in *KmsKeySpec) DeepCopy() *KmsKeySpec {
	if in == nil {
		return nil
	}
	out := new(KmsKeySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KmsKeyStatus) DeepCopyInto(out *KmsKeyStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KmsKeyStatus.
func (in *KmsKeyStatus) DeepCopy() *KmsKeyStatus {
	if in == nil {
		return nil
	}
	out := new(KmsKeyStatus)
	in.DeepCopyInto(out)
	return out
}