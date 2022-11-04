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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type FirewallDomainListObservation struct {

	// The ARN (Amazon Resource Name) of the domain list.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The ID of the domain list.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type FirewallDomainListParameters struct {

	// A array of domains for the firewall domain list.
	// +kubebuilder:validation:Optional
	Domains []*string `json:"domains,omitempty" tf:"domains,omitempty"`

	// A name that lets you identify the domain list, to manage and use it.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// A map of tags to assign to the resource. f configured with a provider default_tags configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// FirewallDomainListSpec defines the desired state of FirewallDomainList
type FirewallDomainListSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FirewallDomainListParameters `json:"forProvider"`
}

// FirewallDomainListStatus defines the observed state of FirewallDomainList.
type FirewallDomainListStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FirewallDomainListObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FirewallDomainList is the Schema for the FirewallDomainLists API. Provides a Route 53 Resolver DNS Firewall domain list resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type FirewallDomainList struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FirewallDomainListSpec   `json:"spec"`
	Status            FirewallDomainListStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FirewallDomainListList contains a list of FirewallDomainLists
type FirewallDomainListList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FirewallDomainList `json:"items"`
}

// Repository type metadata.
var (
	FirewallDomainList_Kind             = "FirewallDomainList"
	FirewallDomainList_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FirewallDomainList_Kind}.String()
	FirewallDomainList_KindAPIVersion   = FirewallDomainList_Kind + "." + CRDGroupVersion.String()
	FirewallDomainList_GroupVersionKind = CRDGroupVersion.WithKind(FirewallDomainList_Kind)
)

func init() {
	SchemeBuilder.Register(&FirewallDomainList{}, &FirewallDomainListList{})
}
