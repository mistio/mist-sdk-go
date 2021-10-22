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

// ExposeMachineRequest struct for ExposeMachineRequest
type ExposeMachineRequest struct {
	// Applies only in GigG8 clouds
	PortForwards map[string]interface{} `json:"port_forwards"`
}

// NewExposeMachineRequest instantiates a new ExposeMachineRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExposeMachineRequest(portForwards map[string]interface{}, ) *ExposeMachineRequest {
	this := ExposeMachineRequest{}
	this.PortForwards = portForwards
	return &this
}

// NewExposeMachineRequestWithDefaults instantiates a new ExposeMachineRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExposeMachineRequestWithDefaults() *ExposeMachineRequest {
	this := ExposeMachineRequest{}
	return &this
}

// GetPortForwards returns the PortForwards field value
func (o *ExposeMachineRequest) GetPortForwards() map[string]interface{} {
	if o == nil  {
		var ret map[string]interface{}
		return ret
	}

	return o.PortForwards
}

// GetPortForwardsOk returns a tuple with the PortForwards field value
// and a boolean to check if the value has been set.
func (o *ExposeMachineRequest) GetPortForwardsOk() (*map[string]interface{}, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PortForwards, true
}

// SetPortForwards sets field value
func (o *ExposeMachineRequest) SetPortForwards(v map[string]interface{}) {
	o.PortForwards = v
}

func (o ExposeMachineRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["port_forwards"] = o.PortForwards
	}
	return json.Marshal(toSerialize)
}

type NullableExposeMachineRequest struct {
	value *ExposeMachineRequest
	isSet bool
}

func (v NullableExposeMachineRequest) Get() *ExposeMachineRequest {
	return v.value
}

func (v *NullableExposeMachineRequest) Set(val *ExposeMachineRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableExposeMachineRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableExposeMachineRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExposeMachineRequest(val *ExposeMachineRequest) *NullableExposeMachineRequest {
	return &NullableExposeMachineRequest{value: val, isSet: true}
}

func (v NullableExposeMachineRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExposeMachineRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


