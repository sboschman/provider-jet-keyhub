/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type KeyHubGroupOnSystemObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	ShortNameInSystem *string `json:"shortNameInSystem,omitempty" tf:"short_name_in_system,omitempty"`
}

type KeyHubGroupOnSystemParameters struct {

	// Display name of the group on provisioned system. (Only on systems that support a display name)
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// The name in the system, value normally to CN of the DN (of the provisioned system). For example: `cn=umbrella,ou=group,dc=example,dc=com`
	// +kubebuilder:validation:Required
	NameInSystem *string `json:"nameInSystem" tf:"name_in_system,omitempty"`

	// The UUID of the group that will become owner of the grouponsystem
	// +kubebuilder:validation:Required
	Owner *string `json:"owner" tf:"owner,omitempty"`

	// Define the provisioning group for the grouponsystem, can be set multiple times. If omitted the owner group will be the provisioning group
	// +kubebuilder:validation:Optional
	Provgroup []ProvgroupParameters `json:"provgroup,omitempty" tf:"provgroup,omitempty"`

	// The UUID of the provisioned system to create the group on
	// +kubebuilder:validation:Required
	System *string `json:"system" tf:"system,omitempty"`

	// Type of the resulting group in the provisioned system, for example: POSIX_GROUP for ldap.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ProvgroupObservation struct {
}

type ProvgroupParameters struct {

	// The UUID of the group that will become a provisioning group for the grouponsystem
	// +kubebuilder:validation:Required
	Group *string `json:"group" tf:"group,omitempty"`

	// The security level. Possible values: `HIGH` (default), `MEDIUM`, `LOW`
	// +kubebuilder:validation:Optional
	Securitylevel *string `json:"securitylevel,omitempty" tf:"securitylevel,omitempty"`

	// If set to true the group on system will be static provisioned
	// +kubebuilder:validation:Optional
	Static *bool `json:"static,omitempty" tf:"static,omitempty"`
}

// KeyHubGroupOnSystemSpec defines the desired state of KeyHubGroupOnSystem
type KeyHubGroupOnSystemSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     KeyHubGroupOnSystemParameters `json:"forProvider"`
}

// KeyHubGroupOnSystemStatus defines the observed state of KeyHubGroupOnSystem.
type KeyHubGroupOnSystemStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        KeyHubGroupOnSystemObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// KeyHubGroupOnSystem is the Schema for the KeyHubGroupOnSystems API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,keyhubjet}
type KeyHubGroupOnSystem struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KeyHubGroupOnSystemSpec   `json:"spec"`
	Status            KeyHubGroupOnSystemStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// KeyHubGroupOnSystemList contains a list of KeyHubGroupOnSystems
type KeyHubGroupOnSystemList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KeyHubGroupOnSystem `json:"items"`
}

// Repository type metadata.
var (
	KeyHubGroupOnSystem_Kind             = "KeyHubGroupOnSystem"
	KeyHubGroupOnSystem_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: KeyHubGroupOnSystem_Kind}.String()
	KeyHubGroupOnSystem_KindAPIVersion   = KeyHubGroupOnSystem_Kind + "." + CRDGroupVersion.String()
	KeyHubGroupOnSystem_GroupVersionKind = CRDGroupVersion.WithKind(KeyHubGroupOnSystem_Kind)
)

func init() {
	SchemeBuilder.Register(&KeyHubGroupOnSystem{}, &KeyHubGroupOnSystemList{})
}