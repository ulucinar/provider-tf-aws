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
func (in *AccessLogSettingsObservation) DeepCopyInto(out *AccessLogSettingsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessLogSettingsObservation.
func (in *AccessLogSettingsObservation) DeepCopy() *AccessLogSettingsObservation {
	if in == nil {
		return nil
	}
	out := new(AccessLogSettingsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessLogSettingsParameters) DeepCopyInto(out *AccessLogSettingsParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessLogSettingsParameters.
func (in *AccessLogSettingsParameters) DeepCopy() *AccessLogSettingsParameters {
	if in == nil {
		return nil
	}
	out := new(AccessLogSettingsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Apigatewayv2Stage) DeepCopyInto(out *Apigatewayv2Stage) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Apigatewayv2Stage.
func (in *Apigatewayv2Stage) DeepCopy() *Apigatewayv2Stage {
	if in == nil {
		return nil
	}
	out := new(Apigatewayv2Stage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Apigatewayv2Stage) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Apigatewayv2StageList) DeepCopyInto(out *Apigatewayv2StageList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Apigatewayv2Stage, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Apigatewayv2StageList.
func (in *Apigatewayv2StageList) DeepCopy() *Apigatewayv2StageList {
	if in == nil {
		return nil
	}
	out := new(Apigatewayv2StageList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Apigatewayv2StageList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Apigatewayv2StageObservation) DeepCopyInto(out *Apigatewayv2StageObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Apigatewayv2StageObservation.
func (in *Apigatewayv2StageObservation) DeepCopy() *Apigatewayv2StageObservation {
	if in == nil {
		return nil
	}
	out := new(Apigatewayv2StageObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Apigatewayv2StageParameters) DeepCopyInto(out *Apigatewayv2StageParameters) {
	*out = *in
	if in.AccessLogSettings != nil {
		in, out := &in.AccessLogSettings, &out.AccessLogSettings
		*out = make([]AccessLogSettingsParameters, len(*in))
		copy(*out, *in)
	}
	if in.AutoDeploy != nil {
		in, out := &in.AutoDeploy, &out.AutoDeploy
		*out = new(bool)
		**out = **in
	}
	if in.ClientCertificateId != nil {
		in, out := &in.ClientCertificateId, &out.ClientCertificateId
		*out = new(string)
		**out = **in
	}
	if in.DefaultRouteSettings != nil {
		in, out := &in.DefaultRouteSettings, &out.DefaultRouteSettings
		*out = make([]DefaultRouteSettingsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DeploymentId != nil {
		in, out := &in.DeploymentId, &out.DeploymentId
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.RouteSettings != nil {
		in, out := &in.RouteSettings, &out.RouteSettings
		*out = make([]RouteSettingsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.StageVariables != nil {
		in, out := &in.StageVariables, &out.StageVariables
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
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
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Apigatewayv2StageParameters.
func (in *Apigatewayv2StageParameters) DeepCopy() *Apigatewayv2StageParameters {
	if in == nil {
		return nil
	}
	out := new(Apigatewayv2StageParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Apigatewayv2StageSpec) DeepCopyInto(out *Apigatewayv2StageSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Apigatewayv2StageSpec.
func (in *Apigatewayv2StageSpec) DeepCopy() *Apigatewayv2StageSpec {
	if in == nil {
		return nil
	}
	out := new(Apigatewayv2StageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Apigatewayv2StageStatus) DeepCopyInto(out *Apigatewayv2StageStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Apigatewayv2StageStatus.
func (in *Apigatewayv2StageStatus) DeepCopy() *Apigatewayv2StageStatus {
	if in == nil {
		return nil
	}
	out := new(Apigatewayv2StageStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefaultRouteSettingsObservation) DeepCopyInto(out *DefaultRouteSettingsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefaultRouteSettingsObservation.
func (in *DefaultRouteSettingsObservation) DeepCopy() *DefaultRouteSettingsObservation {
	if in == nil {
		return nil
	}
	out := new(DefaultRouteSettingsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefaultRouteSettingsParameters) DeepCopyInto(out *DefaultRouteSettingsParameters) {
	*out = *in
	if in.DataTraceEnabled != nil {
		in, out := &in.DataTraceEnabled, &out.DataTraceEnabled
		*out = new(bool)
		**out = **in
	}
	if in.DetailedMetricsEnabled != nil {
		in, out := &in.DetailedMetricsEnabled, &out.DetailedMetricsEnabled
		*out = new(bool)
		**out = **in
	}
	if in.LoggingLevel != nil {
		in, out := &in.LoggingLevel, &out.LoggingLevel
		*out = new(string)
		**out = **in
	}
	if in.ThrottlingBurstLimit != nil {
		in, out := &in.ThrottlingBurstLimit, &out.ThrottlingBurstLimit
		*out = new(int64)
		**out = **in
	}
	if in.ThrottlingRateLimit != nil {
		in, out := &in.ThrottlingRateLimit, &out.ThrottlingRateLimit
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefaultRouteSettingsParameters.
func (in *DefaultRouteSettingsParameters) DeepCopy() *DefaultRouteSettingsParameters {
	if in == nil {
		return nil
	}
	out := new(DefaultRouteSettingsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteSettingsObservation) DeepCopyInto(out *RouteSettingsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteSettingsObservation.
func (in *RouteSettingsObservation) DeepCopy() *RouteSettingsObservation {
	if in == nil {
		return nil
	}
	out := new(RouteSettingsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteSettingsParameters) DeepCopyInto(out *RouteSettingsParameters) {
	*out = *in
	if in.DataTraceEnabled != nil {
		in, out := &in.DataTraceEnabled, &out.DataTraceEnabled
		*out = new(bool)
		**out = **in
	}
	if in.DetailedMetricsEnabled != nil {
		in, out := &in.DetailedMetricsEnabled, &out.DetailedMetricsEnabled
		*out = new(bool)
		**out = **in
	}
	if in.LoggingLevel != nil {
		in, out := &in.LoggingLevel, &out.LoggingLevel
		*out = new(string)
		**out = **in
	}
	if in.ThrottlingBurstLimit != nil {
		in, out := &in.ThrottlingBurstLimit, &out.ThrottlingBurstLimit
		*out = new(int64)
		**out = **in
	}
	if in.ThrottlingRateLimit != nil {
		in, out := &in.ThrottlingRateLimit, &out.ThrottlingRateLimit
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteSettingsParameters.
func (in *RouteSettingsParameters) DeepCopy() *RouteSettingsParameters {
	if in == nil {
		return nil
	}
	out := new(RouteSettingsParameters)
	in.DeepCopyInto(out)
	return out
}