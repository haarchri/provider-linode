/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type DeviceObservation struct {

	// When the Firewall Device was last created.
	// When this Firewall Device was created.
	Created *string `json:"created,omitempty" tf:"created,omitempty"`

	// The unique ID of the entity to attach.
	// The ID of the entity to create a Firewall device for.
	EntityID *float64 `json:"entityId,omitempty" tf:"entity_id,omitempty"`

	// The type of the entity to attach. (default: linode)
	// The type of the entity to create a Firewall device for.
	EntityType *string `json:"entityType,omitempty" tf:"entity_type,omitempty"`

	// The unique ID of the target Firewall.
	// The ID of the Firewall to access.
	FirewallID *float64 `json:"firewallId,omitempty" tf:"firewall_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// When the Firewall Device was last updated.
	// When this Firewall Device was updated.
	Updated *string `json:"updated,omitempty" tf:"updated,omitempty"`
}

type DeviceParameters struct {

	// The unique ID of the entity to attach.
	// The ID of the entity to create a Firewall device for.
	// +crossplane:generate:reference:type=github.com/linode/provider-linode/apis/instance/v1alpha1.Instance
	// +kubebuilder:validation:Optional
	EntityID *float64 `json:"entityId,omitempty" tf:"entity_id,omitempty"`

	// Reference to a Instance in instance to populate entityId.
	// +kubebuilder:validation:Optional
	EntityIDRef *v1.Reference `json:"entityIdRef,omitempty" tf:"-"`

	// Selector for a Instance in instance to populate entityId.
	// +kubebuilder:validation:Optional
	EntityIDSelector *v1.Selector `json:"entityIdSelector,omitempty" tf:"-"`

	// The type of the entity to attach. (default: linode)
	// The type of the entity to create a Firewall device for.
	// +kubebuilder:validation:Optional
	EntityType *string `json:"entityType,omitempty" tf:"entity_type,omitempty"`

	// The unique ID of the target Firewall.
	// The ID of the Firewall to access.
	// +crossplane:generate:reference:type=github.com/linode/provider-linode/apis/firewall/v1alpha1.Device
	// +kubebuilder:validation:Optional
	FirewallID *float64 `json:"firewallId,omitempty" tf:"firewall_id,omitempty"`

	// Reference to a Device in firewall to populate firewallId.
	// +kubebuilder:validation:Optional
	FirewallIDRef *v1.Reference `json:"firewallIdRef,omitempty" tf:"-"`

	// Selector for a Device in firewall to populate firewallId.
	// +kubebuilder:validation:Optional
	FirewallIDSelector *v1.Selector `json:"firewallIdSelector,omitempty" tf:"-"`
}

// DeviceSpec defines the desired state of Device
type DeviceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DeviceParameters `json:"forProvider"`
}

// DeviceStatus defines the observed state of Device.
type DeviceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DeviceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Device is the Schema for the Devices API. Manages a Linode Firewall Device.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,linode}
type Device struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DeviceSpec   `json:"spec"`
	Status            DeviceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DeviceList contains a list of Devices
type DeviceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Device `json:"items"`
}

// Repository type metadata.
var (
	Device_Kind             = "Device"
	Device_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Device_Kind}.String()
	Device_KindAPIVersion   = Device_Kind + "." + CRDGroupVersion.String()
	Device_GroupVersionKind = CRDGroupVersion.WithKind(Device_Kind)
)

func init() {
	SchemeBuilder.Register(&Device{}, &DeviceList{})
}
