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
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	ResourceKindDocumentDBRole = "DocumentDBRole"
	ResourceDocumentDBRole     = "documentdbrole"
	ResourceDocumentDBRoles    = "documentdbroles"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=documentdbroles,singular=documentdbrole,categories={vault,appscode,all}
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Status",type="string",JSONPath=".status.phase"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
type DocumentDBRole struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DocumentDBRoleSpec `json:"spec,omitempty"`
	Status            RoleStatus         `json:"status,omitempty"`
}

// DocumentDBRoleSpec describes a dynamic-role binding against the
// DocumentDB database secret engine. The OpenBao `documentdb-database-plugin`
// (sigilr/openbao#9) reuses the mongo-driver to talk to a DocumentDB
// gateway over the MongoDB wire protocol, so `creation_statements` is
// the same JSON role-document format MongoDB uses.
type DocumentDBRoleSpec struct {
	// SecretEngineRef is the name of a Secret Engine
	SecretEngineRef core.LocalObjectReference `json:"secretEngineRef"`

	// Specifies the TTL for the leases associated with this role.
	// Accepts time suffixed strings ("1h") or an integer number of seconds.
	// Defaults to system/engine default TTL time
	DefaultTTL string `json:"defaultTTL,omitempty"`

	// Specifies the maximum TTL for the leases associated with this role.
	// Accepts time suffixed strings ("1h") or an integer number of seconds.
	// Defaults to system/engine default TTL time.
	MaxTTL string `json:"maxTTL,omitempty"`

	// Specifies the database statements executed to create and configure a user.
	// DocumentDB reuses the MongoDB driver, so each entry is a JSON role
	// document (e.g. `{ "db": "admin", "roles": [{ "role": "readWrite" }] }`).
	CreationStatements []string `json:"creationStatements"`

	// Specifies the database statements to be executed to revoke a user.
	RevocationStatements []string `json:"revocationStatements,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type DocumentDBRoleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	// Items is a list of DocumentDBRole objects
	Items []DocumentDBRole `json:"items,omitempty"`
}
