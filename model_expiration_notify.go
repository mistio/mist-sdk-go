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

// ExpirationNotify Notify user before machine expiration
type ExpirationNotify struct {
	Value int32 `json:"value"`
	Period string `json:"period"`
}

// NewExpirationNotify instantiates a new ExpirationNotify object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExpirationNotify(value int32, period string) *ExpirationNotify {
	this := ExpirationNotify{}
	this.Value = value
	this.Period = period
	return &this
}

// NewExpirationNotifyWithDefaults instantiates a new ExpirationNotify object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExpirationNotifyWithDefaults() *ExpirationNotify {
	this := ExpirationNotify{}
	return &this
}

// GetValue returns the Value field value
func (o *ExpirationNotify) GetValue() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *ExpirationNotify) GetValueOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *ExpirationNotify) SetValue(v int32) {
	o.Value = v
}

// GetPeriod returns the Period field value
func (o *ExpirationNotify) GetPeriod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Period
}

// GetPeriodOk returns a tuple with the Period field value
// and a boolean to check if the value has been set.
func (o *ExpirationNotify) GetPeriodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Period, true
}

// SetPeriod sets field value
func (o *ExpirationNotify) SetPeriod(v string) {
	o.Period = v
}

func (o ExpirationNotify) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["value"] = o.Value
	}
	if true {
		toSerialize["period"] = o.Period
	}
	return json.Marshal(toSerialize)
}

type NullableExpirationNotify struct {
	value *ExpirationNotify
	isSet bool
}

func (v NullableExpirationNotify) Get() *ExpirationNotify {
	return v.value
}

func (v *NullableExpirationNotify) Set(val *ExpirationNotify) {
	v.value = val
	v.isSet = true
}

func (v NullableExpirationNotify) IsSet() bool {
	return v.isSet
}

func (v *NullableExpirationNotify) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExpirationNotify(val *ExpirationNotify) *NullableExpirationNotify {
	return &NullableExpirationNotify{value: val, isSet: true}
}

func (v NullableExpirationNotify) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExpirationNotify) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


