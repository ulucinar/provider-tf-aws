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
// +groupName=cloudwatch.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/cloudwatch/v1alpha1"
)

type CloudwatchLogDestinationPolicyObservation struct {
}

type CloudwatchLogDestinationPolicyParameters struct {
	AccessPolicy string `json:"accessPolicy" tf:"access_policy"`

	DestinationName string `json:"destinationName" tf:"destination_name"`
}

// CloudwatchLogDestinationPolicySpec defines the desired state of CloudwatchLogDestinationPolicy
type CloudwatchLogDestinationPolicySpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       CloudwatchLogDestinationPolicyParameters `json:"forProvider"`
}

// CloudwatchLogDestinationPolicyStatus defines the observed state of CloudwatchLogDestinationPolicy.
type CloudwatchLogDestinationPolicyStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          CloudwatchLogDestinationPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CloudwatchLogDestinationPolicy is the Schema for the CloudwatchLogDestinationPolicys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type CloudwatchLogDestinationPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CloudwatchLogDestinationPolicySpec   `json:"spec"`
	Status            CloudwatchLogDestinationPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CloudwatchLogDestinationPolicyList contains a list of CloudwatchLogDestinationPolicys
type CloudwatchLogDestinationPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudwatchLogDestinationPolicy `json:"items"`
}

// Repository type metadata.
var (
	CloudwatchLogDestinationPolicyKind             = "CloudwatchLogDestinationPolicy"
	CloudwatchLogDestinationPolicyGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: CloudwatchLogDestinationPolicyKind}.String()
	CloudwatchLogDestinationPolicyKindAPIVersion   = CloudwatchLogDestinationPolicyKind + "." + v1alpha1.GroupVersion.String()
	CloudwatchLogDestinationPolicyGroupVersionKind = v1alpha1.GroupVersion.WithKind(CloudwatchLogDestinationPolicyKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&CloudwatchLogDestinationPolicy{}, &CloudwatchLogDestinationPolicyList{})
}