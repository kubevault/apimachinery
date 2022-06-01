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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "kubevault.dev/apimachinery/client/clientset/versioned/typed/engine/v1alpha1"

	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeEngineV1alpha1 struct {
	*testing.Fake
}

func (c *FakeEngineV1alpha1) AWSRoles(namespace string) v1alpha1.AWSRoleInterface {
	return &FakeAWSRoles{c, namespace}
}

func (c *FakeEngineV1alpha1) AzureRoles(namespace string) v1alpha1.AzureRoleInterface {
	return &FakeAzureRoles{c, namespace}
}

func (c *FakeEngineV1alpha1) ElasticsearchRoles(namespace string) v1alpha1.ElasticsearchRoleInterface {
	return &FakeElasticsearchRoles{c, namespace}
}

func (c *FakeEngineV1alpha1) GCPRoles(namespace string) v1alpha1.GCPRoleInterface {
	return &FakeGCPRoles{c, namespace}
}

func (c *FakeEngineV1alpha1) MariaDBRoles(namespace string) v1alpha1.MariaDBRoleInterface {
	return &FakeMariaDBRoles{c, namespace}
}

func (c *FakeEngineV1alpha1) MongoDBRoles(namespace string) v1alpha1.MongoDBRoleInterface {
	return &FakeMongoDBRoles{c, namespace}
}

func (c *FakeEngineV1alpha1) MySQLRoles(namespace string) v1alpha1.MySQLRoleInterface {
	return &FakeMySQLRoles{c, namespace}
}

func (c *FakeEngineV1alpha1) PostgresRoles(namespace string) v1alpha1.PostgresRoleInterface {
	return &FakePostgresRoles{c, namespace}
}

func (c *FakeEngineV1alpha1) SecretAccessRequests(namespace string) v1alpha1.SecretAccessRequestInterface {
	return &FakeSecretAccessRequests{c, namespace}
}

func (c *FakeEngineV1alpha1) SecretEngines(namespace string) v1alpha1.SecretEngineInterface {
	return &FakeSecretEngines{c, namespace}
}

func (c *FakeEngineV1alpha1) SecretRoleBindings(namespace string) v1alpha1.SecretRoleBindingInterface {
	return &FakeSecretRoleBindings{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeEngineV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
