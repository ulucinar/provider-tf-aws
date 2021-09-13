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

type VpcEndpointConnectionNotificationObservation struct {
	NotificationType string `json:"notificationType,omitempty" tf:"notification_type"`

	State string `json:"state,omitempty" tf:"state"`
}

type VpcEndpointConnectionNotificationParameters struct {

	// +kubebuilder:validation:Required
	ConnectionEvents []string `json:"connectionEvents" tf:"connection_events"`

	// +kubebuilder:validation:Required
	ConnectionNotificationArn string `json:"connectionNotificationArn" tf:"connection_notification_arn"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	VpcEndpointID *string `json:"vpcEndpointId,omitempty" tf:"vpc_endpoint_id"`

	// +kubebuilder:validation:Optional
	VpcEndpointServiceID *string `json:"vpcEndpointServiceId,omitempty" tf:"vpc_endpoint_service_id"`
}

// VpcEndpointConnectionNotificationSpec defines the desired state of VpcEndpointConnectionNotification
type VpcEndpointConnectionNotificationSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       VpcEndpointConnectionNotificationParameters `json:"forProvider"`
}

// VpcEndpointConnectionNotificationStatus defines the observed state of VpcEndpointConnectionNotification.
type VpcEndpointConnectionNotificationStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          VpcEndpointConnectionNotificationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VpcEndpointConnectionNotification is the Schema for the VpcEndpointConnectionNotifications API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type VpcEndpointConnectionNotification struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VpcEndpointConnectionNotificationSpec   `json:"spec"`
	Status            VpcEndpointConnectionNotificationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VpcEndpointConnectionNotificationList contains a list of VpcEndpointConnectionNotifications
type VpcEndpointConnectionNotificationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VpcEndpointConnectionNotification `json:"items"`
}

// Repository type metadata.
var (
	VpcEndpointConnectionNotificationKind             = "VpcEndpointConnectionNotification"
	VpcEndpointConnectionNotificationGroupKind        = schema.GroupKind{Group: Group, Kind: VpcEndpointConnectionNotificationKind}.String()
	VpcEndpointConnectionNotificationKindAPIVersion   = VpcEndpointConnectionNotificationKind + "." + GroupVersion.String()
	VpcEndpointConnectionNotificationGroupVersionKind = GroupVersion.WithKind(VpcEndpointConnectionNotificationKind)
)

func init() {
	SchemeBuilder.Register(&VpcEndpointConnectionNotification{}, &VpcEndpointConnectionNotificationList{})
}
