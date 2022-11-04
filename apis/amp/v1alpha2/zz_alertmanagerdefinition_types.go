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

type AlertManagerDefinitionObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type AlertManagerDefinitionParameters struct {

	// +kubebuilder:validation:Required
	Definition *string `json:"definition" tf:"definition,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Required
	WorkspaceID *string `json:"workspaceId" tf:"workspace_id,omitempty"`
}

// AlertManagerDefinitionSpec defines the desired state of AlertManagerDefinition
type AlertManagerDefinitionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AlertManagerDefinitionParameters `json:"forProvider"`
}

// AlertManagerDefinitionStatus defines the observed state of AlertManagerDefinition.
type AlertManagerDefinitionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AlertManagerDefinitionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AlertManagerDefinition is the Schema for the AlertManagerDefinitions API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type AlertManagerDefinition struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AlertManagerDefinitionSpec   `json:"spec"`
	Status            AlertManagerDefinitionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AlertManagerDefinitionList contains a list of AlertManagerDefinitions
type AlertManagerDefinitionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AlertManagerDefinition `json:"items"`
}

// Repository type metadata.
var (
	AlertManagerDefinition_Kind             = "AlertManagerDefinition"
	AlertManagerDefinition_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AlertManagerDefinition_Kind}.String()
	AlertManagerDefinition_KindAPIVersion   = AlertManagerDefinition_Kind + "." + CRDGroupVersion.String()
	AlertManagerDefinition_GroupVersionKind = CRDGroupVersion.WithKind(AlertManagerDefinition_Kind)
)

func init() {
	SchemeBuilder.Register(&AlertManagerDefinition{}, &AlertManagerDefinitionList{})
}