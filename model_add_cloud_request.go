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

// AddCloudRequest struct for AddCloudRequest
type AddCloudRequest struct {
	Name string `json:"name"`
	Provider SupportedProviders `json:"provider"`
	Credentials map[string]interface{} `json:"credentials"`
	Features *CloudFeatures `json:"features,omitempty"`
}

// NewAddCloudRequest instantiates a new AddCloudRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddCloudRequest(name string, provider SupportedProviders, credentials map[string]interface{}) *AddCloudRequest {
	this := AddCloudRequest{}
	this.Name = name
	this.Provider = provider
	this.Credentials = credentials
	return &this
}

// NewAddCloudRequestWithDefaults instantiates a new AddCloudRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddCloudRequestWithDefaults() *AddCloudRequest {
	this := AddCloudRequest{}
	return &this
}

// GetName returns the Name field value
func (o *AddCloudRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AddCloudRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AddCloudRequest) SetName(v string) {
	o.Name = v
}

// GetProvider returns the Provider field value
func (o *AddCloudRequest) GetProvider() SupportedProviders {
	if o == nil {
		var ret SupportedProviders
		return ret
	}

	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *AddCloudRequest) GetProviderOk() (*SupportedProviders, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value
func (o *AddCloudRequest) SetProvider(v SupportedProviders) {
	o.Provider = v
}

// GetCredentials returns the Credentials field value
func (o *AddCloudRequest) GetCredentials() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value
// and a boolean to check if the value has been set.
func (o *AddCloudRequest) GetCredentialsOk() (map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Credentials, true
}

// SetCredentials sets field value
func (o *AddCloudRequest) SetCredentials(v map[string]interface{}) {
	o.Credentials = v
}

// GetFeatures returns the Features field value if set, zero value otherwise.
func (o *AddCloudRequest) GetFeatures() CloudFeatures {
	if o == nil || o.Features == nil {
		var ret CloudFeatures
		return ret
	}
	return *o.Features
}

// GetFeaturesOk returns a tuple with the Features field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddCloudRequest) GetFeaturesOk() (*CloudFeatures, bool) {
	if o == nil || o.Features == nil {
		return nil, false
	}
	return o.Features, true
}

// HasFeatures returns a boolean if a field has been set.
func (o *AddCloudRequest) HasFeatures() bool {
	if o != nil && o.Features != nil {
		return true
	}

	return false
}

// SetFeatures gets a reference to the given CloudFeatures and assigns it to the Features field.
func (o *AddCloudRequest) SetFeatures(v CloudFeatures) {
	o.Features = &v
}

func (o AddCloudRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
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

type NullableAddCloudRequest struct {
	value *AddCloudRequest
	isSet bool
}

func (v NullableAddCloudRequest) Get() *AddCloudRequest {
	return v.value
}

func (v *NullableAddCloudRequest) Set(val *AddCloudRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddCloudRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddCloudRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddCloudRequest(val *AddCloudRequest) *NullableAddCloudRequest {
	return &NullableAddCloudRequest{value: val, isSet: true}
}

func (v NullableAddCloudRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddCloudRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


