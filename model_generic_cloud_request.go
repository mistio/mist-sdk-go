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

// GenericCloudRequest struct for GenericCloudRequest
type GenericCloudRequest struct {
	Provider string `json:"provider"`
	Credentials map[string]interface{} `json:"credentials"`
	Features *CloudFeatures `json:"features,omitempty"`
}

// NewGenericCloudRequest instantiates a new GenericCloudRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenericCloudRequest(provider string, credentials map[string]interface{}) *GenericCloudRequest {
	this := GenericCloudRequest{}
	this.Provider = provider
	this.Credentials = credentials
	return &this
}

// NewGenericCloudRequestWithDefaults instantiates a new GenericCloudRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenericCloudRequestWithDefaults() *GenericCloudRequest {
	this := GenericCloudRequest{}
	return &this
}

// GetProvider returns the Provider field value
func (o *GenericCloudRequest) GetProvider() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *GenericCloudRequest) GetProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value
func (o *GenericCloudRequest) SetProvider(v string) {
	o.Provider = v
}

// GetCredentials returns the Credentials field value
func (o *GenericCloudRequest) GetCredentials() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value
// and a boolean to check if the value has been set.
func (o *GenericCloudRequest) GetCredentialsOk() (map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Credentials, true
}

// SetCredentials sets field value
func (o *GenericCloudRequest) SetCredentials(v map[string]interface{}) {
	o.Credentials = v
}

// GetFeatures returns the Features field value if set, zero value otherwise.
func (o *GenericCloudRequest) GetFeatures() CloudFeatures {
	if o == nil || o.Features == nil {
		var ret CloudFeatures
		return ret
	}
	return *o.Features
}

// GetFeaturesOk returns a tuple with the Features field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericCloudRequest) GetFeaturesOk() (*CloudFeatures, bool) {
	if o == nil || o.Features == nil {
		return nil, false
	}
	return o.Features, true
}

// HasFeatures returns a boolean if a field has been set.
func (o *GenericCloudRequest) HasFeatures() bool {
	if o != nil && o.Features != nil {
		return true
	}

	return false
}

// SetFeatures gets a reference to the given CloudFeatures and assigns it to the Features field.
func (o *GenericCloudRequest) SetFeatures(v CloudFeatures) {
	o.Features = &v
}

func (o GenericCloudRequest) MarshalJSON() ([]byte, error) {
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

type NullableGenericCloudRequest struct {
	value *GenericCloudRequest
	isSet bool
}

func (v NullableGenericCloudRequest) Get() *GenericCloudRequest {
	return v.value
}

func (v *NullableGenericCloudRequest) Set(val *GenericCloudRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGenericCloudRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGenericCloudRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenericCloudRequest(val *GenericCloudRequest) *NullableGenericCloudRequest {
	return &NullableGenericCloudRequest{value: val, isSet: true}
}

func (v NullableGenericCloudRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenericCloudRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


