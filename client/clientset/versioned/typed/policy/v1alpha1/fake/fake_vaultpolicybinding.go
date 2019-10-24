/*
Copyright 2019 The Kube Vault Authors.

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
	v1alpha1 "kubevault.dev/operator/apis/policy/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeVaultPolicyBindings implements VaultPolicyBindingInterface
type FakeVaultPolicyBindings struct {
	Fake *FakePolicyV1alpha1
	ns   string
}

var vaultpolicybindingsResource = schema.GroupVersionResource{Group: "policy.kubevault.com", Version: "v1alpha1", Resource: "vaultpolicybindings"}

var vaultpolicybindingsKind = schema.GroupVersionKind{Group: "policy.kubevault.com", Version: "v1alpha1", Kind: "VaultPolicyBinding"}

// Get takes name of the vaultPolicyBinding, and returns the corresponding vaultPolicyBinding object, and an error if there is any.
func (c *FakeVaultPolicyBindings) Get(name string, options v1.GetOptions) (result *v1alpha1.VaultPolicyBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(vaultpolicybindingsResource, c.ns, name), &v1alpha1.VaultPolicyBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VaultPolicyBinding), err
}

// List takes label and field selectors, and returns the list of VaultPolicyBindings that match those selectors.
func (c *FakeVaultPolicyBindings) List(opts v1.ListOptions) (result *v1alpha1.VaultPolicyBindingList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(vaultpolicybindingsResource, vaultpolicybindingsKind, c.ns, opts), &v1alpha1.VaultPolicyBindingList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.VaultPolicyBindingList{ListMeta: obj.(*v1alpha1.VaultPolicyBindingList).ListMeta}
	for _, item := range obj.(*v1alpha1.VaultPolicyBindingList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested vaultPolicyBindings.
func (c *FakeVaultPolicyBindings) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(vaultpolicybindingsResource, c.ns, opts))

}

// Create takes the representation of a vaultPolicyBinding and creates it.  Returns the server's representation of the vaultPolicyBinding, and an error, if there is any.
func (c *FakeVaultPolicyBindings) Create(vaultPolicyBinding *v1alpha1.VaultPolicyBinding) (result *v1alpha1.VaultPolicyBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(vaultpolicybindingsResource, c.ns, vaultPolicyBinding), &v1alpha1.VaultPolicyBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VaultPolicyBinding), err
}

// Update takes the representation of a vaultPolicyBinding and updates it. Returns the server's representation of the vaultPolicyBinding, and an error, if there is any.
func (c *FakeVaultPolicyBindings) Update(vaultPolicyBinding *v1alpha1.VaultPolicyBinding) (result *v1alpha1.VaultPolicyBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(vaultpolicybindingsResource, c.ns, vaultPolicyBinding), &v1alpha1.VaultPolicyBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VaultPolicyBinding), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeVaultPolicyBindings) UpdateStatus(vaultPolicyBinding *v1alpha1.VaultPolicyBinding) (*v1alpha1.VaultPolicyBinding, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(vaultpolicybindingsResource, "status", c.ns, vaultPolicyBinding), &v1alpha1.VaultPolicyBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VaultPolicyBinding), err
}

// Delete takes name of the vaultPolicyBinding and deletes it. Returns an error if one occurs.
func (c *FakeVaultPolicyBindings) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(vaultpolicybindingsResource, c.ns, name), &v1alpha1.VaultPolicyBinding{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeVaultPolicyBindings) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(vaultpolicybindingsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.VaultPolicyBindingList{})
	return err
}

// Patch applies the patch and returns the patched vaultPolicyBinding.
func (c *FakeVaultPolicyBindings) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.VaultPolicyBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(vaultpolicybindingsResource, c.ns, name, pt, data, subresources...), &v1alpha1.VaultPolicyBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VaultPolicyBinding), err
}
