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

// CreateClusterResponse struct for CreateClusterResponse
type CreateClusterResponse struct {
	JobId *string `json:"jobId,omitempty"`
}

// NewCreateClusterResponse instantiates a new CreateClusterResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateClusterResponse() *CreateClusterResponse {
	this := CreateClusterResponse{}
	return &this
}

// NewCreateClusterResponseWithDefaults instantiates a new CreateClusterResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateClusterResponseWithDefaults() *CreateClusterResponse {
	this := CreateClusterResponse{}
	return &this
}

// GetJobId returns the JobId field value if set, zero value otherwise.
func (o *CreateClusterResponse) GetJobId() string {
	if o == nil || o.JobId == nil {
		var ret string
		return ret
	}
	return *o.JobId
}

// GetJobIdOk returns a tuple with the JobId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateClusterResponse) GetJobIdOk() (*string, bool) {
	if o == nil || o.JobId == nil {
		return nil, false
	}
	return o.JobId, true
}

// HasJobId returns a boolean if a field has been set.
func (o *CreateClusterResponse) HasJobId() bool {
	if o != nil && o.JobId != nil {
		return true
	}

	return false
}

// SetJobId gets a reference to the given string and assigns it to the JobId field.
func (o *CreateClusterResponse) SetJobId(v string) {
	o.JobId = &v
}

func (o CreateClusterResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.JobId != nil {
		toSerialize["jobId"] = o.JobId
	}
	return json.Marshal(toSerialize)
}

type NullableCreateClusterResponse struct {
	value *CreateClusterResponse
	isSet bool
}

func (v NullableCreateClusterResponse) Get() *CreateClusterResponse {
	return v.value
}

func (v *NullableCreateClusterResponse) Set(val *CreateClusterResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateClusterResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateClusterResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateClusterResponse(val *CreateClusterResponse) *NullableCreateClusterResponse {
	return &NullableCreateClusterResponse{value: val, isSet: true}
}

func (v NullableCreateClusterResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateClusterResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


