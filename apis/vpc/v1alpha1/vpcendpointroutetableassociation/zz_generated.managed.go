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
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"

// GetCondition of this VpcEndpointRouteTableAssociation.
func (mg *VpcEndpointRouteTableAssociation) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this VpcEndpointRouteTableAssociation.
func (mg *VpcEndpointRouteTableAssociation) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this VpcEndpointRouteTableAssociation.
func (mg *VpcEndpointRouteTableAssociation) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this VpcEndpointRouteTableAssociation.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *VpcEndpointRouteTableAssociation) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this VpcEndpointRouteTableAssociation.
func (mg *VpcEndpointRouteTableAssociation) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this VpcEndpointRouteTableAssociation.
func (mg *VpcEndpointRouteTableAssociation) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this VpcEndpointRouteTableAssociation.
func (mg *VpcEndpointRouteTableAssociation) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this VpcEndpointRouteTableAssociation.
func (mg *VpcEndpointRouteTableAssociation) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this VpcEndpointRouteTableAssociation.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *VpcEndpointRouteTableAssociation) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this VpcEndpointRouteTableAssociation.
func (mg *VpcEndpointRouteTableAssociation) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}