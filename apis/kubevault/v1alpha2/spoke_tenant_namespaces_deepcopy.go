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

// Hand-written deepcopy for SpokeTenantNamespaces. Kept out of zz_generated.deepcopy.go
// so it survives a `make gen` (which will regenerate the canonical copies); the generator
// produces identical functions, so removing this file after codegen is a no-op.

package v1alpha2

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto copies the receiver into out.
func (in *SpokeTenantNamespaces) DeepCopyInto(out *SpokeTenantNamespaces) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy returns a deep copy of the receiver.
func (in *SpokeTenantNamespaces) DeepCopy() *SpokeTenantNamespaces {
	if in == nil {
		return nil
	}
	out := new(SpokeTenantNamespaces)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject returns a runtime.Object deep copy of the receiver.
func (in *SpokeTenantNamespaces) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto copies the receiver into out.
func (in *SpokeTenantNamespacesSpec) DeepCopyInto(out *SpokeTenantNamespacesSpec) {
	*out = *in
	out.HubVaultRef = in.HubVaultRef
}

// DeepCopy returns a deep copy of the receiver.
func (in *SpokeTenantNamespacesSpec) DeepCopy() *SpokeTenantNamespacesSpec {
	if in == nil {
		return nil
	}
	out := new(SpokeTenantNamespacesSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto copies the receiver into out.
func (in *SpokeTenantNamespacesStatus) DeepCopyInto(out *SpokeTenantNamespacesStatus) {
	*out = *in
	if in.OrgIDs != nil {
		out.OrgIDs = make([]string, len(in.OrgIDs))
		copy(out.OrgIDs, in.OrgIDs)
	}
}

// DeepCopy returns a deep copy of the receiver.
func (in *SpokeTenantNamespacesStatus) DeepCopy() *SpokeTenantNamespacesStatus {
	if in == nil {
		return nil
	}
	out := new(SpokeTenantNamespacesStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto copies the receiver into out.
func (in *SpokeTenantNamespacesList) DeepCopyInto(out *SpokeTenantNamespacesList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		out.Items = make([]SpokeTenantNamespaces, len(in.Items))
		for i := range in.Items {
			in.Items[i].DeepCopyInto(&out.Items[i])
		}
	}
}

// DeepCopy returns a deep copy of the receiver.
func (in *SpokeTenantNamespacesList) DeepCopy() *SpokeTenantNamespacesList {
	if in == nil {
		return nil
	}
	out := new(SpokeTenantNamespacesList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject returns a runtime.Object deep copy of the receiver.
func (in *SpokeTenantNamespacesList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
