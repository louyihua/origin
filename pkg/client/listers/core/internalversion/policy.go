// This file was automatically generated by lister-gen with arguments: --input-dirs=[github.com/openshift/origin/pkg/authorization/api,github.com/openshift/origin/pkg/authorization/api/v1,github.com/openshift/origin/pkg/build/api,github.com/openshift/origin/pkg/build/api/v1,github.com/openshift/origin/pkg/deploy/api,github.com/openshift/origin/pkg/deploy/api/v1,github.com/openshift/origin/pkg/image/api,github.com/openshift/origin/pkg/image/api/v1,github.com/openshift/origin/pkg/oauth/api,github.com/openshift/origin/pkg/oauth/api/v1,github.com/openshift/origin/pkg/project/api,github.com/openshift/origin/pkg/project/api/v1,github.com/openshift/origin/pkg/quota/api,github.com/openshift/origin/pkg/quota/api/v1,github.com/openshift/origin/pkg/route/api,github.com/openshift/origin/pkg/route/api/v1,github.com/openshift/origin/pkg/sdn/api,github.com/openshift/origin/pkg/sdn/api/v1,github.com/openshift/origin/pkg/template/api,github.com/openshift/origin/pkg/template/api/v1,github.com/openshift/origin/pkg/user/api,github.com/openshift/origin/pkg/user/api/v1] --logtostderr=true

package internalversion

import (
	api "github.com/openshift/origin/pkg/authorization/api"
	"k8s.io/kubernetes/pkg/api/errors"
	"k8s.io/kubernetes/pkg/client/cache"
	"k8s.io/kubernetes/pkg/labels"
)

// PolicyLister helps list Policies.
type PolicyLister interface {
	// List lists all Policies in the indexer.
	List(selector labels.Selector) (ret []*api.Policy, err error)
	// Policies returns an object that can list and get Policies.
	Policies(namespace string) PolicyNamespaceLister
	PolicyListerExpansion
}

// policyLister implements the PolicyLister interface.
type policyLister struct {
	indexer cache.Indexer
}

// NewPolicyLister returns a new PolicyLister.
func NewPolicyLister(indexer cache.Indexer) PolicyLister {
	return &policyLister{indexer: indexer}
}

// List lists all Policies in the indexer.
func (s *policyLister) List(selector labels.Selector) (ret []*api.Policy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*api.Policy))
	})
	return ret, err
}

// Policies returns an object that can list and get Policies.
func (s *policyLister) Policies(namespace string) PolicyNamespaceLister {
	return policyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// PolicyNamespaceLister helps list and get Policies.
type PolicyNamespaceLister interface {
	// List lists all Policies in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*api.Policy, err error)
	// Get retrieves the Policy from the indexer for a given namespace and name.
	Get(name string) (*api.Policy, error)
	PolicyNamespaceListerExpansion
}

// policyNamespaceLister implements the PolicyNamespaceLister
// interface.
type policyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Policies in the indexer for a given namespace.
func (s policyNamespaceLister) List(selector labels.Selector) (ret []*api.Policy, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*api.Policy))
	})
	return ret, err
}

// Get retrieves the Policy from the indexer for a given namespace and name.
func (s policyNamespaceLister) Get(name string) (*api.Policy, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(api.Resource("policy"), name)
	}
	return obj.(*api.Policy), nil
}
