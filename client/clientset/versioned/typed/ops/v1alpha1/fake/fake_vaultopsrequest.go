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

	v1alpha1 "kubevault.dev/apimachinery/apis/ops/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeVaultOpsRequests implements VaultOpsRequestInterface
type FakeVaultOpsRequests struct {
	Fake *FakeOpsV1alpha1
	ns   string
}

var vaultopsrequestsResource = v1alpha1.SchemeGroupVersion.WithResource("vaultopsrequests")

var vaultopsrequestsKind = v1alpha1.SchemeGroupVersion.WithKind("VaultOpsRequest")

// Get takes name of the vaultOpsRequest, and returns the corresponding vaultOpsRequest object, and an error if there is any.
func (c *FakeVaultOpsRequests) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.VaultOpsRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(vaultopsrequestsResource, c.ns, name), &v1alpha1.VaultOpsRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VaultOpsRequest), err
}

// List takes label and field selectors, and returns the list of VaultOpsRequests that match those selectors.
func (c *FakeVaultOpsRequests) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.VaultOpsRequestList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(vaultopsrequestsResource, vaultopsrequestsKind, c.ns, opts), &v1alpha1.VaultOpsRequestList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.VaultOpsRequestList{ListMeta: obj.(*v1alpha1.VaultOpsRequestList).ListMeta}
	for _, item := range obj.(*v1alpha1.VaultOpsRequestList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested vaultOpsRequests.
func (c *FakeVaultOpsRequests) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(vaultopsrequestsResource, c.ns, opts))

}

// Create takes the representation of a vaultOpsRequest and creates it.  Returns the server's representation of the vaultOpsRequest, and an error, if there is any.
func (c *FakeVaultOpsRequests) Create(ctx context.Context, vaultOpsRequest *v1alpha1.VaultOpsRequest, opts v1.CreateOptions) (result *v1alpha1.VaultOpsRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(vaultopsrequestsResource, c.ns, vaultOpsRequest), &v1alpha1.VaultOpsRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VaultOpsRequest), err
}

// Update takes the representation of a vaultOpsRequest and updates it. Returns the server's representation of the vaultOpsRequest, and an error, if there is any.
func (c *FakeVaultOpsRequests) Update(ctx context.Context, vaultOpsRequest *v1alpha1.VaultOpsRequest, opts v1.UpdateOptions) (result *v1alpha1.VaultOpsRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(vaultopsrequestsResource, c.ns, vaultOpsRequest), &v1alpha1.VaultOpsRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VaultOpsRequest), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeVaultOpsRequests) UpdateStatus(ctx context.Context, vaultOpsRequest *v1alpha1.VaultOpsRequest, opts v1.UpdateOptions) (*v1alpha1.VaultOpsRequest, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(vaultopsrequestsResource, "status", c.ns, vaultOpsRequest), &v1alpha1.VaultOpsRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VaultOpsRequest), err
}

// Delete takes name of the vaultOpsRequest and deletes it. Returns an error if one occurs.
func (c *FakeVaultOpsRequests) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(vaultopsrequestsResource, c.ns, name, opts), &v1alpha1.VaultOpsRequest{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeVaultOpsRequests) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(vaultopsrequestsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.VaultOpsRequestList{})
	return err
}

// Patch applies the patch and returns the patched vaultOpsRequest.
func (c *FakeVaultOpsRequests) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.VaultOpsRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(vaultopsrequestsResource, c.ns, name, pt, data, subresources...), &v1alpha1.VaultOpsRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VaultOpsRequest), err
}
