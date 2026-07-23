/*
Copyright AppsCode Inc. and Contributors

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

// Hand-written deepcopy for NamespaceSlice. Kept out of zz_generated.deepcopy.go
// so it survives a `make gen` (which will regenerate the canonical copies); the generator
// produces identical functions, so removing this file after codegen is a no-op.

package v1alpha2

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto copies the receiver into out.
func (in *NamespaceSlice) DeepCopyInto(out *NamespaceSlice) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy returns a deep copy of the receiver.
func (in *NamespaceSlice) DeepCopy() *NamespaceSlice {
	if in == nil {
		return nil
	}
	out := new(NamespaceSlice)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject returns a runtime.Object deep copy of the receiver.
func (in *NamespaceSlice) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto copies the receiver into out.
func (in *NamespaceSliceSpec) DeepCopyInto(out *NamespaceSliceSpec) {
	*out = *in
	out.HubVaultRef = in.HubVaultRef
	if in.Namespaces != nil {
		out.Namespaces = make([]NamespaceSliceEntry, len(in.Namespaces))
		for i := range in.Namespaces {
			in.Namespaces[i].DeepCopyInto(&out.Namespaces[i])
		}
	}
}

// DeepCopy returns a deep copy of the receiver.
func (in *NamespaceSliceSpec) DeepCopy() *NamespaceSliceSpec {
	if in == nil {
		return nil
	}
	out := new(NamespaceSliceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto copies the receiver into out.
func (in *NamespaceSliceEntry) DeepCopyInto(out *NamespaceSliceEntry) {
	*out = *in
	in.Conditions.DeepCopyInto(&out.Conditions)
}

// DeepCopy returns a deep copy of the receiver.
func (in *NamespaceSliceEntry) DeepCopy() *NamespaceSliceEntry {
	if in == nil {
		return nil
	}
	out := new(NamespaceSliceEntry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto copies the receiver into out.
func (in *NamespaceSliceConditions) DeepCopyInto(out *NamespaceSliceConditions) {
	*out = *in
	if in.Ready != nil {
		out.Ready = new(bool)
		*out.Ready = *in.Ready
	}
}

// DeepCopy returns a deep copy of the receiver.
func (in *NamespaceSliceConditions) DeepCopy() *NamespaceSliceConditions {
	if in == nil {
		return nil
	}
	out := new(NamespaceSliceConditions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto copies the receiver into out.
func (in *NamespaceSliceStatus) DeepCopyInto(out *NamespaceSliceStatus) {
	*out = *in
}

// DeepCopy returns a deep copy of the receiver.
func (in *NamespaceSliceStatus) DeepCopy() *NamespaceSliceStatus {
	if in == nil {
		return nil
	}
	out := new(NamespaceSliceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto copies the receiver into out.
func (in *NamespaceSliceList) DeepCopyInto(out *NamespaceSliceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		out.Items = make([]NamespaceSlice, len(in.Items))
		for i := range in.Items {
			in.Items[i].DeepCopyInto(&out.Items[i])
		}
	}
}

// DeepCopy returns a deep copy of the receiver.
func (in *NamespaceSliceList) DeepCopy() *NamespaceSliceList {
	if in == nil {
		return nil
	}
	out := new(NamespaceSliceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject returns a runtime.Object deep copy of the receiver.
func (in *NamespaceSliceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
