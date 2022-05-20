/*
Mist API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.0
Contact: api@mist.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mist_sdk

import (
	"encoding/json"
)

// DockerNet struct for DockerNet
type DockerNet struct {
	// ContainerPort,HostPort pairs. For example \"80\"/\"80\", \"123/udp\"/\"123\"
	PortBindings map[string]interface{} `json:"port_bindings,omitempty"`
}

// NewDockerNet instantiates a new DockerNet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDockerNet() *DockerNet {
	this := DockerNet{}
	return &this
}

// NewDockerNetWithDefaults instantiates a new DockerNet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDockerNetWithDefaults() *DockerNet {
	this := DockerNet{}
	return &this
}

// GetPortBindings returns the PortBindings field value if set, zero value otherwise.
func (o *DockerNet) GetPortBindings() map[string]interface{} {
	if o == nil || o.PortBindings == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.PortBindings
}

// GetPortBindingsOk returns a tuple with the PortBindings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DockerNet) GetPortBindingsOk() (map[string]interface{}, bool) {
	if o == nil || o.PortBindings == nil {
		return nil, false
	}
	return o.PortBindings, true
}

// HasPortBindings returns a boolean if a field has been set.
func (o *DockerNet) HasPortBindings() bool {
	if o != nil && o.PortBindings != nil {
		return true
	}

	return false
}

// SetPortBindings gets a reference to the given map[string]interface{} and assigns it to the PortBindings field.
func (o *DockerNet) SetPortBindings(v map[string]interface{}) {
	o.PortBindings = v
}

func (o DockerNet) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PortBindings != nil {
		toSerialize["port_bindings"] = o.PortBindings
	}
	return json.Marshal(toSerialize)
}

type NullableDockerNet struct {
	value *DockerNet
	isSet bool
}

func (v NullableDockerNet) Get() *DockerNet {
	return v.value
}

func (v *NullableDockerNet) Set(val *DockerNet) {
	v.value = val
	v.isSet = true
}

func (v NullableDockerNet) IsSet() bool {
	return v.isSet
}

func (v *NullableDockerNet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDockerNet(val *DockerNet) *NullableDockerNet {
	return &NullableDockerNet{value: val, isSet: true}
}

func (v NullableDockerNet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDockerNet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


