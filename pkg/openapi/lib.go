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

package openapi

import (
	vaulitapi "kubevault.dev/apimachinery/apis/kubevault/v1alpha2"

	"k8s.io/apimachinery/pkg/runtime"
	openapinamer "k8s.io/apiserver/pkg/endpoints/openapi"
	genericapiserver "k8s.io/apiserver/pkg/server"
)

func ConfigureOpenAPI(scheme *runtime.Scheme, serverConfig *genericapiserver.RecommendedConfig) {
	ignorePrefixes := []string{
		"/swaggerapi",

		"/apis/mutators.engine.kubevault.com/v1alpha1",
		"/apis/mutators.engine.kubevault.com/v1alpha1/secretaccessrequestwebhooks",

		"/apis/mutators.kubevault.com/v1alpha1",
		"/apis/mutators.kubevault.com/v1alpha1/vaultserverwebhooks",

		"/apis/mutators.policy.kubevault.com/v1alpha1",
		"/apis/mutators.policy.kubevault.com/v1alpha1/vaultpolicybindingwebhooks",

		"/apis/validators.engine.kubevault.com/v1alpha1",
		"/apis/validators.engine.kubevault.com/v1alpha1/secretaccessrequestwebhooks",

		"/apis/validators.kubevault.com/v1alpha1",
		"/apis/validators.kubevault.com/v1alpha1/vaultserverwebhooks",
	}

	serverConfig.OpenAPIConfig = genericapiserver.DefaultOpenAPIConfig(vaulitapi.GetOpenAPIDefinitions, openapinamer.NewDefinitionNamer(scheme))
	serverConfig.OpenAPIConfig.Info.Title = "kubevault-webhook-server"
	serverConfig.OpenAPIConfig.Info.Version = vaulitapi.SchemeGroupVersion.Version
	serverConfig.OpenAPIConfig.IgnorePrefixes = ignorePrefixes

	serverConfig.OpenAPIV3Config = genericapiserver.DefaultOpenAPIV3Config(vaulitapi.GetOpenAPIDefinitions, openapinamer.NewDefinitionNamer(scheme))
	serverConfig.OpenAPIV3Config.Info.Title = "kubevault-webhook-server"
	serverConfig.OpenAPIV3Config.Info.Version = vaulitapi.SchemeGroupVersion.Version
	serverConfig.OpenAPIV3Config.IgnorePrefixes = ignorePrefixes
}
