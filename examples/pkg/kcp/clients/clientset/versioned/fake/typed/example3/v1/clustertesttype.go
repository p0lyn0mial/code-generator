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
	"encoding/json"
	"fmt"

	"github.com/kcp-dev/logicalcluster/v2"

	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/testing"

	example3v1 "acme.corp/pkg/apis/example3/v1"
	applyconfigurationsexample3v1 "acme.corp/pkg/generated/applyconfigurations/example3/v1"
	example3v1client "acme.corp/pkg/generated/clientset/versioned/typed/example3/v1"
)

var clusterTestTypesResource = schema.GroupVersionResource{Group: "example3.some.corp", Version: "V1", Resource: "clustertesttypes"}
var clusterTestTypesKind = schema.GroupVersionKind{Group: "example3.some.corp", Version: "V1", Kind: "ClusterTestType"}

type clusterTestTypesClusterClient struct {
	*kcptesting.Fake
}

// Cluster scopes the client down to a particular cluster.
func (c *clusterTestTypesClusterClient) Cluster(cluster logicalcluster.Name) example3v1client.ClusterTestTypeInterface {
	if cluster == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &clusterTestTypesClient{Fake: c.Fake, Cluster: cluster}
}

// List takes label and field selectors, and returns the list of ClusterTestTypes that match those selectors across all clusters.
func (c *clusterTestTypesClusterClient) List(ctx context.Context, opts metav1.ListOptions) (*example3v1.ClusterTestTypeList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootListAction(clusterTestTypesResource, clusterTestTypesKind, logicalcluster.Wildcard, opts), &example3v1.ClusterTestTypeList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &example3v1.ClusterTestTypeList{ListMeta: obj.(*example3v1.ClusterTestTypeList).ListMeta}
	for _, item := range obj.(*example3v1.ClusterTestTypeList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested ClusterTestTypes across all clusters.
func (c *clusterTestTypesClusterClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewRootWatchAction(clusterTestTypesResource, logicalcluster.Wildcard, opts))
}

type clusterTestTypesClient struct {
	*kcptesting.Fake
	Cluster logicalcluster.Name
}

func (c *clusterTestTypesClient) Create(ctx context.Context, clusterTestType *example3v1.ClusterTestType, opts metav1.CreateOptions) (*example3v1.ClusterTestType, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootCreateAction(clusterTestTypesResource, c.Cluster, clusterTestType), &example3v1.ClusterTestType{})
	if obj == nil {
		return nil, err
	}
	return obj.(*example3v1.ClusterTestType), err
}

func (c *clusterTestTypesClient) Update(ctx context.Context, clusterTestType *example3v1.ClusterTestType, opts metav1.UpdateOptions) (*example3v1.ClusterTestType, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootUpdateAction(clusterTestTypesResource, c.Cluster, clusterTestType), &example3v1.ClusterTestType{})
	if obj == nil {
		return nil, err
	}
	return obj.(*example3v1.ClusterTestType), err
}

func (c *clusterTestTypesClient) UpdateStatus(ctx context.Context, clusterTestType *example3v1.ClusterTestType, opts metav1.UpdateOptions) (*example3v1.ClusterTestType, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootUpdateSubresourceAction(clusterTestTypesResource, c.Cluster, "status", clusterTestType), &example3v1.ClusterTestType{})
	if obj == nil {
		return nil, err
	}
	return obj.(*example3v1.ClusterTestType), err
}

func (c *clusterTestTypesClient) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.Invokes(kcptesting.NewRootDeleteActionWithOptions(clusterTestTypesResource, c.Cluster, name, opts), &example3v1.ClusterTestType{})
	return err
}

func (c *clusterTestTypesClient) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := kcptesting.NewRootDeleteCollectionAction(clusterTestTypesResource, c.Cluster, listOpts)

	_, err := c.Fake.Invokes(action, &example3v1.ClusterTestTypeList{})
	return err
}

func (c *clusterTestTypesClient) Get(ctx context.Context, name string, options metav1.GetOptions) (*example3v1.ClusterTestType, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootGetAction(clusterTestTypesResource, c.Cluster, name), &example3v1.ClusterTestType{})
	if obj == nil {
		return nil, err
	}
	return obj.(*example3v1.ClusterTestType), err
}

// List takes label and field selectors, and returns the list of ClusterTestTypes that match those selectors.
func (c *clusterTestTypesClient) List(ctx context.Context, opts metav1.ListOptions) (*example3v1.ClusterTestTypeList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootListAction(clusterTestTypesResource, clusterTestTypesKind, c.Cluster, opts), &example3v1.ClusterTestTypeList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &example3v1.ClusterTestTypeList{ListMeta: obj.(*example3v1.ClusterTestTypeList).ListMeta}
	for _, item := range obj.(*example3v1.ClusterTestTypeList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *clusterTestTypesClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewRootWatchAction(clusterTestTypesResource, c.Cluster, opts))
}

func (c *clusterTestTypesClient) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*example3v1.ClusterTestType, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction(clusterTestTypesResource, c.Cluster, name, pt, data, subresources...), &example3v1.ClusterTestType{})
	if obj == nil {
		return nil, err
	}
	return obj.(*example3v1.ClusterTestType), err
}

func (c *clusterTestTypesClient) Apply(ctx context.Context, applyConfiguration *applyconfigurationsexample3v1.ClusterTestTypeApplyConfiguration, opts metav1.ApplyOptions) (*example3v1.ClusterTestType, error) {
	if applyConfiguration == nil {
		return nil, fmt.Errorf("applyConfiguration provided to Apply must not be nil")
	}
	data, err := json.Marshal(applyConfiguration)
	if err != nil {
		return nil, err
	}
	name := applyConfiguration.Name
	if name == nil {
		return nil, fmt.Errorf("applyConfiguration.Name must be provided to Apply")
	}
	obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction(clusterTestTypesResource, c.Cluster, *name, types.ApplyPatchType, data), &example3v1.ClusterTestType{})
	if obj == nil {
		return nil, err
	}
	return obj.(*example3v1.ClusterTestType), err
}

func (c *clusterTestTypesClient) ApplyStatus(ctx context.Context, applyConfiguration *applyconfigurationsexample3v1.ClusterTestTypeApplyConfiguration, opts metav1.ApplyOptions) (*example3v1.ClusterTestType, error) {
	if applyConfiguration == nil {
		return nil, fmt.Errorf("applyConfiguration provided to Apply must not be nil")
	}
	data, err := json.Marshal(applyConfiguration)
	if err != nil {
		return nil, err
	}
	name := applyConfiguration.Name
	if name == nil {
		return nil, fmt.Errorf("applyConfiguration.Name must be provided to Apply")
	}
	obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction(clusterTestTypesResource, c.Cluster, *name, types.ApplyPatchType, data, "status"), &example3v1.ClusterTestType{})
	if obj == nil {
		return nil, err
	}
	return obj.(*example3v1.ClusterTestType), err
}
