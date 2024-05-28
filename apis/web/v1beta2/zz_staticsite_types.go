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

type StaticSiteIdentityInitParameters struct {

	// A list of Managed Identity IDs which should be assigned to this Static Site resource.
	// +listType=set
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// The Type of Managed Identity assigned to this Static Site resource. Possible values are SystemAssigned, UserAssigned and SystemAssigned, UserAssigned.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type StaticSiteIdentityObservation struct {

	// A list of Managed Identity IDs which should be assigned to this Static Site resource.
	// +listType=set
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// The Principal ID associated with this Managed Service Identity.
	PrincipalID *string `json:"principalId,omitempty" tf:"principal_id,omitempty"`

	// The ID of the Static Web App.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	// The Type of Managed Identity assigned to this Static Site resource. Possible values are SystemAssigned, UserAssigned and SystemAssigned, UserAssigned.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type StaticSiteIdentityParameters struct {

	// A list of Managed Identity IDs which should be assigned to this Static Site resource.
	// +kubebuilder:validation:Optional
	// +listType=set
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// The Type of Managed Identity assigned to this Static Site resource. Possible values are SystemAssigned, UserAssigned and SystemAssigned, UserAssigned.
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

type StaticSiteInitParameters struct {

	// A key-value pair of App Settings.
	// +mapType=granular
	AppSettings map[string]*string `json:"appSettings,omitempty" tf:"app_settings,omitempty"`

	// An identity block as defined below.
	Identity *StaticSiteIdentityInitParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// The Azure Region where the Static Web App should exist. Changing this forces a new Static Web App to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Specifies the SKU size of the Static Web App. Possible values are Free or Standard. Defaults to Free.
	SkuSize *string `json:"skuSize,omitempty" tf:"sku_size,omitempty"`

	// Specifies the SKU tier of the Static Web App. Possible values are Free or Standard. Defaults to Free.
	SkuTier *string `json:"skuTier,omitempty" tf:"sku_tier,omitempty"`

	// A mapping of tags to assign to the resource.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type StaticSiteObservation struct {

	// A key-value pair of App Settings.
	// +mapType=granular
	AppSettings map[string]*string `json:"appSettings,omitempty" tf:"app_settings,omitempty"`

	// The default host name of the Static Web App.
	DefaultHostName *string `json:"defaultHostName,omitempty" tf:"default_host_name,omitempty"`

	// The ID of the Static Web App.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// An identity block as defined below.
	Identity *StaticSiteIdentityObservation `json:"identity,omitempty" tf:"identity,omitempty"`

	// The Azure Region where the Static Web App should exist. Changing this forces a new Static Web App to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The name of the Resource Group where the Static Web App should exist. Changing this forces a new Static Web App to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Specifies the SKU size of the Static Web App. Possible values are Free or Standard. Defaults to Free.
	SkuSize *string `json:"skuSize,omitempty" tf:"sku_size,omitempty"`

	// Specifies the SKU tier of the Static Web App. Possible values are Free or Standard. Defaults to Free.
	SkuTier *string `json:"skuTier,omitempty" tf:"sku_tier,omitempty"`

	// A mapping of tags to assign to the resource.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type StaticSiteParameters struct {

	// A key-value pair of App Settings.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	AppSettings map[string]*string `json:"appSettings,omitempty" tf:"app_settings,omitempty"`

	// An identity block as defined below.
	// +kubebuilder:validation:Optional
	Identity *StaticSiteIdentityParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// The Azure Region where the Static Web App should exist. Changing this forces a new Static Web App to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The name of the Resource Group where the Static Web App should exist. Changing this forces a new Static Web App to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// Specifies the SKU size of the Static Web App. Possible values are Free or Standard. Defaults to Free.
	// +kubebuilder:validation:Optional
	SkuSize *string `json:"skuSize,omitempty" tf:"sku_size,omitempty"`

	// Specifies the SKU tier of the Static Web App. Possible values are Free or Standard. Defaults to Free.
	// +kubebuilder:validation:Optional
	SkuTier *string `json:"skuTier,omitempty" tf:"sku_tier,omitempty"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// StaticSiteSpec defines the desired state of StaticSite
type StaticSiteSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     StaticSiteParameters `json:"forProvider"`
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
	InitProvider StaticSiteInitParameters `json:"initProvider,omitempty"`
}

// StaticSiteStatus defines the observed state of StaticSite.
type StaticSiteStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        StaticSiteObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// StaticSite is the Schema for the StaticSites API. Manages a Static Site.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type StaticSite struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || (has(self.initProvider) && has(self.initProvider.location))",message="spec.forProvider.location is a required parameter"
	Spec   StaticSiteSpec   `json:"spec"`
	Status StaticSiteStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// StaticSiteList contains a list of StaticSites
type StaticSiteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StaticSite `json:"items"`
}

// Repository type metadata.
var (
	StaticSite_Kind             = "StaticSite"
	StaticSite_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: StaticSite_Kind}.String()
	StaticSite_KindAPIVersion   = StaticSite_Kind + "." + CRDGroupVersion.String()
	StaticSite_GroupVersionKind = CRDGroupVersion.WithKind(StaticSite_Kind)
)

func init() {
	SchemeBuilder.Register(&StaticSite{}, &StaticSiteList{})
}