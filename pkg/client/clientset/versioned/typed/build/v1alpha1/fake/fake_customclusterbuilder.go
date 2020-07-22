/*
 * Copyright 2019 The original author or authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/pivotal/kpack/pkg/apis/build/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCustomClusterBuilders implements CustomClusterBuilderInterface
type FakeCustomClusterBuilders struct {
	Fake *FakeKpackV1alpha1
}

var customclusterbuildersResource = schema.GroupVersionResource{Group: "kpack.io", Version: "v1alpha1", Resource: "customclusterbuilders"}

var customclusterbuildersKind = schema.GroupVersionKind{Group: "kpack.io", Version: "v1alpha1", Kind: "CustomClusterBuilder"}

// Get takes name of the customClusterBuilder, and returns the corresponding customClusterBuilder object, and an error if there is any.
func (c *FakeCustomClusterBuilders) Get(name string, options v1.GetOptions) (result *v1alpha1.CustomClusterBuilder, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(customclusterbuildersResource, name), &v1alpha1.CustomClusterBuilder{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CustomClusterBuilder), err
}

// List takes label and field selectors, and returns the list of CustomClusterBuilders that match those selectors.
func (c *FakeCustomClusterBuilders) List(opts v1.ListOptions) (result *v1alpha1.CustomClusterBuilderList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(customclusterbuildersResource, customclusterbuildersKind, opts), &v1alpha1.CustomClusterBuilderList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.CustomClusterBuilderList{ListMeta: obj.(*v1alpha1.CustomClusterBuilderList).ListMeta}
	for _, item := range obj.(*v1alpha1.CustomClusterBuilderList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested customClusterBuilders.
func (c *FakeCustomClusterBuilders) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(customclusterbuildersResource, opts))
}

// Create takes the representation of a customClusterBuilder and creates it.  Returns the server's representation of the customClusterBuilder, and an error, if there is any.
func (c *FakeCustomClusterBuilders) Create(customClusterBuilder *v1alpha1.CustomClusterBuilder) (result *v1alpha1.CustomClusterBuilder, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(customclusterbuildersResource, customClusterBuilder), &v1alpha1.CustomClusterBuilder{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CustomClusterBuilder), err
}

// Update takes the representation of a customClusterBuilder and updates it. Returns the server's representation of the customClusterBuilder, and an error, if there is any.
func (c *FakeCustomClusterBuilders) Update(customClusterBuilder *v1alpha1.CustomClusterBuilder) (result *v1alpha1.CustomClusterBuilder, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(customclusterbuildersResource, customClusterBuilder), &v1alpha1.CustomClusterBuilder{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CustomClusterBuilder), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCustomClusterBuilders) UpdateStatus(customClusterBuilder *v1alpha1.CustomClusterBuilder) (*v1alpha1.CustomClusterBuilder, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(customclusterbuildersResource, "status", customClusterBuilder), &v1alpha1.CustomClusterBuilder{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CustomClusterBuilder), err
}

// Delete takes name of the customClusterBuilder and deletes it. Returns an error if one occurs.
func (c *FakeCustomClusterBuilders) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(customclusterbuildersResource, name), &v1alpha1.CustomClusterBuilder{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCustomClusterBuilders) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(customclusterbuildersResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.CustomClusterBuilderList{})
	return err
}

// Patch applies the patch and returns the patched customClusterBuilder.
func (c *FakeCustomClusterBuilders) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CustomClusterBuilder, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(customclusterbuildersResource, name, pt, data, subresources...), &v1alpha1.CustomClusterBuilder{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CustomClusterBuilder), err
}