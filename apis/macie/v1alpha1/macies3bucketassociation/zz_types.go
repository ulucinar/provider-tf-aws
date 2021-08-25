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
// +groupName=macie.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/macie/v1alpha1"
)

type ClassificationTypeObservation struct {
}

type ClassificationTypeParameters struct {
	Continuous *string `json:"continuous,omitempty" tf:"continuous"`

	OneTime *string `json:"oneTime,omitempty" tf:"one_time"`
}

type MacieS3BucketAssociationObservation struct {
}

type MacieS3BucketAssociationParameters struct {
	BucketName string `json:"bucketName" tf:"bucket_name"`

	ClassificationType []ClassificationTypeParameters `json:"classificationType,omitempty" tf:"classification_type"`

	MemberAccountId *string `json:"memberAccountId,omitempty" tf:"member_account_id"`

	Prefix *string `json:"prefix,omitempty" tf:"prefix"`
}

// MacieS3BucketAssociationSpec defines the desired state of MacieS3BucketAssociation
type MacieS3BucketAssociationSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       MacieS3BucketAssociationParameters `json:"forProvider"`
}

// MacieS3BucketAssociationStatus defines the observed state of MacieS3BucketAssociation.
type MacieS3BucketAssociationStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          MacieS3BucketAssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// MacieS3BucketAssociation is the Schema for the MacieS3BucketAssociations API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type MacieS3BucketAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MacieS3BucketAssociationSpec   `json:"spec"`
	Status            MacieS3BucketAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MacieS3BucketAssociationList contains a list of MacieS3BucketAssociations
type MacieS3BucketAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MacieS3BucketAssociation `json:"items"`
}

// Repository type metadata.
var (
	MacieS3BucketAssociationKind             = "MacieS3BucketAssociation"
	MacieS3BucketAssociationGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: MacieS3BucketAssociationKind}.String()
	MacieS3BucketAssociationKindAPIVersion   = MacieS3BucketAssociationKind + "." + v1alpha1.GroupVersion.String()
	MacieS3BucketAssociationGroupVersionKind = v1alpha1.GroupVersion.WithKind(MacieS3BucketAssociationKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&MacieS3BucketAssociation{}, &MacieS3BucketAssociationList{})
}