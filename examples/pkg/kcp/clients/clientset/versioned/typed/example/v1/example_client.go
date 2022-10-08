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
	"net/http"

	kcpclient "github.com/kcp-dev/apimachinery/pkg/client"
	"github.com/kcp-dev/logicalcluster/v2"

	"k8s.io/client-go/rest"

	examplev1 "acme.corp/pkg/generated/clientset/versioned/typed/example/v1"
)

type ExampleV1ClusterInterface interface {
	ExampleV1ClusterScoper
	TestTypesClusterGetter
	ClusterTestTypesClusterGetter
	WithoutVerbTypesClusterGetter
}

type ExampleV1ClusterScoper interface {
	Cluster(logicalcluster.Name) examplev1.ExampleV1Interface
}

type ExampleV1ClusterClient struct {
	clientCache kcpclient.Cache[*examplev1.ExampleV1Client]
}

func (c *ExampleV1ClusterClient) Cluster(name logicalcluster.Name) examplev1.ExampleV1Interface {
	if name == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}
	return c.clientCache.ClusterOrDie(name)
}

func (c *ExampleV1ClusterClient) TestTypes() TestTypeClusterInterface {
	return &testTypesClusterInterface{clientCache: c.clientCache}
}

func (c *ExampleV1ClusterClient) ClusterTestTypes() ClusterTestTypeClusterInterface {
	return &clusterTestTypesClusterInterface{clientCache: c.clientCache}
}

func (c *ExampleV1ClusterClient) WithoutVerbTypes() WithoutVerbTypeClusterInterface {
	return &withoutVerbTypesClusterInterface{clientCache: c.clientCache}
}

// NewForConfig creates a new ExampleV1ClusterClient for the given config.
// NewForConfig is equivalent to NewForConfigAndClient(c, httpClient),
// where httpClient was generated with rest.HTTPClientFor(c).
func NewForConfig(c *rest.Config) (*ExampleV1ClusterClient, error) {
	client, err := rest.HTTPClientFor(c)
	if err != nil {
		return nil, err
	}
	return NewForConfigAndClient(c, client)
}

// NewForConfigAndClient creates a new ExampleV1ClusterClient for the given config and http client.
// Note the http client provided takes precedence over the configured transport values.
func NewForConfigAndClient(c *rest.Config, h *http.Client) (*ExampleV1ClusterClient, error) {
	cache := kcpclient.NewCache(c, h, &kcpclient.Constructor[*examplev1.ExampleV1Client]{
		NewForConfigAndClient: examplev1.NewForConfigAndClient,
	})
	if _, err := cache.Cluster(logicalcluster.New("root")); err != nil {
		return nil, err
	}
	return &ExampleV1ClusterClient{clientCache: cache}, nil
}

// NewForConfigOrDie creates a new ExampleV1ClusterClient for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *ExampleV1ClusterClient {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}
