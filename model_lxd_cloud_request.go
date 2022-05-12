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

// LxdCloudRequest struct for LxdCloudRequest
type LxdCloudRequest struct {
	Provider string `json:"provider"`
	Credentials LxdCredentials `json:"credentials"`
	Features *CloudFeatures `json:"features,omitempty"`
}

// NewLxdCloudRequest instantiates a new LxdCloudRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLxdCloudRequest(provider string, credentials LxdCredentials) *LxdCloudRequest {
	this := LxdCloudRequest{}
	this.Provider = provider
	this.Credentials = credentials
	return &this
}

// NewLxdCloudRequestWithDefaults instantiates a new LxdCloudRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLxdCloudRequestWithDefaults() *LxdCloudRequest {
	this := LxdCloudRequest{}
	return &this
}

// GetProvider returns the Provider field value
func (o *LxdCloudRequest) GetProvider() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *LxdCloudRequest) GetProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value
func (o *LxdCloudRequest) SetProvider(v string) {
	o.Provider = v
}

// GetCredentials returns the Credentials field value
func (o *LxdCloudRequest) GetCredentials() LxdCredentials {
	if o == nil {
		var ret LxdCredentials
		return ret
	}

	return o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value
// and a boolean to check if the value has been set.
func (o *LxdCloudRequest) GetCredentialsOk() (*LxdCredentials, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Credentials, true
}

// SetCredentials sets field value
func (o *LxdCloudRequest) SetCredentials(v LxdCredentials) {
	o.Credentials = v
}

// GetFeatures returns the Features field value if set, zero value otherwise.
func (o *LxdCloudRequest) GetFeatures() CloudFeatures {
	if o == nil || o.Features == nil {
		var ret CloudFeatures
		return ret
	}
	return *o.Features
}

// GetFeaturesOk returns a tuple with the Features field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LxdCloudRequest) GetFeaturesOk() (*CloudFeatures, bool) {
	if o == nil || o.Features == nil {
		return nil, false
	}
	return o.Features, true
}

// HasFeatures returns a boolean if a field has been set.
func (o *LxdCloudRequest) HasFeatures() bool {
	if o != nil && o.Features != nil {
		return true
	}

	return false
}

// SetFeatures gets a reference to the given CloudFeatures and assigns it to the Features field.
func (o *LxdCloudRequest) SetFeatures(v CloudFeatures) {
	o.Features = &v
}

func (o LxdCloudRequest) MarshalJSON() ([]byte, error) {
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

type NullableLxdCloudRequest struct {
	value *LxdCloudRequest
	isSet bool
}

func (v NullableLxdCloudRequest) Get() *LxdCloudRequest {
	return v.value
}

func (v *NullableLxdCloudRequest) Set(val *LxdCloudRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableLxdCloudRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableLxdCloudRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLxdCloudRequest(val *LxdCloudRequest) *NullableLxdCloudRequest {
	return &NullableLxdCloudRequest{value: val, isSet: true}
}

func (v NullableLxdCloudRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLxdCloudRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


