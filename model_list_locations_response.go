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

// ListLocationsResponse struct for ListLocationsResponse
type ListLocationsResponse struct {
	Data []Location `json:"data,omitempty"`
	Meta *ResponseMetadata `json:"meta,omitempty"`
}

// NewListLocationsResponse instantiates a new ListLocationsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListLocationsResponse() *ListLocationsResponse {
	this := ListLocationsResponse{}
	return &this
}

// NewListLocationsResponseWithDefaults instantiates a new ListLocationsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListLocationsResponseWithDefaults() *ListLocationsResponse {
	this := ListLocationsResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ListLocationsResponse) GetData() []Location {
	if o == nil || o.Data == nil {
		var ret []Location
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListLocationsResponse) GetDataOk() ([]Location, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ListLocationsResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []Location and assigns it to the Data field.
func (o *ListLocationsResponse) SetData(v []Location) {
	o.Data = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ListLocationsResponse) GetMeta() ResponseMetadata {
	if o == nil || o.Meta == nil {
		var ret ResponseMetadata
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListLocationsResponse) GetMetaOk() (*ResponseMetadata, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ListLocationsResponse) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given ResponseMetadata and assigns it to the Meta field.
func (o *ListLocationsResponse) SetMeta(v ResponseMetadata) {
	o.Meta = &v
}

func (o ListLocationsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	return json.Marshal(toSerialize)
}

type NullableListLocationsResponse struct {
	value *ListLocationsResponse
	isSet bool
}

func (v NullableListLocationsResponse) Get() *ListLocationsResponse {
	return v.value
}

func (v *NullableListLocationsResponse) Set(val *ListLocationsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListLocationsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListLocationsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListLocationsResponse(val *ListLocationsResponse) *NullableListLocationsResponse {
	return &NullableListLocationsResponse{value: val, isSet: true}
}

func (v NullableListLocationsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListLocationsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


