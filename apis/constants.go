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

package apis

import (
	core "k8s.io/api/core/v1"
)

const (
	Finalizer = "kubevault.com"

	// required fields:
	// - Secret.Data["token"] - a vault token
	SecretTypeTokenAuth core.SecretType = "kubevault.com/token"

	// required for SecretTypeTokenAut
	TokenAuthTokenKey = "token"

	// required fields:
	// - Secret.Data["access_key_id"] - aws access key id
	// - Secret.Data["secret_access_key"] - aws access secret key
	SecretTypeAWSAuth core.SecretType = "kubevault.com/aws"

	// required for SecretTypeAWSAuth
	AWSAuthAccessKeyIDKey = "access_key_id"
	// required for SecretTypeAWSAuth
	AWSAuthAccessSecretKey = "secret_access_key"
	// optional for SecretTypeAWSAuth
	AWSAuthSecurityTokenKey = "security_token"

	// required fields:
	// - Secret.Data["sa.json"] - gcp access secret key
	SecretTypeGCPAuth core.SecretType = "kubevault.com/gcp"
	// required for SecretTypeGCPAuth
	GCPAuthSACredentialJson = "sa.json"

	// - Secret.Data["msiToken"] - azure managed service identity (MSI)  jwt token
	SecretTypeAzureAuth = "kubevault.com/azure"

	// required for SecretTypeAzureAuth
	AzureMSIToken = "msiToken"
)

const (
	// moved from operator/pkg/controller/vault.go
	TLSCACertKey = "ca.crt"
)

const (
	VaultAuthK8sRole    = "role"
	VaultAuthApprole    = "role"
	VaultAuthLDAPGroups = "groups"
	VaultAuthLDAPUsers  = "users"
	VaultAuthJWTRole    = "role"
	VaultAuthOIDCRole   = "role"
)

const (
	CertificatePath            = "/etc/vault/tls"
	VaultServerCertsVolumeName = "vault-server-certs"
	VaultClientCertsVolumeName = "vault-client-certs"
)

// List of possible condition types for a KubeVault object

const (
	VaultServerInitializing        = "Initializing"
	VaultServerInitialized         = "Initialized"
	VaultServerUnsealing           = "Unsealing"
	VaultServerUnsealed            = "Unsealed"
	VaultServerAcceptingConnection = "AcceptingConnection"
	AllReplicasAreReady            = "AllReplicasReady"
	SomeReplicasAreNotReady        = "SomeReplicasNotReady"
	VaultServerPaused              = "Paused"
	VaultReadWriteOK               = "ReadWriteOK"

	// health check constants
	VaultHealthCheckPaused = "HealthCheckPaused"
	RaftLeaderHealthy      = "RaftLeaderHealthy"

	// ClientTrafficPinned latches the <vault-name>-primary Service to the active
	// (leader) node when spec.exposePrimary is true. It is set True once the
	// operator observes at least one vault pod carrying the active-node label, at
	// which point GetPrimaryService narrows the -primary Service selector to that
	// label. It is one-way: once True it stays True until exposePrimary is cleared
	// (false), so a leader election that momentarily leaves no pod labelled never
	// widens the -primary Service back to the standbys (which would serve stale
	// reads). The always-all-nodes <vault-name> Service is unaffected either way.
	// False with reason NoActiveNodeLabelled means exposePrimary is true but no pod
	// is labelled while a node is unsealed and ready: the pin did not take effect,
	// and the -primary Service is still selecting every node. See
	// design/primary-service-routing.md.
	ClientTrafficPinned = "ClientTrafficPinned"

	// NoActiveNodeLabelled is the reason on ClientTrafficPinned=False:
	// spec.exposePrimary is true but no vault pod carries the active-node label
	// while a node is unsealed and ready. The usual cause is the vault
	// ServiceAccount missing the pods get,patch grant, which the webhook cannot
	// catch. It is the only signal a user gets that the pin silently failed to
	// turn on.
	NoActiveNodeLabelled = "NoActiveNodeLabelled"

	// OCM spoke relay placement constants
	VaultServerRelayPlacementResolved    = "RelayPlacementResolved"
	VaultServerRelayHubInitialized       = "RelayHubInitialized"
	VaultServerRelayManifestWorksApplied = "RelayManifestWorksApplied"
	VaultServerRelaysReady               = "RelaysReady"

	// Tenant-isolation constants (see design/tenant-namespace-design.md).
	// VaultServerTenantIsolationReady is set when spec.isolateTenants is on and the
	// backend is confirmed namespace-capable (OpenBao / Vault Enterprise).
	VaultServerTenantIsolationReady = "TenantIsolationReady"
	// VaultServerTenantIsolationUnsupported is set when spec.isolateTenants is on but
	// the backend does not support namespaces; isolation is inert and SecretEngines
	// resolve to the root namespace.
	VaultServerTenantIsolationUnsupported = "TenantIsolationUnsupported"

	// Cross data center DR constants.
	// VaultServerLeaderInActiveDC is true when the active Vault node runs in the data
	// center holding the primary-DC Lease.
	VaultServerLeaderInActiveDC = "LeaderInActiveDC"
	// VaultServerDRCoordinationReady is false when the primary-DC Lease cannot be read
	// or has no holder, or when the guard token is missing, invalid or not delivered to
	// every Member cluster.
	VaultServerDRCoordinationReady = "DRCoordinationReady"
	// VaultServerUnsealKeysDistributed (kubernetesSecret unseal mode only) is true when
	// the unseal keys are applied in every Member cluster.
	VaultServerUnsealKeysDistributed = "UnsealKeysDistributed"
	// VaultServerKubernetesAuthConfigured is false when the operator cannot configure
	// Vault's Kubernetes auth (for example, no usable hub API address).
	VaultServerKubernetesAuthConfigured = "KubernetesAuthConfigured"
	// VaultServerInitKeysConsistent (distributed, kubernetesSecret unseal mode) is
	// false when Vault's initialization and the hub key Secret disagree: Vault is
	// initialized but the Secret has no keys, or the Secret has keys but Vault is
	// not initialized. The operator then never initializes Vault.
	VaultServerInitKeysConsistent = "InitKeysConsistent"
)

const (
	// SpokeRelayFinalizer guards hub-side cleanup of per-cluster spoke relay
	// resources (ManifestWorks, ServiceAccounts, bootstrap tokens).
	SpokeRelayFinalizer = "kubevault.com/spoke-relays"

	// ManagedByHubLabelValue marks resources authored on the hub and delivered
	// to managed clusters via ManifestWork. Spoke-side controllers must not
	// mutate objects carrying app.kubernetes.io/managed-by with this value.
	ManagedByHubLabelValue = "kubevault-hub"

	// Labels stamped on ManifestWorks (and other per-cluster hub resources)
	// to map them back to the owning VaultServer. Cross-namespace owner
	// references are not allowed, so labels replace ownership.
	LabelVaultServerName      = "kubevault.com/vaultserver-name"
	LabelVaultServerNamespace = "kubevault.com/vaultserver-namespace"

	// Labels stamped on the companion ManifestWorks of a distributed VaultServer.
	// They use their own keys, not LabelVaultServerName/Namespace, so the spoke
	// relay controller, which lists ManifestWorks by those, never matches them.
	LabelDCDRVaultServerName      = "dcdr.kubevault.com/vaultserver-name"
	LabelDCDRVaultServerNamespace = "dcdr.kubevault.com/vaultserver-namespace"
	LabelDCDRPurpose              = "dcdr.kubevault.com/purpose"

	// LabelDCDRPod is the per-pod label each Vault pod's unsealer puts on its own pod
	// in a distributed VaultServer; its value is the per-pod name <pod>-<dc>, which
	// the per-pod Service and ServiceExport select.
	LabelDCDRPod = "vaultserver.kubevault.com/pod"
)

const (
	// VaultRelayGRPCProxyPort is the hub-side gRPC proxy port used by spoke relays.
	VaultRelayGRPCProxyPort = 50053

	// AnnotationMigrateNamespace authorizes a one-time, lease-dropping move of a single
	// already-provisioned SecretEngine to its newly-desired OpenBao namespace. Set the
	// value to "true" on the SecretEngine; the operator clears it after migrating.
	// See design/tenant-namespace-design.md §5.2.
	AnnotationMigrateNamespace = "kubevault.com/migrate-namespace"

	// AnnotationMigrateVaultSecrets authorizes bulk migration of every SecretEngine whose
	// database lives in a client-org Namespace and whose spec.vaultRef matches one of the
	// listed Vault-server AppBindings. The value is a JSON array of AppBinding refs in
	// "namespace/name" form, e.g. ["kv/vault","kv/vault-2"]. Set on the client-org
	// Namespace; the operator prunes processed refs (removing the annotation once empty).
	AnnotationMigrateVaultSecrets = "kubevault.com/migrate-vault-secrets"
)

const (
	ResourceKindStatefulSet = "StatefulSet"
)

const (
	VaultAPIPort = 8200
)
