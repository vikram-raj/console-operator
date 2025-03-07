// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	consolev1 "github.com/openshift/api/console/v1"
)

// ConsolePluginProxyEndpointApplyConfiguration represents a declarative configuration of the ConsolePluginProxyEndpoint type for use
// with apply.
type ConsolePluginProxyEndpointApplyConfiguration struct {
	Type    *consolev1.ConsolePluginProxyType                  `json:"type,omitempty"`
	Service *ConsolePluginProxyServiceConfigApplyConfiguration `json:"service,omitempty"`
}

// ConsolePluginProxyEndpointApplyConfiguration constructs a declarative configuration of the ConsolePluginProxyEndpoint type for use with
// apply.
func ConsolePluginProxyEndpoint() *ConsolePluginProxyEndpointApplyConfiguration {
	return &ConsolePluginProxyEndpointApplyConfiguration{}
}

// WithType sets the Type field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Type field is set to the value of the last call.
func (b *ConsolePluginProxyEndpointApplyConfiguration) WithType(value consolev1.ConsolePluginProxyType) *ConsolePluginProxyEndpointApplyConfiguration {
	b.Type = &value
	return b
}

// WithService sets the Service field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Service field is set to the value of the last call.
func (b *ConsolePluginProxyEndpointApplyConfiguration) WithService(value *ConsolePluginProxyServiceConfigApplyConfiguration) *ConsolePluginProxyEndpointApplyConfiguration {
	b.Service = value
	return b
}
