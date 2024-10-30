//go:build !ignore_autogenerated

// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by controller-gen. DO NOT EDIT.

package v1beta2

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdentityInitParameters) DeepCopyInto(out *IdentityInitParameters) {
	*out = *in
	if in.IdentityIds != nil {
		in, out := &in.IdentityIds, &out.IdentityIds
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdentityInitParameters.
func (in *IdentityInitParameters) DeepCopy() *IdentityInitParameters {
	if in == nil {
		return nil
	}
	out := new(IdentityInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdentityObservation) DeepCopyInto(out *IdentityObservation) {
	*out = *in
	if in.IdentityIds != nil {
		in, out := &in.IdentityIds, &out.IdentityIds
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.PrincipalID != nil {
		in, out := &in.PrincipalID, &out.PrincipalID
		*out = new(string)
		**out = **in
	}
	if in.TenantID != nil {
		in, out := &in.TenantID, &out.TenantID
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdentityObservation.
func (in *IdentityObservation) DeepCopy() *IdentityObservation {
	if in == nil {
		return nil
	}
	out := new(IdentityObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdentityParameters) DeepCopyInto(out *IdentityParameters) {
	*out = *in
	if in.IdentityIds != nil {
		in, out := &in.IdentityIds, &out.IdentityIds
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdentityParameters.
func (in *IdentityParameters) DeepCopy() *IdentityParameters {
	if in == nil {
		return nil
	}
	out := new(IdentityParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PatchScheduleInitParameters) DeepCopyInto(out *PatchScheduleInitParameters) {
	*out = *in
	if in.DayOfWeek != nil {
		in, out := &in.DayOfWeek, &out.DayOfWeek
		*out = new(string)
		**out = **in
	}
	if in.MaintenanceWindow != nil {
		in, out := &in.MaintenanceWindow, &out.MaintenanceWindow
		*out = new(string)
		**out = **in
	}
	if in.StartHourUtc != nil {
		in, out := &in.StartHourUtc, &out.StartHourUtc
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PatchScheduleInitParameters.
func (in *PatchScheduleInitParameters) DeepCopy() *PatchScheduleInitParameters {
	if in == nil {
		return nil
	}
	out := new(PatchScheduleInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PatchScheduleObservation) DeepCopyInto(out *PatchScheduleObservation) {
	*out = *in
	if in.DayOfWeek != nil {
		in, out := &in.DayOfWeek, &out.DayOfWeek
		*out = new(string)
		**out = **in
	}
	if in.MaintenanceWindow != nil {
		in, out := &in.MaintenanceWindow, &out.MaintenanceWindow
		*out = new(string)
		**out = **in
	}
	if in.StartHourUtc != nil {
		in, out := &in.StartHourUtc, &out.StartHourUtc
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PatchScheduleObservation.
func (in *PatchScheduleObservation) DeepCopy() *PatchScheduleObservation {
	if in == nil {
		return nil
	}
	out := new(PatchScheduleObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PatchScheduleParameters) DeepCopyInto(out *PatchScheduleParameters) {
	*out = *in
	if in.DayOfWeek != nil {
		in, out := &in.DayOfWeek, &out.DayOfWeek
		*out = new(string)
		**out = **in
	}
	if in.MaintenanceWindow != nil {
		in, out := &in.MaintenanceWindow, &out.MaintenanceWindow
		*out = new(string)
		**out = **in
	}
	if in.StartHourUtc != nil {
		in, out := &in.StartHourUtc, &out.StartHourUtc
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PatchScheduleParameters.
func (in *PatchScheduleParameters) DeepCopy() *PatchScheduleParameters {
	if in == nil {
		return nil
	}
	out := new(PatchScheduleParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisCache) DeepCopyInto(out *RedisCache) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisCache.
func (in *RedisCache) DeepCopy() *RedisCache {
	if in == nil {
		return nil
	}
	out := new(RedisCache)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RedisCache) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisCacheInitParameters) DeepCopyInto(out *RedisCacheInitParameters) {
	*out = *in
	if in.Capacity != nil {
		in, out := &in.Capacity, &out.Capacity
		*out = new(float64)
		**out = **in
	}
	if in.EnableNonSSLPort != nil {
		in, out := &in.EnableNonSSLPort, &out.EnableNonSSLPort
		*out = new(bool)
		**out = **in
	}
	if in.Family != nil {
		in, out := &in.Family, &out.Family
		*out = new(string)
		**out = **in
	}
	if in.Identity != nil {
		in, out := &in.Identity, &out.Identity
		*out = new(IdentityInitParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.MinimumTLSVersion != nil {
		in, out := &in.MinimumTLSVersion, &out.MinimumTLSVersion
		*out = new(string)
		**out = **in
	}
	if in.NonSSLPortEnabled != nil {
		in, out := &in.NonSSLPortEnabled, &out.NonSSLPortEnabled
		*out = new(bool)
		**out = **in
	}
	if in.PatchSchedule != nil {
		in, out := &in.PatchSchedule, &out.PatchSchedule
		*out = make([]PatchScheduleInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PrivateStaticIPAddress != nil {
		in, out := &in.PrivateStaticIPAddress, &out.PrivateStaticIPAddress
		*out = new(string)
		**out = **in
	}
	if in.PublicNetworkAccessEnabled != nil {
		in, out := &in.PublicNetworkAccessEnabled, &out.PublicNetworkAccessEnabled
		*out = new(bool)
		**out = **in
	}
	if in.RedisConfiguration != nil {
		in, out := &in.RedisConfiguration, &out.RedisConfiguration
		*out = new(RedisConfigurationInitParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.RedisVersion != nil {
		in, out := &in.RedisVersion, &out.RedisVersion
		*out = new(string)
		**out = **in
	}
	if in.ReplicasPerMaster != nil {
		in, out := &in.ReplicasPerMaster, &out.ReplicasPerMaster
		*out = new(float64)
		**out = **in
	}
	if in.ReplicasPerPrimary != nil {
		in, out := &in.ReplicasPerPrimary, &out.ReplicasPerPrimary
		*out = new(float64)
		**out = **in
	}
	if in.ShardCount != nil {
		in, out := &in.ShardCount, &out.ShardCount
		*out = new(float64)
		**out = **in
	}
	if in.SkuName != nil {
		in, out := &in.SkuName, &out.SkuName
		*out = new(string)
		**out = **in
	}
	if in.SubnetID != nil {
		in, out := &in.SubnetID, &out.SubnetID
		*out = new(string)
		**out = **in
	}
	if in.SubnetIDRef != nil {
		in, out := &in.SubnetIDRef, &out.SubnetIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.SubnetIDSelector != nil {
		in, out := &in.SubnetIDSelector, &out.SubnetIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.TenantSettings != nil {
		in, out := &in.TenantSettings, &out.TenantSettings
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Zones != nil {
		in, out := &in.Zones, &out.Zones
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisCacheInitParameters.
func (in *RedisCacheInitParameters) DeepCopy() *RedisCacheInitParameters {
	if in == nil {
		return nil
	}
	out := new(RedisCacheInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisCacheList) DeepCopyInto(out *RedisCacheList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RedisCache, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisCacheList.
func (in *RedisCacheList) DeepCopy() *RedisCacheList {
	if in == nil {
		return nil
	}
	out := new(RedisCacheList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RedisCacheList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisCacheObservation) DeepCopyInto(out *RedisCacheObservation) {
	*out = *in
	if in.Capacity != nil {
		in, out := &in.Capacity, &out.Capacity
		*out = new(float64)
		**out = **in
	}
	if in.EnableNonSSLPort != nil {
		in, out := &in.EnableNonSSLPort, &out.EnableNonSSLPort
		*out = new(bool)
		**out = **in
	}
	if in.Family != nil {
		in, out := &in.Family, &out.Family
		*out = new(string)
		**out = **in
	}
	if in.HostName != nil {
		in, out := &in.HostName, &out.HostName
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Identity != nil {
		in, out := &in.Identity, &out.Identity
		*out = new(IdentityObservation)
		(*in).DeepCopyInto(*out)
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.MinimumTLSVersion != nil {
		in, out := &in.MinimumTLSVersion, &out.MinimumTLSVersion
		*out = new(string)
		**out = **in
	}
	if in.NonSSLPortEnabled != nil {
		in, out := &in.NonSSLPortEnabled, &out.NonSSLPortEnabled
		*out = new(bool)
		**out = **in
	}
	if in.PatchSchedule != nil {
		in, out := &in.PatchSchedule, &out.PatchSchedule
		*out = make([]PatchScheduleObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(float64)
		**out = **in
	}
	if in.PrivateStaticIPAddress != nil {
		in, out := &in.PrivateStaticIPAddress, &out.PrivateStaticIPAddress
		*out = new(string)
		**out = **in
	}
	if in.PublicNetworkAccessEnabled != nil {
		in, out := &in.PublicNetworkAccessEnabled, &out.PublicNetworkAccessEnabled
		*out = new(bool)
		**out = **in
	}
	if in.RedisConfiguration != nil {
		in, out := &in.RedisConfiguration, &out.RedisConfiguration
		*out = new(RedisConfigurationObservation)
		(*in).DeepCopyInto(*out)
	}
	if in.RedisVersion != nil {
		in, out := &in.RedisVersion, &out.RedisVersion
		*out = new(string)
		**out = **in
	}
	if in.ReplicasPerMaster != nil {
		in, out := &in.ReplicasPerMaster, &out.ReplicasPerMaster
		*out = new(float64)
		**out = **in
	}
	if in.ReplicasPerPrimary != nil {
		in, out := &in.ReplicasPerPrimary, &out.ReplicasPerPrimary
		*out = new(float64)
		**out = **in
	}
	if in.ResourceGroupName != nil {
		in, out := &in.ResourceGroupName, &out.ResourceGroupName
		*out = new(string)
		**out = **in
	}
	if in.SSLPort != nil {
		in, out := &in.SSLPort, &out.SSLPort
		*out = new(float64)
		**out = **in
	}
	if in.ShardCount != nil {
		in, out := &in.ShardCount, &out.ShardCount
		*out = new(float64)
		**out = **in
	}
	if in.SkuName != nil {
		in, out := &in.SkuName, &out.SkuName
		*out = new(string)
		**out = **in
	}
	if in.SubnetID != nil {
		in, out := &in.SubnetID, &out.SubnetID
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.TenantSettings != nil {
		in, out := &in.TenantSettings, &out.TenantSettings
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Zones != nil {
		in, out := &in.Zones, &out.Zones
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisCacheObservation.
func (in *RedisCacheObservation) DeepCopy() *RedisCacheObservation {
	if in == nil {
		return nil
	}
	out := new(RedisCacheObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisCacheParameters) DeepCopyInto(out *RedisCacheParameters) {
	*out = *in
	if in.Capacity != nil {
		in, out := &in.Capacity, &out.Capacity
		*out = new(float64)
		**out = **in
	}
	if in.EnableNonSSLPort != nil {
		in, out := &in.EnableNonSSLPort, &out.EnableNonSSLPort
		*out = new(bool)
		**out = **in
	}
	if in.Family != nil {
		in, out := &in.Family, &out.Family
		*out = new(string)
		**out = **in
	}
	if in.Identity != nil {
		in, out := &in.Identity, &out.Identity
		*out = new(IdentityParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.MinimumTLSVersion != nil {
		in, out := &in.MinimumTLSVersion, &out.MinimumTLSVersion
		*out = new(string)
		**out = **in
	}
	if in.NonSSLPortEnabled != nil {
		in, out := &in.NonSSLPortEnabled, &out.NonSSLPortEnabled
		*out = new(bool)
		**out = **in
	}
	if in.PatchSchedule != nil {
		in, out := &in.PatchSchedule, &out.PatchSchedule
		*out = make([]PatchScheduleParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PrivateStaticIPAddress != nil {
		in, out := &in.PrivateStaticIPAddress, &out.PrivateStaticIPAddress
		*out = new(string)
		**out = **in
	}
	if in.PublicNetworkAccessEnabled != nil {
		in, out := &in.PublicNetworkAccessEnabled, &out.PublicNetworkAccessEnabled
		*out = new(bool)
		**out = **in
	}
	if in.RedisConfiguration != nil {
		in, out := &in.RedisConfiguration, &out.RedisConfiguration
		*out = new(RedisConfigurationParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.RedisVersion != nil {
		in, out := &in.RedisVersion, &out.RedisVersion
		*out = new(string)
		**out = **in
	}
	if in.ReplicasPerMaster != nil {
		in, out := &in.ReplicasPerMaster, &out.ReplicasPerMaster
		*out = new(float64)
		**out = **in
	}
	if in.ReplicasPerPrimary != nil {
		in, out := &in.ReplicasPerPrimary, &out.ReplicasPerPrimary
		*out = new(float64)
		**out = **in
	}
	if in.ResourceGroupName != nil {
		in, out := &in.ResourceGroupName, &out.ResourceGroupName
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupNameRef != nil {
		in, out := &in.ResourceGroupNameRef, &out.ResourceGroupNameRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ResourceGroupNameSelector != nil {
		in, out := &in.ResourceGroupNameSelector, &out.ResourceGroupNameSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.ShardCount != nil {
		in, out := &in.ShardCount, &out.ShardCount
		*out = new(float64)
		**out = **in
	}
	if in.SkuName != nil {
		in, out := &in.SkuName, &out.SkuName
		*out = new(string)
		**out = **in
	}
	if in.SubnetID != nil {
		in, out := &in.SubnetID, &out.SubnetID
		*out = new(string)
		**out = **in
	}
	if in.SubnetIDRef != nil {
		in, out := &in.SubnetIDRef, &out.SubnetIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.SubnetIDSelector != nil {
		in, out := &in.SubnetIDSelector, &out.SubnetIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.TenantSettings != nil {
		in, out := &in.TenantSettings, &out.TenantSettings
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Zones != nil {
		in, out := &in.Zones, &out.Zones
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisCacheParameters.
func (in *RedisCacheParameters) DeepCopy() *RedisCacheParameters {
	if in == nil {
		return nil
	}
	out := new(RedisCacheParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisCacheSpec) DeepCopyInto(out *RedisCacheSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisCacheSpec.
func (in *RedisCacheSpec) DeepCopy() *RedisCacheSpec {
	if in == nil {
		return nil
	}
	out := new(RedisCacheSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisCacheStatus) DeepCopyInto(out *RedisCacheStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisCacheStatus.
func (in *RedisCacheStatus) DeepCopy() *RedisCacheStatus {
	if in == nil {
		return nil
	}
	out := new(RedisCacheStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisConfigurationInitParameters) DeepCopyInto(out *RedisConfigurationInitParameters) {
	*out = *in
	if in.ActiveDirectoryAuthenticationEnabled != nil {
		in, out := &in.ActiveDirectoryAuthenticationEnabled, &out.ActiveDirectoryAuthenticationEnabled
		*out = new(bool)
		**out = **in
	}
	if in.AofBackupEnabled != nil {
		in, out := &in.AofBackupEnabled, &out.AofBackupEnabled
		*out = new(bool)
		**out = **in
	}
	if in.AofStorageConnectionString0SecretRef != nil {
		in, out := &in.AofStorageConnectionString0SecretRef, &out.AofStorageConnectionString0SecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.AofStorageConnectionString1SecretRef != nil {
		in, out := &in.AofStorageConnectionString1SecretRef, &out.AofStorageConnectionString1SecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.AuthenticationEnabled != nil {
		in, out := &in.AuthenticationEnabled, &out.AuthenticationEnabled
		*out = new(bool)
		**out = **in
	}
	if in.DataPersistenceAuthenticationMethod != nil {
		in, out := &in.DataPersistenceAuthenticationMethod, &out.DataPersistenceAuthenticationMethod
		*out = new(string)
		**out = **in
	}
	if in.EnableAuthentication != nil {
		in, out := &in.EnableAuthentication, &out.EnableAuthentication
		*out = new(bool)
		**out = **in
	}
	if in.MaxfragmentationmemoryReserved != nil {
		in, out := &in.MaxfragmentationmemoryReserved, &out.MaxfragmentationmemoryReserved
		*out = new(float64)
		**out = **in
	}
	if in.MaxmemoryDelta != nil {
		in, out := &in.MaxmemoryDelta, &out.MaxmemoryDelta
		*out = new(float64)
		**out = **in
	}
	if in.MaxmemoryPolicy != nil {
		in, out := &in.MaxmemoryPolicy, &out.MaxmemoryPolicy
		*out = new(string)
		**out = **in
	}
	if in.MaxmemoryReserved != nil {
		in, out := &in.MaxmemoryReserved, &out.MaxmemoryReserved
		*out = new(float64)
		**out = **in
	}
	if in.NotifyKeySpaceEvents != nil {
		in, out := &in.NotifyKeySpaceEvents, &out.NotifyKeySpaceEvents
		*out = new(string)
		**out = **in
	}
	if in.RdbBackupEnabled != nil {
		in, out := &in.RdbBackupEnabled, &out.RdbBackupEnabled
		*out = new(bool)
		**out = **in
	}
	if in.RdbBackupFrequency != nil {
		in, out := &in.RdbBackupFrequency, &out.RdbBackupFrequency
		*out = new(float64)
		**out = **in
	}
	if in.RdbBackupMaxSnapshotCount != nil {
		in, out := &in.RdbBackupMaxSnapshotCount, &out.RdbBackupMaxSnapshotCount
		*out = new(float64)
		**out = **in
	}
	if in.RdbStorageConnectionStringSecretRef != nil {
		in, out := &in.RdbStorageConnectionStringSecretRef, &out.RdbStorageConnectionStringSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.StorageAccountSubscriptionID != nil {
		in, out := &in.StorageAccountSubscriptionID, &out.StorageAccountSubscriptionID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisConfigurationInitParameters.
func (in *RedisConfigurationInitParameters) DeepCopy() *RedisConfigurationInitParameters {
	if in == nil {
		return nil
	}
	out := new(RedisConfigurationInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisConfigurationObservation) DeepCopyInto(out *RedisConfigurationObservation) {
	*out = *in
	if in.ActiveDirectoryAuthenticationEnabled != nil {
		in, out := &in.ActiveDirectoryAuthenticationEnabled, &out.ActiveDirectoryAuthenticationEnabled
		*out = new(bool)
		**out = **in
	}
	if in.AofBackupEnabled != nil {
		in, out := &in.AofBackupEnabled, &out.AofBackupEnabled
		*out = new(bool)
		**out = **in
	}
	if in.AuthenticationEnabled != nil {
		in, out := &in.AuthenticationEnabled, &out.AuthenticationEnabled
		*out = new(bool)
		**out = **in
	}
	if in.DataPersistenceAuthenticationMethod != nil {
		in, out := &in.DataPersistenceAuthenticationMethod, &out.DataPersistenceAuthenticationMethod
		*out = new(string)
		**out = **in
	}
	if in.EnableAuthentication != nil {
		in, out := &in.EnableAuthentication, &out.EnableAuthentication
		*out = new(bool)
		**out = **in
	}
	if in.Maxclients != nil {
		in, out := &in.Maxclients, &out.Maxclients
		*out = new(float64)
		**out = **in
	}
	if in.MaxfragmentationmemoryReserved != nil {
		in, out := &in.MaxfragmentationmemoryReserved, &out.MaxfragmentationmemoryReserved
		*out = new(float64)
		**out = **in
	}
	if in.MaxmemoryDelta != nil {
		in, out := &in.MaxmemoryDelta, &out.MaxmemoryDelta
		*out = new(float64)
		**out = **in
	}
	if in.MaxmemoryPolicy != nil {
		in, out := &in.MaxmemoryPolicy, &out.MaxmemoryPolicy
		*out = new(string)
		**out = **in
	}
	if in.MaxmemoryReserved != nil {
		in, out := &in.MaxmemoryReserved, &out.MaxmemoryReserved
		*out = new(float64)
		**out = **in
	}
	if in.NotifyKeySpaceEvents != nil {
		in, out := &in.NotifyKeySpaceEvents, &out.NotifyKeySpaceEvents
		*out = new(string)
		**out = **in
	}
	if in.RdbBackupEnabled != nil {
		in, out := &in.RdbBackupEnabled, &out.RdbBackupEnabled
		*out = new(bool)
		**out = **in
	}
	if in.RdbBackupFrequency != nil {
		in, out := &in.RdbBackupFrequency, &out.RdbBackupFrequency
		*out = new(float64)
		**out = **in
	}
	if in.RdbBackupMaxSnapshotCount != nil {
		in, out := &in.RdbBackupMaxSnapshotCount, &out.RdbBackupMaxSnapshotCount
		*out = new(float64)
		**out = **in
	}
	if in.StorageAccountSubscriptionID != nil {
		in, out := &in.StorageAccountSubscriptionID, &out.StorageAccountSubscriptionID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisConfigurationObservation.
func (in *RedisConfigurationObservation) DeepCopy() *RedisConfigurationObservation {
	if in == nil {
		return nil
	}
	out := new(RedisConfigurationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisConfigurationParameters) DeepCopyInto(out *RedisConfigurationParameters) {
	*out = *in
	if in.ActiveDirectoryAuthenticationEnabled != nil {
		in, out := &in.ActiveDirectoryAuthenticationEnabled, &out.ActiveDirectoryAuthenticationEnabled
		*out = new(bool)
		**out = **in
	}
	if in.AofBackupEnabled != nil {
		in, out := &in.AofBackupEnabled, &out.AofBackupEnabled
		*out = new(bool)
		**out = **in
	}
	if in.AofStorageConnectionString0SecretRef != nil {
		in, out := &in.AofStorageConnectionString0SecretRef, &out.AofStorageConnectionString0SecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.AofStorageConnectionString1SecretRef != nil {
		in, out := &in.AofStorageConnectionString1SecretRef, &out.AofStorageConnectionString1SecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.AuthenticationEnabled != nil {
		in, out := &in.AuthenticationEnabled, &out.AuthenticationEnabled
		*out = new(bool)
		**out = **in
	}
	if in.DataPersistenceAuthenticationMethod != nil {
		in, out := &in.DataPersistenceAuthenticationMethod, &out.DataPersistenceAuthenticationMethod
		*out = new(string)
		**out = **in
	}
	if in.EnableAuthentication != nil {
		in, out := &in.EnableAuthentication, &out.EnableAuthentication
		*out = new(bool)
		**out = **in
	}
	if in.MaxfragmentationmemoryReserved != nil {
		in, out := &in.MaxfragmentationmemoryReserved, &out.MaxfragmentationmemoryReserved
		*out = new(float64)
		**out = **in
	}
	if in.MaxmemoryDelta != nil {
		in, out := &in.MaxmemoryDelta, &out.MaxmemoryDelta
		*out = new(float64)
		**out = **in
	}
	if in.MaxmemoryPolicy != nil {
		in, out := &in.MaxmemoryPolicy, &out.MaxmemoryPolicy
		*out = new(string)
		**out = **in
	}
	if in.MaxmemoryReserved != nil {
		in, out := &in.MaxmemoryReserved, &out.MaxmemoryReserved
		*out = new(float64)
		**out = **in
	}
	if in.NotifyKeySpaceEvents != nil {
		in, out := &in.NotifyKeySpaceEvents, &out.NotifyKeySpaceEvents
		*out = new(string)
		**out = **in
	}
	if in.RdbBackupEnabled != nil {
		in, out := &in.RdbBackupEnabled, &out.RdbBackupEnabled
		*out = new(bool)
		**out = **in
	}
	if in.RdbBackupFrequency != nil {
		in, out := &in.RdbBackupFrequency, &out.RdbBackupFrequency
		*out = new(float64)
		**out = **in
	}
	if in.RdbBackupMaxSnapshotCount != nil {
		in, out := &in.RdbBackupMaxSnapshotCount, &out.RdbBackupMaxSnapshotCount
		*out = new(float64)
		**out = **in
	}
	if in.RdbStorageConnectionStringSecretRef != nil {
		in, out := &in.RdbStorageConnectionStringSecretRef, &out.RdbStorageConnectionStringSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.StorageAccountSubscriptionID != nil {
		in, out := &in.StorageAccountSubscriptionID, &out.StorageAccountSubscriptionID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisConfigurationParameters.
func (in *RedisConfigurationParameters) DeepCopy() *RedisConfigurationParameters {
	if in == nil {
		return nil
	}
	out := new(RedisConfigurationParameters)
	in.DeepCopyInto(out)
	return out
}
