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

type MskConfigurationObservation struct {
	Arn string `json:"arn" tf:"arn"`

	LatestRevision int64 `json:"latestRevision" tf:"latest_revision"`
}

type MskConfigurationParameters struct {
	Description *string `json:"description,omitempty" tf:"description"`

	KafkaVersions []string `json:"kafkaVersions,omitempty" tf:"kafka_versions"`

	Name string `json:"name" tf:"name"`

	ServerProperties string `json:"serverProperties" tf:"server_properties"`
}

// MskConfigurationSpec defines the desired state of MskConfiguration
type MskConfigurationSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       MskConfigurationParameters `json:"forProvider"`
}

// MskConfigurationStatus defines the observed state of MskConfiguration.
type MskConfigurationStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          MskConfigurationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// MskConfiguration is the Schema for the MskConfigurations API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type MskConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MskConfigurationSpec   `json:"spec"`
	Status            MskConfigurationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MskConfigurationList contains a list of MskConfigurations
type MskConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MskConfiguration `json:"items"`
}

// Repository type metadata.
var (
	MskConfigurationKind             = "MskConfiguration"
	MskConfigurationGroupKind        = schema.GroupKind{Group: Group, Kind: MskConfigurationKind}.String()
	MskConfigurationKindAPIVersion   = MskConfigurationKind + "." + GroupVersion.String()
	MskConfigurationGroupVersionKind = GroupVersion.WithKind(MskConfigurationKind)
)

func init() {
	SchemeBuilder.Register(&MskConfiguration{}, &MskConfigurationList{})
}