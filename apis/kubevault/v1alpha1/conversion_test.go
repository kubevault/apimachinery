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
)

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
