// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	versioned "github.com/open-cluster-management/api/client/cluster/clientset/versioned"
	internalinterfaces "github.com/open-cluster-management/api/client/cluster/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/open-cluster-management/api/client/cluster/listers/cluster/v1alpha1"
	clusterv1alpha1 "github.com/open-cluster-management/api/cluster/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ManagedClusterSetInformer provides access to a shared informer and lister for
// ManagedClusterSets.
type ManagedClusterSetInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ManagedClusterSetLister
}

type managedClusterSetInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewManagedClusterSetInformer constructs a new informer for ManagedClusterSet type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewManagedClusterSetInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredManagedClusterSetInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredManagedClusterSetInformer constructs a new informer for ManagedClusterSet type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredManagedClusterSetInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ClusterV1alpha1().ManagedClusterSets().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ClusterV1alpha1().ManagedClusterSets().Watch(context.TODO(), options)
			},
		},
		&clusterv1alpha1.ManagedClusterSet{},
		resyncPeriod,
		indexers,
	)
}

func (f *managedClusterSetInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredManagedClusterSetInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *managedClusterSetInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&clusterv1alpha1.ManagedClusterSet{}, f.defaultInformer)
}

func (f *managedClusterSetInformer) Lister() v1alpha1.ManagedClusterSetLister {
	return v1alpha1.NewManagedClusterSetLister(f.Informer().GetIndexer())
}