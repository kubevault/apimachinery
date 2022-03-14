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

package main

import (
	"io/ioutil"
	"os"
	"path/filepath"

	cataloginstall "kubevault.dev/apimachinery/apis/catalog/install"
	catalogv1alpha1 "kubevault.dev/apimachinery/apis/catalog/v1alpha1"
	engineinstall "kubevault.dev/apimachinery/apis/engine/install"
	enginev1alpha1 "kubevault.dev/apimachinery/apis/engine/v1alpha1"
	vaultinstall "kubevault.dev/apimachinery/apis/kubevault/install"
	vaultv1alpha1 "kubevault.dev/apimachinery/apis/kubevault/v1alpha1"
	vaultv1alpha2 "kubevault.dev/apimachinery/apis/kubevault/v1alpha2"
	policyinstall "kubevault.dev/apimachinery/apis/policy/install"
	policyv1alpha1 "kubevault.dev/apimachinery/apis/policy/v1alpha1"

	"github.com/go-openapi/spec"
	gort "gomodules.xyz/runtime"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/klog/v2"
	"k8s.io/kube-openapi/pkg/common"
	"kmodules.xyz/client-go/openapi"
)

func generateSwaggerJson() {
	var (
		Scheme = runtime.NewScheme()
		Codecs = serializer.NewCodecFactory(Scheme)
	)

	vaultinstall.Install(Scheme)
	cataloginstall.Install(Scheme)
	policyinstall.Install(Scheme)
	engineinstall.Install(Scheme)

	apispec, err := openapi.RenderOpenAPISpec(openapi.Config{
		Scheme: Scheme,
		Codecs: Codecs,
		Info: spec.InfoProps{
			Title:   "KubeVault",
			Version: "v0.3.0",
			Contact: &spec.ContactInfo{
				Name:  "AppsCode Inc.",
				URL:   "https://appscode.com",
				Email: "kubevault@appscode.com",
			},
			License: &spec.License{
				Name: "Apache 2.0",
				URL:  "https://www.apache.org/licenses/LICENSE-2.0.html",
			},
		},
		OpenAPIDefinitions: []common.GetOpenAPIDefinitions{
			vaultv1alpha2.GetOpenAPIDefinitions,
			vaultv1alpha1.GetOpenAPIDefinitions,
			catalogv1alpha1.GetOpenAPIDefinitions,
			policyv1alpha1.GetOpenAPIDefinitions,
			enginev1alpha1.GetOpenAPIDefinitions,
		},
		//nolint:govet
		Resources: []openapi.TypeInfo{
			{vaultv1alpha2.SchemeGroupVersion, vaultv1alpha2.ResourceVaultServers, vaultv1alpha2.ResourceKindVaultServer, true},
			{vaultv1alpha1.SchemeGroupVersion, vaultv1alpha1.ResourceVaultServers, vaultv1alpha1.ResourceKindVaultServer, true},
			{catalogv1alpha1.SchemeGroupVersion, catalogv1alpha1.ResourceVaultServerVersions, catalogv1alpha1.ResourceKindVaultServerVersion, false},
			{policyv1alpha1.SchemeGroupVersion, policyv1alpha1.ResourceVaultPolicies, policyv1alpha1.ResourceKindVaultPolicy, true},
			{policyv1alpha1.SchemeGroupVersion, policyv1alpha1.ResourceVaultPolicyBindings, policyv1alpha1.ResourceKindVaultPolicyBinding, true},
			{enginev1alpha1.SchemeGroupVersion, enginev1alpha1.ResourceSecretEngines, enginev1alpha1.ResourceKindSecretEngine, true},
			{enginev1alpha1.SchemeGroupVersion, enginev1alpha1.ResourceSecretRoleBindings, enginev1alpha1.ResourceKindSecretRoleBinding, true},
			{enginev1alpha1.SchemeGroupVersion, enginev1alpha1.ResourceSecretAccessRequests, enginev1alpha1.ResourceKindSecretAccessRequest, true},
			{enginev1alpha1.SchemeGroupVersion, enginev1alpha1.ResourceAWSRoles, enginev1alpha1.ResourceKindAWSRole, true},
			{enginev1alpha1.SchemeGroupVersion, enginev1alpha1.ResourceAzureRoles, enginev1alpha1.ResourceKindAzureRole, true},
			{enginev1alpha1.SchemeGroupVersion, enginev1alpha1.ResourceGCPRoles, enginev1alpha1.ResourceKindGCPRole, true},
			{enginev1alpha1.SchemeGroupVersion, enginev1alpha1.ResourceMongoDBRoles, enginev1alpha1.ResourceKindMongoDBRole, true},
			{enginev1alpha1.SchemeGroupVersion, enginev1alpha1.ResourceMySQLRoles, enginev1alpha1.ResourceKindMySQLRole, true},
			{enginev1alpha1.SchemeGroupVersion, enginev1alpha1.ResourcePostgresRoles, enginev1alpha1.ResourceKindPostgresRole, true},
		},
	})
	if err != nil {
		klog.Fatal(err)
	}

	filename := gort.GOPath() + "/src/kubevault.dev/apimachinery/api/openapi-spec/swagger.json"
	err = os.MkdirAll(filepath.Dir(filename), 0o755)
	if err != nil {
		klog.Fatal(err)
	}
	err = ioutil.WriteFile(filename, []byte(apispec), 0o644)
	if err != nil {
		klog.Fatal(err)
	}
}

func main() {
	generateSwaggerJson()
}
