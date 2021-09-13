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

type IamPolicyObservation struct {
	Arn string `json:"arn,omitempty" tf:"arn"`

	PolicyID string `json:"policyId,omitempty" tf:"policy_id"`
}

type IamPolicyParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name"`

	// +kubebuilder:validation:Optional
	NamePrefix *string `json:"namePrefix,omitempty" tf:"name_prefix"`

	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path"`

	// +kubebuilder:validation:Required
	Policy string `json:"policy" tf:"policy"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	// +kubebuilder:validation:Optional
	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

// IamPolicySpec defines the desired state of IamPolicy
type IamPolicySpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       IamPolicyParameters `json:"forProvider"`
}

// IamPolicyStatus defines the observed state of IamPolicy.
type IamPolicyStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          IamPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// IamPolicy is the Schema for the IamPolicys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type IamPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IamPolicySpec   `json:"spec"`
	Status            IamPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IamPolicyList contains a list of IamPolicys
type IamPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IamPolicy `json:"items"`
}

// Repository type metadata.
var (
	IamPolicyKind             = "IamPolicy"
	IamPolicyGroupKind        = schema.GroupKind{Group: Group, Kind: IamPolicyKind}.String()
	IamPolicyKindAPIVersion   = IamPolicyKind + "." + GroupVersion.String()
	IamPolicyGroupVersionKind = GroupVersion.WithKind(IamPolicyKind)
)

func init() {
	SchemeBuilder.Register(&IamPolicy{}, &IamPolicyList{})
}
