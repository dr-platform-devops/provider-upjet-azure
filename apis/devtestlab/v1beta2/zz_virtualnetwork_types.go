// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AllowedPortsInitParameters struct {

	// The port on the Virtual Machine that the traffic will be sent to.
	BackendPort *float64 `json:"backendPort,omitempty" tf:"backend_port,omitempty"`

	// The transport protocol that the traffic will use. Possible values are TCP and UDP.
	TransportProtocol *string `json:"transportProtocol,omitempty" tf:"transport_protocol,omitempty"`
}

type AllowedPortsObservation struct {

	// The port on the Virtual Machine that the traffic will be sent to.
	BackendPort *float64 `json:"backendPort,omitempty" tf:"backend_port,omitempty"`

	// The transport protocol that the traffic will use. Possible values are TCP and UDP.
	TransportProtocol *string `json:"transportProtocol,omitempty" tf:"transport_protocol,omitempty"`
}

type AllowedPortsParameters struct {

	// The port on the Virtual Machine that the traffic will be sent to.
	// +kubebuilder:validation:Optional
	BackendPort *float64 `json:"backendPort,omitempty" tf:"backend_port,omitempty"`

	// The transport protocol that the traffic will use. Possible values are TCP and UDP.
	// +kubebuilder:validation:Optional
	TransportProtocol *string `json:"transportProtocol,omitempty" tf:"transport_protocol,omitempty"`
}

type SharedPublicIPAddressInitParameters struct {

	// A list of allowed_ports blocks as defined below.
	AllowedPorts []AllowedPortsInitParameters `json:"allowedPorts,omitempty" tf:"allowed_ports,omitempty"`
}

type SharedPublicIPAddressObservation struct {

	// A list of allowed_ports blocks as defined below.
	AllowedPorts []AllowedPortsObservation `json:"allowedPorts,omitempty" tf:"allowed_ports,omitempty"`
}

type SharedPublicIPAddressParameters struct {

	// A list of allowed_ports blocks as defined below.
	// +kubebuilder:validation:Optional
	AllowedPorts []AllowedPortsParameters `json:"allowedPorts,omitempty" tf:"allowed_ports,omitempty"`
}

type SubnetInitParameters struct {

	// A shared_public_ip_address block as defined below.
	SharedPublicIPAddress *SharedPublicIPAddressInitParameters `json:"sharedPublicIpAddress,omitempty" tf:"shared_public_ip_address,omitempty"`

	// Can this subnet be used for creating Virtual Machines? Possible values are Allow, Default and Deny. Defaults to Allow.
	UseInVirtualMachineCreation *string `json:"useInVirtualMachineCreation,omitempty" tf:"use_in_virtual_machine_creation,omitempty"`

	// Can Virtual Machines in this Subnet use Public IP Addresses? Possible values are Allow, Default and Deny. Defaults to Allow.
	UsePublicIPAddress *string `json:"usePublicIpAddress,omitempty" tf:"use_public_ip_address,omitempty"`
}

type SubnetObservation struct {

	// The name of the Subnet for this Virtual Network.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A shared_public_ip_address block as defined below.
	SharedPublicIPAddress *SharedPublicIPAddressObservation `json:"sharedPublicIpAddress,omitempty" tf:"shared_public_ip_address,omitempty"`

	// Can this subnet be used for creating Virtual Machines? Possible values are Allow, Default and Deny. Defaults to Allow.
	UseInVirtualMachineCreation *string `json:"useInVirtualMachineCreation,omitempty" tf:"use_in_virtual_machine_creation,omitempty"`

	// Can Virtual Machines in this Subnet use Public IP Addresses? Possible values are Allow, Default and Deny. Defaults to Allow.
	UsePublicIPAddress *string `json:"usePublicIpAddress,omitempty" tf:"use_public_ip_address,omitempty"`
}

type SubnetParameters struct {

	// A shared_public_ip_address block as defined below.
	// +kubebuilder:validation:Optional
	SharedPublicIPAddress *SharedPublicIPAddressParameters `json:"sharedPublicIpAddress,omitempty" tf:"shared_public_ip_address,omitempty"`

	// Can this subnet be used for creating Virtual Machines? Possible values are Allow, Default and Deny. Defaults to Allow.
	// +kubebuilder:validation:Optional
	UseInVirtualMachineCreation *string `json:"useInVirtualMachineCreation,omitempty" tf:"use_in_virtual_machine_creation,omitempty"`

	// Can Virtual Machines in this Subnet use Public IP Addresses? Possible values are Allow, Default and Deny. Defaults to Allow.
	// +kubebuilder:validation:Optional
	UsePublicIPAddress *string `json:"usePublicIpAddress,omitempty" tf:"use_public_ip_address,omitempty"`
}

type VirtualNetworkInitParameters struct {

	// A description for the Virtual Network.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies the name of the Dev Test Lab in which the Virtual Network should be created. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/devtestlab/v1beta1.Lab
	LabName *string `json:"labName,omitempty" tf:"lab_name,omitempty"`

	// Reference to a Lab in devtestlab to populate labName.
	// +kubebuilder:validation:Optional
	LabNameRef *v1.Reference `json:"labNameRef,omitempty" tf:"-"`

	// Selector for a Lab in devtestlab to populate labName.
	// +kubebuilder:validation:Optional
	LabNameSelector *v1.Selector `json:"labNameSelector,omitempty" tf:"-"`

	// Specifies the name of the Dev Test Virtual Network. Changing this forces a new resource to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The name of the resource group in which the Dev Test Lab resource exists. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// A subnet block as defined below.
	Subnet *SubnetInitParameters `json:"subnet,omitempty" tf:"subnet,omitempty"`

	// A mapping of tags to assign to the resource.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type VirtualNetworkObservation struct {

	// A description for the Virtual Network.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the Dev Test Virtual Network.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the name of the Dev Test Lab in which the Virtual Network should be created. Changing this forces a new resource to be created.
	LabName *string `json:"labName,omitempty" tf:"lab_name,omitempty"`

	// Specifies the name of the Dev Test Virtual Network. Changing this forces a new resource to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The name of the resource group in which the Dev Test Lab resource exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// A subnet block as defined below.
	Subnet *SubnetObservation `json:"subnet,omitempty" tf:"subnet,omitempty"`

	// A mapping of tags to assign to the resource.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The unique immutable identifier of the Dev Test Virtual Network.
	UniqueIdentifier *string `json:"uniqueIdentifier,omitempty" tf:"unique_identifier,omitempty"`
}

type VirtualNetworkParameters struct {

	// A description for the Virtual Network.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies the name of the Dev Test Lab in which the Virtual Network should be created. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/devtestlab/v1beta1.Lab
	// +kubebuilder:validation:Optional
	LabName *string `json:"labName,omitempty" tf:"lab_name,omitempty"`

	// Reference to a Lab in devtestlab to populate labName.
	// +kubebuilder:validation:Optional
	LabNameRef *v1.Reference `json:"labNameRef,omitempty" tf:"-"`

	// Selector for a Lab in devtestlab to populate labName.
	// +kubebuilder:validation:Optional
	LabNameSelector *v1.Selector `json:"labNameSelector,omitempty" tf:"-"`

	// Specifies the name of the Dev Test Virtual Network. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The name of the resource group in which the Dev Test Lab resource exists. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// A subnet block as defined below.
	// +kubebuilder:validation:Optional
	Subnet *SubnetParameters `json:"subnet,omitempty" tf:"subnet,omitempty"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// VirtualNetworkSpec defines the desired state of VirtualNetwork
type VirtualNetworkSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VirtualNetworkParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider VirtualNetworkInitParameters `json:"initProvider,omitempty"`
}

// VirtualNetworkStatus defines the observed state of VirtualNetwork.
type VirtualNetworkStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VirtualNetworkObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// VirtualNetwork is the Schema for the VirtualNetworks API. Manages a Virtual Network within a DevTest Lab.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type VirtualNetwork struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   VirtualNetworkSpec   `json:"spec"`
	Status VirtualNetworkStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VirtualNetworkList contains a list of VirtualNetworks
type VirtualNetworkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualNetwork `json:"items"`
}

// Repository type metadata.
var (
	VirtualNetwork_Kind             = "VirtualNetwork"
	VirtualNetwork_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VirtualNetwork_Kind}.String()
	VirtualNetwork_KindAPIVersion   = VirtualNetwork_Kind + "." + CRDGroupVersion.String()
	VirtualNetwork_GroupVersionKind = CRDGroupVersion.WithKind(VirtualNetwork_Kind)
)

func init() {
	SchemeBuilder.Register(&VirtualNetwork{}, &VirtualNetworkList{})
}
