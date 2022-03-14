package v1alpha1

import (
	"unsafe"

	"kubevault.dev/apimachinery/apis/kubevault/v1alpha2"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/conversion"
	apiv1 "kmodules.xyz/client-go/api/v1"
	monitoringagentapiapiv1 "kmodules.xyz/monitoring-agent-api/api/v1"
)

func Convert_v1alpha1_VaultServerSpec_To_v1alpha2_VaultServerSpec(in *VaultServerSpec, out *v1alpha2.VaultServerSpec, s conversion.Scope) error {
	out.Version = in.Version
	out.Replicas = (*int32)(unsafe.Pointer(in.Replicas))
	out.ConfigSecretRef = (*v1.LocalObjectReference)(unsafe.Pointer(in.ConfigSecret))
	out.DataSources = *(*[]v1.VolumeSource)(unsafe.Pointer(&in.DataSources))
	out.TLS = (*apiv1.TLSConfig)(unsafe.Pointer(in.TLS))
	if err := Convert_v1alpha1_BackendStorageSpec_To_v1alpha2_BackendStorageSpec(&in.Backend, &out.Backend, s); err != nil {
		return err
	}
	out.Unsealer = (*v1alpha2.UnsealerSpec)(unsafe.Pointer(in.Unsealer))
	out.AuthMethods = *(*[]v1alpha2.AuthMethod)(unsafe.Pointer(&in.AuthMethods))
	out.Monitor = (*monitoringagentapiapiv1.AgentSpec)(unsafe.Pointer(in.Monitor))
	out.PodTemplate = in.PodTemplate
	out.ServiceTemplates = *(*[]v1alpha2.NamedServiceTemplateSpec)(unsafe.Pointer(&in.ServiceTemplates))
	out.Halted = in.Halted
	out.TerminationPolicy = v1alpha2.TerminationPolicy(in.TerminationPolicy)
	out.AllowedSecretEngines = (*v1alpha2.AllowedSecretEngines)(unsafe.Pointer(in.AllowedSecretEngines))
	return nil
}

func Convert_v1alpha2_VaultServerSpec_To_v1alpha1_VaultServerSpec(in *v1alpha2.VaultServerSpec, out *VaultServerSpec, s conversion.Scope) error {
	out.Version = in.Version
	out.Replicas = (*int32)(unsafe.Pointer(in.Replicas))
	out.ConfigSecret = (*v1.LocalObjectReference)(unsafe.Pointer(in.ConfigSecretRef))
	out.DataSources = *(*[]v1.VolumeSource)(unsafe.Pointer(&in.DataSources))
	out.TLS = (*apiv1.TLSConfig)(unsafe.Pointer(in.TLS))
	if err := Convert_v1alpha2_BackendStorageSpec_To_v1alpha1_BackendStorageSpec(&in.Backend, &out.Backend, s); err != nil {
		return err
	}
	out.Unsealer = (*UnsealerSpec)(unsafe.Pointer(in.Unsealer))
	out.AuthMethods = *(*[]AuthMethod)(unsafe.Pointer(&in.AuthMethods))
	out.Monitor = (*monitoringagentapiapiv1.AgentSpec)(unsafe.Pointer(in.Monitor))
	out.PodTemplate = in.PodTemplate
	out.ServiceTemplates = *(*[]NamedServiceTemplateSpec)(unsafe.Pointer(&in.ServiceTemplates))
	out.Halted = in.Halted
	out.TerminationPolicy = TerminationPolicy(in.TerminationPolicy)
	out.AllowedSecretEngines = (*AllowedSecretEngines)(unsafe.Pointer(in.AllowedSecretEngines))
	return nil
}
