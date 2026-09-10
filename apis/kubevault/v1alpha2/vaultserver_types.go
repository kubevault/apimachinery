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
	"time"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kmapi "kmodules.xyz/client-go/api/v1"
	appcat "kmodules.xyz/custom-resources/apis/appcatalog/v1alpha1"
	mona "kmodules.xyz/monitoring-agent-api/api/v1"
	ofst "kmodules.xyz/offshoot-api/api/v1"
)

const (
	ResourceKindVaultServer = "VaultServer"
	ResourceVaultServer     = "vaultserver"
	ResourceVaultServers    = "vaultservers"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:resource:path=vaultservers,singular=vaultserver,shortName=vs,categories={vault,appscode,all}
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Replicas",type="string",JSONPath=".spec.replicas"
// +kubebuilder:printcolumn:name="Version",type="string",JSONPath=".spec.version"
// +kubebuilder:printcolumn:name="Status",type="string",JSONPath=".status.phase"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
type VaultServer struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VaultServerSpec   `json:"spec,omitempty"`
	Status            VaultServerStatus `json:"status,omitempty"`
}

type VaultServerSpec struct {
	// Version of VaultServer to be deployed.
	Version string `json:"version"`

	// Number of instances to deploy for a VaultServer.
	Replicas *int32 `json:"replicas,omitempty"`

	// ConfigSecret is an optional field to provide extra configuration for vault.
	// This secret contain extra config for vault
	// File name should be 'vault.hcl'.
	// If specified, this file will be appended to the controller configuration file.
	// +optional
	ConfigSecret *core.LocalObjectReference `json:"configSecret,omitempty"`

	// DataSources is a list of Configmaps/Secrets in the same namespace as the VaultServer
	// object, which shall be mounted into the VaultServer Pods.
	// The data are mounted into /etc/vault/data/<name>.
	// The first data will be named as "data-0", second one will be named as "data-1" and so on.
	// +optional
	DataSources []core.VolumeSource `json:"dataSources,omitempty"`

	// TLS policy of vault nodes
	// +optional
	TLS *kmapi.TLSConfig `json:"tls,omitempty"`

	// backend storage configuration for vault
	Backend BackendStorageSpec `json:"backend"`

	// Unsealer configuration for vault
	// +optional
	Unsealer *UnsealerSpec `json:"unsealer,omitempty"`

	// Specifies the list of auth methods to enable
	// +optional
	AuthMethods []AuthMethod `json:"authMethods,omitempty"`

	// Monitor is used monitor database instance
	// +optional
	Monitor *mona.AgentSpec `json:"monitor,omitempty"`

	// PodTemplate is an optional configuration for pods used to run vault
	// +optional
	PodTemplate ofst.PodTemplateSpec `json:"podTemplate,omitempty"`

	// ServiceTemplates is an optional configuration for services used to expose database
	// +optional
	ServiceTemplates []NamedServiceTemplateSpec `json:"serviceTemplates,omitempty"`

	// Indicates that the vault server is halted and all offshoot Kubernetes resources except PVCs are deleted.
	// +optional
	Halted bool `json:"halted,omitempty"`

	// TerminationPolicy controls the delete operation for vault server
	// +optional
	TerminationPolicy TerminationPolicy `json:"terminationPolicy,omitempty"`

	// AllowedSecretEngines defines the types of Secret Engines that MAY be attached to a
	// Listener and the trusted namespaces where those Route resources MAY be
	// present.
	//
	// Although a client request may match multiple route rules, only one rule
	// may ultimately receive the request. Matching precedence MUST be
	// determined in order of the following criteria:
	//
	// * The most specific match as defined by the Route type.
	// * The oldest Route based on creation timestamp. For example, a Route with
	//   a creation timestamp of "2020-09-08 01:02:03" is given precedence over
	//   a Route with a creation timestamp of "2020-09-08 01:02:04".
	// * If everything else is equivalent, the Route appearing first in
	//   alphabetical order (namespace/name) should be given precedence. For
	//   example, foo/bar is given precedence over foo/baz.
	//
	// All valid rules within a Route attached to this Listener should be
	// implemented. Invalid Route rules can be ignored (sometimes that will mean
	// the full Route). If a Route rule transitions from valid to invalid,
	// support for that Route rule should be dropped to ensure consistency. For
	// example, even if a filter specified by a Route rule is invalid, the rest
	// of the rules within that Route should still be supported.
	//
	// Support: Core
	// +kubebuilder:default={namespaces:{from: Same}}
	// +optional
	AllowedSecretEngines *AllowedSecretEngines `json:"allowedSecretEngines,omitempty"`

	// HealthChecker defines attributes of the health checker
	// +optional
	// +kubebuilder:default={periodSeconds: 10, timeoutSeconds: 10, failureThreshold: 1}
	HealthChecker kmapi.HealthCheckSpec `json:"healthChecker"`

	// ExposePrimary, when true, makes the operator create an additional
	// <vault-name>-primary Service whose selector narrows to the active (leader)
	// node, alongside the always-all-nodes <vault-name> Service. A client that
	// requires strict read-after-write consistency, and cannot tolerate reading from
	// a lagging standby, binds to the primary Service; everything else, including the
	// default AppBinding, keeps using the all-nodes Service and is unaffected.
	//
	// Requires an HA-capable storage backend and a supported distribution; the
	// admission webhook rejects it otherwise. The primary Service has no endpoints
	// during a leader election (the brief window with no active node), which is the
	// cost of the guarantee. Defaults to false. See design/primary-service-routing.md.
	// +optional
	ExposePrimary bool `json:"exposePrimary,omitempty"`

	// IsolateTenants is the master opt-in for per-tenant OpenBao namespace isolation.
	// When true (and the backend is namespace-capable, i.e. OpenBao / Vault Enterprise),
	// SecretEngines on this server whose referenced database lives in a client-org
	// namespace are provisioned in a matching OpenBao namespace keyed on the org-id, and
	// an explicit SecretEngine.spec.namespace is also honored. When false (the default),
	// everything stays in the root namespace and any SecretEngine.spec.namespace is rejected.
	// Enabling this on a running server is safe: existing mounts are never moved on their own
	// (migration is admin-authorized). See design/tenant-namespace-design.md.
	// +optional
	IsolateTenants bool `json:"isolateTenants,omitempty"`

	// RelayPlacementRef points to an OCM Placement object (cluster.open-cluster-management.io/v1beta1)
	// in the same namespace as the VaultServer. When set, the operator deploys a VaultRelay to
	// every managed cluster selected by the Placement, using one ManifestWork per cluster.
	// Requires the OCM hub CRDs (Placement, PlacementDecision, ManifestWork) to be installed;
	// the field is ignored with a warning condition otherwise.
	// +optional
	RelayPlacementRef *core.LocalObjectReference `json:"relayPlacementRef,omitempty"`

	// RelayTemplate customizes the VaultRelays stamped out for clusters selected by
	// RelayPlacementRef. Per-cluster fields (spokeName, hubVaultRef, join material)
	// are filled in by the operator.
	// +optional
	RelayTemplate *VaultRelayTemplate `json:"relayTemplate,omitempty"`
}

// VaultRelayTemplate is the subset of VaultRelaySpec a hub admin may pre-set for
// placement-driven spoke relay deployments.
type VaultRelayTemplate struct {
	// Namespace on the managed cluster where the VaultRelay and its companion
	// resources are created. Defaults to the VaultServer's namespace.
	// +optional
	Namespace string `json:"namespace,omitempty"`

	// PodTemplate is an optional configuration for the spoke-relay pods.
	// +optional
	PodTemplate ofst.PodTemplateSpec `json:"podTemplate,omitempty"`

	// BootstrapTokenTTL controls the TTL of hub bootstrap tokens minted per spoke
	// (and therefore their rotation period). Default 24h, minimum 1h.
	// +optional
	BootstrapTokenTTL *metav1.Duration `json:"bootstrapTokenTTL,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type VaultServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VaultServer `json:"items,omitempty"`
}

type VaultServerStatus struct {
	// ObservedGeneration is the most recent generation observed for this resource. It corresponds to the
	// resource's generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	// Phase indicates the state this Vault server jumps in.
	// +optional
	Phase VaultServerPhase `json:"phase,omitempty"`

	// Initialized indicates if the Vault service is initialized.
	// +optional
	Initialized bool `json:"initialized,omitempty"`

	// ServiceName is the LB service for accessing vault nodes.
	// +optional
	ServiceName string `json:"serviceName,omitempty"`

	// ClientPort is the port for vault client to access.
	// It's the same on client LB service and vault nodes.
	// +optional
	ClientPort int64 `json:"clientPort,omitempty"`

	// VaultStatus is the set of Vault node specific statuses: Active, Standby, and Sealed
	// +optional
	VaultStatus VaultStatus `json:"vaultStatus,omitempty"`

	// PodNames of updated Vault nodes. Updated means the Vault container image version
	// matches the spec's version.
	// +optional
	UpdatedNodes []string `json:"updatedNodes,omitempty"`

	// Represents the latest available observations of a VaultServer current state.
	// +optional
	Conditions []kmapi.Condition `json:"conditions,omitempty"`

	// Status of the vault auth methods
	// +optional
	AuthMethodStatus []AuthMethodStatus `json:"authMethodStatus,omitempty"`

	// RelayPlacement summarizes spoke relay rollout when spec.relayPlacementRef is set.
	// +optional
	RelayPlacement *RelayPlacementStatus `json:"relayPlacement,omitempty"`
}

// RelayPlacementStatus summarizes the rollout of spoke relays to managed clusters.
type RelayPlacementStatus struct {
	// Placement is the resolved Placement name.
	// +optional
	Placement string `json:"placement,omitempty"`

	// Selected is the number of clusters currently listed in the PlacementDecisions.
	// +optional
	Selected int32 `json:"selected,omitempty"`

	// Applied is the number of clusters whose ManifestWork has condition Applied=True.
	// +optional
	Applied int32 `json:"applied,omitempty"`

	// Ready is the number of clusters whose VaultRelay reports phase Connected
	// (scraped via ManifestWork status feedback).
	// +optional
	Ready int32 `json:"ready,omitempty"`

	// Clusters holds per-cluster detail.
	// +optional
	Clusters []SpokeClusterStatus `json:"clusters,omitempty"`
}

// SpokeClusterStatus is the per managed cluster rollout state.
type SpokeClusterStatus struct {
	// ClusterName is the ManagedCluster name.
	ClusterName string `json:"clusterName"`

	// Phase mirrors the spoke VaultRelay phase (Pending|Connected|Disconnected|Error)
	// plus hub-side values (WorkApplied, WorkProgressing, WorkDegraded).
	// +optional
	Phase string `json:"phase,omitempty"`

	// TokenExpiry is when the current bootstrap token for this spoke expires.
	// +optional
	TokenExpiry *metav1.Time `json:"tokenExpiry,omitempty"`

	// CertExpiry is when this spoke's mTLS client certificate expires, as
	// observed by the hub relay backend (relay/spokes). Nil when unknown: the
	// spoke is not connected, or the hub captured no verified peer cert.
	// +optional
	CertExpiry *metav1.Time `json:"certExpiry,omitempty"`
}

// AllowedSecretEngines defines which Secret Engines may be attached to this Listener.
type AllowedSecretEngines struct {
	// Namespaces indicates namespaces from which Secret Engines may be attached to this
	// Listener. This is restricted to the namespace of this VaultServer by default.
	//
	// +optional
	// +kubebuilder:default={from: Same}
	Namespaces *SecretEngineNamespaces `json:"namespaces,omitempty"`

	// SecretEngines specifies the types of Secret Engines that are allowed to bind
	// to this VaultServer. When unspecified or empty, all types of Secret Engines
	// are allowed.
	//
	// +optional
	SecretEngines []SecretEngineType `json:"secretEngines,omitempty"`
}

// +kubebuilder:validation:Enum=kv;pki;aws;azure;gcp;postgres;mongodb;mysql;mariadb;elasticsearch;redis;db2;druid;hanadb;hazelcast;ignite;kafka;memcached;milvus;mssqlserver;neo4j;oracle;qdrant;rabbitmq;solr;weaviate;zookeeper
type SecretEngineType string

const (
	SecretEngineTypeKV            SecretEngineType = "kv"
	SecretEngineTypePKI           SecretEngineType = "pki"
	SecretEngineTypeAWS           SecretEngineType = "aws"
	SecretEngineTypeAzure         SecretEngineType = "azure"
	SecretEngineTypeGCP           SecretEngineType = "gcp"
	SecretEngineTypePostgres      SecretEngineType = "postgres"
	SecretEngineTypeMongoDB       SecretEngineType = "mongodb"
	SecretEngineTypeMySQL         SecretEngineType = "mysql"
	SecretEngineTypeMariaDB       SecretEngineType = "mariadb"
	SecretEngineTypeElasticsearch SecretEngineType = "elasticsearch"
	SecretEngineTypeRedis         SecretEngineType = "redis"
	SecretEngineTypeDB2           SecretEngineType = "db2"
	SecretEngineTypeDruid         SecretEngineType = "druid"
	SecretEngineTypeHanaDB        SecretEngineType = "hanadb"
	SecretEngineTypeHazelcast     SecretEngineType = "hazelcast"
	SecretEngineTypeIgnite        SecretEngineType = "ignite"
	SecretEngineTypeKafka         SecretEngineType = "kafka"
	SecretEngineTypeMemcached     SecretEngineType = "memcached"
	SecretEngineTypeMilvus        SecretEngineType = "milvus"
	SecretEngineTypeMSSQLServer   SecretEngineType = "mssqlserver"
	SecretEngineTypeNeo4j         SecretEngineType = "neo4j"
	SecretEngineTypeOracle        SecretEngineType = "oracle"
	SecretEngineTypeQdrant        SecretEngineType = "qdrant"
	SecretEngineTypeRabbitMQ      SecretEngineType = "rabbitmq"
	SecretEngineTypeSolr          SecretEngineType = "solr"
	SecretEngineTypeWeaviate      SecretEngineType = "weaviate"
	SecretEngineTypeZooKeeper     SecretEngineType = "zookeeper"
)

// FromNamespaces specifies namespace from which Secret Engines may be attached to a
// VaultServer.
//
// +kubebuilder:validation:Enum=All;Selector;Same
type FromNamespaces string

const (
	// Secret Engines in all namespaces may be attached to this VaultServer.
	NamespacesFromAll FromNamespaces = "All"
	// Only Secret Engines in namespaces selected by the selector may be attached to
	// this VaultServer.
	NamespacesFromSelector FromNamespaces = "Selector"
	// Only Secret Engines in the same namespace as the VaultServer may be attached to this
	// VaultServer.
	NamespacesFromSame FromNamespaces = "Same"
)

// SecretEngineNamespaces indicate which namespaces Secret Engines should be selected from.
type SecretEngineNamespaces struct {
	// From indicates where Secret Engines will be selected for this VaultServer. Possible
	// values are:
	// * All: Secret Engines in all namespaces may be used by this VaultServer.
	// * Selector: Secret Engines in namespaces selected by the selector may be used by
	//   this VaultServer.
	// * Same: Only Secret Engines in the same namespace may be used by this VaultServer.
	//
	// +optional
	// +kubebuilder:default=Same
	From *FromNamespaces `json:"from,omitempty"`

	// Selector must be specified when From is set to "Selector". In that case,
	// only Secret Engines in Namespaces matching this Selector will be selected by this
	// VaultServer. This field is ignored for other values of "From".
	//
	// +optional
	Selector *metav1.LabelSelector `json:"selector,omitempty"`
}

type VaultStatus struct {
	// PodName of the active (leader) Vault node. The active node is unsealed and
	// every write is committed there. Standby nodes serve reads and forward
	// writes to it. The <vault-name>-primary Service, created when
	// spec.exposePrimary is true, points at this node alone; the <vault-name>
	// Service always points at all nodes.
	// Empty while no node is active, for example during a leader election.
	// +optional
	Active string `json:"active,omitempty"`

	// PodNames of the standby Vault nodes. Standby nodes are unsealed.
	// Standby nodes do not process requests, and instead redirect to the active Vault.
	// +optional
	Standby []string `json:"standby,omitempty"`

	// PodNames of Sealed Vault nodes. Sealed nodes MUST be unsealed to
	// become standby or leader.
	// +optional
	Sealed []string `json:"sealed,omitempty"`

	// PodNames of Unsealed Vault nodes.
	// +optional
	Unsealed []string `json:"unsealed,omitempty"`
}

// TLSPolicy defines the TLS policy of the vault nodes
// If this is not set, operator will auto-gen TLS assets and secrets.
type TLSPolicy struct {
	// TLSSecret is the secret containing TLS certs used by each vault node
	// for the communication between the vault server and its clients.
	// The secret should contain three files:
	// 	- tls.crt
	// 	- tls.key
	//
	// The server certificate must allow the following wildcard domains:
	// 	- localhost
	// 	- *.<namespace>.pod
	// 	- <vaultServer-name>.<namespace>.svc
	TLSSecret string `json:"tlsSecret"`

	// CABundle is a PEM encoded CA bundle which will be used to validate the serving certificate.
	// +optional
	CABundle []byte `json:"caBundle,omitempty"`
}

// TODO : set defaults and validation
// BackendStorageSpec defines storage backend configuration of vault
//
// FoundationDB is intentionally not represented here: unlike every other
// backend, it requires CGo and the native libfdb_c client library, so it
// isn't compiled into standard OpenBao/Vault images and can't be assumed
// available in a generic VaultServer deployment.
//
// New backends added here (CockroachDB, Cassandra, Zookeeper, CouchDB,
// MSSQL, Spanner, Aerospike, OCI, AlicloudOSS) are v1alpha2-only for now;
// v1alpha1 already predates several v1alpha2-only conventions used by
// their spec types (SecretRef-typed credentials, DatabaseRef, HAEnabled),
// so mirroring them would need bespoke field-mapping conversion functions
// rather than the mechanical 1:1 copy used elsewhere. v1alpha2 is the
// storage/hub version and the only one the operator's storage package
// consumes; a v1alpha1 backport can follow the existing
// Convert_v1alpha1_<X>Spec_To_v1alpha2_<X>Spec pattern if ever needed.
type BackendStorageSpec struct {
	// ref: https://www.vaultproject.io/docs/configuration/storage/in-memory.html
	// +optional
	Inmem *InmemSpec `json:"inmem,omitempty"`

	// +optional
	Etcd *EtcdSpec `json:"etcd,omitempty"`

	// +optional
	Gcs *GcsSpec `json:"gcs,omitempty"`

	// +optional
	S3 *S3Spec `json:"s3,omitempty"`

	// +optional
	Azure *AzureSpec `json:"azure,omitempty"`

	// +optional
	PostgreSQL *PostgreSQLSpec `json:"postgresql,omitempty"`

	// +optional
	MySQL *MySQLSpec `json:"mysql,omitempty"`

	// +optional
	File *FileSpec `json:"file,omitempty"`

	// +optional
	DynamoDB *DynamoDBSpec `json:"dynamodb,omitempty"`

	// +optional
	Swift *SwiftSpec `json:"swift,omitempty"`

	// +optional
	Consul *ConsulSpec `json:"consul,omitempty"`

	// +optional
	Raft *RaftSpec `json:"raft,omitempty"`

	// +optional
	CockroachDB *CockroachDBSpec `json:"cockroachdb,omitempty"`

	// +optional
	Cassandra *CassandraSpec `json:"cassandra,omitempty"`

	// +optional
	Zookeeper *ZookeeperSpec `json:"zookeeper,omitempty"`

	// +optional
	CouchDB *CouchDBSpec `json:"couchdb,omitempty"`

	// +optional
	MSSQL *MSSQLSpec `json:"mssql,omitempty"`

	// +optional
	Spanner *SpannerSpec `json:"spanner,omitempty"`

	// +optional
	Aerospike *AerospikeSpec `json:"aerospike,omitempty"`

	// +optional
	OCI *OCISpec `json:"oci,omitempty"`

	// +optional
	AlicloudOSS *AlicloudOSSSpec `json:"alicloudoss,omitempty"`
}

// ref: https://www.vaultproject.io/docs/configuration/storage/consul.html
//
// ConsulSpec defines the configuration to set up consul as backend storage in vault
type ConsulSpec struct {
	// Specifies the address of the Consul agent to communicate with.
	// This can be an IP address, DNS record, or unix socket.
	// +optional
	Address string `json:"address,omitempty"`

	// Specifies the check interval used to send health check information
	// back to Consul.
	// This is specified using a label suffix like "30s" or "1h".
	// +optional
	CheckTimeout string `json:"checkTimeout,omitempty"`

	// Specifies the Consul consistency mode.
	// Possible values are "default" or "strong".
	// +optional
	ConsistencyMode string `json:"consistencyMode,omitempty"`

	// Specifies whether Vault should register itself with Consul.
	// Possible values are "true" or "false"
	// +optional
	DisableRegistration string `json:"disableRegistration,omitempty"`

	// Specifies the maximum number of concurrent requests to Consul.
	// +optional
	MaxParallel string `json:"maxParallel,omitempty"`

	// Specifies the path in Consul's key-value store
	// where Vault data will be stored.
	// +optional
	Path string `json:"path,omitempty"`

	// Specifies the scheme to use when communicating with Consul.
	// This can be set to "http" or "https".
	// +optional
	Scheme string `json:"scheme,omitempty"`

	// Specifies the name of the service to register in Consul.
	// +optional
	Service string `json:"service,omitempty"`

	// Specifies a comma-separated list of tags
	// to attach to the service registration in Consul.
	// +optional
	ServiceTags string `json:"serviceTags,omitempty"`

	// Specifies a service-specific address to set on the service registration
	// in Consul.
	// If unset, Vault will use what it knows to be the HA redirect address
	// - which is usually desirable.
	// Setting this parameter to "" will tell Consul to leverage the configuration
	// of the node the service is registered on dynamically.
	// +optional
	ServiceAddress string `json:"serviceAddress,omitempty"`

	// Specifies the secret name that contains ACL token with permission
	// to read and write from the path in Consul's key-value store.
	// secret data:
	//  - aclToken:<value>
	// +optional
	ACLTokenSecretRef *core.LocalObjectReference `json:"aclTokenSecretRef,omitempty"`

	// Specifies the minimum allowed session TTL.
	// Consul server has a lower limit of 10s on the session TTL by default.
	// +optional
	SessionTTL string `json:"sessionTTL,omitempty"`

	// Specifies the wait time before a lock lock acquisition is made.
	// This affects the minimum time it takes to cancel a lock acquisition.
	// +optional
	LockWaitTime string `json:"lockWaitTime,omitempty"`

	// Specifies the secret name that contains tls_ca_file, tls_cert_file and tls_key_file
	// for consul communication
	// Secret data:
	//  - ca.crt
	//  - tls.crt
	//  - tls.key
	// +optional
	TLSSecretRef *core.LocalObjectReference `json:"tlsSecretRef,omitempty"`

	// Specifies the minimum TLS version to use.
	// Accepted values are "tls10", "tls11" or "tls12".
	// +optional
	TLSMinVersion string `json:"tlsMinVersion,omitempty"`

	// Specifies if the TLS host verification should be disabled.
	// It is highly discouraged that you disable this option.
	// +optional
	TLSSkipVerify bool `json:"tlsSkipVerify,omitempty"`
}

// ref: https://www.vaultproject.io/docs/configuration/storage/in-memory.html
type InmemSpec struct{}

// TODO : set defaults and validation
// vault doc: https://www.vaultproject.io/docs/configuration/storage/etcd.html
//
// EtcdSpec defines configuration to set up etcd as backend storage in vault
type EtcdSpec struct {
	// Specifies the addresses of the etcd instances
	Address string `json:"address"`

	// Specifies the version of the API to communicate with etcd
	// +optional
	EtcdApi string `json:"etcdApi,omitempty"`

	// Specifies if high availability should be enabled
	// +optional
	HAEnable bool `json:"haEnable,omitempty"`

	// Specifies the path in etcd where vault data will be stored
	// +optional
	Path string `json:"path,omitempty"`

	// Specifies whether to sync list of available etcd services on startup
	// +optional
	Sync bool `json:"sync,omitempty"`

	// Specifies the domain name to query for SRV records describing cluster endpoints
	// +optional
	DiscoverySrv string `json:"discoverySrv,omitempty"`

	// Specifies the secret name that contain username and password to use when authenticating with the etcd server
	// secret data:
	//  - username:<value>
	//  - password:<value>
	// +optional
	CredentialSecretRef *core.LocalObjectReference `json:"credentialSecretRef,omitempty"`

	// Specifies the secret name that contains tls_ca_file, tls_cert_file and tls_key_file for etcd communication
	// secret data:
	//  - ca.crt
	//  - tls.crt
	//  - tls.key
	// +optional
	TLSSecretRef *core.LocalObjectReference `json:"tlsSecretRef,omitempty"`
}

// vault doc: https://www.vaultproject.io/docs/configuration/storage/google-cloud-storage.html
//
// GcsSpec defines configuration to set up Google Cloud Storage as backend storage in vault
type GcsSpec struct {
	// Specifies the name of the bucket to use for storage.
	Bucket string `json:"bucket"`

	// Specifies the maximum size (in kilobytes) to send in a single request. If set to 0,
	// it will attempt to send the whole object at once, but will not retry any failures.
	// +optional
	ChunkSize string `json:"chunkSize,omitempty"`

	//  Specifies the maximum number of parallel operations to take place.
	// +optional
	MaxParallel int64 `json:"maxParallel,omitempty"`

	// Specifies if high availability mode is enabled.
	// +optional
	HAEnabled bool `json:"haEnabled,omitempty"`

	// Secret containing Google application credential
	// secret data:
	//  - sa.json:<value>
	// +optional
	CredentialSecretRef *core.LocalObjectReference `json:"credentialSecretRef,omitempty"`
}

// vault doc: https://www.vaultproject.io/docs/configuration/storage/s3.html
//
// S3Spec defines configuration to set up Amazon S3 Storage as backend storage in vault
type S3Spec struct {
	// Specifies the name of the bucket to use for storage.
	Bucket string `json:"bucket"`

	// Specifies an alternative, AWS compatible, S3 endpoint.
	// +optional
	Endpoint string `json:"endpoint,omitempty"`

	// Specifies the AWS region
	// +optional
	Region string `json:"region,omitempty"`

	// Specifies the secret name containing AWS session token, AWS access key and AWS secret key
	// secret data:
	//  - access_key=<value>
	//  - secret_key=<value>
	//  - session_token=<value>
	// +optional
	CredentialSecretRef *core.LocalObjectReference `json:"credentialSecretRef,omitempty"`

	// Specifies the maximum number of parallel operations to take place.
	// +optional
	MaxParallel int64 `json:"maxParallel,omitempty"`

	// Specifies whether to use host bucket style domains with the configured endpoint.
	// +optional
	ForcePathStyle bool `json:"forcePathStyle,omitempty"`

	// Specifies if SSL should be used for the endpoint connection
	// +optional
	DisableSSL bool `json:"disableSSL,omitempty"`
}

// vault doc: https://www.vaultproject.io/docs/configuration/storage/azure.html
//
// AzureSpec defines configuration to set up Google Cloud Storage as backend storage in vault
type AzureSpec struct {
	// Specifies the Azure Storage account name.
	AccountName string `json:"accountName"`

	// Specifies the secret containing Azure Storage account key.
	// secret data:
	//  - account_key:<value>
	// +optional
	CredentialSecretRef *core.LocalObjectReference `json:"credentialSecretRef,omitempty"`

	// Specifies the Azure Storage Blob container name.
	Container string `json:"container"`

	//  Specifies the maximum number of concurrent operations to take place.
	// +optional
	MaxParallel int64 `json:"maxParallel,omitempty"`
}

// vault doc: https://www.vaultproject.io/docs/configuration/storage/postgresql.html
//
// PostgreSQLSpec defines configuration to set up PostgreSQL storage as backend storage in vault
type PostgreSQLSpec struct {
	// Specifies the address of the Postgres host.
	// if DatabaseRef is set then Address will be generated from it
	// This must be set if DatabaseRef is empty, validate from ValidatingWebhook
	// host example: <db-name>.<db-ns>.svc:3306
	// +optional
	Address string `json:"address"`

	//  - username=<value>
	//  - password=<value>
	//  - connection_url="postgres://<username>:<password>@<host>:<port>/<db_name>"
	CredentialSecretRef *core.LocalObjectReference `json:"credentialSecretRef,omitempty"`

	// DatabaseRef contains the info of KubeDB managed Database
	// This will be used to generate the "Address" field
	DatabaseRef *appcat.AppReference `json:"databaseRef,omitempty"`

	// SSLMode for both standalone and clusters. [disable;require;verify-ca;verify-full]
	SSLMode PostgresSSLMode `json:"sslMode,omitempty"`

	// Specifies the name of the table in which to write Vault data.
	// This table must already exist (Vault will not attempt to create it).
	// +optional
	// +kubebuilder:default:="vault_kv_store"
	Table string `json:"table,omitempty"`

	//  Specifies the maximum number of concurrent requests to take place.
	// +optional
	MaxParallel int64 `json:"maxParallel,omitempty"`

	// Default not set. Sets the maximum number of connections in the idle connection pool.
	// See golang docs on SetMaxIdleConns(https://pkg.go.dev/database/sql#DB.SetMaxIdleConns) for more information. Requires 1.2 or later.
	// +optional
	MaxIdleConnection int64 `json:"maxIdleConnection,omitempty"`

	// High Availability Parameter
	// Default not enabled, requires 9.5 or later
	// Specifies if high availability mode is enabled. This is a boolean value, but it is specified as a string like "true" or "false".
	// +optional
	// +kubebuilder:default:="false"
	HAEnabled string `json:"haEnabled,omitempty"`

	// Specifies the name of the table to use for storing high availability information. This table must already exist (Vault will not attempt to create it).
	// +optional
	// +kubebuilder:default:="vault_ha_locks"
	HATable string `json:"haTable,omitempty"`
}

// +kubebuilder:validation:Enum=disable;require;verify-ca;verify-full
type PostgresSSLMode string

const (
	// PostgresSSLModeDisable represents `disable` sslMode. It ensures that the server does not use TLS/SSL.
	PostgresSSLModeDisable PostgresSSLMode = "disable"

	// Always SSL (skip verification)
	PostgressSSLModeRequire PostgresSSLMode = "require"

	// Always SSL (verify that the certificate presented by the server was signed by a trusted CA)
	PostgressSSLModeVerifyCA PostgresSSLMode = "verify-ca"

	// PostgresSSLModeVerifyFull represents `verify-full` sslmode. I want my data encrypted, and I accept the overhead.
	// I want to be sure that I connect to a server I trust, and that it's the one I specify.
	PostgresSSLModeVerifyFull PostgresSSLMode = "verify-full"
)

// vault doc: https://www.vaultproject.io/docs/configuration/storage/mysql.html
//
// MySQLSpec defines configuration to set up MySQL Storage as backend storage in vault
type MySQLSpec struct {
	// Specifies the address of the MySQL host.
	// if DatabaseRef is set then Address will be generated from it
	// This must be set if DatabaseRef is empty, validate from ValidatingWebhook
	// host example: <db-name>.<db-ns>.svc:3306
	// +optional
	Address string `json:"address"`

	// Specifies the name of the database. If the database does not exist, Vault will attempt to create it.
	// +optional
	// +kubebuilder:default:="vault"
	Database string `json:"database,omitempty"`

	// Specifies the name of the table. If the table does not exist, Vault will attempt to create it.
	// +optional
	// +kubebuilder:default:="vault"
	Table string `json:"table,omitempty"`

	// Specifies the MySQL username and password to connect to the database
	// secret data:
	//  - username=<value>
	//  - password=<value>
	CredentialSecretRef *core.LocalObjectReference `json:"credentialSecretRef,omitempty"`

	// Specifies the name of the secret containing the CA certificate to connect using TLS.
	// secret data:
	//  - ca.crt=<value>
	// +optional
	TLSSecretRef *core.LocalObjectReference `json:"tlsSecretRef,omitempty"`

	//  Specifies the maximum number of concurrent requests to take place.
	// +optional
	MaxParallel int64 `json:"maxParallel,omitempty"`

	// DatabaseRef contains the info of KubeDB managed Database
	// This will be used to generate the "Address" field
	// +optional
	DatabaseRef *appcat.AppReference `json:"databaseRef,omitempty"`

	PlaintextCredentialTransmission string `json:"plaintextCredentialTransmission,omitempty"`

	// Specifies the maximum number of idle connections to the database.
	// A zero uses value defaults to 2 idle connections and a negative value disables idle connections.
	// If larger than max_parallel it will be reduced to be equal.
	// +optional
	MaxIdleConnection int64 `json:"maxIdleConnection,omitempty"`

	// Specifies the maximum amount of time in seconds that a connection may be reused. If <= 0s connections are reused forever.
	// +optional
	MaxConnectionLifetime int64 `json:"maxConnectionLifetime,omitempty"`

	// High Availability Parameter
	// Specifies if high availability mode is enabled. This is a boolean value, but it is specified as a string like "true" or "false".
	// +optional
	// +kubebuilder:default:="true"
	HAEnabled string `json:"haEnabled,omitempty"`

	// High Availability Parameter
	// Specifies the name of the table to use for storing high availability information.
	// By default, this is the name of the table suffixed with _lock. If the table does not exist, Vault will attempt to create it.
	// +optional
	// +kubebuilder:default:="vault_lock"
	LockTable string `json:"lockTable,omitempty"`
}

// vault doc: https://www.vaultproject.io/docs/configuration/storage/filesystem.html
//
// FileSpec defines configuration to set up File system Storage as backend storage in vault
type FileSpec struct {
	// The absolute path on disk to the directory where the data will be stored.
	// If the directory does not exist, Vault will create it.
	Path string `json:"path"`

	// volumeClaimTemplate is a claim that pods are allowed to reference.
	// The VaultServer controller is responsible for deploying the claim
	// and update the volumeMounts in the Vault server container in the template.
	VolumeClaimTemplate ofst.PersistentVolumeClaim `json:"volumeClaimTemplate"`
}

// vault doc: https://www.vaultproject.io/docs/configuration/storage/dynamodb.html
//
// DynamoDBSpec defines configuration to set up DynamoDB Storage as backend storage in vault
type DynamoDBSpec struct {
	// Specifies an alternative, AWS compatible, DynamoDB endpoint.
	// +optional
	Endpoint string `json:"endpoint,omitempty"`

	// Specifies the AWS region
	// +optional
	Region string `json:"region,omitempty"`

	// Specifies whether this backend should be used to run Vault in high availability mode.
	// +optional
	HaEnabled bool `json:"haEnabled,omitempty"`

	// Specifies the maximum number of reads consumed per second on the table
	// +optional
	ReadCapacity int64 `json:"readCapacity,omitempty"`

	// Specifies the maximum number of writes performed per second on the table.
	// +optional
	WriteCapacity int64 `json:"writeCapacity,omitempty"`

	// Specifies the name of the DynamoDB table in which to store Vault data.
	// If the specified table does not yet exist, it will be created during initialization.
	// default: vault-dynamodb-backend
	// +optional
	Table string `json:"table,omitempty"`

	// Specifies the secret name containing AWS session token, AWS access key and AWS secret key
	// secret data:
	//  - access_key=<value>
	//  - secret_key=<value>
	//  - session_token=<value>
	// +optional
	CredentialSecretRef *core.LocalObjectReference `json:"credentialSecretRef,omitempty"`

	// Specifies the maximum number of parallel operations to take place.
	// +optional
	MaxParallel int64 `json:"maxParallel,omitempty"`
}

// vault doc: https://www.vaultproject.io/docs/configuration/storage/swift.html
//
// SwiftSpec defines configuration to set up Swift Storage as backend storage in vault
type SwiftSpec struct {
	// Specifies the OpenStack authentication endpoint.
	AuthURL string `json:"authURL"`

	// Specifies the name of the Swift container.
	Container string `json:"container"`

	// Specifies the name of the secret containing the OpenStack account/username and password
	// Specifies secret containing auth token from alternate authentication.
	// secret data:
	//  - username=<value>
	//  - password=<value>
	//  - auth_token=<value>
	// +optional
	CredentialSecretRef *core.LocalObjectReference `json:"credentialSecretRef,omitempty"`

	// Specifies the name of the tenant. If left blank, this will default to the default tenant of the username.
	// +optional
	Tenant string `json:"tenant,omitempty"`

	// Specifies the name of the region.
	// +optional
	Region string `json:"region,omitempty"`

	// Specifies the id of the tenant.
	// +optional
	TenantID string `json:"tenantID,omitempty"`

	// Specifies the name of the user domain.
	// +optional
	Domain string `json:"domain,omitempty"`

	// Specifies the name of the project's domain.
	// +optional
	ProjectDomain string `json:"projectDomain,omitempty"`

	// Specifies the id of the trust.
	// +optional
	TrustID string `json:"trustID,omitempty"`

	// Specifies storage URL from alternate authentication.
	// +optional
	StorageURL string `json:"storageURL,omitempty"`

	//  Specifies the maximum number of concurrent requests to take place.
	// +optional
	MaxParallel int64 `json:"maxParallel,omitempty"`
}

// UnsealerSpec contain the configuration for auto vault initialize/unseal
type UnsealerSpec struct {
	// Total count of secret shares that exist
	// +optional
	SecretShares int64 `json:"secretShares,omitempty"`

	// Minimum required secret shares to unseal
	// +optional
	SecretThreshold int64 `json:"secretThreshold,omitempty"`

	// How often to attempt to unseal the vault instance
	// +optional
	RetryPeriodSeconds time.Duration `json:"retryPeriodSeconds,omitempty"` //nolint

	// overwrite existing unseal keys and root tokens, possibly dangerous!
	// +optional
	OverwriteExisting bool `json:"overwriteExisting,omitempty"`

	// should the root token be stored in the key store (default true)
	// +optional
	StoreRootToken bool `json:"storeRootToken,omitempty"`

	// mode contains unseal mechanism
	// +optional
	Mode ModeSpec `json:"mode,omitempty"`
}

// ModeSpec contain unseal mechanism
type ModeSpec struct {
	// +optional
	KubernetesSecret *KubernetesSecretSpec `json:"kubernetesSecret,omitempty"`

	// +optional
	GoogleKmsGcs *GoogleKmsGcsSpec `json:"googleKmsGcs,omitempty"`

	// +optional
	AwsKmsSsm *AwsKmsSsmSpec `json:"awsKmsSsm,omitempty"`

	// +optional
	AzureKeyVault *AzureKeyVault `json:"azureKeyVault,omitempty"`
}

// KubernetesSecretSpec contain the fields that required to unseal using kubernetes secret
type KubernetesSecretSpec struct {
	SecretName string `json:"secretName"`
}

// GoogleKmsGcsSpec contain the fields that required to unseal vault using google kms
type GoogleKmsGcsSpec struct {
	// The name of the Google Cloud KMS crypto key to use
	KmsCryptoKey string `json:"kmsCryptoKey"`

	// The name of the Google Cloud KMS key ring to use
	KmsKeyRing string `json:"kmsKeyRing"`

	// The Google Cloud KMS location to use (eg. 'global', 'europe-west1')
	KmsLocation string `json:"kmsLocation"`

	// The Google Cloud KMS project to use
	KmsProject string `json:"kmsProject"`

	// The name of the Google Cloud Storage bucket to store values in
	Bucket string `json:"bucket"`

	// Secret containing Google application credential
	// secret data:
	//  - sa.json:<value>
	// +optional
	CredentialSecretRef *core.LocalObjectReference `json:"credentialSecretRef,omitempty"`
}

// AwsKmsSsmSpec contain the fields that required to unseal vault using aws kms ssm
type AwsKmsSsmSpec struct {
	// The ID or ARN of the AWS KMS key to encrypt values
	KmsKeyID string `json:"kmsKeyID"`

	// +optional
	// An optional Key prefix for SSM Parameter store
	SsmKeyPrefix string `json:"ssmKeyPrefix,omitempty"`

	Region string `json:"region,omitempty"`

	// Specifies the secret name containing AWS access key and AWS secret key
	// secret data:
	//  - access_key:<value>
	//  - secret_key:<value>
	// +optional
	CredentialSecretRef *core.LocalObjectReference `json:"credentialSecretRef,omitempty"`

	// Used to make AWS KMS requests. This is useful,
	// for example, when connecting to KMS over a VPC Endpoint.
	// If not set, Vault will use the default API endpoint for your region.
	Endpoint string `json:"endpoint,omitempty"`
}

// RaftSpec defines the configuration for the Raft integrated storage.
// https://www.vaultproject.io/docs/configuration/storage/raft
type RaftSpec struct {
	// An integer multiplier used by servers to scale key Raft timing parameters.
	// Tuning this affects the time it takes Vault to detect leader failures and to perform leader elections,
	// at the expense of requiring more network and CPU resources for better performance.
	// default: 0
	// +optional
	PerformanceMultiplier int64 `json:"performanceMultiplier,omitempty"`

	// This controls how many log entries are left in the log store on disk after a snapshot is made.
	// default: 10000
	// +optional
	TrailingLogs *int64 `json:"trailingLogs,omitempty"`

	// This controls the minimum number of raft commit entries between snapshots that are saved to disk.
	// default: 8192
	// +optional
	SnapshotThreshold *int64 `json:"snapshotThreshold,omitempty"`

	// This configures the maximum number of bytes for a raft entry. It applies to both Put operations and transactions.
	// default: 1048576
	// +optional
	MaxEntrySize *int64 `json:"maxEntrySize,omitempty"`

	// This is the interval after which autopilot will pick up any state changes.
	// default: ""
	// +optional
	AutopilotReconcileInterval string `json:"autopilotReconcileInterval,omitempty"`

	// Storage to specify how storage shall be used.
	Storage *core.PersistentVolumeClaimSpec `json:"storage,omitempty"`
}

// vault doc: https://openbao.org/docs/configuration/storage/cockroachdb/
//
// CockroachDBSpec defines configuration to set up CockroachDB as backend storage in vault.
// CockroachDB speaks the PostgreSQL wire protocol, but is configured as a dedicated backend
// (rather than reusing PostgreSQLSpec) so its schema and locking queries can account for
// CockroachDB-specific behavior.
type CockroachDBSpec struct {
	// Specifies the address of the CockroachDB host.
	// if DatabaseRef is set then Address will be generated from it
	// This must be set if DatabaseRef is empty, validate from ValidatingWebhook
	// +optional
	Address string `json:"address"`

	// Specifies the secret name containing username, password and, optionally, a
	// full connection_url to connect to the CockroachDB cluster.
	// secret data:
	//  - username=<value>
	//  - password=<value>
	//  - connection_url="postgresql://<username>:<password>@<host>:<port>/<db_name>?sslmode=<sslMode>"
	// +optional
	CredentialSecretRef *core.LocalObjectReference `json:"credentialSecretRef,omitempty"`

	// DatabaseRef contains the info of KubeDB managed Database
	// This will be used to generate the "Address" field
	// +optional
	DatabaseRef *appcat.AppReference `json:"databaseRef,omitempty"`

	// SSLMode for both standalone and clusters. [disable;require;verify-ca;verify-full]
	// +optional
	SSLMode PostgresSSLMode `json:"sslMode,omitempty"`

	// Specifies the name of the table in which to write Vault data.
	// +optional
	// +kubebuilder:default:="openbao_kv_store"
	Table string `json:"table,omitempty"`

	// Specifies the maximum number of concurrent requests to take place.
	// +optional
	MaxParallel int64 `json:"maxParallel,omitempty"`

	// Specifies the maximum number of concurrent interactive transactions to take place.
	// +optional
	TransactionMaxParallel int64 `json:"transactionMaxParallel,omitempty"`

	// When set, Vault will not attempt to create the storage table(s); they must already exist.
	// +optional
	SkipCreateTable bool `json:"skipCreateTable,omitempty"`

	// High Availability Parameter
	// Specifies if high availability mode is enabled. This is a boolean value, but it is specified as a string like "true" or "false".
	// +optional
	// +kubebuilder:default:="false"
	HAEnabled string `json:"haEnabled,omitempty"`

	// Specifies the name of the table to use for storing high availability information.
	// +optional
	// +kubebuilder:default:="openbao_ha_locks"
	HATable string `json:"haTable,omitempty"`
}

// vault doc: https://openbao.org/docs/configuration/storage/cassandra/
//
// CassandraSpec defines configuration to set up Cassandra as backend storage in vault
type CassandraSpec struct {
	// Specifies a comma-separated list of Cassandra hosts to connect to. All hosts must
	// listen on the same port; include the port in each host as "<host>:<port>" if it is
	// not the CQL native protocol default.
	Hosts string `json:"hosts"`

	// Specifies the keyspace that is used for storing the Vault data. The keyspace must
	// already exist, be reachable, and writable.
	// +optional
	// +kubebuilder:default:="vault"
	Keyspace string `json:"keyspace,omitempty"`

	// Specifies the table in the keyspace that is used for storing the Vault data. The
	// table must already exist.
	// +optional
	// +kubebuilder:default:="entries"
	Table string `json:"table,omitempty"`

	// Specifies the consistency level for read and write operations. Must be one of ANY,
	// ONE, TWO, THREE, QUORUM, ALL, LOCAL_QUORUM, EACH_QUORUM, or LOCAL_ONE.
	// +optional
	Consistency string `json:"consistency,omitempty"`

	// Specifies the CQL protocol version to use. Set to "3" or higher to use
	// username/password authentication with a proto version that requires it.
	// +optional
	ProtocolVersion string `json:"protocolVersion,omitempty"`

	// Specifies the secret name containing username and password to use for
	// authentication (PasswordAuthenticator). Requires protocolVersion of "2" or higher.
	// secret data:
	//  - username=<value>
	//  - password=<value>
	// +optional
	CredentialSecretRef *core.LocalObjectReference `json:"credentialSecretRef,omitempty"`

	// When set to a positive integer, enables Cassandra's SimpleRetryPolicy with the
	// given number of retries for queries that time out or fail.
	// +optional
	SimpleRetryPolicyRetries string `json:"simpleRetryPolicyRetries,omitempty"`

	// Specifies the timeout, in seconds, for the initial connection to the cluster.
	// +optional
	InitialConnectionTimeout string `json:"initialConnectionTimeout,omitempty"`

	// Specifies the timeout, in seconds, for individual queries against the cluster.
	// +optional
	ConnectionTimeout string `json:"connectionTimeout,omitempty"`

	// Set to enable a TLS connection to Cassandra.
	// +optional
	TLSEnabled bool `json:"tlsEnabled,omitempty"`

	// Specifies the secret name that contains a PEM-encoded certificate bundle (and,
	// optionally, a private key) used for the TLS connection to Cassandra.
	// secret data:
	//  - pem_bundle=<value>
	// +optional
	TLSSecretRef *core.LocalObjectReference `json:"tlsSecretRef,omitempty"`

	// Disables verification of the Cassandra server's certificate chain and host name.
	// Not recommended for production use.
	// +optional
	TLSSkipVerify bool `json:"tlsSkipVerify,omitempty"`

	// Specifies the minimum acceptable TLS version. One of tls10, tls11, tls12, or tls13.
	// +optional
	TLSMinVersion string `json:"tlsMinVersion,omitempty"`
}

// vault doc: https://openbao.org/docs/configuration/storage/zookeeper/
//
// ZookeeperSpec defines configuration to set up ZooKeeper as backend storage in vault
type ZookeeperSpec struct {
	// Specifies the addresses of the ZooKeeper instances as a comma-separated list.
	// +optional
	// +kubebuilder:default:="localhost:2181"
	Address string `json:"address,omitempty"`

	// Specifies the path in ZooKeeper's tree where Vault data will be stored.
	// +optional
	// +kubebuilder:default:="vault/"
	Path string `json:"path,omitempty"`

	// Specifies the ACL scheme:id applied to every znode Vault creates. Defaults to
	// "world:anyone", i.e. unrestricted access.
	// +optional
	ZnodeOwner string `json:"znodeOwner,omitempty"`

	// Specifies the secret name containing a scheme:auth pair passed to ZooKeeper's
	// AddAuth API immediately after connecting, so the client authenticates as a
	// specific principal.
	// secret data:
	//  - authInfo=<scheme:auth>
	// +optional
	AuthInfoSecretRef *core.LocalObjectReference `json:"authInfoSecretRef,omitempty"`

	// Enables a TLS connection to ZooKeeper.
	// +optional
	TLSEnabled bool `json:"tlsEnabled,omitempty"`

	// Specifies the secret name that contains ca.crt, tls.crt and tls.key for ZooKeeper
	// communication.
	// secret data:
	//  - ca.crt
	//  - tls.crt
	//  - tls.key
	// +optional
	TLSSecretRef *core.LocalObjectReference `json:"tlsSecretRef,omitempty"`

	// Disables verification of the ZooKeeper server's certificate chain and host name.
	// Not recommended for production use.
	// +optional
	TLSSkipVerify bool `json:"tlsSkipVerify,omitempty"`

	// When set, verifies the server's IP address against the certificate instead of its
	// DNS name. Only consulted when TLSSkipVerify is false.
	// +optional
	TLSVerifyIP bool `json:"tlsVerifyIP,omitempty"`

	// Specifies the minimum acceptable TLS version. One of tls10, tls11, tls12, or tls13.
	// +optional
	TLSMinVersion string `json:"tlsMinVersion,omitempty"`
}

// vault doc: https://openbao.org/docs/configuration/storage/couchdb/
//
// CouchDBSpec defines configuration to set up CouchDB as backend storage in vault
type CouchDBSpec struct {
	// Specifies the full URL to the CouchDB database to use, including the database
	// name. The database must already exist; Vault does not create it.
	Endpoint string `json:"endpoint"`

	// Specifies the secret name containing the CouchDB username and password to
	// connect with.
	// secret data:
	//  - username=<value>
	//  - password=<value>
	// +optional
	CredentialSecretRef *core.LocalObjectReference `json:"credentialSecretRef,omitempty"`

	// Specifies the maximum number of concurrent requests to CouchDB.
	// +optional
	MaxParallel int64 `json:"maxParallel,omitempty"`
}

// vault doc: https://openbao.org/docs/configuration/storage/mssql/
//
// MSSQLSpec defines configuration to set up Microsoft SQL Server as backend storage in vault
type MSSQLSpec struct {
	// Specifies the address of the MSSQL host.
	Server string `json:"server"`

	// Specifies the port of the MSSQL host. Defaults to the driver's standard port
	// (1433) when unset.
	// +optional
	Port string `json:"port,omitempty"`

	// Specifies the secret name containing the MSSQL username and password to connect
	// with.
	// secret data:
	//  - username=<value>
	//  - password=<value>
	// +optional
	CredentialSecretRef *core.LocalObjectReference `json:"credentialSecretRef,omitempty"`

	// Specifies the name of the database to use. Vault will attempt to create it if it
	// does not already exist.
	// +optional
	// +kubebuilder:default:="openbao"
	Database string `json:"database,omitempty"`

	// Specifies the name of the table in which to write Vault data. Vault will attempt
	// to create it if missing.
	// +optional
	// +kubebuilder:default:="openbao"
	Table string `json:"table,omitempty"`

	// Specifies the schema within the database that the table lives in. Vault will
	// attempt to create it if missing (requires permission to run CREATE SCHEMA).
	// +optional
	// +kubebuilder:default:="dbo"
	Schema string `json:"schema,omitempty"`

	// Specifies the application name to report to the server.
	// +optional
	// +kubebuilder:default:="openbao"
	AppName string `json:"appName,omitempty"`

	// Specifies the connection timeout, in seconds.
	// +optional
	// +kubebuilder:default:="30"
	ConnectionTimeout string `json:"connectionTimeout,omitempty"`

	// Specifies the driver's internal log level bitmask.
	// +optional
	LogLevel string `json:"logLevel,omitempty"`

	// Specifies the maximum number of concurrent requests to MSSQL.
	// +optional
	MaxParallel int64 `json:"maxParallel,omitempty"`
}

// vault doc: https://openbao.org/docs/configuration/storage/google-cloud-spanner/
//
// SpannerSpec defines configuration to set up Google Cloud Spanner as backend storage in vault
type SpannerSpec struct {
	// Specifies the full name of the Spanner database, in the form
	// projects/<project>/instances/<instance>/databases/<database>.
	Database string `json:"database"`

	// Specifies the name of the table to use for Vault data.
	// +optional
	// +kubebuilder:default:="Vault"
	Table string `json:"table,omitempty"`

	// High Availability Parameter
	// Specifies if high availability mode is enabled. This is a boolean value, but it is specified as a string like "true" or "false".
	// +optional
	// +kubebuilder:default:="false"
	HAEnabled string `json:"haEnabled,omitempty"`

	// Specifies the name of the table used for HA leader election. Defaults to the
	// value of Table suffixed with "HA".
	// +optional
	HATable string `json:"haTable,omitempty"`

	// Specifies the maximum number of concurrent requests to Spanner.
	// +optional
	MaxParallel int64 `json:"maxParallel,omitempty"`

	// Secret containing the Google application credential used to reach Spanner.
	// secret data:
	//  - sa.json:<value>
	// +optional
	CredentialSecretRef *core.LocalObjectReference `json:"credentialSecretRef,omitempty"`
}

// vault doc: https://openbao.org/docs/configuration/storage/aerospike/
//
// AerospikeSpec defines configuration to set up Aerospike as backend storage in vault
type AerospikeSpec struct {
	// Specifies the hostname of the Aerospike server to connect to, when HostList is
	// not set.
	// +optional
	// +kubebuilder:default:="127.0.0.1"
	Hostname string `json:"hostname,omitempty"`

	// Specifies the port of the Aerospike server to connect to, when HostList is not
	// set.
	// +optional
	// +kubebuilder:default:="3000"
	Port string `json:"port,omitempty"`

	// A comma-separated list of host[:port] entries describing the Aerospike cluster's
	// seed nodes. Takes precedence over Hostname/Port when set.
	// +optional
	HostList string `json:"hostList,omitempty"`

	// Specifies the Aerospike namespace to store data in. The namespace must already
	// exist on the server; Vault does not create it.
	// +optional
	// +kubebuilder:default:="test"
	Namespace string `json:"namespace,omitempty"`

	// Specifies the Aerospike set to store data in, within Namespace.
	// +optional
	Set string `json:"set,omitempty"`

	// Specifies the secret name containing username and password to authenticate to
	// the Aerospike cluster with, if it has security enabled.
	// secret data:
	//  - username=<value>
	//  - password=<value>
	// +optional
	CredentialSecretRef *core.LocalObjectReference `json:"credentialSecretRef,omitempty"`

	// Either INTERNAL (Aerospike's own credential store) or EXTERNAL (e.g. LDAP).
	// +optional
	// +kubebuilder:default:="INTERNAL"
	AuthMode string `json:"authMode,omitempty"`

	// If set, the client verifies it is connected to a cluster with this name.
	// +optional
	ClusterName string `json:"clusterName,omitempty"`

	// Socket connection timeout, in milliseconds.
	// +optional
	Timeout int64 `json:"timeout,omitempty"`

	// Idle connection timeout, in milliseconds. 0 disables idle connection trimming.
	// +optional
	IdleTimeout int64 `json:"idleTimeout,omitempty"`
}

// vault doc: https://openbao.org/docs/configuration/storage/oci/
//
// OCISpec defines configuration to set up OCI Object Storage as backend storage in vault
type OCISpec struct {
	// Specifies the name of the OCI Object Storage bucket to store data in. It must
	// already exist; Vault does not create it.
	BucketName string `json:"bucketName"`

	// Specifies the Object Storage namespace the bucket belongs to.
	NamespaceName string `json:"namespaceName"`

	// Specifies the OCI region to use. Defaults to the region configured in the
	// resolved OCI configuration provider.
	// +optional
	Region string `json:"region,omitempty"`

	// When true, authenticates using an OCI API-key configuration file (see
	// CredentialSecretRef). When false (the default), authenticates using instance
	// principal credentials, intended for Vault instances running on an OCI compute
	// instance.
	// +optional
	AuthTypeAPIKey bool `json:"authTypeAPIKey,omitempty"`

	// Specifies the secret name containing an OCI API-key configuration file (in the
	// same format as the default "~/.oci/config" file, including the referenced
	// private key). Only consulted when AuthTypeAPIKey is true.
	// secret data:
	//  - config=<value>
	// +optional
	CredentialSecretRef *core.LocalObjectReference `json:"credentialSecretRef,omitempty"`

	// Enables High Availability support. Requires LockBucketName to also be set.
	// +optional
	HAEnabled bool `json:"haEnabled,omitempty"`

	// Specifies the name of a second bucket used to store HA lock records. Required
	// when HAEnabled is true. Must already exist and should generally be a different
	// bucket than BucketName.
	// +optional
	LockBucketName string `json:"lockBucketName,omitempty"`
}

// vault doc: https://openbao.org/docs/configuration/storage/alicloudoss/
//
// AlicloudOSSSpec defines configuration to set up Alicloud OSS as backend storage in vault
type AlicloudOSSSpec struct {
	// Specifies the OSS endpoint to connect to, e.g. "http://oss-us-east-1.aliyuncs.com".
	Endpoint string `json:"endpoint"`

	// Specifies the name of the OSS bucket to store data in. It must already exist;
	// Vault does not create it.
	Bucket string `json:"bucket"`

	// Specifies the secret name containing the Alicloud access key ID and secret to
	// connect with.
	// secret data:
	//  - access_key=<value>
	//  - secret_key=<value>
	// +optional
	CredentialSecretRef *core.LocalObjectReference `json:"credentialSecretRef,omitempty"`

	// Specifies the maximum number of concurrent requests to OSS.
	// +optional
	MaxParallel int64 `json:"maxParallel,omitempty"`
}

// AzureKeyVault contain the fields that required to unseal vault using azure key vault
type AzureKeyVault struct {
	// Azure key vault url, for example https://myvault.vault.azure.net
	VaultBaseURL string `json:"vaultBaseURL"`

	// The cloud environment identifier
	// default: "AZUREPUBLICCLOUD"
	// +optional
	Cloud string `json:"cloud,omitempty"`

	// The AAD Tenant ID
	TenantID string `json:"tenantID"`

	// Specifies the name of secret containing client cert and client cert password
	// secret data:
	//  - client-cert:<value>
	// 	- client-cert-password: <value>
	// +optional
	TLSSecretRef *core.LocalObjectReference `json:"tlsSecretRef,omitempty"`

	// Specifies the name of secret containing client id and client secret of AAD application
	// secret data:
	//  - client-id:<value>
	//  - client-secret:<value>
	// +optional
	CredentialSecretRef *core.LocalObjectReference `json:"credentialSecretRef,omitempty"`

	// Use managed service identity for the virtual machine
	// +optional
	UseManagedIdentity bool `json:"useManagedIdentity,omitempty"`
}

// +kubebuilder:validation:Enum=kubernetes;aws;gcp;userpass;cert;azure;jwt;oidc
type AuthMethodType string

const (
	AuthTypeKubernetes AuthMethodType = "kubernetes"
	AuthTypeAws        AuthMethodType = "aws"
	AuthTypeGcp        AuthMethodType = "gcp"
	AuthTypeUserPass   AuthMethodType = "userpass"
	AuthTypeCert       AuthMethodType = "cert"
	AuthTypeAzure      AuthMethodType = "azure"
	AuthTypeJWT        AuthMethodType = "jwt"
	AuthTypeOIDC       AuthMethodType = "oidc"
)

// AuthMethod contains the information to enable vault auth method
// links: https://www.vaultproject.io/api/system/auth.html
type AuthMethod struct {
	//  Specifies the name of the authentication method type, such as "github" or "token".
	Type AuthMethodType `json:"type"`

	// Specifies the path in which to enable the auth method.
	// Default value is the same as the 'type'
	Path string `json:"path"`

	// Specifies a human-friendly description of the auth method.
	// +optional
	Description string `json:"description,omitempty"`

	// Kubernetes auth config
	KubernetesConfig *KubernetesConfig `json:"kubernetesConfig,omitempty"`

	// OIDC auth config
	OIDCConfig *JWTOIDCConfig `json:"oidcConfig,omitempty"`

	// JWT auth config
	JWTConfig *JWTOIDCConfig `json:"jwtConfig,omitempty"`

	// Specifies the name of the auth plugin to use based from the name in the plugin catalog.
	// Applies only to plugin methods.
	// +optional
	PluginName string `json:"pluginName,omitempty"`

	// Specifies if the auth method is a local only. Local auth methods are not replicated nor (if a secondary) removed by replication.
	// +optional
	Local bool `json:"local,omitempty"`
}

// +kubebuilder:validation:Enum=EnableSucceeded;EnableFailed;DisableSucceeded;DisableFailed
type AuthMethodEnableDisableStatus string

const (
	AuthMethodEnableSucceeded  AuthMethodEnableDisableStatus = "EnableSucceeded"
	AuthMethodEnableFailed     AuthMethodEnableDisableStatus = "EnableFailed"
	AuthMethodDisableSucceeded AuthMethodEnableDisableStatus = "DisableSucceeded"
	AuthMethodDisableFailed    AuthMethodEnableDisableStatus = "DisableFailed"
)

// AuthMethodStatus specifies the status of the auth method maintained by the auth method controller
type AuthMethodStatus struct {
	//  Specifies the name of the authentication method type, such as "github" or "token".
	Type AuthMethodType `json:"type"`

	// Specifies the path in which to enable the auth method.
	Path string `json:"path"`

	// Specifies whether auth method is enabled or not
	Status AuthMethodEnableDisableStatus `json:"status"`

	// Specifies the reason why failed to enable auth method
	// +optional
	Reason string `json:"reason,omitempty"`
}

type KubernetesConfig struct {
	// The default lease duration, specified as a string duration like "5s" or "30m".
	// +optional
	DefaultLeaseTTL string `json:"defaultLeaseTTL,omitempty"`

	// The maximum lease duration, specified as a string duration like "5s" or "30m".
	// +optional
	MaxLeaseTTL string `json:"maxLeaseTTL,omitempty"`

	// The name of the plugin in the plugin catalog to use.
	// +optional
	PluginName string `json:"pluginName,omitempty"`

	// List of keys that will not be HMAC'd by audit devices in the request data object.
	// +optional
	AuditNonHMACRequestKeys []string `json:"auditNonHMACRequestKeys,omitempty"`

	// List of keys that will not be HMAC'd by audit devices in the response data object.
	// +optional
	AuditNonHMACResponseKeys []string `json:"auditNonHMACResponseKeys,omitempty"`

	// Speficies whether to show this mount in the UI-specific listing endpoint.
	// +optional
	ListingVisibility string `json:"listingVisibility,omitempty"`

	// List of headers to whitelist and pass from the request to the backend.
	// +optional
	PassthroughRequestHeaders []string `json:"passthroughRequestHeaders,omitempty"`
}

type JWTOIDCConfig struct {
	// The default lease duration, specified as a string duration like "5s" or "30m".
	// +optional
	DefaultLeaseTTL string `json:"defaultLeaseTTL,omitempty"`

	// The maximum lease duration, specified as a string duration like "5s" or "30m".
	// +optional
	MaxLeaseTTL string `json:"maxLeaseTTL,omitempty"`

	// The name of the plugin in the plugin catalog to use.
	// +optional
	PluginName string `json:"pluginName,omitempty"`

	// List of keys that will not be HMAC'd by audit devices in the request data object.
	// +optional
	AuditNonHMACRequestKeys []string `json:"auditNonHMACRequestKeys,omitempty"`

	// List of keys that will not be HMAC'd by audit devices in the response data object.
	// +optional
	AuditNonHMACResponseKeys []string `json:"auditNonHMACResponseKeys,omitempty"`

	// Speficies whether to show this mount in the UI-specific listing endpoint.
	// +optional
	ListingVisibility string `json:"listingVisibility,omitempty"`

	// List of headers to whitelist and pass from the request to the backend.
	// +optional
	PassthroughRequestHeaders []string `json:"passthroughRequestHeaders,omitempty"`

	// CredentialSecretRef
	// +optional
	CredentialSecretRef *core.LocalObjectReference `json:"credentialSecretRef,omitempty"`

	// TLSSecretRef
	// +optional
	TLSSecretRef *core.LocalObjectReference `json:"tlsSecretRef,omitempty"`

	// common configuration parameters
	// The OIDC Discovery URL, without any .well-known component (base path). Cannot be used with "jwks_url" or "jwt_validation_pubkeys".
	// +optional
	OIDCDiscoveryURL string `json:"oidcDiscoveryURL,omitempty"`

	// The OAuth Client ID from the provider for OIDC roles.
	// +optional
	OIDCClientID string `json:"oidcClientID,omitempty"`

	// The response mode to be used in the OAuth2 request. Allowed values are "query" and "form_post". Defaults to "query".
	// If using Vault namespaces, and oidc_response_mode is "form_post", then "namespace_in_state" should be set to false.
	// +optional
	OIDCResponseMode string `json:"oidcResponseMode,omitempty"`

	// (comma-separated string, or array of strings: <optional>) - The response types to request.
	// Allowed values are "code" and "id_token". Defaults to "code". Note: "id_token" may only be used if "oidc_response_mode" is set to "form_post".
	// +optional
	OIDCResponseTypes string `json:"oidcResponseTypes,omitempty"`

	// The default role to use if none is provided during login
	// +optional
	DefaultRole string `json:"defaultRole,omitempty"`

	// Configuration options for provider-specific handling.
	// Providers with specific handling include: Azure, Google. The options are described in each provider's section in OIDC Provider Setup.
	// +optional
	ProviderConfig map[string]string `json:"providerConfig,omitempty"`

	// JWKS URL to use to authenticate signatures. Cannot be used with "oidc_discovery_url" or "jwt_validation_pubkeys".
	// +optional
	JWKSURL string `json:"jwksURL,omitempty"`

	// (comma-separated string, or array of strings: <optional>)
	// A list of PEM-encoded public keys to use to authenticate signatures locally. Cannot be used with "jwks_url" or "oidc_discovery_url".
	JWTValidationPubkeys []string `json:"jwtValidationPubkeys,omitempty"`

	// (comma-separated string, or array of strings: <optional>)
	// A list of supported signing algorithms. Defaults to [RS256] for OIDC roles. Defaults to all available algorithms for JWT roles.
	// +optional
	JWTSupportedAlgs []string `json:"jwtSupportedAlgs,omitempty"`

	// The value against which to match the iss claim in a JWT.
	// +optional
	BoundIssuer string `json:"boundIssuer,omitempty"`
}
