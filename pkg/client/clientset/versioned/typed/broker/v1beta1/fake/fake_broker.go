/*
Copyright 2020 Google LLC

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
	v1beta1 "github.com/google/knative-gcp/pkg/apis/broker/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeBrokers implements BrokerInterface
type FakeBrokers struct {
	Fake *FakeEventingV1beta1
	ns   string
}

var brokersResource = schema.GroupVersionResource{Group: "eventing.knative.dev", Version: "v1beta1", Resource: "brokers"}

var brokersKind = schema.GroupVersionKind{Group: "eventing.knative.dev", Version: "v1beta1", Kind: "Broker"}

// Get takes name of the broker, and returns the corresponding broker object, and an error if there is any.
func (c *FakeBrokers) Get(name string, options v1.GetOptions) (result *v1beta1.Broker, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(brokersResource, c.ns, name), &v1beta1.Broker{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Broker), err
}

// List takes label and field selectors, and returns the list of Brokers that match those selectors.
func (c *FakeBrokers) List(opts v1.ListOptions) (result *v1beta1.BrokerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(brokersResource, brokersKind, c.ns, opts), &v1beta1.BrokerList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.BrokerList{ListMeta: obj.(*v1beta1.BrokerList).ListMeta}
	for _, item := range obj.(*v1beta1.BrokerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested brokers.
func (c *FakeBrokers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(brokersResource, c.ns, opts))

}

// Create takes the representation of a broker and creates it.  Returns the server's representation of the broker, and an error, if there is any.
func (c *FakeBrokers) Create(broker *v1beta1.Broker) (result *v1beta1.Broker, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(brokersResource, c.ns, broker), &v1beta1.Broker{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Broker), err
}

// Update takes the representation of a broker and updates it. Returns the server's representation of the broker, and an error, if there is any.
func (c *FakeBrokers) Update(broker *v1beta1.Broker) (result *v1beta1.Broker, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(brokersResource, c.ns, broker), &v1beta1.Broker{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Broker), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeBrokers) UpdateStatus(broker *v1beta1.Broker) (*v1beta1.Broker, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(brokersResource, "status", c.ns, broker), &v1beta1.Broker{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Broker), err
}

// Delete takes name of the broker and deletes it. Returns an error if one occurs.
func (c *FakeBrokers) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(brokersResource, c.ns, name), &v1beta1.Broker{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeBrokers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(brokersResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1beta1.BrokerList{})
	return err
}

// Patch applies the patch and returns the patched broker.
func (c *FakeBrokers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.Broker, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(brokersResource, c.ns, name, pt, data, subresources...), &v1beta1.Broker{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Broker), err
}
