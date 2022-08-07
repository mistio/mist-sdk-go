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

// GoogleCloudRequest struct for GoogleCloudRequest
type GoogleCloudRequest struct {
	Provider string `json:"provider"`
	Credentials GoogleCredentials `json:"credentials"`
	Features *GoogleCloudFeatures `json:"features,omitempty"`
}

// NewGoogleCloudRequest instantiates a new GoogleCloudRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGoogleCloudRequest(provider string, credentials GoogleCredentials) *GoogleCloudRequest {
	this := GoogleCloudRequest{}
	this.Provider = provider
	this.Credentials = credentials
	return &this
}

// NewGoogleCloudRequestWithDefaults instantiates a new GoogleCloudRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGoogleCloudRequestWithDefaults() *GoogleCloudRequest {
	this := GoogleCloudRequest{}
	return &this
}

// GetProvider returns the Provider field value
func (o *GoogleCloudRequest) GetProvider() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *GoogleCloudRequest) GetProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value
func (o *GoogleCloudRequest) SetProvider(v string) {
	o.Provider = v
}

// GetCredentials returns the Credentials field value
func (o *GoogleCloudRequest) GetCredentials() GoogleCredentials {
	if o == nil {
		var ret GoogleCredentials
		return ret
	}

	return o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value
// and a boolean to check if the value has been set.
func (o *GoogleCloudRequest) GetCredentialsOk() (*GoogleCredentials, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Credentials, true
}

// SetCredentials sets field value
func (o *GoogleCloudRequest) SetCredentials(v GoogleCredentials) {
	o.Credentials = v
}

// GetFeatures returns the Features field value if set, zero value otherwise.
func (o *GoogleCloudRequest) GetFeatures() GoogleCloudFeatures {
	if o == nil || o.Features == nil {
		var ret GoogleCloudFeatures
		return ret
	}
	return *o.Features
}

// GetFeaturesOk returns a tuple with the Features field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleCloudRequest) GetFeaturesOk() (*GoogleCloudFeatures, bool) {
	if o == nil || o.Features == nil {
		return nil, false
	}
	return o.Features, true
}

// HasFeatures returns a boolean if a field has been set.
func (o *GoogleCloudRequest) HasFeatures() bool {
	if o != nil && o.Features != nil {
		return true
	}

	return false
}

// SetFeatures gets a reference to the given GoogleCloudFeatures and assigns it to the Features field.
func (o *GoogleCloudRequest) SetFeatures(v GoogleCloudFeatures) {
	o.Features = &v
}

func (o GoogleCloudRequest) MarshalJSON() ([]byte, error) {
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

type NullableGoogleCloudRequest struct {
	value *GoogleCloudRequest
	isSet bool
}

func (v NullableGoogleCloudRequest) Get() *GoogleCloudRequest {
	return v.value
}

func (v *NullableGoogleCloudRequest) Set(val *GoogleCloudRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGoogleCloudRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGoogleCloudRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGoogleCloudRequest(val *GoogleCloudRequest) *NullableGoogleCloudRequest {
	return &NullableGoogleCloudRequest{value: val, isSet: true}
}

func (v NullableGoogleCloudRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGoogleCloudRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


