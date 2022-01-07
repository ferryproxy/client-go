/*
Copyright 2022 Shiming Zhang.

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

	ferryv1alpha1 "github.com/ferry-proxy/api/apis/ferry/v1alpha1"
	versioned "github.com/ferry-proxy/client-go/generated/clientset/versioned"
	internalinterfaces "github.com/ferry-proxy/client-go/generated/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/ferry-proxy/client-go/generated/listers/ferry/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ClusterInformationInformer provides access to a shared informer and lister for
// ClusterInformations.
type ClusterInformationInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ClusterInformationLister
}

type clusterInformationInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewClusterInformationInformer constructs a new informer for ClusterInformation type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewClusterInformationInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredClusterInformationInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredClusterInformationInformer constructs a new informer for ClusterInformation type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredClusterInformationInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.FerryV1alpha1().ClusterInformations(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.FerryV1alpha1().ClusterInformations(namespace).Watch(context.TODO(), options)
			},
		},
		&ferryv1alpha1.ClusterInformation{},
		resyncPeriod,
		indexers,
	)
}

func (f *clusterInformationInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredClusterInformationInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *clusterInformationInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&ferryv1alpha1.ClusterInformation{}, f.defaultInformer)
}

func (f *clusterInformationInformer) Lister() v1alpha1.ClusterInformationLister {
	return v1alpha1.NewClusterInformationLister(f.Informer().GetIndexer())
}
