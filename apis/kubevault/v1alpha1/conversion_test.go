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
	"fmt"
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

func TestConvert_v1alpha1_AzureSpec_To_v1alpha2_AzureSpec(t *testing.T) {
	testData := []struct {
		testName string
		spec     *AzureSpec
	}{
		{
			testName: "test-0: should be successful",
			spec: &AzureSpec{
				AccountName:      "acc",
				AccountKeySecret: "secret",
				Container:        "container",
				MaxParallel:      0,
			},
		},
		{
			testName: "test-1: should be successful",
			spec: &AzureSpec{
				AccountName:      "",
				AccountKeySecret: "",
				Container:        "",
				MaxParallel:      10,
			},
		},
	}

	for idx := range testData {
		test := testData[idx]
		t.Run(test.testName, func(t *testing.T) {
			azurev1alph1 := test.spec
			azurev1alph2 := &v1alpha2.AzureSpec{}
			if err := Convert_v1alpha1_AzureSpec_To_v1alpha2_AzureSpec(azurev1alph1, azurev1alph2, nil); err != nil {
				t.Error(err)
			}

			newazurev1alpha1 := &AzureSpec{}
			if err := Convert_v1alpha2_AzureSpec_To_v1alpha1_AzureSpec(azurev1alph2, newazurev1alpha1, nil); err != nil {
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

func TestConvert_v1alpha2_AzureSpec_To_v1alpha1_AzureSpec(t *testing.T) {
	testData := []struct {
		testName string
		spec     *v1alpha2.AzureSpec
	}{
		{
			testName: "test-0: should be successful",
			spec: &v1alpha2.AzureSpec{
				AccountName: "acc",
				CredentialSecretRef: &v1.LocalObjectReference{
					Name: "cred",
				},
				Container:   "",
				MaxParallel: 0,
			},
		},
		{
			testName: "test-1: should be successful",
			spec: &v1alpha2.AzureSpec{
				AccountName:         "",
				CredentialSecretRef: nil,
				Container:           "",
				MaxParallel:         10,
			},
		},
	}

	for idx := range testData {
		test := testData[idx]
		t.Run(test.testName, func(t *testing.T) {
			azurev1alph2 := test.spec
			azurev1alph1 := &AzureSpec{}
			if err := Convert_v1alpha2_AzureSpec_To_v1alpha1_AzureSpec(azurev1alph2, azurev1alph1, nil); err != nil {
				t.Error(err)
			}

			newazurev1alpha2 := &v1alpha2.AzureSpec{}
			if err := Convert_v1alpha1_AzureSpec_To_v1alpha2_AzureSpec(azurev1alph1, newazurev1alpha2, nil); err != nil {
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

func TestConvert_v1alpha1_GoogleKmsGcsSpec_To_v1alpha2_GoogleKmsGcsSpec(t *testing.T) {
	testData := []struct {
		testName string
		spec     *GoogleKmsGcsSpec
	}{
		{
			testName: "test-0: should be successful",
			spec: &GoogleKmsGcsSpec{
				KmsCryptoKey:     "kms",
				KmsKeyRing:       "ring",
				KmsLocation:      "location",
				KmsProject:       "project",
				Bucket:           "bucket",
				CredentialSecret: "secret",
			},
		},
		{
			testName: "test-1: should be successful",
			spec: &GoogleKmsGcsSpec{
				KmsCryptoKey:     "",
				KmsKeyRing:       "",
				KmsLocation:      "",
				KmsProject:       "",
				Bucket:           "",
				CredentialSecret: "",
			},
		},
	}

	for idx := range testData {
		test := testData[idx]
		t.Run(test.testName, func(t *testing.T) {
			gcsv1alph1 := test.spec
			gcsv1alph2 := &v1alpha2.GoogleKmsGcsSpec{}
			if err := Convert_v1alpha1_GoogleKmsGcsSpec_To_v1alpha2_GoogleKmsGcsSpec(gcsv1alph1, gcsv1alph2, nil); err != nil {
				t.Error(err)
			}

			newgcsv1alpha1 := &GoogleKmsGcsSpec{}
			if err := Convert_v1alpha2_GoogleKmsGcsSpec_To_v1alpha1_GoogleKmsGcsSpec(gcsv1alph2, newgcsv1alpha1, nil); err != nil {
				t.Error(err)
			}

			old, err := json.Marshal(gcsv1alph1)
			if err != nil {
				t.Error(err)
			}

			new, err := json.Marshal(newgcsv1alpha1)
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

func TestConvert_v1alpha2_GoogleKmsGcsSpec_To_v1alpha1_GoogleKmsGcsSpec(t *testing.T) {
	testData := []struct {
		testName string
		spec     *v1alpha2.GoogleKmsGcsSpec
	}{
		{
			testName: "test-0: should be successful",
			spec: &v1alpha2.GoogleKmsGcsSpec{
				KmsCryptoKey: "kms",
				KmsKeyRing:   "ring",
				KmsLocation:  "location",
				KmsProject:   "project",
				Bucket:       "bucket",
				CredentialSecretRef: &v1.LocalObjectReference{
					Name: "secret",
				},
			},
		},
		{
			testName: "test-1: should be successful",
			spec: &v1alpha2.GoogleKmsGcsSpec{
				KmsCryptoKey:        "",
				KmsKeyRing:          "",
				KmsLocation:         "",
				KmsProject:          "",
				Bucket:              "",
				CredentialSecretRef: nil,
			},
		},
	}

	for idx := range testData {
		test := testData[idx]
		t.Run(test.testName, func(t *testing.T) {
			gcsv1alph2 := test.spec
			gcsv1alph1 := &GoogleKmsGcsSpec{}
			if err := Convert_v1alpha2_GoogleKmsGcsSpec_To_v1alpha1_GoogleKmsGcsSpec(gcsv1alph2, gcsv1alph1, nil); err != nil {
				t.Error(err)
			}

			newgcsv1alpha2 := &v1alpha2.GoogleKmsGcsSpec{}
			if err := Convert_v1alpha1_GoogleKmsGcsSpec_To_v1alpha2_GoogleKmsGcsSpec(gcsv1alph1, newgcsv1alpha2, nil); err != nil {
				t.Error(err)
			}

			old, err := json.Marshal(gcsv1alph2)
			if err != nil {
				t.Error(err)
			}

			new, err := json.Marshal(newgcsv1alpha2)
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

func TestConvert_v1alpha1_GcsSpec_To_v1alpha2_GcsSpec(t *testing.T) {
	testData := []struct {
		testName string
		spec     *GcsSpec
	}{
		{
			testName: "test-0: should be successful",
			spec: &GcsSpec{
				Bucket:           "bucket",
				ChunkSize:        "size",
				MaxParallel:      0,
				HAEnabled:        false,
				CredentialSecret: "cred",
			},
		},
		{
			testName: "test-1: should be successful",
			spec: &GcsSpec{
				Bucket:           "",
				ChunkSize:        "",
				MaxParallel:      10,
				HAEnabled:        false,
				CredentialSecret: "",
			},
		},
	}

	for idx := range testData {
		test := testData[idx]
		t.Run(test.testName, func(t *testing.T) {
			gcsv1alph1 := test.spec
			gcsv1alph2 := &v1alpha2.GcsSpec{}
			if err := Convert_v1alpha1_GcsSpec_To_v1alpha2_GcsSpec(gcsv1alph1, gcsv1alph2, nil); err != nil {
				t.Error(err)
			}

			newgcsv1alpha1 := &GcsSpec{}
			if err := Convert_v1alpha2_GcsSpec_To_v1alpha1_GcsSpec(gcsv1alph2, newgcsv1alpha1, nil); err != nil {
				t.Error(err)
			}

			old, err := json.Marshal(gcsv1alph1)
			if err != nil {
				t.Error(err)
			}

			new, err := json.Marshal(newgcsv1alpha1)
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

func TestConvert_v1alpha2_GcsSpec_To_v1alpha1_GcsSpec(t *testing.T) {
	testData := []struct {
		testName string
		spec     *v1alpha2.GcsSpec
	}{
		{
			testName: "test-0: should be successful",
			spec: &v1alpha2.GcsSpec{
				Bucket:      "bucket",
				ChunkSize:   "size",
				MaxParallel: 0,
				HAEnabled:   false,
				CredentialSecretRef: &v1.LocalObjectReference{
					Name: "cred",
				},
			},
		},
		{
			testName: "test-1: should be successful",
			spec: &v1alpha2.GcsSpec{
				Bucket:              "",
				ChunkSize:           "",
				MaxParallel:         10,
				HAEnabled:           true,
				CredentialSecretRef: nil,
			},
		},
	}

	for idx := range testData {
		test := testData[idx]
		t.Run(test.testName, func(t *testing.T) {
			gcsv1alph2 := test.spec
			gcsv1alph1 := &GcsSpec{}
			if err := Convert_v1alpha2_GcsSpec_To_v1alpha1_GcsSpec(gcsv1alph2, gcsv1alph1, nil); err != nil {
				t.Error(err)
			}

			newgcsv1alpha2 := &v1alpha2.GcsSpec{}
			if err := Convert_v1alpha1_GcsSpec_To_v1alpha2_GcsSpec(gcsv1alph1, newgcsv1alpha2, nil); err != nil {
				t.Error(err)
			}

			old, err := json.Marshal(gcsv1alph2)
			if err != nil {
				t.Error(err)
			}

			new, err := json.Marshal(newgcsv1alpha2)
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

func TestConvert_v1alpha1_EtcdSpec_To_v1alpha2_EtcdSpec(t *testing.T) {
	testData := []struct {
		testName string
		spec     *EtcdSpec
	}{
		{
			testName: "test-0: should be successful",
			spec: &EtcdSpec{
				Address:              "address",
				EtcdApi:              "api",
				HAEnable:             false,
				Path:                 "path",
				Sync:                 false,
				DiscoverySrv:         "srv",
				CredentialSecretName: "cred",
				TLSSecretName:        "tls",
			},
		},
		{
			testName: "test-1: should be successful",
			spec: &EtcdSpec{
				Address:              "",
				EtcdApi:              "",
				HAEnable:             true,
				Path:                 "",
				Sync:                 true,
				DiscoverySrv:         "",
				CredentialSecretName: "",
				TLSSecretName:        "",
			},
		},
	}

	for idx := range testData {
		test := testData[idx]
		t.Run(test.testName, func(t *testing.T) {
			etcdv1alph1 := test.spec
			etcdsv1alph2 := &v1alpha2.EtcdSpec{}
			if err := Convert_v1alpha1_EtcdSpec_To_v1alpha2_EtcdSpec(etcdv1alph1, etcdsv1alph2, nil); err != nil {
				t.Error(err)
			}

			newetcdv1alpha1 := &EtcdSpec{}
			if err := Convert_v1alpha2_EtcdSpec_To_v1alpha1_EtcdSpec(etcdsv1alph2, newetcdv1alpha1, nil); err != nil {
				t.Error(err)
			}

			old, err := json.Marshal(etcdv1alph1)
			if err != nil {
				t.Error(err)
			}

			new, err := json.Marshal(newetcdv1alpha1)
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

func TestConvert_v1alpha2_EtcdSpec_To_v1alpha1_EtcdSpec(t *testing.T) {
	testData := []struct {
		testName string
		spec     *v1alpha2.EtcdSpec
	}{
		{
			testName: "test-0: should be successful",
			spec: &v1alpha2.EtcdSpec{
				Address:      "addr",
				EtcdApi:      "api",
				HAEnable:     false,
				Path:         "path",
				Sync:         false,
				DiscoverySrv: "srv",
				CredentialSecretRef: &v1.LocalObjectReference{
					Name: "cred",
				},
				TLSSecretRef: &v1.LocalObjectReference{
					Name: "tls",
				},
			},
		},
		{
			testName: "test-1: should be successful",
			spec: &v1alpha2.EtcdSpec{
				Address:             "",
				EtcdApi:             "",
				HAEnable:            false,
				Path:                "",
				Sync:                false,
				DiscoverySrv:        "",
				CredentialSecretRef: nil,
				TLSSecretRef:        nil,
			},
		},
	}

	for idx := range testData {
		test := testData[idx]
		t.Run(test.testName, func(t *testing.T) {
			etcdv1alph2 := test.spec
			etcdv1alph1 := &EtcdSpec{}
			if err := Convert_v1alpha2_EtcdSpec_To_v1alpha1_EtcdSpec(etcdv1alph2, etcdv1alph1, nil); err != nil {
				t.Error(err)
			}

			newetcdv1alpha2 := &v1alpha2.EtcdSpec{}
			if err := Convert_v1alpha1_EtcdSpec_To_v1alpha2_EtcdSpec(etcdv1alph1, newetcdv1alpha2, nil); err != nil {
				t.Error(err)
			}

			old, err := json.Marshal(etcdv1alph2)
			if err != nil {
				t.Error(err)
			}

			new, err := json.Marshal(newetcdv1alpha2)
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

func TestConvert_v1alpha1_DynamoDBSpec_To_v1alpha2_DynamoDBSpec(t *testing.T) {
	testData := []struct {
		testName string
		spec     *DynamoDBSpec
	}{
		{
			testName: "test-0: should be successful",
			spec: &DynamoDBSpec{
				Endpoint:           "endpoint",
				Region:             "region",
				HaEnabled:          false,
				ReadCapacity:       0,
				WriteCapacity:      1,
				Table:              "",
				CredentialSecret:   "cred",
				SessionTokenSecret: "",
				MaxParallel:        0,
			},
		},
		{
			testName: "test-1: should be successful",
			spec: &DynamoDBSpec{
				Endpoint:           "",
				Region:             "",
				HaEnabled:          false,
				ReadCapacity:       10,
				WriteCapacity:      1,
				Table:              "",
				CredentialSecret:   "",
				SessionTokenSecret: "",
				MaxParallel:        2,
			},
		},
	}

	for idx := range testData {
		test := testData[idx]
		t.Run(test.testName, func(t *testing.T) {
			dynv1alph1 := test.spec
			dynsv1alph2 := &v1alpha2.DynamoDBSpec{}
			if err := Convert_v1alpha1_DynamoDBSpec_To_v1alpha2_DynamoDBSpec(dynv1alph1, dynsv1alph2, nil); err != nil {
				t.Error(err)
			}

			newdynv1alpha1 := &DynamoDBSpec{}
			if err := Convert_v1alpha2_DynamoDBSpec_To_v1alpha1_DynamoDBSpec(dynsv1alph2, newdynv1alpha1, nil); err != nil {
				t.Error(err)
			}

			old, err := json.Marshal(dynv1alph1)
			if err != nil {
				t.Error(err)
			}

			new, err := json.Marshal(newdynv1alpha1)
			if err != nil {
				t.Error(err)
			}

			differ := diff.New()
			d, err := differ.Compare(old, new)
			if err != nil {
				t.Error(err)
			}

			if d.Modified() {
				fmt.Println("====delta: ", d.Deltas())
				t.Error("modified")
			}
		})
	}
}

func TestConvert_v1alpha2_DynamoDBSpec_To_v1alpha1_DynamoDBSpec(t *testing.T) {
	testData := []struct {
		testName string
		spec     *v1alpha2.DynamoDBSpec
	}{
		{
			testName: "test-0: should be successful",
			spec: &v1alpha2.DynamoDBSpec{
				Endpoint:      "",
				Region:        "",
				HaEnabled:     false,
				ReadCapacity:  0,
				WriteCapacity: 0,
				Table:         "",
				CredentialSecretRef: &v1.LocalObjectReference{
					Name: "cred",
				},
				MaxParallel: 0,
			},
		},
		{
			testName: "test-1: should be successful",
			spec: &v1alpha2.DynamoDBSpec{
				Endpoint:            "",
				Region:              "",
				HaEnabled:           false,
				ReadCapacity:        0,
				WriteCapacity:       0,
				Table:               "",
				CredentialSecretRef: nil,
				MaxParallel:         0,
			},
		},
	}

	for idx := range testData {
		test := testData[idx]
		t.Run(test.testName, func(t *testing.T) {
			dynv1alph2 := test.spec
			dynv1alph1 := &DynamoDBSpec{}
			if err := Convert_v1alpha2_DynamoDBSpec_To_v1alpha1_DynamoDBSpec(dynv1alph2, dynv1alph1, nil); err != nil {
				t.Error(err)
			}

			newdynv1alpha2 := &v1alpha2.DynamoDBSpec{}
			if err := Convert_v1alpha1_DynamoDBSpec_To_v1alpha2_DynamoDBSpec(dynv1alph1, newdynv1alpha2, nil); err != nil {
				t.Error(err)
			}

			old, err := json.Marshal(dynv1alph2)
			if err != nil {
				t.Error(err)
			}

			new, err := json.Marshal(newdynv1alpha2)
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
