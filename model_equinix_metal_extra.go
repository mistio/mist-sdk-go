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

// EquinixMetalExtra struct for EquinixMetalExtra
type EquinixMetalExtra struct {
	// Project UUID, if not given the first one available will be selected
	ProjectId *string `json:"project_id,omitempty"`
}

// NewEquinixMetalExtra instantiates a new EquinixMetalExtra object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEquinixMetalExtra() *EquinixMetalExtra {
	this := EquinixMetalExtra{}
	return &this
}

// NewEquinixMetalExtraWithDefaults instantiates a new EquinixMetalExtra object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEquinixMetalExtraWithDefaults() *EquinixMetalExtra {
	this := EquinixMetalExtra{}
	return &this
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *EquinixMetalExtra) GetProjectId() string {
	if o == nil || o.ProjectId == nil {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquinixMetalExtra) GetProjectIdOk() (*string, bool) {
	if o == nil || o.ProjectId == nil {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *EquinixMetalExtra) HasProjectId() bool {
	if o != nil && o.ProjectId != nil {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *EquinixMetalExtra) SetProjectId(v string) {
	o.ProjectId = &v
}

func (o EquinixMetalExtra) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ProjectId != nil {
		toSerialize["project_id"] = o.ProjectId
	}
	return json.Marshal(toSerialize)
}

type NullableEquinixMetalExtra struct {
	value *EquinixMetalExtra
	isSet bool
}

func (v NullableEquinixMetalExtra) Get() *EquinixMetalExtra {
	return v.value
}

func (v *NullableEquinixMetalExtra) Set(val *EquinixMetalExtra) {
	v.value = val
	v.isSet = true
}

func (v NullableEquinixMetalExtra) IsSet() bool {
	return v.isSet
}

func (v *NullableEquinixMetalExtra) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEquinixMetalExtra(val *EquinixMetalExtra) *NullableEquinixMetalExtra {
	return &NullableEquinixMetalExtra{value: val, isSet: true}
}

func (v NullableEquinixMetalExtra) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEquinixMetalExtra) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

