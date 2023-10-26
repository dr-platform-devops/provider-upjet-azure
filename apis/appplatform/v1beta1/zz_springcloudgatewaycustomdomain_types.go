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

type SpringCloudGatewayCustomDomainInitParameters struct {

	// The name which should be used for this Spring Cloud Gateway Custom Domain. Changing this forces a new Spring Cloud Gateway Custom Domain to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the thumbprint of the Spring Cloud Certificate that binds to the Spring Cloud Gateway Custom Domain.
	Thumbprint *string `json:"thumbprint,omitempty" tf:"thumbprint,omitempty"`
}

type SpringCloudGatewayCustomDomainObservation struct {

	// The ID of the Spring Cloud Gateway Custom Domain.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name which should be used for this Spring Cloud Gateway Custom Domain. Changing this forces a new Spring Cloud Gateway Custom Domain to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the Spring Cloud Gateway. Changing this forces a new Spring Cloud Gateway Custom Domain to be created.
	SpringCloudGatewayID *string `json:"springCloudGatewayId,omitempty" tf:"spring_cloud_gateway_id,omitempty"`

	// Specifies the thumbprint of the Spring Cloud Certificate that binds to the Spring Cloud Gateway Custom Domain.
	Thumbprint *string `json:"thumbprint,omitempty" tf:"thumbprint,omitempty"`
}

type SpringCloudGatewayCustomDomainParameters struct {

	// The name which should be used for this Spring Cloud Gateway Custom Domain. Changing this forces a new Spring Cloud Gateway Custom Domain to be created.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the Spring Cloud Gateway. Changing this forces a new Spring Cloud Gateway Custom Domain to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/appplatform/v1beta1.SpringCloudGateway
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SpringCloudGatewayID *string `json:"springCloudGatewayId,omitempty" tf:"spring_cloud_gateway_id,omitempty"`

	// Reference to a SpringCloudGateway in appplatform to populate springCloudGatewayId.
	// +kubebuilder:validation:Optional
	SpringCloudGatewayIDRef *v1.Reference `json:"springCloudGatewayIdRef,omitempty" tf:"-"`

	// Selector for a SpringCloudGateway in appplatform to populate springCloudGatewayId.
	// +kubebuilder:validation:Optional
	SpringCloudGatewayIDSelector *v1.Selector `json:"springCloudGatewayIdSelector,omitempty" tf:"-"`

	// Specifies the thumbprint of the Spring Cloud Certificate that binds to the Spring Cloud Gateway Custom Domain.
	// +kubebuilder:validation:Optional
	Thumbprint *string `json:"thumbprint,omitempty" tf:"thumbprint,omitempty"`
}

// SpringCloudGatewayCustomDomainSpec defines the desired state of SpringCloudGatewayCustomDomain
type SpringCloudGatewayCustomDomainSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SpringCloudGatewayCustomDomainParameters `json:"forProvider"`
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
	InitProvider SpringCloudGatewayCustomDomainInitParameters `json:"initProvider,omitempty"`
}

// SpringCloudGatewayCustomDomainStatus defines the observed state of SpringCloudGatewayCustomDomain.
type SpringCloudGatewayCustomDomainStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SpringCloudGatewayCustomDomainObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SpringCloudGatewayCustomDomain is the Schema for the SpringCloudGatewayCustomDomains API. Manages a Spring Cloud Gateway Custom Domain.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type SpringCloudGatewayCustomDomain struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   SpringCloudGatewayCustomDomainSpec   `json:"spec"`
	Status SpringCloudGatewayCustomDomainStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SpringCloudGatewayCustomDomainList contains a list of SpringCloudGatewayCustomDomains
type SpringCloudGatewayCustomDomainList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SpringCloudGatewayCustomDomain `json:"items"`
}

// Repository type metadata.
var (
	SpringCloudGatewayCustomDomain_Kind             = "SpringCloudGatewayCustomDomain"
	SpringCloudGatewayCustomDomain_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SpringCloudGatewayCustomDomain_Kind}.String()
	SpringCloudGatewayCustomDomain_KindAPIVersion   = SpringCloudGatewayCustomDomain_Kind + "." + CRDGroupVersion.String()
	SpringCloudGatewayCustomDomain_GroupVersionKind = CRDGroupVersion.WithKind(SpringCloudGatewayCustomDomain_Kind)
)

func init() {
	SchemeBuilder.Register(&SpringCloudGatewayCustomDomain{}, &SpringCloudGatewayCustomDomainList{})
}
