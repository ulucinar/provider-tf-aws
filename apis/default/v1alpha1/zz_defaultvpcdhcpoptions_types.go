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

type DefaultVpcDhcpOptionsObservation struct {
	Arn string `json:"arn,omitempty" tf:"arn"`

	DomainName string `json:"domainName,omitempty" tf:"domain_name"`

	DomainNameServers string `json:"domainNameServers,omitempty" tf:"domain_name_servers"`

	NtpServers string `json:"ntpServers,omitempty" tf:"ntp_servers"`
}

type DefaultVpcDhcpOptionsParameters struct {

	// +kubebuilder:validation:Optional
	NetbiosNameServers []string `json:"netbiosNameServers,omitempty" tf:"netbios_name_servers"`

	// +kubebuilder:validation:Optional
	NetbiosNodeType *string `json:"netbiosNodeType,omitempty" tf:"netbios_node_type"`

	// +kubebuilder:validation:Optional
	OwnerID *string `json:"ownerId,omitempty" tf:"owner_id"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	// +kubebuilder:validation:Optional
	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

// DefaultVpcDhcpOptionsSpec defines the desired state of DefaultVpcDhcpOptions
type DefaultVpcDhcpOptionsSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       DefaultVpcDhcpOptionsParameters `json:"forProvider"`
}

// DefaultVpcDhcpOptionsStatus defines the observed state of DefaultVpcDhcpOptions.
type DefaultVpcDhcpOptionsStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          DefaultVpcDhcpOptionsObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DefaultVpcDhcpOptions is the Schema for the DefaultVpcDhcpOptionss API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type DefaultVpcDhcpOptions struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DefaultVpcDhcpOptionsSpec   `json:"spec"`
	Status            DefaultVpcDhcpOptionsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DefaultVpcDhcpOptionsList contains a list of DefaultVpcDhcpOptionss
type DefaultVpcDhcpOptionsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DefaultVpcDhcpOptions `json:"items"`
}

// Repository type metadata.
var (
	DefaultVpcDhcpOptionsKind             = "DefaultVpcDhcpOptions"
	DefaultVpcDhcpOptionsGroupKind        = schema.GroupKind{Group: Group, Kind: DefaultVpcDhcpOptionsKind}.String()
	DefaultVpcDhcpOptionsKindAPIVersion   = DefaultVpcDhcpOptionsKind + "." + GroupVersion.String()
	DefaultVpcDhcpOptionsGroupVersionKind = GroupVersion.WithKind(DefaultVpcDhcpOptionsKind)
)

func init() {
	SchemeBuilder.Register(&DefaultVpcDhcpOptions{}, &DefaultVpcDhcpOptionsList{})
}
