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
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kmapi "kmodules.xyz/client-go/api/v1"
	ofst "kmodules.xyz/offshoot-api/api/v1"
)

const (
	ResourceKindVaultAgent = "VaultAgent"
	ResourceVaultAgent     = "vaultagent"
	ResourceVaultAgents    = "vaultagents"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=vaultagents,singular=vaultagent,shortName=va,categories={vault,appscode,all}
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Spoke",type="string",JSONPath=".spec.spokeName"
// +kubebuilder:printcolumn:name="Status",type="string",JSONPath=".status.phase"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
type VaultAgent struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VaultAgentSpec   `json:"spec,omitempty"`
	Status            VaultAgentStatus `json:"status,omitempty"`
}

type VaultAgentSpec struct {
	// HubVaultRef references the hub VaultServer
	HubVaultRef HubVaultReference `json:"hubVaultRef"`

	// SpokeName is the unique identifier for this spoke cluster
	SpokeName string `json:"spokeName"`

	// TokenSecretRef references a secret containing vault token for authentication
	// Secret data:
	//  - token: <vault-token>
	// +optional
	TokenSecretRef *core.LocalObjectReference `json:"tokenSecretRef,omitempty"`

	// Image is the spoke-agent container image
	// +optional
	Image string `json:"image,omitempty"`

	// TLS configuration for gRPC connection
	// +optional
	TLS *VaultAgentTLSConfig `json:"tls,omitempty"`

	// Reconnect settings for automatic reconnection
	// +optional
	// +kubebuilder:default={enabled: true, backoffSeconds: 5, maxBackoffSeconds: 300}
	Reconnect *ReconnectConfig `json:"reconnect,omitempty"`

	// PodTemplate is an optional configuration for the spoke-agent pod
	// +optional
	PodTemplate ofst.PodTemplateSpec `json:"podTemplate,omitempty"`
}

// HubVaultReference contains information to connect to hub vault
type HubVaultReference struct {
	// Name of the VaultServer in hub cluster
	Name string `json:"name"`

	// Namespace of the VaultServer in hub cluster
	Namespace string `json:"namespace"`

	// Address is the hub vault URL (e.g., http://10.2.0.88:30820)
	Address string `json:"address"`

	// GRPCPort is the port for gRPC proxy communication
	// +optional
	// +kubebuilder:default=50053
	GRPCPort int32 `json:"grpcPort,omitempty"`
}

// VaultAgentTLSConfig contains TLS configuration for spoke-agent
type VaultAgentTLSConfig struct {
	// Enabled indicates whether TLS is enabled
	// +optional
	Enabled bool `json:"enabled,omitempty"`

	// CASecret contains the CA certificate
	// Secret data:
	//  - ca.crt: <value>
	// +optional
	CASecret *core.LocalObjectReference `json:"caSecret,omitempty"`

	// CertSecret contains the client certificate and key
	// Secret data:
	//  - tls.crt: <value>
	//  - tls.key: <value>
	// +optional
	CertSecret *core.LocalObjectReference `json:"certSecret,omitempty"`
}

// ReconnectConfig contains automatic reconnection settings
type ReconnectConfig struct {
	// Enabled indicates whether auto-reconnect is enabled
	// +optional
	// +kubebuilder:default=true
	Enabled bool `json:"enabled,omitempty"`

	// BackoffSeconds is the initial backoff duration in seconds
	// +optional
	// +kubebuilder:default=5
	BackoffSeconds int32 `json:"backoffSeconds,omitempty"`

	// MaxBackoffSeconds is the maximum backoff duration in seconds
	// +optional
	// +kubebuilder:default=300
	MaxBackoffSeconds int32 `json:"maxBackoffSeconds,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type VaultAgentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VaultAgent `json:"items,omitempty"`
}

// +kubebuilder:validation:Enum=Pending;Connected;Disconnected;Error
type VaultAgentPhase string

const (
	VaultAgentPhasePending      VaultAgentPhase = "Pending"
	VaultAgentPhaseConnected    VaultAgentPhase = "Connected"
	VaultAgentPhaseDisconnected VaultAgentPhase = "Disconnected"
	VaultAgentPhaseError        VaultAgentPhase = "Error"
)

type VaultAgentStatus struct {
	// ObservedGeneration is the most recent generation observed
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	// Phase indicates the current state of the VaultAgent
	// +optional
	Phase VaultAgentPhase `json:"phase,omitempty"`

	// LastHeartbeat is the timestamp of the last successful heartbeat
	// +optional
	LastHeartbeat *metav1.Time `json:"lastHeartbeat,omitempty"`

	// AppBindingRef references the created AppBinding for hub vault
	// +optional
	AppBindingRef *kmapi.ObjectReference `json:"appBindingRef,omitempty"`

	// PodName is the name of the spoke-agent pod
	// +optional
	PodName string `json:"podName,omitempty"`

	// Conditions represent the latest available observations of the VaultAgent's state
	// +optional
	Conditions []kmapi.Condition `json:"conditions,omitempty"`
}
