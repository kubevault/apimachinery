/*
Copyright The KubeVault Authors.

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

	v1alpha1 "kubevault.dev/operator/apis/engine/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeDatabaseAccessRequests implements DatabaseAccessRequestInterface
type FakeDatabaseAccessRequests struct {
	Fake *FakeEngineV1alpha1
	ns   string
}

var databaseaccessrequestsResource = schema.GroupVersionResource{Group: "engine.kubevault.com", Version: "v1alpha1", Resource: "databaseaccessrequests"}

var databaseaccessrequestsKind = schema.GroupVersionKind{Group: "engine.kubevault.com", Version: "v1alpha1", Kind: "DatabaseAccessRequest"}

// Get takes name of the databaseAccessRequest, and returns the corresponding databaseAccessRequest object, and an error if there is any.
func (c *FakeDatabaseAccessRequests) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.DatabaseAccessRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(databaseaccessrequestsResource, c.ns, name), &v1alpha1.DatabaseAccessRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DatabaseAccessRequest), err
}

// List takes label and field selectors, and returns the list of DatabaseAccessRequests that match those selectors.
func (c *FakeDatabaseAccessRequests) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.DatabaseAccessRequestList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(databaseaccessrequestsResource, databaseaccessrequestsKind, c.ns, opts), &v1alpha1.DatabaseAccessRequestList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DatabaseAccessRequestList{ListMeta: obj.(*v1alpha1.DatabaseAccessRequestList).ListMeta}
	for _, item := range obj.(*v1alpha1.DatabaseAccessRequestList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested databaseAccessRequests.
func (c *FakeDatabaseAccessRequests) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(databaseaccessrequestsResource, c.ns, opts))

}

// Create takes the representation of a databaseAccessRequest and creates it.  Returns the server's representation of the databaseAccessRequest, and an error, if there is any.
func (c *FakeDatabaseAccessRequests) Create(ctx context.Context, databaseAccessRequest *v1alpha1.DatabaseAccessRequest, opts v1.CreateOptions) (result *v1alpha1.DatabaseAccessRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(databaseaccessrequestsResource, c.ns, databaseAccessRequest), &v1alpha1.DatabaseAccessRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DatabaseAccessRequest), err
}

// Update takes the representation of a databaseAccessRequest and updates it. Returns the server's representation of the databaseAccessRequest, and an error, if there is any.
func (c *FakeDatabaseAccessRequests) Update(ctx context.Context, databaseAccessRequest *v1alpha1.DatabaseAccessRequest, opts v1.UpdateOptions) (result *v1alpha1.DatabaseAccessRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(databaseaccessrequestsResource, c.ns, databaseAccessRequest), &v1alpha1.DatabaseAccessRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DatabaseAccessRequest), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDatabaseAccessRequests) UpdateStatus(ctx context.Context, databaseAccessRequest *v1alpha1.DatabaseAccessRequest, opts v1.UpdateOptions) (*v1alpha1.DatabaseAccessRequest, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(databaseaccessrequestsResource, "status", c.ns, databaseAccessRequest), &v1alpha1.DatabaseAccessRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DatabaseAccessRequest), err
}

// Delete takes name of the databaseAccessRequest and deletes it. Returns an error if one occurs.
func (c *FakeDatabaseAccessRequests) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(databaseaccessrequestsResource, c.ns, name), &v1alpha1.DatabaseAccessRequest{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDatabaseAccessRequests) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(databaseaccessrequestsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.DatabaseAccessRequestList{})
	return err
}

// Patch applies the patch and returns the patched databaseAccessRequest.
func (c *FakeDatabaseAccessRequests) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.DatabaseAccessRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(databaseaccessrequestsResource, c.ns, name, pt, data, subresources...), &v1alpha1.DatabaseAccessRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DatabaseAccessRequest), err
}
