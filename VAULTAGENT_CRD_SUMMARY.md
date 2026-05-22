# VaultAgent CRD Implementation Summary

## Files Created in apimachinery

### 1. Core Type Definitions
**File:** `apis/kubevault/v1alpha2/vault_agent_types.go`

**Defines:**
- `VaultAgent` CRD with spec and status
- `VaultAgentSpec` - configuration for spoke agent
- `HubVaultReference` - hub vault connection details
- `VaultAgentTLSConfig` - TLS settings
- `ReconnectConfig` - auto-reconnect settings
- `VaultAgentStatus` - runtime status
- `VaultAgentPhase` - lifecycle phases (Pending, Connected, Disconnected, Error)

**Key Fields:**
```yaml
spec:
  hubVaultRef:
    name: vault3
    namespace: demo
    address: "http://10.2.0.88:30820"
    grpcPort: 50053
  spokeName: "spoke-cluster-1"
  image: "ghcr.io/kubevault/spoke-agent:latest"
  tls:
    enabled: false
    caSecret: vault-ca
    certSecret: spoke-cert
  reconnect:
    enabled: true
    backoffSeconds: 5
    maxBackoffSeconds: 300
  podTemplate: {}
```

### 2. Helper Functions
**File:** `apis/kubevault/v1alpha2/vault_agent_helpers.go`

**Functions:**
- `OffshootName()` - returns resource name
- `ServiceAccountName()` - returns SA name
- `PodName()` - returns spoke-agent pod name
- `AppBindingName()` - returns AppBinding name for hub vault
- `OffshootSelectors()` - returns label selectors
- `OffshootLabels()` - returns labels
- `GetGRPCPort()` - returns gRPC port with default (50053)
- `GetImage()` - returns image with default
- `SetDefaults()` - sets default values

### 3. Registration
**File:** `apis/kubevault/v1alpha2/register.go`

**Changes:**
- Added `VaultAgent` and `VaultAgentList` to scheme registration

### 4. Constants
**File:** `apis/kubevault/v1alpha2/types.go`

**Added:**
- `VaultAgentContainerName = "spoke-agent"`

## Generated Files

### 1. DeepCopy Methods
**File:** `apis/kubevault/v1alpha2/zz_generated.deepcopy.go`

Auto-generated DeepCopy methods for:
- `VaultAgent`
- `VaultAgentList`
- `VaultAgentSpec`
- `VaultAgentStatus`
- `HubVaultReference`
- `VaultAgentTLSConfig`
- `ReconnectConfig`

### 2. CRD Manifest
**File:** `crds/kubevault.com_vaultagents.yaml`

Complete Kubernetes CRD definition with:
- OpenAPI v3 schema validation
- Subresource: status
- Print columns: Spoke, Status, Age
- Short name: `va`
- Categories: vault, appscode, all

### 3. Client Code
**Generated in:** `client/`

- **Clientset:** `client/clientset/versioned/typed/kubevault/v1alpha2/vaultagent.go`
  - CRUD operations for VaultAgent
  - Status subresource operations

- **Lister:** `client/listers/kubevault/v1alpha2/vaultagent.go`
  - List VaultAgents with label selectors
  - Get VaultAgent by namespace/name

- **Informer:** `client/informers/externalversions/kubevault/v1alpha2/vaultagent.go`
  - Watch VaultAgent changes
  - Cache for efficient access

- **Fake Client:** `client/clientset/versioned/typed/kubevault/v1alpha2/fake/fake_vaultagent.go`
  - For unit testing

## Usage Example

### Create VaultAgent
```yaml
apiVersion: kubevault.com/v1alpha2
kind: VaultAgent
metadata:
  name: hub-vault-agent
  namespace: spoke-ns
spec:
  hubVaultRef:
    name: vault3
    namespace: demo
    address: "http://10.2.0.88:30820"
    grpcPort: 50053
  spokeName: "spoke-cluster-1"
  reconnect:
    enabled: true
    backoffSeconds: 5
    maxBackoffSeconds: 300
```

### Apply CRD
```bash
kubectl apply -f crds/kubevault.com_vaultagents.yaml
```

### Create VaultAgent Instance
```bash
kubectl apply -f vaultagent.yaml
```

### Check Status
```bash
kubectl get vaultagent -n spoke-ns
kubectl get va -n spoke-ns  # using short name
kubectl describe vaultagent hub-vault-agent -n spoke-ns
```

## Next Steps for Operator

1. **Update operator go.mod** to use new apimachinery version
2. **Create VaultAgent controller** in operator repo:
   - `operator/pkg/controller/vault_agent.go`
3. **Implement reconciliation logic:**
   - Deploy spoke-agent pod
   - Create AppBinding for hub vault
   - Update status with connection state
4. **Modify SecretEngine controller** to detect remote vault via AppBinding labels

## Integration Points

### AppBinding Creation
The VaultAgent controller will create an AppBinding like:
```yaml
apiVersion: appcatalog.appscode.com/v1alpha1
kind: AppBinding
metadata:
  name: hub-vault-appbinding
  namespace: spoke-ns
  labels:
    vault-type: remote  # ← Key label for detection
spec:
  type: kubevault.com/vault
  clientConfig:
    url: http://10.2.0.88:30820
    service:
      name: vault3
      namespace: demo
  parameters:
    spokeName: spoke-cluster-1
    grpcPort: 50053
    vaultType: remote
```

### SecretEngine Detection
```go
// In SecretEngine controller
appBinding, _ := getAppBinding(se.Spec.VaultRef)
vaultType := appBinding.Labels["vault-type"]
isRemote := vaultType == "remote"

if isRemote {
    // Add spoke_name to database config
    // Use remote-postgres-proxy plugin
}
```

## Summary

✅ **Completed:**
- VaultAgent CRD type definitions
- Helper functions
- Code generation (deepcopy, clients, informers)
- CRD manifest
- Registration in scheme

🔄 **Next (in operator repo):**
- VaultAgent controller implementation
- Spoke-agent pod deployment logic
- AppBinding creation logic
- SecretEngine remote detection logic
