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
// +groupName=storagegateway.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/storagegateway/v1alpha1"
)

type GatewayNetworkInterfaceObservation struct {
	Ipv4Address string `json:"ipv4Address" tf:"ipv4_address"`
}

type GatewayNetworkInterfaceParameters struct {
}

type SmbActiveDirectorySettingsObservation struct {
	ActiveDirectoryStatus string `json:"activeDirectoryStatus" tf:"active_directory_status"`
}

type SmbActiveDirectorySettingsParameters struct {
	DomainControllers []string `json:"domainControllers,omitempty" tf:"domain_controllers"`

	DomainName string `json:"domainName" tf:"domain_name"`

	OrganizationalUnit *string `json:"organizationalUnit,omitempty" tf:"organizational_unit"`

	Password string `json:"password" tf:"password"`

	TimeoutInSeconds *int64 `json:"timeoutInSeconds,omitempty" tf:"timeout_in_seconds"`

	Username string `json:"username" tf:"username"`
}

type StoragegatewayGatewayObservation struct {
	Arn string `json:"arn" tf:"arn"`

	Ec2InstanceId string `json:"ec2InstanceId" tf:"ec2_instance_id"`

	EndpointType string `json:"endpointType" tf:"endpoint_type"`

	GatewayId string `json:"gatewayId" tf:"gateway_id"`

	GatewayNetworkInterface []GatewayNetworkInterfaceObservation `json:"gatewayNetworkInterface" tf:"gateway_network_interface"`

	HostEnvironment string `json:"hostEnvironment" tf:"host_environment"`
}

type StoragegatewayGatewayParameters struct {
	ActivationKey *string `json:"activationKey,omitempty" tf:"activation_key"`

	AverageDownloadRateLimitInBitsPerSec *int64 `json:"averageDownloadRateLimitInBitsPerSec,omitempty" tf:"average_download_rate_limit_in_bits_per_sec"`

	AverageUploadRateLimitInBitsPerSec *int64 `json:"averageUploadRateLimitInBitsPerSec,omitempty" tf:"average_upload_rate_limit_in_bits_per_sec"`

	CloudwatchLogGroupArn *string `json:"cloudwatchLogGroupArn,omitempty" tf:"cloudwatch_log_group_arn"`

	GatewayIpAddress *string `json:"gatewayIpAddress,omitempty" tf:"gateway_ip_address"`

	GatewayName string `json:"gatewayName" tf:"gateway_name"`

	GatewayTimezone string `json:"gatewayTimezone" tf:"gateway_timezone"`

	GatewayType *string `json:"gatewayType,omitempty" tf:"gateway_type"`

	GatewayVpcEndpoint *string `json:"gatewayVpcEndpoint,omitempty" tf:"gateway_vpc_endpoint"`

	MediumChangerType *string `json:"mediumChangerType,omitempty" tf:"medium_changer_type"`

	SmbActiveDirectorySettings []SmbActiveDirectorySettingsParameters `json:"smbActiveDirectorySettings,omitempty" tf:"smb_active_directory_settings"`

	SmbFileShareVisibility *bool `json:"smbFileShareVisibility,omitempty" tf:"smb_file_share_visibility"`

	SmbGuestPassword *string `json:"smbGuestPassword,omitempty" tf:"smb_guest_password"`

	SmbSecurityStrategy *string `json:"smbSecurityStrategy,omitempty" tf:"smb_security_strategy"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`

	TapeDriveType *string `json:"tapeDriveType,omitempty" tf:"tape_drive_type"`
}

// StoragegatewayGatewaySpec defines the desired state of StoragegatewayGateway
type StoragegatewayGatewaySpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       StoragegatewayGatewayParameters `json:"forProvider"`
}

// StoragegatewayGatewayStatus defines the observed state of StoragegatewayGateway.
type StoragegatewayGatewayStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          StoragegatewayGatewayObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// StoragegatewayGateway is the Schema for the StoragegatewayGateways API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type StoragegatewayGateway struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StoragegatewayGatewaySpec   `json:"spec"`
	Status            StoragegatewayGatewayStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// StoragegatewayGatewayList contains a list of StoragegatewayGateways
type StoragegatewayGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StoragegatewayGateway `json:"items"`
}

// Repository type metadata.
var (
	StoragegatewayGatewayKind             = "StoragegatewayGateway"
	StoragegatewayGatewayGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: StoragegatewayGatewayKind}.String()
	StoragegatewayGatewayKindAPIVersion   = StoragegatewayGatewayKind + "." + v1alpha1.GroupVersion.String()
	StoragegatewayGatewayGroupVersionKind = v1alpha1.GroupVersion.WithKind(StoragegatewayGatewayKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&StoragegatewayGateway{}, &StoragegatewayGatewayList{})
}