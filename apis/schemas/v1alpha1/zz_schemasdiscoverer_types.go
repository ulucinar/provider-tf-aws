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

type SchemasDiscovererObservation struct {
	Arn string `json:"arn" tf:"arn"`
}

type SchemasDiscovererParameters struct {
	Description *string `json:"description,omitempty" tf:"description"`

	SourceArn string `json:"sourceArn" tf:"source_arn"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

// SchemasDiscovererSpec defines the desired state of SchemasDiscoverer
type SchemasDiscovererSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       SchemasDiscovererParameters `json:"forProvider"`
}

// SchemasDiscovererStatus defines the observed state of SchemasDiscoverer.
type SchemasDiscovererStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          SchemasDiscovererObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SchemasDiscoverer is the Schema for the SchemasDiscoverers API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type SchemasDiscoverer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SchemasDiscovererSpec   `json:"spec"`
	Status            SchemasDiscovererStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SchemasDiscovererList contains a list of SchemasDiscoverers
type SchemasDiscovererList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SchemasDiscoverer `json:"items"`
}

// Repository type metadata.
var (
	SchemasDiscovererKind             = "SchemasDiscoverer"
	SchemasDiscovererGroupKind        = schema.GroupKind{Group: Group, Kind: SchemasDiscovererKind}.String()
	SchemasDiscovererKindAPIVersion   = SchemasDiscovererKind + "." + GroupVersion.String()
	SchemasDiscovererGroupVersionKind = GroupVersion.WithKind(SchemasDiscovererKind)
)

func init() {
	SchemeBuilder.Register(&SchemasDiscoverer{}, &SchemasDiscovererList{})
}