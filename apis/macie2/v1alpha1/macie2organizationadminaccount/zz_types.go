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

// Code generated by terrajet. DO NOT EDIT.

// +kubebuilder:object:generate=true
// +groupName=macie2.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/macie2/v1alpha1"
)

type Macie2OrganizationAdminAccountObservation struct {
}

type Macie2OrganizationAdminAccountParameters struct {
	AdminAccountId string `json:"adminAccountId" tf:"admin_account_id"`
}

// Macie2OrganizationAdminAccountSpec defines the desired state of Macie2OrganizationAdminAccount
type Macie2OrganizationAdminAccountSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       Macie2OrganizationAdminAccountParameters `json:"forProvider"`
}

// Macie2OrganizationAdminAccountStatus defines the observed state of Macie2OrganizationAdminAccount.
type Macie2OrganizationAdminAccountStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          Macie2OrganizationAdminAccountObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Macie2OrganizationAdminAccount is the Schema for the Macie2OrganizationAdminAccounts API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Macie2OrganizationAdminAccount struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Macie2OrganizationAdminAccountSpec   `json:"spec"`
	Status            Macie2OrganizationAdminAccountStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Macie2OrganizationAdminAccountList contains a list of Macie2OrganizationAdminAccounts
type Macie2OrganizationAdminAccountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Macie2OrganizationAdminAccount `json:"items"`
}

// Repository type metadata.
var (
	Macie2OrganizationAdminAccountKind             = "Macie2OrganizationAdminAccount"
	Macie2OrganizationAdminAccountGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: Macie2OrganizationAdminAccountKind}.String()
	Macie2OrganizationAdminAccountKindAPIVersion   = Macie2OrganizationAdminAccountKind + "." + v1alpha1.GroupVersion.String()
	Macie2OrganizationAdminAccountGroupVersionKind = v1alpha1.GroupVersion.WithKind(Macie2OrganizationAdminAccountKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&Macie2OrganizationAdminAccount{}, &Macie2OrganizationAdminAccountList{})
}