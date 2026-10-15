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
	core "k8s.io/api/core/v1"
	appcat "kmodules.xyz/custom-resources/apis/appcatalog/v1alpha1"
)

// vault doc: https://openbao.org/docs/configuration/storage/mssql/
//
// MSSQLSpec defines configuration to set up Microsoft SQL Server as backend storage in vault
type MSSQLSpec struct {
	// Specifies the address of the MSSQL host.
	// +optional
	Server string `json:"server,omitempty"`

	// Specifies the port of the MSSQL host. Defaults to the driver's standard port
	// (1433) when unset.
	// +optional
	Port string `json:"port,omitempty"`

	// Specifies the secret name containing the MSSQL username and password to connect
	// with.
	// secret data:
	//  - username=<value>
	//  - password=<value>
	// +optional
	CredentialSecretRef *core.LocalObjectReference `json:"credentialSecretRef,omitempty"`

	// Specifies the name of the database to use. Vault will attempt to create it if it
	// does not already exist.
	// +optional
	// +kubebuilder:default:="openbao"
	Database string `json:"database,omitempty"`

	// Specifies the name of the table in which to write Vault data. Vault will attempt
	// to create it if missing.
	// +optional
	// +kubebuilder:default:="openbao"
	Table string `json:"table,omitempty"`

	// Specifies the schema within the database that the table lives in. Vault will
	// attempt to create it if missing (requires permission to run CREATE SCHEMA).
	// +optional
	// +kubebuilder:default:="dbo"
	Schema string `json:"schema,omitempty"`

	// Specifies the application name to report to the server.
	// +optional
	// +kubebuilder:default:="openbao"
	AppName string `json:"appName,omitempty"`

	// Specifies the connection timeout, in seconds.
	// +optional
	// +kubebuilder:default:="30"
	ConnectionTimeout string `json:"connectionTimeout,omitempty"`

	// Specifies the driver's internal log level bitmask.
	// +optional
	LogLevel string `json:"logLevel,omitempty"`

	// Specifies the maximum number of concurrent requests to MSSQL.
	// +optional
	MaxParallel int64 `json:"maxParallel,omitempty"`

	// DatabaseRef contains the info of a KubeDB managed MSSQLServer.
	// If set, Server will be generated from the referenced AppBinding.
	// +optional
	DatabaseRef *appcat.AppReference `json:"databaseRef,omitempty"`
}
