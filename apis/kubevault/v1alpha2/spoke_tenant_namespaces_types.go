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

package v1alpha2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kmapi "kmodules.xyz/client-go/api/v1"
)

const (
	ResourceKindSpokeTenantNamespaces = "SpokeTenantNamespaces"
	ResourceSpokeTenantNamespaces     = "spoketenantnamespaces"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=spoketenantnamespaces,singular=spoketenantnamespaces,scope=Cluster,shortName=stns,categories={vault,appscode,all}
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Orgs",type="integer",JSONPath=".status.orgCount"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"

// SpokeTenantNamespaces is a spoke-cluster CR maintained by the KubeVault operator on a
// managed (spoke) cluster. Its status carries the deduplicated, validated set of OpenBao
// org namespaces the spoke's client-org databases require on the hub. The hub reads it
// through an OCM ManagedClusterView and idempotently creates each namespace — the spoke
// never creates hub namespaces itself (design/tenant-namespace-hub-spoke-design.md §5.3).
type SpokeTenantNamespaces struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SpokeTenantNamespacesSpec   `json:"spec,omitempty"`
	Status            SpokeTenantNamespacesStatus `json:"status,omitempty"`
}

type SpokeTenantNamespacesSpec struct {
	// HubVaultRef identifies the hub VaultServer AppBinding (the RemoteAgent AppBinding on
	// this spoke) whose org namespaces this object tracks.
	// +optional
	HubVaultRef kmapi.ObjectReference `json:"hubVaultRef,omitempty"`
}

type SpokeTenantNamespacesStatus struct {
	// ObservedGeneration is the most recent generation observed for this resource.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	// OrgIDs is the deduplicated, validated set of OpenBao namespaces (org-ids) that the
	// spoke's client-org databases targeting the hub require. The hub ensures each exists.
	// +optional
	OrgIDs []string `json:"orgIDs,omitempty"`

	// OrgCount mirrors len(OrgIDs) for a cheap print column / bounded status-feedback scalar.
	// +optional
	OrgCount int32 `json:"orgCount,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type SpokeTenantNamespacesList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SpokeTenantNamespaces `json:"items,omitempty"`
}
