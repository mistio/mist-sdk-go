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

// ListScriptsResponse struct for ListScriptsResponse
type ListScriptsResponse struct {
	Data []Script `json:"data,omitempty"`
	Meta *ResponseMetadata `json:"meta,omitempty"`
}

// NewListScriptsResponse instantiates a new ListScriptsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListScriptsResponse() *ListScriptsResponse {
	this := ListScriptsResponse{}
	return &this
}

// NewListScriptsResponseWithDefaults instantiates a new ListScriptsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListScriptsResponseWithDefaults() *ListScriptsResponse {
	this := ListScriptsResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ListScriptsResponse) GetData() []Script {
	if o == nil || o.Data == nil {
		var ret []Script
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListScriptsResponse) GetDataOk() ([]Script, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ListScriptsResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []Script and assigns it to the Data field.
func (o *ListScriptsResponse) SetData(v []Script) {
	o.Data = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ListScriptsResponse) GetMeta() ResponseMetadata {
	if o == nil || o.Meta == nil {
		var ret ResponseMetadata
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListScriptsResponse) GetMetaOk() (*ResponseMetadata, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ListScriptsResponse) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given ResponseMetadata and assigns it to the Meta field.
func (o *ListScriptsResponse) SetMeta(v ResponseMetadata) {
	o.Meta = &v
}

func (o ListScriptsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	return json.Marshal(toSerialize)
}

type NullableListScriptsResponse struct {
	value *ListScriptsResponse
	isSet bool
}

func (v NullableListScriptsResponse) Get() *ListScriptsResponse {
	return v.value
}

func (v *NullableListScriptsResponse) Set(val *ListScriptsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListScriptsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListScriptsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListScriptsResponse(val *ListScriptsResponse) *NullableListScriptsResponse {
	return &NullableListScriptsResponse{value: val, isSet: true}
}

func (v NullableListScriptsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListScriptsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


