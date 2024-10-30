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

type LoadBalancerRuleInitParameters struct {

	// A list of reference to a Backend Address Pool over which this Load Balancing Rule operates.
	BackendAddressPoolIds []*string `json:"backendAddressPoolIds,omitempty" tf:"backend_address_pool_ids,omitempty"`

	// The port used for internal connections on the endpoint. Possible values range between 0 and 65535, inclusive. A port of 0 means "Any Port".
	BackendPort *float64 `json:"backendPort,omitempty" tf:"backend_port,omitempty"`

	// Is snat enabled for this Load Balancer Rule? Default false.
	DisableOutboundSnat *bool `json:"disableOutboundSnat,omitempty" tf:"disable_outbound_snat,omitempty"`

	// Are the Floating IPs enabled for this Load Balancer Rule? A "floating” IP is reassigned to a secondary server in case the primary server fails. Required to configure a SQL AlwaysOn Availability Group. Defaults to false.
	EnableFloatingIP *bool `json:"enableFloatingIp,omitempty" tf:"enable_floating_ip,omitempty"`

	// Is TCP Reset enabled for this Load Balancer Rule?
	EnableTCPReset *bool `json:"enableTcpReset,omitempty" tf:"enable_tcp_reset,omitempty"`

	// The name of the frontend IP configuration to which the rule is associated.
	FrontendIPConfigurationName *string `json:"frontendIpConfigurationName,omitempty" tf:"frontend_ip_configuration_name,omitempty"`

	// The port for the external endpoint. Port numbers for each Rule must be unique within the Load Balancer. Possible values range between 0 and 65534, inclusive. A port of 0 means "Any Port".
	FrontendPort *float64 `json:"frontendPort,omitempty" tf:"frontend_port,omitempty"`

	// Specifies the idle timeout in minutes for TCP connections. Valid values are between 4 and 100 minutes. Defaults to 4 minutes.
	IdleTimeoutInMinutes *float64 `json:"idleTimeoutInMinutes,omitempty" tf:"idle_timeout_in_minutes,omitempty"`

	// Specifies the load balancing distribution type to be used by the Load Balancer. Possible values are: Default – The load balancer is configured to use a 5 tuple hash to map traffic to available servers. SourceIP – The load balancer is configured to use a 2 tuple hash to map traffic to available servers. SourceIPProtocol – The load balancer is configured to use a 3 tuple hash to map traffic to available servers. Also known as Session Persistence, where in the Azure portal the options are called None, Client IP and Client IP and Protocol respectively. Defaults to Default.
	LoadDistribution *string `json:"loadDistribution,omitempty" tf:"load_distribution,omitempty"`

	// A reference to a Probe used by this Load Balancing Rule.
	ProbeID *string `json:"probeId,omitempty" tf:"probe_id,omitempty"`

	// The transport protocol for the external endpoint. Possible values are Tcp, Udp or All.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`
}

type LoadBalancerRuleObservation struct {

	// A list of reference to a Backend Address Pool over which this Load Balancing Rule operates.
	BackendAddressPoolIds []*string `json:"backendAddressPoolIds,omitempty" tf:"backend_address_pool_ids,omitempty"`

	// The port used for internal connections on the endpoint. Possible values range between 0 and 65535, inclusive. A port of 0 means "Any Port".
	BackendPort *float64 `json:"backendPort,omitempty" tf:"backend_port,omitempty"`

	// Is snat enabled for this Load Balancer Rule? Default false.
	DisableOutboundSnat *bool `json:"disableOutboundSnat,omitempty" tf:"disable_outbound_snat,omitempty"`

	// Are the Floating IPs enabled for this Load Balancer Rule? A "floating” IP is reassigned to a secondary server in case the primary server fails. Required to configure a SQL AlwaysOn Availability Group. Defaults to false.
	EnableFloatingIP *bool `json:"enableFloatingIp,omitempty" tf:"enable_floating_ip,omitempty"`

	// Is TCP Reset enabled for this Load Balancer Rule?
	EnableTCPReset *bool `json:"enableTcpReset,omitempty" tf:"enable_tcp_reset,omitempty"`

	// The ID of the Load Balancer Rule.
	FrontendIPConfigurationID *string `json:"frontendIpConfigurationId,omitempty" tf:"frontend_ip_configuration_id,omitempty"`

	// The name of the frontend IP configuration to which the rule is associated.
	FrontendIPConfigurationName *string `json:"frontendIpConfigurationName,omitempty" tf:"frontend_ip_configuration_name,omitempty"`

	// The port for the external endpoint. Port numbers for each Rule must be unique within the Load Balancer. Possible values range between 0 and 65534, inclusive. A port of 0 means "Any Port".
	FrontendPort *float64 `json:"frontendPort,omitempty" tf:"frontend_port,omitempty"`

	// The ID of the Load Balancer Rule.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the idle timeout in minutes for TCP connections. Valid values are between 4 and 100 minutes. Defaults to 4 minutes.
	IdleTimeoutInMinutes *float64 `json:"idleTimeoutInMinutes,omitempty" tf:"idle_timeout_in_minutes,omitempty"`

	// Specifies the load balancing distribution type to be used by the Load Balancer. Possible values are: Default – The load balancer is configured to use a 5 tuple hash to map traffic to available servers. SourceIP – The load balancer is configured to use a 2 tuple hash to map traffic to available servers. SourceIPProtocol – The load balancer is configured to use a 3 tuple hash to map traffic to available servers. Also known as Session Persistence, where in the Azure portal the options are called None, Client IP and Client IP and Protocol respectively. Defaults to Default.
	LoadDistribution *string `json:"loadDistribution,omitempty" tf:"load_distribution,omitempty"`

	// The ID of the Load Balancer in which to create the Rule. Changing this forces a new resource to be created.
	LoadbalancerID *string `json:"loadbalancerId,omitempty" tf:"loadbalancer_id,omitempty"`

	// A reference to a Probe used by this Load Balancing Rule.
	ProbeID *string `json:"probeId,omitempty" tf:"probe_id,omitempty"`

	// The transport protocol for the external endpoint. Possible values are Tcp, Udp or All.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`
}

type LoadBalancerRuleParameters struct {

	// A list of reference to a Backend Address Pool over which this Load Balancing Rule operates.
	// +kubebuilder:validation:Optional
	BackendAddressPoolIds []*string `json:"backendAddressPoolIds,omitempty" tf:"backend_address_pool_ids,omitempty"`

	// The port used for internal connections on the endpoint. Possible values range between 0 and 65535, inclusive. A port of 0 means "Any Port".
	// +kubebuilder:validation:Optional
	BackendPort *float64 `json:"backendPort,omitempty" tf:"backend_port,omitempty"`

	// Is snat enabled for this Load Balancer Rule? Default false.
	// +kubebuilder:validation:Optional
	DisableOutboundSnat *bool `json:"disableOutboundSnat,omitempty" tf:"disable_outbound_snat,omitempty"`

	// Are the Floating IPs enabled for this Load Balancer Rule? A "floating” IP is reassigned to a secondary server in case the primary server fails. Required to configure a SQL AlwaysOn Availability Group. Defaults to false.
	// +kubebuilder:validation:Optional
	EnableFloatingIP *bool `json:"enableFloatingIp,omitempty" tf:"enable_floating_ip,omitempty"`

	// Is TCP Reset enabled for this Load Balancer Rule?
	// +kubebuilder:validation:Optional
	EnableTCPReset *bool `json:"enableTcpReset,omitempty" tf:"enable_tcp_reset,omitempty"`

	// The name of the frontend IP configuration to which the rule is associated.
	// +kubebuilder:validation:Optional
	FrontendIPConfigurationName *string `json:"frontendIpConfigurationName,omitempty" tf:"frontend_ip_configuration_name,omitempty"`

	// The port for the external endpoint. Port numbers for each Rule must be unique within the Load Balancer. Possible values range between 0 and 65534, inclusive. A port of 0 means "Any Port".
	// +kubebuilder:validation:Optional
	FrontendPort *float64 `json:"frontendPort,omitempty" tf:"frontend_port,omitempty"`

	// Specifies the idle timeout in minutes for TCP connections. Valid values are between 4 and 100 minutes. Defaults to 4 minutes.
	// +kubebuilder:validation:Optional
	IdleTimeoutInMinutes *float64 `json:"idleTimeoutInMinutes,omitempty" tf:"idle_timeout_in_minutes,omitempty"`

	// Specifies the load balancing distribution type to be used by the Load Balancer. Possible values are: Default – The load balancer is configured to use a 5 tuple hash to map traffic to available servers. SourceIP – The load balancer is configured to use a 2 tuple hash to map traffic to available servers. SourceIPProtocol – The load balancer is configured to use a 3 tuple hash to map traffic to available servers. Also known as Session Persistence, where in the Azure portal the options are called None, Client IP and Client IP and Protocol respectively. Defaults to Default.
	// +kubebuilder:validation:Optional
	LoadDistribution *string `json:"loadDistribution,omitempty" tf:"load_distribution,omitempty"`

	// The ID of the Load Balancer in which to create the Rule. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/network/v1beta1.LoadBalancer
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	LoadbalancerID *string `json:"loadbalancerId,omitempty" tf:"loadbalancer_id,omitempty"`

	// Reference to a LoadBalancer in network to populate loadbalancerId.
	// +kubebuilder:validation:Optional
	LoadbalancerIDRef *v1.Reference `json:"loadbalancerIdRef,omitempty" tf:"-"`

	// Selector for a LoadBalancer in network to populate loadbalancerId.
	// +kubebuilder:validation:Optional
	LoadbalancerIDSelector *v1.Selector `json:"loadbalancerIdSelector,omitempty" tf:"-"`

	// A reference to a Probe used by this Load Balancing Rule.
	// +kubebuilder:validation:Optional
	ProbeID *string `json:"probeId,omitempty" tf:"probe_id,omitempty"`

	// The transport protocol for the external endpoint. Possible values are Tcp, Udp or All.
	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`
}

// LoadBalancerRuleSpec defines the desired state of LoadBalancerRule
type LoadBalancerRuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LoadBalancerRuleParameters `json:"forProvider"`
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
	InitProvider LoadBalancerRuleInitParameters `json:"initProvider,omitempty"`
}

// LoadBalancerRuleStatus defines the observed state of LoadBalancerRule.
type LoadBalancerRuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LoadBalancerRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// LoadBalancerRule is the Schema for the LoadBalancerRules API. Manages a Load Balancer Rule.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type LoadBalancerRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.backendPort) || (has(self.initProvider) && has(self.initProvider.backendPort))",message="spec.forProvider.backendPort is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.frontendIpConfigurationName) || (has(self.initProvider) && has(self.initProvider.frontendIpConfigurationName))",message="spec.forProvider.frontendIpConfigurationName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.frontendPort) || (has(self.initProvider) && has(self.initProvider.frontendPort))",message="spec.forProvider.frontendPort is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.protocol) || (has(self.initProvider) && has(self.initProvider.protocol))",message="spec.forProvider.protocol is a required parameter"
	Spec   LoadBalancerRuleSpec   `json:"spec"`
	Status LoadBalancerRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LoadBalancerRuleList contains a list of LoadBalancerRules
type LoadBalancerRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LoadBalancerRule `json:"items"`
}

// Repository type metadata.
var (
	LoadBalancerRule_Kind             = "LoadBalancerRule"
	LoadBalancerRule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LoadBalancerRule_Kind}.String()
	LoadBalancerRule_KindAPIVersion   = LoadBalancerRule_Kind + "." + CRDGroupVersion.String()
	LoadBalancerRule_GroupVersionKind = CRDGroupVersion.WithKind(LoadBalancerRule_Kind)
)

func init() {
	SchemeBuilder.Register(&LoadBalancerRule{}, &LoadBalancerRuleList{})
}
