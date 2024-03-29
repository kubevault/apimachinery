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

// FakeSecretAccessRequests implements SecretAccessRequestInterface
type FakeSecretAccessRequests struct {
	Fake *FakeEngineV1alpha1
	ns   string
}

var secretaccessrequestsResource = v1alpha1.SchemeGroupVersion.WithResource("secretaccessrequests")

var secretaccessrequestsKind = v1alpha1.SchemeGroupVersion.WithKind("SecretAccessRequest")

// Get takes name of the secretAccessRequest, and returns the corresponding secretAccessRequest object, and an error if there is any.
func (c *FakeSecretAccessRequests) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.SecretAccessRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(secretaccessrequestsResource, c.ns, name), &v1alpha1.SecretAccessRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SecretAccessRequest), err
}

// List takes label and field selectors, and returns the list of SecretAccessRequests that match those selectors.
func (c *FakeSecretAccessRequests) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.SecretAccessRequestList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(secretaccessrequestsResource, secretaccessrequestsKind, c.ns, opts), &v1alpha1.SecretAccessRequestList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SecretAccessRequestList{ListMeta: obj.(*v1alpha1.SecretAccessRequestList).ListMeta}
	for _, item := range obj.(*v1alpha1.SecretAccessRequestList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested secretAccessRequests.
func (c *FakeSecretAccessRequests) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(secretaccessrequestsResource, c.ns, opts))

}

// Create takes the representation of a secretAccessRequest and creates it.  Returns the server's representation of the secretAccessRequest, and an error, if there is any.
func (c *FakeSecretAccessRequests) Create(ctx context.Context, secretAccessRequest *v1alpha1.SecretAccessRequest, opts v1.CreateOptions) (result *v1alpha1.SecretAccessRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(secretaccessrequestsResource, c.ns, secretAccessRequest), &v1alpha1.SecretAccessRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SecretAccessRequest), err
}

// Update takes the representation of a secretAccessRequest and updates it. Returns the server's representation of the secretAccessRequest, and an error, if there is any.
func (c *FakeSecretAccessRequests) Update(ctx context.Context, secretAccessRequest *v1alpha1.SecretAccessRequest, opts v1.UpdateOptions) (result *v1alpha1.SecretAccessRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(secretaccessrequestsResource, c.ns, secretAccessRequest), &v1alpha1.SecretAccessRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SecretAccessRequest), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSecretAccessRequests) UpdateStatus(ctx context.Context, secretAccessRequest *v1alpha1.SecretAccessRequest, opts v1.UpdateOptions) (*v1alpha1.SecretAccessRequest, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(secretaccessrequestsResource, "status", c.ns, secretAccessRequest), &v1alpha1.SecretAccessRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SecretAccessRequest), err
}

// Delete takes name of the secretAccessRequest and deletes it. Returns an error if one occurs.
func (c *FakeSecretAccessRequests) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(secretaccessrequestsResource, c.ns, name, opts), &v1alpha1.SecretAccessRequest{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSecretAccessRequests) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(secretaccessrequestsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.SecretAccessRequestList{})
	return err
}

// Patch applies the patch and returns the patched secretAccessRequest.
func (c *FakeSecretAccessRequests) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.SecretAccessRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(secretaccessrequestsResource, c.ns, name, pt, data, subresources...), &v1alpha1.SecretAccessRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SecretAccessRequest), err
}
