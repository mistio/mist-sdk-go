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

// RackspaceCloudRequest struct for RackspaceCloudRequest
type RackspaceCloudRequest struct {
	Provider string `json:"provider"`
	Credentials RackspaceCredentials `json:"credentials"`
	Features *CloudFeatures `json:"features,omitempty"`
}

// NewRackspaceCloudRequest instantiates a new RackspaceCloudRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRackspaceCloudRequest(provider string, credentials RackspaceCredentials, ) *RackspaceCloudRequest {
	this := RackspaceCloudRequest{}
	this.Provider = provider
	this.Credentials = credentials
	return &this
}

// NewRackspaceCloudRequestWithDefaults instantiates a new RackspaceCloudRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRackspaceCloudRequestWithDefaults() *RackspaceCloudRequest {
	this := RackspaceCloudRequest{}
	return &this
}

// GetProvider returns the Provider field value
func (o *RackspaceCloudRequest) GetProvider() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *RackspaceCloudRequest) GetProviderOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value
func (o *RackspaceCloudRequest) SetProvider(v string) {
	o.Provider = v
}

// GetCredentials returns the Credentials field value
func (o *RackspaceCloudRequest) GetCredentials() RackspaceCredentials {
	if o == nil  {
		var ret RackspaceCredentials
		return ret
	}

	return o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value
// and a boolean to check if the value has been set.
func (o *RackspaceCloudRequest) GetCredentialsOk() (*RackspaceCredentials, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Credentials, true
}

// SetCredentials sets field value
func (o *RackspaceCloudRequest) SetCredentials(v RackspaceCredentials) {
	o.Credentials = v
}

// GetFeatures returns the Features field value if set, zero value otherwise.
func (o *RackspaceCloudRequest) GetFeatures() CloudFeatures {
	if o == nil || o.Features == nil {
		var ret CloudFeatures
		return ret
	}
	return *o.Features
}

// GetFeaturesOk returns a tuple with the Features field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RackspaceCloudRequest) GetFeaturesOk() (*CloudFeatures, bool) {
	if o == nil || o.Features == nil {
		return nil, false
	}
	return o.Features, true
}

// HasFeatures returns a boolean if a field has been set.
func (o *RackspaceCloudRequest) HasFeatures() bool {
	if o != nil && o.Features != nil {
		return true
	}

	return false
}

// SetFeatures gets a reference to the given CloudFeatures and assigns it to the Features field.
func (o *RackspaceCloudRequest) SetFeatures(v CloudFeatures) {
	o.Features = &v
}

func (o RackspaceCloudRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["provider"] = o.Provider
	}
	if true {
		toSerialize["credentials"] = o.Credentials
	}
	if o.Features != nil {
		toSerialize["features"] = o.Features
	}
	return json.Marshal(toSerialize)
}

type NullableRackspaceCloudRequest struct {
	value *RackspaceCloudRequest
	isSet bool
}

func (v NullableRackspaceCloudRequest) Get() *RackspaceCloudRequest {
	return v.value
}

func (v *NullableRackspaceCloudRequest) Set(val *RackspaceCloudRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRackspaceCloudRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRackspaceCloudRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRackspaceCloudRequest(val *RackspaceCloudRequest) *NullableRackspaceCloudRequest {
	return &NullableRackspaceCloudRequest{value: val, isSet: true}
}

func (v NullableRackspaceCloudRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRackspaceCloudRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


