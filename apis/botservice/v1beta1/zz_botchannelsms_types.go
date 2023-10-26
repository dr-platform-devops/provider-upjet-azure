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

type BotChannelSMSInitParameters struct {

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The phone number for the SMS Channel.
	PhoneNumber *string `json:"phoneNumber,omitempty" tf:"phone_number,omitempty"`

	// The account security identifier (SID) for the SMS Channel.
	SMSChannelAccountSecurityID *string `json:"smsChannelAccountSecurityId,omitempty" tf:"sms_channel_account_security_id,omitempty"`
}

type BotChannelSMSObservation struct {

	// The name of the Bot Resource this channel will be associated with. Changing this forces a new resource to be created.
	BotName *string `json:"botName,omitempty" tf:"bot_name,omitempty"`

	// The ID of the SMS Integration for a Bot Channel.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The phone number for the SMS Channel.
	PhoneNumber *string `json:"phoneNumber,omitempty" tf:"phone_number,omitempty"`

	// The name of the resource group where the SMS Channel should be created. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// The account security identifier (SID) for the SMS Channel.
	SMSChannelAccountSecurityID *string `json:"smsChannelAccountSecurityId,omitempty" tf:"sms_channel_account_security_id,omitempty"`
}

type BotChannelSMSParameters struct {

	// The name of the Bot Resource this channel will be associated with. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/botservice/v1beta1.BotChannelsRegistration
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("name",false)
	// +kubebuilder:validation:Optional
	BotName *string `json:"botName,omitempty" tf:"bot_name,omitempty"`

	// Reference to a BotChannelsRegistration in botservice to populate botName.
	// +kubebuilder:validation:Optional
	BotNameRef *v1.Reference `json:"botNameRef,omitempty" tf:"-"`

	// Selector for a BotChannelsRegistration in botservice to populate botName.
	// +kubebuilder:validation:Optional
	BotNameSelector *v1.Selector `json:"botNameSelector,omitempty" tf:"-"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The phone number for the SMS Channel.
	// +kubebuilder:validation:Optional
	PhoneNumber *string `json:"phoneNumber,omitempty" tf:"phone_number,omitempty"`

	// The name of the resource group where the SMS Channel should be created. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The account security identifier (SID) for the SMS Channel.
	// +kubebuilder:validation:Optional
	SMSChannelAccountSecurityID *string `json:"smsChannelAccountSecurityId,omitempty" tf:"sms_channel_account_security_id,omitempty"`

	// The authorization token for the SMS Channel.
	// +kubebuilder:validation:Optional
	SMSChannelAuthTokenSecretRef v1.SecretKeySelector `json:"smsChannelAuthTokenSecretRef" tf:"-"`
}

// BotChannelSMSSpec defines the desired state of BotChannelSMS
type BotChannelSMSSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BotChannelSMSParameters `json:"forProvider"`
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
	InitProvider BotChannelSMSInitParameters `json:"initProvider,omitempty"`
}

// BotChannelSMSStatus defines the observed state of BotChannelSMS.
type BotChannelSMSStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BotChannelSMSObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BotChannelSMS is the Schema for the BotChannelSMSs API. Manages a SMS integration for a Bot Channel
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure},path=botchannelsms
type BotChannelSMS struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || (has(self.initProvider) && has(self.initProvider.location))",message="spec.forProvider.location is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.phoneNumber) || (has(self.initProvider) && has(self.initProvider.phoneNumber))",message="spec.forProvider.phoneNumber is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.smsChannelAccountSecurityId) || (has(self.initProvider) && has(self.initProvider.smsChannelAccountSecurityId))",message="spec.forProvider.smsChannelAccountSecurityId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.smsChannelAuthTokenSecretRef)",message="spec.forProvider.smsChannelAuthTokenSecretRef is a required parameter"
	Spec   BotChannelSMSSpec   `json:"spec"`
	Status BotChannelSMSStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BotChannelSMSList contains a list of BotChannelSMSs
type BotChannelSMSList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BotChannelSMS `json:"items"`
}

// Repository type metadata.
var (
	BotChannelSMS_Kind             = "BotChannelSMS"
	BotChannelSMS_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BotChannelSMS_Kind}.String()
	BotChannelSMS_KindAPIVersion   = BotChannelSMS_Kind + "." + CRDGroupVersion.String()
	BotChannelSMS_GroupVersionKind = CRDGroupVersion.WithKind(BotChannelSMS_Kind)
)

func init() {
	SchemeBuilder.Register(&BotChannelSMS{}, &BotChannelSMSList{})
}
