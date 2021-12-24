/*
 * Mist API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 2.0.0
 * Contact: api@mist.io
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mist_sdk

import (
	"encoding/json"
)

// KVMNetNetworks struct for KVMNetNetworks
type KVMNetNetworks struct {
	// The IPv4 address for the default Gateway
	Gateway *string `json:"gateway,omitempty"`
	// The IPv4 address to statically assign to the interface
	Ip *string `json:"ip,omitempty"`
	// Name or ID of the network, if only this field is provided a dynamic IP address will be assigned
	Network string `json:"network"`
	// The primary interface, which will be assigned a routing rule for the default GW
	Primary *string `json:"primary,omitempty"`
}

// NewKVMNetNetworks instantiates a new KVMNetNetworks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKVMNetNetworks(network string, ) *KVMNetNetworks {
	this := KVMNetNetworks{}
	this.Network = network
	return &this
}

// NewKVMNetNetworksWithDefaults instantiates a new KVMNetNetworks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKVMNetNetworksWithDefaults() *KVMNetNetworks {
	this := KVMNetNetworks{}
	return &this
}

// GetGateway returns the Gateway field value if set, zero value otherwise.
func (o *KVMNetNetworks) GetGateway() string {
	if o == nil || o.Gateway == nil {
		var ret string
		return ret
	}
	return *o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KVMNetNetworks) GetGatewayOk() (*string, bool) {
	if o == nil || o.Gateway == nil {
		return nil, false
	}
	return o.Gateway, true
}

// HasGateway returns a boolean if a field has been set.
func (o *KVMNetNetworks) HasGateway() bool {
	if o != nil && o.Gateway != nil {
		return true
	}

	return false
}

// SetGateway gets a reference to the given string and assigns it to the Gateway field.
func (o *KVMNetNetworks) SetGateway(v string) {
	o.Gateway = &v
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *KVMNetNetworks) GetIp() string {
	if o == nil || o.Ip == nil {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KVMNetNetworks) GetIpOk() (*string, bool) {
	if o == nil || o.Ip == nil {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *KVMNetNetworks) HasIp() bool {
	if o != nil && o.Ip != nil {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *KVMNetNetworks) SetIp(v string) {
	o.Ip = &v
}

// GetNetwork returns the Network field value
func (o *KVMNetNetworks) GetNetwork() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Network
}

// GetNetworkOk returns a tuple with the Network field value
// and a boolean to check if the value has been set.
func (o *KVMNetNetworks) GetNetworkOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Network, true
}

// SetNetwork sets field value
func (o *KVMNetNetworks) SetNetwork(v string) {
	o.Network = v
}

// GetPrimary returns the Primary field value if set, zero value otherwise.
func (o *KVMNetNetworks) GetPrimary() string {
	if o == nil || o.Primary == nil {
		var ret string
		return ret
	}
	return *o.Primary
}

// GetPrimaryOk returns a tuple with the Primary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KVMNetNetworks) GetPrimaryOk() (*string, bool) {
	if o == nil || o.Primary == nil {
		return nil, false
	}
	return o.Primary, true
}

// HasPrimary returns a boolean if a field has been set.
func (o *KVMNetNetworks) HasPrimary() bool {
	if o != nil && o.Primary != nil {
		return true
	}

	return false
}

// SetPrimary gets a reference to the given string and assigns it to the Primary field.
func (o *KVMNetNetworks) SetPrimary(v string) {
	o.Primary = &v
}

func (o KVMNetNetworks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Gateway != nil {
		toSerialize["gateway"] = o.Gateway
	}
	if o.Ip != nil {
		toSerialize["ip"] = o.Ip
	}
	if true {
		toSerialize["network"] = o.Network
	}
	if o.Primary != nil {
		toSerialize["primary"] = o.Primary
	}
	return json.Marshal(toSerialize)
}

type NullableKVMNetNetworks struct {
	value *KVMNetNetworks
	isSet bool
}

func (v NullableKVMNetNetworks) Get() *KVMNetNetworks {
	return v.value
}

func (v *NullableKVMNetNetworks) Set(val *KVMNetNetworks) {
	v.value = val
	v.isSet = true
}

func (v NullableKVMNetNetworks) IsSet() bool {
	return v.isSet
}

func (v *NullableKVMNetNetworks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKVMNetNetworks(val *KVMNetNetworks) *NullableKVMNetNetworks {
	return &NullableKVMNetNetworks{value: val, isSet: true}
}

func (v NullableKVMNetNetworks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKVMNetNetworks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


