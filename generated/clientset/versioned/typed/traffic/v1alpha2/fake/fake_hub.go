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

	v1alpha2 "github.com/ferry-proxy/api/apis/traffic/v1alpha2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeHubs implements HubInterface
type FakeHubs struct {
	Fake *FakeTrafficV1alpha2
	ns   string
}

var hubsResource = schema.GroupVersionResource{Group: "traffic", Version: "v1alpha2", Resource: "hubs"}

var hubsKind = schema.GroupVersionKind{Group: "traffic", Version: "v1alpha2", Kind: "Hub"}

// Get takes name of the hub, and returns the corresponding hub object, and an error if there is any.
func (c *FakeHubs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha2.Hub, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(hubsResource, c.ns, name), &v1alpha2.Hub{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.Hub), err
}

// List takes label and field selectors, and returns the list of Hubs that match those selectors.
func (c *FakeHubs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha2.HubList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(hubsResource, hubsKind, c.ns, opts), &v1alpha2.HubList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha2.HubList{ListMeta: obj.(*v1alpha2.HubList).ListMeta}
	for _, item := range obj.(*v1alpha2.HubList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested hubs.
func (c *FakeHubs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(hubsResource, c.ns, opts))

}

// Create takes the representation of a hub and creates it.  Returns the server's representation of the hub, and an error, if there is any.
func (c *FakeHubs) Create(ctx context.Context, hub *v1alpha2.Hub, opts v1.CreateOptions) (result *v1alpha2.Hub, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(hubsResource, c.ns, hub), &v1alpha2.Hub{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.Hub), err
}

// Update takes the representation of a hub and updates it. Returns the server's representation of the hub, and an error, if there is any.
func (c *FakeHubs) Update(ctx context.Context, hub *v1alpha2.Hub, opts v1.UpdateOptions) (result *v1alpha2.Hub, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(hubsResource, c.ns, hub), &v1alpha2.Hub{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.Hub), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeHubs) UpdateStatus(ctx context.Context, hub *v1alpha2.Hub, opts v1.UpdateOptions) (*v1alpha2.Hub, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(hubsResource, "status", c.ns, hub), &v1alpha2.Hub{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.Hub), err
}

// Delete takes name of the hub and deletes it. Returns an error if one occurs.
func (c *FakeHubs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(hubsResource, c.ns, name, opts), &v1alpha2.Hub{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeHubs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(hubsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha2.HubList{})
	return err
}

// Patch applies the patch and returns the patched hub.
func (c *FakeHubs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha2.Hub, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(hubsResource, c.ns, name, pt, data, subresources...), &v1alpha2.Hub{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.Hub), err
}
