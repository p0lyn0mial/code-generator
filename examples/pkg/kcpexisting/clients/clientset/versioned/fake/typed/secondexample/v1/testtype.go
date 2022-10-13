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

	secondexamplev1 "acme.corp/pkg/apis/secondexample/v1"
	applyconfigurationssecondexamplev1 "acme.corp/pkg/generated/applyconfigurations/secondexample/v1"
	secondexamplev1client "acme.corp/pkg/generated/clientset/versioned/typed/secondexample/v1"
	kcpsecondexamplev1 "acme.corp/pkg/kcpexisting/clients/clientset/versioned/typed/secondexample/v1"
)

var testTypesResource = schema.GroupVersionResource{Group: "secondexample", Version: "v1", Resource: "testtypes"}
var testTypesKind = schema.GroupVersionKind{Group: "secondexample", Version: "v1", Kind: "TestType"}

type testTypesClusterClient struct {
	*kcptesting.Fake
}

// Cluster scopes the client down to a particular cluster.
func (c *testTypesClusterClient) Cluster(cluster logicalcluster.Name) kcpsecondexamplev1.TestTypesNamespacer {
	if cluster == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &testTypesNamespacer{Fake: c.Fake, Cluster: cluster}
}

// List takes label and field selectors, and returns the list of TestTypes that match those selectors across all clusters.
func (c *testTypesClusterClient) List(ctx context.Context, opts metav1.ListOptions) (*secondexamplev1.TestTypeList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewListAction(testTypesResource, testTypesKind, logicalcluster.Wildcard, metav1.NamespaceAll, opts), &secondexamplev1.TestTypeList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &secondexamplev1.TestTypeList{ListMeta: obj.(*secondexamplev1.TestTypeList).ListMeta}
	for _, item := range obj.(*secondexamplev1.TestTypeList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested TestTypes across all clusters.
func (c *testTypesClusterClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewWatchAction(testTypesResource, logicalcluster.Wildcard, metav1.NamespaceAll, opts))
}

type testTypesNamespacer struct {
	*kcptesting.Fake
	Cluster logicalcluster.Name
}

func (n *testTypesNamespacer) Namespace(namespace string) secondexamplev1client.TestTypeInterface {
	return &testTypesClient{Fake: n.Fake, Cluster: n.Cluster, Namespace: namespace}
}

type testTypesClient struct {
	*kcptesting.Fake
	Cluster   logicalcluster.Name
	Namespace string
}

func (c *testTypesClient) Create(ctx context.Context, testType *secondexamplev1.TestType, opts metav1.CreateOptions) (*secondexamplev1.TestType, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewCreateAction(testTypesResource, c.Cluster, c.Namespace, testType), &secondexamplev1.TestType{})
	if obj == nil {
		return nil, err
	}
	return obj.(*secondexamplev1.TestType), err
}

func (c *testTypesClient) Update(ctx context.Context, testType *secondexamplev1.TestType, opts metav1.UpdateOptions) (*secondexamplev1.TestType, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewUpdateAction(testTypesResource, c.Cluster, c.Namespace, testType), &secondexamplev1.TestType{})
	if obj == nil {
		return nil, err
	}
	return obj.(*secondexamplev1.TestType), err
}

func (c *testTypesClient) UpdateStatus(ctx context.Context, testType *secondexamplev1.TestType, opts metav1.UpdateOptions) (*secondexamplev1.TestType, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewUpdateSubresourceAction(testTypesResource, c.Cluster, "status", c.Namespace, testType), &secondexamplev1.TestType{})
	if obj == nil {
		return nil, err
	}
	return obj.(*secondexamplev1.TestType), err
}

func (c *testTypesClient) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.Invokes(kcptesting.NewDeleteActionWithOptions(testTypesResource, c.Cluster, c.Namespace, name, opts), &secondexamplev1.TestType{})
	return err
}

func (c *testTypesClient) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := kcptesting.NewDeleteCollectionAction(testTypesResource, c.Cluster, c.Namespace, listOpts)

	_, err := c.Fake.Invokes(action, &secondexamplev1.TestTypeList{})
	return err
}

func (c *testTypesClient) Get(ctx context.Context, name string, options metav1.GetOptions) (*secondexamplev1.TestType, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewGetAction(testTypesResource, c.Cluster, c.Namespace, name), &secondexamplev1.TestType{})
	if obj == nil {
		return nil, err
	}
	return obj.(*secondexamplev1.TestType), err
}

// List takes label and field selectors, and returns the list of TestTypes that match those selectors.
func (c *testTypesClient) List(ctx context.Context, opts metav1.ListOptions) (*secondexamplev1.TestTypeList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewListAction(testTypesResource, testTypesKind, c.Cluster, c.Namespace, opts), &secondexamplev1.TestTypeList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &secondexamplev1.TestTypeList{ListMeta: obj.(*secondexamplev1.TestTypeList).ListMeta}
	for _, item := range obj.(*secondexamplev1.TestTypeList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *testTypesClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewWatchAction(testTypesResource, c.Cluster, c.Namespace, opts))
}

func (c *testTypesClient) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*secondexamplev1.TestType, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(testTypesResource, c.Cluster, c.Namespace, name, pt, data, subresources...), &secondexamplev1.TestType{})
	if obj == nil {
		return nil, err
	}
	return obj.(*secondexamplev1.TestType), err
}

func (c *testTypesClient) Apply(ctx context.Context, applyConfiguration *applyconfigurationssecondexamplev1.TestTypeApplyConfiguration, opts metav1.ApplyOptions) (*secondexamplev1.TestType, error) {
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
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(testTypesResource, c.Cluster, c.Namespace, *name, types.ApplyPatchType, data), &secondexamplev1.TestType{})
	if obj == nil {
		return nil, err
	}
	return obj.(*secondexamplev1.TestType), err
}

func (c *testTypesClient) ApplyStatus(ctx context.Context, applyConfiguration *applyconfigurationssecondexamplev1.TestTypeApplyConfiguration, opts metav1.ApplyOptions) (*secondexamplev1.TestType, error) {
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
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(testTypesResource, c.Cluster, c.Namespace, *name, types.ApplyPatchType, data, "status"), &secondexamplev1.TestType{})
	if obj == nil {
		return nil, err
	}
	return obj.(*secondexamplev1.TestType), err
}
