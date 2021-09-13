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

type DefaultNetworkAclObservation struct {
	Arn string `json:"arn,omitempty" tf:"arn"`

	OwnerID string `json:"ownerId,omitempty" tf:"owner_id"`

	VpcID string `json:"vpcId,omitempty" tf:"vpc_id"`
}

type DefaultNetworkAclParameters struct {

	// +kubebuilder:validation:Required
	DefaultNetworkACLID string `json:"defaultNetworkAclId" tf:"default_network_acl_id"`

	// +kubebuilder:validation:Optional
	Egress []EgressParameters `json:"egress,omitempty" tf:"egress"`

	// +kubebuilder:validation:Optional
	Ingress []IngressParameters `json:"ingress,omitempty" tf:"ingress"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	SubnetIds []string `json:"subnetIds,omitempty" tf:"subnet_ids"`

	// +kubebuilder:validation:Optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	// +kubebuilder:validation:Optional
	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

type EgressObservation struct {
}

type EgressParameters struct {

	// +kubebuilder:validation:Required
	Action string `json:"action" tf:"action"`

	// +kubebuilder:validation:Optional
	CidrBlock *string `json:"cidrBlock,omitempty" tf:"cidr_block"`

	// +kubebuilder:validation:Required
	FromPort int64 `json:"fromPort" tf:"from_port"`

	// +kubebuilder:validation:Optional
	IPv6CidrBlock *string `json:"ipv6CidrBlock,omitempty" tf:"ipv6_cidr_block"`

	// +kubebuilder:validation:Optional
	IcmpCode *int64 `json:"icmpCode,omitempty" tf:"icmp_code"`

	// +kubebuilder:validation:Optional
	IcmpType *int64 `json:"icmpType,omitempty" tf:"icmp_type"`

	// +kubebuilder:validation:Required
	Protocol string `json:"protocol" tf:"protocol"`

	// +kubebuilder:validation:Required
	RuleNo int64 `json:"ruleNo" tf:"rule_no"`

	// +kubebuilder:validation:Required
	ToPort int64 `json:"toPort" tf:"to_port"`
}

type IngressObservation struct {
}

type IngressParameters struct {

	// +kubebuilder:validation:Required
	Action string `json:"action" tf:"action"`

	// +kubebuilder:validation:Optional
	CidrBlock *string `json:"cidrBlock,omitempty" tf:"cidr_block"`

	// +kubebuilder:validation:Required
	FromPort int64 `json:"fromPort" tf:"from_port"`

	// +kubebuilder:validation:Optional
	IPv6CidrBlock *string `json:"ipv6CidrBlock,omitempty" tf:"ipv6_cidr_block"`

	// +kubebuilder:validation:Optional
	IcmpCode *int64 `json:"icmpCode,omitempty" tf:"icmp_code"`

	// +kubebuilder:validation:Optional
	IcmpType *int64 `json:"icmpType,omitempty" tf:"icmp_type"`

	// +kubebuilder:validation:Required
	Protocol string `json:"protocol" tf:"protocol"`

	// +kubebuilder:validation:Required
	RuleNo int64 `json:"ruleNo" tf:"rule_no"`

	// +kubebuilder:validation:Required
	ToPort int64 `json:"toPort" tf:"to_port"`
}

// DefaultNetworkAclSpec defines the desired state of DefaultNetworkAcl
type DefaultNetworkAclSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       DefaultNetworkAclParameters `json:"forProvider"`
}

// DefaultNetworkAclStatus defines the observed state of DefaultNetworkAcl.
type DefaultNetworkAclStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          DefaultNetworkAclObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DefaultNetworkAcl is the Schema for the DefaultNetworkAcls API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type DefaultNetworkAcl struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DefaultNetworkAclSpec   `json:"spec"`
	Status            DefaultNetworkAclStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DefaultNetworkAclList contains a list of DefaultNetworkAcls
type DefaultNetworkAclList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DefaultNetworkAcl `json:"items"`
}

// Repository type metadata.
var (
	DefaultNetworkAclKind             = "DefaultNetworkAcl"
	DefaultNetworkAclGroupKind        = schema.GroupKind{Group: Group, Kind: DefaultNetworkAclKind}.String()
	DefaultNetworkAclKindAPIVersion   = DefaultNetworkAclKind + "." + GroupVersion.String()
	DefaultNetworkAclGroupVersionKind = GroupVersion.WithKind(DefaultNetworkAclKind)
)

func init() {
	SchemeBuilder.Register(&DefaultNetworkAcl{}, &DefaultNetworkAclList{})
}
