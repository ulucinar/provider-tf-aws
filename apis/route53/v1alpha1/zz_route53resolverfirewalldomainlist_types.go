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

type Route53ResolverFirewallDomainListObservation struct {
	Arn string `json:"arn,omitempty" tf:"arn"`
}

type Route53ResolverFirewallDomainListParameters struct {

	// +kubebuilder:validation:Optional
	Domains []string `json:"domains,omitempty" tf:"domains"`

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	// +kubebuilder:validation:Optional
	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

// Route53ResolverFirewallDomainListSpec defines the desired state of Route53ResolverFirewallDomainList
type Route53ResolverFirewallDomainListSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       Route53ResolverFirewallDomainListParameters `json:"forProvider"`
}

// Route53ResolverFirewallDomainListStatus defines the observed state of Route53ResolverFirewallDomainList.
type Route53ResolverFirewallDomainListStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          Route53ResolverFirewallDomainListObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Route53ResolverFirewallDomainList is the Schema for the Route53ResolverFirewallDomainLists API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type Route53ResolverFirewallDomainList struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Route53ResolverFirewallDomainListSpec   `json:"spec"`
	Status            Route53ResolverFirewallDomainListStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Route53ResolverFirewallDomainListList contains a list of Route53ResolverFirewallDomainLists
type Route53ResolverFirewallDomainListList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Route53ResolverFirewallDomainList `json:"items"`
}

// Repository type metadata.
var (
	Route53ResolverFirewallDomainListKind             = "Route53ResolverFirewallDomainList"
	Route53ResolverFirewallDomainListGroupKind        = schema.GroupKind{Group: Group, Kind: Route53ResolverFirewallDomainListKind}.String()
	Route53ResolverFirewallDomainListKindAPIVersion   = Route53ResolverFirewallDomainListKind + "." + GroupVersion.String()
	Route53ResolverFirewallDomainListGroupVersionKind = GroupVersion.WithKind(Route53ResolverFirewallDomainListKind)
)

func init() {
	SchemeBuilder.Register(&Route53ResolverFirewallDomainList{}, &Route53ResolverFirewallDomainListList{})
}
