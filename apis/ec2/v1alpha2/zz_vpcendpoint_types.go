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

type DNSEntryObservation struct {
	DNSName *string `json:"dnsName,omitempty" tf:"dns_name,omitempty"`

	HostedZoneID *string `json:"hostedZoneId,omitempty" tf:"hosted_zone_id,omitempty"`
}

type DNSEntryParameters struct {
}

type DNSOptionsObservation struct {
}

type DNSOptionsParameters struct {

	// +kubebuilder:validation:Optional
	DNSRecordIPType *string `json:"dnsRecordIpType,omitempty" tf:"dns_record_ip_type,omitempty"`
}

type VPCEndpointObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	CidrBlocks []*string `json:"cidrBlocks,omitempty" tf:"cidr_blocks,omitempty"`

	DNSEntry []DNSEntryObservation `json:"dnsEntry,omitempty" tf:"dns_entry,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	NetworkInterfaceIds []*string `json:"networkInterfaceIds,omitempty" tf:"network_interface_ids,omitempty"`

	OwnerID *string `json:"ownerId,omitempty" tf:"owner_id,omitempty"`

	PrefixListID *string `json:"prefixListId,omitempty" tf:"prefix_list_id,omitempty"`

	RequesterManaged *bool `json:"requesterManaged,omitempty" tf:"requester_managed,omitempty"`

	State *string `json:"state,omitempty" tf:"state,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type VPCEndpointParameters struct {

	// +kubebuilder:validation:Optional
	AutoAccept *bool `json:"autoAccept,omitempty" tf:"auto_accept,omitempty"`

	// +kubebuilder:validation:Optional
	DNSOptions []DNSOptionsParameters `json:"dnsOptions,omitempty" tf:"dns_options,omitempty"`

	// +kubebuilder:validation:Optional
	IPAddressType *string `json:"ipAddressType,omitempty" tf:"ip_address_type,omitempty"`

	// +kubebuilder:validation:Optional
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`

	// +kubebuilder:validation:Optional
	PrivateDNSEnabled *bool `json:"privateDnsEnabled,omitempty" tf:"private_dns_enabled,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +crossplane:generate:reference:type=RouteTable
	// +kubebuilder:validation:Optional
	RouteTableIds []*string `json:"routeTableIds,omitempty" tf:"route_table_ids,omitempty"`

	// References to RouteTable to populate routeTableIds.
	// +kubebuilder:validation:Optional
	RouteTableIdsRefs []v1.Reference `json:"routeTableIdsRefs,omitempty" tf:"-"`

	// Selector for a list of RouteTable to populate routeTableIds.
	// +kubebuilder:validation:Optional
	RouteTableIdsSelector *v1.Selector `json:"routeTableIdsSelector,omitempty" tf:"-"`

	// +crossplane:generate:reference:type=SecurityGroup
	// +kubebuilder:validation:Optional
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// References to SecurityGroup to populate securityGroupIds.
	// +kubebuilder:validation:Optional
	SecurityGroupIdsRefs []v1.Reference `json:"securityGroupIdsRefs,omitempty" tf:"-"`

	// Selector for a list of SecurityGroup to populate securityGroupIds.
	// +kubebuilder:validation:Optional
	SecurityGroupIdsSelector *v1.Selector `json:"securityGroupIdsSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	ServiceName *string `json:"serviceName" tf:"service_name,omitempty"`

	// +crossplane:generate:reference:type=Subnet
	// +kubebuilder:validation:Optional
	SubnetIds []*string `json:"subnetIds,omitempty" tf:"subnet_ids,omitempty"`

	// References to Subnet to populate subnetIds.
	// +kubebuilder:validation:Optional
	SubnetIdsRefs []v1.Reference `json:"subnetIdsRefs,omitempty" tf:"-"`

	// Selector for a list of Subnet to populate subnetIds.
	// +kubebuilder:validation:Optional
	SubnetIdsSelector *v1.Selector `json:"subnetIdsSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	VPCEndpointType *string `json:"vpcEndpointType,omitempty" tf:"vpc_endpoint_type,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-aws/apis/ec2/v1alpha2.VPC
	// +kubebuilder:validation:Optional
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// Reference to a VPC in ec2 to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDRef *v1.Reference `json:"vpcIdRef,omitempty" tf:"-"`

	// Selector for a VPC in ec2 to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDSelector *v1.Selector `json:"vpcIdSelector,omitempty" tf:"-"`
}

// VPCEndpointSpec defines the desired state of VPCEndpoint
type VPCEndpointSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VPCEndpointParameters `json:"forProvider"`
}

// VPCEndpointStatus defines the observed state of VPCEndpoint.
type VPCEndpointStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VPCEndpointObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VPCEndpoint is the Schema for the VPCEndpoints API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type VPCEndpoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VPCEndpointSpec   `json:"spec"`
	Status            VPCEndpointStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VPCEndpointList contains a list of VPCEndpoints
type VPCEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VPCEndpoint `json:"items"`
}

// Repository type metadata.
var (
	VPCEndpoint_Kind             = "VPCEndpoint"
	VPCEndpoint_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VPCEndpoint_Kind}.String()
	VPCEndpoint_KindAPIVersion   = VPCEndpoint_Kind + "." + CRDGroupVersion.String()
	VPCEndpoint_GroupVersionKind = CRDGroupVersion.WithKind(VPCEndpoint_Kind)
)

func init() {
	SchemeBuilder.Register(&VPCEndpoint{}, &VPCEndpointList{})
}
