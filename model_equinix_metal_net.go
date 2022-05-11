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

// EquinixMetalNet struct for EquinixMetalNet
type EquinixMetalNet struct {
	// Specify addresses to be created with your machine.By default Equinix configures public IPv4, public IPv6, and private IPv4
	IpAddresses []EquinixMetalNetIpAddresses `json:"ip_addresses,omitempty"`
}

// NewEquinixMetalNet instantiates a new EquinixMetalNet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEquinixMetalNet() *EquinixMetalNet {
	this := EquinixMetalNet{}
	return &this
}

// NewEquinixMetalNetWithDefaults instantiates a new EquinixMetalNet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEquinixMetalNetWithDefaults() *EquinixMetalNet {
	this := EquinixMetalNet{}
	return &this
}

// GetIpAddresses returns the IpAddresses field value if set, zero value otherwise.
func (o *EquinixMetalNet) GetIpAddresses() []EquinixMetalNetIpAddresses {
	if o == nil || o.IpAddresses == nil {
		var ret []EquinixMetalNetIpAddresses
		return ret
	}
	return o.IpAddresses
}

// GetIpAddressesOk returns a tuple with the IpAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquinixMetalNet) GetIpAddressesOk() ([]EquinixMetalNetIpAddresses, bool) {
	if o == nil || o.IpAddresses == nil {
		return nil, false
	}
	return o.IpAddresses, true
}

// HasIpAddresses returns a boolean if a field has been set.
func (o *EquinixMetalNet) HasIpAddresses() bool {
	if o != nil && o.IpAddresses != nil {
		return true
	}

	return false
}

// SetIpAddresses gets a reference to the given []EquinixMetalNetIpAddresses and assigns it to the IpAddresses field.
func (o *EquinixMetalNet) SetIpAddresses(v []EquinixMetalNetIpAddresses) {
	o.IpAddresses = v
}

func (o EquinixMetalNet) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IpAddresses != nil {
		toSerialize["ip_addresses"] = o.IpAddresses
	}
	return json.Marshal(toSerialize)
}

type NullableEquinixMetalNet struct {
	value *EquinixMetalNet
	isSet bool
}

func (v NullableEquinixMetalNet) Get() *EquinixMetalNet {
	return v.value
}

func (v *NullableEquinixMetalNet) Set(val *EquinixMetalNet) {
	v.value = val
	v.isSet = true
}

func (v NullableEquinixMetalNet) IsSet() bool {
	return v.isSet
}

func (v *NullableEquinixMetalNet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEquinixMetalNet(val *EquinixMetalNet) *NullableEquinixMetalNet {
	return &NullableEquinixMetalNet{value: val, isSet: true}
}

func (v NullableEquinixMetalNet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEquinixMetalNet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


