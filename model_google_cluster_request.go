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

// GoogleClusterRequest struct for GoogleClusterRequest
type GoogleClusterRequest struct {
	// The name of the location to create the cluster in
	Location string `json:"location"`
	Provider string `json:"provider"`
}

// NewGoogleClusterRequest instantiates a new GoogleClusterRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGoogleClusterRequest(location string, provider string, ) *GoogleClusterRequest {
	this := GoogleClusterRequest{}
	this.Location = location
	this.Provider = provider
	return &this
}

// NewGoogleClusterRequestWithDefaults instantiates a new GoogleClusterRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGoogleClusterRequestWithDefaults() *GoogleClusterRequest {
	this := GoogleClusterRequest{}
	return &this
}

// GetLocation returns the Location field value
func (o *GoogleClusterRequest) GetLocation() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Location
}

// GetLocationOk returns a tuple with the Location field value
// and a boolean to check if the value has been set.
func (o *GoogleClusterRequest) GetLocationOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Location, true
}

// SetLocation sets field value
func (o *GoogleClusterRequest) SetLocation(v string) {
	o.Location = v
}

// GetProvider returns the Provider field value
func (o *GoogleClusterRequest) GetProvider() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *GoogleClusterRequest) GetProviderOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value
func (o *GoogleClusterRequest) SetProvider(v string) {
	o.Provider = v
}

func (o GoogleClusterRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["location"] = o.Location
	}
	if true {
		toSerialize["provider"] = o.Provider
	}
	return json.Marshal(toSerialize)
}

type NullableGoogleClusterRequest struct {
	value *GoogleClusterRequest
	isSet bool
}

func (v NullableGoogleClusterRequest) Get() *GoogleClusterRequest {
	return v.value
}

func (v *NullableGoogleClusterRequest) Set(val *GoogleClusterRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGoogleClusterRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGoogleClusterRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGoogleClusterRequest(val *GoogleClusterRequest) *NullableGoogleClusterRequest {
	return &NullableGoogleClusterRequest{value: val, isSet: true}
}

func (v NullableGoogleClusterRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGoogleClusterRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


