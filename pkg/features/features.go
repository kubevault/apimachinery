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

package features

import (
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/component-base/featuregate"
)

const (
	// Every feature gate should add method here following this template:
	//
	// MyFeature featuregate.Feature = "MyFeature"
	//
	// Feature gates should be listed in alphabetical, case-sensitive
	// (upper before any lower case character) order. This reduces the risk
	// of code conflicts because changes are more likely to be scattered
	// across the file.
	//
	// These gate the installation of a database engine's SecretEngine
	// role CRD (e.g. PostgresRole). They deliberately exclude the
	// cloud-IAM/PKI secret engines (AWS, Azure, GCP, PKI), which are not
	// database plugins and are always installed.

	// Enables the DB2 secret engine.
	DB2 featuregate.Feature = "DB2"

	// Enables the Druid secret engine.
	Druid featuregate.Feature = "Druid"

	// Enables the Elasticsearch secret engine.
	Elasticsearch featuregate.Feature = "Elasticsearch"

	// Enables the HanaDB secret engine.
	HanaDB featuregate.Feature = "HanaDB"

	// Enables the Hazelcast secret engine.
	Hazelcast featuregate.Feature = "Hazelcast"

	// Enables the Ignite secret engine.
	Ignite featuregate.Feature = "Ignite"

	// Enables the Kafka secret engine.
	Kafka featuregate.Feature = "Kafka"

	// Enables the MariaDB secret engine.
	MariaDB featuregate.Feature = "MariaDB"

	// Enables the Memcached secret engine.
	Memcached featuregate.Feature = "Memcached"

	// Enables the Milvus secret engine.
	Milvus featuregate.Feature = "Milvus"

	// Enables the MongoDB secret engine.
	MongoDB featuregate.Feature = "MongoDB"

	// Enables the Microsoft SQL Server secret engine.
	MSSQLServer featuregate.Feature = "MSSQLServer"

	// Enables the MySQL secret engine.
	MySQL featuregate.Feature = "MySQL"

	// Enables the Neo4j secret engine.
	Neo4j featuregate.Feature = "Neo4j"

	// Enables the Oracle Database secret engine.
	Oracle featuregate.Feature = "Oracle"

	// Enables the Postgres secret engine.
	Postgres featuregate.Feature = "Postgres"

	// Enables the Qdrant secret engine.
	Qdrant featuregate.Feature = "Qdrant"

	// Enables the RabbitMQ secret engine.
	RabbitMQ featuregate.Feature = "RabbitMQ"

	// Enables the Redis secret engine.
	Redis featuregate.Feature = "Redis"

	// Enables the Solr secret engine.
	Solr featuregate.Feature = "Solr"

	// Enables the Weaviate secret engine.
	Weaviate featuregate.Feature = "Weaviate"

	// Enables the ZooKeeper secret engine.
	ZooKeeper featuregate.Feature = "ZooKeeper"
)

func init() {
	runtime.Must(DefaultMutableFeatureGate.Add(defaultKubeVaultFeatureGates))
}

// defaultKubeVaultFeatureGates consists of all known KubeVault database
// secret-engine feature keys. To add a new database engine, define a key
// for it above and add it here. The features will be available throughout
// KubeVault binaries (operator, crd-manager).
var defaultKubeVaultFeatureGates = map[featuregate.Feature]featuregate.FeatureSpec{
	DB2:           {Default: false, PreRelease: featuregate.Alpha},
	Druid:         {Default: false, PreRelease: featuregate.Alpha},
	Elasticsearch: {Default: false, PreRelease: featuregate.Alpha},
	HanaDB:        {Default: false, PreRelease: featuregate.Alpha},
	Hazelcast:     {Default: false, PreRelease: featuregate.Alpha},
	Ignite:        {Default: false, PreRelease: featuregate.Alpha},
	Kafka:         {Default: false, PreRelease: featuregate.Alpha},
	MariaDB:       {Default: false, PreRelease: featuregate.Alpha},
	Memcached:     {Default: false, PreRelease: featuregate.Alpha},
	Milvus:        {Default: false, PreRelease: featuregate.Alpha},
	MongoDB:       {Default: false, PreRelease: featuregate.Alpha},
	MSSQLServer:   {Default: false, PreRelease: featuregate.Alpha},
	MySQL:         {Default: true, PreRelease: featuregate.GA},
	Neo4j:         {Default: false, PreRelease: featuregate.Alpha},
	Oracle:        {Default: false, PreRelease: featuregate.Alpha},
	Postgres:      {Default: true, PreRelease: featuregate.GA},
	Qdrant:        {Default: false, PreRelease: featuregate.Alpha},
	RabbitMQ:      {Default: false, PreRelease: featuregate.Alpha},
	Redis:         {Default: false, PreRelease: featuregate.Alpha},
	Solr:          {Default: false, PreRelease: featuregate.Alpha},
	Weaviate:      {Default: false, PreRelease: featuregate.Alpha},
	ZooKeeper:     {Default: false, PreRelease: featuregate.Alpha},
}
