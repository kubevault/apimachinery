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

	v1alpha1 "kubevault.dev/apimachinery/apis/engine/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeMariaDBRoles implements MariaDBRoleInterface
type FakeMariaDBRoles struct {
	Fake *FakeEngineV1alpha1
	ns   string
}

var mariadbrolesResource = v1alpha1.SchemeGroupVersion.WithResource("mariadbroles")

var mariadbrolesKind = v1alpha1.SchemeGroupVersion.WithKind("MariaDBRole")

// Get takes name of the mariaDBRole, and returns the corresponding mariaDBRole object, and an error if there is any.
func (c *FakeMariaDBRoles) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.MariaDBRole, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(mariadbrolesResource, c.ns, name), &v1alpha1.MariaDBRole{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MariaDBRole), err
}

// List takes label and field selectors, and returns the list of MariaDBRoles that match those selectors.
func (c *FakeMariaDBRoles) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.MariaDBRoleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(mariadbrolesResource, mariadbrolesKind, c.ns, opts), &v1alpha1.MariaDBRoleList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.MariaDBRoleList{ListMeta: obj.(*v1alpha1.MariaDBRoleList).ListMeta}
	for _, item := range obj.(*v1alpha1.MariaDBRoleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested mariaDBRoles.
func (c *FakeMariaDBRoles) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(mariadbrolesResource, c.ns, opts))

}

// Create takes the representation of a mariaDBRole and creates it.  Returns the server's representation of the mariaDBRole, and an error, if there is any.
func (c *FakeMariaDBRoles) Create(ctx context.Context, mariaDBRole *v1alpha1.MariaDBRole, opts v1.CreateOptions) (result *v1alpha1.MariaDBRole, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(mariadbrolesResource, c.ns, mariaDBRole), &v1alpha1.MariaDBRole{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MariaDBRole), err
}

// Update takes the representation of a mariaDBRole and updates it. Returns the server's representation of the mariaDBRole, and an error, if there is any.
func (c *FakeMariaDBRoles) Update(ctx context.Context, mariaDBRole *v1alpha1.MariaDBRole, opts v1.UpdateOptions) (result *v1alpha1.MariaDBRole, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(mariadbrolesResource, c.ns, mariaDBRole), &v1alpha1.MariaDBRole{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MariaDBRole), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeMariaDBRoles) UpdateStatus(ctx context.Context, mariaDBRole *v1alpha1.MariaDBRole, opts v1.UpdateOptions) (*v1alpha1.MariaDBRole, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(mariadbrolesResource, "status", c.ns, mariaDBRole), &v1alpha1.MariaDBRole{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MariaDBRole), err
}

// Delete takes name of the mariaDBRole and deletes it. Returns an error if one occurs.
func (c *FakeMariaDBRoles) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(mariadbrolesResource, c.ns, name, opts), &v1alpha1.MariaDBRole{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMariaDBRoles) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(mariadbrolesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.MariaDBRoleList{})
	return err
}

// Patch applies the patch and returns the patched mariaDBRole.
func (c *FakeMariaDBRoles) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.MariaDBRole, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(mariadbrolesResource, c.ns, name, pt, data, subresources...), &v1alpha1.MariaDBRole{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MariaDBRole), err
}
