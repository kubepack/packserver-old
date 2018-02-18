/*
Copyright 2018 The Kubepack Authors.

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
package internalversion

import (
	apps "github.com/kubepack/packserver/apis/apps"
	scheme "github.com/kubepack/packserver/client/clientset/internalversion/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ManifestsGetter has a method to return a ManifestInterface.
// A group's client should implement this interface.
type ManifestsGetter interface {
	Manifests(namespace string) ManifestInterface
}

// ManifestInterface has methods to work with Manifest resources.
type ManifestInterface interface {
	Create(*apps.Manifest) (*apps.Manifest, error)
	Update(*apps.Manifest) (*apps.Manifest, error)
	UpdateStatus(*apps.Manifest) (*apps.Manifest, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*apps.Manifest, error)
	List(opts v1.ListOptions) (*apps.ManifestList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *apps.Manifest, err error)
	ManifestExpansion
}

// manifests implements ManifestInterface
type manifests struct {
	client rest.Interface
	ns     string
}

// newManifests returns a Manifests
func newManifests(c *AppsClient, namespace string) *manifests {
	return &manifests{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the manifest, and returns the corresponding manifest object, and an error if there is any.
func (c *manifests) Get(name string, options v1.GetOptions) (result *apps.Manifest, err error) {
	result = &apps.Manifest{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("manifests").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Manifests that match those selectors.
func (c *manifests) List(opts v1.ListOptions) (result *apps.ManifestList, err error) {
	result = &apps.ManifestList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("manifests").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested manifests.
func (c *manifests) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("manifests").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a manifest and creates it.  Returns the server's representation of the manifest, and an error, if there is any.
func (c *manifests) Create(manifest *apps.Manifest) (result *apps.Manifest, err error) {
	result = &apps.Manifest{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("manifests").
		Body(manifest).
		Do().
		Into(result)
	return
}

// Update takes the representation of a manifest and updates it. Returns the server's representation of the manifest, and an error, if there is any.
func (c *manifests) Update(manifest *apps.Manifest) (result *apps.Manifest, err error) {
	result = &apps.Manifest{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("manifests").
		Name(manifest.Name).
		Body(manifest).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *manifests) UpdateStatus(manifest *apps.Manifest) (result *apps.Manifest, err error) {
	result = &apps.Manifest{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("manifests").
		Name(manifest.Name).
		SubResource("status").
		Body(manifest).
		Do().
		Into(result)
	return
}

// Delete takes name of the manifest and deletes it. Returns an error if one occurs.
func (c *manifests) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("manifests").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *manifests) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("manifests").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched manifest.
func (c *manifests) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *apps.Manifest, err error) {
	result = &apps.Manifest{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("manifests").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
