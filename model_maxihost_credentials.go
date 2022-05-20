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

// MaxihostCredentials struct for MaxihostCredentials
type MaxihostCredentials struct {
	// Your Maxihost API token
	Token string `json:"token"`
}

// NewMaxihostCredentials instantiates a new MaxihostCredentials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMaxihostCredentials(token string) *MaxihostCredentials {
	this := MaxihostCredentials{}
	this.Token = token
	return &this
}

// NewMaxihostCredentialsWithDefaults instantiates a new MaxihostCredentials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMaxihostCredentialsWithDefaults() *MaxihostCredentials {
	this := MaxihostCredentials{}
	return &this
}

// GetToken returns the Token field value
func (o *MaxihostCredentials) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *MaxihostCredentials) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *MaxihostCredentials) SetToken(v string) {
	o.Token = v
}

func (o MaxihostCredentials) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["token"] = o.Token
	}
	return json.Marshal(toSerialize)
}

type NullableMaxihostCredentials struct {
	value *MaxihostCredentials
	isSet bool
}

func (v NullableMaxihostCredentials) Get() *MaxihostCredentials {
	return v.value
}

func (v *NullableMaxihostCredentials) Set(val *MaxihostCredentials) {
	v.value = val
	v.isSet = true
}

func (v NullableMaxihostCredentials) IsSet() bool {
	return v.isSet
}

func (v *NullableMaxihostCredentials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMaxihostCredentials(val *MaxihostCredentials) *NullableMaxihostCredentials {
	return &NullableMaxihostCredentials{value: val, isSet: true}
}

func (v NullableMaxihostCredentials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMaxihostCredentials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


