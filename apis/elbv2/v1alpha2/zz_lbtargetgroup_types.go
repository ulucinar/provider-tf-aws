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

package v1alpha2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type HealthCheckObservation struct {
}

type HealthCheckParameters struct {

	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Optional
	HealthyThreshold *float64 `json:"healthyThreshold,omitempty" tf:"healthy_threshold,omitempty"`

	// +kubebuilder:validation:Optional
	Interval *float64 `json:"interval,omitempty" tf:"interval,omitempty"`

	// +kubebuilder:validation:Optional
	Matcher *string `json:"matcher,omitempty" tf:"matcher,omitempty"`

	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// +kubebuilder:validation:Optional
	Port *string `json:"port,omitempty" tf:"port,omitempty"`

	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// +kubebuilder:validation:Optional
	Timeout *float64 `json:"timeout,omitempty" tf:"timeout,omitempty"`

	// +kubebuilder:validation:Optional
	UnhealthyThreshold *float64 `json:"unhealthyThreshold,omitempty" tf:"unhealthy_threshold,omitempty"`
}

type LBTargetGroupObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	ArnSuffix *string `json:"arnSuffix,omitempty" tf:"arn_suffix,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type LBTargetGroupParameters struct {

	// +kubebuilder:validation:Optional
	ConnectionTermination *bool `json:"connectionTermination,omitempty" tf:"connection_termination,omitempty"`

	// +kubebuilder:validation:Optional
	DeregistrationDelay *string `json:"deregistrationDelay,omitempty" tf:"deregistration_delay,omitempty"`

	// +kubebuilder:validation:Optional
	HealthCheck []HealthCheckParameters `json:"healthCheck,omitempty" tf:"health_check,omitempty"`

	// +kubebuilder:validation:Optional
	LambdaMultiValueHeadersEnabled *bool `json:"lambdaMultiValueHeadersEnabled,omitempty" tf:"lambda_multi_value_headers_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	LoadBalancingAlgorithmType *string `json:"loadBalancingAlgorithmType,omitempty" tf:"load_balancing_algorithm_type,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// +kubebuilder:validation:Optional
	PreserveClientIP *string `json:"preserveClientIp,omitempty" tf:"preserve_client_ip,omitempty"`

	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// +kubebuilder:validation:Optional
	ProtocolVersion *string `json:"protocolVersion,omitempty" tf:"protocol_version,omitempty"`

	// +kubebuilder:validation:Optional
	ProxyProtocolV2 *bool `json:"proxyProtocolV2,omitempty" tf:"proxy_protocol_v2,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	SlowStart *float64 `json:"slowStart,omitempty" tf:"slow_start,omitempty"`

	// +kubebuilder:validation:Optional
	Stickiness []LBTargetGroupStickinessParameters `json:"stickiness,omitempty" tf:"stickiness,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	TargetType *string `json:"targetType,omitempty" tf:"target_type,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-aws/apis/ec2/v1alpha2.VPC
	// +crossplane:generate:reference:refFieldName=VpcIdRef
	// +crossplane:generate:reference:selectorFieldName=VpcIdSelector
	// +kubebuilder:validation:Optional
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// +kubebuilder:validation:Optional
	VpcIdRef *v1.Reference `json:"vpcIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	VpcIdSelector *v1.Selector `json:"vpcIdSelector,omitempty" tf:"-"`
}

type LBTargetGroupStickinessObservation struct {
}

type LBTargetGroupStickinessParameters struct {

	// +kubebuilder:validation:Optional
	CookieDuration *float64 `json:"cookieDuration,omitempty" tf:"cookie_duration,omitempty"`

	// +kubebuilder:validation:Optional
	CookieName *string `json:"cookieName,omitempty" tf:"cookie_name,omitempty"`

	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

// LBTargetGroupSpec defines the desired state of LBTargetGroup
type LBTargetGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LBTargetGroupParameters `json:"forProvider"`
}

// LBTargetGroupStatus defines the observed state of LBTargetGroup.
type LBTargetGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LBTargetGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LBTargetGroup is the Schema for the LBTargetGroups API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type LBTargetGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LBTargetGroupSpec   `json:"spec"`
	Status            LBTargetGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LBTargetGroupList contains a list of LBTargetGroups
type LBTargetGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LBTargetGroup `json:"items"`
}

// Repository type metadata.
var (
	LBTargetGroup_Kind             = "LBTargetGroup"
	LBTargetGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LBTargetGroup_Kind}.String()
	LBTargetGroup_KindAPIVersion   = LBTargetGroup_Kind + "." + CRDGroupVersion.String()
	LBTargetGroup_GroupVersionKind = CRDGroupVersion.WithKind(LBTargetGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&LBTargetGroup{}, &LBTargetGroupList{})
}
