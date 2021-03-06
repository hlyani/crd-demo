/*
Copyright The Kubernetes Authors.

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
	crddemov1 "crddemo/pkg/apis/crddemo/v1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeMydemos implements MydemoInterface
type FakeMydemos struct {
	Fake *FakeCrddemoV1
	ns   string
}

var mydemosResource = schema.GroupVersionResource{Group: "crddemo.k8s.io", Version: "v1", Resource: "mydemos"}

var mydemosKind = schema.GroupVersionKind{Group: "crddemo.k8s.io", Version: "v1", Kind: "Mydemo"}

// Get takes name of the mydemo, and returns the corresponding mydemo object, and an error if there is any.
func (c *FakeMydemos) Get(ctx context.Context, name string, options v1.GetOptions) (result *crddemov1.Mydemo, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(mydemosResource, c.ns, name), &crddemov1.Mydemo{})

	if obj == nil {
		return nil, err
	}
	return obj.(*crddemov1.Mydemo), err
}

// List takes label and field selectors, and returns the list of Mydemos that match those selectors.
func (c *FakeMydemos) List(ctx context.Context, opts v1.ListOptions) (result *crddemov1.MydemoList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(mydemosResource, mydemosKind, c.ns, opts), &crddemov1.MydemoList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &crddemov1.MydemoList{ListMeta: obj.(*crddemov1.MydemoList).ListMeta}
	for _, item := range obj.(*crddemov1.MydemoList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested mydemos.
func (c *FakeMydemos) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(mydemosResource, c.ns, opts))

}

// Create takes the representation of a mydemo and creates it.  Returns the server's representation of the mydemo, and an error, if there is any.
func (c *FakeMydemos) Create(ctx context.Context, mydemo *crddemov1.Mydemo, opts v1.CreateOptions) (result *crddemov1.Mydemo, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(mydemosResource, c.ns, mydemo), &crddemov1.Mydemo{})

	if obj == nil {
		return nil, err
	}
	return obj.(*crddemov1.Mydemo), err
}

// Update takes the representation of a mydemo and updates it. Returns the server's representation of the mydemo, and an error, if there is any.
func (c *FakeMydemos) Update(ctx context.Context, mydemo *crddemov1.Mydemo, opts v1.UpdateOptions) (result *crddemov1.Mydemo, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(mydemosResource, c.ns, mydemo), &crddemov1.Mydemo{})

	if obj == nil {
		return nil, err
	}
	return obj.(*crddemov1.Mydemo), err
}

// Delete takes name of the mydemo and deletes it. Returns an error if one occurs.
func (c *FakeMydemos) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(mydemosResource, c.ns, name), &crddemov1.Mydemo{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMydemos) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(mydemosResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &crddemov1.MydemoList{})
	return err
}

// Patch applies the patch and returns the patched mydemo.
func (c *FakeMydemos) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *crddemov1.Mydemo, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(mydemosResource, c.ns, name, pt, data, subresources...), &crddemov1.Mydemo{})

	if obj == nil {
		return nil, err
	}
	return obj.(*crddemov1.Mydemo), err
}
