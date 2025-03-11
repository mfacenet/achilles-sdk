//go:build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"

	"github.com/reddit/achilles-sdk-api/api"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TestClaim) DeepCopyInto(out *TestClaim) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TestClaim.
func (in *TestClaim) DeepCopy() *TestClaim {
	if in == nil {
		return nil
	}
	out := new(TestClaim)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TestClaim) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TestClaimList) DeepCopyInto(out *TestClaimList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]*TestClaim, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(TestClaim)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TestClaimList.
func (in *TestClaimList) DeepCopy() *TestClaimList {
	if in == nil {
		return nil
	}
	out := new(TestClaimList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TestClaimList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TestClaimSpec) DeepCopyInto(out *TestClaimSpec) {
	*out = *in
	if in.ClaimedRef != nil {
		in, out := &in.ClaimedRef, &out.ClaimedRef
		*out = new(api.TypedObjectRef)
		**out = **in
	}
	if in.ConfigMapName != nil {
		in, out := &in.ConfigMapName, &out.ConfigMapName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TestClaimSpec.
func (in *TestClaimSpec) DeepCopy() *TestClaimSpec {
	if in == nil {
		return nil
	}
	out := new(TestClaimSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TestClaimStatus) DeepCopyInto(out *TestClaimStatus) {
	*out = *in
	in.ConditionedStatus.DeepCopyInto(&out.ConditionedStatus)
	if in.ResourceRefs != nil {
		in, out := &in.ResourceRefs, &out.ResourceRefs
		*out = make([]api.TypedObjectRef, len(*in))
		copy(*out, *in)
	}
	if in.ConfigMapName != nil {
		in, out := &in.ConfigMapName, &out.ConfigMapName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TestClaimStatus.
func (in *TestClaimStatus) DeepCopy() *TestClaimStatus {
	if in == nil {
		return nil
	}
	out := new(TestClaimStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TestClaimed) DeepCopyInto(out *TestClaimed) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TestClaimed.
func (in *TestClaimed) DeepCopy() *TestClaimed {
	if in == nil {
		return nil
	}
	out := new(TestClaimed)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TestClaimed) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TestClaimedList) DeepCopyInto(out *TestClaimedList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]*TestClaimed, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(TestClaimed)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TestClaimedList.
func (in *TestClaimedList) DeepCopy() *TestClaimedList {
	if in == nil {
		return nil
	}
	out := new(TestClaimedList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TestClaimedList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TestClaimedSpec) DeepCopyInto(out *TestClaimedSpec) {
	*out = *in
	if in.ClaimRef != nil {
		in, out := &in.ClaimRef, &out.ClaimRef
		*out = new(api.TypedObjectRef)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TestClaimedSpec.
func (in *TestClaimedSpec) DeepCopy() *TestClaimedSpec {
	if in == nil {
		return nil
	}
	out := new(TestClaimedSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TestClaimedStatus) DeepCopyInto(out *TestClaimedStatus) {
	*out = *in
	in.ConditionedStatus.DeepCopyInto(&out.ConditionedStatus)
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]api.TypedObjectRef, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TestClaimedStatus.
func (in *TestClaimedStatus) DeepCopy() *TestClaimedStatus {
	if in == nil {
		return nil
	}
	out := new(TestClaimedStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TestResourceWithoutSubresource) DeepCopyInto(out *TestResourceWithoutSubresource) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TestResourceWithoutSubresource.
func (in *TestResourceWithoutSubresource) DeepCopy() *TestResourceWithoutSubresource {
	if in == nil {
		return nil
	}
	out := new(TestResourceWithoutSubresource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TestResourceWithoutSubresource) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TestResourceWithoutSubresourceList) DeepCopyInto(out *TestResourceWithoutSubresourceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]*TestResourceWithoutSubresource, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(TestResourceWithoutSubresource)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TestResourceWithoutSubresourceList.
func (in *TestResourceWithoutSubresourceList) DeepCopy() *TestResourceWithoutSubresourceList {
	if in == nil {
		return nil
	}
	out := new(TestResourceWithoutSubresourceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TestResourceWithoutSubresourceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TestResourceWithoutSubresourceSpec) DeepCopyInto(out *TestResourceWithoutSubresourceSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TestResourceWithoutSubresourceSpec.
func (in *TestResourceWithoutSubresourceSpec) DeepCopy() *TestResourceWithoutSubresourceSpec {
	if in == nil {
		return nil
	}
	out := new(TestResourceWithoutSubresourceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TestResourceWithoutSubresourceStatus) DeepCopyInto(out *TestResourceWithoutSubresourceStatus) {
	*out = *in
	in.ConditionedStatus.DeepCopyInto(&out.ConditionedStatus)
	if in.ResourceRefs != nil {
		in, out := &in.ResourceRefs, &out.ResourceRefs
		*out = make([]api.TypedObjectRef, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TestResourceWithoutSubresourceStatus.
func (in *TestResourceWithoutSubresourceStatus) DeepCopy() *TestResourceWithoutSubresourceStatus {
	if in == nil {
		return nil
	}
	out := new(TestResourceWithoutSubresourceStatus)
	in.DeepCopyInto(out)
	return out
}
