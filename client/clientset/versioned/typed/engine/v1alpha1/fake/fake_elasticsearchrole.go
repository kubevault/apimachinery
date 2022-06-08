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
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeElasticsearchRoles implements ElasticsearchRoleInterface
type FakeElasticsearchRoles struct {
	Fake *FakeEngineV1alpha1
	ns   string
}

var elasticsearchrolesResource = schema.GroupVersionResource{Group: "engine.kubevault.com", Version: "v1alpha1", Resource: "elasticsearchroles"}

var elasticsearchrolesKind = schema.GroupVersionKind{Group: "engine.kubevault.com", Version: "v1alpha1", Kind: "ElasticsearchRole"}

// Get takes name of the elasticsearchRole, and returns the corresponding elasticsearchRole object, and an error if there is any.
func (c *FakeElasticsearchRoles) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ElasticsearchRole, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(elasticsearchrolesResource, c.ns, name), &v1alpha1.ElasticsearchRole{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ElasticsearchRole), err
}

// List takes label and field selectors, and returns the list of ElasticsearchRoles that match those selectors.
func (c *FakeElasticsearchRoles) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ElasticsearchRoleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(elasticsearchrolesResource, elasticsearchrolesKind, c.ns, opts), &v1alpha1.ElasticsearchRoleList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ElasticsearchRoleList{ListMeta: obj.(*v1alpha1.ElasticsearchRoleList).ListMeta}
	for _, item := range obj.(*v1alpha1.ElasticsearchRoleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested elasticsearchRoles.
func (c *FakeElasticsearchRoles) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(elasticsearchrolesResource, c.ns, opts))

}

// Create takes the representation of a elasticsearchRole and creates it.  Returns the server's representation of the elasticsearchRole, and an error, if there is any.
func (c *FakeElasticsearchRoles) Create(ctx context.Context, elasticsearchRole *v1alpha1.ElasticsearchRole, opts v1.CreateOptions) (result *v1alpha1.ElasticsearchRole, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(elasticsearchrolesResource, c.ns, elasticsearchRole), &v1alpha1.ElasticsearchRole{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ElasticsearchRole), err
}

// Update takes the representation of a elasticsearchRole and updates it. Returns the server's representation of the elasticsearchRole, and an error, if there is any.
func (c *FakeElasticsearchRoles) Update(ctx context.Context, elasticsearchRole *v1alpha1.ElasticsearchRole, opts v1.UpdateOptions) (result *v1alpha1.ElasticsearchRole, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(elasticsearchrolesResource, c.ns, elasticsearchRole), &v1alpha1.ElasticsearchRole{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ElasticsearchRole), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeElasticsearchRoles) UpdateStatus(ctx context.Context, elasticsearchRole *v1alpha1.ElasticsearchRole, opts v1.UpdateOptions) (*v1alpha1.ElasticsearchRole, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(elasticsearchrolesResource, "status", c.ns, elasticsearchRole), &v1alpha1.ElasticsearchRole{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ElasticsearchRole), err
}

// Delete takes name of the elasticsearchRole and deletes it. Returns an error if one occurs.
func (c *FakeElasticsearchRoles) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(elasticsearchrolesResource, c.ns, name, opts), &v1alpha1.ElasticsearchRole{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeElasticsearchRoles) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(elasticsearchrolesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ElasticsearchRoleList{})
	return err
}

// Patch applies the patch and returns the patched elasticsearchRole.
func (c *FakeElasticsearchRoles) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ElasticsearchRole, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(elasticsearchrolesResource, c.ns, name, pt, data, subresources...), &v1alpha1.ElasticsearchRole{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ElasticsearchRole), err
}
