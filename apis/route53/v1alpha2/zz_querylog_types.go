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

type QueryLogObservation struct {

	// The Amazon Resource Name (ARN) of the Query Logging Config.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The query logging configuration ID
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type QueryLogParameters struct {

	// CloudWatch log group ARN to send query logs.
	// +kubebuilder:validation:Required
	CloudwatchLogGroupArn *string `json:"cloudwatchLogGroupArn" tf:"cloudwatch_log_group_arn,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Route53 hosted zone ID to enable query logs.
	// +kubebuilder:validation:Required
	ZoneID *string `json:"zoneId" tf:"zone_id,omitempty"`
}

// QueryLogSpec defines the desired state of QueryLog
type QueryLogSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     QueryLogParameters `json:"forProvider"`
}

// QueryLogStatus defines the observed state of QueryLog.
type QueryLogStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        QueryLogObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// QueryLog is the Schema for the QueryLogs API. Provides a Route53 query logging configuration resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type QueryLog struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              QueryLogSpec   `json:"spec"`
	Status            QueryLogStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// QueryLogList contains a list of QueryLogs
type QueryLogList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []QueryLog `json:"items"`
}

// Repository type metadata.
var (
	QueryLog_Kind             = "QueryLog"
	QueryLog_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: QueryLog_Kind}.String()
	QueryLog_KindAPIVersion   = QueryLog_Kind + "." + CRDGroupVersion.String()
	QueryLog_GroupVersionKind = CRDGroupVersion.WithKind(QueryLog_Kind)
)

func init() {
	SchemeBuilder.Register(&QueryLog{}, &QueryLogList{})
}
