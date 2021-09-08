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

// GetCondition of this NetworkfirewallFirewall.
func (mg *NetworkfirewallFirewall) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this NetworkfirewallFirewall.
func (mg *NetworkfirewallFirewall) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this NetworkfirewallFirewall.
func (mg *NetworkfirewallFirewall) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this NetworkfirewallFirewall.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *NetworkfirewallFirewall) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this NetworkfirewallFirewall.
func (mg *NetworkfirewallFirewall) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this NetworkfirewallFirewall.
func (mg *NetworkfirewallFirewall) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this NetworkfirewallFirewall.
func (mg *NetworkfirewallFirewall) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this NetworkfirewallFirewall.
func (mg *NetworkfirewallFirewall) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this NetworkfirewallFirewall.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *NetworkfirewallFirewall) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this NetworkfirewallFirewall.
func (mg *NetworkfirewallFirewall) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this NetworkfirewallFirewallPolicy.
func (mg *NetworkfirewallFirewallPolicy) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this NetworkfirewallFirewallPolicy.
func (mg *NetworkfirewallFirewallPolicy) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this NetworkfirewallFirewallPolicy.
func (mg *NetworkfirewallFirewallPolicy) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this NetworkfirewallFirewallPolicy.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *NetworkfirewallFirewallPolicy) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this NetworkfirewallFirewallPolicy.
func (mg *NetworkfirewallFirewallPolicy) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this NetworkfirewallFirewallPolicy.
func (mg *NetworkfirewallFirewallPolicy) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this NetworkfirewallFirewallPolicy.
func (mg *NetworkfirewallFirewallPolicy) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this NetworkfirewallFirewallPolicy.
func (mg *NetworkfirewallFirewallPolicy) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this NetworkfirewallFirewallPolicy.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *NetworkfirewallFirewallPolicy) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this NetworkfirewallFirewallPolicy.
func (mg *NetworkfirewallFirewallPolicy) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this NetworkfirewallLoggingConfiguration.
func (mg *NetworkfirewallLoggingConfiguration) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this NetworkfirewallLoggingConfiguration.
func (mg *NetworkfirewallLoggingConfiguration) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this NetworkfirewallLoggingConfiguration.
func (mg *NetworkfirewallLoggingConfiguration) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this NetworkfirewallLoggingConfiguration.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *NetworkfirewallLoggingConfiguration) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this NetworkfirewallLoggingConfiguration.
func (mg *NetworkfirewallLoggingConfiguration) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this NetworkfirewallLoggingConfiguration.
func (mg *NetworkfirewallLoggingConfiguration) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this NetworkfirewallLoggingConfiguration.
func (mg *NetworkfirewallLoggingConfiguration) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this NetworkfirewallLoggingConfiguration.
func (mg *NetworkfirewallLoggingConfiguration) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this NetworkfirewallLoggingConfiguration.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *NetworkfirewallLoggingConfiguration) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this NetworkfirewallLoggingConfiguration.
func (mg *NetworkfirewallLoggingConfiguration) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this NetworkfirewallResourcePolicy.
func (mg *NetworkfirewallResourcePolicy) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this NetworkfirewallResourcePolicy.
func (mg *NetworkfirewallResourcePolicy) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this NetworkfirewallResourcePolicy.
func (mg *NetworkfirewallResourcePolicy) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this NetworkfirewallResourcePolicy.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *NetworkfirewallResourcePolicy) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this NetworkfirewallResourcePolicy.
func (mg *NetworkfirewallResourcePolicy) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this NetworkfirewallResourcePolicy.
func (mg *NetworkfirewallResourcePolicy) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this NetworkfirewallResourcePolicy.
func (mg *NetworkfirewallResourcePolicy) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this NetworkfirewallResourcePolicy.
func (mg *NetworkfirewallResourcePolicy) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this NetworkfirewallResourcePolicy.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *NetworkfirewallResourcePolicy) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this NetworkfirewallResourcePolicy.
func (mg *NetworkfirewallResourcePolicy) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this NetworkfirewallRuleGroup.
func (mg *NetworkfirewallRuleGroup) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this NetworkfirewallRuleGroup.
func (mg *NetworkfirewallRuleGroup) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this NetworkfirewallRuleGroup.
func (mg *NetworkfirewallRuleGroup) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this NetworkfirewallRuleGroup.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *NetworkfirewallRuleGroup) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this NetworkfirewallRuleGroup.
func (mg *NetworkfirewallRuleGroup) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this NetworkfirewallRuleGroup.
func (mg *NetworkfirewallRuleGroup) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this NetworkfirewallRuleGroup.
func (mg *NetworkfirewallRuleGroup) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this NetworkfirewallRuleGroup.
func (mg *NetworkfirewallRuleGroup) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this NetworkfirewallRuleGroup.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *NetworkfirewallRuleGroup) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this NetworkfirewallRuleGroup.
func (mg *NetworkfirewallRuleGroup) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}