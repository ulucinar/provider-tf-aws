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

// GetCondition of this Wafv2IpSet.
func (mg *Wafv2IpSet) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this Wafv2IpSet.
func (mg *Wafv2IpSet) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this Wafv2IpSet.
func (mg *Wafv2IpSet) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this Wafv2IpSet.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *Wafv2IpSet) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this Wafv2IpSet.
func (mg *Wafv2IpSet) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this Wafv2IpSet.
func (mg *Wafv2IpSet) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this Wafv2IpSet.
func (mg *Wafv2IpSet) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this Wafv2IpSet.
func (mg *Wafv2IpSet) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this Wafv2IpSet.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *Wafv2IpSet) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this Wafv2IpSet.
func (mg *Wafv2IpSet) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this Wafv2RegexPatternSet.
func (mg *Wafv2RegexPatternSet) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this Wafv2RegexPatternSet.
func (mg *Wafv2RegexPatternSet) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this Wafv2RegexPatternSet.
func (mg *Wafv2RegexPatternSet) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this Wafv2RegexPatternSet.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *Wafv2RegexPatternSet) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this Wafv2RegexPatternSet.
func (mg *Wafv2RegexPatternSet) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this Wafv2RegexPatternSet.
func (mg *Wafv2RegexPatternSet) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this Wafv2RegexPatternSet.
func (mg *Wafv2RegexPatternSet) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this Wafv2RegexPatternSet.
func (mg *Wafv2RegexPatternSet) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this Wafv2RegexPatternSet.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *Wafv2RegexPatternSet) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this Wafv2RegexPatternSet.
func (mg *Wafv2RegexPatternSet) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this Wafv2RuleGroup.
func (mg *Wafv2RuleGroup) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this Wafv2RuleGroup.
func (mg *Wafv2RuleGroup) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this Wafv2RuleGroup.
func (mg *Wafv2RuleGroup) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this Wafv2RuleGroup.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *Wafv2RuleGroup) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this Wafv2RuleGroup.
func (mg *Wafv2RuleGroup) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this Wafv2RuleGroup.
func (mg *Wafv2RuleGroup) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this Wafv2RuleGroup.
func (mg *Wafv2RuleGroup) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this Wafv2RuleGroup.
func (mg *Wafv2RuleGroup) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this Wafv2RuleGroup.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *Wafv2RuleGroup) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this Wafv2RuleGroup.
func (mg *Wafv2RuleGroup) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this Wafv2WebAcl.
func (mg *Wafv2WebAcl) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this Wafv2WebAcl.
func (mg *Wafv2WebAcl) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this Wafv2WebAcl.
func (mg *Wafv2WebAcl) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this Wafv2WebAcl.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *Wafv2WebAcl) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this Wafv2WebAcl.
func (mg *Wafv2WebAcl) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this Wafv2WebAcl.
func (mg *Wafv2WebAcl) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this Wafv2WebAcl.
func (mg *Wafv2WebAcl) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this Wafv2WebAcl.
func (mg *Wafv2WebAcl) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this Wafv2WebAcl.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *Wafv2WebAcl) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this Wafv2WebAcl.
func (mg *Wafv2WebAcl) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this Wafv2WebAclAssociation.
func (mg *Wafv2WebAclAssociation) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this Wafv2WebAclAssociation.
func (mg *Wafv2WebAclAssociation) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this Wafv2WebAclAssociation.
func (mg *Wafv2WebAclAssociation) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this Wafv2WebAclAssociation.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *Wafv2WebAclAssociation) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this Wafv2WebAclAssociation.
func (mg *Wafv2WebAclAssociation) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this Wafv2WebAclAssociation.
func (mg *Wafv2WebAclAssociation) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this Wafv2WebAclAssociation.
func (mg *Wafv2WebAclAssociation) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this Wafv2WebAclAssociation.
func (mg *Wafv2WebAclAssociation) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this Wafv2WebAclAssociation.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *Wafv2WebAclAssociation) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this Wafv2WebAclAssociation.
func (mg *Wafv2WebAclAssociation) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this Wafv2WebAclLoggingConfiguration.
func (mg *Wafv2WebAclLoggingConfiguration) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this Wafv2WebAclLoggingConfiguration.
func (mg *Wafv2WebAclLoggingConfiguration) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this Wafv2WebAclLoggingConfiguration.
func (mg *Wafv2WebAclLoggingConfiguration) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this Wafv2WebAclLoggingConfiguration.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *Wafv2WebAclLoggingConfiguration) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this Wafv2WebAclLoggingConfiguration.
func (mg *Wafv2WebAclLoggingConfiguration) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this Wafv2WebAclLoggingConfiguration.
func (mg *Wafv2WebAclLoggingConfiguration) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this Wafv2WebAclLoggingConfiguration.
func (mg *Wafv2WebAclLoggingConfiguration) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this Wafv2WebAclLoggingConfiguration.
func (mg *Wafv2WebAclLoggingConfiguration) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this Wafv2WebAclLoggingConfiguration.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *Wafv2WebAclLoggingConfiguration) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this Wafv2WebAclLoggingConfiguration.
func (mg *Wafv2WebAclLoggingConfiguration) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}