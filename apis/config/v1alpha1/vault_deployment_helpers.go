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
	"encoding/json"
	"fmt"

	appcat "kmodules.xyz/custom-resources/apis/appcatalog/v1alpha1"
)

// GetVaultDeployment is the single sanctioned way to determine whether a
// Vault AppBinding points at a Local vault or a hub vault reached through
// the OpenBao spoke agent (RemoteAgent). It returns the normalized
// deployment type and, for RemoteAgent, the spoke name.
//
// Rules:
//   - missing parameters, or parameters without vaultType => Local
//   - vaultType "Local"                                   => Local
//   - vaultType "RemoteAgent" or legacy "remote"          => RemoteAgent;
//     spokeName must be present, otherwise an error is returned
//   - any other value                                     => error
//
// The `vault-type: remote` AppBinding label is a list-filter convenience
// only and must never be used for routing decisions.
func GetVaultDeployment(ab *appcat.AppBinding) (VaultDeploymentType, string, error) {
	if ab == nil {
		return "", "", fmt.Errorf("AppBinding is nil")
	}
	if ab.Spec.Parameters == nil || len(ab.Spec.Parameters.Raw) == 0 {
		return VaultDeploymentLocal, "", nil
	}

	var cfg VaultServerConfiguration
	if err := json.Unmarshal(ab.Spec.Parameters.Raw, &cfg); err != nil {
		return "", "", fmt.Errorf("failed to parse parameters of AppBinding %s/%s: %w", ab.Namespace, ab.Name, err)
	}

	switch cfg.VaultType {
	case "", VaultDeploymentLocal:
		return VaultDeploymentLocal, "", nil
	case VaultDeploymentRemoteAgent, vaultDeploymentLegacyRemote:
		if cfg.SpokeName == "" {
			return "", "", fmt.Errorf("AppBinding %s/%s has vaultType %q but no spokeName", ab.Namespace, ab.Name, cfg.VaultType)
		}
		return VaultDeploymentRemoteAgent, cfg.SpokeName, nil
	default:
		return "", "", fmt.Errorf("AppBinding %s/%s has unknown vaultType %q", ab.Namespace, ab.Name, cfg.VaultType)
	}
}
