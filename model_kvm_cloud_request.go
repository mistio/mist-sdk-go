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

// KvmCloudRequest struct for KvmCloudRequest
type KvmCloudRequest struct {
	Provider string `json:"provider"`
	Credentials KvmCredentials `json:"credentials"`
	Features *KvmCloudFeatures `json:"features,omitempty"`
}

// NewKvmCloudRequest instantiates a new KvmCloudRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKvmCloudRequest(provider string, credentials KvmCredentials) *KvmCloudRequest {
	this := KvmCloudRequest{}
	this.Provider = provider
	this.Credentials = credentials
	return &this
}

// NewKvmCloudRequestWithDefaults instantiates a new KvmCloudRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKvmCloudRequestWithDefaults() *KvmCloudRequest {
	this := KvmCloudRequest{}
	return &this
}

// GetProvider returns the Provider field value
func (o *KvmCloudRequest) GetProvider() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *KvmCloudRequest) GetProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value
func (o *KvmCloudRequest) SetProvider(v string) {
	o.Provider = v
}

// GetCredentials returns the Credentials field value
func (o *KvmCloudRequest) GetCredentials() KvmCredentials {
	if o == nil {
		var ret KvmCredentials
		return ret
	}

	return o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value
// and a boolean to check if the value has been set.
func (o *KvmCloudRequest) GetCredentialsOk() (*KvmCredentials, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Credentials, true
}

// SetCredentials sets field value
func (o *KvmCloudRequest) SetCredentials(v KvmCredentials) {
	o.Credentials = v
}

// GetFeatures returns the Features field value if set, zero value otherwise.
func (o *KvmCloudRequest) GetFeatures() KvmCloudFeatures {
	if o == nil || o.Features == nil {
		var ret KvmCloudFeatures
		return ret
	}
	return *o.Features
}

// GetFeaturesOk returns a tuple with the Features field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KvmCloudRequest) GetFeaturesOk() (*KvmCloudFeatures, bool) {
	if o == nil || o.Features == nil {
		return nil, false
	}
	return o.Features, true
}

// HasFeatures returns a boolean if a field has been set.
func (o *KvmCloudRequest) HasFeatures() bool {
	if o != nil && o.Features != nil {
		return true
	}

	return false
}

// SetFeatures gets a reference to the given KvmCloudFeatures and assigns it to the Features field.
func (o *KvmCloudRequest) SetFeatures(v KvmCloudFeatures) {
	o.Features = &v
}

func (o KvmCloudRequest) MarshalJSON() ([]byte, error) {
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

type NullableKvmCloudRequest struct {
	value *KvmCloudRequest
	isSet bool
}

func (v NullableKvmCloudRequest) Get() *KvmCloudRequest {
	return v.value
}

func (v *NullableKvmCloudRequest) Set(val *KvmCloudRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableKvmCloudRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableKvmCloudRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKvmCloudRequest(val *KvmCloudRequest) *NullableKvmCloudRequest {
	return &NullableKvmCloudRequest{value: val, isSet: true}
}

func (v NullableKvmCloudRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKvmCloudRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


