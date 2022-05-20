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

// EquinixCredentials struct for EquinixCredentials
type EquinixCredentials struct {
	// Your Equinix Metal API key
	Apikey string `json:"apikey"`
}

// NewEquinixCredentials instantiates a new EquinixCredentials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEquinixCredentials(apikey string) *EquinixCredentials {
	this := EquinixCredentials{}
	this.Apikey = apikey
	return &this
}

// NewEquinixCredentialsWithDefaults instantiates a new EquinixCredentials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEquinixCredentialsWithDefaults() *EquinixCredentials {
	this := EquinixCredentials{}
	return &this
}

// GetApikey returns the Apikey field value
func (o *EquinixCredentials) GetApikey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Apikey
}

// GetApikeyOk returns a tuple with the Apikey field value
// and a boolean to check if the value has been set.
func (o *EquinixCredentials) GetApikeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Apikey, true
}

// SetApikey sets field value
func (o *EquinixCredentials) SetApikey(v string) {
	o.Apikey = v
}

func (o EquinixCredentials) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["apikey"] = o.Apikey
	}
	return json.Marshal(toSerialize)
}

type NullableEquinixCredentials struct {
	value *EquinixCredentials
	isSet bool
}

func (v NullableEquinixCredentials) Get() *EquinixCredentials {
	return v.value
}

func (v *NullableEquinixCredentials) Set(val *EquinixCredentials) {
	v.value = val
	v.isSet = true
}

func (v NullableEquinixCredentials) IsSet() bool {
	return v.isSet
}

func (v *NullableEquinixCredentials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEquinixCredentials(val *EquinixCredentials) *NullableEquinixCredentials {
	return &NullableEquinixCredentials{value: val, isSet: true}
}

func (v NullableEquinixCredentials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEquinixCredentials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


