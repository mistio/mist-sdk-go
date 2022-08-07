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

// VultrCloudRequest struct for VultrCloudRequest
type VultrCloudRequest struct {
	Provider string `json:"provider"`
	Credentials VultrCredentials `json:"credentials"`
	Features *VultrCloudFeatures `json:"features,omitempty"`
}

// NewVultrCloudRequest instantiates a new VultrCloudRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVultrCloudRequest(provider string, credentials VultrCredentials) *VultrCloudRequest {
	this := VultrCloudRequest{}
	this.Provider = provider
	this.Credentials = credentials
	return &this
}

// NewVultrCloudRequestWithDefaults instantiates a new VultrCloudRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVultrCloudRequestWithDefaults() *VultrCloudRequest {
	this := VultrCloudRequest{}
	return &this
}

// GetProvider returns the Provider field value
func (o *VultrCloudRequest) GetProvider() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *VultrCloudRequest) GetProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value
func (o *VultrCloudRequest) SetProvider(v string) {
	o.Provider = v
}

// GetCredentials returns the Credentials field value
func (o *VultrCloudRequest) GetCredentials() VultrCredentials {
	if o == nil {
		var ret VultrCredentials
		return ret
	}

	return o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value
// and a boolean to check if the value has been set.
func (o *VultrCloudRequest) GetCredentialsOk() (*VultrCredentials, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Credentials, true
}

// SetCredentials sets field value
func (o *VultrCloudRequest) SetCredentials(v VultrCredentials) {
	o.Credentials = v
}

// GetFeatures returns the Features field value if set, zero value otherwise.
func (o *VultrCloudRequest) GetFeatures() VultrCloudFeatures {
	if o == nil || o.Features == nil {
		var ret VultrCloudFeatures
		return ret
	}
	return *o.Features
}

// GetFeaturesOk returns a tuple with the Features field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VultrCloudRequest) GetFeaturesOk() (*VultrCloudFeatures, bool) {
	if o == nil || o.Features == nil {
		return nil, false
	}
	return o.Features, true
}

// HasFeatures returns a boolean if a field has been set.
func (o *VultrCloudRequest) HasFeatures() bool {
	if o != nil && o.Features != nil {
		return true
	}

	return false
}

// SetFeatures gets a reference to the given VultrCloudFeatures and assigns it to the Features field.
func (o *VultrCloudRequest) SetFeatures(v VultrCloudFeatures) {
	o.Features = &v
}

func (o VultrCloudRequest) MarshalJSON() ([]byte, error) {
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

type NullableVultrCloudRequest struct {
	value *VultrCloudRequest
	isSet bool
}

func (v NullableVultrCloudRequest) Get() *VultrCloudRequest {
	return v.value
}

func (v *NullableVultrCloudRequest) Set(val *VultrCloudRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableVultrCloudRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableVultrCloudRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVultrCloudRequest(val *VultrCloudRequest) *NullableVultrCloudRequest {
	return &NullableVultrCloudRequest{value: val, isSet: true}
}

func (v NullableVultrCloudRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVultrCloudRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


