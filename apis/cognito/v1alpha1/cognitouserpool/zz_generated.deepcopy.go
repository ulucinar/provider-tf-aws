// +build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccountRecoverySettingObservation) DeepCopyInto(out *AccountRecoverySettingObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccountRecoverySettingObservation.
func (in *AccountRecoverySettingObservation) DeepCopy() *AccountRecoverySettingObservation {
	if in == nil {
		return nil
	}
	out := new(AccountRecoverySettingObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccountRecoverySettingParameters) DeepCopyInto(out *AccountRecoverySettingParameters) {
	*out = *in
	if in.RecoveryMechanism != nil {
		in, out := &in.RecoveryMechanism, &out.RecoveryMechanism
		*out = make([]RecoveryMechanismParameters, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccountRecoverySettingParameters.
func (in *AccountRecoverySettingParameters) DeepCopy() *AccountRecoverySettingParameters {
	if in == nil {
		return nil
	}
	out := new(AccountRecoverySettingParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdminCreateUserConfigObservation) DeepCopyInto(out *AdminCreateUserConfigObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdminCreateUserConfigObservation.
func (in *AdminCreateUserConfigObservation) DeepCopy() *AdminCreateUserConfigObservation {
	if in == nil {
		return nil
	}
	out := new(AdminCreateUserConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdminCreateUserConfigParameters) DeepCopyInto(out *AdminCreateUserConfigParameters) {
	*out = *in
	if in.AllowAdminCreateUserOnly != nil {
		in, out := &in.AllowAdminCreateUserOnly, &out.AllowAdminCreateUserOnly
		*out = new(bool)
		**out = **in
	}
	if in.InviteMessageTemplate != nil {
		in, out := &in.InviteMessageTemplate, &out.InviteMessageTemplate
		*out = make([]InviteMessageTemplateParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdminCreateUserConfigParameters.
func (in *AdminCreateUserConfigParameters) DeepCopy() *AdminCreateUserConfigParameters {
	if in == nil {
		return nil
	}
	out := new(AdminCreateUserConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CognitoUserPool) DeepCopyInto(out *CognitoUserPool) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CognitoUserPool.
func (in *CognitoUserPool) DeepCopy() *CognitoUserPool {
	if in == nil {
		return nil
	}
	out := new(CognitoUserPool)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CognitoUserPool) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CognitoUserPoolList) DeepCopyInto(out *CognitoUserPoolList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CognitoUserPool, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CognitoUserPoolList.
func (in *CognitoUserPoolList) DeepCopy() *CognitoUserPoolList {
	if in == nil {
		return nil
	}
	out := new(CognitoUserPoolList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CognitoUserPoolList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CognitoUserPoolObservation) DeepCopyInto(out *CognitoUserPoolObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CognitoUserPoolObservation.
func (in *CognitoUserPoolObservation) DeepCopy() *CognitoUserPoolObservation {
	if in == nil {
		return nil
	}
	out := new(CognitoUserPoolObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CognitoUserPoolParameters) DeepCopyInto(out *CognitoUserPoolParameters) {
	*out = *in
	if in.AccountRecoverySetting != nil {
		in, out := &in.AccountRecoverySetting, &out.AccountRecoverySetting
		*out = make([]AccountRecoverySettingParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AdminCreateUserConfig != nil {
		in, out := &in.AdminCreateUserConfig, &out.AdminCreateUserConfig
		*out = make([]AdminCreateUserConfigParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AliasAttributes != nil {
		in, out := &in.AliasAttributes, &out.AliasAttributes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AutoVerifiedAttributes != nil {
		in, out := &in.AutoVerifiedAttributes, &out.AutoVerifiedAttributes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.DeviceConfiguration != nil {
		in, out := &in.DeviceConfiguration, &out.DeviceConfiguration
		*out = make([]DeviceConfigurationParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.EmailConfiguration != nil {
		in, out := &in.EmailConfiguration, &out.EmailConfiguration
		*out = make([]EmailConfigurationParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.EmailVerificationMessage != nil {
		in, out := &in.EmailVerificationMessage, &out.EmailVerificationMessage
		*out = new(string)
		**out = **in
	}
	if in.EmailVerificationSubject != nil {
		in, out := &in.EmailVerificationSubject, &out.EmailVerificationSubject
		*out = new(string)
		**out = **in
	}
	if in.LambdaConfig != nil {
		in, out := &in.LambdaConfig, &out.LambdaConfig
		*out = make([]LambdaConfigParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MfaConfiguration != nil {
		in, out := &in.MfaConfiguration, &out.MfaConfiguration
		*out = new(string)
		**out = **in
	}
	if in.PasswordPolicy != nil {
		in, out := &in.PasswordPolicy, &out.PasswordPolicy
		*out = make([]PasswordPolicyParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Schema != nil {
		in, out := &in.Schema, &out.Schema
		*out = make([]SchemaParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SmsAuthenticationMessage != nil {
		in, out := &in.SmsAuthenticationMessage, &out.SmsAuthenticationMessage
		*out = new(string)
		**out = **in
	}
	if in.SmsConfiguration != nil {
		in, out := &in.SmsConfiguration, &out.SmsConfiguration
		*out = make([]SmsConfigurationParameters, len(*in))
		copy(*out, *in)
	}
	if in.SmsVerificationMessage != nil {
		in, out := &in.SmsVerificationMessage, &out.SmsVerificationMessage
		*out = new(string)
		**out = **in
	}
	if in.SoftwareTokenMfaConfiguration != nil {
		in, out := &in.SoftwareTokenMfaConfiguration, &out.SoftwareTokenMfaConfiguration
		*out = make([]SoftwareTokenMfaConfigurationParameters, len(*in))
		copy(*out, *in)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.TagsAll != nil {
		in, out := &in.TagsAll, &out.TagsAll
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.UserPoolAddOns != nil {
		in, out := &in.UserPoolAddOns, &out.UserPoolAddOns
		*out = make([]UserPoolAddOnsParameters, len(*in))
		copy(*out, *in)
	}
	if in.UsernameAttributes != nil {
		in, out := &in.UsernameAttributes, &out.UsernameAttributes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.UsernameConfiguration != nil {
		in, out := &in.UsernameConfiguration, &out.UsernameConfiguration
		*out = make([]UsernameConfigurationParameters, len(*in))
		copy(*out, *in)
	}
	if in.VerificationMessageTemplate != nil {
		in, out := &in.VerificationMessageTemplate, &out.VerificationMessageTemplate
		*out = make([]VerificationMessageTemplateParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CognitoUserPoolParameters.
func (in *CognitoUserPoolParameters) DeepCopy() *CognitoUserPoolParameters {
	if in == nil {
		return nil
	}
	out := new(CognitoUserPoolParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CognitoUserPoolSpec) DeepCopyInto(out *CognitoUserPoolSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CognitoUserPoolSpec.
func (in *CognitoUserPoolSpec) DeepCopy() *CognitoUserPoolSpec {
	if in == nil {
		return nil
	}
	out := new(CognitoUserPoolSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CognitoUserPoolStatus) DeepCopyInto(out *CognitoUserPoolStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CognitoUserPoolStatus.
func (in *CognitoUserPoolStatus) DeepCopy() *CognitoUserPoolStatus {
	if in == nil {
		return nil
	}
	out := new(CognitoUserPoolStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomEmailSenderObservation) DeepCopyInto(out *CustomEmailSenderObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomEmailSenderObservation.
func (in *CustomEmailSenderObservation) DeepCopy() *CustomEmailSenderObservation {
	if in == nil {
		return nil
	}
	out := new(CustomEmailSenderObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomEmailSenderParameters) DeepCopyInto(out *CustomEmailSenderParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomEmailSenderParameters.
func (in *CustomEmailSenderParameters) DeepCopy() *CustomEmailSenderParameters {
	if in == nil {
		return nil
	}
	out := new(CustomEmailSenderParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomSmsSenderObservation) DeepCopyInto(out *CustomSmsSenderObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomSmsSenderObservation.
func (in *CustomSmsSenderObservation) DeepCopy() *CustomSmsSenderObservation {
	if in == nil {
		return nil
	}
	out := new(CustomSmsSenderObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomSmsSenderParameters) DeepCopyInto(out *CustomSmsSenderParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomSmsSenderParameters.
func (in *CustomSmsSenderParameters) DeepCopy() *CustomSmsSenderParameters {
	if in == nil {
		return nil
	}
	out := new(CustomSmsSenderParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeviceConfigurationObservation) DeepCopyInto(out *DeviceConfigurationObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeviceConfigurationObservation.
func (in *DeviceConfigurationObservation) DeepCopy() *DeviceConfigurationObservation {
	if in == nil {
		return nil
	}
	out := new(DeviceConfigurationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeviceConfigurationParameters) DeepCopyInto(out *DeviceConfigurationParameters) {
	*out = *in
	if in.ChallengeRequiredOnNewDevice != nil {
		in, out := &in.ChallengeRequiredOnNewDevice, &out.ChallengeRequiredOnNewDevice
		*out = new(bool)
		**out = **in
	}
	if in.DeviceOnlyRememberedOnUserPrompt != nil {
		in, out := &in.DeviceOnlyRememberedOnUserPrompt, &out.DeviceOnlyRememberedOnUserPrompt
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeviceConfigurationParameters.
func (in *DeviceConfigurationParameters) DeepCopy() *DeviceConfigurationParameters {
	if in == nil {
		return nil
	}
	out := new(DeviceConfigurationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EmailConfigurationObservation) DeepCopyInto(out *EmailConfigurationObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EmailConfigurationObservation.
func (in *EmailConfigurationObservation) DeepCopy() *EmailConfigurationObservation {
	if in == nil {
		return nil
	}
	out := new(EmailConfigurationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EmailConfigurationParameters) DeepCopyInto(out *EmailConfigurationParameters) {
	*out = *in
	if in.ConfigurationSet != nil {
		in, out := &in.ConfigurationSet, &out.ConfigurationSet
		*out = new(string)
		**out = **in
	}
	if in.EmailSendingAccount != nil {
		in, out := &in.EmailSendingAccount, &out.EmailSendingAccount
		*out = new(string)
		**out = **in
	}
	if in.FromEmailAddress != nil {
		in, out := &in.FromEmailAddress, &out.FromEmailAddress
		*out = new(string)
		**out = **in
	}
	if in.ReplyToEmailAddress != nil {
		in, out := &in.ReplyToEmailAddress, &out.ReplyToEmailAddress
		*out = new(string)
		**out = **in
	}
	if in.SourceArn != nil {
		in, out := &in.SourceArn, &out.SourceArn
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EmailConfigurationParameters.
func (in *EmailConfigurationParameters) DeepCopy() *EmailConfigurationParameters {
	if in == nil {
		return nil
	}
	out := new(EmailConfigurationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InviteMessageTemplateObservation) DeepCopyInto(out *InviteMessageTemplateObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InviteMessageTemplateObservation.
func (in *InviteMessageTemplateObservation) DeepCopy() *InviteMessageTemplateObservation {
	if in == nil {
		return nil
	}
	out := new(InviteMessageTemplateObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InviteMessageTemplateParameters) DeepCopyInto(out *InviteMessageTemplateParameters) {
	*out = *in
	if in.EmailMessage != nil {
		in, out := &in.EmailMessage, &out.EmailMessage
		*out = new(string)
		**out = **in
	}
	if in.EmailSubject != nil {
		in, out := &in.EmailSubject, &out.EmailSubject
		*out = new(string)
		**out = **in
	}
	if in.SmsMessage != nil {
		in, out := &in.SmsMessage, &out.SmsMessage
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InviteMessageTemplateParameters.
func (in *InviteMessageTemplateParameters) DeepCopy() *InviteMessageTemplateParameters {
	if in == nil {
		return nil
	}
	out := new(InviteMessageTemplateParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LambdaConfigObservation) DeepCopyInto(out *LambdaConfigObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LambdaConfigObservation.
func (in *LambdaConfigObservation) DeepCopy() *LambdaConfigObservation {
	if in == nil {
		return nil
	}
	out := new(LambdaConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LambdaConfigParameters) DeepCopyInto(out *LambdaConfigParameters) {
	*out = *in
	if in.CreateAuthChallenge != nil {
		in, out := &in.CreateAuthChallenge, &out.CreateAuthChallenge
		*out = new(string)
		**out = **in
	}
	if in.CustomEmailSender != nil {
		in, out := &in.CustomEmailSender, &out.CustomEmailSender
		*out = make([]CustomEmailSenderParameters, len(*in))
		copy(*out, *in)
	}
	if in.CustomMessage != nil {
		in, out := &in.CustomMessage, &out.CustomMessage
		*out = new(string)
		**out = **in
	}
	if in.CustomSmsSender != nil {
		in, out := &in.CustomSmsSender, &out.CustomSmsSender
		*out = make([]CustomSmsSenderParameters, len(*in))
		copy(*out, *in)
	}
	if in.DefineAuthChallenge != nil {
		in, out := &in.DefineAuthChallenge, &out.DefineAuthChallenge
		*out = new(string)
		**out = **in
	}
	if in.KmsKeyId != nil {
		in, out := &in.KmsKeyId, &out.KmsKeyId
		*out = new(string)
		**out = **in
	}
	if in.PostAuthentication != nil {
		in, out := &in.PostAuthentication, &out.PostAuthentication
		*out = new(string)
		**out = **in
	}
	if in.PostConfirmation != nil {
		in, out := &in.PostConfirmation, &out.PostConfirmation
		*out = new(string)
		**out = **in
	}
	if in.PreAuthentication != nil {
		in, out := &in.PreAuthentication, &out.PreAuthentication
		*out = new(string)
		**out = **in
	}
	if in.PreSignUp != nil {
		in, out := &in.PreSignUp, &out.PreSignUp
		*out = new(string)
		**out = **in
	}
	if in.PreTokenGeneration != nil {
		in, out := &in.PreTokenGeneration, &out.PreTokenGeneration
		*out = new(string)
		**out = **in
	}
	if in.UserMigration != nil {
		in, out := &in.UserMigration, &out.UserMigration
		*out = new(string)
		**out = **in
	}
	if in.VerifyAuthChallengeResponse != nil {
		in, out := &in.VerifyAuthChallengeResponse, &out.VerifyAuthChallengeResponse
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LambdaConfigParameters.
func (in *LambdaConfigParameters) DeepCopy() *LambdaConfigParameters {
	if in == nil {
		return nil
	}
	out := new(LambdaConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NumberAttributeConstraintsObservation) DeepCopyInto(out *NumberAttributeConstraintsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NumberAttributeConstraintsObservation.
func (in *NumberAttributeConstraintsObservation) DeepCopy() *NumberAttributeConstraintsObservation {
	if in == nil {
		return nil
	}
	out := new(NumberAttributeConstraintsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NumberAttributeConstraintsParameters) DeepCopyInto(out *NumberAttributeConstraintsParameters) {
	*out = *in
	if in.MaxValue != nil {
		in, out := &in.MaxValue, &out.MaxValue
		*out = new(string)
		**out = **in
	}
	if in.MinValue != nil {
		in, out := &in.MinValue, &out.MinValue
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NumberAttributeConstraintsParameters.
func (in *NumberAttributeConstraintsParameters) DeepCopy() *NumberAttributeConstraintsParameters {
	if in == nil {
		return nil
	}
	out := new(NumberAttributeConstraintsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PasswordPolicyObservation) DeepCopyInto(out *PasswordPolicyObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PasswordPolicyObservation.
func (in *PasswordPolicyObservation) DeepCopy() *PasswordPolicyObservation {
	if in == nil {
		return nil
	}
	out := new(PasswordPolicyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PasswordPolicyParameters) DeepCopyInto(out *PasswordPolicyParameters) {
	*out = *in
	if in.MinimumLength != nil {
		in, out := &in.MinimumLength, &out.MinimumLength
		*out = new(int64)
		**out = **in
	}
	if in.RequireLowercase != nil {
		in, out := &in.RequireLowercase, &out.RequireLowercase
		*out = new(bool)
		**out = **in
	}
	if in.RequireNumbers != nil {
		in, out := &in.RequireNumbers, &out.RequireNumbers
		*out = new(bool)
		**out = **in
	}
	if in.RequireSymbols != nil {
		in, out := &in.RequireSymbols, &out.RequireSymbols
		*out = new(bool)
		**out = **in
	}
	if in.RequireUppercase != nil {
		in, out := &in.RequireUppercase, &out.RequireUppercase
		*out = new(bool)
		**out = **in
	}
	if in.TemporaryPasswordValidityDays != nil {
		in, out := &in.TemporaryPasswordValidityDays, &out.TemporaryPasswordValidityDays
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PasswordPolicyParameters.
func (in *PasswordPolicyParameters) DeepCopy() *PasswordPolicyParameters {
	if in == nil {
		return nil
	}
	out := new(PasswordPolicyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RecoveryMechanismObservation) DeepCopyInto(out *RecoveryMechanismObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RecoveryMechanismObservation.
func (in *RecoveryMechanismObservation) DeepCopy() *RecoveryMechanismObservation {
	if in == nil {
		return nil
	}
	out := new(RecoveryMechanismObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RecoveryMechanismParameters) DeepCopyInto(out *RecoveryMechanismParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RecoveryMechanismParameters.
func (in *RecoveryMechanismParameters) DeepCopy() *RecoveryMechanismParameters {
	if in == nil {
		return nil
	}
	out := new(RecoveryMechanismParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SchemaObservation) DeepCopyInto(out *SchemaObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SchemaObservation.
func (in *SchemaObservation) DeepCopy() *SchemaObservation {
	if in == nil {
		return nil
	}
	out := new(SchemaObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SchemaParameters) DeepCopyInto(out *SchemaParameters) {
	*out = *in
	if in.DeveloperOnlyAttribute != nil {
		in, out := &in.DeveloperOnlyAttribute, &out.DeveloperOnlyAttribute
		*out = new(bool)
		**out = **in
	}
	if in.Mutable != nil {
		in, out := &in.Mutable, &out.Mutable
		*out = new(bool)
		**out = **in
	}
	if in.NumberAttributeConstraints != nil {
		in, out := &in.NumberAttributeConstraints, &out.NumberAttributeConstraints
		*out = make([]NumberAttributeConstraintsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Required != nil {
		in, out := &in.Required, &out.Required
		*out = new(bool)
		**out = **in
	}
	if in.StringAttributeConstraints != nil {
		in, out := &in.StringAttributeConstraints, &out.StringAttributeConstraints
		*out = make([]StringAttributeConstraintsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SchemaParameters.
func (in *SchemaParameters) DeepCopy() *SchemaParameters {
	if in == nil {
		return nil
	}
	out := new(SchemaParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SmsConfigurationObservation) DeepCopyInto(out *SmsConfigurationObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SmsConfigurationObservation.
func (in *SmsConfigurationObservation) DeepCopy() *SmsConfigurationObservation {
	if in == nil {
		return nil
	}
	out := new(SmsConfigurationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SmsConfigurationParameters) DeepCopyInto(out *SmsConfigurationParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SmsConfigurationParameters.
func (in *SmsConfigurationParameters) DeepCopy() *SmsConfigurationParameters {
	if in == nil {
		return nil
	}
	out := new(SmsConfigurationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SoftwareTokenMfaConfigurationObservation) DeepCopyInto(out *SoftwareTokenMfaConfigurationObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SoftwareTokenMfaConfigurationObservation.
func (in *SoftwareTokenMfaConfigurationObservation) DeepCopy() *SoftwareTokenMfaConfigurationObservation {
	if in == nil {
		return nil
	}
	out := new(SoftwareTokenMfaConfigurationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SoftwareTokenMfaConfigurationParameters) DeepCopyInto(out *SoftwareTokenMfaConfigurationParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SoftwareTokenMfaConfigurationParameters.
func (in *SoftwareTokenMfaConfigurationParameters) DeepCopy() *SoftwareTokenMfaConfigurationParameters {
	if in == nil {
		return nil
	}
	out := new(SoftwareTokenMfaConfigurationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StringAttributeConstraintsObservation) DeepCopyInto(out *StringAttributeConstraintsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StringAttributeConstraintsObservation.
func (in *StringAttributeConstraintsObservation) DeepCopy() *StringAttributeConstraintsObservation {
	if in == nil {
		return nil
	}
	out := new(StringAttributeConstraintsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StringAttributeConstraintsParameters) DeepCopyInto(out *StringAttributeConstraintsParameters) {
	*out = *in
	if in.MaxLength != nil {
		in, out := &in.MaxLength, &out.MaxLength
		*out = new(string)
		**out = **in
	}
	if in.MinLength != nil {
		in, out := &in.MinLength, &out.MinLength
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StringAttributeConstraintsParameters.
func (in *StringAttributeConstraintsParameters) DeepCopy() *StringAttributeConstraintsParameters {
	if in == nil {
		return nil
	}
	out := new(StringAttributeConstraintsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserPoolAddOnsObservation) DeepCopyInto(out *UserPoolAddOnsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserPoolAddOnsObservation.
func (in *UserPoolAddOnsObservation) DeepCopy() *UserPoolAddOnsObservation {
	if in == nil {
		return nil
	}
	out := new(UserPoolAddOnsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserPoolAddOnsParameters) DeepCopyInto(out *UserPoolAddOnsParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserPoolAddOnsParameters.
func (in *UserPoolAddOnsParameters) DeepCopy() *UserPoolAddOnsParameters {
	if in == nil {
		return nil
	}
	out := new(UserPoolAddOnsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UsernameConfigurationObservation) DeepCopyInto(out *UsernameConfigurationObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UsernameConfigurationObservation.
func (in *UsernameConfigurationObservation) DeepCopy() *UsernameConfigurationObservation {
	if in == nil {
		return nil
	}
	out := new(UsernameConfigurationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UsernameConfigurationParameters) DeepCopyInto(out *UsernameConfigurationParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UsernameConfigurationParameters.
func (in *UsernameConfigurationParameters) DeepCopy() *UsernameConfigurationParameters {
	if in == nil {
		return nil
	}
	out := new(UsernameConfigurationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VerificationMessageTemplateObservation) DeepCopyInto(out *VerificationMessageTemplateObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VerificationMessageTemplateObservation.
func (in *VerificationMessageTemplateObservation) DeepCopy() *VerificationMessageTemplateObservation {
	if in == nil {
		return nil
	}
	out := new(VerificationMessageTemplateObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VerificationMessageTemplateParameters) DeepCopyInto(out *VerificationMessageTemplateParameters) {
	*out = *in
	if in.DefaultEmailOption != nil {
		in, out := &in.DefaultEmailOption, &out.DefaultEmailOption
		*out = new(string)
		**out = **in
	}
	if in.EmailMessage != nil {
		in, out := &in.EmailMessage, &out.EmailMessage
		*out = new(string)
		**out = **in
	}
	if in.EmailMessageByLink != nil {
		in, out := &in.EmailMessageByLink, &out.EmailMessageByLink
		*out = new(string)
		**out = **in
	}
	if in.EmailSubject != nil {
		in, out := &in.EmailSubject, &out.EmailSubject
		*out = new(string)
		**out = **in
	}
	if in.EmailSubjectByLink != nil {
		in, out := &in.EmailSubjectByLink, &out.EmailSubjectByLink
		*out = new(string)
		**out = **in
	}
	if in.SmsMessage != nil {
		in, out := &in.SmsMessage, &out.SmsMessage
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VerificationMessageTemplateParameters.
func (in *VerificationMessageTemplateParameters) DeepCopy() *VerificationMessageTemplateParameters {
	if in == nil {
		return nil
	}
	out := new(VerificationMessageTemplateParameters)
	in.DeepCopyInto(out)
	return out
}