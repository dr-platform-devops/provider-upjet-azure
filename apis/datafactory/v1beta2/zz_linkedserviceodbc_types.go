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

type LinkedServiceOdbcBasicAuthenticationInitParameters struct {

	// The password associated with the username, which can be used to authenticate to the ODBC endpoint.
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef" tf:"-"`

	// The username which can be used to authenticate to the ODBC endpoint.
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type LinkedServiceOdbcBasicAuthenticationObservation struct {

	// The username which can be used to authenticate to the ODBC endpoint.
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type LinkedServiceOdbcBasicAuthenticationParameters struct {

	// The password associated with the username, which can be used to authenticate to the ODBC endpoint.
	// +kubebuilder:validation:Optional
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef" tf:"-"`

	// The username which can be used to authenticate to the ODBC endpoint.
	// +kubebuilder:validation:Optional
	Username *string `json:"username" tf:"username,omitempty"`
}

type LinkedServiceOdbcInitParameters struct {

	// A map of additional properties to associate with the Data Factory Linked Service ODBC.
	// +mapType=granular
	AdditionalProperties map[string]*string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// List of tags that can be used for describing the Data Factory Linked Service ODBC.
	Annotations []*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// A basic_authentication block as defined below.
	BasicAuthentication *LinkedServiceOdbcBasicAuthenticationInitParameters `json:"basicAuthentication,omitempty" tf:"basic_authentication,omitempty"`

	// The connection string in which to authenticate with ODBC.
	ConnectionString *string `json:"connectionString,omitempty" tf:"connection_string,omitempty"`

	// The description for the Data Factory Linked Service ODBC.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The integration runtime reference to associate with the Data Factory Linked Service ODBC.
	IntegrationRuntimeName *string `json:"integrationRuntimeName,omitempty" tf:"integration_runtime_name,omitempty"`

	// A map of parameters to associate with the Data Factory Linked Service ODBC.
	// +mapType=granular
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

type LinkedServiceOdbcObservation struct {

	// A map of additional properties to associate with the Data Factory Linked Service ODBC.
	// +mapType=granular
	AdditionalProperties map[string]*string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// List of tags that can be used for describing the Data Factory Linked Service ODBC.
	Annotations []*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// A basic_authentication block as defined below.
	BasicAuthentication *LinkedServiceOdbcBasicAuthenticationObservation `json:"basicAuthentication,omitempty" tf:"basic_authentication,omitempty"`

	// The connection string in which to authenticate with ODBC.
	ConnectionString *string `json:"connectionString,omitempty" tf:"connection_string,omitempty"`

	// The Data Factory ID in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryID *string `json:"dataFactoryId,omitempty" tf:"data_factory_id,omitempty"`

	// The description for the Data Factory Linked Service ODBC.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the Data Factory ODBC Linked Service.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The integration runtime reference to associate with the Data Factory Linked Service ODBC.
	IntegrationRuntimeName *string `json:"integrationRuntimeName,omitempty" tf:"integration_runtime_name,omitempty"`

	// A map of parameters to associate with the Data Factory Linked Service ODBC.
	// +mapType=granular
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

type LinkedServiceOdbcParameters struct {

	// A map of additional properties to associate with the Data Factory Linked Service ODBC.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	AdditionalProperties map[string]*string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// List of tags that can be used for describing the Data Factory Linked Service ODBC.
	// +kubebuilder:validation:Optional
	Annotations []*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// A basic_authentication block as defined below.
	// +kubebuilder:validation:Optional
	BasicAuthentication *LinkedServiceOdbcBasicAuthenticationParameters `json:"basicAuthentication,omitempty" tf:"basic_authentication,omitempty"`

	// The connection string in which to authenticate with ODBC.
	// +kubebuilder:validation:Optional
	ConnectionString *string `json:"connectionString,omitempty" tf:"connection_string,omitempty"`

	// The Data Factory ID in which to associate the Linked Service with. Changing this forces a new resource.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/datafactory/v1beta2.Factory
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	DataFactoryID *string `json:"dataFactoryId,omitempty" tf:"data_factory_id,omitempty"`

	// Reference to a Factory in datafactory to populate dataFactoryId.
	// +kubebuilder:validation:Optional
	DataFactoryIDRef *v1.Reference `json:"dataFactoryIdRef,omitempty" tf:"-"`

	// Selector for a Factory in datafactory to populate dataFactoryId.
	// +kubebuilder:validation:Optional
	DataFactoryIDSelector *v1.Selector `json:"dataFactoryIdSelector,omitempty" tf:"-"`

	// The description for the Data Factory Linked Service ODBC.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The integration runtime reference to associate with the Data Factory Linked Service ODBC.
	// +kubebuilder:validation:Optional
	IntegrationRuntimeName *string `json:"integrationRuntimeName,omitempty" tf:"integration_runtime_name,omitempty"`

	// A map of parameters to associate with the Data Factory Linked Service ODBC.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

// LinkedServiceOdbcSpec defines the desired state of LinkedServiceOdbc
type LinkedServiceOdbcSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LinkedServiceOdbcParameters `json:"forProvider"`
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
	InitProvider LinkedServiceOdbcInitParameters `json:"initProvider,omitempty"`
}

// LinkedServiceOdbcStatus defines the observed state of LinkedServiceOdbc.
type LinkedServiceOdbcStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LinkedServiceOdbcObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// LinkedServiceOdbc is the Schema for the LinkedServiceOdbcs API. Manages a Linked Service (connection) between a Database and Azure Data Factory through ODBC protocol.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type LinkedServiceOdbc struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.connectionString) || (has(self.initProvider) && has(self.initProvider.connectionString))",message="spec.forProvider.connectionString is a required parameter"
	Spec   LinkedServiceOdbcSpec   `json:"spec"`
	Status LinkedServiceOdbcStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LinkedServiceOdbcList contains a list of LinkedServiceOdbcs
type LinkedServiceOdbcList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LinkedServiceOdbc `json:"items"`
}

// Repository type metadata.
var (
	LinkedServiceOdbc_Kind             = "LinkedServiceOdbc"
	LinkedServiceOdbc_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LinkedServiceOdbc_Kind}.String()
	LinkedServiceOdbc_KindAPIVersion   = LinkedServiceOdbc_Kind + "." + CRDGroupVersion.String()
	LinkedServiceOdbc_GroupVersionKind = CRDGroupVersion.WithKind(LinkedServiceOdbc_Kind)
)

func init() {
	SchemeBuilder.Register(&LinkedServiceOdbc{}, &LinkedServiceOdbcList{})
}
