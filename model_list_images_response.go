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

// ListImagesResponse struct for ListImagesResponse
type ListImagesResponse struct {
	Data *[]Image `json:"data,omitempty"`
	Meta *ResponseMetadata `json:"meta,omitempty"`
}

// NewListImagesResponse instantiates a new ListImagesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListImagesResponse() *ListImagesResponse {
	this := ListImagesResponse{}
	return &this
}

// NewListImagesResponseWithDefaults instantiates a new ListImagesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListImagesResponseWithDefaults() *ListImagesResponse {
	this := ListImagesResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ListImagesResponse) GetData() []Image {
	if o == nil || o.Data == nil {
		var ret []Image
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListImagesResponse) GetDataOk() (*[]Image, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ListImagesResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []Image and assigns it to the Data field.
func (o *ListImagesResponse) SetData(v []Image) {
	o.Data = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ListImagesResponse) GetMeta() ResponseMetadata {
	if o == nil || o.Meta == nil {
		var ret ResponseMetadata
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListImagesResponse) GetMetaOk() (*ResponseMetadata, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ListImagesResponse) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given ResponseMetadata and assigns it to the Meta field.
func (o *ListImagesResponse) SetMeta(v ResponseMetadata) {
	o.Meta = &v
}

func (o ListImagesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	return json.Marshal(toSerialize)
}

type NullableListImagesResponse struct {
	value *ListImagesResponse
	isSet bool
}

func (v NullableListImagesResponse) Get() *ListImagesResponse {
	return v.value
}

func (v *NullableListImagesResponse) Set(val *ListImagesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListImagesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListImagesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListImagesResponse(val *ListImagesResponse) *NullableListImagesResponse {
	return &NullableListImagesResponse{value: val, isSet: true}
}

func (v NullableListImagesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListImagesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


