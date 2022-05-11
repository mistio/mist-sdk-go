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

// AzureNet struct for AzureNet
type AzureNet struct {
	// A new or existing network If not provided a `mist-resource_group-location` network will be used.
	Network *string `json:"network,omitempty"`
}

// NewAzureNet instantiates a new AzureNet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureNet() *AzureNet {
	this := AzureNet{}
	return &this
}

// NewAzureNetWithDefaults instantiates a new AzureNet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureNetWithDefaults() *AzureNet {
	this := AzureNet{}
	return &this
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *AzureNet) GetNetwork() string {
	if o == nil || o.Network == nil {
		var ret string
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureNet) GetNetworkOk() (*string, bool) {
	if o == nil || o.Network == nil {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *AzureNet) HasNetwork() bool {
	if o != nil && o.Network != nil {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given string and assigns it to the Network field.
func (o *AzureNet) SetNetwork(v string) {
	o.Network = &v
}

func (o AzureNet) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Network != nil {
		toSerialize["network"] = o.Network
	}
	return json.Marshal(toSerialize)
}

type NullableAzureNet struct {
	value *AzureNet
	isSet bool
}

func (v NullableAzureNet) Get() *AzureNet {
	return v.value
}

func (v *NullableAzureNet) Set(val *AzureNet) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureNet) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureNet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureNet(val *AzureNet) *NullableAzureNet {
	return &NullableAzureNet{value: val, isSet: true}
}

func (v NullableAzureNet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureNet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


