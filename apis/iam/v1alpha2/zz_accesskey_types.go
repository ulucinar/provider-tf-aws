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

// Code generated by upjet. DO NOT EDIT.

package v1alpha2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AccessKeyObservation struct {
	CreateDate *string `json:"createDate,omitempty" tf:"create_date,omitempty"`

	EncryptedSecret *string `json:"encryptedSecret,omitempty" tf:"encrypted_secret,omitempty"`

	EncryptedSesSMTPPasswordV4 *string `json:"encryptedSesSmtpPasswordV4,omitempty" tf:"encrypted_ses_smtp_password_v4,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	KeyFingerprint *string `json:"keyFingerprint,omitempty" tf:"key_fingerprint,omitempty"`
}

type AccessKeyParameters struct {

	// +kubebuilder:validation:Optional
	PgpKey *string `json:"pgpKey,omitempty" tf:"pgp_key,omitempty"`

	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// +crossplane:generate:reference:type=User
	// +kubebuilder:validation:Optional
	User *string `json:"user,omitempty" tf:"user,omitempty"`

	// Reference to a User to populate user.
	// +kubebuilder:validation:Optional
	UserRef *v1.Reference `json:"userRef,omitempty" tf:"-"`

	// Selector for a User to populate user.
	// +kubebuilder:validation:Optional
	UserSelector *v1.Selector `json:"userSelector,omitempty" tf:"-"`
}

// AccessKeySpec defines the desired state of AccessKey
type AccessKeySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AccessKeyParameters `json:"forProvider"`
}

// AccessKeyStatus defines the observed state of AccessKey.
type AccessKeyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AccessKeyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AccessKey is the Schema for the AccessKeys API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type AccessKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AccessKeySpec   `json:"spec"`
	Status            AccessKeyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AccessKeyList contains a list of AccessKeys
type AccessKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AccessKey `json:"items"`
}

// Repository type metadata.
var (
	AccessKey_Kind             = "AccessKey"
	AccessKey_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AccessKey_Kind}.String()
	AccessKey_KindAPIVersion   = AccessKey_Kind + "." + CRDGroupVersion.String()
	AccessKey_GroupVersionKind = CRDGroupVersion.WithKind(AccessKey_Kind)
)

func init() {
	SchemeBuilder.Register(&AccessKey{}, &AccessKeyList{})
}
