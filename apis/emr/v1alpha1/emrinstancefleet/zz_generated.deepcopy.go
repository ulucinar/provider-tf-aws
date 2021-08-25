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
func (in *ConfigurationsObservation) DeepCopyInto(out *ConfigurationsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigurationsObservation.
func (in *ConfigurationsObservation) DeepCopy() *ConfigurationsObservation {
	if in == nil {
		return nil
	}
	out := new(ConfigurationsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigurationsParameters) DeepCopyInto(out *ConfigurationsParameters) {
	*out = *in
	if in.Classification != nil {
		in, out := &in.Classification, &out.Classification
		*out = new(string)
		**out = **in
	}
	if in.Properties != nil {
		in, out := &in.Properties, &out.Properties
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigurationsParameters.
func (in *ConfigurationsParameters) DeepCopy() *ConfigurationsParameters {
	if in == nil {
		return nil
	}
	out := new(ConfigurationsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EbsConfigObservation) DeepCopyInto(out *EbsConfigObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EbsConfigObservation.
func (in *EbsConfigObservation) DeepCopy() *EbsConfigObservation {
	if in == nil {
		return nil
	}
	out := new(EbsConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EbsConfigParameters) DeepCopyInto(out *EbsConfigParameters) {
	*out = *in
	if in.Iops != nil {
		in, out := &in.Iops, &out.Iops
		*out = new(int64)
		**out = **in
	}
	if in.VolumesPerInstance != nil {
		in, out := &in.VolumesPerInstance, &out.VolumesPerInstance
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EbsConfigParameters.
func (in *EbsConfigParameters) DeepCopy() *EbsConfigParameters {
	if in == nil {
		return nil
	}
	out := new(EbsConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EmrInstanceFleet) DeepCopyInto(out *EmrInstanceFleet) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EmrInstanceFleet.
func (in *EmrInstanceFleet) DeepCopy() *EmrInstanceFleet {
	if in == nil {
		return nil
	}
	out := new(EmrInstanceFleet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EmrInstanceFleet) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EmrInstanceFleetList) DeepCopyInto(out *EmrInstanceFleetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EmrInstanceFleet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EmrInstanceFleetList.
func (in *EmrInstanceFleetList) DeepCopy() *EmrInstanceFleetList {
	if in == nil {
		return nil
	}
	out := new(EmrInstanceFleetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EmrInstanceFleetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EmrInstanceFleetObservation) DeepCopyInto(out *EmrInstanceFleetObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EmrInstanceFleetObservation.
func (in *EmrInstanceFleetObservation) DeepCopy() *EmrInstanceFleetObservation {
	if in == nil {
		return nil
	}
	out := new(EmrInstanceFleetObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EmrInstanceFleetParameters) DeepCopyInto(out *EmrInstanceFleetParameters) {
	*out = *in
	if in.InstanceTypeConfigs != nil {
		in, out := &in.InstanceTypeConfigs, &out.InstanceTypeConfigs
		*out = make([]InstanceTypeConfigsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.LaunchSpecifications != nil {
		in, out := &in.LaunchSpecifications, &out.LaunchSpecifications
		*out = make([]LaunchSpecificationsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.TargetOnDemandCapacity != nil {
		in, out := &in.TargetOnDemandCapacity, &out.TargetOnDemandCapacity
		*out = new(int64)
		**out = **in
	}
	if in.TargetSpotCapacity != nil {
		in, out := &in.TargetSpotCapacity, &out.TargetSpotCapacity
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EmrInstanceFleetParameters.
func (in *EmrInstanceFleetParameters) DeepCopy() *EmrInstanceFleetParameters {
	if in == nil {
		return nil
	}
	out := new(EmrInstanceFleetParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EmrInstanceFleetSpec) DeepCopyInto(out *EmrInstanceFleetSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EmrInstanceFleetSpec.
func (in *EmrInstanceFleetSpec) DeepCopy() *EmrInstanceFleetSpec {
	if in == nil {
		return nil
	}
	out := new(EmrInstanceFleetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EmrInstanceFleetStatus) DeepCopyInto(out *EmrInstanceFleetStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EmrInstanceFleetStatus.
func (in *EmrInstanceFleetStatus) DeepCopy() *EmrInstanceFleetStatus {
	if in == nil {
		return nil
	}
	out := new(EmrInstanceFleetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceTypeConfigsObservation) DeepCopyInto(out *InstanceTypeConfigsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceTypeConfigsObservation.
func (in *InstanceTypeConfigsObservation) DeepCopy() *InstanceTypeConfigsObservation {
	if in == nil {
		return nil
	}
	out := new(InstanceTypeConfigsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceTypeConfigsParameters) DeepCopyInto(out *InstanceTypeConfigsParameters) {
	*out = *in
	if in.BidPrice != nil {
		in, out := &in.BidPrice, &out.BidPrice
		*out = new(string)
		**out = **in
	}
	if in.BidPriceAsPercentageOfOnDemandPrice != nil {
		in, out := &in.BidPriceAsPercentageOfOnDemandPrice, &out.BidPriceAsPercentageOfOnDemandPrice
		*out = new(float64)
		**out = **in
	}
	if in.Configurations != nil {
		in, out := &in.Configurations, &out.Configurations
		*out = make([]ConfigurationsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.EbsConfig != nil {
		in, out := &in.EbsConfig, &out.EbsConfig
		*out = make([]EbsConfigParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.WeightedCapacity != nil {
		in, out := &in.WeightedCapacity, &out.WeightedCapacity
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceTypeConfigsParameters.
func (in *InstanceTypeConfigsParameters) DeepCopy() *InstanceTypeConfigsParameters {
	if in == nil {
		return nil
	}
	out := new(InstanceTypeConfigsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LaunchSpecificationsObservation) DeepCopyInto(out *LaunchSpecificationsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LaunchSpecificationsObservation.
func (in *LaunchSpecificationsObservation) DeepCopy() *LaunchSpecificationsObservation {
	if in == nil {
		return nil
	}
	out := new(LaunchSpecificationsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LaunchSpecificationsParameters) DeepCopyInto(out *LaunchSpecificationsParameters) {
	*out = *in
	if in.OnDemandSpecification != nil {
		in, out := &in.OnDemandSpecification, &out.OnDemandSpecification
		*out = make([]OnDemandSpecificationParameters, len(*in))
		copy(*out, *in)
	}
	if in.SpotSpecification != nil {
		in, out := &in.SpotSpecification, &out.SpotSpecification
		*out = make([]SpotSpecificationParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LaunchSpecificationsParameters.
func (in *LaunchSpecificationsParameters) DeepCopy() *LaunchSpecificationsParameters {
	if in == nil {
		return nil
	}
	out := new(LaunchSpecificationsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OnDemandSpecificationObservation) DeepCopyInto(out *OnDemandSpecificationObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OnDemandSpecificationObservation.
func (in *OnDemandSpecificationObservation) DeepCopy() *OnDemandSpecificationObservation {
	if in == nil {
		return nil
	}
	out := new(OnDemandSpecificationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OnDemandSpecificationParameters) DeepCopyInto(out *OnDemandSpecificationParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OnDemandSpecificationParameters.
func (in *OnDemandSpecificationParameters) DeepCopy() *OnDemandSpecificationParameters {
	if in == nil {
		return nil
	}
	out := new(OnDemandSpecificationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpotSpecificationObservation) DeepCopyInto(out *SpotSpecificationObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpotSpecificationObservation.
func (in *SpotSpecificationObservation) DeepCopy() *SpotSpecificationObservation {
	if in == nil {
		return nil
	}
	out := new(SpotSpecificationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpotSpecificationParameters) DeepCopyInto(out *SpotSpecificationParameters) {
	*out = *in
	if in.BlockDurationMinutes != nil {
		in, out := &in.BlockDurationMinutes, &out.BlockDurationMinutes
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpotSpecificationParameters.
func (in *SpotSpecificationParameters) DeepCopy() *SpotSpecificationParameters {
	if in == nil {
		return nil
	}
	out := new(SpotSpecificationParameters)
	in.DeepCopyInto(out)
	return out
}