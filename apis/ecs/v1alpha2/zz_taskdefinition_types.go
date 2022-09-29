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

type AuthorizationConfigObservation struct {
}

type AuthorizationConfigParameters struct {

	// +kubebuilder:validation:Optional
	AccessPointID *string `json:"accessPointId,omitempty" tf:"access_point_id,omitempty"`

	// +kubebuilder:validation:Optional
	IAM *string `json:"iam,omitempty" tf:"iam,omitempty"`
}

type DockerVolumeConfigurationObservation struct {
}

type DockerVolumeConfigurationParameters struct {

	// +kubebuilder:validation:Optional
	Autoprovision *bool `json:"autoprovision,omitempty" tf:"autoprovision,omitempty"`

	// +kubebuilder:validation:Optional
	Driver *string `json:"driver,omitempty" tf:"driver,omitempty"`

	// +kubebuilder:validation:Optional
	DriverOpts map[string]*string `json:"driverOpts,omitempty" tf:"driver_opts,omitempty"`

	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// +kubebuilder:validation:Optional
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`
}

type EFSVolumeConfigurationObservation struct {
}

type EFSVolumeConfigurationParameters struct {

	// +kubebuilder:validation:Optional
	AuthorizationConfig []AuthorizationConfigParameters `json:"authorizationConfig,omitempty" tf:"authorization_config,omitempty"`

	// +kubebuilder:validation:Required
	FileSystemID *string `json:"fileSystemId" tf:"file_system_id,omitempty"`

	// +kubebuilder:validation:Optional
	RootDirectory *string `json:"rootDirectory,omitempty" tf:"root_directory,omitempty"`

	// +kubebuilder:validation:Optional
	TransitEncryption *string `json:"transitEncryption,omitempty" tf:"transit_encryption,omitempty"`

	// +kubebuilder:validation:Optional
	TransitEncryptionPort *float64 `json:"transitEncryptionPort,omitempty" tf:"transit_encryption_port,omitempty"`
}

type EphemeralStorageObservation struct {
}

type EphemeralStorageParameters struct {

	// +kubebuilder:validation:Required
	SizeInGib *float64 `json:"sizeInGib" tf:"size_in_gib,omitempty"`
}

type FSXWindowsFileServerVolumeConfigurationAuthorizationConfigObservation struct {
}

type FSXWindowsFileServerVolumeConfigurationAuthorizationConfigParameters struct {

	// +kubebuilder:validation:Required
	CredentialsParameter *string `json:"credentialsParameter" tf:"credentials_parameter,omitempty"`

	// +kubebuilder:validation:Required
	Domain *string `json:"domain" tf:"domain,omitempty"`
}

type FSXWindowsFileServerVolumeConfigurationObservation struct {
}

type FSXWindowsFileServerVolumeConfigurationParameters struct {

	// +kubebuilder:validation:Required
	AuthorizationConfig []FSXWindowsFileServerVolumeConfigurationAuthorizationConfigParameters `json:"authorizationConfig" tf:"authorization_config,omitempty"`

	// +kubebuilder:validation:Required
	FileSystemID *string `json:"fileSystemId" tf:"file_system_id,omitempty"`

	// +kubebuilder:validation:Required
	RootDirectory *string `json:"rootDirectory" tf:"root_directory,omitempty"`
}

type InferenceAcceleratorObservation struct {
}

type InferenceAcceleratorParameters struct {

	// +kubebuilder:validation:Required
	DeviceName *string `json:"deviceName" tf:"device_name,omitempty"`

	// +kubebuilder:validation:Required
	DeviceType *string `json:"deviceType" tf:"device_type,omitempty"`
}

type ProxyConfigurationObservation struct {
}

type ProxyConfigurationParameters struct {

	// +kubebuilder:validation:Required
	ContainerName *string `json:"containerName" tf:"container_name,omitempty"`

	// +kubebuilder:validation:Optional
	Properties map[string]*string `json:"properties,omitempty" tf:"properties,omitempty"`

	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type RuntimePlatformObservation struct {
}

type RuntimePlatformParameters struct {

	// +kubebuilder:validation:Optional
	CPUArchitecture *string `json:"cpuArchitecture,omitempty" tf:"cpu_architecture,omitempty"`

	// +kubebuilder:validation:Optional
	OperatingSystemFamily *string `json:"operatingSystemFamily,omitempty" tf:"operating_system_family,omitempty"`
}

type TaskDefinitionObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Revision *float64 `json:"revision,omitempty" tf:"revision,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type TaskDefinitionParameters struct {

	// +kubebuilder:validation:Optional
	CPU *string `json:"cpu,omitempty" tf:"cpu,omitempty"`

	// +kubebuilder:validation:Required
	ContainerDefinitions *string `json:"containerDefinitions" tf:"container_definitions,omitempty"`

	// +kubebuilder:validation:Optional
	EphemeralStorage []EphemeralStorageParameters `json:"ephemeralStorage,omitempty" tf:"ephemeral_storage,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-aws/apis/iam/v1alpha2.Role
	// +crossplane:generate:reference:extractor=github.com/crossplane-contrib/provider-jet-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	ExecutionRoleArn *string `json:"executionRoleArn,omitempty" tf:"execution_role_arn,omitempty"`

	// +kubebuilder:validation:Optional
	ExecutionRoleArnRef *v1.Reference `json:"executionRoleArnRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ExecutionRoleArnSelector *v1.Selector `json:"executionRoleArnSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	Family *string `json:"family" tf:"family,omitempty"`

	// +kubebuilder:validation:Optional
	InferenceAccelerator []InferenceAcceleratorParameters `json:"inferenceAccelerator,omitempty" tf:"inference_accelerator,omitempty"`

	// +kubebuilder:validation:Optional
	IpcMode *string `json:"ipcMode,omitempty" tf:"ipc_mode,omitempty"`

	// +kubebuilder:validation:Optional
	Memory *string `json:"memory,omitempty" tf:"memory,omitempty"`

	// +kubebuilder:validation:Optional
	NetworkMode *string `json:"networkMode,omitempty" tf:"network_mode,omitempty"`

	// +kubebuilder:validation:Optional
	PidMode *string `json:"pidMode,omitempty" tf:"pid_mode,omitempty"`

	// +kubebuilder:validation:Optional
	PlacementConstraints []TaskDefinitionPlacementConstraintsParameters `json:"placementConstraints,omitempty" tf:"placement_constraints,omitempty"`

	// +kubebuilder:validation:Optional
	ProxyConfiguration []ProxyConfigurationParameters `json:"proxyConfiguration,omitempty" tf:"proxy_configuration,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	RequiresCompatibilities []*string `json:"requiresCompatibilities,omitempty" tf:"requires_compatibilities,omitempty"`

	// +kubebuilder:validation:Optional
	RuntimePlatform []RuntimePlatformParameters `json:"runtimePlatform,omitempty" tf:"runtime_platform,omitempty"`

	// +kubebuilder:validation:Optional
	SkipDestroy *bool `json:"skipDestroy,omitempty" tf:"skip_destroy,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	TaskRoleArn *string `json:"taskRoleArn,omitempty" tf:"task_role_arn,omitempty"`

	// +kubebuilder:validation:Optional
	Volume []VolumeParameters `json:"volume,omitempty" tf:"volume,omitempty"`
}

type TaskDefinitionPlacementConstraintsObservation struct {
}

type TaskDefinitionPlacementConstraintsParameters struct {

	// +kubebuilder:validation:Optional
	Expression *string `json:"expression,omitempty" tf:"expression,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type VolumeObservation struct {
}

type VolumeParameters struct {

	// +kubebuilder:validation:Optional
	DockerVolumeConfiguration []DockerVolumeConfigurationParameters `json:"dockerVolumeConfiguration,omitempty" tf:"docker_volume_configuration,omitempty"`

	// +kubebuilder:validation:Optional
	EFSVolumeConfiguration []EFSVolumeConfigurationParameters `json:"efsVolumeConfiguration,omitempty" tf:"efs_volume_configuration,omitempty"`

	// +kubebuilder:validation:Optional
	FSXWindowsFileServerVolumeConfiguration []FSXWindowsFileServerVolumeConfigurationParameters `json:"fsxWindowsFileServerVolumeConfiguration,omitempty" tf:"fsx_windows_file_server_volume_configuration,omitempty"`

	// +kubebuilder:validation:Optional
	HostPath *string `json:"hostPath,omitempty" tf:"host_path,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`
}

// TaskDefinitionSpec defines the desired state of TaskDefinition
type TaskDefinitionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TaskDefinitionParameters `json:"forProvider"`
}

// TaskDefinitionStatus defines the observed state of TaskDefinition.
type TaskDefinitionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TaskDefinitionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TaskDefinition is the Schema for the TaskDefinitions API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type TaskDefinition struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TaskDefinitionSpec   `json:"spec"`
	Status            TaskDefinitionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TaskDefinitionList contains a list of TaskDefinitions
type TaskDefinitionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TaskDefinition `json:"items"`
}

// Repository type metadata.
var (
	TaskDefinition_Kind             = "TaskDefinition"
	TaskDefinition_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TaskDefinition_Kind}.String()
	TaskDefinition_KindAPIVersion   = TaskDefinition_Kind + "." + CRDGroupVersion.String()
	TaskDefinition_GroupVersionKind = CRDGroupVersion.WithKind(TaskDefinition_Kind)
)

func init() {
	SchemeBuilder.Register(&TaskDefinition{}, &TaskDefinitionList{})
}
