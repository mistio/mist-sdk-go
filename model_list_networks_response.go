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

// ListNetworksResponse struct for ListNetworksResponse
type ListNetworksResponse struct {
	Data *[]Network `json:"data,omitempty"`
	Meta *ResponseMetadata `json:"meta,omitempty"`
}

// NewListNetworksResponse instantiates a new ListNetworksResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListNetworksResponse() *ListNetworksResponse {
	this := ListNetworksResponse{}
	return &this
}

// NewListNetworksResponseWithDefaults instantiates a new ListNetworksResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListNetworksResponseWithDefaults() *ListNetworksResponse {
	this := ListNetworksResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ListNetworksResponse) GetData() []Network {
	if o == nil || o.Data == nil {
		var ret []Network
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListNetworksResponse) GetDataOk() (*[]Network, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ListNetworksResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []Network and assigns it to the Data field.
func (o *ListNetworksResponse) SetData(v []Network) {
	o.Data = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ListNetworksResponse) GetMeta() ResponseMetadata {
	if o == nil || o.Meta == nil {
		var ret ResponseMetadata
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListNetworksResponse) GetMetaOk() (*ResponseMetadata, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ListNetworksResponse) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given ResponseMetadata and assigns it to the Meta field.
func (o *ListNetworksResponse) SetMeta(v ResponseMetadata) {
	o.Meta = &v
}

func (o ListNetworksResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	return json.Marshal(toSerialize)
}

type NullableListNetworksResponse struct {
	value *ListNetworksResponse
	isSet bool
}

func (v NullableListNetworksResponse) Get() *ListNetworksResponse {
	return v.value
}

func (v *NullableListNetworksResponse) Set(val *ListNetworksResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListNetworksResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListNetworksResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListNetworksResponse(val *ListNetworksResponse) *NullableListNetworksResponse {
	return &NullableListNetworksResponse{value: val, isSet: true}
}

func (v NullableListNetworksResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListNetworksResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


