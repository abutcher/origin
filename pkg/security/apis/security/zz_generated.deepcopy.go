// +build !ignore_autogenerated_openshift

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package security

import (
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	api "k8s.io/kubernetes/pkg/api"
	reflect "reflect"
	unsafe "unsafe"
)

func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
//
// Deprecated: deepcopy registration will go away when static deepcopy is fully implemented.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*AllowedFlexVolume).DeepCopyInto(out.(*AllowedFlexVolume))
			return nil
		}, InType: reflect.TypeOf(new(AllowedFlexVolume))},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*FSGroupStrategyOptions).DeepCopyInto(out.(*FSGroupStrategyOptions))
			return nil
		}, InType: reflect.TypeOf(new(FSGroupStrategyOptions))},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*FSGroupStrategyType).DeepCopyInto(out.(*FSGroupStrategyType))
			return nil
		}, InType: reflect.TypeOf(new(FSGroupStrategyType))},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*FSType).DeepCopyInto(out.(*FSType))
			return nil
		}, InType: reflect.TypeOf(new(FSType))},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*IDRange).DeepCopyInto(out.(*IDRange))
			return nil
		}, InType: reflect.TypeOf(new(IDRange))},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*PodSecurityPolicyReview).DeepCopyInto(out.(*PodSecurityPolicyReview))
			return nil
		}, InType: reflect.TypeOf(new(PodSecurityPolicyReview))},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*PodSecurityPolicyReviewSpec).DeepCopyInto(out.(*PodSecurityPolicyReviewSpec))
			return nil
		}, InType: reflect.TypeOf(new(PodSecurityPolicyReviewSpec))},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*PodSecurityPolicyReviewStatus).DeepCopyInto(out.(*PodSecurityPolicyReviewStatus))
			return nil
		}, InType: reflect.TypeOf(new(PodSecurityPolicyReviewStatus))},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*PodSecurityPolicySelfSubjectReview).DeepCopyInto(out.(*PodSecurityPolicySelfSubjectReview))
			return nil
		}, InType: reflect.TypeOf(new(PodSecurityPolicySelfSubjectReview))},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*PodSecurityPolicySelfSubjectReviewSpec).DeepCopyInto(out.(*PodSecurityPolicySelfSubjectReviewSpec))
			return nil
		}, InType: reflect.TypeOf(new(PodSecurityPolicySelfSubjectReviewSpec))},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*PodSecurityPolicySubjectReview).DeepCopyInto(out.(*PodSecurityPolicySubjectReview))
			return nil
		}, InType: reflect.TypeOf(new(PodSecurityPolicySubjectReview))},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*PodSecurityPolicySubjectReviewSpec).DeepCopyInto(out.(*PodSecurityPolicySubjectReviewSpec))
			return nil
		}, InType: reflect.TypeOf(new(PodSecurityPolicySubjectReviewSpec))},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*PodSecurityPolicySubjectReviewStatus).DeepCopyInto(out.(*PodSecurityPolicySubjectReviewStatus))
			return nil
		}, InType: reflect.TypeOf(new(PodSecurityPolicySubjectReviewStatus))},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*RunAsUserStrategyOptions).DeepCopyInto(out.(*RunAsUserStrategyOptions))
			return nil
		}, InType: reflect.TypeOf(new(RunAsUserStrategyOptions))},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*RunAsUserStrategyType).DeepCopyInto(out.(*RunAsUserStrategyType))
			return nil
		}, InType: reflect.TypeOf(new(RunAsUserStrategyType))},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*SELinuxContextStrategyOptions).DeepCopyInto(out.(*SELinuxContextStrategyOptions))
			return nil
		}, InType: reflect.TypeOf(new(SELinuxContextStrategyOptions))},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*SELinuxContextStrategyType).DeepCopyInto(out.(*SELinuxContextStrategyType))
			return nil
		}, InType: reflect.TypeOf(new(SELinuxContextStrategyType))},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*SecurityContextConstraints).DeepCopyInto(out.(*SecurityContextConstraints))
			return nil
		}, InType: reflect.TypeOf(new(SecurityContextConstraints))},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*SecurityContextConstraintsList).DeepCopyInto(out.(*SecurityContextConstraintsList))
			return nil
		}, InType: reflect.TypeOf(new(SecurityContextConstraintsList))},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ServiceAccountPodSecurityPolicyReviewStatus).DeepCopyInto(out.(*ServiceAccountPodSecurityPolicyReviewStatus))
			return nil
		}, InType: reflect.TypeOf(new(ServiceAccountPodSecurityPolicyReviewStatus))},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*SupplementalGroupsStrategyOptions).DeepCopyInto(out.(*SupplementalGroupsStrategyOptions))
			return nil
		}, InType: reflect.TypeOf(new(SupplementalGroupsStrategyOptions))},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*SupplementalGroupsStrategyType).DeepCopyInto(out.(*SupplementalGroupsStrategyType))
			return nil
		}, InType: reflect.TypeOf(new(SupplementalGroupsStrategyType))},
	)
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AllowedFlexVolume) DeepCopyInto(out *AllowedFlexVolume) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AllowedFlexVolume.
func (in *AllowedFlexVolume) DeepCopy() *AllowedFlexVolume {
	if in == nil {
		return nil
	}
	out := new(AllowedFlexVolume)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FSGroupStrategyOptions) DeepCopyInto(out *FSGroupStrategyOptions) {
	*out = *in
	if in.Ranges != nil {
		in, out := &in.Ranges, &out.Ranges
		*out = make([]IDRange, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FSGroupStrategyOptions.
func (in *FSGroupStrategyOptions) DeepCopy() *FSGroupStrategyOptions {
	if in == nil {
		return nil
	}
	out := new(FSGroupStrategyOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FSGroupStrategyType) DeepCopyInto(out *FSGroupStrategyType) {
	{
		in := (*string)(unsafe.Pointer(in))
		out := (*string)(unsafe.Pointer(out))
		*out = *in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FSGroupStrategyType.
func (in *FSGroupStrategyType) DeepCopy() *FSGroupStrategyType {
	if in == nil {
		return nil
	}
	out := new(FSGroupStrategyType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FSType) DeepCopyInto(out *FSType) {
	{
		in := (*string)(unsafe.Pointer(in))
		out := (*string)(unsafe.Pointer(out))
		*out = *in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FSType.
func (in *FSType) DeepCopy() *FSType {
	if in == nil {
		return nil
	}
	out := new(FSType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IDRange) DeepCopyInto(out *IDRange) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IDRange.
func (in *IDRange) DeepCopy() *IDRange {
	if in == nil {
		return nil
	}
	out := new(IDRange)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodSecurityPolicyReview) DeepCopyInto(out *PodSecurityPolicyReview) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodSecurityPolicyReview.
func (in *PodSecurityPolicyReview) DeepCopy() *PodSecurityPolicyReview {
	if in == nil {
		return nil
	}
	out := new(PodSecurityPolicyReview)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PodSecurityPolicyReview) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodSecurityPolicyReviewSpec) DeepCopyInto(out *PodSecurityPolicyReviewSpec) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
	if in.ServiceAccountNames != nil {
		in, out := &in.ServiceAccountNames, &out.ServiceAccountNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodSecurityPolicyReviewSpec.
func (in *PodSecurityPolicyReviewSpec) DeepCopy() *PodSecurityPolicyReviewSpec {
	if in == nil {
		return nil
	}
	out := new(PodSecurityPolicyReviewSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodSecurityPolicyReviewStatus) DeepCopyInto(out *PodSecurityPolicyReviewStatus) {
	*out = *in
	if in.AllowedServiceAccounts != nil {
		in, out := &in.AllowedServiceAccounts, &out.AllowedServiceAccounts
		*out = make([]ServiceAccountPodSecurityPolicyReviewStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodSecurityPolicyReviewStatus.
func (in *PodSecurityPolicyReviewStatus) DeepCopy() *PodSecurityPolicyReviewStatus {
	if in == nil {
		return nil
	}
	out := new(PodSecurityPolicyReviewStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodSecurityPolicySelfSubjectReview) DeepCopyInto(out *PodSecurityPolicySelfSubjectReview) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodSecurityPolicySelfSubjectReview.
func (in *PodSecurityPolicySelfSubjectReview) DeepCopy() *PodSecurityPolicySelfSubjectReview {
	if in == nil {
		return nil
	}
	out := new(PodSecurityPolicySelfSubjectReview)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PodSecurityPolicySelfSubjectReview) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodSecurityPolicySelfSubjectReviewSpec) DeepCopyInto(out *PodSecurityPolicySelfSubjectReviewSpec) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodSecurityPolicySelfSubjectReviewSpec.
func (in *PodSecurityPolicySelfSubjectReviewSpec) DeepCopy() *PodSecurityPolicySelfSubjectReviewSpec {
	if in == nil {
		return nil
	}
	out := new(PodSecurityPolicySelfSubjectReviewSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodSecurityPolicySubjectReview) DeepCopyInto(out *PodSecurityPolicySubjectReview) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodSecurityPolicySubjectReview.
func (in *PodSecurityPolicySubjectReview) DeepCopy() *PodSecurityPolicySubjectReview {
	if in == nil {
		return nil
	}
	out := new(PodSecurityPolicySubjectReview)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PodSecurityPolicySubjectReview) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodSecurityPolicySubjectReviewSpec) DeepCopyInto(out *PodSecurityPolicySubjectReviewSpec) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
	if in.Groups != nil {
		in, out := &in.Groups, &out.Groups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodSecurityPolicySubjectReviewSpec.
func (in *PodSecurityPolicySubjectReviewSpec) DeepCopy() *PodSecurityPolicySubjectReviewSpec {
	if in == nil {
		return nil
	}
	out := new(PodSecurityPolicySubjectReviewSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodSecurityPolicySubjectReviewStatus) DeepCopyInto(out *PodSecurityPolicySubjectReviewStatus) {
	*out = *in
	if in.AllowedBy != nil {
		in, out := &in.AllowedBy, &out.AllowedBy
		if *in == nil {
			*out = nil
		} else {
			*out = new(api.ObjectReference)
			**out = **in
		}
	}
	in.Template.DeepCopyInto(&out.Template)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodSecurityPolicySubjectReviewStatus.
func (in *PodSecurityPolicySubjectReviewStatus) DeepCopy() *PodSecurityPolicySubjectReviewStatus {
	if in == nil {
		return nil
	}
	out := new(PodSecurityPolicySubjectReviewStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RunAsUserStrategyOptions) DeepCopyInto(out *RunAsUserStrategyOptions) {
	*out = *in
	if in.UID != nil {
		in, out := &in.UID, &out.UID
		if *in == nil {
			*out = nil
		} else {
			*out = new(int64)
			**out = **in
		}
	}
	if in.UIDRangeMin != nil {
		in, out := &in.UIDRangeMin, &out.UIDRangeMin
		if *in == nil {
			*out = nil
		} else {
			*out = new(int64)
			**out = **in
		}
	}
	if in.UIDRangeMax != nil {
		in, out := &in.UIDRangeMax, &out.UIDRangeMax
		if *in == nil {
			*out = nil
		} else {
			*out = new(int64)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RunAsUserStrategyOptions.
func (in *RunAsUserStrategyOptions) DeepCopy() *RunAsUserStrategyOptions {
	if in == nil {
		return nil
	}
	out := new(RunAsUserStrategyOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RunAsUserStrategyType) DeepCopyInto(out *RunAsUserStrategyType) {
	{
		in := (*string)(unsafe.Pointer(in))
		out := (*string)(unsafe.Pointer(out))
		*out = *in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RunAsUserStrategyType.
func (in *RunAsUserStrategyType) DeepCopy() *RunAsUserStrategyType {
	if in == nil {
		return nil
	}
	out := new(RunAsUserStrategyType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SELinuxContextStrategyOptions) DeepCopyInto(out *SELinuxContextStrategyOptions) {
	*out = *in
	if in.SELinuxOptions != nil {
		in, out := &in.SELinuxOptions, &out.SELinuxOptions
		if *in == nil {
			*out = nil
		} else {
			*out = new(api.SELinuxOptions)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SELinuxContextStrategyOptions.
func (in *SELinuxContextStrategyOptions) DeepCopy() *SELinuxContextStrategyOptions {
	if in == nil {
		return nil
	}
	out := new(SELinuxContextStrategyOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SELinuxContextStrategyType) DeepCopyInto(out *SELinuxContextStrategyType) {
	{
		in := (*string)(unsafe.Pointer(in))
		out := (*string)(unsafe.Pointer(out))
		*out = *in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SELinuxContextStrategyType.
func (in *SELinuxContextStrategyType) DeepCopy() *SELinuxContextStrategyType {
	if in == nil {
		return nil
	}
	out := new(SELinuxContextStrategyType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecurityContextConstraints) DeepCopyInto(out *SecurityContextConstraints) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Priority != nil {
		in, out := &in.Priority, &out.Priority
		if *in == nil {
			*out = nil
		} else {
			*out = new(int32)
			**out = **in
		}
	}
	if in.DefaultAddCapabilities != nil {
		in, out := &in.DefaultAddCapabilities, &out.DefaultAddCapabilities
		*out = make([]api.Capability, len(*in))
		copy(*out, *in)
	}
	if in.RequiredDropCapabilities != nil {
		in, out := &in.RequiredDropCapabilities, &out.RequiredDropCapabilities
		*out = make([]api.Capability, len(*in))
		copy(*out, *in)
	}
	if in.AllowedCapabilities != nil {
		in, out := &in.AllowedCapabilities, &out.AllowedCapabilities
		*out = make([]api.Capability, len(*in))
		copy(*out, *in)
	}
	if in.Volumes != nil {
		in, out := &in.Volumes, &out.Volumes
		*out = make([]FSType, len(*in))
		copy(*out, *in)
	}
	if in.AllowedFlexVolumes != nil {
		in, out := &in.AllowedFlexVolumes, &out.AllowedFlexVolumes
		*out = make([]AllowedFlexVolume, len(*in))
		copy(*out, *in)
	}
	in.SELinuxContext.DeepCopyInto(&out.SELinuxContext)
	in.RunAsUser.DeepCopyInto(&out.RunAsUser)
	in.SupplementalGroups.DeepCopyInto(&out.SupplementalGroups)
	in.FSGroup.DeepCopyInto(&out.FSGroup)
	if in.SeccompProfiles != nil {
		in, out := &in.SeccompProfiles, &out.SeccompProfiles
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Users != nil {
		in, out := &in.Users, &out.Users
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Groups != nil {
		in, out := &in.Groups, &out.Groups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecurityContextConstraints.
func (in *SecurityContextConstraints) DeepCopy() *SecurityContextConstraints {
	if in == nil {
		return nil
	}
	out := new(SecurityContextConstraints)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SecurityContextConstraints) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecurityContextConstraintsList) DeepCopyInto(out *SecurityContextConstraintsList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SecurityContextConstraints, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecurityContextConstraintsList.
func (in *SecurityContextConstraintsList) DeepCopy() *SecurityContextConstraintsList {
	if in == nil {
		return nil
	}
	out := new(SecurityContextConstraintsList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SecurityContextConstraintsList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceAccountPodSecurityPolicyReviewStatus) DeepCopyInto(out *ServiceAccountPodSecurityPolicyReviewStatus) {
	*out = *in
	in.PodSecurityPolicySubjectReviewStatus.DeepCopyInto(&out.PodSecurityPolicySubjectReviewStatus)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceAccountPodSecurityPolicyReviewStatus.
func (in *ServiceAccountPodSecurityPolicyReviewStatus) DeepCopy() *ServiceAccountPodSecurityPolicyReviewStatus {
	if in == nil {
		return nil
	}
	out := new(ServiceAccountPodSecurityPolicyReviewStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SupplementalGroupsStrategyOptions) DeepCopyInto(out *SupplementalGroupsStrategyOptions) {
	*out = *in
	if in.Ranges != nil {
		in, out := &in.Ranges, &out.Ranges
		*out = make([]IDRange, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SupplementalGroupsStrategyOptions.
func (in *SupplementalGroupsStrategyOptions) DeepCopy() *SupplementalGroupsStrategyOptions {
	if in == nil {
		return nil
	}
	out := new(SupplementalGroupsStrategyOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SupplementalGroupsStrategyType) DeepCopyInto(out *SupplementalGroupsStrategyType) {
	{
		in := (*string)(unsafe.Pointer(in))
		out := (*string)(unsafe.Pointer(out))
		*out = *in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SupplementalGroupsStrategyType.
func (in *SupplementalGroupsStrategyType) DeepCopy() *SupplementalGroupsStrategyType {
	if in == nil {
		return nil
	}
	out := new(SupplementalGroupsStrategyType)
	in.DeepCopyInto(out)
	return out
}
