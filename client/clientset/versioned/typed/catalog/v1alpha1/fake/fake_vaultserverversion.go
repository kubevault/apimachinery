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
	"context"

	v1alpha1 "kubevault.dev/apimachinery/apis/catalog/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeVaultServerVersions implements VaultServerVersionInterface
type FakeVaultServerVersions struct {
	Fake *FakeCatalogV1alpha1
}

var vaultserverversionsResource = v1alpha1.SchemeGroupVersion.WithResource("vaultserverversions")

var vaultserverversionsKind = v1alpha1.SchemeGroupVersion.WithKind("VaultServerVersion")

// Get takes name of the vaultServerVersion, and returns the corresponding vaultServerVersion object, and an error if there is any.
func (c *FakeVaultServerVersions) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.VaultServerVersion, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(vaultserverversionsResource, name), &v1alpha1.VaultServerVersion{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VaultServerVersion), err
}

// List takes label and field selectors, and returns the list of VaultServerVersions that match those selectors.
func (c *FakeVaultServerVersions) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.VaultServerVersionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(vaultserverversionsResource, vaultserverversionsKind, opts), &v1alpha1.VaultServerVersionList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.VaultServerVersionList{ListMeta: obj.(*v1alpha1.VaultServerVersionList).ListMeta}
	for _, item := range obj.(*v1alpha1.VaultServerVersionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested vaultServerVersions.
func (c *FakeVaultServerVersions) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(vaultserverversionsResource, opts))
}

// Create takes the representation of a vaultServerVersion and creates it.  Returns the server's representation of the vaultServerVersion, and an error, if there is any.
func (c *FakeVaultServerVersions) Create(ctx context.Context, vaultServerVersion *v1alpha1.VaultServerVersion, opts v1.CreateOptions) (result *v1alpha1.VaultServerVersion, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(vaultserverversionsResource, vaultServerVersion), &v1alpha1.VaultServerVersion{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VaultServerVersion), err
}

// Update takes the representation of a vaultServerVersion and updates it. Returns the server's representation of the vaultServerVersion, and an error, if there is any.
func (c *FakeVaultServerVersions) Update(ctx context.Context, vaultServerVersion *v1alpha1.VaultServerVersion, opts v1.UpdateOptions) (result *v1alpha1.VaultServerVersion, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(vaultserverversionsResource, vaultServerVersion), &v1alpha1.VaultServerVersion{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VaultServerVersion), err
}

// Delete takes name of the vaultServerVersion and deletes it. Returns an error if one occurs.
func (c *FakeVaultServerVersions) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(vaultserverversionsResource, name, opts), &v1alpha1.VaultServerVersion{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeVaultServerVersions) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(vaultserverversionsResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.VaultServerVersionList{})
	return err
}

// Patch applies the patch and returns the patched vaultServerVersion.
func (c *FakeVaultServerVersions) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.VaultServerVersion, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(vaultserverversionsResource, name, pt, data, subresources...), &v1alpha1.VaultServerVersion{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VaultServerVersion), err
}
