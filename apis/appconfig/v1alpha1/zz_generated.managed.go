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

// GetCondition of this AppconfigApplication.
func (mg *AppconfigApplication) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this AppconfigApplication.
func (mg *AppconfigApplication) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this AppconfigApplication.
func (mg *AppconfigApplication) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this AppconfigApplication.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *AppconfigApplication) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this AppconfigApplication.
func (mg *AppconfigApplication) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this AppconfigApplication.
func (mg *AppconfigApplication) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this AppconfigApplication.
func (mg *AppconfigApplication) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this AppconfigApplication.
func (mg *AppconfigApplication) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this AppconfigApplication.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *AppconfigApplication) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this AppconfigApplication.
func (mg *AppconfigApplication) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this AppconfigConfigurationProfile.
func (mg *AppconfigConfigurationProfile) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this AppconfigConfigurationProfile.
func (mg *AppconfigConfigurationProfile) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this AppconfigConfigurationProfile.
func (mg *AppconfigConfigurationProfile) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this AppconfigConfigurationProfile.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *AppconfigConfigurationProfile) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this AppconfigConfigurationProfile.
func (mg *AppconfigConfigurationProfile) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this AppconfigConfigurationProfile.
func (mg *AppconfigConfigurationProfile) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this AppconfigConfigurationProfile.
func (mg *AppconfigConfigurationProfile) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this AppconfigConfigurationProfile.
func (mg *AppconfigConfigurationProfile) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this AppconfigConfigurationProfile.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *AppconfigConfigurationProfile) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this AppconfigConfigurationProfile.
func (mg *AppconfigConfigurationProfile) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this AppconfigDeployment.
func (mg *AppconfigDeployment) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this AppconfigDeployment.
func (mg *AppconfigDeployment) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this AppconfigDeployment.
func (mg *AppconfigDeployment) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this AppconfigDeployment.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *AppconfigDeployment) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this AppconfigDeployment.
func (mg *AppconfigDeployment) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this AppconfigDeployment.
func (mg *AppconfigDeployment) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this AppconfigDeployment.
func (mg *AppconfigDeployment) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this AppconfigDeployment.
func (mg *AppconfigDeployment) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this AppconfigDeployment.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *AppconfigDeployment) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this AppconfigDeployment.
func (mg *AppconfigDeployment) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this AppconfigDeploymentStrategy.
func (mg *AppconfigDeploymentStrategy) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this AppconfigDeploymentStrategy.
func (mg *AppconfigDeploymentStrategy) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this AppconfigDeploymentStrategy.
func (mg *AppconfigDeploymentStrategy) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this AppconfigDeploymentStrategy.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *AppconfigDeploymentStrategy) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this AppconfigDeploymentStrategy.
func (mg *AppconfigDeploymentStrategy) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this AppconfigDeploymentStrategy.
func (mg *AppconfigDeploymentStrategy) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this AppconfigDeploymentStrategy.
func (mg *AppconfigDeploymentStrategy) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this AppconfigDeploymentStrategy.
func (mg *AppconfigDeploymentStrategy) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this AppconfigDeploymentStrategy.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *AppconfigDeploymentStrategy) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this AppconfigDeploymentStrategy.
func (mg *AppconfigDeploymentStrategy) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this AppconfigEnvironment.
func (mg *AppconfigEnvironment) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this AppconfigEnvironment.
func (mg *AppconfigEnvironment) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this AppconfigEnvironment.
func (mg *AppconfigEnvironment) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this AppconfigEnvironment.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *AppconfigEnvironment) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this AppconfigEnvironment.
func (mg *AppconfigEnvironment) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this AppconfigEnvironment.
func (mg *AppconfigEnvironment) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this AppconfigEnvironment.
func (mg *AppconfigEnvironment) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this AppconfigEnvironment.
func (mg *AppconfigEnvironment) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this AppconfigEnvironment.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *AppconfigEnvironment) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this AppconfigEnvironment.
func (mg *AppconfigEnvironment) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this AppconfigHostedConfigurationVersion.
func (mg *AppconfigHostedConfigurationVersion) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this AppconfigHostedConfigurationVersion.
func (mg *AppconfigHostedConfigurationVersion) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this AppconfigHostedConfigurationVersion.
func (mg *AppconfigHostedConfigurationVersion) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this AppconfigHostedConfigurationVersion.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *AppconfigHostedConfigurationVersion) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this AppconfigHostedConfigurationVersion.
func (mg *AppconfigHostedConfigurationVersion) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this AppconfigHostedConfigurationVersion.
func (mg *AppconfigHostedConfigurationVersion) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this AppconfigHostedConfigurationVersion.
func (mg *AppconfigHostedConfigurationVersion) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this AppconfigHostedConfigurationVersion.
func (mg *AppconfigHostedConfigurationVersion) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this AppconfigHostedConfigurationVersion.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *AppconfigHostedConfigurationVersion) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this AppconfigHostedConfigurationVersion.
func (mg *AppconfigHostedConfigurationVersion) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}