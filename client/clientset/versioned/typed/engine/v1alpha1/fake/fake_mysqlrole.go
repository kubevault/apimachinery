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

// FakeMySQLRoles implements MySQLRoleInterface
type FakeMySQLRoles struct {
	Fake *FakeEngineV1alpha1
	ns   string
}

var mysqlrolesResource = v1alpha1.SchemeGroupVersion.WithResource("mysqlroles")

var mysqlrolesKind = v1alpha1.SchemeGroupVersion.WithKind("MySQLRole")

// Get takes name of the mySQLRole, and returns the corresponding mySQLRole object, and an error if there is any.
func (c *FakeMySQLRoles) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.MySQLRole, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(mysqlrolesResource, c.ns, name), &v1alpha1.MySQLRole{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MySQLRole), err
}

// List takes label and field selectors, and returns the list of MySQLRoles that match those selectors.
func (c *FakeMySQLRoles) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.MySQLRoleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(mysqlrolesResource, mysqlrolesKind, c.ns, opts), &v1alpha1.MySQLRoleList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.MySQLRoleList{ListMeta: obj.(*v1alpha1.MySQLRoleList).ListMeta}
	for _, item := range obj.(*v1alpha1.MySQLRoleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested mySQLRoles.
func (c *FakeMySQLRoles) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(mysqlrolesResource, c.ns, opts))

}

// Create takes the representation of a mySQLRole and creates it.  Returns the server's representation of the mySQLRole, and an error, if there is any.
func (c *FakeMySQLRoles) Create(ctx context.Context, mySQLRole *v1alpha1.MySQLRole, opts v1.CreateOptions) (result *v1alpha1.MySQLRole, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(mysqlrolesResource, c.ns, mySQLRole), &v1alpha1.MySQLRole{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MySQLRole), err
}

// Update takes the representation of a mySQLRole and updates it. Returns the server's representation of the mySQLRole, and an error, if there is any.
func (c *FakeMySQLRoles) Update(ctx context.Context, mySQLRole *v1alpha1.MySQLRole, opts v1.UpdateOptions) (result *v1alpha1.MySQLRole, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(mysqlrolesResource, c.ns, mySQLRole), &v1alpha1.MySQLRole{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MySQLRole), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeMySQLRoles) UpdateStatus(ctx context.Context, mySQLRole *v1alpha1.MySQLRole, opts v1.UpdateOptions) (*v1alpha1.MySQLRole, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(mysqlrolesResource, "status", c.ns, mySQLRole), &v1alpha1.MySQLRole{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MySQLRole), err
}

// Delete takes name of the mySQLRole and deletes it. Returns an error if one occurs.
func (c *FakeMySQLRoles) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(mysqlrolesResource, c.ns, name, opts), &v1alpha1.MySQLRole{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMySQLRoles) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(mysqlrolesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.MySQLRoleList{})
	return err
}

// Patch applies the patch and returns the patched mySQLRole.
func (c *FakeMySQLRoles) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.MySQLRole, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(mysqlrolesResource, c.ns, name, pt, data, subresources...), &v1alpha1.MySQLRole{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MySQLRole), err
}
