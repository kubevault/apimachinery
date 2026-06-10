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
	"testing"

	"k8s.io/apimachinery/pkg/runtime"
	appcat "kmodules.xyz/custom-resources/apis/appcatalog/v1alpha1"
)

func appBindingWithParams(params string) *appcat.AppBinding {
	ab := &appcat.AppBinding{}
	ab.Namespace = "demo"
	ab.Name = "vault"
	if params != "" {
		ab.Spec.Parameters = &runtime.RawExtension{Raw: []byte(params)}
	}
	return ab
}

func TestGetVaultDeployment(t *testing.T) {
	tests := []struct {
		name      string
		params    string
		wantType  VaultDeploymentType
		wantSpoke string
		wantErr   bool
	}{
		{
			name:     "no parameters means local",
			params:   "",
			wantType: VaultDeploymentLocal,
		},
		{
			name:     "parameters without vaultType means local",
			params:   `{"apiVersion":"config.kubevault.com/v1alpha1","kind":"VaultServerConfiguration","path":"kubernetes"}`,
			wantType: VaultDeploymentLocal,
		},
		{
			name:     "explicit Local",
			params:   `{"vaultType":"Local"}`,
			wantType: VaultDeploymentLocal,
		},
		{
			name:      "RemoteAgent with spokeName",
			params:    `{"vaultType":"RemoteAgent","spokeName":"cluster-1"}`,
			wantType:  VaultDeploymentRemoteAgent,
			wantSpoke: "cluster-1",
		},
		{
			name:      "legacy remote value normalized",
			params:    `{"vaultType":"remote","spokeName":"cluster-1"}`,
			wantType:  VaultDeploymentRemoteAgent,
			wantSpoke: "cluster-1",
		},
		{
			name:    "RemoteAgent without spokeName fails",
			params:  `{"vaultType":"RemoteAgent"}`,
			wantErr: true,
		},
		{
			name:    "unknown vaultType fails",
			params:  `{"vaultType":"Hybrid"}`,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			typ, spoke, err := GetVaultDeployment(appBindingWithParams(tt.params))
			if (err != nil) != tt.wantErr {
				t.Fatalf("err = %v, wantErr = %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			if typ != tt.wantType || spoke != tt.wantSpoke {
				t.Fatalf("got (%q, %q), want (%q, %q)", typ, spoke, tt.wantType, tt.wantSpoke)
			}
		})
	}
}
