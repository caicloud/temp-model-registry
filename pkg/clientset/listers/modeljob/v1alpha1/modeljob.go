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
	v1alpha1 "github.com/kleveross/klever-model-registry/pkg/apis/modeljob/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ModelJobLister helps list ModelJobs.
type ModelJobLister interface {
	// List lists all ModelJobs in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ModelJob, err error)
	// ModelJobs returns an object that can list and get ModelJobs.
	ModelJobs(namespace string) ModelJobNamespaceLister
	ModelJobListerExpansion
}

// modelJobLister implements the ModelJobLister interface.
type modelJobLister struct {
	indexer cache.Indexer
}

// NewModelJobLister returns a new ModelJobLister.
func NewModelJobLister(indexer cache.Indexer) ModelJobLister {
	return &modelJobLister{indexer: indexer}
}

// List lists all ModelJobs in the indexer.
func (s *modelJobLister) List(selector labels.Selector) (ret []*v1alpha1.ModelJob, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ModelJob))
	})
	return ret, err
}

// ModelJobs returns an object that can list and get ModelJobs.
func (s *modelJobLister) ModelJobs(namespace string) ModelJobNamespaceLister {
	return modelJobNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ModelJobNamespaceLister helps list and get ModelJobs.
type ModelJobNamespaceLister interface {
	// List lists all ModelJobs in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.ModelJob, err error)
	// Get retrieves the ModelJob from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.ModelJob, error)
	ModelJobNamespaceListerExpansion
}

// modelJobNamespaceLister implements the ModelJobNamespaceLister
// interface.
type modelJobNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ModelJobs in the indexer for a given namespace.
func (s modelJobNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ModelJob, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ModelJob))
	})
	return ret, err
}

// Get retrieves the ModelJob from the indexer for a given namespace and name.
func (s modelJobNamespaceLister) Get(name string) (*v1alpha1.ModelJob, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("modeljob"), name)
	}
	return obj.(*v1alpha1.ModelJob), nil
}
