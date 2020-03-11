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

package v1alpha1

import (
	"time"

	v1alpha1 "github.com/google/knative-gcp/pkg/apis/security/v1alpha1"
	scheme "github.com/google/knative-gcp/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// EventPoliciesGetter has a method to return a EventPolicyInterface.
// A group's client should implement this interface.
type EventPoliciesGetter interface {
	EventPolicies(namespace string) EventPolicyInterface
}

// EventPolicyInterface has methods to work with EventPolicy resources.
type EventPolicyInterface interface {
	Create(*v1alpha1.EventPolicy) (*v1alpha1.EventPolicy, error)
	Update(*v1alpha1.EventPolicy) (*v1alpha1.EventPolicy, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.EventPolicy, error)
	List(opts v1.ListOptions) (*v1alpha1.EventPolicyList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.EventPolicy, err error)
	EventPolicyExpansion
}

// eventPolicies implements EventPolicyInterface
type eventPolicies struct {
	client rest.Interface
	ns     string
}

// newEventPolicies returns a EventPolicies
func newEventPolicies(c *SecurityV1alpha1Client, namespace string) *eventPolicies {
	return &eventPolicies{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the eventPolicy, and returns the corresponding eventPolicy object, and an error if there is any.
func (c *eventPolicies) Get(name string, options v1.GetOptions) (result *v1alpha1.EventPolicy, err error) {
	result = &v1alpha1.EventPolicy{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("eventpolicies").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of EventPolicies that match those selectors.
func (c *eventPolicies) List(opts v1.ListOptions) (result *v1alpha1.EventPolicyList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.EventPolicyList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("eventpolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested eventPolicies.
func (c *eventPolicies) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("eventpolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a eventPolicy and creates it.  Returns the server's representation of the eventPolicy, and an error, if there is any.
func (c *eventPolicies) Create(eventPolicy *v1alpha1.EventPolicy) (result *v1alpha1.EventPolicy, err error) {
	result = &v1alpha1.EventPolicy{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("eventpolicies").
		Body(eventPolicy).
		Do().
		Into(result)
	return
}

// Update takes the representation of a eventPolicy and updates it. Returns the server's representation of the eventPolicy, and an error, if there is any.
func (c *eventPolicies) Update(eventPolicy *v1alpha1.EventPolicy) (result *v1alpha1.EventPolicy, err error) {
	result = &v1alpha1.EventPolicy{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("eventpolicies").
		Name(eventPolicy.Name).
		Body(eventPolicy).
		Do().
		Into(result)
	return
}

// Delete takes name of the eventPolicy and deletes it. Returns an error if one occurs.
func (c *eventPolicies) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("eventpolicies").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *eventPolicies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("eventpolicies").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched eventPolicy.
func (c *eventPolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.EventPolicy, err error) {
	result = &v1alpha1.EventPolicy{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("eventpolicies").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
