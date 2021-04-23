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

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	enginev1alpha1 "kubevault.dev/apimachinery/apis/engine/v1alpha1"
	versioned "kubevault.dev/apimachinery/client/clientset/versioned"
	internalinterfaces "kubevault.dev/apimachinery/client/informers/externalversions/internalinterfaces"
	v1alpha1 "kubevault.dev/apimachinery/client/listers/engine/v1alpha1"
)

// MySQLRoleInformer provides access to a shared informer and lister for
// MySQLRoles.
type MySQLRoleInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.MySQLRoleLister
}

type mySQLRoleInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewMySQLRoleInformer constructs a new informer for MySQLRole type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewMySQLRoleInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredMySQLRoleInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredMySQLRoleInformer constructs a new informer for MySQLRole type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredMySQLRoleInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.EngineV1alpha1().MySQLRoles(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.EngineV1alpha1().MySQLRoles(namespace).Watch(context.TODO(), options)
			},
		},
		&enginev1alpha1.MySQLRole{},
		resyncPeriod,
		indexers,
	)
}

func (f *mySQLRoleInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredMySQLRoleInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *mySQLRoleInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&enginev1alpha1.MySQLRole{}, f.defaultInformer)
}

func (f *mySQLRoleInformer) Lister() v1alpha1.MySQLRoleLister {
	return v1alpha1.NewMySQLRoleLister(f.Informer().GetIndexer())
}
