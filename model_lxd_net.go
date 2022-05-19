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

// LXDNet struct for LXDNet
type LXDNet struct {
	// An array of network names or IDs to launch the container into
	Networks []string `json:"networks,omitempty"`
}

// NewLXDNet instantiates a new LXDNet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLXDNet() *LXDNet {
	this := LXDNet{}
	return &this
}

// NewLXDNetWithDefaults instantiates a new LXDNet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLXDNetWithDefaults() *LXDNet {
	this := LXDNet{}
	return &this
}

// GetNetworks returns the Networks field value if set, zero value otherwise.
func (o *LXDNet) GetNetworks() []string {
	if o == nil || o.Networks == nil {
		var ret []string
		return ret
	}
	return o.Networks
}

// GetNetworksOk returns a tuple with the Networks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LXDNet) GetNetworksOk() ([]string, bool) {
	if o == nil || o.Networks == nil {
		return nil, false
	}
	return o.Networks, true
}

// HasNetworks returns a boolean if a field has been set.
func (o *LXDNet) HasNetworks() bool {
	if o != nil && o.Networks != nil {
		return true
	}

	return false
}

// SetNetworks gets a reference to the given []string and assigns it to the Networks field.
func (o *LXDNet) SetNetworks(v []string) {
	o.Networks = v
}

func (o LXDNet) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Networks != nil {
		toSerialize["networks"] = o.Networks
	}
	return json.Marshal(toSerialize)
}

type NullableLXDNet struct {
	value *LXDNet
	isSet bool
}

func (v NullableLXDNet) Get() *LXDNet {
	return v.value
}

func (v *NullableLXDNet) Set(val *LXDNet) {
	v.value = val
	v.isSet = true
}

func (v NullableLXDNet) IsSet() bool {
	return v.isSet
}

func (v *NullableLXDNet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLXDNet(val *LXDNet) *NullableLXDNet {
	return &NullableLXDNet{value: val, isSet: true}
}

func (v NullableLXDNet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLXDNet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


