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

// TagResourcesRequest struct for TagResourcesRequest
type TagResourcesRequest struct {
	Resources *[]TagResourcesRequestResources `json:"resources,omitempty"`
}

// NewTagResourcesRequest instantiates a new TagResourcesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTagResourcesRequest() *TagResourcesRequest {
	this := TagResourcesRequest{}
	return &this
}

// NewTagResourcesRequestWithDefaults instantiates a new TagResourcesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTagResourcesRequestWithDefaults() *TagResourcesRequest {
	this := TagResourcesRequest{}
	return &this
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *TagResourcesRequest) GetResources() []TagResourcesRequestResources {
	if o == nil || o.Resources == nil {
		var ret []TagResourcesRequestResources
		return ret
	}
	return *o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagResourcesRequest) GetResourcesOk() (*[]TagResourcesRequestResources, bool) {
	if o == nil || o.Resources == nil {
		return nil, false
	}
	return o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *TagResourcesRequest) HasResources() bool {
	if o != nil && o.Resources != nil {
		return true
	}

	return false
}

// SetResources gets a reference to the given []TagResourcesRequestResources and assigns it to the Resources field.
func (o *TagResourcesRequest) SetResources(v []TagResourcesRequestResources) {
	o.Resources = &v
}

func (o TagResourcesRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Resources != nil {
		toSerialize["resources"] = o.Resources
	}
	return json.Marshal(toSerialize)
}

type NullableTagResourcesRequest struct {
	value *TagResourcesRequest
	isSet bool
}

func (v NullableTagResourcesRequest) Get() *TagResourcesRequest {
	return v.value
}

func (v *NullableTagResourcesRequest) Set(val *TagResourcesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTagResourcesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTagResourcesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTagResourcesRequest(val *TagResourcesRequest) *NullableTagResourcesRequest {
	return &NullableTagResourcesRequest{value: val, isSet: true}
}

func (v NullableTagResourcesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTagResourcesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


