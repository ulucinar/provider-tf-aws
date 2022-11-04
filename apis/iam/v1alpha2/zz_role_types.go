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

type InlinePolicyObservation struct {
}

type InlinePolicyParameters struct {

	// Friendly name of the role. See IAM Identifiers for more information.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Policy document as a JSON formatted string.
	// +kubebuilder:validation:Optional
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`
}

type RoleObservation struct {

	// Amazon Resource Name (ARN) specifying the role.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Creation date of the IAM role.
	CreateDate *string `json:"createDate,omitempty" tf:"create_date,omitempty"`

	// Name of the role.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Set of exclusive IAM managed policy ARNs to attach to the IAM role. Configuring an empty set (i.e.
	ManagedPolicyArns []*string `json:"managedPolicyArns,omitempty" tf:"managed_policy_arns,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// Stable and unique string identifying the role.
	UniqueID *string `json:"uniqueId,omitempty" tf:"unique_id,omitempty"`
}

type RoleParameters struct {

	// Policy that grants an entity permission to assume the role.
	// +kubebuilder:validation:Required
	AssumeRolePolicy *string `json:"assumeRolePolicy" tf:"assume_role_policy,omitempty"`

	// Description of the role.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Whether to force detaching any policies the role has before destroying it. Defaults to false.
	// +kubebuilder:validation:Optional
	ForceDetachPolicies *bool `json:"forceDetachPolicies,omitempty" tf:"force_detach_policies,omitempty"`

	// Configuration block defining an exclusive set of IAM inline policies associated with the IAM role. See below. Configuring one empty block (i.e.
	// +kubebuilder:validation:Optional
	InlinePolicy []InlinePolicyParameters `json:"inlinePolicy,omitempty" tf:"inline_policy,omitempty"`

	// Maximum session duration (in seconds) that you want to set for the specified role. If you do not specify a value for this setting, the default maximum of one hour is applied. This setting can have a value from 1 hour to 12 hours.
	// +kubebuilder:validation:Optional
	MaxSessionDuration *float64 `json:"maxSessionDuration,omitempty" tf:"max_session_duration,omitempty"`

	// Path to the role. See IAM Identifiers for more information.
	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// ARN of the policy that is used to set the permissions boundary for the role.
	// +kubebuilder:validation:Optional
	PermissionsBoundary *string `json:"permissionsBoundary,omitempty" tf:"permissions_boundary,omitempty"`

	// Key-value mapping of tags for the IAM role. If configured with a provider default_tags configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// RoleSpec defines the desired state of Role
type RoleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RoleParameters `json:"forProvider"`
}

// RoleStatus defines the observed state of Role.
type RoleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RoleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Role is the Schema for the Roles API. Provides an IAM role.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type Role struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RoleSpec   `json:"spec"`
	Status            RoleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RoleList contains a list of Roles
type RoleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Role `json:"items"`
}

// Repository type metadata.
var (
	Role_Kind             = "Role"
	Role_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Role_Kind}.String()
	Role_KindAPIVersion   = Role_Kind + "." + CRDGroupVersion.String()
	Role_GroupVersionKind = CRDGroupVersion.WithKind(Role_Kind)
)

func init() {
	SchemeBuilder.Register(&Role{}, &RoleList{})
}
