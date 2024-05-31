// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type RedisCacheInitParameters struct {

	// The location where to use cache from. Possible values are default and valid Azure regions. Defaults to default.
	CacheLocation *string `json:"cacheLocation,omitempty" tf:"cache_location,omitempty"`

	// The connection string to the Cache for Redis.
	ConnectionStringSecretRef v1.SecretKeySelector `json:"connectionStringSecretRef" tf:"-"`

	// The description of the API Management Redis Cache.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The resource ID of the Cache for Redis.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/cache/v1beta2.RedisCache
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	RedisCacheID *string `json:"redisCacheId,omitempty" tf:"redis_cache_id,omitempty"`

	// Reference to a RedisCache in cache to populate redisCacheId.
	// +kubebuilder:validation:Optional
	RedisCacheIDRef *v1.Reference `json:"redisCacheIdRef,omitempty" tf:"-"`

	// Selector for a RedisCache in cache to populate redisCacheId.
	// +kubebuilder:validation:Optional
	RedisCacheIDSelector *v1.Selector `json:"redisCacheIdSelector,omitempty" tf:"-"`
}

type RedisCacheObservation struct {

	// The resource ID of the API Management Service from which to create this external cache. Changing this forces a new API Management Redis Cache to be created.
	APIManagementID *string `json:"apiManagementId,omitempty" tf:"api_management_id,omitempty"`

	// The location where to use cache from. Possible values are default and valid Azure regions. Defaults to default.
	CacheLocation *string `json:"cacheLocation,omitempty" tf:"cache_location,omitempty"`

	// The description of the API Management Redis Cache.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the API Management Redis Cache.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The resource ID of the Cache for Redis.
	RedisCacheID *string `json:"redisCacheId,omitempty" tf:"redis_cache_id,omitempty"`
}

type RedisCacheParameters struct {

	// The resource ID of the API Management Service from which to create this external cache. Changing this forces a new API Management Redis Cache to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/apimanagement/v1beta2.Management
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	APIManagementID *string `json:"apiManagementId,omitempty" tf:"api_management_id,omitempty"`

	// Reference to a Management in apimanagement to populate apiManagementId.
	// +kubebuilder:validation:Optional
	APIManagementIDRef *v1.Reference `json:"apiManagementIdRef,omitempty" tf:"-"`

	// Selector for a Management in apimanagement to populate apiManagementId.
	// +kubebuilder:validation:Optional
	APIManagementIDSelector *v1.Selector `json:"apiManagementIdSelector,omitempty" tf:"-"`

	// The location where to use cache from. Possible values are default and valid Azure regions. Defaults to default.
	// +kubebuilder:validation:Optional
	CacheLocation *string `json:"cacheLocation,omitempty" tf:"cache_location,omitempty"`

	// The connection string to the Cache for Redis.
	// +kubebuilder:validation:Optional
	ConnectionStringSecretRef v1.SecretKeySelector `json:"connectionStringSecretRef" tf:"-"`

	// The description of the API Management Redis Cache.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The resource ID of the Cache for Redis.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/cache/v1beta2.RedisCache
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	RedisCacheID *string `json:"redisCacheId,omitempty" tf:"redis_cache_id,omitempty"`

	// Reference to a RedisCache in cache to populate redisCacheId.
	// +kubebuilder:validation:Optional
	RedisCacheIDRef *v1.Reference `json:"redisCacheIdRef,omitempty" tf:"-"`

	// Selector for a RedisCache in cache to populate redisCacheId.
	// +kubebuilder:validation:Optional
	RedisCacheIDSelector *v1.Selector `json:"redisCacheIdSelector,omitempty" tf:"-"`
}

// RedisCacheSpec defines the desired state of RedisCache
type RedisCacheSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RedisCacheParameters `json:"forProvider"`
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
	InitProvider RedisCacheInitParameters `json:"initProvider,omitempty"`
}

// RedisCacheStatus defines the observed state of RedisCache.
type RedisCacheStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RedisCacheObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// RedisCache is the Schema for the RedisCaches API. Manages a API Management Redis Cache.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type RedisCache struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.connectionStringSecretRef)",message="spec.forProvider.connectionStringSecretRef is a required parameter"
	Spec   RedisCacheSpec   `json:"spec"`
	Status RedisCacheStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RedisCacheList contains a list of RedisCaches
type RedisCacheList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RedisCache `json:"items"`
}

// Repository type metadata.
var (
	RedisCache_Kind             = "RedisCache"
	RedisCache_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RedisCache_Kind}.String()
	RedisCache_KindAPIVersion   = RedisCache_Kind + "." + CRDGroupVersion.String()
	RedisCache_GroupVersionKind = CRDGroupVersion.WithKind(RedisCache_Kind)
)

func init() {
	SchemeBuilder.Register(&RedisCache{}, &RedisCacheList{})
}
