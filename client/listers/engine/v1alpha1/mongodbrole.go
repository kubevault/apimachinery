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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1alpha1 "kubevault.dev/apimachinery/apis/engine/v1alpha1"
)

// MongoDBRoleLister helps list MongoDBRoles.
// All objects returned here must be treated as read-only.
type MongoDBRoleLister interface {
	// List lists all MongoDBRoles in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.MongoDBRole, err error)
	// MongoDBRoles returns an object that can list and get MongoDBRoles.
	MongoDBRoles(namespace string) MongoDBRoleNamespaceLister
	MongoDBRoleListerExpansion
}

// mongoDBRoleLister implements the MongoDBRoleLister interface.
type mongoDBRoleLister struct {
	indexer cache.Indexer
}

// NewMongoDBRoleLister returns a new MongoDBRoleLister.
func NewMongoDBRoleLister(indexer cache.Indexer) MongoDBRoleLister {
	return &mongoDBRoleLister{indexer: indexer}
}

// List lists all MongoDBRoles in the indexer.
func (s *mongoDBRoleLister) List(selector labels.Selector) (ret []*v1alpha1.MongoDBRole, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.MongoDBRole))
	})
	return ret, err
}

// MongoDBRoles returns an object that can list and get MongoDBRoles.
func (s *mongoDBRoleLister) MongoDBRoles(namespace string) MongoDBRoleNamespaceLister {
	return mongoDBRoleNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// MongoDBRoleNamespaceLister helps list and get MongoDBRoles.
// All objects returned here must be treated as read-only.
type MongoDBRoleNamespaceLister interface {
	// List lists all MongoDBRoles in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.MongoDBRole, err error)
	// Get retrieves the MongoDBRole from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.MongoDBRole, error)
	MongoDBRoleNamespaceListerExpansion
}

// mongoDBRoleNamespaceLister implements the MongoDBRoleNamespaceLister
// interface.
type mongoDBRoleNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all MongoDBRoles in the indexer for a given namespace.
func (s mongoDBRoleNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.MongoDBRole, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.MongoDBRole))
	})
	return ret, err
}

// Get retrieves the MongoDBRole from the indexer for a given namespace and name.
func (s mongoDBRoleNamespaceLister) Get(name string) (*v1alpha1.MongoDBRole, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("mongodbrole"), name)
	}
	return obj.(*v1alpha1.MongoDBRole), nil
}
