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

type DNSEntryObservation struct {
	DNSName string `json:"dnsName,omitempty" tf:"dns_name"`

	HostedZoneID string `json:"hostedZoneId,omitempty" tf:"hosted_zone_id"`
}

type DNSEntryParameters struct {
}

type VpcEndpointObservation struct {
	Arn string `json:"arn,omitempty" tf:"arn"`

	CidrBlocks []string `json:"cidrBlocks,omitempty" tf:"cidr_blocks"`

	DNSEntry []DNSEntryObservation `json:"dnsEntry,omitempty" tf:"dns_entry"`

	NetworkInterfaceIds []string `json:"networkInterfaceIds,omitempty" tf:"network_interface_ids"`

	OwnerID string `json:"ownerId,omitempty" tf:"owner_id"`

	PrefixListID string `json:"prefixListId,omitempty" tf:"prefix_list_id"`

	RequesterManaged bool `json:"requesterManaged,omitempty" tf:"requester_managed"`

	State string `json:"state,omitempty" tf:"state"`
}

type VpcEndpointParameters struct {

	// +kubebuilder:validation:Optional
	AutoAccept *bool `json:"autoAccept,omitempty" tf:"auto_accept"`

	// +kubebuilder:validation:Optional
	Policy *string `json:"policy,omitempty" tf:"policy"`

	// +kubebuilder:validation:Optional
	PrivateDNSEnabled *bool `json:"privateDnsEnabled,omitempty" tf:"private_dns_enabled"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	RouteTableIds []string `json:"routeTableIds,omitempty" tf:"route_table_ids"`

	// +kubebuilder:validation:Optional
	SecurityGroupIds []string `json:"securityGroupIds,omitempty" tf:"security_group_ids"`

	// +kubebuilder:validation:Required
	ServiceName string `json:"serviceName" tf:"service_name"`

	// +kubebuilder:validation:Optional
	SubnetIds []string `json:"subnetIds,omitempty" tf:"subnet_ids"`

	// +kubebuilder:validation:Optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	// +kubebuilder:validation:Optional
	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`

	// +kubebuilder:validation:Optional
	VpcEndpointType *string `json:"vpcEndpointType,omitempty" tf:"vpc_endpoint_type"`

	// +kubebuilder:validation:Required
	VpcID string `json:"vpcId" tf:"vpc_id"`
}

// VpcEndpointSpec defines the desired state of VpcEndpoint
type VpcEndpointSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       VpcEndpointParameters `json:"forProvider"`
}

// VpcEndpointStatus defines the observed state of VpcEndpoint.
type VpcEndpointStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          VpcEndpointObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VpcEndpoint is the Schema for the VpcEndpoints API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type VpcEndpoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VpcEndpointSpec   `json:"spec"`
	Status            VpcEndpointStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VpcEndpointList contains a list of VpcEndpoints
type VpcEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VpcEndpoint `json:"items"`
}

// Repository type metadata.
var (
	VpcEndpointKind             = "VpcEndpoint"
	VpcEndpointGroupKind        = schema.GroupKind{Group: Group, Kind: VpcEndpointKind}.String()
	VpcEndpointKindAPIVersion   = VpcEndpointKind + "." + GroupVersion.String()
	VpcEndpointGroupVersionKind = GroupVersion.WithKind(VpcEndpointKind)
)

func init() {
	SchemeBuilder.Register(&VpcEndpoint{}, &VpcEndpointList{})
}
