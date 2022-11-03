//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The KCP Authors.

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

// Code generated by kcp code-generator. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	kcpcache "github.com/kcp-dev/apimachinery/pkg/cache"
	kcpinformers "github.com/kcp-dev/apimachinery/third_party/informers"
	"github.com/kcp-dev/logicalcluster/v2"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"

	example3v1 "acme.corp/pkg/apis/example3/v1"
	scopedclientset "acme.corp/pkg/generated/clientset/versioned"
	clientset "acme.corp/pkg/kcp/clients/clientset/versioned"
	"acme.corp/pkg/kcp/clients/informers/externalversions/internalinterfaces"
	example3v1listers "acme.corp/pkg/kcp/clients/listers/example3/v1"
)

// TestTypeClusterInformer provides access to a shared informer and lister for
// TestTypes.
type TestTypeClusterInformer interface {
	Cluster(logicalcluster.Name) TestTypeInformer
	Informer() kcpcache.ScopeableSharedIndexInformer
	Lister() example3v1listers.TestTypeClusterLister
}

type testTypeClusterInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewTestTypeClusterInformer constructs a new informer for TestType type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewTestTypeClusterInformer(client clientset.ClusterInterface, resyncPeriod time.Duration, indexers cache.Indexers) kcpcache.ScopeableSharedIndexInformer {
	return NewFilteredTestTypeClusterInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredTestTypeClusterInformer constructs a new informer for TestType type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredTestTypeClusterInformer(client clientset.ClusterInterface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) kcpcache.ScopeableSharedIndexInformer {
	return kcpinformers.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Example3V1().TestTypes().List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Example3V1().TestTypes().Watch(context.TODO(), options)
			},
		},
		&example3v1.TestType{},
		resyncPeriod,
		indexers,
	)
}

func (f *testTypeClusterInformer) defaultInformer(client clientset.ClusterInterface, resyncPeriod time.Duration) kcpcache.ScopeableSharedIndexInformer {
	return NewFilteredTestTypeClusterInformer(client, resyncPeriod, cache.Indexers{
		kcpcache.ClusterIndexName:             kcpcache.ClusterIndexFunc,
		kcpcache.ClusterAndNamespaceIndexName: kcpcache.ClusterAndNamespaceIndexFunc},
		f.tweakListOptions,
	)
}

func (f *testTypeClusterInformer) Informer() kcpcache.ScopeableSharedIndexInformer {
	return f.factory.InformerFor(&example3v1.TestType{}, f.defaultInformer)
}

func (f *testTypeClusterInformer) Lister() example3v1listers.TestTypeClusterLister {
	return example3v1listers.NewTestTypeClusterLister(f.Informer().GetIndexer())
}

// TestTypeInformer provides access to a shared informer and lister for
// TestTypes.
type TestTypeInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() example3v1listers.TestTypeLister
}

func (f *testTypeClusterInformer) Cluster(cluster logicalcluster.Name) TestTypeInformer {
	return &testTypeInformer{
		informer: f.Informer().Cluster(cluster),
		lister:   f.Lister().Cluster(cluster),
	}
}

type testTypeInformer struct {
	informer cache.SharedIndexInformer
	lister   example3v1listers.TestTypeLister
}

func (f *testTypeInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

func (f *testTypeInformer) Lister() example3v1listers.TestTypeLister {
	return f.lister
}

type testTypeScopedInformer struct {
	factory          internalinterfaces.SharedScopedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

func (f *testTypeScopedInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&example3v1.TestType{}, f.defaultInformer)
}

func (f *testTypeScopedInformer) Lister() example3v1listers.TestTypeLister {
	return example3v1listers.NewTestTypeLister(f.Informer().GetIndexer())
}

// NewTestTypeInformer constructs a new informer for TestType type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewTestTypeInformer(client scopedclientset.Interface, resyncPeriod time.Duration, namespace string, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredTestTypeInformer(client, resyncPeriod, namespace, indexers, nil)
}

// NewFilteredTestTypeInformer constructs a new informer for TestType type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredTestTypeInformer(client scopedclientset.Interface, resyncPeriod time.Duration, namespace string, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Example3V1().TestTypes(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Example3V1().TestTypes(namespace).Watch(context.TODO(), options)
			},
		},
		&example3v1.TestType{},
		resyncPeriod,
		indexers,
	)
}

func (f *testTypeScopedInformer) defaultInformer(client scopedclientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredTestTypeInformer(client, resyncPeriod, f.namespace, cache.Indexers{
		cache.NamespaceIndex: cache.MetaNamespaceIndexFunc,
	}, f.tweakListOptions)
}
