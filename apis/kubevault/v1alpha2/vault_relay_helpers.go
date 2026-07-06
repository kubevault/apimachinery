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
	"fmt"

	"kubevault.dev/apimachinery/apis/kubevault"
	"kubevault.dev/apimachinery/crds"

	"kmodules.xyz/client-go/apiextensions"
	meta_util "kmodules.xyz/client-go/meta"
)

func (*VaultRelay) Hub() {}

func (va VaultRelay) CustomResourceDefinition() *apiextensions.CustomResourceDefinition {
	return crds.MustCustomResourceDefinition(SchemeGroupVersion.WithResource(ResourceVaultRelays))
}

func (va VaultRelay) ResourceFQN() string {
	return fmt.Sprintf("%s.%s", ResourceVaultRelays, kubevault.GroupName)
}

func (va VaultRelay) GetKey() string {
	return va.Namespace + "/" + va.Name
}

func (va VaultRelay) OffshootName() string {
	return va.Name
}

func (va VaultRelay) ServiceAccountName() string {
	return va.Name
}

func (va VaultRelay) PodName() string {
	return meta_util.NameWithSuffix(va.Name, "agent")
}

func (va VaultRelay) AppBindingName() string {
	return meta_util.NameWithSuffix(va.Name, "hub-vault")
}

func (va VaultRelay) OffshootSelectors() map[string]string {
	return map[string]string{
		meta_util.NameLabelKey:      va.ResourceFQN(),
		meta_util.InstanceLabelKey:  va.Name,
		meta_util.ManagedByLabelKey: kubevault.GroupName,
	}
}

func (va VaultRelay) OffshootLabels() map[string]string {
	return meta_util.FilterKeys("kubevault.com", va.OffshootSelectors(), va.Labels)
}

func (va VaultRelay) IsValid() error {
	return nil
}

// GetGRPCPort returns the gRPC port with default value
func (va VaultRelay) GetGRPCPort() int32 {
	if va.Spec.HubVaultRef.GRPCPort == 0 {
		return 50053
	}
	return va.Spec.HubVaultRef.GRPCPort
}

// GetImage returns the spoke-agent image with default value
func (va VaultRelay) GetImage() string {
	if va.Spec.Image == "" {
		return "ghcr.io/kubevault/spoke-agent:latest"
	}
	return va.Spec.Image
}

// SetDefaults sets default values for VaultRelay
func (va *VaultRelay) SetDefaults() {
	if va.Spec.HubVaultRef.GRPCPort == 0 {
		va.Spec.HubVaultRef.GRPCPort = 50053
	}

	if va.Spec.Reconnect == nil {
		va.Spec.Reconnect = &ReconnectConfig{
			Enabled:           true,
			BackoffSeconds:    5,
			MaxBackoffSeconds: 300,
		}
	}
}
