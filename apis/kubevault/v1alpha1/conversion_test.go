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
	"encoding/json"
	"testing"

	"kubevault.dev/apimachinery/apis/kubevault/v1alpha2"

	diff "github.com/yudai/gojsondiff"
	v1 "k8s.io/api/core/v1"
	appcat "kmodules.xyz/custom-resources/apis/appcatalog/v1alpha1"
)

func TestConvert_v1alpha1_MySQLSpec_To_v1alpha2_MySQLSpec(t *testing.T) {
	testData := []struct {
		testName string
		spec     *MySQLSpec
	}{
		{
			testName: "test-0: should be successful",
			spec: &MySQLSpec{
				Address:              "address-v1alpha1",
				Database:             "database",
				Table:                "table",
				UserCredentialSecret: "",
				TLSCASecret:          "",
				MaxParallel:          0,
			},
		},
		{
			testName: "test-1: should be successful",
			spec: &MySQLSpec{
				Address:              "address-v1alpha1",
				Database:             "database",
				Table:                "table",
				UserCredentialSecret: "secret",
				TLSCASecret:          "tls",
				MaxParallel:          10,
			},
		},
	}

	for idx := range testData {
		test := testData[idx]
		t.Run(test.testName, func(t *testing.T) {
			mysqlv1alph1 := test.spec
			mysqlv1alph2 := &v1alpha2.MySQLSpec{}
			if err := Convert_v1alpha1_MySQLSpec_To_v1alpha2_MySQLSpec(mysqlv1alph1, mysqlv1alph2, nil); err != nil {
				t.Error(err)
			}

			newmysqlv1alpha1 := &MySQLSpec{}
			if err := Convert_v1alpha2_MySQLSpec_To_v1alpha1_MySQLSpec(mysqlv1alph2, newmysqlv1alpha1, nil); err != nil {
				t.Error(err)
			}

			old, err := json.Marshal(mysqlv1alph1)
			if err != nil {
				t.Error(err)
			}

			new, err := json.Marshal(newmysqlv1alpha1)
			if err != nil {
				t.Error(err)
			}

			differ := diff.New()
			d, err := differ.Compare(old, new)
			if err != nil {
				t.Error(err)
			}

			if d.Modified() {
				t.Error("modified")
			}
		})
	}
}

func TestConvert_v1alpha2_MySQLSpec_To_v1alpha1_MySQLSpec(t *testing.T) {
	testData := []struct {
		testName string
		spec     *v1alpha2.MySQLSpec
	}{
		{
			testName: "test-0: should be successful",
			spec: &v1alpha2.MySQLSpec{
				Address:  "address",
				Database: "database",
				Table:    "table",
				CredentialSecretRef: &v1.LocalObjectReference{
					Name: "cred",
				},
				TLSSecretRef: &v1.LocalObjectReference{
					Name: "tls",
				},
				MaxParallel: 0,
				DatabaseRef: &appcat.AppReference{
					Name:      "database",
					Namespace: "ns",
				},
				PlaintextCredentialTransmission: "",
				MaxIdleConnection:               0,
				MaxConnectionLifetime:           0,
				HAEnabled:                       "",
				LockTable:                       "",
			},
		},
		{
			testName: "test-1: should be successful",
			spec: &v1alpha2.MySQLSpec{
				Address:                         "address",
				Database:                        "database",
				Table:                           "table",
				CredentialSecretRef:             nil,
				TLSSecretRef:                    nil,
				MaxParallel:                     0,
				DatabaseRef:                     nil,
				PlaintextCredentialTransmission: "",
				MaxIdleConnection:               0,
				MaxConnectionLifetime:           0,
				HAEnabled:                       "",
				LockTable:                       "",
			},
		},
	}

	for idx := range testData {
		test := testData[idx]
		t.Run(test.testName, func(t *testing.T) {
			mysqlv1alph2 := test.spec
			mysqlv1alph1 := &MySQLSpec{}
			if err := Convert_v1alpha2_MySQLSpec_To_v1alpha1_MySQLSpec(mysqlv1alph2, mysqlv1alph1, nil); err != nil {
				t.Error(err)
			}

			newmysqlv1alpha2 := &v1alpha2.MySQLSpec{}
			if err := Convert_v1alpha1_MySQLSpec_To_v1alpha2_MySQLSpec(mysqlv1alph1, newmysqlv1alpha2, nil); err != nil {
				t.Error(err)
			}

			old, err := json.Marshal(mysqlv1alph2)
			if err != nil {
				t.Error(err)
			}

			new, err := json.Marshal(newmysqlv1alpha2)
			if err != nil {
				t.Error(err)
			}

			differ := diff.New()
			d, err := differ.Compare(old, new)
			if err != nil {
				t.Error(err)
			}

			if d.Modified() {
				t.Error("modified")
			}
		})
	}
}

func TestConvert_v1alpha1_PostgreSQLSpec_To_v1alpha2_PostgreSQLSpec(t *testing.T) {
	testData := []struct {
		testName string
		spec     *PostgreSQLSpec
	}{
		{
			testName: "test-0: should be successful",
			spec: &PostgreSQLSpec{
				ConnectionURLSecret: "secret",
				Table:               "table",
				MaxParallel:         0,
			},
		},
		{
			testName: "test-1: should be successful",
			spec: &PostgreSQLSpec{
				ConnectionURLSecret: "",
				Table:               "",
				MaxParallel:         0,
			},
		},
	}

	for idx := range testData {
		test := testData[idx]
		t.Run(test.testName, func(t *testing.T) {
			pgv1alph1 := test.spec
			pgv1alph2 := &v1alpha2.PostgreSQLSpec{}
			if err := Convert_v1alpha1_PostgreSQLSpec_To_v1alpha2_PostgreSQLSpec(pgv1alph1, pgv1alph2, nil); err != nil {
				t.Error(err)
			}

			newpgv1alpha1 := &PostgreSQLSpec{}
			if err := Convert_v1alpha2_PostgreSQLSpec_To_v1alpha1_PostgreSQLSpec(pgv1alph2, newpgv1alpha1, nil); err != nil {
				t.Error(err)
			}

			old, err := json.Marshal(pgv1alph1)
			if err != nil {
				t.Error(err)
			}

			new, err := json.Marshal(newpgv1alpha1)
			if err != nil {
				t.Error(err)
			}

			differ := diff.New()
			d, err := differ.Compare(old, new)
			if err != nil {
				t.Error(err)
			}

			if d.Modified() {
				t.Error("modified")
			}
		})
	}
}

func TestConvert_v1alpha2_PostgreSQLSpec_To_v1alpha1_PostgreSQLSpec(t *testing.T) {
	testData := []struct {
		testName string
		spec     *v1alpha2.PostgreSQLSpec
	}{
		{
			testName: "test-0: should be successful",
			spec: &v1alpha2.PostgreSQLSpec{
				Address: "addr",
				CredentialSecretRef: &v1.LocalObjectReference{
					Name: "cred",
				},
				DatabaseRef: &appcat.AppReference{
					Name:      "db",
					Namespace: "ns",
				},
				SSLMode:           "",
				Table:             "",
				MaxParallel:       0,
				MaxIdleConnection: 0,
				HAEnabled:         "",
				HATable:           "",
			},
		},
		{
			testName: "test-1: should be successful",
			spec: &v1alpha2.PostgreSQLSpec{
				Address:             "addr",
				CredentialSecretRef: nil,
				DatabaseRef:         nil,
				SSLMode:             "",
				Table:               "table",
				MaxParallel:         10,
				MaxIdleConnection:   10,
				HAEnabled:           "",
				HATable:             "",
			},
		},
	}

	for idx := range testData {
		test := testData[idx]
		t.Run(test.testName, func(t *testing.T) {
			pgv1alph2 := test.spec
			pgv1alph1 := &PostgreSQLSpec{}
			if err := Convert_v1alpha2_PostgreSQLSpec_To_v1alpha1_PostgreSQLSpec(pgv1alph2, pgv1alph1, nil); err != nil {
				t.Error(err)
			}

			newpgv1alpha2 := &v1alpha2.PostgreSQLSpec{}
			if err := Convert_v1alpha1_PostgreSQLSpec_To_v1alpha2_PostgreSQLSpec(pgv1alph1, newpgv1alpha2, nil); err != nil {
				t.Error(err)
			}

			old, err := json.Marshal(pgv1alph2)
			if err != nil {
				t.Error(err)
			}

			new, err := json.Marshal(newpgv1alpha2)
			if err != nil {
				t.Error(err)
			}

			differ := diff.New()
			d, err := differ.Compare(old, new)
			if err != nil {
				t.Error(err)
			}

			if d.Modified() {
				t.Error("modified")
			}
		})
	}
}

func TestConvert_v1alpha1_AwsKmsSsmSpec_To_v1alpha2_AwsKmsSsmSpec(t *testing.T) {
	testData := []struct {
		testName string
		spec     *AwsKmsSsmSpec
	}{
		{
			testName: "test-0: should be successful",
			spec: &AwsKmsSsmSpec{
				KmsKeyID:         "kmskeyid-v1alpha1",
				SsmKeyPrefix:     "ssmkeyprefix-v1alpha1",
				Region:           "region-v1alpha1",
				CredentialSecret: "",
				Endpoint:         "endpoint",
			},
		},
		{
			testName: "test-1: should be successful",
			spec: &AwsKmsSsmSpec{
				KmsKeyID:         "kmskeyid-v1alpha1",
				SsmKeyPrefix:     "ssmkeyprefix-v1alpha1",
				Region:           "region-v1alpha1",
				CredentialSecret: "cred",
				Endpoint:         "endpoint",
			},
		},
	}

	for idx := range testData {
		test := testData[idx]
		t.Run(test.testName, func(t *testing.T) {
			awskmsv1alph1 := test.spec
			awskmsv1alph2 := &v1alpha2.AwsKmsSsmSpec{}
			if err := Convert_v1alpha1_AwsKmsSsmSpec_To_v1alpha2_AwsKmsSsmSpec(awskmsv1alph1, awskmsv1alph2, nil); err != nil {
				t.Error(err)
			}

			newawskmsv1alpha1 := &AwsKmsSsmSpec{}
			if err := Convert_v1alpha2_AwsKmsSsmSpec_To_v1alpha1_AwsKmsSsmSpec(awskmsv1alph2, newawskmsv1alpha1, nil); err != nil {
				t.Error(err)
			}

			old, err := json.Marshal(awskmsv1alph1)
			if err != nil {
				t.Error(err)
			}

			new, err := json.Marshal(newawskmsv1alpha1)
			if err != nil {
				t.Error(err)
			}

			differ := diff.New()
			d, err := differ.Compare(old, new)
			if err != nil {
				t.Error(err)
			}

			if d.Modified() {
				t.Error("modified")
			}
		})
	}
}

func TestConvert_v1alpha2_AwsKmsSsmSpec_To_v1alpha1_AwsKmsSsmSpec(t *testing.T) {
	testData := []struct {
		testName string
		spec     *v1alpha2.AwsKmsSsmSpec
	}{
		{
			testName: "test-0: should be successful",
			spec: &v1alpha2.AwsKmsSsmSpec{
				KmsKeyID:            "kmskeyid-v1alpha2",
				SsmKeyPrefix:        "ssmkeyprefix-v1alpha2",
				Region:              "region-v1alpha2",
				CredentialSecretRef: nil,
				Endpoint:            "endpoint",
			},
		},
		{
			testName: "test-1: should be successful",
			spec: &v1alpha2.AwsKmsSsmSpec{
				KmsKeyID:     "kmskeyid-v1alpha2",
				SsmKeyPrefix: "ssmkeyprefix-v1alpha2",
				Region:       "region-v1alpha2",
				CredentialSecretRef: &v1.LocalObjectReference{
					Name: "cred",
				},
				Endpoint: "endpoint",
			},
		},
	}

	for idx := range testData {
		test := testData[idx]
		t.Run(test.testName, func(t *testing.T) {
			awskmsv1alph2 := test.spec
			awskmsv1alph1 := &AwsKmsSsmSpec{}
			if err := Convert_v1alpha2_AwsKmsSsmSpec_To_v1alpha1_AwsKmsSsmSpec(awskmsv1alph2, awskmsv1alph1, nil); err != nil {
				t.Error(err)
			}

			newawskmsv1alpha2 := &v1alpha2.AwsKmsSsmSpec{}
			if err := Convert_v1alpha1_AwsKmsSsmSpec_To_v1alpha2_AwsKmsSsmSpec(awskmsv1alph1, newawskmsv1alpha2, nil); err != nil {
				t.Error(err)
			}

			old, err := json.Marshal(awskmsv1alph2)
			if err != nil {
				t.Error(err)
			}

			new, err := json.Marshal(newawskmsv1alpha2)
			if err != nil {
				t.Error(err)
			}

			differ := diff.New()
			d, err := differ.Compare(old, new)
			if err != nil {
				t.Error(err)
			}

			if d.Modified() {
				t.Error("modified")
			}
		})
	}
}

func TestConvert_v1alpha1_AzureKeyVault_To_v1alpha2_AzureKeyVault(t *testing.T) {
	testData := []struct {
		testName string
		spec     *AzureKeyVault
	}{
		{
			testName: "test-0: should be successful",
			spec: &AzureKeyVault{
				VaultBaseURL:       "url",
				Cloud:              "cloud",
				TenantID:           "id",
				ClientCertSecret:   "client-secret",
				AADClientSecret:    "aad-secret",
				UseManagedIdentity: false,
			},
		},
		{
			testName: "test-1: should be successful",
			spec: &AzureKeyVault{
				VaultBaseURL:       "",
				Cloud:              "",
				TenantID:           "",
				ClientCertSecret:   "",
				AADClientSecret:    "",
				UseManagedIdentity: true,
			},
		},
	}

	for idx := range testData {
		test := testData[idx]
		t.Run(test.testName, func(t *testing.T) {
			azurev1alph1 := test.spec
			azurev1alph2 := &v1alpha2.AzureKeyVault{}
			if err := Convert_v1alpha1_AzureKeyVault_To_v1alpha2_AzureKeyVault(azurev1alph1, azurev1alph2, nil); err != nil {
				t.Error(err)
			}

			newazurev1alpha1 := &AzureKeyVault{}
			if err := Convert_v1alpha2_AzureKeyVault_To_v1alpha1_AzureKeyVault(azurev1alph2, newazurev1alpha1, nil); err != nil {
				t.Error(err)
			}

			old, err := json.Marshal(azurev1alph1)
			if err != nil {
				t.Error(err)
			}

			new, err := json.Marshal(newazurev1alpha1)
			if err != nil {
				t.Error(err)
			}

			differ := diff.New()
			d, err := differ.Compare(old, new)
			if err != nil {
				t.Error(err)
			}

			if d.Modified() {
				t.Error("modified")
			}
		})
	}
}

func TestConvert_v1alpha2_AzureKeyVault_To_v1alpha1_AzureKeyVault(t *testing.T) {
	testData := []struct {
		testName string
		spec     *v1alpha2.AzureKeyVault
	}{
		{
			testName: "test-0: should be successful",
			spec: &v1alpha2.AzureKeyVault{
				VaultBaseURL: "url",
				Cloud:        "cloud",
				TenantID:     "tenant",
				TLSSecretRef: &v1.LocalObjectReference{
					Name: "tls",
				},
				CredentialSecretRef: &v1.LocalObjectReference{
					Name: "cred",
				},
				UseManagedIdentity: false,
			},
		},
		{
			testName: "test-1: should be successful",
			spec: &v1alpha2.AzureKeyVault{
				VaultBaseURL:        "",
				Cloud:               "",
				TenantID:            "",
				TLSSecretRef:        nil,
				CredentialSecretRef: nil,
				UseManagedIdentity:  false,
			},
		},
	}

	for idx := range testData {
		test := testData[idx]
		t.Run(test.testName, func(t *testing.T) {
			azurev1alph2 := test.spec
			azurev1alph1 := &AzureKeyVault{}
			if err := Convert_v1alpha2_AzureKeyVault_To_v1alpha1_AzureKeyVault(azurev1alph2, azurev1alph1, nil); err != nil {
				t.Error(err)
			}

			newazurev1alpha2 := &v1alpha2.AzureKeyVault{}
			if err := Convert_v1alpha1_AzureKeyVault_To_v1alpha2_AzureKeyVault(azurev1alph1, newazurev1alpha2, nil); err != nil {
				t.Error(err)
			}

			old, err := json.Marshal(azurev1alph2)
			if err != nil {
				t.Error(err)
			}

			new, err := json.Marshal(newazurev1alpha2)
			if err != nil {
				t.Error(err)
			}

			differ := diff.New()
			d, err := differ.Compare(old, new)
			if err != nil {
				t.Error(err)
			}

			if d.Modified() {
				t.Error("modified")
			}
		})
	}
}
