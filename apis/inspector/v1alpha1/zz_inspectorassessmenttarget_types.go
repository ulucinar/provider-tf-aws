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

type InspectorAssessmentTargetObservation struct {
	Arn string `json:"arn" tf:"arn"`
}

type InspectorAssessmentTargetParameters struct {
	Name string `json:"name" tf:"name"`

	ResourceGroupArn *string `json:"resourceGroupArn,omitempty" tf:"resource_group_arn"`
}

// InspectorAssessmentTargetSpec defines the desired state of InspectorAssessmentTarget
type InspectorAssessmentTargetSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       InspectorAssessmentTargetParameters `json:"forProvider"`
}

// InspectorAssessmentTargetStatus defines the observed state of InspectorAssessmentTarget.
type InspectorAssessmentTargetStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          InspectorAssessmentTargetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// InspectorAssessmentTarget is the Schema for the InspectorAssessmentTargets API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type InspectorAssessmentTarget struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InspectorAssessmentTargetSpec   `json:"spec"`
	Status            InspectorAssessmentTargetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InspectorAssessmentTargetList contains a list of InspectorAssessmentTargets
type InspectorAssessmentTargetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []InspectorAssessmentTarget `json:"items"`
}

// Repository type metadata.
var (
	InspectorAssessmentTargetKind             = "InspectorAssessmentTarget"
	InspectorAssessmentTargetGroupKind        = schema.GroupKind{Group: Group, Kind: InspectorAssessmentTargetKind}.String()
	InspectorAssessmentTargetKindAPIVersion   = InspectorAssessmentTargetKind + "." + GroupVersion.String()
	InspectorAssessmentTargetGroupVersionKind = GroupVersion.WithKind(InspectorAssessmentTargetKind)
)

func init() {
	SchemeBuilder.Register(&InspectorAssessmentTarget{}, &InspectorAssessmentTargetList{})
}