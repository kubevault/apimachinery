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

func TestGetVaultDeploymentMode(t *testing.T) {
	tests := []struct {
		name      string
		params    string
		wantType  DeploymentMode
		wantSpoke string
		wantErr   bool
	}{
		{
			name:     "no parameters means local",
			params:   "",
			wantType: DeploymentModeLocal,
		},
		{
			name:     "parameters without deploymentMode means local",
			params:   `{"apiVersion":"config.kubevault.com/v1alpha1","kind":"VaultServerConfiguration","path":"kubernetes"}`,
			wantType: DeploymentModeLocal,
		},
		{
			name:     "explicit Local",
			params:   `{"deploymentMode":"Local"}`,
			wantType: DeploymentModeLocal,
		},
		{
			name:      "RemoteRelay with spokeName",
			params:    `{"deploymentMode":"RemoteRelay","spokeName":"cluster-1"}`,
			wantType:  DeploymentModeRemoteRelay,
			wantSpoke: "cluster-1",
		},
		{
			name:    "legacy remote value rejected",
			params:  `{"deploymentMode":"remote","spokeName":"cluster-1"}`,
			wantErr: true,
		},
		{
			name:    "RemoteRelay without spokeName fails",
			params:  `{"deploymentMode":"RemoteRelay"}`,
			wantErr: true,
		},
		{
			name:    "unknown deploymentMode fails",
			params:  `{"deploymentMode":"Hybrid"}`,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			typ, spoke, err := GetVaultDeploymentMode(appBindingWithParams(tt.params))
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
