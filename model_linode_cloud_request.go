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

// LinodeCloudRequest struct for LinodeCloudRequest
type LinodeCloudRequest struct {
	Provider string `json:"provider"`
	Credentials LinodeCredentials `json:"credentials"`
	Features *CloudFeatures `json:"features,omitempty"`
}

// NewLinodeCloudRequest instantiates a new LinodeCloudRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinodeCloudRequest(provider string, credentials LinodeCredentials, ) *LinodeCloudRequest {
	this := LinodeCloudRequest{}
	this.Provider = provider
	this.Credentials = credentials
	return &this
}

// NewLinodeCloudRequestWithDefaults instantiates a new LinodeCloudRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinodeCloudRequestWithDefaults() *LinodeCloudRequest {
	this := LinodeCloudRequest{}
	return &this
}

// GetProvider returns the Provider field value
func (o *LinodeCloudRequest) GetProvider() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *LinodeCloudRequest) GetProviderOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value
func (o *LinodeCloudRequest) SetProvider(v string) {
	o.Provider = v
}

// GetCredentials returns the Credentials field value
func (o *LinodeCloudRequest) GetCredentials() LinodeCredentials {
	if o == nil  {
		var ret LinodeCredentials
		return ret
	}

	return o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value
// and a boolean to check if the value has been set.
func (o *LinodeCloudRequest) GetCredentialsOk() (*LinodeCredentials, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Credentials, true
}

// SetCredentials sets field value
func (o *LinodeCloudRequest) SetCredentials(v LinodeCredentials) {
	o.Credentials = v
}

// GetFeatures returns the Features field value if set, zero value otherwise.
func (o *LinodeCloudRequest) GetFeatures() CloudFeatures {
	if o == nil || o.Features == nil {
		var ret CloudFeatures
		return ret
	}
	return *o.Features
}

// GetFeaturesOk returns a tuple with the Features field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinodeCloudRequest) GetFeaturesOk() (*CloudFeatures, bool) {
	if o == nil || o.Features == nil {
		return nil, false
	}
	return o.Features, true
}

// HasFeatures returns a boolean if a field has been set.
func (o *LinodeCloudRequest) HasFeatures() bool {
	if o != nil && o.Features != nil {
		return true
	}

	return false
}

// SetFeatures gets a reference to the given CloudFeatures and assigns it to the Features field.
func (o *LinodeCloudRequest) SetFeatures(v CloudFeatures) {
	o.Features = &v
}

func (o LinodeCloudRequest) MarshalJSON() ([]byte, error) {
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

type NullableLinodeCloudRequest struct {
	value *LinodeCloudRequest
	isSet bool
}

func (v NullableLinodeCloudRequest) Get() *LinodeCloudRequest {
	return v.value
}

func (v *NullableLinodeCloudRequest) Set(val *LinodeCloudRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableLinodeCloudRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableLinodeCloudRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinodeCloudRequest(val *LinodeCloudRequest) *NullableLinodeCloudRequest {
	return &NullableLinodeCloudRequest{value: val, isSet: true}
}

func (v NullableLinodeCloudRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinodeCloudRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


