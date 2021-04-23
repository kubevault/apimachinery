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

package v1alpha1

import (
	"context"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1alpha1 "kubevault.dev/apimachinery/apis/engine/v1alpha1"
	scheme "kubevault.dev/apimachinery/client/clientset/versioned/scheme"
)

// DatabaseAccessRequestsGetter has a method to return a DatabaseAccessRequestInterface.
// A group's client should implement this interface.
type DatabaseAccessRequestsGetter interface {
	DatabaseAccessRequests(namespace string) DatabaseAccessRequestInterface
}

// DatabaseAccessRequestInterface has methods to work with DatabaseAccessRequest resources.
type DatabaseAccessRequestInterface interface {
	Create(ctx context.Context, databaseAccessRequest *v1alpha1.DatabaseAccessRequest, opts v1.CreateOptions) (*v1alpha1.DatabaseAccessRequest, error)
	Update(ctx context.Context, databaseAccessRequest *v1alpha1.DatabaseAccessRequest, opts v1.UpdateOptions) (*v1alpha1.DatabaseAccessRequest, error)
	UpdateStatus(ctx context.Context, databaseAccessRequest *v1alpha1.DatabaseAccessRequest, opts v1.UpdateOptions) (*v1alpha1.DatabaseAccessRequest, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.DatabaseAccessRequest, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.DatabaseAccessRequestList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.DatabaseAccessRequest, err error)
	DatabaseAccessRequestExpansion
}

// databaseAccessRequests implements DatabaseAccessRequestInterface
type databaseAccessRequests struct {
	client rest.Interface
	ns     string
}

// newDatabaseAccessRequests returns a DatabaseAccessRequests
func newDatabaseAccessRequests(c *EngineV1alpha1Client, namespace string) *databaseAccessRequests {
	return &databaseAccessRequests{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the databaseAccessRequest, and returns the corresponding databaseAccessRequest object, and an error if there is any.
func (c *databaseAccessRequests) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.DatabaseAccessRequest, err error) {
	result = &v1alpha1.DatabaseAccessRequest{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("databaseaccessrequests").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of DatabaseAccessRequests that match those selectors.
func (c *databaseAccessRequests) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.DatabaseAccessRequestList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.DatabaseAccessRequestList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("databaseaccessrequests").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested databaseAccessRequests.
func (c *databaseAccessRequests) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("databaseaccessrequests").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a databaseAccessRequest and creates it.  Returns the server's representation of the databaseAccessRequest, and an error, if there is any.
func (c *databaseAccessRequests) Create(ctx context.Context, databaseAccessRequest *v1alpha1.DatabaseAccessRequest, opts v1.CreateOptions) (result *v1alpha1.DatabaseAccessRequest, err error) {
	result = &v1alpha1.DatabaseAccessRequest{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("databaseaccessrequests").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(databaseAccessRequest).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a databaseAccessRequest and updates it. Returns the server's representation of the databaseAccessRequest, and an error, if there is any.
func (c *databaseAccessRequests) Update(ctx context.Context, databaseAccessRequest *v1alpha1.DatabaseAccessRequest, opts v1.UpdateOptions) (result *v1alpha1.DatabaseAccessRequest, err error) {
	result = &v1alpha1.DatabaseAccessRequest{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("databaseaccessrequests").
		Name(databaseAccessRequest.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(databaseAccessRequest).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *databaseAccessRequests) UpdateStatus(ctx context.Context, databaseAccessRequest *v1alpha1.DatabaseAccessRequest, opts v1.UpdateOptions) (result *v1alpha1.DatabaseAccessRequest, err error) {
	result = &v1alpha1.DatabaseAccessRequest{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("databaseaccessrequests").
		Name(databaseAccessRequest.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(databaseAccessRequest).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the databaseAccessRequest and deletes it. Returns an error if one occurs.
func (c *databaseAccessRequests) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("databaseaccessrequests").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *databaseAccessRequests) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("databaseaccessrequests").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched databaseAccessRequest.
func (c *databaseAccessRequests) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.DatabaseAccessRequest, err error) {
	result = &v1alpha1.DatabaseAccessRequest{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("databaseaccessrequests").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
