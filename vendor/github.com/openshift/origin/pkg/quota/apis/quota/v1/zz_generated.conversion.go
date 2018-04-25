// +build !ignore_autogenerated_openshift

// This file was autogenerated by conversion-gen. Do not edit it manually!

package v1

import (
	v1 "github.com/openshift/api/quota/v1"
	quota "github.com/openshift/origin/pkg/quota/apis/quota"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	core_v1 "k8s.io/kubernetes/pkg/apis/core/v1"
	unsafe "unsafe"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedConversionFuncs(
		Convert_v1_AppliedClusterResourceQuota_To_quota_AppliedClusterResourceQuota,
		Convert_quota_AppliedClusterResourceQuota_To_v1_AppliedClusterResourceQuota,
		Convert_v1_AppliedClusterResourceQuotaList_To_quota_AppliedClusterResourceQuotaList,
		Convert_quota_AppliedClusterResourceQuotaList_To_v1_AppliedClusterResourceQuotaList,
		Convert_v1_ClusterResourceQuota_To_quota_ClusterResourceQuota,
		Convert_quota_ClusterResourceQuota_To_v1_ClusterResourceQuota,
		Convert_v1_ClusterResourceQuotaList_To_quota_ClusterResourceQuotaList,
		Convert_quota_ClusterResourceQuotaList_To_v1_ClusterResourceQuotaList,
		Convert_v1_ClusterResourceQuotaSelector_To_quota_ClusterResourceQuotaSelector,
		Convert_quota_ClusterResourceQuotaSelector_To_v1_ClusterResourceQuotaSelector,
		Convert_v1_ClusterResourceQuotaSpec_To_quota_ClusterResourceQuotaSpec,
		Convert_quota_ClusterResourceQuotaSpec_To_v1_ClusterResourceQuotaSpec,
		Convert_v1_ClusterResourceQuotaStatus_To_quota_ClusterResourceQuotaStatus,
		Convert_quota_ClusterResourceQuotaStatus_To_v1_ClusterResourceQuotaStatus,
	)
}

func autoConvert_v1_AppliedClusterResourceQuota_To_quota_AppliedClusterResourceQuota(in *v1.AppliedClusterResourceQuota, out *quota.AppliedClusterResourceQuota, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1_ClusterResourceQuotaSpec_To_quota_ClusterResourceQuotaSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1_ClusterResourceQuotaStatus_To_quota_ClusterResourceQuotaStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1_AppliedClusterResourceQuota_To_quota_AppliedClusterResourceQuota is an autogenerated conversion function.
func Convert_v1_AppliedClusterResourceQuota_To_quota_AppliedClusterResourceQuota(in *v1.AppliedClusterResourceQuota, out *quota.AppliedClusterResourceQuota, s conversion.Scope) error {
	return autoConvert_v1_AppliedClusterResourceQuota_To_quota_AppliedClusterResourceQuota(in, out, s)
}

func autoConvert_quota_AppliedClusterResourceQuota_To_v1_AppliedClusterResourceQuota(in *quota.AppliedClusterResourceQuota, out *v1.AppliedClusterResourceQuota, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_quota_ClusterResourceQuotaSpec_To_v1_ClusterResourceQuotaSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_quota_ClusterResourceQuotaStatus_To_v1_ClusterResourceQuotaStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_quota_AppliedClusterResourceQuota_To_v1_AppliedClusterResourceQuota is an autogenerated conversion function.
func Convert_quota_AppliedClusterResourceQuota_To_v1_AppliedClusterResourceQuota(in *quota.AppliedClusterResourceQuota, out *v1.AppliedClusterResourceQuota, s conversion.Scope) error {
	return autoConvert_quota_AppliedClusterResourceQuota_To_v1_AppliedClusterResourceQuota(in, out, s)
}

func autoConvert_v1_AppliedClusterResourceQuotaList_To_quota_AppliedClusterResourceQuotaList(in *v1.AppliedClusterResourceQuotaList, out *quota.AppliedClusterResourceQuotaList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]quota.AppliedClusterResourceQuota, len(*in))
		for i := range *in {
			if err := Convert_v1_AppliedClusterResourceQuota_To_quota_AppliedClusterResourceQuota(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1_AppliedClusterResourceQuotaList_To_quota_AppliedClusterResourceQuotaList is an autogenerated conversion function.
func Convert_v1_AppliedClusterResourceQuotaList_To_quota_AppliedClusterResourceQuotaList(in *v1.AppliedClusterResourceQuotaList, out *quota.AppliedClusterResourceQuotaList, s conversion.Scope) error {
	return autoConvert_v1_AppliedClusterResourceQuotaList_To_quota_AppliedClusterResourceQuotaList(in, out, s)
}

func autoConvert_quota_AppliedClusterResourceQuotaList_To_v1_AppliedClusterResourceQuotaList(in *quota.AppliedClusterResourceQuotaList, out *v1.AppliedClusterResourceQuotaList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]v1.AppliedClusterResourceQuota, len(*in))
		for i := range *in {
			if err := Convert_quota_AppliedClusterResourceQuota_To_v1_AppliedClusterResourceQuota(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_quota_AppliedClusterResourceQuotaList_To_v1_AppliedClusterResourceQuotaList is an autogenerated conversion function.
func Convert_quota_AppliedClusterResourceQuotaList_To_v1_AppliedClusterResourceQuotaList(in *quota.AppliedClusterResourceQuotaList, out *v1.AppliedClusterResourceQuotaList, s conversion.Scope) error {
	return autoConvert_quota_AppliedClusterResourceQuotaList_To_v1_AppliedClusterResourceQuotaList(in, out, s)
}

func autoConvert_v1_ClusterResourceQuota_To_quota_ClusterResourceQuota(in *v1.ClusterResourceQuota, out *quota.ClusterResourceQuota, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1_ClusterResourceQuotaSpec_To_quota_ClusterResourceQuotaSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1_ClusterResourceQuotaStatus_To_quota_ClusterResourceQuotaStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1_ClusterResourceQuota_To_quota_ClusterResourceQuota is an autogenerated conversion function.
func Convert_v1_ClusterResourceQuota_To_quota_ClusterResourceQuota(in *v1.ClusterResourceQuota, out *quota.ClusterResourceQuota, s conversion.Scope) error {
	return autoConvert_v1_ClusterResourceQuota_To_quota_ClusterResourceQuota(in, out, s)
}

func autoConvert_quota_ClusterResourceQuota_To_v1_ClusterResourceQuota(in *quota.ClusterResourceQuota, out *v1.ClusterResourceQuota, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_quota_ClusterResourceQuotaSpec_To_v1_ClusterResourceQuotaSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_quota_ClusterResourceQuotaStatus_To_v1_ClusterResourceQuotaStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_quota_ClusterResourceQuota_To_v1_ClusterResourceQuota is an autogenerated conversion function.
func Convert_quota_ClusterResourceQuota_To_v1_ClusterResourceQuota(in *quota.ClusterResourceQuota, out *v1.ClusterResourceQuota, s conversion.Scope) error {
	return autoConvert_quota_ClusterResourceQuota_To_v1_ClusterResourceQuota(in, out, s)
}

func autoConvert_v1_ClusterResourceQuotaList_To_quota_ClusterResourceQuotaList(in *v1.ClusterResourceQuotaList, out *quota.ClusterResourceQuotaList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]quota.ClusterResourceQuota, len(*in))
		for i := range *in {
			if err := Convert_v1_ClusterResourceQuota_To_quota_ClusterResourceQuota(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1_ClusterResourceQuotaList_To_quota_ClusterResourceQuotaList is an autogenerated conversion function.
func Convert_v1_ClusterResourceQuotaList_To_quota_ClusterResourceQuotaList(in *v1.ClusterResourceQuotaList, out *quota.ClusterResourceQuotaList, s conversion.Scope) error {
	return autoConvert_v1_ClusterResourceQuotaList_To_quota_ClusterResourceQuotaList(in, out, s)
}

func autoConvert_quota_ClusterResourceQuotaList_To_v1_ClusterResourceQuotaList(in *quota.ClusterResourceQuotaList, out *v1.ClusterResourceQuotaList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]v1.ClusterResourceQuota, len(*in))
		for i := range *in {
			if err := Convert_quota_ClusterResourceQuota_To_v1_ClusterResourceQuota(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_quota_ClusterResourceQuotaList_To_v1_ClusterResourceQuotaList is an autogenerated conversion function.
func Convert_quota_ClusterResourceQuotaList_To_v1_ClusterResourceQuotaList(in *quota.ClusterResourceQuotaList, out *v1.ClusterResourceQuotaList, s conversion.Scope) error {
	return autoConvert_quota_ClusterResourceQuotaList_To_v1_ClusterResourceQuotaList(in, out, s)
}

func autoConvert_v1_ClusterResourceQuotaSelector_To_quota_ClusterResourceQuotaSelector(in *v1.ClusterResourceQuotaSelector, out *quota.ClusterResourceQuotaSelector, s conversion.Scope) error {
	out.LabelSelector = (*meta_v1.LabelSelector)(unsafe.Pointer(in.LabelSelector))
	out.AnnotationSelector = *(*map[string]string)(unsafe.Pointer(&in.AnnotationSelector))
	return nil
}

// Convert_v1_ClusterResourceQuotaSelector_To_quota_ClusterResourceQuotaSelector is an autogenerated conversion function.
func Convert_v1_ClusterResourceQuotaSelector_To_quota_ClusterResourceQuotaSelector(in *v1.ClusterResourceQuotaSelector, out *quota.ClusterResourceQuotaSelector, s conversion.Scope) error {
	return autoConvert_v1_ClusterResourceQuotaSelector_To_quota_ClusterResourceQuotaSelector(in, out, s)
}

func autoConvert_quota_ClusterResourceQuotaSelector_To_v1_ClusterResourceQuotaSelector(in *quota.ClusterResourceQuotaSelector, out *v1.ClusterResourceQuotaSelector, s conversion.Scope) error {
	out.LabelSelector = (*meta_v1.LabelSelector)(unsafe.Pointer(in.LabelSelector))
	out.AnnotationSelector = *(*map[string]string)(unsafe.Pointer(&in.AnnotationSelector))
	return nil
}

// Convert_quota_ClusterResourceQuotaSelector_To_v1_ClusterResourceQuotaSelector is an autogenerated conversion function.
func Convert_quota_ClusterResourceQuotaSelector_To_v1_ClusterResourceQuotaSelector(in *quota.ClusterResourceQuotaSelector, out *v1.ClusterResourceQuotaSelector, s conversion.Scope) error {
	return autoConvert_quota_ClusterResourceQuotaSelector_To_v1_ClusterResourceQuotaSelector(in, out, s)
}

func autoConvert_v1_ClusterResourceQuotaSpec_To_quota_ClusterResourceQuotaSpec(in *v1.ClusterResourceQuotaSpec, out *quota.ClusterResourceQuotaSpec, s conversion.Scope) error {
	if err := Convert_v1_ClusterResourceQuotaSelector_To_quota_ClusterResourceQuotaSelector(&in.Selector, &out.Selector, s); err != nil {
		return err
	}
	if err := core_v1.Convert_v1_ResourceQuotaSpec_To_core_ResourceQuotaSpec(&in.Quota, &out.Quota, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1_ClusterResourceQuotaSpec_To_quota_ClusterResourceQuotaSpec is an autogenerated conversion function.
func Convert_v1_ClusterResourceQuotaSpec_To_quota_ClusterResourceQuotaSpec(in *v1.ClusterResourceQuotaSpec, out *quota.ClusterResourceQuotaSpec, s conversion.Scope) error {
	return autoConvert_v1_ClusterResourceQuotaSpec_To_quota_ClusterResourceQuotaSpec(in, out, s)
}

func autoConvert_quota_ClusterResourceQuotaSpec_To_v1_ClusterResourceQuotaSpec(in *quota.ClusterResourceQuotaSpec, out *v1.ClusterResourceQuotaSpec, s conversion.Scope) error {
	if err := Convert_quota_ClusterResourceQuotaSelector_To_v1_ClusterResourceQuotaSelector(&in.Selector, &out.Selector, s); err != nil {
		return err
	}
	if err := core_v1.Convert_core_ResourceQuotaSpec_To_v1_ResourceQuotaSpec(&in.Quota, &out.Quota, s); err != nil {
		return err
	}
	return nil
}

// Convert_quota_ClusterResourceQuotaSpec_To_v1_ClusterResourceQuotaSpec is an autogenerated conversion function.
func Convert_quota_ClusterResourceQuotaSpec_To_v1_ClusterResourceQuotaSpec(in *quota.ClusterResourceQuotaSpec, out *v1.ClusterResourceQuotaSpec, s conversion.Scope) error {
	return autoConvert_quota_ClusterResourceQuotaSpec_To_v1_ClusterResourceQuotaSpec(in, out, s)
}

func autoConvert_v1_ClusterResourceQuotaStatus_To_quota_ClusterResourceQuotaStatus(in *v1.ClusterResourceQuotaStatus, out *quota.ClusterResourceQuotaStatus, s conversion.Scope) error {
	if err := core_v1.Convert_v1_ResourceQuotaStatus_To_core_ResourceQuotaStatus(&in.Total, &out.Total, s); err != nil {
		return err
	}
	if err := Convert_v1_ResourceQuotasStatusByNamespace_To_quota_ResourceQuotasStatusByNamespace(&in.Namespaces, &out.Namespaces, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1_ClusterResourceQuotaStatus_To_quota_ClusterResourceQuotaStatus is an autogenerated conversion function.
func Convert_v1_ClusterResourceQuotaStatus_To_quota_ClusterResourceQuotaStatus(in *v1.ClusterResourceQuotaStatus, out *quota.ClusterResourceQuotaStatus, s conversion.Scope) error {
	return autoConvert_v1_ClusterResourceQuotaStatus_To_quota_ClusterResourceQuotaStatus(in, out, s)
}

func autoConvert_quota_ClusterResourceQuotaStatus_To_v1_ClusterResourceQuotaStatus(in *quota.ClusterResourceQuotaStatus, out *v1.ClusterResourceQuotaStatus, s conversion.Scope) error {
	if err := core_v1.Convert_core_ResourceQuotaStatus_To_v1_ResourceQuotaStatus(&in.Total, &out.Total, s); err != nil {
		return err
	}
	if err := Convert_quota_ResourceQuotasStatusByNamespace_To_v1_ResourceQuotasStatusByNamespace(&in.Namespaces, &out.Namespaces, s); err != nil {
		return err
	}
	return nil
}

// Convert_quota_ClusterResourceQuotaStatus_To_v1_ClusterResourceQuotaStatus is an autogenerated conversion function.
func Convert_quota_ClusterResourceQuotaStatus_To_v1_ClusterResourceQuotaStatus(in *quota.ClusterResourceQuotaStatus, out *v1.ClusterResourceQuotaStatus, s conversion.Scope) error {
	return autoConvert_quota_ClusterResourceQuotaStatus_To_v1_ClusterResourceQuotaStatus(in, out, s)
}