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

type AzureadBasedServicePrincipalInitParameters struct {

	// Specifies the Ledger Role to grant this AzureAD Service Principal. Possible values are Administrator, Contributor and Reader.
	LedgerRoleName *string `json:"ledgerRoleName,omitempty" tf:"ledger_role_name,omitempty"`

	// Specifies the Principal ID of the AzureAD Service Principal.
	PrincipalID *string `json:"principalId,omitempty" tf:"principal_id,omitempty"`

	// Specifies the Tenant ID for this AzureAD Service Principal.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

type AzureadBasedServicePrincipalObservation struct {

	// Specifies the Ledger Role to grant this AzureAD Service Principal. Possible values are Administrator, Contributor and Reader.
	LedgerRoleName *string `json:"ledgerRoleName,omitempty" tf:"ledger_role_name,omitempty"`

	// Specifies the Principal ID of the AzureAD Service Principal.
	PrincipalID *string `json:"principalId,omitempty" tf:"principal_id,omitempty"`

	// Specifies the Tenant ID for this AzureAD Service Principal.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

type AzureadBasedServicePrincipalParameters struct {

	// Specifies the Ledger Role to grant this AzureAD Service Principal. Possible values are Administrator, Contributor and Reader.
	// +kubebuilder:validation:Optional
	LedgerRoleName *string `json:"ledgerRoleName" tf:"ledger_role_name,omitempty"`

	// Specifies the Principal ID of the AzureAD Service Principal.
	// +kubebuilder:validation:Optional
	PrincipalID *string `json:"principalId" tf:"principal_id,omitempty"`

	// Specifies the Tenant ID for this AzureAD Service Principal.
	// +kubebuilder:validation:Optional
	TenantID *string `json:"tenantId" tf:"tenant_id,omitempty"`
}

type CertificateBasedSecurityPrincipalInitParameters struct {

	// Specifies the Ledger Role to grant this Certificate Security Principal. Possible values are Administrator, Contributor and Reader.
	LedgerRoleName *string `json:"ledgerRoleName,omitempty" tf:"ledger_role_name,omitempty"`

	// The public key, in PEM format, of the certificate used by this identity to authenticate with the Confidential Ledger.
	PemPublicKey *string `json:"pemPublicKey,omitempty" tf:"pem_public_key,omitempty"`
}

type CertificateBasedSecurityPrincipalObservation struct {

	// Specifies the Ledger Role to grant this Certificate Security Principal. Possible values are Administrator, Contributor and Reader.
	LedgerRoleName *string `json:"ledgerRoleName,omitempty" tf:"ledger_role_name,omitempty"`

	// The public key, in PEM format, of the certificate used by this identity to authenticate with the Confidential Ledger.
	PemPublicKey *string `json:"pemPublicKey,omitempty" tf:"pem_public_key,omitempty"`
}

type CertificateBasedSecurityPrincipalParameters struct {

	// Specifies the Ledger Role to grant this Certificate Security Principal. Possible values are Administrator, Contributor and Reader.
	// +kubebuilder:validation:Optional
	LedgerRoleName *string `json:"ledgerRoleName" tf:"ledger_role_name,omitempty"`

	// The public key, in PEM format, of the certificate used by this identity to authenticate with the Confidential Ledger.
	// +kubebuilder:validation:Optional
	PemPublicKey *string `json:"pemPublicKey" tf:"pem_public_key,omitempty"`
}

type LedgerInitParameters struct {

	// A list of azuread_based_service_principal blocks as defined below.
	AzureadBasedServicePrincipal []AzureadBasedServicePrincipalInitParameters `json:"azureadBasedServicePrincipal,omitempty" tf:"azuread_based_service_principal,omitempty"`

	// A list of certificate_based_security_principal blocks as defined below.
	CertificateBasedSecurityPrincipal []CertificateBasedSecurityPrincipalInitParameters `json:"certificateBasedSecurityPrincipal,omitempty" tf:"certificate_based_security_principal,omitempty"`

	// Specifies the type of Confidential Ledger. Possible values are Private and Public. Changing this forces a new resource to be created.
	LedgerType *string `json:"ledgerType,omitempty" tf:"ledger_type,omitempty"`

	// Specifies the supported Azure location where the Confidential Ledger exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// A mapping of tags to assign to the Confidential Ledger.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type LedgerObservation struct {

	// A list of azuread_based_service_principal blocks as defined below.
	AzureadBasedServicePrincipal []AzureadBasedServicePrincipalObservation `json:"azureadBasedServicePrincipal,omitempty" tf:"azuread_based_service_principal,omitempty"`

	// A list of certificate_based_security_principal blocks as defined below.
	CertificateBasedSecurityPrincipal []CertificateBasedSecurityPrincipalObservation `json:"certificateBasedSecurityPrincipal,omitempty" tf:"certificate_based_security_principal,omitempty"`

	// The ID of this Confidential Ledger.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The Identity Service Endpoint for this Confidential Ledger.
	IdentityServiceEndpoint *string `json:"identityServiceEndpoint,omitempty" tf:"identity_service_endpoint,omitempty"`

	// The Endpoint for this Confidential Ledger.
	LedgerEndpoint *string `json:"ledgerEndpoint,omitempty" tf:"ledger_endpoint,omitempty"`

	// Specifies the type of Confidential Ledger. Possible values are Private and Public. Changing this forces a new resource to be created.
	LedgerType *string `json:"ledgerType,omitempty" tf:"ledger_type,omitempty"`

	// Specifies the supported Azure location where the Confidential Ledger exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The name of the Resource Group where the Confidential Ledger exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// A mapping of tags to assign to the Confidential Ledger.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type LedgerParameters struct {

	// A list of azuread_based_service_principal blocks as defined below.
	// +kubebuilder:validation:Optional
	AzureadBasedServicePrincipal []AzureadBasedServicePrincipalParameters `json:"azureadBasedServicePrincipal,omitempty" tf:"azuread_based_service_principal,omitempty"`

	// A list of certificate_based_security_principal blocks as defined below.
	// +kubebuilder:validation:Optional
	CertificateBasedSecurityPrincipal []CertificateBasedSecurityPrincipalParameters `json:"certificateBasedSecurityPrincipal,omitempty" tf:"certificate_based_security_principal,omitempty"`

	// Specifies the type of Confidential Ledger. Possible values are Private and Public. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	LedgerType *string `json:"ledgerType,omitempty" tf:"ledger_type,omitempty"`

	// Specifies the supported Azure location where the Confidential Ledger exists. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The name of the Resource Group where the Confidential Ledger exists. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// A mapping of tags to assign to the Confidential Ledger.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// LedgerSpec defines the desired state of Ledger
type LedgerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LedgerParameters `json:"forProvider"`
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
	InitProvider LedgerInitParameters `json:"initProvider,omitempty"`
}

// LedgerStatus defines the observed state of Ledger.
type LedgerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LedgerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Ledger is the Schema for the Ledgers API. Manages a Confidential Ledger.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type Ledger struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.azureadBasedServicePrincipal) || (has(self.initProvider) && has(self.initProvider.azureadBasedServicePrincipal))",message="spec.forProvider.azureadBasedServicePrincipal is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.ledgerType) || (has(self.initProvider) && has(self.initProvider.ledgerType))",message="spec.forProvider.ledgerType is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || (has(self.initProvider) && has(self.initProvider.location))",message="spec.forProvider.location is a required parameter"
	Spec   LedgerSpec   `json:"spec"`
	Status LedgerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LedgerList contains a list of Ledgers
type LedgerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Ledger `json:"items"`
}

// Repository type metadata.
var (
	Ledger_Kind             = "Ledger"
	Ledger_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Ledger_Kind}.String()
	Ledger_KindAPIVersion   = Ledger_Kind + "." + CRDGroupVersion.String()
	Ledger_GroupVersionKind = CRDGroupVersion.WithKind(Ledger_Kind)
)

func init() {
	SchemeBuilder.Register(&Ledger{}, &LedgerList{})
}
