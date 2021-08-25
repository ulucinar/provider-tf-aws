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
// +groupName=dms.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/dms/v1alpha1"
)

type DmsReplicationSubnetGroupObservation struct {
	ReplicationSubnetGroupArn string `json:"replicationSubnetGroupArn" tf:"replication_subnet_group_arn"`

	VpcId string `json:"vpcId" tf:"vpc_id"`
}

type DmsReplicationSubnetGroupParameters struct {
	ReplicationSubnetGroupDescription string `json:"replicationSubnetGroupDescription" tf:"replication_subnet_group_description"`

	ReplicationSubnetGroupId string `json:"replicationSubnetGroupId" tf:"replication_subnet_group_id"`

	SubnetIds []string `json:"subnetIds" tf:"subnet_ids"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

// DmsReplicationSubnetGroupSpec defines the desired state of DmsReplicationSubnetGroup
type DmsReplicationSubnetGroupSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       DmsReplicationSubnetGroupParameters `json:"forProvider"`
}

// DmsReplicationSubnetGroupStatus defines the observed state of DmsReplicationSubnetGroup.
type DmsReplicationSubnetGroupStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          DmsReplicationSubnetGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DmsReplicationSubnetGroup is the Schema for the DmsReplicationSubnetGroups API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type DmsReplicationSubnetGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DmsReplicationSubnetGroupSpec   `json:"spec"`
	Status            DmsReplicationSubnetGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DmsReplicationSubnetGroupList contains a list of DmsReplicationSubnetGroups
type DmsReplicationSubnetGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DmsReplicationSubnetGroup `json:"items"`
}

// Repository type metadata.
var (
	DmsReplicationSubnetGroupKind             = "DmsReplicationSubnetGroup"
	DmsReplicationSubnetGroupGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: DmsReplicationSubnetGroupKind}.String()
	DmsReplicationSubnetGroupKindAPIVersion   = DmsReplicationSubnetGroupKind + "." + v1alpha1.GroupVersion.String()
	DmsReplicationSubnetGroupGroupVersionKind = v1alpha1.GroupVersion.WithKind(DmsReplicationSubnetGroupKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&DmsReplicationSubnetGroup{}, &DmsReplicationSubnetGroupList{})
}