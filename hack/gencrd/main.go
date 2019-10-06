package main

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/appscode/go/log"
	gort "github.com/appscode/go/runtime"
	"github.com/go-openapi/spec"
	"github.com/golang/glog"
	crd_api "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/kube-openapi/pkg/common"
	crdutils "kmodules.xyz/client-go/apiextensions/v1beta1"
	"kmodules.xyz/client-go/openapi"
	"kubevault.dev/operator/apis"
	cataloginstall "kubevault.dev/operator/apis/catalog/install"
	catalogv1alpha1 "kubevault.dev/operator/apis/catalog/v1alpha1"
	engineinstall "kubevault.dev/operator/apis/engine/install"
	enginev1alpha1 "kubevault.dev/operator/apis/engine/v1alpha1"
	vaultinstall "kubevault.dev/operator/apis/kubevault/install"
	vaultv1alpha1 "kubevault.dev/operator/apis/kubevault/v1alpha1"
	policyinstall "kubevault.dev/operator/apis/policy/install"
	policyv1alpha1 "kubevault.dev/operator/apis/policy/v1alpha1"
)

func generateCRDDefinitions() {
	apis.EnableStatusSubresource = true

	filename := gort.GOPath() + "/src/kubevault.dev/operator/apis/kubevault/v1alpha1/crds.yaml"
	os.Remove(filename)

	err := os.MkdirAll(filepath.Join(gort.GOPath(), "/src/kubevault.dev/operator/api/crds"), 0755)
	if err != nil {
		log.Fatal(err)
	}

	crds := []*crd_api.CustomResourceDefinition{
		vaultv1alpha1.VaultServer{}.CustomResourceDefinition(),
		catalogv1alpha1.VaultServerVersion{}.CustomResourceDefinition(),
		policyv1alpha1.VaultPolicy{}.CustomResourceDefinition(),
		policyv1alpha1.VaultPolicyBinding{}.CustomResourceDefinition(),
		enginev1alpha1.SecretEngine{}.CustomResourceDefinition(),
		enginev1alpha1.AWSRole{}.CustomResourceDefinition(),
		enginev1alpha1.AWSAccessKeyRequest{}.CustomResourceDefinition(),
		enginev1alpha1.AzureRole{}.CustomResourceDefinition(),
		enginev1alpha1.AzureAccessKeyRequest{}.CustomResourceDefinition(),
		enginev1alpha1.GCPRole{}.CustomResourceDefinition(),
		enginev1alpha1.GCPAccessKeyRequest{}.CustomResourceDefinition(),
		enginev1alpha1.DatabaseAccessRequest{}.CustomResourceDefinition(),
		enginev1alpha1.MongoDBRole{}.CustomResourceDefinition(),
		enginev1alpha1.MySQLRole{}.CustomResourceDefinition(),
		enginev1alpha1.PostgresRole{}.CustomResourceDefinition(),
	}
	for _, crd := range crds {
		filename := filepath.Join(gort.GOPath(), "/src/kubevault.dev/operator/api/crds", crd.Spec.Names.Singular+".yaml")
		f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
		if err != nil {
			log.Fatal(err)
		}
		crdutils.MarshallCrd(f, crd, "yaml")
		f.Close()
	}
}

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
			Version: "v0.2.0",
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
			vaultv1alpha1.GetOpenAPIDefinitions,
			catalogv1alpha1.GetOpenAPIDefinitions,
			policyv1alpha1.GetOpenAPIDefinitions,
			enginev1alpha1.GetOpenAPIDefinitions,
		},
		Resources: []openapi.TypeInfo{
			{vaultv1alpha1.SchemeGroupVersion, vaultv1alpha1.ResourceVaultServers, vaultv1alpha1.ResourceKindVaultServer, true},
			{catalogv1alpha1.SchemeGroupVersion, catalogv1alpha1.ResourceVaultServerVersions, catalogv1alpha1.ResourceKindVaultServerVersion, false},
			{policyv1alpha1.SchemeGroupVersion, policyv1alpha1.ResourceVaultPolicies, policyv1alpha1.ResourceKindVaultPolicy, true},
			{policyv1alpha1.SchemeGroupVersion, policyv1alpha1.ResourceVaultPolicyBindings, policyv1alpha1.ResourceKindVaultPolicyBinding, true},
			{enginev1alpha1.SchemeGroupVersion, enginev1alpha1.ResourceSecretEngines, enginev1alpha1.ResourceKindSecretEngine, true},
			{enginev1alpha1.SchemeGroupVersion, enginev1alpha1.ResourceAWSRoles, enginev1alpha1.ResourceKindAWSRole, true},
			{enginev1alpha1.SchemeGroupVersion, enginev1alpha1.ResourceAWSAccessKeyRequests, enginev1alpha1.ResourceKindAWSAccessKeyRequest, true},
			{enginev1alpha1.SchemeGroupVersion, enginev1alpha1.ResourceAzureRoles, enginev1alpha1.ResourceKindAzureRole, true},
			{enginev1alpha1.SchemeGroupVersion, enginev1alpha1.ResourceAzureAccessKeyRequests, enginev1alpha1.ResourceKindAzureAccessKeyRequest, true},
			{enginev1alpha1.SchemeGroupVersion, enginev1alpha1.ResourceGCPRoles, enginev1alpha1.ResourceKindGCPRole, true},
			{enginev1alpha1.SchemeGroupVersion, enginev1alpha1.ResourceGCPAccessKeyRequests, enginev1alpha1.ResourceKindGCPAccessKeyRequest, true},
			{enginev1alpha1.SchemeGroupVersion, enginev1alpha1.ResourceDatabaseAccessRequests, enginev1alpha1.ResourceKindDatabaseAccessRequest, true},
			{enginev1alpha1.SchemeGroupVersion, enginev1alpha1.ResourceMongoDBRoles, enginev1alpha1.ResourceKindMongoDBRole, true},
			{enginev1alpha1.SchemeGroupVersion, enginev1alpha1.ResourceMySQLRoles, enginev1alpha1.ResourceKindMySQLRole, true},
			{enginev1alpha1.SchemeGroupVersion, enginev1alpha1.ResourcePostgresRoles, enginev1alpha1.ResourceKindPostgresRole, true},
		},
	})
	if err != nil {
		glog.Fatal(err)
	}

	filename := gort.GOPath() + "/src/kubevault.dev/operator/api/openapi-spec/swagger.json"
	err = os.MkdirAll(filepath.Dir(filename), 0755)
	if err != nil {
		glog.Fatal(err)
	}
	err = ioutil.WriteFile(filename, []byte(apispec), 0644)
	if err != nil {
		glog.Fatal(err)
	}
}

func main() {
	// generateCRDDefinitions()
	generateSwaggerJson()
}
