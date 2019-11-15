/*
Copyright The KubeVault Authors.

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
	"kubevault.dev/operator/api/crds"

	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"sigs.k8s.io/yaml"
)

func (_ GCPAccessKeyRequest) CustomResourceDefinition() *apiextensions.CustomResourceDefinition {
	data := crds.MustAsset("engine.kubevault.com_gcpaccesskeyrequests.yaml")
	var out apiextensions.CustomResourceDefinition
	utilruntime.Must(yaml.Unmarshal(data, &out))
	return &out
}

func (d GCPAccessKeyRequest) IsValid() error {
	return nil
}
