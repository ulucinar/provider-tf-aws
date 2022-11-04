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

type UserPolicyAttachmentObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type UserPolicyAttachmentParameters struct {

	// The ARN of the policy you want to apply
	// +crossplane:generate:reference:type=Policy
	// +crossplane:generate:reference:extractor=github.com/crossplane-contrib/provider-jet-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	PolicyArn *string `json:"policyArn,omitempty" tf:"policy_arn,omitempty"`

	// Reference to a Policy to populate policyArn.
	// +kubebuilder:validation:Optional
	PolicyArnRef *v1.Reference `json:"policyArnRef,omitempty" tf:"-"`

	// Selector for a Policy to populate policyArn.
	// +kubebuilder:validation:Optional
	PolicyArnSelector *v1.Selector `json:"policyArnSelector,omitempty" tf:"-"`

	// The user the policy should be applied to
	// +crossplane:generate:reference:type=User
	// +kubebuilder:validation:Optional
	User *string `json:"user,omitempty" tf:"user,omitempty"`

	// Reference to a User to populate user.
	// +kubebuilder:validation:Optional
	UserRef *v1.Reference `json:"userRef,omitempty" tf:"-"`

	// Selector for a User to populate user.
	// +kubebuilder:validation:Optional
	UserSelector *v1.Selector `json:"userSelector,omitempty" tf:"-"`
}

// UserPolicyAttachmentSpec defines the desired state of UserPolicyAttachment
type UserPolicyAttachmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     UserPolicyAttachmentParameters `json:"forProvider"`
}

// UserPolicyAttachmentStatus defines the observed state of UserPolicyAttachment.
type UserPolicyAttachmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        UserPolicyAttachmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// UserPolicyAttachment is the Schema for the UserPolicyAttachments API. Attaches a Managed IAM Policy to an IAM user
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type UserPolicyAttachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              UserPolicyAttachmentSpec   `json:"spec"`
	Status            UserPolicyAttachmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// UserPolicyAttachmentList contains a list of UserPolicyAttachments
type UserPolicyAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []UserPolicyAttachment `json:"items"`
}

// Repository type metadata.
var (
	UserPolicyAttachment_Kind             = "UserPolicyAttachment"
	UserPolicyAttachment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: UserPolicyAttachment_Kind}.String()
	UserPolicyAttachment_KindAPIVersion   = UserPolicyAttachment_Kind + "." + CRDGroupVersion.String()
	UserPolicyAttachment_GroupVersionKind = CRDGroupVersion.WithKind(UserPolicyAttachment_Kind)
)

func init() {
	SchemeBuilder.Register(&UserPolicyAttachment{}, &UserPolicyAttachmentList{})
}
