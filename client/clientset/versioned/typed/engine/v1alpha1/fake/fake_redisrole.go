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

// FakeRedisRoles implements RedisRoleInterface
type FakeRedisRoles struct {
	Fake *FakeEngineV1alpha1
	ns   string
}

var redisrolesResource = v1alpha1.SchemeGroupVersion.WithResource("redisroles")

var redisrolesKind = v1alpha1.SchemeGroupVersion.WithKind("RedisRole")

// Get takes name of the redisRole, and returns the corresponding redisRole object, and an error if there is any.
func (c *FakeRedisRoles) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.RedisRole, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(redisrolesResource, c.ns, name), &v1alpha1.RedisRole{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RedisRole), err
}

// List takes label and field selectors, and returns the list of RedisRoles that match those selectors.
func (c *FakeRedisRoles) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.RedisRoleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(redisrolesResource, redisrolesKind, c.ns, opts), &v1alpha1.RedisRoleList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.RedisRoleList{ListMeta: obj.(*v1alpha1.RedisRoleList).ListMeta}
	for _, item := range obj.(*v1alpha1.RedisRoleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested redisRoles.
func (c *FakeRedisRoles) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(redisrolesResource, c.ns, opts))

}

// Create takes the representation of a redisRole and creates it.  Returns the server's representation of the redisRole, and an error, if there is any.
func (c *FakeRedisRoles) Create(ctx context.Context, redisRole *v1alpha1.RedisRole, opts v1.CreateOptions) (result *v1alpha1.RedisRole, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(redisrolesResource, c.ns, redisRole), &v1alpha1.RedisRole{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RedisRole), err
}

// Update takes the representation of a redisRole and updates it. Returns the server's representation of the redisRole, and an error, if there is any.
func (c *FakeRedisRoles) Update(ctx context.Context, redisRole *v1alpha1.RedisRole, opts v1.UpdateOptions) (result *v1alpha1.RedisRole, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(redisrolesResource, c.ns, redisRole), &v1alpha1.RedisRole{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RedisRole), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeRedisRoles) UpdateStatus(ctx context.Context, redisRole *v1alpha1.RedisRole, opts v1.UpdateOptions) (*v1alpha1.RedisRole, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(redisrolesResource, "status", c.ns, redisRole), &v1alpha1.RedisRole{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RedisRole), err
}

// Delete takes name of the redisRole and deletes it. Returns an error if one occurs.
func (c *FakeRedisRoles) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(redisrolesResource, c.ns, name, opts), &v1alpha1.RedisRole{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRedisRoles) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(redisrolesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.RedisRoleList{})
	return err
}

// Patch applies the patch and returns the patched redisRole.
func (c *FakeRedisRoles) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.RedisRole, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(redisrolesResource, c.ns, name, pt, data, subresources...), &v1alpha1.RedisRole{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RedisRole), err
}
