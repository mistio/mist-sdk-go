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

// EditCloudRequestAnyOf struct for EditCloudRequestAnyOf
type EditCloudRequestAnyOf struct {
	// Updated name
	Name string `json:"name"`
}

// NewEditCloudRequestAnyOf instantiates a new EditCloudRequestAnyOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEditCloudRequestAnyOf(name string) *EditCloudRequestAnyOf {
	this := EditCloudRequestAnyOf{}
	this.Name = name
	return &this
}

// NewEditCloudRequestAnyOfWithDefaults instantiates a new EditCloudRequestAnyOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEditCloudRequestAnyOfWithDefaults() *EditCloudRequestAnyOf {
	this := EditCloudRequestAnyOf{}
	return &this
}

// GetName returns the Name field value
func (o *EditCloudRequestAnyOf) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EditCloudRequestAnyOf) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EditCloudRequestAnyOf) SetName(v string) {
	o.Name = v
}

func (o EditCloudRequestAnyOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableEditCloudRequestAnyOf struct {
	value *EditCloudRequestAnyOf
	isSet bool
}

func (v NullableEditCloudRequestAnyOf) Get() *EditCloudRequestAnyOf {
	return v.value
}

func (v *NullableEditCloudRequestAnyOf) Set(val *EditCloudRequestAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEditCloudRequestAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEditCloudRequestAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEditCloudRequestAnyOf(val *EditCloudRequestAnyOf) *NullableEditCloudRequestAnyOf {
	return &NullableEditCloudRequestAnyOf{value: val, isSet: true}
}

func (v NullableEditCloudRequestAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEditCloudRequestAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

