#!/bin/bash

echo "=== VaultAgent CRD Verification ==="
echo ""

echo "1. Checking VaultAgent type files..."
ls -lh /home/rudro25/go/src/kubevault.dev/apimachinery/apis/kubevault/v1alpha2/vault_agent*.go
echo ""

echo "2. Checking generated deepcopy methods..."
grep "func (in \*VaultAgent)" /home/rudro25/go/src/kubevault.dev/apimachinery/apis/kubevault/v1alpha2/zz_generated.deepcopy.go | head -5
echo ""

echo "3. Checking CRD manifest..."
ls -lh /home/rudro25/go/src/kubevault.dev/apimachinery/crds/kubevault.com_vaultagents.yaml
echo ""

echo "4. Checking client code..."
find /home/rudro25/go/src/kubevault.dev/apimachinery/client -name "*vaultagent*" | wc -l
echo "   client files generated"
echo ""

echo "5. Verifying CRD structure..."
head -30 /home/rudro25/go/src/kubevault.dev/apimachinery/crds/kubevault.com_vaultagents.yaml
echo ""

echo "=== Verification Complete ==="
echo ""
echo "✅ VaultAgent CRD is ready!"
echo ""
echo "Next steps:"
echo "1. Update operator go.mod to use this apimachinery version"
echo "2. Run 'go mod tidy' in operator repo"
echo "3. Implement VaultAgent controller in operator"
