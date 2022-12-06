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

// VsphereCloudRequest struct for VsphereCloudRequest
type VsphereCloudRequest struct {
	Provider string `json:"provider"`
	Credentials VsphereCredentials `json:"credentials"`
	Features *VsphereCloudFeatures `json:"features,omitempty"`
}

// NewVsphereCloudRequest instantiates a new VsphereCloudRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVsphereCloudRequest(provider string, credentials VsphereCredentials) *VsphereCloudRequest {
	this := VsphereCloudRequest{}
	this.Provider = provider
	this.Credentials = credentials
	return &this
}

// NewVsphereCloudRequestWithDefaults instantiates a new VsphereCloudRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVsphereCloudRequestWithDefaults() *VsphereCloudRequest {
	this := VsphereCloudRequest{}
	return &this
}

// GetProvider returns the Provider field value
func (o *VsphereCloudRequest) GetProvider() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *VsphereCloudRequest) GetProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value
func (o *VsphereCloudRequest) SetProvider(v string) {
	o.Provider = v
}

// GetCredentials returns the Credentials field value
func (o *VsphereCloudRequest) GetCredentials() VsphereCredentials {
	if o == nil {
		var ret VsphereCredentials
		return ret
	}

	return o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value
// and a boolean to check if the value has been set.
func (o *VsphereCloudRequest) GetCredentialsOk() (*VsphereCredentials, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Credentials, true
}

// SetCredentials sets field value
func (o *VsphereCloudRequest) SetCredentials(v VsphereCredentials) {
	o.Credentials = v
}

// GetFeatures returns the Features field value if set, zero value otherwise.
func (o *VsphereCloudRequest) GetFeatures() VsphereCloudFeatures {
	if o == nil || o.Features == nil {
		var ret VsphereCloudFeatures
		return ret
	}
	return *o.Features
}

// GetFeaturesOk returns a tuple with the Features field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VsphereCloudRequest) GetFeaturesOk() (*VsphereCloudFeatures, bool) {
	if o == nil || o.Features == nil {
		return nil, false
	}
	return o.Features, true
}

// HasFeatures returns a boolean if a field has been set.
func (o *VsphereCloudRequest) HasFeatures() bool {
	if o != nil && o.Features != nil {
		return true
	}

	return false
}

// SetFeatures gets a reference to the given VsphereCloudFeatures and assigns it to the Features field.
func (o *VsphereCloudRequest) SetFeatures(v VsphereCloudFeatures) {
	o.Features = &v
}

func (o VsphereCloudRequest) MarshalJSON() ([]byte, error) {
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

type NullableVsphereCloudRequest struct {
	value *VsphereCloudRequest
	isSet bool
}

func (v NullableVsphereCloudRequest) Get() *VsphereCloudRequest {
	return v.value
}

func (v *NullableVsphereCloudRequest) Set(val *VsphereCloudRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableVsphereCloudRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableVsphereCloudRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVsphereCloudRequest(val *VsphereCloudRequest) *NullableVsphereCloudRequest {
	return &NullableVsphereCloudRequest{value: val, isSet: true}
}

func (v NullableVsphereCloudRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVsphereCloudRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


