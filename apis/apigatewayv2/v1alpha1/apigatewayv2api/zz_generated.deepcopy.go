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
func (in *Apigatewayv2Api) DeepCopyInto(out *Apigatewayv2Api) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Apigatewayv2Api.
func (in *Apigatewayv2Api) DeepCopy() *Apigatewayv2Api {
	if in == nil {
		return nil
	}
	out := new(Apigatewayv2Api)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Apigatewayv2Api) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Apigatewayv2ApiList) DeepCopyInto(out *Apigatewayv2ApiList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Apigatewayv2Api, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Apigatewayv2ApiList.
func (in *Apigatewayv2ApiList) DeepCopy() *Apigatewayv2ApiList {
	if in == nil {
		return nil
	}
	out := new(Apigatewayv2ApiList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Apigatewayv2ApiList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Apigatewayv2ApiObservation) DeepCopyInto(out *Apigatewayv2ApiObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Apigatewayv2ApiObservation.
func (in *Apigatewayv2ApiObservation) DeepCopy() *Apigatewayv2ApiObservation {
	if in == nil {
		return nil
	}
	out := new(Apigatewayv2ApiObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Apigatewayv2ApiParameters) DeepCopyInto(out *Apigatewayv2ApiParameters) {
	*out = *in
	if in.ApiKeySelectionExpression != nil {
		in, out := &in.ApiKeySelectionExpression, &out.ApiKeySelectionExpression
		*out = new(string)
		**out = **in
	}
	if in.Body != nil {
		in, out := &in.Body, &out.Body
		*out = new(string)
		**out = **in
	}
	if in.CorsConfiguration != nil {
		in, out := &in.CorsConfiguration, &out.CorsConfiguration
		*out = make([]CorsConfigurationParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CredentialsArn != nil {
		in, out := &in.CredentialsArn, &out.CredentialsArn
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.DisableExecuteApiEndpoint != nil {
		in, out := &in.DisableExecuteApiEndpoint, &out.DisableExecuteApiEndpoint
		*out = new(bool)
		**out = **in
	}
	if in.FailOnWarnings != nil {
		in, out := &in.FailOnWarnings, &out.FailOnWarnings
		*out = new(bool)
		**out = **in
	}
	if in.RouteKey != nil {
		in, out := &in.RouteKey, &out.RouteKey
		*out = new(string)
		**out = **in
	}
	if in.RouteSelectionExpression != nil {
		in, out := &in.RouteSelectionExpression, &out.RouteSelectionExpression
		*out = new(string)
		**out = **in
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
	if in.Target != nil {
		in, out := &in.Target, &out.Target
		*out = new(string)
		**out = **in
	}
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Apigatewayv2ApiParameters.
func (in *Apigatewayv2ApiParameters) DeepCopy() *Apigatewayv2ApiParameters {
	if in == nil {
		return nil
	}
	out := new(Apigatewayv2ApiParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Apigatewayv2ApiSpec) DeepCopyInto(out *Apigatewayv2ApiSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Apigatewayv2ApiSpec.
func (in *Apigatewayv2ApiSpec) DeepCopy() *Apigatewayv2ApiSpec {
	if in == nil {
		return nil
	}
	out := new(Apigatewayv2ApiSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Apigatewayv2ApiStatus) DeepCopyInto(out *Apigatewayv2ApiStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Apigatewayv2ApiStatus.
func (in *Apigatewayv2ApiStatus) DeepCopy() *Apigatewayv2ApiStatus {
	if in == nil {
		return nil
	}
	out := new(Apigatewayv2ApiStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CorsConfigurationObservation) DeepCopyInto(out *CorsConfigurationObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CorsConfigurationObservation.
func (in *CorsConfigurationObservation) DeepCopy() *CorsConfigurationObservation {
	if in == nil {
		return nil
	}
	out := new(CorsConfigurationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CorsConfigurationParameters) DeepCopyInto(out *CorsConfigurationParameters) {
	*out = *in
	if in.AllowCredentials != nil {
		in, out := &in.AllowCredentials, &out.AllowCredentials
		*out = new(bool)
		**out = **in
	}
	if in.AllowHeaders != nil {
		in, out := &in.AllowHeaders, &out.AllowHeaders
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AllowMethods != nil {
		in, out := &in.AllowMethods, &out.AllowMethods
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AllowOrigins != nil {
		in, out := &in.AllowOrigins, &out.AllowOrigins
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ExposeHeaders != nil {
		in, out := &in.ExposeHeaders, &out.ExposeHeaders
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.MaxAge != nil {
		in, out := &in.MaxAge, &out.MaxAge
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CorsConfigurationParameters.
func (in *CorsConfigurationParameters) DeepCopy() *CorsConfigurationParameters {
	if in == nil {
		return nil
	}
	out := new(CorsConfigurationParameters)
	in.DeepCopyInto(out)
	return out
}