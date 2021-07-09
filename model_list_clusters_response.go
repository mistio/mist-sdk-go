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

// ListClustersResponse struct for ListClustersResponse
type ListClustersResponse struct {
	Data *[]Cluster `json:"data,omitempty"`
	Meta *ResponseMetadata `json:"meta,omitempty"`
}

// NewListClustersResponse instantiates a new ListClustersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListClustersResponse() *ListClustersResponse {
	this := ListClustersResponse{}
	return &this
}

// NewListClustersResponseWithDefaults instantiates a new ListClustersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListClustersResponseWithDefaults() *ListClustersResponse {
	this := ListClustersResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ListClustersResponse) GetData() []Cluster {
	if o == nil || o.Data == nil {
		var ret []Cluster
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListClustersResponse) GetDataOk() (*[]Cluster, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ListClustersResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []Cluster and assigns it to the Data field.
func (o *ListClustersResponse) SetData(v []Cluster) {
	o.Data = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ListClustersResponse) GetMeta() ResponseMetadata {
	if o == nil || o.Meta == nil {
		var ret ResponseMetadata
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListClustersResponse) GetMetaOk() (*ResponseMetadata, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ListClustersResponse) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given ResponseMetadata and assigns it to the Meta field.
func (o *ListClustersResponse) SetMeta(v ResponseMetadata) {
	o.Meta = &v
}

func (o ListClustersResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	return json.Marshal(toSerialize)
}

type NullableListClustersResponse struct {
	value *ListClustersResponse
	isSet bool
}

func (v NullableListClustersResponse) Get() *ListClustersResponse {
	return v.value
}

func (v *NullableListClustersResponse) Set(val *ListClustersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListClustersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListClustersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListClustersResponse(val *ListClustersResponse) *NullableListClustersResponse {
	return &NullableListClustersResponse{value: val, isSet: true}
}

func (v NullableListClustersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListClustersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


