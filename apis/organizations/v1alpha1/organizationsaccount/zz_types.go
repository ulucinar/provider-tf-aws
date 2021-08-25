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

// +kubebuilder:object:generate=true
// +groupName=organizations.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/organizations/v1alpha1"
)

type OrganizationsAccountObservation struct {
	Arn string `json:"arn" tf:"arn"`

	JoinedMethod string `json:"joinedMethod" tf:"joined_method"`

	JoinedTimestamp string `json:"joinedTimestamp" tf:"joined_timestamp"`

	Status string `json:"status" tf:"status"`
}

type OrganizationsAccountParameters struct {
	Email string `json:"email" tf:"email"`

	IamUserAccessToBilling *string `json:"iamUserAccessToBilling,omitempty" tf:"iam_user_access_to_billing"`

	Name string `json:"name" tf:"name"`

	ParentId *string `json:"parentId,omitempty" tf:"parent_id"`

	RoleName *string `json:"roleName,omitempty" tf:"role_name"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

// OrganizationsAccountSpec defines the desired state of OrganizationsAccount
type OrganizationsAccountSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       OrganizationsAccountParameters `json:"forProvider"`
}

// OrganizationsAccountStatus defines the observed state of OrganizationsAccount.
type OrganizationsAccountStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          OrganizationsAccountObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// OrganizationsAccount is the Schema for the OrganizationsAccounts API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type OrganizationsAccount struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              OrganizationsAccountSpec   `json:"spec"`
	Status            OrganizationsAccountStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OrganizationsAccountList contains a list of OrganizationsAccounts
type OrganizationsAccountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OrganizationsAccount `json:"items"`
}

// Repository type metadata.
var (
	OrganizationsAccountKind             = "OrganizationsAccount"
	OrganizationsAccountGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: OrganizationsAccountKind}.String()
	OrganizationsAccountKindAPIVersion   = OrganizationsAccountKind + "." + v1alpha1.GroupVersion.String()
	OrganizationsAccountGroupVersionKind = v1alpha1.GroupVersion.WithKind(OrganizationsAccountKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&OrganizationsAccount{}, &OrganizationsAccountList{})
}