// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ManagedPrivateEndpointInitParameters struct {

	// Specifies the sub resource name which the Stream Analytics Private Endpoint is able to connect to. Changing this forces a new resource to be created.
	SubresourceName *string `json:"subresourceName,omitempty" tf:"subresource_name,omitempty"`
}

type ManagedPrivateEndpointObservation struct {

	// The ID of the Stream Analytics.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the Resource Group where the Stream Analytics Managed Private Endpoint should exist. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// The name of the Stream Analytics Cluster where the Managed Private Endpoint should be created. Changing this forces a new resource to be created.
	StreamAnalyticsClusterName *string `json:"streamAnalyticsClusterName,omitempty" tf:"stream_analytics_cluster_name,omitempty"`

	// Specifies the sub resource name which the Stream Analytics Private Endpoint is able to connect to. Changing this forces a new resource to be created.
	SubresourceName *string `json:"subresourceName,omitempty" tf:"subresource_name,omitempty"`

	// The ID of the Private Link Enabled Remote Resource which this Stream Analytics Private endpoint should be connected to. Changing this forces a new resource to be created.
	TargetResourceID *string `json:"targetResourceId,omitempty" tf:"target_resource_id,omitempty"`
}

type ManagedPrivateEndpointParameters struct {

	// The name of the Resource Group where the Stream Analytics Managed Private Endpoint should exist. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The name of the Stream Analytics Cluster where the Managed Private Endpoint should be created. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/streamanalytics/v1beta1.Cluster
	// +kubebuilder:validation:Optional
	StreamAnalyticsClusterName *string `json:"streamAnalyticsClusterName,omitempty" tf:"stream_analytics_cluster_name,omitempty"`

	// Reference to a Cluster in streamanalytics to populate streamAnalyticsClusterName.
	// +kubebuilder:validation:Optional
	StreamAnalyticsClusterNameRef *v1.Reference `json:"streamAnalyticsClusterNameRef,omitempty" tf:"-"`

	// Selector for a Cluster in streamanalytics to populate streamAnalyticsClusterName.
	// +kubebuilder:validation:Optional
	StreamAnalyticsClusterNameSelector *v1.Selector `json:"streamAnalyticsClusterNameSelector,omitempty" tf:"-"`

	// Specifies the sub resource name which the Stream Analytics Private Endpoint is able to connect to. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	SubresourceName *string `json:"subresourceName,omitempty" tf:"subresource_name,omitempty"`

	// The ID of the Private Link Enabled Remote Resource which this Stream Analytics Private endpoint should be connected to. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/storage/v1beta1.Account
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	TargetResourceID *string `json:"targetResourceId,omitempty" tf:"target_resource_id,omitempty"`

	// Reference to a Account in storage to populate targetResourceId.
	// +kubebuilder:validation:Optional
	TargetResourceIDRef *v1.Reference `json:"targetResourceIdRef,omitempty" tf:"-"`

	// Selector for a Account in storage to populate targetResourceId.
	// +kubebuilder:validation:Optional
	TargetResourceIDSelector *v1.Selector `json:"targetResourceIdSelector,omitempty" tf:"-"`
}

// ManagedPrivateEndpointSpec defines the desired state of ManagedPrivateEndpoint
type ManagedPrivateEndpointSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ManagedPrivateEndpointParameters `json:"forProvider"`
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
	InitProvider ManagedPrivateEndpointInitParameters `json:"initProvider,omitempty"`
}

// ManagedPrivateEndpointStatus defines the observed state of ManagedPrivateEndpoint.
type ManagedPrivateEndpointStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ManagedPrivateEndpointObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ManagedPrivateEndpoint is the Schema for the ManagedPrivateEndpoints API. Manages a Stream Analytics Managed Private Endpoint.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ManagedPrivateEndpoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.subresourceName) || (has(self.initProvider) && has(self.initProvider.subresourceName))",message="spec.forProvider.subresourceName is a required parameter"
	Spec   ManagedPrivateEndpointSpec   `json:"spec"`
	Status ManagedPrivateEndpointStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ManagedPrivateEndpointList contains a list of ManagedPrivateEndpoints
type ManagedPrivateEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ManagedPrivateEndpoint `json:"items"`
}

// Repository type metadata.
var (
	ManagedPrivateEndpoint_Kind             = "ManagedPrivateEndpoint"
	ManagedPrivateEndpoint_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ManagedPrivateEndpoint_Kind}.String()
	ManagedPrivateEndpoint_KindAPIVersion   = ManagedPrivateEndpoint_Kind + "." + CRDGroupVersion.String()
	ManagedPrivateEndpoint_GroupVersionKind = CRDGroupVersion.WithKind(ManagedPrivateEndpoint_Kind)
)

func init() {
	SchemeBuilder.Register(&ManagedPrivateEndpoint{}, &ManagedPrivateEndpointList{})
}
