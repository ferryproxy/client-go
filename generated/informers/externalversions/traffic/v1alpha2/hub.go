/*
Copyright 2022 FerryProxy.

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

package v1alpha2

import (
	"context"
	time "time"

	trafficv1alpha2 "github.com/ferryproxy/api/apis/traffic/v1alpha2"
	versioned "github.com/ferryproxy/client-go/generated/clientset/versioned"
	internalinterfaces "github.com/ferryproxy/client-go/generated/informers/externalversions/internalinterfaces"
	v1alpha2 "github.com/ferryproxy/client-go/generated/listers/traffic/v1alpha2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// HubInformer provides access to a shared informer and lister for
// Hubs.
type HubInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha2.HubLister
}

type hubInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewHubInformer constructs a new informer for Hub type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewHubInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredHubInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredHubInformer constructs a new informer for Hub type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredHubInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TrafficV1alpha2().Hubs(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TrafficV1alpha2().Hubs(namespace).Watch(context.TODO(), options)
			},
		},
		&trafficv1alpha2.Hub{},
		resyncPeriod,
		indexers,
	)
}

func (f *hubInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredHubInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *hubInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&trafficv1alpha2.Hub{}, f.defaultInformer)
}

func (f *hubInformer) Lister() v1alpha2.HubLister {
	return v1alpha2.NewHubLister(f.Informer().GetIndexer())
}
