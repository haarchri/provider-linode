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

type DiskObservation_2 struct {

	// A list of public SSH keys that will be automatically appended to the root user’s ~/.ssh/authorized_keys file when deploying from an Image.
	// A list of public SSH keys that will be automatically appended to the root user’s ~/.ssh/authorized_keys file when deploying from an Image.
	AuthorizedKeys []*string `json:"authorizedKeys,omitempty" tf:"authorized_keys,omitempty"`

	// A list of usernames. If the usernames have associated SSH keys, the keys will be appended to the
	// A list of usernames. If the usernames have associated SSH keys, the keys will be appended to the root users ~/.ssh/authorized_keys file automatically when deploying from an Image.
	AuthorizedUsers []*string `json:"authorizedUsers,omitempty" tf:"authorized_users,omitempty"`

	// When this disk was created.
	// When this disk was created.
	Created *string `json:"created,omitempty" tf:"created,omitempty"`

	// The filesystem of this disk. (raw, swap, ext3, ext4, initrd)
	// The filesystem of this disk.
	Filesystem *string `json:"filesystem,omitempty" tf:"filesystem,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// An Image ID to deploy the Linode Disk from.
	// An Image ID to deploy the Linode Disk from.
	Image *string `json:"image,omitempty" tf:"image,omitempty"`

	// The Disk's label for display purposes only.
	// The Disk’s label is for display purposes only.
	Label *string `json:"label,omitempty" tf:"label,omitempty"`

	// The ID of the Linode to create this Disk under.
	// The ID of the Linode to assign this disk to.
	LinodeID *float64 `json:"linodeId,omitempty" tf:"linode_id,omitempty"`

	// The size of the Disk in MB. NOTE: Resizing a disk will trigger a Linode reboot.
	// The size of the Disk in MB.
	Size *float64 `json:"size,omitempty" tf:"size,omitempty"`

	// A StackScript ID that will cause the referenced StackScript to be run during deployment of this Disk.
	// A StackScript ID that will cause the referenced StackScript to be run during deployment of this Linode.
	StackscriptID *float64 `json:"stackscriptId,omitempty" tf:"stackscript_id,omitempty"`

	// A brief description of this Disk's current state.
	// A brief description of this Disk's current state.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// When this disk was last updated.
	// When this disk was last updated.
	Updated *string `json:"updated,omitempty" tf:"updated,omitempty"`
}

type DiskParameters_2 struct {

	// A list of public SSH keys that will be automatically appended to the root user’s ~/.ssh/authorized_keys file when deploying from an Image.
	// A list of public SSH keys that will be automatically appended to the root user’s ~/.ssh/authorized_keys file when deploying from an Image.
	// +kubebuilder:validation:Optional
	AuthorizedKeys []*string `json:"authorizedKeys,omitempty" tf:"authorized_keys,omitempty"`

	// A list of usernames. If the usernames have associated SSH keys, the keys will be appended to the
	// A list of usernames. If the usernames have associated SSH keys, the keys will be appended to the root users ~/.ssh/authorized_keys file automatically when deploying from an Image.
	// +kubebuilder:validation:Optional
	AuthorizedUsers []*string `json:"authorizedUsers,omitempty" tf:"authorized_users,omitempty"`

	// The filesystem of this disk. (raw, swap, ext3, ext4, initrd)
	// The filesystem of this disk.
	// +kubebuilder:validation:Optional
	Filesystem *string `json:"filesystem,omitempty" tf:"filesystem,omitempty"`

	// An Image ID to deploy the Linode Disk from.
	// An Image ID to deploy the Linode Disk from.
	// +kubebuilder:validation:Optional
	Image *string `json:"image,omitempty" tf:"image,omitempty"`

	// The Disk's label for display purposes only.
	// The Disk’s label is for display purposes only.
	// +kubebuilder:validation:Optional
	Label *string `json:"label,omitempty" tf:"label,omitempty"`

	// The ID of the Linode to create this Disk under.
	// The ID of the Linode to assign this disk to.
	// +crossplane:generate:reference:type=Instance
	// +kubebuilder:validation:Optional
	LinodeID *float64 `json:"linodeId,omitempty" tf:"linode_id,omitempty"`

	// Reference to a Instance to populate linodeId.
	// +kubebuilder:validation:Optional
	LinodeIDRef *v1.Reference `json:"linodeIdRef,omitempty" tf:"-"`

	// Selector for a Instance to populate linodeId.
	// +kubebuilder:validation:Optional
	LinodeIDSelector *v1.Selector `json:"linodeIdSelector,omitempty" tf:"-"`

	// The root user’s password on a newly-created Linode Disk when deploying from an Image.
	// This sets the root user’s password on a newly-created Linode Disk when deploying from an Image.
	// +kubebuilder:validation:Optional
	RootPassSecretRef *v1.SecretKeySelector `json:"rootPassSecretRef,omitempty" tf:"-"`

	// The size of the Disk in MB. NOTE: Resizing a disk will trigger a Linode reboot.
	// The size of the Disk in MB.
	// +kubebuilder:validation:Optional
	Size *float64 `json:"size,omitempty" tf:"size,omitempty"`

	// An object containing responses to any User Defined Fields present in the StackScript being deployed to this Disk. Only accepted if stackscript_id is given.
	// An object containing responses to any User Defined Fields present in the StackScript being deployed to this Disk. Only accepted if 'stackscript_id' is given. The required values depend on the StackScript being deployed.
	// +kubebuilder:validation:Optional
	StackscriptDataSecretRef *v1.SecretReference `json:"stackscriptDataSecretRef,omitempty" tf:"-"`

	// A StackScript ID that will cause the referenced StackScript to be run during deployment of this Disk.
	// A StackScript ID that will cause the referenced StackScript to be run during deployment of this Linode.
	// +crossplane:generate:reference:type=github.com/linode/provider-linode/apis/stackscript/v1alpha1.Stackscript
	// +kubebuilder:validation:Optional
	StackscriptID *float64 `json:"stackscriptId,omitempty" tf:"stackscript_id,omitempty"`

	// Reference to a Stackscript in stackscript to populate stackscriptId.
	// +kubebuilder:validation:Optional
	StackscriptIDRef *v1.Reference `json:"stackscriptIdRef,omitempty" tf:"-"`

	// Selector for a Stackscript in stackscript to populate stackscriptId.
	// +kubebuilder:validation:Optional
	StackscriptIDSelector *v1.Selector `json:"stackscriptIdSelector,omitempty" tf:"-"`
}

// DiskSpec defines the desired state of Disk
type DiskSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DiskParameters_2 `json:"forProvider"`
}

// DiskStatus defines the observed state of Disk.
type DiskStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DiskObservation_2 `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Disk is the Schema for the Disks API. Manages a Linode Instance Disk.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,linode}
type Disk struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.label)",message="label is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.size)",message="size is a required parameter"
	Spec   DiskSpec   `json:"spec"`
	Status DiskStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DiskList contains a list of Disks
type DiskList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Disk `json:"items"`
}

// Repository type metadata.
var (
	Disk_Kind             = "Disk"
	Disk_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Disk_Kind}.String()
	Disk_KindAPIVersion   = Disk_Kind + "." + CRDGroupVersion.String()
	Disk_GroupVersionKind = CRDGroupVersion.WithKind(Disk_Kind)
)

func init() {
	SchemeBuilder.Register(&Disk{}, &DiskList{})
}
