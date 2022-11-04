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

type HealthCheckObservation struct {
}

type HealthCheckParameters struct {

	// Whether health checks are enabled. Defaults to true.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Number of consecutive health checks successes required before considering an unhealthy target healthy. Defaults to 3.
	// +kubebuilder:validation:Optional
	HealthyThreshold *float64 `json:"healthyThreshold,omitempty" tf:"healthy_threshold,omitempty"`

	// Approximate amount of time, in seconds, between health checks of an individual target. Minimum value 5 seconds, Maximum value 300 seconds. For lambda target groups, it needs to be greater as the timeout of the underlying lambda. Default 30 seconds.
	// +kubebuilder:validation:Optional
	Interval *float64 `json:"interval,omitempty" tf:"interval,omitempty"`

	// 299" or "0-99"). Required for HTTP/HTTPS/GRPC ALB. Only applies to Application Load Balancers (i.e., HTTP/HTTPS/GRPC) not Network Load Balancers (i.e., TCP).
	// +kubebuilder:validation:Optional
	Matcher *string `json:"matcher,omitempty" tf:"matcher,omitempty"`

	// (May be required) Destination for the health check request. Required for HTTP/HTTPS ALB and HTTP NLB. Only applies to HTTP/HTTPS.
	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// Port to use to connect with the target. Valid values are either ports 1-65535, or traffic-port. Defaults to traffic-port.
	// +kubebuilder:validation:Optional
	Port *string `json:"port,omitempty" tf:"port,omitempty"`

	// Protocol to use to connect with the target. Defaults to HTTP. Not applicable when target_type is lambda.
	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// Amount of time, in seconds, during which no response means a failed health check. For Application Load Balancers, the range is 2 to 120 seconds, and the default is 5 seconds for the instance target type and 30 seconds for the lambda target type. For Network Load Balancers, you cannot set a custom value, and the default is 10 seconds for TCP and HTTPS health checks and 6 seconds for HTTP health checks.
	// +kubebuilder:validation:Optional
	Timeout *float64 `json:"timeout,omitempty" tf:"timeout,omitempty"`

	// Number of consecutive health check failures required before considering the target unhealthy. For Network Load Balancers, this value must be the same as the healthy_threshold. Defaults to 3.
	// +kubebuilder:validation:Optional
	UnhealthyThreshold *float64 `json:"unhealthyThreshold,omitempty" tf:"unhealthy_threshold,omitempty"`
}

type LBTargetGroupObservation struct {

	// ARN of the Target Group (matches id).
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// ARN suffix for use with CloudWatch Metrics.
	ArnSuffix *string `json:"arnSuffix,omitempty" tf:"arn_suffix,omitempty"`

	// ARN of the Target Group (matches arn).
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type LBTargetGroupParameters struct {

	// Whether to terminate connections at the end of the deregistration timeout on Network Load Balancers. See doc for more information. Default is false.
	// +kubebuilder:validation:Optional
	ConnectionTermination *bool `json:"connectionTermination,omitempty" tf:"connection_termination,omitempty"`

	// Amount time for Elastic Load Balancing to wait before changing the state of a deregistering target from draining to unused. The range is 0-3600 seconds. The default value is 300 seconds.
	// +kubebuilder:validation:Optional
	DeregistrationDelay *string `json:"deregistrationDelay,omitempty" tf:"deregistration_delay,omitempty"`

	// Health Check configuration block. Detailed below.
	// +kubebuilder:validation:Optional
	HealthCheck []HealthCheckParameters `json:"healthCheck,omitempty" tf:"health_check,omitempty"`

	// Whether the request and response headers exchanged between the load balancer and the Lambda function include arrays of values or strings. Only applies when target_type is lambda. Default is false.
	// +kubebuilder:validation:Optional
	LambdaMultiValueHeadersEnabled *bool `json:"lambdaMultiValueHeadersEnabled,omitempty" tf:"lambda_multi_value_headers_enabled,omitempty"`

	// Determines how the load balancer selects targets when routing requests. Only applicable for Application Load Balancer Target Groups. The value is round_robin or least_outstanding_requests. The default is round_robin.
	// +kubebuilder:validation:Optional
	LoadBalancingAlgorithmType *string `json:"loadBalancingAlgorithmType,omitempty" tf:"load_balancing_algorithm_type,omitempty"`

	// Name of the target group.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// (May be required, Forces new resource) Port on which targets receive traffic, unless overridden when registering a specific target. Required when target_type is instance, ip or alb. Does not apply when target_type is lambda.
	// +kubebuilder:validation:Optional
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Whether client IP preservation is enabled. See doc for more information.
	// +kubebuilder:validation:Optional
	PreserveClientIP *string `json:"preserveClientIp,omitempty" tf:"preserve_client_ip,omitempty"`

	// (May be required, Forces new resource) Protocol to use for routing traffic to the targets. Should be one of GENEVE, HTTP, HTTPS, TCP, TCP_UDP, TLS, or UDP. Required when target_type is instance, ip or alb. Does not apply when target_type is lambda.
	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// Only applicable when protocol is HTTP or HTTPS. The protocol version. Specify GRPC to send requests to targets using gRPC. Specify HTTP2 to send requests to targets using HTTP/2. The default is HTTP1, which sends requests to targets using HTTP/1.1
	// +kubebuilder:validation:Optional
	ProtocolVersion *string `json:"protocolVersion,omitempty" tf:"protocol_version,omitempty"`

	// Whether to enable support for proxy protocol v2 on Network Load Balancers. See doc for more information. Default is false.
	// +kubebuilder:validation:Optional
	ProxyProtocolV2 *bool `json:"proxyProtocolV2,omitempty" tf:"proxy_protocol_v2,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Amount time for targets to warm up before the load balancer sends them a full share of requests. The range is 30-900 seconds or 0 to disable. The default value is 0 seconds.
	// +kubebuilder:validation:Optional
	SlowStart *float64 `json:"slowStart,omitempty" tf:"slow_start,omitempty"`

	// Stickiness configuration block. Detailed below.
	// +kubebuilder:validation:Optional
	Stickiness []LBTargetGroupStickinessParameters `json:"stickiness,omitempty" tf:"stickiness,omitempty"`

	// Map of tags to assign to the resource. If configured with a provider default_tags configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// (May be required, Forces new resource) Type of target that you must specify when registering targets with this target group. See doc for supported values. The default is instance.
	// +kubebuilder:validation:Optional
	TargetType *string `json:"targetType,omitempty" tf:"target_type,omitempty"`

	// Identifier of the VPC in which to create the target group. Required when target_type is instance, ip or alb. Does not apply when target_type is lambda.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-aws/apis/ec2/v1alpha2.VPC
	// +kubebuilder:validation:Optional
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// Reference to a VPC in ec2 to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDRef *v1.Reference `json:"vpcIdRef,omitempty" tf:"-"`

	// Selector for a VPC in ec2 to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDSelector *v1.Selector `json:"vpcIdSelector,omitempty" tf:"-"`
}

type LBTargetGroupStickinessObservation struct {
}

type LBTargetGroupStickinessParameters struct {

	// Only used when the type is lb_cookie. The time period, in seconds, during which requests from a client should be routed to the same target. After this time period expires, the load balancer-generated cookie is considered stale. The range is 1 second to 1 week (604800 seconds). The default value is 1 day (86400 seconds).
	// +kubebuilder:validation:Optional
	CookieDuration *float64 `json:"cookieDuration,omitempty" tf:"cookie_duration,omitempty"`

	// Name of the application based cookie. AWSALB, AWSALBAPP, and AWSALBTG prefixes are reserved and cannot be used. Only needed when type is app_cookie.
	// +kubebuilder:validation:Optional
	CookieName *string `json:"cookieName,omitempty" tf:"cookie_name,omitempty"`

	// Whether health checks are enabled. Defaults to true.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The type of sticky sessions. The only current possible values are lb_cookie, app_cookie for ALBs, and source_ip for NLBs.
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

// LBTargetGroup is the Schema for the LBTargetGroups API. Provides a Target Group resource for use with Load Balancers.
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
