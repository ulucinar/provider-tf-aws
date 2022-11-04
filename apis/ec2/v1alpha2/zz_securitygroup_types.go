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

type EgressObservation struct {
}

type EgressParameters struct {

	// List of CIDR blocks.
	// +kubebuilder:validation:Optional
	CidrBlocks []*string `json:"cidrBlocks,omitempty" tf:"cidr_blocks"`

	// Security group description. Cannot be "". NOTE: This field maps to the AWS GroupDescription attribute, for which there is no Update API. If you'd like to classify your security groups in a way that can be updated, use tags.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description"`

	// Start port (or ICMP type number if protocol is icmp or icmpv6).
	// +kubebuilder:validation:Optional
	FromPort *float64 `json:"fromPort,omitempty" tf:"from_port"`

	// List of IPv6 CIDR blocks.
	// +kubebuilder:validation:Optional
	IPv6CidrBlocks []*string `json:"ipv6CidrBlocks,omitempty" tf:"ipv6_cidr_blocks"`

	// List of Prefix List IDs.
	// +kubebuilder:validation:Optional
	PrefixListIds []*string `json:"prefixListIds,omitempty" tf:"prefix_list_ids"`

	// Protocol. If you select a protocol of -1 (semantically equivalent to all, which is not a valid value here), you must specify a from_port and to_port equal to 0.  The supported values are defined in the IpProtocol argument on the IpPermission API reference.12.
	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol"`

	// References to SecurityGroup to populate securityGroups.
	// +kubebuilder:validation:Optional
	SecurityGroupRefs []v1.Reference `json:"securityGroupRefs,omitempty" tf:"-"`

	// Selector for a list of SecurityGroup to populate securityGroups.
	// +kubebuilder:validation:Optional
	SecurityGroupSelector *v1.Selector `json:"securityGroupSelector,omitempty" tf:"-"`

	// List of security group Group Names if using EC2-Classic, or Group IDs if using a VPC.
	// +crossplane:generate:reference:type=SecurityGroup
	// +crossplane:generate:reference:refFieldName=SecurityGroupRefs
	// +crossplane:generate:reference:selectorFieldName=SecurityGroupSelector
	// +kubebuilder:validation:Optional
	SecurityGroups []*string `json:"securityGroups,omitempty" tf:"security_groups"`

	// Whether the security group itself will be added as a source to this ingress rule.
	// +kubebuilder:validation:Optional
	Self *bool `json:"self,omitempty" tf:"self"`

	// End range port (or ICMP code if protocol is icmp).
	// +kubebuilder:validation:Optional
	ToPort *float64 `json:"toPort,omitempty" tf:"to_port"`
}

type IngressObservation struct {
}

type IngressParameters struct {

	// List of CIDR blocks.
	// +kubebuilder:validation:Optional
	CidrBlocks []*string `json:"cidrBlocks,omitempty" tf:"cidr_blocks"`

	// Security group description. Cannot be "". NOTE: This field maps to the AWS GroupDescription attribute, for which there is no Update API. If you'd like to classify your security groups in a way that can be updated, use tags.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description"`

	// Start port (or ICMP type number if protocol is icmp or icmpv6).
	// +kubebuilder:validation:Optional
	FromPort *float64 `json:"fromPort,omitempty" tf:"from_port"`

	// List of IPv6 CIDR blocks.
	// +kubebuilder:validation:Optional
	IPv6CidrBlocks []*string `json:"ipv6CidrBlocks,omitempty" tf:"ipv6_cidr_blocks"`

	// List of Prefix List IDs.
	// +kubebuilder:validation:Optional
	PrefixListIds []*string `json:"prefixListIds,omitempty" tf:"prefix_list_ids"`

	// Protocol. If you select a protocol of -1 (semantically equivalent to all, which is not a valid value here), you must specify a from_port and to_port equal to 0.  The supported values are defined in the IpProtocol argument on the IpPermission API reference.12.
	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol"`

	// References to SecurityGroup to populate securityGroups.
	// +kubebuilder:validation:Optional
	SecurityGroupRefs []v1.Reference `json:"securityGroupRefs,omitempty" tf:"-"`

	// Selector for a list of SecurityGroup to populate securityGroups.
	// +kubebuilder:validation:Optional
	SecurityGroupSelector *v1.Selector `json:"securityGroupSelector,omitempty" tf:"-"`

	// List of security group Group Names if using EC2-Classic, or Group IDs if using a VPC.
	// +crossplane:generate:reference:type=SecurityGroup
	// +crossplane:generate:reference:refFieldName=SecurityGroupRefs
	// +crossplane:generate:reference:selectorFieldName=SecurityGroupSelector
	// +kubebuilder:validation:Optional
	SecurityGroups []*string `json:"securityGroups,omitempty" tf:"security_groups"`

	// Whether the security group itself will be added as a source to this ingress rule.
	// +kubebuilder:validation:Optional
	Self *bool `json:"self,omitempty" tf:"self"`

	// End range port (or ICMP code if protocol is icmp).
	// +kubebuilder:validation:Optional
	ToPort *float64 `json:"toPort,omitempty" tf:"to_port"`
}

type SecurityGroupObservation struct {

	// ARN of the security group.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Configuration block for egress rules. Can be specified multiple times for each egress rule. Each egress block supports fields documented below. This argument is processed in attribute-as-blocks mode.
	// +kubebuilder:validation:Required
	Egress []EgressObservation `json:"egress,omitempty" tf:"egress,omitempty"`

	// ID of the security group.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Configuration block for ingress rules. Can be specified multiple times for each ingress rule. Each ingress block supports fields documented below. This argument is processed in attribute-as-blocks mode.
	// +kubebuilder:validation:Required
	Ingress []IngressObservation `json:"ingress,omitempty" tf:"ingress,omitempty"`

	// Owner ID.
	OwnerID *string `json:"ownerId,omitempty" tf:"owner_id,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type SecurityGroupParameters struct {

	// Security group description. Cannot be "". NOTE: This field maps to the AWS GroupDescription attribute, for which there is no Update API. If you'd like to classify your security groups in a way that can be updated, use tags.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Configuration block for egress rules. Can be specified multiple times for each egress rule. Each egress block supports fields documented below. This argument is processed in attribute-as-blocks mode.
	// +kubebuilder:validation:Required
	Egress []EgressParameters `json:"egress" tf:"egress,omitempty"`

	// Configuration block for ingress rules. Can be specified multiple times for each ingress rule. Each ingress block supports fields documented below. This argument is processed in attribute-as-blocks mode.
	// +kubebuilder:validation:Required
	Ingress []IngressParameters `json:"ingress" tf:"ingress,omitempty"`

	// Name of the security group.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Creates a unique name beginning with the specified prefix. Conflicts with name.
	// +kubebuilder:validation:Optional
	NamePrefix *string `json:"namePrefix,omitempty" tf:"name_prefix,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// This is normally not needed, however certain AWS services such as Elastic Map Reduce may automatically add required rules to security groups used with the service, and those rules may contain a cyclic dependency that prevent the security groups from being destroyed without removing the dependency first. Default false.
	// +kubebuilder:validation:Optional
	RevokeRulesOnDelete *bool `json:"revokeRulesOnDelete,omitempty" tf:"revoke_rules_on_delete,omitempty"`

	// Map of tags to assign to the resource. If configured with a provider default_tags configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// VPC ID.
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

// SecurityGroupSpec defines the desired state of SecurityGroup
type SecurityGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SecurityGroupParameters `json:"forProvider"`
}

// SecurityGroupStatus defines the observed state of SecurityGroup.
type SecurityGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SecurityGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SecurityGroup is the Schema for the SecurityGroups API. Provides a security group resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type SecurityGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SecurityGroupSpec   `json:"spec"`
	Status            SecurityGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecurityGroupList contains a list of SecurityGroups
type SecurityGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecurityGroup `json:"items"`
}

// Repository type metadata.
var (
	SecurityGroup_Kind             = "SecurityGroup"
	SecurityGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SecurityGroup_Kind}.String()
	SecurityGroup_KindAPIVersion   = SecurityGroup_Kind + "." + CRDGroupVersion.String()
	SecurityGroup_GroupVersionKind = CRDGroupVersion.WithKind(SecurityGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&SecurityGroup{}, &SecurityGroupList{})
}
