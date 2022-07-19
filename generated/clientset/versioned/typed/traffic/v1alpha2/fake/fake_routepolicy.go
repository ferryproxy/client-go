/*
Copyright 2022 FerryProxy Authors.

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

	v1alpha2 "github.com/ferryproxy/api/apis/traffic/v1alpha2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeRoutePolicies implements RoutePolicyInterface
type FakeRoutePolicies struct {
	Fake *FakeTrafficV1alpha2
	ns   string
}

var routepoliciesResource = schema.GroupVersionResource{Group: "traffic", Version: "v1alpha2", Resource: "routepolicies"}

var routepoliciesKind = schema.GroupVersionKind{Group: "traffic", Version: "v1alpha2", Kind: "RoutePolicy"}

// Get takes name of the routePolicy, and returns the corresponding routePolicy object, and an error if there is any.
func (c *FakeRoutePolicies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha2.RoutePolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(routepoliciesResource, c.ns, name), &v1alpha2.RoutePolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.RoutePolicy), err
}

// List takes label and field selectors, and returns the list of RoutePolicies that match those selectors.
func (c *FakeRoutePolicies) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha2.RoutePolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(routepoliciesResource, routepoliciesKind, c.ns, opts), &v1alpha2.RoutePolicyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha2.RoutePolicyList{ListMeta: obj.(*v1alpha2.RoutePolicyList).ListMeta}
	for _, item := range obj.(*v1alpha2.RoutePolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested routePolicies.
func (c *FakeRoutePolicies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(routepoliciesResource, c.ns, opts))

}

// Create takes the representation of a routePolicy and creates it.  Returns the server's representation of the routePolicy, and an error, if there is any.
func (c *FakeRoutePolicies) Create(ctx context.Context, routePolicy *v1alpha2.RoutePolicy, opts v1.CreateOptions) (result *v1alpha2.RoutePolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(routepoliciesResource, c.ns, routePolicy), &v1alpha2.RoutePolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.RoutePolicy), err
}

// Update takes the representation of a routePolicy and updates it. Returns the server's representation of the routePolicy, and an error, if there is any.
func (c *FakeRoutePolicies) Update(ctx context.Context, routePolicy *v1alpha2.RoutePolicy, opts v1.UpdateOptions) (result *v1alpha2.RoutePolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(routepoliciesResource, c.ns, routePolicy), &v1alpha2.RoutePolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.RoutePolicy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeRoutePolicies) UpdateStatus(ctx context.Context, routePolicy *v1alpha2.RoutePolicy, opts v1.UpdateOptions) (*v1alpha2.RoutePolicy, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(routepoliciesResource, "status", c.ns, routePolicy), &v1alpha2.RoutePolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.RoutePolicy), err
}

// Delete takes name of the routePolicy and deletes it. Returns an error if one occurs.
func (c *FakeRoutePolicies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(routepoliciesResource, c.ns, name, opts), &v1alpha2.RoutePolicy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRoutePolicies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(routepoliciesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha2.RoutePolicyList{})
	return err
}

// Patch applies the patch and returns the patched routePolicy.
func (c *FakeRoutePolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha2.RoutePolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(routepoliciesResource, c.ns, name, pt, data, subresources...), &v1alpha2.RoutePolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.RoutePolicy), err
}
