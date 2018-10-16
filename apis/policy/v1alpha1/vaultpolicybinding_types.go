package v1alpha1

import (
	"github.com/appscode/go/encoding/json/types"
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	ResourceKindVaultPolicyBinding = "VaultPolicyBinding"
	ResourceVaultPolicyBinding     = "vaultpolicybinding"
	ResourceVaultPolicyBindings    = "vaultpolicybindings"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type VaultPolicyBinding struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VaultPolicyBindingSpec   `json:"spec,omitempty"`
	Status            VaultPolicyBindingStatus `json:"status,omitempty"`
}

// links: https://www.vaultproject.io/api/auth/kubernetes/index.html#parameters-1
type VaultPolicyBindingSpec struct {
	// Specifies the path where kubernetes auth is enabled
	// default : kubernetes
	AuthPath string `json:"authPath,omitempty"`

	// Specifies the names of the VaultPolicy
	Policies []string `json:"policies"`

	// Specifies the names of the service account to bind with policy
	ServiceAccountNames []string `json:"serviceAccountNames"`

	// Specifies the namespaces of the service account
	ServiceAccountNamespaces []string `json:"serviceAccountNamespaces"`

	//Specifies the TTL period of tokens issued using this role in seconds.
	TTL string `json:"TTL,omiempty"`

	//Specifies the maximum allowed lifetime of tokens issued in seconds using this role.
	MaxTTL string `json:"maxTTL,omitempty"`

	// If set, indicates that the token generated using this role should never expire.
	// The token should be renewed within the duration specified by this value.
	// At each renewal, the token's TTL will be set to the value of this parameter.
	Period string `json:"period,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type VaultPolicyBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VaultPolicyBinding `json:"items,omitempty"`
}

// ServiceAccountReference contains name and namespace of the service account
type ServiceAccountReference struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}

type PolicyBindingStatus string

const (
	PolicyBindingSuccess PolicyBindingStatus = "Success"
	PolicyBindingFailed  PolicyBindingStatus = "Failed"
)

type VaultPolicyBindingStatus struct {
	// observedGeneration is the most recent generation observed for this resource. It corresponds to the
	// resource's generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration *types.IntHash `json:"observedGeneration,omitempty"`

	// Status indicates whether successfully bind the policy to service account in vault or not or in progress
	Status PolicyBindingStatus `json:"status,omitempty"`

	// Represents the latest available observations of a VaultPolicyBinding.
	Conditions []PolicyBindingCondition `json:"conditions,omitempty"`
}

type PolicyBindingConditionType string

// These are valid conditions of a VaultPolicyBinding.
const (
	PolicyBindingConditionFailure PolicyBindingConditionType = "Failure"
)

// PolicyBindingCondition describes the state of a VaultPolicyBinding at a certain point.
type PolicyBindingCondition struct {
	// Type of PolicyBindingCondition condition.
	Type PolicyBindingConditionType `json:"type,omitempty"`

	// Status of the condition, one of True, False, Unknown.
	Status core.ConditionStatus `json:"status,omitempty"`

	// The reason for the condition's.
	Reason string `json:"reason,omitempty"`

	// A human readable message indicating details about the transition.
	Message string `json:"message,omitempty"`
}
