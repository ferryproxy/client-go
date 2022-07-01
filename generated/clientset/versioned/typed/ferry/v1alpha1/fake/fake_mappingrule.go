/*
Copyright 2022 Shiming Zhang.

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
// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/ferry-proxy/api/apis/ferry/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeMappingRules implements MappingRuleInterface
type FakeMappingRules struct {
	Fake *FakeFerryV1alpha1
	ns   string
}

var mappingrulesResource = schema.GroupVersionResource{Group: "ferry", Version: "v1alpha1", Resource: "mappingrules"}

var mappingrulesKind = schema.GroupVersionKind{Group: "ferry", Version: "v1alpha1", Kind: "MappingRule"}

// Get takes name of the mappingRule, and returns the corresponding mappingRule object, and an error if there is any.
func (c *FakeMappingRules) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.MappingRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(mappingrulesResource, c.ns, name), &v1alpha1.MappingRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MappingRule), err
}

// List takes label and field selectors, and returns the list of MappingRules that match those selectors.
func (c *FakeMappingRules) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.MappingRuleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(mappingrulesResource, mappingrulesKind, c.ns, opts), &v1alpha1.MappingRuleList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.MappingRuleList{ListMeta: obj.(*v1alpha1.MappingRuleList).ListMeta}
	for _, item := range obj.(*v1alpha1.MappingRuleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested mappingRules.
func (c *FakeMappingRules) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(mappingrulesResource, c.ns, opts))

}

// Create takes the representation of a mappingRule and creates it.  Returns the server's representation of the mappingRule, and an error, if there is any.
func (c *FakeMappingRules) Create(ctx context.Context, mappingRule *v1alpha1.MappingRule, opts v1.CreateOptions) (result *v1alpha1.MappingRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(mappingrulesResource, c.ns, mappingRule), &v1alpha1.MappingRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MappingRule), err
}

// Update takes the representation of a mappingRule and updates it. Returns the server's representation of the mappingRule, and an error, if there is any.
func (c *FakeMappingRules) Update(ctx context.Context, mappingRule *v1alpha1.MappingRule, opts v1.UpdateOptions) (result *v1alpha1.MappingRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(mappingrulesResource, c.ns, mappingRule), &v1alpha1.MappingRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MappingRule), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeMappingRules) UpdateStatus(ctx context.Context, mappingRule *v1alpha1.MappingRule, opts v1.UpdateOptions) (*v1alpha1.MappingRule, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(mappingrulesResource, "status", c.ns, mappingRule), &v1alpha1.MappingRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MappingRule), err
}

// Delete takes name of the mappingRule and deletes it. Returns an error if one occurs.
func (c *FakeMappingRules) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(mappingrulesResource, c.ns, name, opts), &v1alpha1.MappingRule{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMappingRules) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(mappingrulesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.MappingRuleList{})
	return err
}

// Patch applies the patch and returns the patched mappingRule.
func (c *FakeMappingRules) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.MappingRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(mappingrulesResource, c.ns, name, pt, data, subresources...), &v1alpha1.MappingRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MappingRule), err
}
