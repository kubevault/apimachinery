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

package v1alpha1

import (
	ofst "kmodules.xyz/offshoot-api/api/v1"
)

// +kubebuilder:validation:Enum=Provisioning;DataRestoring;Ready;Critical;NotReady;Halted
type DatabasePhase string

const (
	// used for Databases that are currently provisioning
	DatabasePhaseProvisioning DatabasePhase = "Provisioning"
	// used for Databases for which data is currently restoring
	DatabasePhaseDataRestoring DatabasePhase = "DataRestoring"
	// used for Databases that are currently ReplicaReady, AcceptingConnection and Ready
	DatabasePhaseReady DatabasePhase = "Ready"
	// used for Databases that can connect, ReplicaReady == false || Ready == false (eg, ES yellow)
	DatabasePhaseCritical DatabasePhase = "Critical"
	// used for Databases that can't connect
	DatabasePhaseNotReady DatabasePhase = "NotReady"
	// used for Databases that are halted
	DatabasePhaseHalted DatabasePhase = "Halted"
)

// +kubebuilder:validation:Enum=Halt;Delete;WipeOut;DoNotTerminate
type TerminationPolicy string

const (
	// Deletes database pods, service but leave the PVCs and stash backup data intact.
	TerminationPolicyHalt TerminationPolicy = "Halt"
	// Deletes database pods, service, pvcs but leave the stash backup data intact.
	TerminationPolicyDelete TerminationPolicy = "Delete"
	// Deletes database pods, service, pvcs and stash backup data.
	TerminationPolicyWipeOut TerminationPolicy = "WipeOut"
	// Rejects attempt to delete database using ValidationWebhook.
	TerminationPolicyDoNotTerminate TerminationPolicy = "DoNotTerminate"
)

// +kubebuilder:validation:Enum=primary;standby;stats
type ServiceAlias string

const (
	PrimaryServiceAlias ServiceAlias = "primary"
	StandbyServiceAlias ServiceAlias = "standby"
	StatsServiceAlias   ServiceAlias = "stats"
)

type NamedServiceTemplateSpec struct {
	// Alias represents the identifier of the service.
	Alias ServiceAlias `json:"alias" protobuf:"bytes,1,opt,name=alias"`

	// ServiceTemplate is an optional configuration for a service used to expose database
	// +optional
	ofst.ServiceTemplateSpec `json:",inline,omitempty" protobuf:"bytes,2,opt,name=serviceTemplateSpec"`
}
