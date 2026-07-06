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
	kubevaultv1alpha2 "kubevault.dev/apimachinery/apis/kubevault/v1alpha2"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	appcat "kmodules.xyz/custom-resources/apis/appcatalog/v1alpha1"
)

const (
	ResourceKindVaultServerConfiguration = "VaultServerConfiguration"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// VaultServerConfiguration defines a Vault Server configuration.
type VaultServerConfiguration struct {
	// +optional
	metav1.TypeMeta `json:",inline,omitempty"`

	// Specifies the path which is used for authentication by this AppBinding.
	// If vault server is provisioned by KubeVault, this is usually `kubernetes`.
	// +optional
	Path string `json:"path,omitempty"`

	// Specifies the vault role name for policy controller
	// It has permission to create policy in vault
	// +optional
	VaultRole string `json:"vaultRole,omitempty"`

	// Specifies the Kubernetes authentication information
	// +optional
	Kubernetes *KubernetesAuthConfig `json:"kubernetes,omitempty"`

	// Specifies the Azure authentication information
	// +optional
	Azure *AzureAuthConfig `json:"azure,omitempty"`

	// Specifies the AWS authentication information
	// +optional
	AWS *AWSAuthConfig `json:"aws,omitempty"`

	// Specifies the Secret name that contains the token with permission for backup/restore
	// +optional
	BackupTokenSecretRef *core.LocalObjectReference `json:"backupTokenSecretRef,omitempty"`

	// Stash defines backup and restore task definitions.
	// +optional
	Stash appcat.StashAddonSpec `json:"stash,omitempty"`

	// backend storage information for vault
	// +optional
	Backend kubevaultv1alpha2.VaultServerBackend `json:"backend,omitempty"`

	// Unsealer configuration for vault
	// +optional
	Unsealer *kubevaultv1alpha2.UnsealerSpec `json:"unsealer,omitempty"`

	// DeploymentMode indicates how this AppBinding reaches the VaultServer:
	// Local (in-cluster vault, the default when absent) or RemoteRelay
	// (a hub vault accessed from a spoke cluster via the OpenBao spoke relay).
	// Consumers must read this through GetVaultDeploymentMode.
	// +optional
	DeploymentMode DeploymentMode `json:"deploymentMode,omitempty"`

	// SpokeName is the spoke cluster identity registered with the hub's
	// relay backend. Required when DeploymentMode is RemoteRelay. The secret
	// engine controllers use it to route database mounts through the
	// hub's remote-<db>-plugin proxies.
	// +optional
	SpokeName string `json:"spokeName,omitempty"`
}

// DeploymentMode distinguishes a locally reachable VaultServer from a
// hub VaultServer reached through the OpenBao spoke relay.
// +kubebuilder:validation:Enum=Local;RemoteRelay
type DeploymentMode string

const (
	// DeploymentModeLocal marks an AppBinding pointing at an in-cluster
	// (or directly reachable) VaultServer. Database secret engines use the
	// built-in <db>-database-plugin family.
	DeploymentModeLocal DeploymentMode = "Local"

	// DeploymentModeRemoteRelay marks an AppBinding authored for a spoke
	// cluster that reaches a hub VaultServer. Database secret engines must
	// use the remote-<db>-plugin family so the hub proxies plugin calls to
	// the spoke relay.
	DeploymentModeRemoteRelay DeploymentMode = "RemoteRelay"
)

// KubernetesAuthConfiguration contains necessary information for
// performing Kubernetes authentication to the Vault server.
type KubernetesAuthConfig struct {
	// Specifies the service account name
	ServiceAccountName string `json:"serviceAccountName"`

	// Specifies the service account name for token reviewer
	// It has system:auth-delegator permission
	// It's jwt token is used on vault kubernetes auth config
	// +optional
	TokenReviewerServiceAccountName string `json:"tokenReviewerServiceAccountName,omitempty"`

	// Specifies to use pod service account for vault csi driver
	// +optional
	UsePodServiceAccountForCSIDriver bool `json:"usePodServiceAccountForCSIDriver,omitempty"`
}

// AzureAuthConfig contains necessary information for
// performing Azure authentication to the Vault server.
type AzureAuthConfig struct {
	// Specifies the subscription ID for the machine
	// that generated the MSI token.
	// +optional
	SubscriptionID string `json:"subscriptionID,omitempty"`

	// Specifies the resource group for the machine
	// that generated the MSI token.
	// +optional
	ResourceGroupName string `json:"resourceGroupName,omitempty"`

	// Specifies the virtual machine name for the machine
	// that generated the MSI token. If VmssName is provided,
	// this value is ignored.
	// +optional
	VmName string `json:"vmName,omitempty"`

	// Specifies the virtual machine scale set name
	// for the machine that generated the MSI token.
	// +optional
	VmssName string `json:"vmssName,omitempty"`
}

// AWSAuthConfig contains necessary information for
// performing AWS authentication to the Vault server.
type AWSAuthConfig struct {
	// Specifies the header value that required
	// if X-Vault-AWS-IAM-Server-ID Header is set in Vault.
	// +optional
	HeaderValue string `json:"headerValue,omitempty"`
}
