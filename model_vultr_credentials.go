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

// VultrCredentials struct for VultrCredentials
type VultrCredentials struct {
	// Your Vultr API key
	Apikey string `json:"apikey"`
}

// NewVultrCredentials instantiates a new VultrCredentials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVultrCredentials(apikey string) *VultrCredentials {
	this := VultrCredentials{}
	this.Apikey = apikey
	return &this
}

// NewVultrCredentialsWithDefaults instantiates a new VultrCredentials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVultrCredentialsWithDefaults() *VultrCredentials {
	this := VultrCredentials{}
	return &this
}

// GetApikey returns the Apikey field value
func (o *VultrCredentials) GetApikey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Apikey
}

// GetApikeyOk returns a tuple with the Apikey field value
// and a boolean to check if the value has been set.
func (o *VultrCredentials) GetApikeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Apikey, true
}

// SetApikey sets field value
func (o *VultrCredentials) SetApikey(v string) {
	o.Apikey = v
}

func (o VultrCredentials) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["apikey"] = o.Apikey
	}
	return json.Marshal(toSerialize)
}

type NullableVultrCredentials struct {
	value *VultrCredentials
	isSet bool
}

func (v NullableVultrCredentials) Get() *VultrCredentials {
	return v.value
}

func (v *NullableVultrCredentials) Set(val *VultrCredentials) {
	v.value = val
	v.isSet = true
}

func (v NullableVultrCredentials) IsSet() bool {
	return v.isSet
}

func (v *NullableVultrCredentials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVultrCredentials(val *VultrCredentials) *NullableVultrCredentials {
	return &NullableVultrCredentials{value: val, isSet: true}
}

func (v NullableVultrCredentials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVultrCredentials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


