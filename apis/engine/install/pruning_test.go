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

package install

import (
	"testing"

	"kubevault.dev/apimachinery/apis/engine/fuzzer"
	"kubevault.dev/apimachinery/apis/engine/v1alpha1"

	clientsetscheme "k8s.io/client-go/kubernetes/scheme"
	crdfuzz "kmodules.xyz/crd-schema-fuzz"
)

func TestPruneTypes(t *testing.T) {
	Install(clientsetscheme.Scheme)

	// CRD v1
	if crd := (v1alpha1.SecretEngine{}).CustomResourceDefinition(); crd.V1 != nil {
		crdfuzz.SchemaFuzzTestForV1CRD(t, clientsetscheme.Scheme, crd.V1, fuzzer.Funcs)
	}
	if crd := (v1alpha1.SecretRoleBinding{}).CustomResourceDefinition(); crd.V1 != nil {
		crdfuzz.SchemaFuzzTestForV1CRD(t, clientsetscheme.Scheme, crd.V1, fuzzer.Funcs)
	}
	if crd := (v1alpha1.SecretAccessRequest{}).CustomResourceDefinition(); crd.V1 != nil {
		crdfuzz.SchemaFuzzTestForV1CRD(t, clientsetscheme.Scheme, crd.V1, fuzzer.Funcs)
	}
	if crd := (v1alpha1.AWSRole{}).CustomResourceDefinition(); crd.V1 != nil {
		crdfuzz.SchemaFuzzTestForV1CRD(t, clientsetscheme.Scheme, crd.V1, fuzzer.Funcs)
	}
	if crd := (v1alpha1.AzureRole{}).CustomResourceDefinition(); crd.V1 != nil {
		crdfuzz.SchemaFuzzTestForV1CRD(t, clientsetscheme.Scheme, crd.V1, fuzzer.Funcs)
	}
	if crd := (v1alpha1.GCPRole{}).CustomResourceDefinition(); crd.V1 != nil {
		crdfuzz.SchemaFuzzTestForV1CRD(t, clientsetscheme.Scheme, crd.V1, fuzzer.Funcs)
	}
	if crd := (v1alpha1.MongoDBRole{}).CustomResourceDefinition(); crd.V1 != nil {
		crdfuzz.SchemaFuzzTestForV1CRD(t, clientsetscheme.Scheme, crd.V1, fuzzer.Funcs)
	}
	if crd := (v1alpha1.MySQLRole{}).CustomResourceDefinition(); crd.V1 != nil {
		crdfuzz.SchemaFuzzTestForV1CRD(t, clientsetscheme.Scheme, crd.V1, fuzzer.Funcs)
	}
	if crd := (v1alpha1.PostgresRole{}).CustomResourceDefinition(); crd.V1 != nil {
		crdfuzz.SchemaFuzzTestForV1CRD(t, clientsetscheme.Scheme, crd.V1, fuzzer.Funcs)
	}
	if crd := (v1alpha1.RedisRole{}).CustomResourceDefinition(); crd.V1 != nil {
		crdfuzz.SchemaFuzzTestForV1CRD(t, clientsetscheme.Scheme, crd.V1, fuzzer.Funcs)
	}
}
