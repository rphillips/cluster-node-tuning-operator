// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	tunedv1 "github.com/openshift/cluster-node-tuning-operator/pkg/apis/tuned/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeProfiles implements ProfileInterface
type FakeProfiles struct {
	Fake *FakeTunedV1
	ns   string
}

var profilesResource = schema.GroupVersionResource{Group: "tuned.openshift.io", Version: "v1", Resource: "profiles"}

var profilesKind = schema.GroupVersionKind{Group: "tuned.openshift.io", Version: "v1", Kind: "Profile"}

// Get takes name of the profile, and returns the corresponding profile object, and an error if there is any.
func (c *FakeProfiles) Get(ctx context.Context, name string, options v1.GetOptions) (result *tunedv1.Profile, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(profilesResource, c.ns, name), &tunedv1.Profile{})

	if obj == nil {
		return nil, err
	}
	return obj.(*tunedv1.Profile), err
}

// List takes label and field selectors, and returns the list of Profiles that match those selectors.
func (c *FakeProfiles) List(ctx context.Context, opts v1.ListOptions) (result *tunedv1.ProfileList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(profilesResource, profilesKind, c.ns, opts), &tunedv1.ProfileList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &tunedv1.ProfileList{ListMeta: obj.(*tunedv1.ProfileList).ListMeta}
	for _, item := range obj.(*tunedv1.ProfileList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested profiles.
func (c *FakeProfiles) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(profilesResource, c.ns, opts))

}

// Create takes the representation of a profile and creates it.  Returns the server's representation of the profile, and an error, if there is any.
func (c *FakeProfiles) Create(ctx context.Context, profile *tunedv1.Profile, opts v1.CreateOptions) (result *tunedv1.Profile, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(profilesResource, c.ns, profile), &tunedv1.Profile{})

	if obj == nil {
		return nil, err
	}
	return obj.(*tunedv1.Profile), err
}

// Update takes the representation of a profile and updates it. Returns the server's representation of the profile, and an error, if there is any.
func (c *FakeProfiles) Update(ctx context.Context, profile *tunedv1.Profile, opts v1.UpdateOptions) (result *tunedv1.Profile, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(profilesResource, c.ns, profile), &tunedv1.Profile{})

	if obj == nil {
		return nil, err
	}
	return obj.(*tunedv1.Profile), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeProfiles) UpdateStatus(ctx context.Context, profile *tunedv1.Profile, opts v1.UpdateOptions) (*tunedv1.Profile, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(profilesResource, "status", c.ns, profile), &tunedv1.Profile{})

	if obj == nil {
		return nil, err
	}
	return obj.(*tunedv1.Profile), err
}

// Delete takes name of the profile and deletes it. Returns an error if one occurs.
func (c *FakeProfiles) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(profilesResource, c.ns, name), &tunedv1.Profile{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeProfiles) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(profilesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &tunedv1.ProfileList{})
	return err
}

// Patch applies the patch and returns the patched profile.
func (c *FakeProfiles) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *tunedv1.Profile, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(profilesResource, c.ns, name, pt, data, subresources...), &tunedv1.Profile{})

	if obj == nil {
		return nil, err
	}
	return obj.(*tunedv1.Profile), err
}
