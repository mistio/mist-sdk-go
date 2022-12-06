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

// KvmCloudFeatures struct for KvmCloudFeatures
type KvmCloudFeatures struct {
	Compute *bool `json:"compute,omitempty"`
}

// NewKvmCloudFeatures instantiates a new KvmCloudFeatures object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKvmCloudFeatures() *KvmCloudFeatures {
	this := KvmCloudFeatures{}
	var compute bool = true
	this.Compute = &compute
	return &this
}

// NewKvmCloudFeaturesWithDefaults instantiates a new KvmCloudFeatures object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKvmCloudFeaturesWithDefaults() *KvmCloudFeatures {
	this := KvmCloudFeatures{}
	var compute bool = true
	this.Compute = &compute
	return &this
}

// GetCompute returns the Compute field value if set, zero value otherwise.
func (o *KvmCloudFeatures) GetCompute() bool {
	if o == nil || o.Compute == nil {
		var ret bool
		return ret
	}
	return *o.Compute
}

// GetComputeOk returns a tuple with the Compute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KvmCloudFeatures) GetComputeOk() (*bool, bool) {
	if o == nil || o.Compute == nil {
		return nil, false
	}
	return o.Compute, true
}

// HasCompute returns a boolean if a field has been set.
func (o *KvmCloudFeatures) HasCompute() bool {
	if o != nil && o.Compute != nil {
		return true
	}

	return false
}

// SetCompute gets a reference to the given bool and assigns it to the Compute field.
func (o *KvmCloudFeatures) SetCompute(v bool) {
	o.Compute = &v
}

func (o KvmCloudFeatures) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Compute != nil {
		toSerialize["compute"] = o.Compute
	}
	return json.Marshal(toSerialize)
}

type NullableKvmCloudFeatures struct {
	value *KvmCloudFeatures
	isSet bool
}

func (v NullableKvmCloudFeatures) Get() *KvmCloudFeatures {
	return v.value
}

func (v *NullableKvmCloudFeatures) Set(val *KvmCloudFeatures) {
	v.value = val
	v.isSet = true
}

func (v NullableKvmCloudFeatures) IsSet() bool {
	return v.isSet
}

func (v *NullableKvmCloudFeatures) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKvmCloudFeatures(val *KvmCloudFeatures) *NullableKvmCloudFeatures {
	return &NullableKvmCloudFeatures{value: val, isSet: true}
}

func (v NullableKvmCloudFeatures) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKvmCloudFeatures) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


