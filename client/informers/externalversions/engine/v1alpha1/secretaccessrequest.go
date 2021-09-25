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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	enginev1alpha1 "kubevault.dev/apimachinery/apis/engine/v1alpha1"
	versioned "kubevault.dev/apimachinery/client/clientset/versioned"
	internalinterfaces "kubevault.dev/apimachinery/client/informers/externalversions/internalinterfaces"
	v1alpha1 "kubevault.dev/apimachinery/client/listers/engine/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// SecretAccessRequestInformer provides access to a shared informer and lister for
// SecretAccessRequests.
type SecretAccessRequestInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.SecretAccessRequestLister
}

type secretAccessRequestInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewSecretAccessRequestInformer constructs a new informer for SecretAccessRequest type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewSecretAccessRequestInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredSecretAccessRequestInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredSecretAccessRequestInformer constructs a new informer for SecretAccessRequest type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredSecretAccessRequestInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.EngineV1alpha1().SecretAccessRequests(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.EngineV1alpha1().SecretAccessRequests(namespace).Watch(context.TODO(), options)
			},
		},
		&enginev1alpha1.SecretAccessRequest{},
		resyncPeriod,
		indexers,
	)
}

func (f *secretAccessRequestInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredSecretAccessRequestInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *secretAccessRequestInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&enginev1alpha1.SecretAccessRequest{}, f.defaultInformer)
}

func (f *secretAccessRequestInformer) Lister() v1alpha1.SecretAccessRequestLister {
	return v1alpha1.NewSecretAccessRequestLister(f.Informer().GetIndexer())
}