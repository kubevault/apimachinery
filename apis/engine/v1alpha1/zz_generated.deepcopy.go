// +build !ignore_autogenerated

/*
Copyright 2019 The Kube Vault Authors.

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/api/rbac/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	appcatalogv1alpha1 "kmodules.xyz/custom-resources/apis/appcatalog/v1alpha1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSAccessKeyRequest) DeepCopyInto(out *AWSAccessKeyRequest) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSAccessKeyRequest.
func (in *AWSAccessKeyRequest) DeepCopy() *AWSAccessKeyRequest {
	if in == nil {
		return nil
	}
	out := new(AWSAccessKeyRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AWSAccessKeyRequest) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSAccessKeyRequestCondition) DeepCopyInto(out *AWSAccessKeyRequestCondition) {
	*out = *in
	in.LastUpdateTime.DeepCopyInto(&out.LastUpdateTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSAccessKeyRequestCondition.
func (in *AWSAccessKeyRequestCondition) DeepCopy() *AWSAccessKeyRequestCondition {
	if in == nil {
		return nil
	}
	out := new(AWSAccessKeyRequestCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSAccessKeyRequestList) DeepCopyInto(out *AWSAccessKeyRequestList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AWSAccessKeyRequest, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSAccessKeyRequestList.
func (in *AWSAccessKeyRequestList) DeepCopy() *AWSAccessKeyRequestList {
	if in == nil {
		return nil
	}
	out := new(AWSAccessKeyRequestList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AWSAccessKeyRequestList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSAccessKeyRequestSpec) DeepCopyInto(out *AWSAccessKeyRequestSpec) {
	*out = *in
	out.RoleRef = in.RoleRef
	if in.Subjects != nil {
		in, out := &in.Subjects, &out.Subjects
		*out = make([]v1.Subject, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSAccessKeyRequestSpec.
func (in *AWSAccessKeyRequestSpec) DeepCopy() *AWSAccessKeyRequestSpec {
	if in == nil {
		return nil
	}
	out := new(AWSAccessKeyRequestSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSAccessKeyRequestStatus) DeepCopyInto(out *AWSAccessKeyRequestStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]AWSAccessKeyRequestCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Secret != nil {
		in, out := &in.Secret, &out.Secret
		*out = new(corev1.LocalObjectReference)
		**out = **in
	}
	if in.Lease != nil {
		in, out := &in.Lease, &out.Lease
		*out = new(Lease)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSAccessKeyRequestStatus.
func (in *AWSAccessKeyRequestStatus) DeepCopy() *AWSAccessKeyRequestStatus {
	if in == nil {
		return nil
	}
	out := new(AWSAccessKeyRequestStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSConfig) DeepCopyInto(out *AWSConfig) {
	*out = *in
	if in.MaxRetries != nil {
		in, out := &in.MaxRetries, &out.MaxRetries
		*out = new(int)
		**out = **in
	}
	if in.LeaseConfig != nil {
		in, out := &in.LeaseConfig, &out.LeaseConfig
		*out = new(LeaseConfig)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSConfig.
func (in *AWSConfig) DeepCopy() *AWSConfig {
	if in == nil {
		return nil
	}
	out := new(AWSConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSRole) DeepCopyInto(out *AWSRole) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSRole.
func (in *AWSRole) DeepCopy() *AWSRole {
	if in == nil {
		return nil
	}
	out := new(AWSRole)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AWSRole) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSRoleCondition) DeepCopyInto(out *AWSRoleCondition) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSRoleCondition.
func (in *AWSRoleCondition) DeepCopy() *AWSRoleCondition {
	if in == nil {
		return nil
	}
	out := new(AWSRoleCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSRoleList) DeepCopyInto(out *AWSRoleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AWSRole, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSRoleList.
func (in *AWSRoleList) DeepCopy() *AWSRoleList {
	if in == nil {
		return nil
	}
	out := new(AWSRoleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AWSRoleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSRoleSpec) DeepCopyInto(out *AWSRoleSpec) {
	*out = *in
	if in.AuthManagerRef != nil {
		in, out := &in.AuthManagerRef, &out.AuthManagerRef
		*out = new(appcatalogv1alpha1.AppReference)
		(*in).DeepCopyInto(*out)
	}
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = new(AWSConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.RoleARNs != nil {
		in, out := &in.RoleARNs, &out.RoleARNs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.PolicyARNs != nil {
		in, out := &in.PolicyARNs, &out.PolicyARNs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSRoleSpec.
func (in *AWSRoleSpec) DeepCopy() *AWSRoleSpec {
	if in == nil {
		return nil
	}
	out := new(AWSRoleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSRoleStatus) DeepCopyInto(out *AWSRoleStatus) {
	*out = *in
	if in.ObservedGeneration != nil {
		in, out := &in.ObservedGeneration, &out.ObservedGeneration
		*out = (*in).DeepCopy()
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]AWSRoleCondition, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSRoleStatus.
func (in *AWSRoleStatus) DeepCopy() *AWSRoleStatus {
	if in == nil {
		return nil
	}
	out := new(AWSRoleStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GCPConfig) DeepCopyInto(out *GCPConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GCPConfig.
func (in *GCPConfig) DeepCopy() *GCPConfig {
	if in == nil {
		return nil
	}
	out := new(GCPConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GCPRole) DeepCopyInto(out *GCPRole) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GCPRole.
func (in *GCPRole) DeepCopy() *GCPRole {
	if in == nil {
		return nil
	}
	out := new(GCPRole)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GCPRole) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GCPRoleCondition) DeepCopyInto(out *GCPRoleCondition) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GCPRoleCondition.
func (in *GCPRoleCondition) DeepCopy() *GCPRoleCondition {
	if in == nil {
		return nil
	}
	out := new(GCPRoleCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GCPRoleList) DeepCopyInto(out *GCPRoleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GCPRole, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GCPRoleList.
func (in *GCPRoleList) DeepCopy() *GCPRoleList {
	if in == nil {
		return nil
	}
	out := new(GCPRoleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GCPRoleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GCPRoleSpec) DeepCopyInto(out *GCPRoleSpec) {
	*out = *in
	if in.AuthManagerRef != nil {
		in, out := &in.AuthManagerRef, &out.AuthManagerRef
		*out = new(appcatalogv1alpha1.AppReference)
		(*in).DeepCopyInto(*out)
	}
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = new(GCPConfig)
		**out = **in
	}
	if in.TokenScopes != nil {
		in, out := &in.TokenScopes, &out.TokenScopes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GCPRoleSpec.
func (in *GCPRoleSpec) DeepCopy() *GCPRoleSpec {
	if in == nil {
		return nil
	}
	out := new(GCPRoleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GCPRoleStatus) DeepCopyInto(out *GCPRoleStatus) {
	*out = *in
	if in.ObservedGeneration != nil {
		in, out := &in.ObservedGeneration, &out.ObservedGeneration
		*out = (*in).DeepCopy()
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]GCPRoleCondition, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GCPRoleStatus.
func (in *GCPRoleStatus) DeepCopy() *GCPRoleStatus {
	if in == nil {
		return nil
	}
	out := new(GCPRoleStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Lease) DeepCopyInto(out *Lease) {
	*out = *in
	out.Duration = in.Duration
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Lease.
func (in *Lease) DeepCopy() *Lease {
	if in == nil {
		return nil
	}
	out := new(Lease)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LeaseConfig) DeepCopyInto(out *LeaseConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LeaseConfig.
func (in *LeaseConfig) DeepCopy() *LeaseConfig {
	if in == nil {
		return nil
	}
	out := new(LeaseConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoleReference) DeepCopyInto(out *RoleReference) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoleReference.
func (in *RoleReference) DeepCopy() *RoleReference {
	if in == nil {
		return nil
	}
	out := new(RoleReference)
	in.DeepCopyInto(out)
	return out
}
