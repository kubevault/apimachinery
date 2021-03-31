/*
Copyright AppsCode Inc. and Contributors

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

package v1alpha1

import (
	"fmt"

	kmapi "kmodules.xyz/client-go/api/v1"
	"kmodules.xyz/client-go/meta"
	ofst "kmodules.xyz/offshoot-api/api/v1"
)

// +kubebuilder:validation:Enum=Provisioning;DataRestoring;Ready;Critical;NotReady;Halted;Sealed;Unsealed;Initializing;Initialized
type VaultServerPhase string

const (
	// used for VaultServer that are currently provisioning
	VaultServerPhaseProvisioning VaultServerPhase = "Provisioning"
	// used for VaultServer for which data is currently restoring
	VaultServerPhaseDataRestoring VaultServerPhase = "DataRestoring"
	// used for VaultServer that are currently ReplicaReady, AcceptingConnection and Ready
	VaultServerPhaseReady VaultServerPhase = "Ready"
	// used for VaultServer that can connect, ReplicaReady == false || Ready == false (eg, ES yellow)
	VaultServerPhaseCritical VaultServerPhase = "Critical"
	// used for VaultServer that can't connect
	VaultServerPhaseNotReady VaultServerPhase = "NotReady"
	// used for VaultServer that are halted
	VaultServerPhaseHalted VaultServerPhase = "Halted"

	// used for VaultServer that are sealed
	VaultServerPhaseSealed VaultServerPhase = "Sealed"
	// used for VaultServer that are unsealed
	VaultServerPhaseUnsealed VaultServerPhase = "Unsealed"
	// used for VaultServer that are initializing
	VaultServerPhaseInitializing VaultServerPhase = "Initializing"
	// used for VaultServer that are initialized
	VaultServerPhaseInitialized VaultServerPhase = "Initialized"
)

// +kubebuilder:validation:Enum=Halt;Delete;WipeOut;DoNotTerminate
type TerminationPolicy string

const (
	// Deletes VaultServer pods, service but leave the PVCs and stash backup data intact.
	TerminationPolicyHalt TerminationPolicy = "Halt"
	// Deletes VaultServer pods, service, pvcs but leave the stash backup data intact.
	TerminationPolicyDelete TerminationPolicy = "Delete"
	// Deletes VaultServer pods, service, pvcs and stash backup data.
	TerminationPolicyWipeOut TerminationPolicy = "WipeOut"
	// Rejects attempt to delete VaultServer using ValidationWebhook.
	TerminationPolicyDoNotTerminate TerminationPolicy = "DoNotTerminate"
)

// +kubebuilder:validation:Enum=primary;vault;stats
type ServiceAlias string

const (
	PrimaryVaultServerService ServiceAlias = "primary"
	StandbyVaultServerService ServiceAlias = "vault"
	StatsVaultServerService   ServiceAlias = "stats"
)

type NamedServiceTemplateSpec struct {
	// Alias represents the identifier of the service.
	Alias ServiceAlias `json:"alias" protobuf:"bytes,1,opt,name=alias"`

	// ServiceTemplate is an optional configuration for a service used to expose VaultServer
	// +optional
	ofst.ServiceTemplateSpec `json:",inline,omitempty" protobuf:"bytes,2,opt,name=serviceTemplateSpec"`
}

// Returns the default certificate secret name for given alias.
func (vs *VaultServer) DefaultCertSecretName(alias string) string {
	return meta.NameWithSuffix(fmt.Sprintf("%s-%s", vs.Name, alias), "certs")
}

// Returns certificate secret name for given alias if exists,
// otherwise returns the default certificate secret name.
func (vs *VaultServer) GetCertSecretName(alias string) string {
	if vs.Spec.TLS != nil {
		sName, valid := kmapi.GetCertificateSecretName(vs.Spec.TLS.Certificates, alias)
		if valid {
			return sName
		}
	}

	return vs.DefaultCertSecretName(alias)
}
