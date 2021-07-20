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

// GetDatapointsResponse struct for GetDatapointsResponse
type GetDatapointsResponse struct {
	Data *Datapoints `json:"data,omitempty"`
	Meta *ResponseMetadata `json:"meta,omitempty"`
}

// NewGetDatapointsResponse instantiates a new GetDatapointsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDatapointsResponse() *GetDatapointsResponse {
	this := GetDatapointsResponse{}
	return &this
}

// NewGetDatapointsResponseWithDefaults instantiates a new GetDatapointsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDatapointsResponseWithDefaults() *GetDatapointsResponse {
	this := GetDatapointsResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetDatapointsResponse) GetData() Datapoints {
	if o == nil || o.Data == nil {
		var ret Datapoints
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDatapointsResponse) GetDataOk() (*Datapoints, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetDatapointsResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given Datapoints and assigns it to the Data field.
func (o *GetDatapointsResponse) SetData(v Datapoints) {
	o.Data = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *GetDatapointsResponse) GetMeta() ResponseMetadata {
	if o == nil || o.Meta == nil {
		var ret ResponseMetadata
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDatapointsResponse) GetMetaOk() (*ResponseMetadata, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *GetDatapointsResponse) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given ResponseMetadata and assigns it to the Meta field.
func (o *GetDatapointsResponse) SetMeta(v ResponseMetadata) {
	o.Meta = &v
}

func (o GetDatapointsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	return json.Marshal(toSerialize)
}

type NullableGetDatapointsResponse struct {
	value *GetDatapointsResponse
	isSet bool
}

func (v NullableGetDatapointsResponse) Get() *GetDatapointsResponse {
	return v.value
}

func (v *NullableGetDatapointsResponse) Set(val *GetDatapointsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDatapointsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDatapointsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDatapointsResponse(val *GetDatapointsResponse) *NullableGetDatapointsResponse {
	return &NullableGetDatapointsResponse{value: val, isSet: true}
}

func (v NullableGetDatapointsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDatapointsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

