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

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type EbsDefaultKmsKeyObservation struct {
}

type EbsDefaultKmsKeyParameters struct {
	KeyArn string `json:"keyArn" tf:"key_arn"`
}

// EbsDefaultKmsKeySpec defines the desired state of EbsDefaultKmsKey
type EbsDefaultKmsKeySpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       EbsDefaultKmsKeyParameters `json:"forProvider"`
}

// EbsDefaultKmsKeyStatus defines the observed state of EbsDefaultKmsKey.
type EbsDefaultKmsKeyStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          EbsDefaultKmsKeyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// EbsDefaultKmsKey is the Schema for the EbsDefaultKmsKeys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type EbsDefaultKmsKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EbsDefaultKmsKeySpec   `json:"spec"`
	Status            EbsDefaultKmsKeyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EbsDefaultKmsKeyList contains a list of EbsDefaultKmsKeys
type EbsDefaultKmsKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EbsDefaultKmsKey `json:"items"`
}

// Repository type metadata.
var (
	EbsDefaultKmsKeyKind             = "EbsDefaultKmsKey"
	EbsDefaultKmsKeyGroupKind        = schema.GroupKind{Group: Group, Kind: EbsDefaultKmsKeyKind}.String()
	EbsDefaultKmsKeyKindAPIVersion   = EbsDefaultKmsKeyKind + "." + GroupVersion.String()
	EbsDefaultKmsKeyGroupVersionKind = GroupVersion.WithKind(EbsDefaultKmsKeyKind)
)

func init() {
	SchemeBuilder.Register(&EbsDefaultKmsKey{}, &EbsDefaultKmsKeyList{})
}