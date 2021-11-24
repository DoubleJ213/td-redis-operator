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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	tdbv1alpha1 "redis-priv-operator/pkg/apis/tdb/v1alpha1"
	clientset "redis-priv-operator/pkg/client/clientset"
	internalinterfaces "redis-priv-operator/pkg/client/informers/internalinterfaces"
	v1alpha1 "redis-priv-operator/pkg/client/listers/tdb/v1alpha1"
	time "time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// MysqlProxyInformer provides access to a shared informer and lister for
// MysqlProxies.
type MysqlProxyInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.MysqlProxyLister
}

type mysqlProxyInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewMysqlProxyInformer constructs a new informer for MysqlProxy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewMysqlProxyInformer(client clientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredMysqlProxyInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredMysqlProxyInformer constructs a new informer for MysqlProxy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredMysqlProxyInformer(client clientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TdbV1alpha1().MysqlProxies(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TdbV1alpha1().MysqlProxies(namespace).Watch(context.TODO(), options)
			},
		},
		&tdbv1alpha1.MysqlProxy{},
		resyncPeriod,
		indexers,
	)
}

func (f *mysqlProxyInformer) defaultInformer(client clientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredMysqlProxyInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *mysqlProxyInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&tdbv1alpha1.MysqlProxy{}, f.defaultInformer)
}

func (f *mysqlProxyInformer) Lister() v1alpha1.MysqlProxyLister {
	return v1alpha1.NewMysqlProxyLister(f.Informer().GetIndexer())
}
