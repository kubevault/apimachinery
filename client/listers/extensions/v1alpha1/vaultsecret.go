/*
Copyright 2018 The Vault Operator Authors.

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

// This file was automatically generated by lister-gen

package v1alpha1

import (
	v1alpha1 "github.com/soter/vault-operator/apis/extensions/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// VaultSecretLister helps list VaultSecrets.
type VaultSecretLister interface {
	// List lists all VaultSecrets in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.VaultSecret, err error)
	// VaultSecrets returns an object that can list and get VaultSecrets.
	VaultSecrets(namespace string) VaultSecretNamespaceLister
	VaultSecretListerExpansion
}

// vaultSecretLister implements the VaultSecretLister interface.
type vaultSecretLister struct {
	indexer cache.Indexer
}

// NewVaultSecretLister returns a new VaultSecretLister.
func NewVaultSecretLister(indexer cache.Indexer) VaultSecretLister {
	return &vaultSecretLister{indexer: indexer}
}

// List lists all VaultSecrets in the indexer.
func (s *vaultSecretLister) List(selector labels.Selector) (ret []*v1alpha1.VaultSecret, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.VaultSecret))
	})
	return ret, err
}

// VaultSecrets returns an object that can list and get VaultSecrets.
func (s *vaultSecretLister) VaultSecrets(namespace string) VaultSecretNamespaceLister {
	return vaultSecretNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// VaultSecretNamespaceLister helps list and get VaultSecrets.
type VaultSecretNamespaceLister interface {
	// List lists all VaultSecrets in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.VaultSecret, err error)
	// Get retrieves the VaultSecret from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.VaultSecret, error)
	VaultSecretNamespaceListerExpansion
}

// vaultSecretNamespaceLister implements the VaultSecretNamespaceLister
// interface.
type vaultSecretNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all VaultSecrets in the indexer for a given namespace.
func (s vaultSecretNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.VaultSecret, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.VaultSecret))
	})
	return ret, err
}

// Get retrieves the VaultSecret from the indexer for a given namespace and name.
func (s vaultSecretNamespaceLister) Get(name string) (*v1alpha1.VaultSecret, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("vaultsecret"), name)
	}
	return obj.(*v1alpha1.VaultSecret), nil
}
