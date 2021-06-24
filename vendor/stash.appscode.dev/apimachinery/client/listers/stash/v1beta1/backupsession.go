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

package v1beta1

import (
	v1beta1 "stash.appscode.dev/apimachinery/apis/stash/v1beta1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// BackupSessionLister helps list BackupSessions.
// All objects returned here must be treated as read-only.
type BackupSessionLister interface {
	// List lists all BackupSessions in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.BackupSession, err error)
	// BackupSessions returns an object that can list and get BackupSessions.
	BackupSessions(namespace string) BackupSessionNamespaceLister
	BackupSessionListerExpansion
}

// backupSessionLister implements the BackupSessionLister interface.
type backupSessionLister struct {
	indexer cache.Indexer
}

// NewBackupSessionLister returns a new BackupSessionLister.
func NewBackupSessionLister(indexer cache.Indexer) BackupSessionLister {
	return &backupSessionLister{indexer: indexer}
}

// List lists all BackupSessions in the indexer.
func (s *backupSessionLister) List(selector labels.Selector) (ret []*v1beta1.BackupSession, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.BackupSession))
	})
	return ret, err
}

// BackupSessions returns an object that can list and get BackupSessions.
func (s *backupSessionLister) BackupSessions(namespace string) BackupSessionNamespaceLister {
	return backupSessionNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// BackupSessionNamespaceLister helps list and get BackupSessions.
// All objects returned here must be treated as read-only.
type BackupSessionNamespaceLister interface {
	// List lists all BackupSessions in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.BackupSession, err error)
	// Get retrieves the BackupSession from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1beta1.BackupSession, error)
	BackupSessionNamespaceListerExpansion
}

// backupSessionNamespaceLister implements the BackupSessionNamespaceLister
// interface.
type backupSessionNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all BackupSessions in the indexer for a given namespace.
func (s backupSessionNamespaceLister) List(selector labels.Selector) (ret []*v1beta1.BackupSession, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.BackupSession))
	})
	return ret, err
}

// Get retrieves the BackupSession from the indexer for a given namespace and name.
func (s backupSessionNamespaceLister) Get(name string) (*v1beta1.BackupSession, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("backupsession"), name)
	}
	return obj.(*v1beta1.BackupSession), nil
}