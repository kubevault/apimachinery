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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	time "time"

	enginev1alpha1 "kubevault.dev/operator/apis/engine/v1alpha1"
	versioned "kubevault.dev/operator/client/clientset/versioned"
	internalinterfaces "kubevault.dev/operator/client/informers/externalversions/internalinterfaces"
	v1alpha1 "kubevault.dev/operator/client/listers/engine/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// AzureAccessKeyRequestInformer provides access to a shared informer and lister for
// AzureAccessKeyRequests.
type AzureAccessKeyRequestInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.AzureAccessKeyRequestLister
}

type azureAccessKeyRequestInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewAzureAccessKeyRequestInformer constructs a new informer for AzureAccessKeyRequest type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewAzureAccessKeyRequestInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredAzureAccessKeyRequestInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredAzureAccessKeyRequestInformer constructs a new informer for AzureAccessKeyRequest type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredAzureAccessKeyRequestInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.EngineV1alpha1().AzureAccessKeyRequests(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.EngineV1alpha1().AzureAccessKeyRequests(namespace).Watch(options)
			},
		},
		&enginev1alpha1.AzureAccessKeyRequest{},
		resyncPeriod,
		indexers,
	)
}

func (f *azureAccessKeyRequestInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredAzureAccessKeyRequestInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *azureAccessKeyRequestInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&enginev1alpha1.AzureAccessKeyRequest{}, f.defaultInformer)
}

func (f *azureAccessKeyRequestInformer) Lister() v1alpha1.AzureAccessKeyRequestLister {
	return v1alpha1.NewAzureAccessKeyRequestLister(f.Informer().GetIndexer())
}
