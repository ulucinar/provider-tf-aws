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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type VPCObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	DefaultNetworkACLID *string `json:"defaultNetworkAclId,omitempty" tf:"default_network_acl_id,omitempty"`

	DefaultRouteTableID *string `json:"defaultRouteTableId,omitempty" tf:"default_route_table_id,omitempty"`

	DefaultSecurityGroupID *string `json:"defaultSecurityGroupId,omitempty" tf:"default_security_group_id,omitempty"`

	DhcpOptionsID *string `json:"dhcpOptionsId,omitempty" tf:"dhcp_options_id,omitempty"`

	IPv6AssociationID *string `json:"ipv6AssociationId,omitempty" tf:"ipv6_association_id,omitempty"`

	IPv6CidrBlock *string `json:"ipv6CidrBlock,omitempty" tf:"ipv6_cidr_block,omitempty"`

	MainRouteTableID *string `json:"mainRouteTableId,omitempty" tf:"main_route_table_id,omitempty"`

	OwnerID *string `json:"ownerId,omitempty" tf:"owner_id,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type VPCParameters struct {

	// +kubebuilder:validation:Optional
	AssignGeneratedIPv6CidrBlock *bool `json:"assignGeneratedIpv6CidrBlock,omitempty" tf:"assign_generated_ipv6_cidr_block,omitempty"`

	// +kubebuilder:validation:Required
	CidrBlock *string `json:"cidrBlock" tf:"cidr_block,omitempty"`

	// +kubebuilder:validation:Optional
	EnableClassiclink *bool `json:"enableClassiclink,omitempty" tf:"enable_classiclink,omitempty"`

	// +kubebuilder:validation:Optional
	EnableClassiclinkDNSSupport *bool `json:"enableClassiclinkDnsSupport,omitempty" tf:"enable_classiclink_dns_support,omitempty"`

	// +kubebuilder:validation:Optional
	EnableDNSHostnames *bool `json:"enableDnsHostnames,omitempty" tf:"enable_dns_hostnames,omitempty"`

	// +kubebuilder:validation:Optional
	EnableDNSSupport *bool `json:"enableDnsSupport,omitempty" tf:"enable_dns_support,omitempty"`

	// +kubebuilder:validation:Optional
	InstanceTenancy *string `json:"instanceTenancy,omitempty" tf:"instance_tenancy,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// VPCSpec defines the desired state of VPC
type VPCSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VPCParameters `json:"forProvider"`
}

// VPCStatus defines the observed state of VPC.
type VPCStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VPCObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VPC is the Schema for the VPCs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type VPC struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VPCSpec   `json:"spec"`
	Status            VPCStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VPCList contains a list of VPCs
type VPCList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VPC `json:"items"`
}

// Repository type metadata.
var (
	VPCKind             = "VPC"
	VPCGroupKind        = schema.GroupKind{Group: Group, Kind: VPCKind}.String()
	VPCKindAPIVersion   = VPCKind + "." + GroupVersion.String()
	VPCGroupVersionKind = GroupVersion.WithKind(VPCKind)
)

func init() {
	SchemeBuilder.Register(&VPC{}, &VPCList{})
}