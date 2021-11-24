/*
Copyright The Kubernetes Authors.

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
	v1alpha1 "redis-priv-operator/pkg/apis/cache/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// RedisStandaloneLister helps list RedisStandalones.
type RedisStandaloneLister interface {
	// List lists all RedisStandalones in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.RedisStandalone, err error)
	// RedisStandalones returns an object that can list and get RedisStandalones.
	RedisStandalones(namespace string) RedisStandaloneNamespaceLister
	RedisStandaloneListerExpansion
}

// redisStandaloneLister implements the RedisStandaloneLister interface.
type redisStandaloneLister struct {
	indexer cache.Indexer
}

// NewRedisStandaloneLister returns a new RedisStandaloneLister.
func NewRedisStandaloneLister(indexer cache.Indexer) RedisStandaloneLister {
	return &redisStandaloneLister{indexer: indexer}
}

// List lists all RedisStandalones in the indexer.
func (s *redisStandaloneLister) List(selector labels.Selector) (ret []*v1alpha1.RedisStandalone, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.RedisStandalone))
	})
	return ret, err
}

// RedisStandalones returns an object that can list and get RedisStandalones.
func (s *redisStandaloneLister) RedisStandalones(namespace string) RedisStandaloneNamespaceLister {
	return redisStandaloneNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// RedisStandaloneNamespaceLister helps list and get RedisStandalones.
type RedisStandaloneNamespaceLister interface {
	// List lists all RedisStandalones in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.RedisStandalone, err error)
	// Get retrieves the RedisStandalone from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.RedisStandalone, error)
	RedisStandaloneNamespaceListerExpansion
}

// redisStandaloneNamespaceLister implements the RedisStandaloneNamespaceLister
// interface.
type redisStandaloneNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all RedisStandalones in the indexer for a given namespace.
func (s redisStandaloneNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.RedisStandalone, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.RedisStandalone))
	})
	return ret, err
}

// Get retrieves the RedisStandalone from the indexer for a given namespace and name.
func (s redisStandaloneNamespaceLister) Get(name string) (*v1alpha1.RedisStandalone, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("redisstandalone"), name)
	}
	return obj.(*v1alpha1.RedisStandalone), nil
}
