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

type StorageDefenderInitParameters struct {

	// The max GB to be scanned per Month. Must be -1 or above 0. Omit this property or set to -1 if no capping is needed. Defaults to -1.
	MalwareScanningOnUploadCapGbPerMonth *float64 `json:"malwareScanningOnUploadCapGbPerMonth,omitempty" tf:"malware_scanning_on_upload_cap_gb_per_month,omitempty"`

	// Whether On Upload malware scanning should be enabled. Defaults to false.
	MalwareScanningOnUploadEnabled *bool `json:"malwareScanningOnUploadEnabled,omitempty" tf:"malware_scanning_on_upload_enabled,omitempty"`

	// Whether the settings defined for this storage account should override the settings defined for the subscription. Defaults to false.
	OverrideSubscriptionSettingsEnabled *bool `json:"overrideSubscriptionSettingsEnabled,omitempty" tf:"override_subscription_settings_enabled,omitempty"`

	// The Event Grid Topic where every scan result will be sent to. When you set an Event Grid custom topic, you must set override_subscription_settings_enabled to true to override the subscription-level settings.
	ScanResultsEventGridTopicID *string `json:"scanResultsEventGridTopicId,omitempty" tf:"scan_results_event_grid_topic_id,omitempty"`

	// Whether Sensitive Data Discovery should be enabled. Defaults to false.
	SensitiveDataDiscoveryEnabled *bool `json:"sensitiveDataDiscoveryEnabled,omitempty" tf:"sensitive_data_discovery_enabled,omitempty"`

	// The ID of the storage account the defender applied to. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/storage/v1beta2.Account
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	StorageAccountID *string `json:"storageAccountId,omitempty" tf:"storage_account_id,omitempty"`

	// Reference to a Account in storage to populate storageAccountId.
	// +kubebuilder:validation:Optional
	StorageAccountIDRef *v1.Reference `json:"storageAccountIdRef,omitempty" tf:"-"`

	// Selector for a Account in storage to populate storageAccountId.
	// +kubebuilder:validation:Optional
	StorageAccountIDSelector *v1.Selector `json:"storageAccountIdSelector,omitempty" tf:"-"`
}

type StorageDefenderObservation struct {

	// The Defender for Storage id.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The max GB to be scanned per Month. Must be -1 or above 0. Omit this property or set to -1 if no capping is needed. Defaults to -1.
	MalwareScanningOnUploadCapGbPerMonth *float64 `json:"malwareScanningOnUploadCapGbPerMonth,omitempty" tf:"malware_scanning_on_upload_cap_gb_per_month,omitempty"`

	// Whether On Upload malware scanning should be enabled. Defaults to false.
	MalwareScanningOnUploadEnabled *bool `json:"malwareScanningOnUploadEnabled,omitempty" tf:"malware_scanning_on_upload_enabled,omitempty"`

	// Whether the settings defined for this storage account should override the settings defined for the subscription. Defaults to false.
	OverrideSubscriptionSettingsEnabled *bool `json:"overrideSubscriptionSettingsEnabled,omitempty" tf:"override_subscription_settings_enabled,omitempty"`

	// The Event Grid Topic where every scan result will be sent to. When you set an Event Grid custom topic, you must set override_subscription_settings_enabled to true to override the subscription-level settings.
	ScanResultsEventGridTopicID *string `json:"scanResultsEventGridTopicId,omitempty" tf:"scan_results_event_grid_topic_id,omitempty"`

	// Whether Sensitive Data Discovery should be enabled. Defaults to false.
	SensitiveDataDiscoveryEnabled *bool `json:"sensitiveDataDiscoveryEnabled,omitempty" tf:"sensitive_data_discovery_enabled,omitempty"`

	// The ID of the storage account the defender applied to. Changing this forces a new resource to be created.
	StorageAccountID *string `json:"storageAccountId,omitempty" tf:"storage_account_id,omitempty"`
}

type StorageDefenderParameters struct {

	// The max GB to be scanned per Month. Must be -1 or above 0. Omit this property or set to -1 if no capping is needed. Defaults to -1.
	// +kubebuilder:validation:Optional
	MalwareScanningOnUploadCapGbPerMonth *float64 `json:"malwareScanningOnUploadCapGbPerMonth,omitempty" tf:"malware_scanning_on_upload_cap_gb_per_month,omitempty"`

	// Whether On Upload malware scanning should be enabled. Defaults to false.
	// +kubebuilder:validation:Optional
	MalwareScanningOnUploadEnabled *bool `json:"malwareScanningOnUploadEnabled,omitempty" tf:"malware_scanning_on_upload_enabled,omitempty"`

	// Whether the settings defined for this storage account should override the settings defined for the subscription. Defaults to false.
	// +kubebuilder:validation:Optional
	OverrideSubscriptionSettingsEnabled *bool `json:"overrideSubscriptionSettingsEnabled,omitempty" tf:"override_subscription_settings_enabled,omitempty"`

	// The Event Grid Topic where every scan result will be sent to. When you set an Event Grid custom topic, you must set override_subscription_settings_enabled to true to override the subscription-level settings.
	// +kubebuilder:validation:Optional
	ScanResultsEventGridTopicID *string `json:"scanResultsEventGridTopicId,omitempty" tf:"scan_results_event_grid_topic_id,omitempty"`

	// Whether Sensitive Data Discovery should be enabled. Defaults to false.
	// +kubebuilder:validation:Optional
	SensitiveDataDiscoveryEnabled *bool `json:"sensitiveDataDiscoveryEnabled,omitempty" tf:"sensitive_data_discovery_enabled,omitempty"`

	// The ID of the storage account the defender applied to. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/storage/v1beta2.Account
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	StorageAccountID *string `json:"storageAccountId,omitempty" tf:"storage_account_id,omitempty"`

	// Reference to a Account in storage to populate storageAccountId.
	// +kubebuilder:validation:Optional
	StorageAccountIDRef *v1.Reference `json:"storageAccountIdRef,omitempty" tf:"-"`

	// Selector for a Account in storage to populate storageAccountId.
	// +kubebuilder:validation:Optional
	StorageAccountIDSelector *v1.Selector `json:"storageAccountIdSelector,omitempty" tf:"-"`
}

// StorageDefenderSpec defines the desired state of StorageDefender
type StorageDefenderSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     StorageDefenderParameters `json:"forProvider"`
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
	InitProvider StorageDefenderInitParameters `json:"initProvider,omitempty"`
}

// StorageDefenderStatus defines the observed state of StorageDefender.
type StorageDefenderStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        StorageDefenderObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// StorageDefender is the Schema for the StorageDefenders API. Manages the Defender for Storage.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type StorageDefender struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StorageDefenderSpec   `json:"spec"`
	Status            StorageDefenderStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// StorageDefenderList contains a list of StorageDefenders
type StorageDefenderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StorageDefender `json:"items"`
}

// Repository type metadata.
var (
	StorageDefender_Kind             = "StorageDefender"
	StorageDefender_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: StorageDefender_Kind}.String()
	StorageDefender_KindAPIVersion   = StorageDefender_Kind + "." + CRDGroupVersion.String()
	StorageDefender_GroupVersionKind = CRDGroupVersion.WithKind(StorageDefender_Kind)
)

func init() {
	SchemeBuilder.Register(&StorageDefender{}, &StorageDefenderList{})
}
