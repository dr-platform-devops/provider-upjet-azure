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

type CloudToDeviceInitParameters struct {

	// The default time to live for cloud-to-device messages, specified as an ISO 8601 timespan duration. This value must be between 1 minute and 48 hours. Defaults to PT1H.
	DefaultTTL *string `json:"defaultTtl,omitempty" tf:"default_ttl,omitempty"`

	// A feedback block as defined below.
	Feedback []FeedbackInitParameters `json:"feedback,omitempty" tf:"feedback,omitempty"`

	// The maximum delivery count for cloud-to-device per-device queues. This value must be between 1 and 100. Defaults to 10.
	MaxDeliveryCount *float64 `json:"maxDeliveryCount,omitempty" tf:"max_delivery_count,omitempty"`
}

type CloudToDeviceObservation struct {

	// The default time to live for cloud-to-device messages, specified as an ISO 8601 timespan duration. This value must be between 1 minute and 48 hours. Defaults to PT1H.
	DefaultTTL *string `json:"defaultTtl,omitempty" tf:"default_ttl,omitempty"`

	// A feedback block as defined below.
	Feedback []FeedbackObservation `json:"feedback,omitempty" tf:"feedback,omitempty"`

	// The maximum delivery count for cloud-to-device per-device queues. This value must be between 1 and 100. Defaults to 10.
	MaxDeliveryCount *float64 `json:"maxDeliveryCount,omitempty" tf:"max_delivery_count,omitempty"`
}

type CloudToDeviceParameters struct {

	// The default time to live for cloud-to-device messages, specified as an ISO 8601 timespan duration. This value must be between 1 minute and 48 hours. Defaults to PT1H.
	// +kubebuilder:validation:Optional
	DefaultTTL *string `json:"defaultTtl,omitempty" tf:"default_ttl,omitempty"`

	// A feedback block as defined below.
	// +kubebuilder:validation:Optional
	Feedback []FeedbackParameters `json:"feedback,omitempty" tf:"feedback,omitempty"`

	// The maximum delivery count for cloud-to-device per-device queues. This value must be between 1 and 100. Defaults to 10.
	// +kubebuilder:validation:Optional
	MaxDeliveryCount *float64 `json:"maxDeliveryCount,omitempty" tf:"max_delivery_count,omitempty"`
}

type EndpointInitParameters struct {
}

type EndpointObservation struct {

	// The type used to authenticate against the endpoint. Possible values are keyBased and identityBased. Defaults to keyBased.
	AuthenticationType *string `json:"authenticationType,omitempty" tf:"authentication_type,omitempty"`

	// Time interval at which blobs are written to storage. Value should be between 60 and 720 seconds. Default value is 300 seconds. This attribute is applicable for endpoint type AzureIotHub.StorageContainer.
	BatchFrequencyInSeconds *float64 `json:"batchFrequencyInSeconds,omitempty" tf:"batch_frequency_in_seconds,omitempty"`

	// The name of storage container in the storage account. This attribute is mandatory for endpoint type AzureIotHub.StorageContainer.
	ContainerName *string `json:"containerName,omitempty" tf:"container_name,omitempty"`

	// Encoding that is used to serialize messages to blobs. Supported values are Avro, AvroDeflate and JSON. Default value is Avro. This attribute is applicable for endpoint type AzureIotHub.StorageContainer. Changing this forces a new resource to be created.
	Encoding *string `json:"encoding,omitempty" tf:"encoding,omitempty"`

	// URI of the Service Bus or Event Hubs Namespace endpoint. This attribute can only be specified and is mandatory when authentication_type is identityBased for endpoint type AzureIotHub.ServiceBusQueue, AzureIotHub.ServiceBusTopic or AzureIotHub.EventHub.
	EndpointURI *string `json:"endpointUri,omitempty" tf:"endpoint_uri,omitempty"`

	// Name of the Service Bus Queue/Topic or Event Hub. This attribute can only be specified and is mandatory when authentication_type is identityBased for endpoint type AzureIotHub.ServiceBusQueue, AzureIotHub.ServiceBusTopic or AzureIotHub.EventHub.
	EntityPath *string `json:"entityPath,omitempty" tf:"entity_path,omitempty"`

	// File name format for the blob. All parameters are mandatory but can be reordered. This attribute is applicable for endpoint type AzureIotHub.StorageContainer. Defaults to {iothub}/{partition}/{YYYY}/{MM}/{DD}/{HH}/{mm}.
	FileNameFormat *string `json:"fileNameFormat,omitempty" tf:"file_name_format,omitempty"`

	// The ID of the User Managed Identity used to authenticate against the endpoint.
	IdentityID *string `json:"identityId,omitempty" tf:"identity_id,omitempty"`

	// Maximum number of bytes for each blob written to storage. Value should be between 10485760(10MB) and 524288000(500MB). Default value is 314572800(300MB). This attribute is applicable for endpoint type AzureIotHub.StorageContainer.
	MaxChunkSizeInBytes *float64 `json:"maxChunkSizeInBytes,omitempty" tf:"max_chunk_size_in_bytes,omitempty"`

	// The name of the endpoint. The name must be unique across endpoint types. The following names are reserved: events, operationsMonitoringEvents, fileNotifications and $default.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The resource group in which the endpoint will be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// The type of the endpoint. Possible values are AzureIotHub.StorageContainer, AzureIotHub.ServiceBusQueue, AzureIotHub.ServiceBusTopic or AzureIotHub.EventHub.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type EndpointParameters struct {
}

type EnrichmentInitParameters struct {
}

type EnrichmentObservation struct {

	// The list of endpoints which will be enriched.
	EndpointNames []*string `json:"endpointNames,omitempty" tf:"endpoint_names,omitempty"`

	// The key of the enrichment.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// The value of the enrichment. Value can be any static string, the name of the IoT Hub sending the message (use $iothubname) or information from the device twin (ex: $twin.tags.latitude)
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type EnrichmentParameters struct {
}

type FallbackRouteInitParameters struct {
}

type FallbackRouteObservation struct {

	// The condition that is evaluated to apply the routing rule. Defaults to true. For grammar, see: https://docs.microsoft.com/azure/iot-hub/iot-hub-devguide-query-language.
	Condition *string `json:"condition,omitempty" tf:"condition,omitempty"`

	// Used to specify whether the fallback route is enabled.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The endpoints to which messages that satisfy the condition are routed. Currently only 1 endpoint is allowed.
	EndpointNames []*string `json:"endpointNames,omitempty" tf:"endpoint_names,omitempty"`

	// The source that the routing rule is to be applied to, such as DeviceMessages. Possible values include: Invalid, DeviceMessages, TwinChangeEvents, DeviceLifecycleEvents, DeviceConnectionStateEvents, DeviceJobLifecycleEvents and DigitalTwinChangeEvents. Defaults to DeviceMessages.
	Source *string `json:"source,omitempty" tf:"source,omitempty"`
}

type FallbackRouteParameters struct {
}

type FeedbackInitParameters struct {

	// The lock duration for the file upload notifications queue, specified as an ISO 8601 timespan duration. This value must be between 5 and 300 seconds. Defaults to PT1M.
	LockDuration *string `json:"lockDuration,omitempty" tf:"lock_duration,omitempty"`

	// The number of times the IoT Hub attempts to deliver a file upload notification message. Defaults to 10.
	MaxDeliveryCount *float64 `json:"maxDeliveryCount,omitempty" tf:"max_delivery_count,omitempty"`

	// The retention time for service-bound feedback messages, specified as an ISO 8601 timespan duration. This value must be between 1 minute and 48 hours. Defaults to PT1H.
	TimeToLive *string `json:"timeToLive,omitempty" tf:"time_to_live,omitempty"`
}

type FeedbackObservation struct {

	// The lock duration for the file upload notifications queue, specified as an ISO 8601 timespan duration. This value must be between 5 and 300 seconds. Defaults to PT1M.
	LockDuration *string `json:"lockDuration,omitempty" tf:"lock_duration,omitempty"`

	// The number of times the IoT Hub attempts to deliver a file upload notification message. Defaults to 10.
	MaxDeliveryCount *float64 `json:"maxDeliveryCount,omitempty" tf:"max_delivery_count,omitempty"`

	// The retention time for service-bound feedback messages, specified as an ISO 8601 timespan duration. This value must be between 1 minute and 48 hours. Defaults to PT1H.
	TimeToLive *string `json:"timeToLive,omitempty" tf:"time_to_live,omitempty"`
}

type FeedbackParameters struct {

	// The lock duration for the file upload notifications queue, specified as an ISO 8601 timespan duration. This value must be between 5 and 300 seconds. Defaults to PT1M.
	// +kubebuilder:validation:Optional
	LockDuration *string `json:"lockDuration,omitempty" tf:"lock_duration,omitempty"`

	// The number of times the IoT Hub attempts to deliver a file upload notification message. Defaults to 10.
	// +kubebuilder:validation:Optional
	MaxDeliveryCount *float64 `json:"maxDeliveryCount,omitempty" tf:"max_delivery_count,omitempty"`

	// The retention time for service-bound feedback messages, specified as an ISO 8601 timespan duration. This value must be between 1 minute and 48 hours. Defaults to PT1H.
	// +kubebuilder:validation:Optional
	TimeToLive *string `json:"timeToLive,omitempty" tf:"time_to_live,omitempty"`
}

type FileUploadInitParameters struct {

	// The type used to authenticate against the storage account. Possible values are keyBased and identityBased. Defaults to keyBased.
	AuthenticationType *string `json:"authenticationType,omitempty" tf:"authentication_type,omitempty"`

	// The connection string for the Azure Storage account to which files are uploaded.
	ConnectionStringSecretRef v1.SecretKeySelector `json:"connectionStringSecretRef" tf:"-"`

	// The name of the root container where the files should be uploaded to. The container need not exist but should be creatable using the connection_string specified.
	ContainerName *string `json:"containerName,omitempty" tf:"container_name,omitempty"`

	// The period of time for which a file upload notification message is available to consume before it expires, specified as an ISO 8601 timespan duration. This value must be between 1 minute and 48 hours. Defaults to PT1H.
	DefaultTTL *string `json:"defaultTtl,omitempty" tf:"default_ttl,omitempty"`

	// The ID of the User Managed Identity used to authenticate against the storage account.
	IdentityID *string `json:"identityId,omitempty" tf:"identity_id,omitempty"`

	// The lock duration for the file upload notifications queue, specified as an ISO 8601 timespan duration. This value must be between 5 and 300 seconds. Defaults to PT1M.
	LockDuration *string `json:"lockDuration,omitempty" tf:"lock_duration,omitempty"`

	// The number of times the IoT Hub attempts to deliver a file upload notification message. Defaults to 10.
	MaxDeliveryCount *float64 `json:"maxDeliveryCount,omitempty" tf:"max_delivery_count,omitempty"`

	// Used to specify whether file notifications are sent to IoT Hub on upload. Defaults to false.
	Notifications *bool `json:"notifications,omitempty" tf:"notifications,omitempty"`

	// The period of time for which the SAS URI generated by IoT Hub for file upload is valid, specified as an ISO 8601 timespan duration. This value must be between 1 minute and 24 hours. Defaults to PT1H.
	SASTTL *string `json:"sasTtl,omitempty" tf:"sas_ttl,omitempty"`
}

type FileUploadObservation struct {

	// The type used to authenticate against the storage account. Possible values are keyBased and identityBased. Defaults to keyBased.
	AuthenticationType *string `json:"authenticationType,omitempty" tf:"authentication_type,omitempty"`

	// The name of the root container where the files should be uploaded to. The container need not exist but should be creatable using the connection_string specified.
	ContainerName *string `json:"containerName,omitempty" tf:"container_name,omitempty"`

	// The period of time for which a file upload notification message is available to consume before it expires, specified as an ISO 8601 timespan duration. This value must be between 1 minute and 48 hours. Defaults to PT1H.
	DefaultTTL *string `json:"defaultTtl,omitempty" tf:"default_ttl,omitempty"`

	// The ID of the User Managed Identity used to authenticate against the storage account.
	IdentityID *string `json:"identityId,omitempty" tf:"identity_id,omitempty"`

	// The lock duration for the file upload notifications queue, specified as an ISO 8601 timespan duration. This value must be between 5 and 300 seconds. Defaults to PT1M.
	LockDuration *string `json:"lockDuration,omitempty" tf:"lock_duration,omitempty"`

	// The number of times the IoT Hub attempts to deliver a file upload notification message. Defaults to 10.
	MaxDeliveryCount *float64 `json:"maxDeliveryCount,omitempty" tf:"max_delivery_count,omitempty"`

	// Used to specify whether file notifications are sent to IoT Hub on upload. Defaults to false.
	Notifications *bool `json:"notifications,omitempty" tf:"notifications,omitempty"`

	// The period of time for which the SAS URI generated by IoT Hub for file upload is valid, specified as an ISO 8601 timespan duration. This value must be between 1 minute and 24 hours. Defaults to PT1H.
	SASTTL *string `json:"sasTtl,omitempty" tf:"sas_ttl,omitempty"`
}

type FileUploadParameters struct {

	// The type used to authenticate against the storage account. Possible values are keyBased and identityBased. Defaults to keyBased.
	// +kubebuilder:validation:Optional
	AuthenticationType *string `json:"authenticationType,omitempty" tf:"authentication_type,omitempty"`

	// The connection string for the Azure Storage account to which files are uploaded.
	// +kubebuilder:validation:Optional
	ConnectionStringSecretRef v1.SecretKeySelector `json:"connectionStringSecretRef" tf:"-"`

	// The name of the root container where the files should be uploaded to. The container need not exist but should be creatable using the connection_string specified.
	// +kubebuilder:validation:Optional
	ContainerName *string `json:"containerName" tf:"container_name,omitempty"`

	// The period of time for which a file upload notification message is available to consume before it expires, specified as an ISO 8601 timespan duration. This value must be between 1 minute and 48 hours. Defaults to PT1H.
	// +kubebuilder:validation:Optional
	DefaultTTL *string `json:"defaultTtl,omitempty" tf:"default_ttl,omitempty"`

	// The ID of the User Managed Identity used to authenticate against the storage account.
	// +kubebuilder:validation:Optional
	IdentityID *string `json:"identityId,omitempty" tf:"identity_id,omitempty"`

	// The lock duration for the file upload notifications queue, specified as an ISO 8601 timespan duration. This value must be between 5 and 300 seconds. Defaults to PT1M.
	// +kubebuilder:validation:Optional
	LockDuration *string `json:"lockDuration,omitempty" tf:"lock_duration,omitempty"`

	// The number of times the IoT Hub attempts to deliver a file upload notification message. Defaults to 10.
	// +kubebuilder:validation:Optional
	MaxDeliveryCount *float64 `json:"maxDeliveryCount,omitempty" tf:"max_delivery_count,omitempty"`

	// Used to specify whether file notifications are sent to IoT Hub on upload. Defaults to false.
	// +kubebuilder:validation:Optional
	Notifications *bool `json:"notifications,omitempty" tf:"notifications,omitempty"`

	// The period of time for which the SAS URI generated by IoT Hub for file upload is valid, specified as an ISO 8601 timespan duration. This value must be between 1 minute and 24 hours. Defaults to PT1H.
	// +kubebuilder:validation:Optional
	SASTTL *string `json:"sasTtl,omitempty" tf:"sas_ttl,omitempty"`
}

type IOTHubInitParameters struct {

	// A cloud_to_device block as defined below.
	CloudToDevice *CloudToDeviceInitParameters `json:"cloudToDevice,omitempty" tf:"cloud_to_device,omitempty"`

	// The number of device-to-cloud partitions used by backing event hubs. Must be between 2 and 128.
	EventHubPartitionCount *float64 `json:"eventHubPartitionCount,omitempty" tf:"event_hub_partition_count,omitempty"`

	// The event hub retention to use in days. Must be between 1 and 7.
	EventHubRetentionInDays *float64 `json:"eventHubRetentionInDays,omitempty" tf:"event_hub_retention_in_days,omitempty"`

	// A file_upload block as defined below.
	FileUpload *FileUploadInitParameters `json:"fileUpload,omitempty" tf:"file_upload,omitempty"`

	// An identity block as defined below.
	Identity *IdentityInitParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// If false, SAS tokens with Iot hub scoped SAS keys cannot be used for authentication. Defaults to true.
	LocalAuthenticationEnabled *bool `json:"localAuthenticationEnabled,omitempty" tf:"local_authentication_enabled,omitempty"`

	// Specifies the supported Azure location where the resource has to be created. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Specifies the minimum TLS version to support for this hub. The only valid value is 1.2. Changing this forces a new resource to be created.
	MinTLSVersion *string `json:"minTlsVersion,omitempty" tf:"min_tls_version,omitempty"`

	// A network_rule_set block as defined below.
	NetworkRuleSet []NetworkRuleSetInitParameters `json:"networkRuleSet,omitempty" tf:"network_rule_set,omitempty"`

	// Is the IotHub resource accessible from a public network?
	PublicNetworkAccessEnabled *bool `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled,omitempty"`

	// A sku block as defined below.
	Sku *SkuInitParameters `json:"sku,omitempty" tf:"sku,omitempty"`

	// A mapping of tags to assign to the resource.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type IOTHubObservation struct {

	// A cloud_to_device block as defined below.
	CloudToDevice *CloudToDeviceObservation `json:"cloudToDevice,omitempty" tf:"cloud_to_device,omitempty"`

	// An endpoint block as defined below.
	Endpoint []EndpointObservation `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	// A enrichment block as defined below.
	Enrichment []EnrichmentObservation `json:"enrichment,omitempty" tf:"enrichment,omitempty"`

	// The EventHub compatible endpoint for events data
	EventHubEventsEndpoint *string `json:"eventHubEventsEndpoint,omitempty" tf:"event_hub_events_endpoint,omitempty"`

	// The EventHub namespace for events data
	EventHubEventsNamespace *string `json:"eventHubEventsNamespace,omitempty" tf:"event_hub_events_namespace,omitempty"`

	// The EventHub compatible path for events data
	EventHubEventsPath *string `json:"eventHubEventsPath,omitempty" tf:"event_hub_events_path,omitempty"`

	// The EventHub compatible endpoint for operational data
	EventHubOperationsEndpoint *string `json:"eventHubOperationsEndpoint,omitempty" tf:"event_hub_operations_endpoint,omitempty"`

	// The EventHub compatible path for operational data
	EventHubOperationsPath *string `json:"eventHubOperationsPath,omitempty" tf:"event_hub_operations_path,omitempty"`

	// The number of device-to-cloud partitions used by backing event hubs. Must be between 2 and 128.
	EventHubPartitionCount *float64 `json:"eventHubPartitionCount,omitempty" tf:"event_hub_partition_count,omitempty"`

	// The event hub retention to use in days. Must be between 1 and 7.
	EventHubRetentionInDays *float64 `json:"eventHubRetentionInDays,omitempty" tf:"event_hub_retention_in_days,omitempty"`

	// A fallback_route block as defined below. If the fallback route is enabled, messages that don't match any of the supplied routes are automatically sent to this route. Defaults to messages/events.
	FallbackRoute *FallbackRouteObservation `json:"fallbackRoute,omitempty" tf:"fallback_route,omitempty"`

	// A file_upload block as defined below.
	FileUpload *FileUploadObservation `json:"fileUpload,omitempty" tf:"file_upload,omitempty"`

	// The hostname of the IotHub Resource.
	HostName *string `json:"hostname,omitempty" tf:"hostname,omitempty"`

	// The ID of the IoTHub.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// An identity block as defined below.
	Identity *IdentityObservation `json:"identity,omitempty" tf:"identity,omitempty"`

	// If false, SAS tokens with Iot hub scoped SAS keys cannot be used for authentication. Defaults to true.
	LocalAuthenticationEnabled *bool `json:"localAuthenticationEnabled,omitempty" tf:"local_authentication_enabled,omitempty"`

	// Specifies the supported Azure location where the resource has to be created. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Specifies the minimum TLS version to support for this hub. The only valid value is 1.2. Changing this forces a new resource to be created.
	MinTLSVersion *string `json:"minTlsVersion,omitempty" tf:"min_tls_version,omitempty"`

	// A network_rule_set block as defined below.
	NetworkRuleSet []NetworkRuleSetObservation `json:"networkRuleSet,omitempty" tf:"network_rule_set,omitempty"`

	// Is the IotHub resource accessible from a public network?
	PublicNetworkAccessEnabled *bool `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled,omitempty"`

	// The name of the resource group under which the IotHub resource has to be created. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// A route block as defined below.
	Route []RouteObservation `json:"route,omitempty" tf:"route,omitempty"`

	// One or more shared_access_policy blocks as defined below.
	SharedAccessPolicy []SharedAccessPolicyObservation `json:"sharedAccessPolicy,omitempty" tf:"shared_access_policy,omitempty"`

	// A sku block as defined below.
	Sku *SkuObservation `json:"sku,omitempty" tf:"sku,omitempty"`

	// A mapping of tags to assign to the resource.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies the type of Managed Service Identity that should be configured on this IoT Hub. Possible values are SystemAssigned, UserAssigned, SystemAssigned, UserAssigned (to enable both).
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type IOTHubParameters struct {

	// A cloud_to_device block as defined below.
	// +kubebuilder:validation:Optional
	CloudToDevice *CloudToDeviceParameters `json:"cloudToDevice,omitempty" tf:"cloud_to_device,omitempty"`

	// The number of device-to-cloud partitions used by backing event hubs. Must be between 2 and 128.
	// +kubebuilder:validation:Optional
	EventHubPartitionCount *float64 `json:"eventHubPartitionCount,omitempty" tf:"event_hub_partition_count,omitempty"`

	// The event hub retention to use in days. Must be between 1 and 7.
	// +kubebuilder:validation:Optional
	EventHubRetentionInDays *float64 `json:"eventHubRetentionInDays,omitempty" tf:"event_hub_retention_in_days,omitempty"`

	// A file_upload block as defined below.
	// +kubebuilder:validation:Optional
	FileUpload *FileUploadParameters `json:"fileUpload,omitempty" tf:"file_upload,omitempty"`

	// An identity block as defined below.
	// +kubebuilder:validation:Optional
	Identity *IdentityParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// If false, SAS tokens with Iot hub scoped SAS keys cannot be used for authentication. Defaults to true.
	// +kubebuilder:validation:Optional
	LocalAuthenticationEnabled *bool `json:"localAuthenticationEnabled,omitempty" tf:"local_authentication_enabled,omitempty"`

	// Specifies the supported Azure location where the resource has to be created. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Specifies the minimum TLS version to support for this hub. The only valid value is 1.2. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	MinTLSVersion *string `json:"minTlsVersion,omitempty" tf:"min_tls_version,omitempty"`

	// A network_rule_set block as defined below.
	// +kubebuilder:validation:Optional
	NetworkRuleSet []NetworkRuleSetParameters `json:"networkRuleSet,omitempty" tf:"network_rule_set,omitempty"`

	// Is the IotHub resource accessible from a public network?
	// +kubebuilder:validation:Optional
	PublicNetworkAccessEnabled *bool `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled,omitempty"`

	// The name of the resource group under which the IotHub resource has to be created. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// A sku block as defined below.
	// +kubebuilder:validation:Optional
	Sku *SkuParameters `json:"sku,omitempty" tf:"sku,omitempty"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type IPRuleInitParameters struct {

	// The desired action for requests captured by this rule. Possible values are Allow. Defaults to Allow.
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// The IP address range in CIDR notation for the IP rule.
	IPMask *string `json:"ipMask,omitempty" tf:"ip_mask,omitempty"`

	// The name of the sku. Possible values are B1, B2, B3, F1, S1, S2, and S3.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type IPRuleObservation struct {

	// The desired action for requests captured by this rule. Possible values are Allow. Defaults to Allow.
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// The IP address range in CIDR notation for the IP rule.
	IPMask *string `json:"ipMask,omitempty" tf:"ip_mask,omitempty"`

	// The name of the sku. Possible values are B1, B2, B3, F1, S1, S2, and S3.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type IPRuleParameters struct {

	// The desired action for requests captured by this rule. Possible values are Allow. Defaults to Allow.
	// +kubebuilder:validation:Optional
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// The IP address range in CIDR notation for the IP rule.
	// +kubebuilder:validation:Optional
	IPMask *string `json:"ipMask" tf:"ip_mask,omitempty"`

	// The name of the sku. Possible values are B1, B2, B3, F1, S1, S2, and S3.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`
}

type IdentityInitParameters struct {

	// Specifies a list of User Assigned Managed Identity IDs to be assigned to this IoT Hub.
	// +listType=set
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// Specifies the type of Managed Service Identity that should be configured on this IoT Hub. Possible values are SystemAssigned, UserAssigned, SystemAssigned, UserAssigned (to enable both).
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type IdentityObservation struct {

	// Specifies a list of User Assigned Managed Identity IDs to be assigned to this IoT Hub.
	// +listType=set
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// The Principal ID associated with this Managed Service Identity.
	PrincipalID *string `json:"principalId,omitempty" tf:"principal_id,omitempty"`

	// The Tenant ID associated with this Managed Service Identity.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	// Specifies the type of Managed Service Identity that should be configured on this IoT Hub. Possible values are SystemAssigned, UserAssigned, SystemAssigned, UserAssigned (to enable both).
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type IdentityParameters struct {

	// Specifies a list of User Assigned Managed Identity IDs to be assigned to this IoT Hub.
	// +kubebuilder:validation:Optional
	// +listType=set
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// Specifies the type of Managed Service Identity that should be configured on this IoT Hub. Possible values are SystemAssigned, UserAssigned, SystemAssigned, UserAssigned (to enable both).
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

type NetworkRuleSetInitParameters struct {

	// Determines if Network Rule Set is also applied to the BuiltIn EventHub EndPoint of the IotHub. Defaults to false.
	ApplyToBuiltinEventHubEndpoint *bool `json:"applyToBuiltinEventhubEndpoint,omitempty" tf:"apply_to_builtin_eventhub_endpoint,omitempty"`

	// Default Action for Network Rule Set. Possible values are Deny, Allow. Defaults to Deny.
	DefaultAction *string `json:"defaultAction,omitempty" tf:"default_action,omitempty"`

	// One or more ip_rule blocks as defined below.
	IPRule []IPRuleInitParameters `json:"ipRule,omitempty" tf:"ip_rule,omitempty"`
}

type NetworkRuleSetObservation struct {

	// Determines if Network Rule Set is also applied to the BuiltIn EventHub EndPoint of the IotHub. Defaults to false.
	ApplyToBuiltinEventHubEndpoint *bool `json:"applyToBuiltinEventhubEndpoint,omitempty" tf:"apply_to_builtin_eventhub_endpoint,omitempty"`

	// Default Action for Network Rule Set. Possible values are Deny, Allow. Defaults to Deny.
	DefaultAction *string `json:"defaultAction,omitempty" tf:"default_action,omitempty"`

	// One or more ip_rule blocks as defined below.
	IPRule []IPRuleObservation `json:"ipRule,omitempty" tf:"ip_rule,omitempty"`
}

type NetworkRuleSetParameters struct {

	// Determines if Network Rule Set is also applied to the BuiltIn EventHub EndPoint of the IotHub. Defaults to false.
	// +kubebuilder:validation:Optional
	ApplyToBuiltinEventHubEndpoint *bool `json:"applyToBuiltinEventhubEndpoint,omitempty" tf:"apply_to_builtin_eventhub_endpoint,omitempty"`

	// Default Action for Network Rule Set. Possible values are Deny, Allow. Defaults to Deny.
	// +kubebuilder:validation:Optional
	DefaultAction *string `json:"defaultAction,omitempty" tf:"default_action,omitempty"`

	// One or more ip_rule blocks as defined below.
	// +kubebuilder:validation:Optional
	IPRule []IPRuleParameters `json:"ipRule,omitempty" tf:"ip_rule,omitempty"`
}

type RouteInitParameters struct {
}

type RouteObservation struct {

	// The condition that is evaluated to apply the routing rule. Defaults to true. For grammar, see: https://docs.microsoft.com/azure/iot-hub/iot-hub-devguide-query-language.
	Condition *string `json:"condition,omitempty" tf:"condition,omitempty"`

	// Used to specify whether a route is enabled.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The list of endpoints to which messages that satisfy the condition are routed.
	EndpointNames []*string `json:"endpointNames,omitempty" tf:"endpoint_names,omitempty"`

	// The name of the route.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The source that the routing rule is to be applied to, such as DeviceMessages. Possible values include: Invalid, DeviceMessages, TwinChangeEvents, DeviceLifecycleEvents, DeviceConnectionStateEvents, DeviceJobLifecycleEvents and DigitalTwinChangeEvents.
	Source *string `json:"source,omitempty" tf:"source,omitempty"`
}

type RouteParameters struct {
}

type SharedAccessPolicyInitParameters struct {
}

type SharedAccessPolicyObservation struct {

	// The name of the shared access policy.
	KeyName *string `json:"keyName,omitempty" tf:"key_name,omitempty"`

	// The permissions assigned to the shared access policy.
	Permissions *string `json:"permissions,omitempty" tf:"permissions,omitempty"`
}

type SharedAccessPolicyParameters struct {
}

type SkuInitParameters struct {

	// The number of provisioned IoT Hub units.
	Capacity *float64 `json:"capacity,omitempty" tf:"capacity,omitempty"`

	// The name of the sku. Possible values are B1, B2, B3, F1, S1, S2, and S3.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type SkuObservation struct {

	// The number of provisioned IoT Hub units.
	Capacity *float64 `json:"capacity,omitempty" tf:"capacity,omitempty"`

	// The name of the sku. Possible values are B1, B2, B3, F1, S1, S2, and S3.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type SkuParameters struct {

	// The number of provisioned IoT Hub units.
	// +kubebuilder:validation:Optional
	Capacity *float64 `json:"capacity" tf:"capacity,omitempty"`

	// The name of the sku. Possible values are B1, B2, B3, F1, S1, S2, and S3.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`
}

// IOTHubSpec defines the desired state of IOTHub
type IOTHubSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IOTHubParameters `json:"forProvider"`
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
	InitProvider IOTHubInitParameters `json:"initProvider,omitempty"`
}

// IOTHubStatus defines the observed state of IOTHub.
type IOTHubStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IOTHubObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// IOTHub is the Schema for the IOTHubs API. Manages an IotHub
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type IOTHub struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || (has(self.initProvider) && has(self.initProvider.location))",message="spec.forProvider.location is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.sku) || (has(self.initProvider) && has(self.initProvider.sku))",message="spec.forProvider.sku is a required parameter"
	Spec   IOTHubSpec   `json:"spec"`
	Status IOTHubStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IOTHubList contains a list of IOTHubs
type IOTHubList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IOTHub `json:"items"`
}

// Repository type metadata.
var (
	IOTHub_Kind             = "IOTHub"
	IOTHub_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: IOTHub_Kind}.String()
	IOTHub_KindAPIVersion   = IOTHub_Kind + "." + CRDGroupVersion.String()
	IOTHub_GroupVersionKind = CRDGroupVersion.WithKind(IOTHub_Kind)
)

func init() {
	SchemeBuilder.Register(&IOTHub{}, &IOTHubList{})
}
