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
	appcat "kmodules.xyz/custom-resources/apis/appcatalog/v1alpha1"
)

const (
	ResourceKindRabbitMQRole = "RabbitMQRole"
	ResourceRabbitMQRole     = "rabbitmqrole"
	ResourceRabbitMQRoles    = "rabbitmqroles"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=rabbitmqroles,singular=rabbitmqrole,categories={vault,appscode,all}
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Status",type="string",JSONPath=".status.phase"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
type RabbitMQRole struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RabbitMQRoleSpec `json:"spec,omitempty"`
	Status            RoleStatus       `json:"status,omitempty"`
}

// RabbitMQRoleSpec describes a dynamic-role binding against the
// RabbitMQ database secret engine. The OpenBao
// `rabbitmq-database-plugin` (sigilr/openbao#8) provisions credentials
// via the RabbitMQ Management HTTP API (using rabbit-hole/v3) so
// `creationStatements` is a single-element JSON role document of the
// form `{"tags":"administrator","vhosts":{"/":{"configure":".*","write":".*","read":".*"}}}`.
// At least one of `tags` or `vhosts` must be set. Revocation is the
// plugin's default DELETE /api/users/<name>, naturally idempotent, so
// no `revocation_statements` field is exposed here.
// https://www.rabbitmq.com/access-control.html
type RabbitMQRoleSpec struct {
	// SecretEngineRef is the name of a Secret Engine
	SecretEngineRef core.LocalObjectReference `json:"secretEngineRef"`

	// CreationStatements is a JSON role document. Example:
	// ['{"tags":"administrator","vhosts":{"/":{"configure":".*","write":".*","read":".*"}}}']
	// At least one of `tags` or `vhosts` must be set.
	CreationStatements []string `json:"creationStatements"`

	// Specifies the TTL for the leases associated with this role.
	// Accepts time suffixed strings ("1h") or an integer number of seconds.
	// Defaults to system/engine default TTL time
	// +optional
	DefaultTTL string `json:"defaultTTL,omitempty"`

	// Specifies the maximum TTL for the leases associated with this role.
	// Accepts time suffixed strings ("1h") or an integer number of seconds.
	// Defaults to system/engine default TTL time.
	// +optional
	MaxTTL string `json:"maxTTL,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type RabbitMQRoleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	// Items is a list of RabbitMQRole objects
	Items []RabbitMQRole `json:"items,omitempty"`
}

// RabbitMQConfiguration defines a RabbitMQ app configuration. The
// OpenBao `rabbitmq-database-plugin` (sigilr/openbao#8) provisions
// credentials via the RabbitMQ Management HTTP API (using
// rabbit-hole/v3), so the connection payload uses `connection_url`
// (the RabbitMQ management HTTP base URL, e.g.
// `http://rabbitmq.demo.svc:15672`). Authentication is HTTP Basic Auth
// (username + password from the AppBinding secret). RabbitMQ is
// dynamic: NewUser/UpdateUser/DeleteUser are all supported. Revocation
// is the plugin's default DELETE /api/users/<name> (idempotent) so no
// `revocation_statements` field is exposed here.
// https://www.rabbitmq.com/access-control.html
type RabbitMQConfiguration struct {
	// Specifies the RabbitMQ database appbinding reference. The
	// AppBinding URL points at the RabbitMQ Management HTTP base URL
	// (e.g. `http://rabbitmq.demo.svc:15672`); the secret contributes
	// username/password used to authenticate against the management
	// API when the plugin issues credential operations.
	DatabaseRef appcat.AppReference `json:"databaseRef"`

	// Specifies the name of the plugin to use for this connection.
	// Default plugin:
	//  - for rabbitmq: rabbitmq-database-plugin
	// +optional
	PluginName string `json:"pluginName,omitempty"`

	// List of the roles allowed to use this connection.
	// Defaults to empty (no roles), if contains a "*" any role can use this connection.
	// +optional
	AllowedRoles []string `json:"allowedRoles,omitempty"`

	// PasswordPolicy is the optional name of a Vault password policy that
	// generates dynamic credentials.
	// +optional
	PasswordPolicy string `json:"passwordPolicy,omitempty"`

	// Insecure disables TLS verification when talking to the RabbitMQ
	// management endpoint. Useful for self-signed development
	// clusters; not recommended in production.
	// +kubebuilder:default:=false
	// +optional
	Insecure bool `json:"insecure,omitempty"`
}
