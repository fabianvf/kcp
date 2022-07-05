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

package v1alpha1

import (
	apimachinerycache "github.com/kcp-dev/apimachinery/pkg/cache"
	"github.com/kcp-dev/logicalcluster"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"

	apiresourcev1alpha1 "github.com/kcp-dev/kcp/pkg/apis/apiresource/v1alpha1"
)

// APIResourceImportLister helps list apiresourcev1alpha1.APIResourceImport.
// All objects returned here must be treated as read-only.
type APIResourceImportClusterLister interface {
	// List lists all apiresourcev1alpha1.APIResourceImport in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*apiresourcev1alpha1.APIResourceImport, err error)

	// Cluster returns an object that can list and get apiresourcev1alpha1.APIResourceImport from the given logical cluster.
	Cluster(cluster logicalcluster.Name) APIResourceImportLister
}

// aPIResourceImportClusterLister implements the APIResourceImportClusterLister interface.
type aPIResourceImportClusterLister struct {
	indexer cache.Indexer
}

// NewAPIResourceImportClusterLister returns a new APIResourceImportClusterLister.
func NewAPIResourceImportClusterLister(indexer cache.Indexer) APIResourceImportClusterLister {
	return &aPIResourceImportClusterLister{indexer: indexer}
}

// List lists all apiresourcev1alpha1.APIResourceImport in the indexer.
func (s *aPIResourceImportClusterLister) List(selector labels.Selector) (ret []*apiresourcev1alpha1.APIResourceImport, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*apiresourcev1alpha1.APIResourceImport))
	})
	return ret, err
}

// Cluster returns an object that can list and get apiresourcev1alpha1.APIResourceImport.
func (s *aPIResourceImportClusterLister) Cluster(cluster logicalcluster.Name) APIResourceImportLister {
	return &aPIResourceImportLister{indexer: s.indexer, cluster: cluster}
}

// APIResourceImportLister helps list apiresourcev1alpha1.APIResourceImport.
// All objects returned here must be treated as read-only.
type APIResourceImportLister interface {
	// List lists all apiresourcev1alpha1.APIResourceImport in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*apiresourcev1alpha1.APIResourceImport, err error)
	// Get retrieves the apiresourcev1alpha1.APIResourceImport from the indexer for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*apiresourcev1alpha1.APIResourceImport, error)
}

// aPIResourceImportLister implements the APIResourceImportLister interface.
type aPIResourceImportLister struct {
	indexer cache.Indexer
	cluster logicalcluster.Name
}

// List lists all apiresourcev1alpha1.APIResourceImport in the indexer.
func (s *aPIResourceImportLister) List(selector labels.Selector) (ret []*apiresourcev1alpha1.APIResourceImport, err error) {
	selectAll := selector == nil || selector.Empty()

	key := apimachinerycache.ToClusterAwareKey(s.cluster.String(), "", "")
	list, err := s.indexer.ByIndex(apimachinerycache.ClusterIndexName, key)
	if err != nil {
		return nil, err
	}

	for i := range list {
		obj := list[i].(*apiresourcev1alpha1.APIResourceImport)
		if selectAll {
			ret = append(ret, obj)
		} else {
			if selector.Matches(labels.Set(obj.GetLabels())) {
				ret = append(ret, obj)
			}
		}
	}

	return ret, err
}

// Get retrieves the apiresourcev1alpha1.APIResourceImport from the indexer for a given name.
func (s aPIResourceImportLister) Get(name string) (*apiresourcev1alpha1.APIResourceImport, error) {
	key := apimachinerycache.ToClusterAwareKey(s.cluster.String(), "", name)
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(apiresourcev1alpha1.Resource("aPIResourceImport"), name)
	}
	return obj.(*apiresourcev1alpha1.APIResourceImport), nil
}
