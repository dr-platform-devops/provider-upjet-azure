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

type LinkedServiceSFTPInitParameters struct {

	// A map of additional properties to associate with the Data Factory Linked Service.
	// +mapType=granular
	AdditionalProperties map[string]*string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// List of tags that can be used for describing the Data Factory Linked Service.
	Annotations []*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// The type of authentication used to connect to the web table source. Valid options are Anonymous, Basic and ClientCertificate.
	AuthenticationType *string `json:"authenticationType,omitempty" tf:"authentication_type,omitempty"`

	// The description for the Data Factory Linked Service.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The SFTP server hostname.
	Host *string `json:"host,omitempty" tf:"host,omitempty"`

	// The host key fingerprint of the SFTP server.
	HostKeyFingerprint *string `json:"hostKeyFingerprint,omitempty" tf:"host_key_fingerprint,omitempty"`

	// The integration runtime reference to associate with the Data Factory Linked Service.
	IntegrationRuntimeName *string `json:"integrationRuntimeName,omitempty" tf:"integration_runtime_name,omitempty"`

	// A map of parameters to associate with the Data Factory Linked Service.
	// +mapType=granular
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// Password to logon to the SFTP Server for Basic Authentication.
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef" tf:"-"`

	// The TCP port number that the SFTP server uses to listen for client connection. Default value is 22.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Whether to validate host key fingerprint while connecting. If set to false, host_key_fingerprint must also be set.
	SkipHostKeyValidation *bool `json:"skipHostKeyValidation,omitempty" tf:"skip_host_key_validation,omitempty"`

	// The username used to log on to the SFTP server.
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type LinkedServiceSFTPObservation struct {

	// A map of additional properties to associate with the Data Factory Linked Service.
	// +mapType=granular
	AdditionalProperties map[string]*string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// List of tags that can be used for describing the Data Factory Linked Service.
	Annotations []*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// The type of authentication used to connect to the web table source. Valid options are Anonymous, Basic and ClientCertificate.
	AuthenticationType *string `json:"authenticationType,omitempty" tf:"authentication_type,omitempty"`

	// The Data Factory ID in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryID *string `json:"dataFactoryId,omitempty" tf:"data_factory_id,omitempty"`

	// The description for the Data Factory Linked Service.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The SFTP server hostname.
	Host *string `json:"host,omitempty" tf:"host,omitempty"`

	// The host key fingerprint of the SFTP server.
	HostKeyFingerprint *string `json:"hostKeyFingerprint,omitempty" tf:"host_key_fingerprint,omitempty"`

	// The ID of the Data Factory Linked Service.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The integration runtime reference to associate with the Data Factory Linked Service.
	IntegrationRuntimeName *string `json:"integrationRuntimeName,omitempty" tf:"integration_runtime_name,omitempty"`

	// A map of parameters to associate with the Data Factory Linked Service.
	// +mapType=granular
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// The TCP port number that the SFTP server uses to listen for client connection. Default value is 22.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Whether to validate host key fingerprint while connecting. If set to false, host_key_fingerprint must also be set.
	SkipHostKeyValidation *bool `json:"skipHostKeyValidation,omitempty" tf:"skip_host_key_validation,omitempty"`

	// The username used to log on to the SFTP server.
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type LinkedServiceSFTPParameters struct {

	// A map of additional properties to associate with the Data Factory Linked Service.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	AdditionalProperties map[string]*string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// List of tags that can be used for describing the Data Factory Linked Service.
	// +kubebuilder:validation:Optional
	Annotations []*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// The type of authentication used to connect to the web table source. Valid options are Anonymous, Basic and ClientCertificate.
	// +kubebuilder:validation:Optional
	AuthenticationType *string `json:"authenticationType,omitempty" tf:"authentication_type,omitempty"`

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

	// The description for the Data Factory Linked Service.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The SFTP server hostname.
	// +kubebuilder:validation:Optional
	Host *string `json:"host,omitempty" tf:"host,omitempty"`

	// The host key fingerprint of the SFTP server.
	// +kubebuilder:validation:Optional
	HostKeyFingerprint *string `json:"hostKeyFingerprint,omitempty" tf:"host_key_fingerprint,omitempty"`

	// The integration runtime reference to associate with the Data Factory Linked Service.
	// +kubebuilder:validation:Optional
	IntegrationRuntimeName *string `json:"integrationRuntimeName,omitempty" tf:"integration_runtime_name,omitempty"`

	// A map of parameters to associate with the Data Factory Linked Service.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// Password to logon to the SFTP Server for Basic Authentication.
	// +kubebuilder:validation:Optional
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef" tf:"-"`

	// The TCP port number that the SFTP server uses to listen for client connection. Default value is 22.
	// +kubebuilder:validation:Optional
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Whether to validate host key fingerprint while connecting. If set to false, host_key_fingerprint must also be set.
	// +kubebuilder:validation:Optional
	SkipHostKeyValidation *bool `json:"skipHostKeyValidation,omitempty" tf:"skip_host_key_validation,omitempty"`

	// The username used to log on to the SFTP server.
	// +kubebuilder:validation:Optional
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

// LinkedServiceSFTPSpec defines the desired state of LinkedServiceSFTP
type LinkedServiceSFTPSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LinkedServiceSFTPParameters `json:"forProvider"`
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
	InitProvider LinkedServiceSFTPInitParameters `json:"initProvider,omitempty"`
}

// LinkedServiceSFTPStatus defines the observed state of LinkedServiceSFTP.
type LinkedServiceSFTPStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LinkedServiceSFTPObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// LinkedServiceSFTP is the Schema for the LinkedServiceSFTPs API. Manages a Linked Service (connection) between an SFTP Server and Azure Data Factory.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type LinkedServiceSFTP struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.authenticationType) || (has(self.initProvider) && has(self.initProvider.authenticationType))",message="spec.forProvider.authenticationType is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.host) || (has(self.initProvider) && has(self.initProvider.host))",message="spec.forProvider.host is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.passwordSecretRef)",message="spec.forProvider.passwordSecretRef is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.port) || (has(self.initProvider) && has(self.initProvider.port))",message="spec.forProvider.port is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.username) || (has(self.initProvider) && has(self.initProvider.username))",message="spec.forProvider.username is a required parameter"
	Spec   LinkedServiceSFTPSpec   `json:"spec"`
	Status LinkedServiceSFTPStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LinkedServiceSFTPList contains a list of LinkedServiceSFTPs
type LinkedServiceSFTPList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LinkedServiceSFTP `json:"items"`
}

// Repository type metadata.
var (
	LinkedServiceSFTP_Kind             = "LinkedServiceSFTP"
	LinkedServiceSFTP_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LinkedServiceSFTP_Kind}.String()
	LinkedServiceSFTP_KindAPIVersion   = LinkedServiceSFTP_Kind + "." + CRDGroupVersion.String()
	LinkedServiceSFTP_GroupVersionKind = CRDGroupVersion.WithKind(LinkedServiceSFTP_Kind)
)

func init() {
	SchemeBuilder.Register(&LinkedServiceSFTP{}, &LinkedServiceSFTPList{})
}
