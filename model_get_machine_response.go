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

// GetMachineResponse struct for GetMachineResponse
type GetMachineResponse struct {
	Data *Machine `json:"data,omitempty"`
	Meta *ResponseMetadata `json:"meta,omitempty"`
}

// NewGetMachineResponse instantiates a new GetMachineResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMachineResponse() *GetMachineResponse {
	this := GetMachineResponse{}
	return &this
}

// NewGetMachineResponseWithDefaults instantiates a new GetMachineResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMachineResponseWithDefaults() *GetMachineResponse {
	this := GetMachineResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetMachineResponse) GetData() Machine {
	if o == nil || o.Data == nil {
		var ret Machine
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMachineResponse) GetDataOk() (*Machine, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetMachineResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given Machine and assigns it to the Data field.
func (o *GetMachineResponse) SetData(v Machine) {
	o.Data = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *GetMachineResponse) GetMeta() ResponseMetadata {
	if o == nil || o.Meta == nil {
		var ret ResponseMetadata
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMachineResponse) GetMetaOk() (*ResponseMetadata, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *GetMachineResponse) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given ResponseMetadata and assigns it to the Meta field.
func (o *GetMachineResponse) SetMeta(v ResponseMetadata) {
	o.Meta = &v
}

func (o GetMachineResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	return json.Marshal(toSerialize)
}

type NullableGetMachineResponse struct {
	value *GetMachineResponse
	isSet bool
}

func (v NullableGetMachineResponse) Get() *GetMachineResponse {
	return v.value
}

func (v *NullableGetMachineResponse) Set(val *GetMachineResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMachineResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMachineResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMachineResponse(val *GetMachineResponse) *NullableGetMachineResponse {
	return &NullableGetMachineResponse{value: val, isSet: true}
}

func (v NullableGetMachineResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMachineResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


