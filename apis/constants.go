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
)

const (
	// VaultRelayGRPCProxyPort is the hub-side gRPC proxy port used by spoke relays.
	VaultRelayGRPCProxyPort = 50053
)

const (
	ResourceKindStatefulSet = "StatefulSet"
)

const (
	VaultAPIPort = 8200
)
