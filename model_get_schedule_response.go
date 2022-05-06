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

// GetScheduleResponse struct for GetScheduleResponse
type GetScheduleResponse struct {
	Data *Schedule `json:"data,omitempty"`
	Meta *ResponseMetadata `json:"meta,omitempty"`
}

// NewGetScheduleResponse instantiates a new GetScheduleResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetScheduleResponse() *GetScheduleResponse {
	this := GetScheduleResponse{}
	return &this
}

// NewGetScheduleResponseWithDefaults instantiates a new GetScheduleResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetScheduleResponseWithDefaults() *GetScheduleResponse {
	this := GetScheduleResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetScheduleResponse) GetData() Schedule {
	if o == nil || o.Data == nil {
		var ret Schedule
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetScheduleResponse) GetDataOk() (*Schedule, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetScheduleResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given Schedule and assigns it to the Data field.
func (o *GetScheduleResponse) SetData(v Schedule) {
	o.Data = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *GetScheduleResponse) GetMeta() ResponseMetadata {
	if o == nil || o.Meta == nil {
		var ret ResponseMetadata
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetScheduleResponse) GetMetaOk() (*ResponseMetadata, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *GetScheduleResponse) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given ResponseMetadata and assigns it to the Meta field.
func (o *GetScheduleResponse) SetMeta(v ResponseMetadata) {
	o.Meta = &v
}

func (o GetScheduleResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	return json.Marshal(toSerialize)
}

type NullableGetScheduleResponse struct {
	value *GetScheduleResponse
	isSet bool
}

func (v NullableGetScheduleResponse) Get() *GetScheduleResponse {
	return v.value
}

func (v *NullableGetScheduleResponse) Set(val *GetScheduleResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetScheduleResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetScheduleResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetScheduleResponse(val *GetScheduleResponse) *NullableGetScheduleResponse {
	return &NullableGetScheduleResponse{value: val, isSet: true}
}

func (v NullableGetScheduleResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetScheduleResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

