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

// GetSizeResponse struct for GetSizeResponse
type GetSizeResponse struct {
	Data *Size `json:"data,omitempty"`
	Meta *ResponseMetadata `json:"meta,omitempty"`
}

// NewGetSizeResponse instantiates a new GetSizeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSizeResponse() *GetSizeResponse {
	this := GetSizeResponse{}
	return &this
}

// NewGetSizeResponseWithDefaults instantiates a new GetSizeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSizeResponseWithDefaults() *GetSizeResponse {
	this := GetSizeResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetSizeResponse) GetData() Size {
	if o == nil || o.Data == nil {
		var ret Size
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSizeResponse) GetDataOk() (*Size, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetSizeResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given Size and assigns it to the Data field.
func (o *GetSizeResponse) SetData(v Size) {
	o.Data = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *GetSizeResponse) GetMeta() ResponseMetadata {
	if o == nil || o.Meta == nil {
		var ret ResponseMetadata
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSizeResponse) GetMetaOk() (*ResponseMetadata, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *GetSizeResponse) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given ResponseMetadata and assigns it to the Meta field.
func (o *GetSizeResponse) SetMeta(v ResponseMetadata) {
	o.Meta = &v
}

func (o GetSizeResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	return json.Marshal(toSerialize)
}

type NullableGetSizeResponse struct {
	value *GetSizeResponse
	isSet bool
}

func (v NullableGetSizeResponse) Get() *GetSizeResponse {
	return v.value
}

func (v *NullableGetSizeResponse) Set(val *GetSizeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSizeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSizeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSizeResponse(val *GetSizeResponse) *NullableGetSizeResponse {
	return &NullableGetSizeResponse{value: val, isSet: true}
}

func (v NullableGetSizeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSizeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


