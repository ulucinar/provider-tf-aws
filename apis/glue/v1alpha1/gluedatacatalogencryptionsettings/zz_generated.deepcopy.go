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
func (in *ConnectionPasswordEncryptionObservation) DeepCopyInto(out *ConnectionPasswordEncryptionObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConnectionPasswordEncryptionObservation.
func (in *ConnectionPasswordEncryptionObservation) DeepCopy() *ConnectionPasswordEncryptionObservation {
	if in == nil {
		return nil
	}
	out := new(ConnectionPasswordEncryptionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConnectionPasswordEncryptionParameters) DeepCopyInto(out *ConnectionPasswordEncryptionParameters) {
	*out = *in
	if in.AwsKmsKeyId != nil {
		in, out := &in.AwsKmsKeyId, &out.AwsKmsKeyId
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConnectionPasswordEncryptionParameters.
func (in *ConnectionPasswordEncryptionParameters) DeepCopy() *ConnectionPasswordEncryptionParameters {
	if in == nil {
		return nil
	}
	out := new(ConnectionPasswordEncryptionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataCatalogEncryptionSettingsObservation) DeepCopyInto(out *DataCatalogEncryptionSettingsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataCatalogEncryptionSettingsObservation.
func (in *DataCatalogEncryptionSettingsObservation) DeepCopy() *DataCatalogEncryptionSettingsObservation {
	if in == nil {
		return nil
	}
	out := new(DataCatalogEncryptionSettingsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataCatalogEncryptionSettingsParameters) DeepCopyInto(out *DataCatalogEncryptionSettingsParameters) {
	*out = *in
	if in.ConnectionPasswordEncryption != nil {
		in, out := &in.ConnectionPasswordEncryption, &out.ConnectionPasswordEncryption
		*out = make([]ConnectionPasswordEncryptionParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.EncryptionAtRest != nil {
		in, out := &in.EncryptionAtRest, &out.EncryptionAtRest
		*out = make([]EncryptionAtRestParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataCatalogEncryptionSettingsParameters.
func (in *DataCatalogEncryptionSettingsParameters) DeepCopy() *DataCatalogEncryptionSettingsParameters {
	if in == nil {
		return nil
	}
	out := new(DataCatalogEncryptionSettingsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EncryptionAtRestObservation) DeepCopyInto(out *EncryptionAtRestObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EncryptionAtRestObservation.
func (in *EncryptionAtRestObservation) DeepCopy() *EncryptionAtRestObservation {
	if in == nil {
		return nil
	}
	out := new(EncryptionAtRestObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EncryptionAtRestParameters) DeepCopyInto(out *EncryptionAtRestParameters) {
	*out = *in
	if in.SseAwsKmsKeyId != nil {
		in, out := &in.SseAwsKmsKeyId, &out.SseAwsKmsKeyId
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EncryptionAtRestParameters.
func (in *EncryptionAtRestParameters) DeepCopy() *EncryptionAtRestParameters {
	if in == nil {
		return nil
	}
	out := new(EncryptionAtRestParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlueDataCatalogEncryptionSettings) DeepCopyInto(out *GlueDataCatalogEncryptionSettings) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlueDataCatalogEncryptionSettings.
func (in *GlueDataCatalogEncryptionSettings) DeepCopy() *GlueDataCatalogEncryptionSettings {
	if in == nil {
		return nil
	}
	out := new(GlueDataCatalogEncryptionSettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GlueDataCatalogEncryptionSettings) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlueDataCatalogEncryptionSettingsList) DeepCopyInto(out *GlueDataCatalogEncryptionSettingsList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GlueDataCatalogEncryptionSettings, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlueDataCatalogEncryptionSettingsList.
func (in *GlueDataCatalogEncryptionSettingsList) DeepCopy() *GlueDataCatalogEncryptionSettingsList {
	if in == nil {
		return nil
	}
	out := new(GlueDataCatalogEncryptionSettingsList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GlueDataCatalogEncryptionSettingsList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlueDataCatalogEncryptionSettingsObservation) DeepCopyInto(out *GlueDataCatalogEncryptionSettingsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlueDataCatalogEncryptionSettingsObservation.
func (in *GlueDataCatalogEncryptionSettingsObservation) DeepCopy() *GlueDataCatalogEncryptionSettingsObservation {
	if in == nil {
		return nil
	}
	out := new(GlueDataCatalogEncryptionSettingsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlueDataCatalogEncryptionSettingsParameters) DeepCopyInto(out *GlueDataCatalogEncryptionSettingsParameters) {
	*out = *in
	if in.CatalogId != nil {
		in, out := &in.CatalogId, &out.CatalogId
		*out = new(string)
		**out = **in
	}
	if in.DataCatalogEncryptionSettings != nil {
		in, out := &in.DataCatalogEncryptionSettings, &out.DataCatalogEncryptionSettings
		*out = make([]DataCatalogEncryptionSettingsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlueDataCatalogEncryptionSettingsParameters.
func (in *GlueDataCatalogEncryptionSettingsParameters) DeepCopy() *GlueDataCatalogEncryptionSettingsParameters {
	if in == nil {
		return nil
	}
	out := new(GlueDataCatalogEncryptionSettingsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlueDataCatalogEncryptionSettingsSpec) DeepCopyInto(out *GlueDataCatalogEncryptionSettingsSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlueDataCatalogEncryptionSettingsSpec.
func (in *GlueDataCatalogEncryptionSettingsSpec) DeepCopy() *GlueDataCatalogEncryptionSettingsSpec {
	if in == nil {
		return nil
	}
	out := new(GlueDataCatalogEncryptionSettingsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlueDataCatalogEncryptionSettingsStatus) DeepCopyInto(out *GlueDataCatalogEncryptionSettingsStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlueDataCatalogEncryptionSettingsStatus.
func (in *GlueDataCatalogEncryptionSettingsStatus) DeepCopy() *GlueDataCatalogEncryptionSettingsStatus {
	if in == nil {
		return nil
	}
	out := new(GlueDataCatalogEncryptionSettingsStatus)
	in.DeepCopyInto(out)
	return out
}