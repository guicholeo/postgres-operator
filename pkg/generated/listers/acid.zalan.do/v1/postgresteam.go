/*
Copyright 2020 Compose, Zalando SE

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/zalando/postgres-operator/pkg/apis/acid.zalan.do/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// PostgresTeamLister helps list PostgresTeams.
// All objects returned here must be treated as read-only.
type PostgresTeamLister interface {
	// List lists all PostgresTeams in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.PostgresTeam, err error)
	// PostgresTeams returns an object that can list and get PostgresTeams.
	PostgresTeams(namespace string) PostgresTeamNamespaceLister
	PostgresTeamListerExpansion
}

// postgresTeamLister implements the PostgresTeamLister interface.
type postgresTeamLister struct {
	indexer cache.Indexer
}

// NewPostgresTeamLister returns a new PostgresTeamLister.
func NewPostgresTeamLister(indexer cache.Indexer) PostgresTeamLister {
	return &postgresTeamLister{indexer: indexer}
}

// List lists all PostgresTeams in the indexer.
func (s *postgresTeamLister) List(selector labels.Selector) (ret []*v1.PostgresTeam, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.PostgresTeam))
	})
	return ret, err
}

// PostgresTeams returns an object that can list and get PostgresTeams.
func (s *postgresTeamLister) PostgresTeams(namespace string) PostgresTeamNamespaceLister {
	return postgresTeamNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// PostgresTeamNamespaceLister helps list and get PostgresTeams.
// All objects returned here must be treated as read-only.
type PostgresTeamNamespaceLister interface {
	// List lists all PostgresTeams in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.PostgresTeam, err error)
	// Get retrieves the PostgresTeam from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.PostgresTeam, error)
	PostgresTeamNamespaceListerExpansion
}

// postgresTeamNamespaceLister implements the PostgresTeamNamespaceLister
// interface.
type postgresTeamNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all PostgresTeams in the indexer for a given namespace.
func (s postgresTeamNamespaceLister) List(selector labels.Selector) (ret []*v1.PostgresTeam, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.PostgresTeam))
	})
	return ret, err
}

// Get retrieves the PostgresTeam from the indexer for a given namespace and name.
func (s postgresTeamNamespaceLister) Get(name string) (*v1.PostgresTeam, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("postgresteam"), name)
	}
	return obj.(*v1.PostgresTeam), nil
}
