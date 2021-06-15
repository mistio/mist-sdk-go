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

// CreateMachineResponseOneOf struct for CreateMachineResponseOneOf
type CreateMachineResponseOneOf struct {
	JobId *string `json:"jobId,omitempty"`
	Plan *map[string]interface{} `json:"plan,omitempty"`
}

// NewCreateMachineResponseOneOf instantiates a new CreateMachineResponseOneOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateMachineResponseOneOf() *CreateMachineResponseOneOf {
	this := CreateMachineResponseOneOf{}
	return &this
}

// NewCreateMachineResponseOneOfWithDefaults instantiates a new CreateMachineResponseOneOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateMachineResponseOneOfWithDefaults() *CreateMachineResponseOneOf {
	this := CreateMachineResponseOneOf{}
	return &this
}

// GetJobId returns the JobId field value if set, zero value otherwise.
func (o *CreateMachineResponseOneOf) GetJobId() string {
	if o == nil || o.JobId == nil {
		var ret string
		return ret
	}
	return *o.JobId
}

// GetJobIdOk returns a tuple with the JobId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMachineResponseOneOf) GetJobIdOk() (*string, bool) {
	if o == nil || o.JobId == nil {
		return nil, false
	}
	return o.JobId, true
}

// HasJobId returns a boolean if a field has been set.
func (o *CreateMachineResponseOneOf) HasJobId() bool {
	if o != nil && o.JobId != nil {
		return true
	}

	return false
}

// SetJobId gets a reference to the given string and assigns it to the JobId field.
func (o *CreateMachineResponseOneOf) SetJobId(v string) {
	o.JobId = &v
}

// GetPlan returns the Plan field value if set, zero value otherwise.
func (o *CreateMachineResponseOneOf) GetPlan() map[string]interface{} {
	if o == nil || o.Plan == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Plan
}

// GetPlanOk returns a tuple with the Plan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMachineResponseOneOf) GetPlanOk() (*map[string]interface{}, bool) {
	if o == nil || o.Plan == nil {
		return nil, false
	}
	return o.Plan, true
}

// HasPlan returns a boolean if a field has been set.
func (o *CreateMachineResponseOneOf) HasPlan() bool {
	if o != nil && o.Plan != nil {
		return true
	}

	return false
}

// SetPlan gets a reference to the given map[string]interface{} and assigns it to the Plan field.
func (o *CreateMachineResponseOneOf) SetPlan(v map[string]interface{}) {
	o.Plan = &v
}

func (o CreateMachineResponseOneOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.JobId != nil {
		toSerialize["jobId"] = o.JobId
	}
	if o.Plan != nil {
		toSerialize["plan"] = o.Plan
	}
	return json.Marshal(toSerialize)
}

type NullableCreateMachineResponseOneOf struct {
	value *CreateMachineResponseOneOf
	isSet bool
}

func (v NullableCreateMachineResponseOneOf) Get() *CreateMachineResponseOneOf {
	return v.value
}

func (v *NullableCreateMachineResponseOneOf) Set(val *CreateMachineResponseOneOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateMachineResponseOneOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateMachineResponseOneOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateMachineResponseOneOf(val *CreateMachineResponseOneOf) *NullableCreateMachineResponseOneOf {
	return &NullableCreateMachineResponseOneOf{value: val, isSet: true}
}

func (v NullableCreateMachineResponseOneOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateMachineResponseOneOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


