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

// ListUsersResponse struct for ListUsersResponse
type ListUsersResponse struct {
	Data *[]User `json:"data,omitempty"`
	Meta *ResponseMetadata `json:"meta,omitempty"`
}

// NewListUsersResponse instantiates a new ListUsersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListUsersResponse() *ListUsersResponse {
	this := ListUsersResponse{}
	return &this
}

// NewListUsersResponseWithDefaults instantiates a new ListUsersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListUsersResponseWithDefaults() *ListUsersResponse {
	this := ListUsersResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ListUsersResponse) GetData() []User {
	if o == nil || o.Data == nil {
		var ret []User
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListUsersResponse) GetDataOk() (*[]User, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ListUsersResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []User and assigns it to the Data field.
func (o *ListUsersResponse) SetData(v []User) {
	o.Data = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ListUsersResponse) GetMeta() ResponseMetadata {
	if o == nil || o.Meta == nil {
		var ret ResponseMetadata
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListUsersResponse) GetMetaOk() (*ResponseMetadata, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ListUsersResponse) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given ResponseMetadata and assigns it to the Meta field.
func (o *ListUsersResponse) SetMeta(v ResponseMetadata) {
	o.Meta = &v
}

func (o ListUsersResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	return json.Marshal(toSerialize)
}

type NullableListUsersResponse struct {
	value *ListUsersResponse
	isSet bool
}

func (v NullableListUsersResponse) Get() *ListUsersResponse {
	return v.value
}

func (v *NullableListUsersResponse) Set(val *ListUsersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListUsersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListUsersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListUsersResponse(val *ListUsersResponse) *NullableListUsersResponse {
	return &NullableListUsersResponse{value: val, isSet: true}
}

func (v NullableListUsersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListUsersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


