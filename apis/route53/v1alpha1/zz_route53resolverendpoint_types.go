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

type IpAddressObservation struct {
	IpId string `json:"ipId" tf:"ip_id"`
}

type IpAddressParameters struct {
	Ip *string `json:"ip,omitempty" tf:"ip"`

	SubnetId string `json:"subnetId" tf:"subnet_id"`
}

type Route53ResolverEndpointObservation struct {
	Arn string `json:"arn" tf:"arn"`

	HostVpcId string `json:"hostVpcId" tf:"host_vpc_id"`
}

type Route53ResolverEndpointParameters struct {
	Direction string `json:"direction" tf:"direction"`

	IpAddress []IpAddressParameters `json:"ipAddress" tf:"ip_address"`

	Name *string `json:"name,omitempty" tf:"name"`

	SecurityGroupIds []string `json:"securityGroupIds" tf:"security_group_ids"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

// Route53ResolverEndpointSpec defines the desired state of Route53ResolverEndpoint
type Route53ResolverEndpointSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       Route53ResolverEndpointParameters `json:"forProvider"`
}

// Route53ResolverEndpointStatus defines the observed state of Route53ResolverEndpoint.
type Route53ResolverEndpointStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          Route53ResolverEndpointObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Route53ResolverEndpoint is the Schema for the Route53ResolverEndpoints API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type Route53ResolverEndpoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Route53ResolverEndpointSpec   `json:"spec"`
	Status            Route53ResolverEndpointStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Route53ResolverEndpointList contains a list of Route53ResolverEndpoints
type Route53ResolverEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Route53ResolverEndpoint `json:"items"`
}

// Repository type metadata.
var (
	Route53ResolverEndpointKind             = "Route53ResolverEndpoint"
	Route53ResolverEndpointGroupKind        = schema.GroupKind{Group: Group, Kind: Route53ResolverEndpointKind}.String()
	Route53ResolverEndpointKindAPIVersion   = Route53ResolverEndpointKind + "." + GroupVersion.String()
	Route53ResolverEndpointGroupVersionKind = GroupVersion.WithKind(Route53ResolverEndpointKind)
)

func init() {
	SchemeBuilder.Register(&Route53ResolverEndpoint{}, &Route53ResolverEndpointList{})
}