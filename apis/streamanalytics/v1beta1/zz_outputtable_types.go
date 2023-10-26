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

type OutputTableInitParameters struct {

	// The number of records for a batch operation. Must be between 1 and 100.
	BatchSize *float64 `json:"batchSize,omitempty" tf:"batch_size,omitempty"`

	// A list of the column names to be removed from output event entities.
	ColumnsToRemove []*string `json:"columnsToRemove,omitempty" tf:"columns_to_remove,omitempty"`

	// The name of the output column that contains the partition key.
	PartitionKey *string `json:"partitionKey,omitempty" tf:"partition_key,omitempty"`

	// The name of the output column that contains the row key.
	RowKey *string `json:"rowKey,omitempty" tf:"row_key,omitempty"`
}

type OutputTableObservation struct {

	// The number of records for a batch operation. Must be between 1 and 100.
	BatchSize *float64 `json:"batchSize,omitempty" tf:"batch_size,omitempty"`

	// A list of the column names to be removed from output event entities.
	ColumnsToRemove []*string `json:"columnsToRemove,omitempty" tf:"columns_to_remove,omitempty"`

	// The ID of the Stream Analytics Output Table.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the output column that contains the partition key.
	PartitionKey *string `json:"partitionKey,omitempty" tf:"partition_key,omitempty"`

	// The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// The name of the output column that contains the row key.
	RowKey *string `json:"rowKey,omitempty" tf:"row_key,omitempty"`

	// The name of the Storage Account.
	StorageAccountName *string `json:"storageAccountName,omitempty" tf:"storage_account_name,omitempty"`

	// The name of the Stream Analytics Job. Changing this forces a new resource to be created.
	StreamAnalyticsJobName *string `json:"streamAnalyticsJobName,omitempty" tf:"stream_analytics_job_name,omitempty"`

	// The name of the table where the stream should be output to.
	Table *string `json:"table,omitempty" tf:"table,omitempty"`
}

type OutputTableParameters struct {

	// The number of records for a batch operation. Must be between 1 and 100.
	// +kubebuilder:validation:Optional
	BatchSize *float64 `json:"batchSize,omitempty" tf:"batch_size,omitempty"`

	// A list of the column names to be removed from output event entities.
	// +kubebuilder:validation:Optional
	ColumnsToRemove []*string `json:"columnsToRemove,omitempty" tf:"columns_to_remove,omitempty"`

	// The name of the output column that contains the partition key.
	// +kubebuilder:validation:Optional
	PartitionKey *string `json:"partitionKey,omitempty" tf:"partition_key,omitempty"`

	// The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The name of the output column that contains the row key.
	// +kubebuilder:validation:Optional
	RowKey *string `json:"rowKey,omitempty" tf:"row_key,omitempty"`

	// The Access Key which should be used to connect to this Storage Account.
	// +kubebuilder:validation:Optional
	StorageAccountKeySecretRef v1.SecretKeySelector `json:"storageAccountKeySecretRef" tf:"-"`

	// The name of the Storage Account.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/storage/v1beta1.Account
	// +kubebuilder:validation:Optional
	StorageAccountName *string `json:"storageAccountName,omitempty" tf:"storage_account_name,omitempty"`

	// Reference to a Account in storage to populate storageAccountName.
	// +kubebuilder:validation:Optional
	StorageAccountNameRef *v1.Reference `json:"storageAccountNameRef,omitempty" tf:"-"`

	// Selector for a Account in storage to populate storageAccountName.
	// +kubebuilder:validation:Optional
	StorageAccountNameSelector *v1.Selector `json:"storageAccountNameSelector,omitempty" tf:"-"`

	// The name of the Stream Analytics Job. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	StreamAnalyticsJobName *string `json:"streamAnalyticsJobName" tf:"stream_analytics_job_name,omitempty"`

	// The name of the table where the stream should be output to.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/storage/v1beta1.Table
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("name",false)
	// +kubebuilder:validation:Optional
	Table *string `json:"table,omitempty" tf:"table,omitempty"`

	// Reference to a Table in storage to populate table.
	// +kubebuilder:validation:Optional
	TableRef *v1.Reference `json:"tableRef,omitempty" tf:"-"`

	// Selector for a Table in storage to populate table.
	// +kubebuilder:validation:Optional
	TableSelector *v1.Selector `json:"tableSelector,omitempty" tf:"-"`
}

// OutputTableSpec defines the desired state of OutputTable
type OutputTableSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     OutputTableParameters `json:"forProvider"`
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
	InitProvider OutputTableInitParameters `json:"initProvider,omitempty"`
}

// OutputTableStatus defines the observed state of OutputTable.
type OutputTableStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        OutputTableObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// OutputTable is the Schema for the OutputTables API. Manages a Stream Analytics Output Table.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type OutputTable struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.batchSize) || (has(self.initProvider) && has(self.initProvider.batchSize))",message="spec.forProvider.batchSize is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.partitionKey) || (has(self.initProvider) && has(self.initProvider.partitionKey))",message="spec.forProvider.partitionKey is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.rowKey) || (has(self.initProvider) && has(self.initProvider.rowKey))",message="spec.forProvider.rowKey is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.storageAccountKeySecretRef)",message="spec.forProvider.storageAccountKeySecretRef is a required parameter"
	Spec   OutputTableSpec   `json:"spec"`
	Status OutputTableStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OutputTableList contains a list of OutputTables
type OutputTableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OutputTable `json:"items"`
}

// Repository type metadata.
var (
	OutputTable_Kind             = "OutputTable"
	OutputTable_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: OutputTable_Kind}.String()
	OutputTable_KindAPIVersion   = OutputTable_Kind + "." + CRDGroupVersion.String()
	OutputTable_GroupVersionKind = CRDGroupVersion.WithKind(OutputTable_Kind)
)

func init() {
	SchemeBuilder.Register(&OutputTable{}, &OutputTableList{})
}
