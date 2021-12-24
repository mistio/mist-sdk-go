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

// AzureCloudRequest struct for AzureCloudRequest
type AzureCloudRequest struct {
	Credentials AzureCredentials `json:"credentials"`
	Features *CloudFeatures `json:"features,omitempty"`
	Provider string `json:"provider"`
}

// NewAzureCloudRequest instantiates a new AzureCloudRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureCloudRequest(credentials AzureCredentials, provider string, ) *AzureCloudRequest {
	this := AzureCloudRequest{}
	this.Credentials = credentials
	this.Provider = provider
	return &this
}

// NewAzureCloudRequestWithDefaults instantiates a new AzureCloudRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureCloudRequestWithDefaults() *AzureCloudRequest {
	this := AzureCloudRequest{}
	return &this
}

// GetCredentials returns the Credentials field value
func (o *AzureCloudRequest) GetCredentials() AzureCredentials {
	if o == nil  {
		var ret AzureCredentials
		return ret
	}

	return o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value
// and a boolean to check if the value has been set.
func (o *AzureCloudRequest) GetCredentialsOk() (*AzureCredentials, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Credentials, true
}

// SetCredentials sets field value
func (o *AzureCloudRequest) SetCredentials(v AzureCredentials) {
	o.Credentials = v
}

// GetFeatures returns the Features field value if set, zero value otherwise.
func (o *AzureCloudRequest) GetFeatures() CloudFeatures {
	if o == nil || o.Features == nil {
		var ret CloudFeatures
		return ret
	}
	return *o.Features
}

// GetFeaturesOk returns a tuple with the Features field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureCloudRequest) GetFeaturesOk() (*CloudFeatures, bool) {
	if o == nil || o.Features == nil {
		return nil, false
	}
	return o.Features, true
}

// HasFeatures returns a boolean if a field has been set.
func (o *AzureCloudRequest) HasFeatures() bool {
	if o != nil && o.Features != nil {
		return true
	}

	return false
}

// SetFeatures gets a reference to the given CloudFeatures and assigns it to the Features field.
func (o *AzureCloudRequest) SetFeatures(v CloudFeatures) {
	o.Features = &v
}

// GetProvider returns the Provider field value
func (o *AzureCloudRequest) GetProvider() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *AzureCloudRequest) GetProviderOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value
func (o *AzureCloudRequest) SetProvider(v string) {
	o.Provider = v
}

func (o AzureCloudRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["credentials"] = o.Credentials
	}
	if o.Features != nil {
		toSerialize["features"] = o.Features
	}
	if true {
		toSerialize["provider"] = o.Provider
	}
	return json.Marshal(toSerialize)
}

type NullableAzureCloudRequest struct {
	value *AzureCloudRequest
	isSet bool
}

func (v NullableAzureCloudRequest) Get() *AzureCloudRequest {
	return v.value
}

func (v *NullableAzureCloudRequest) Set(val *AzureCloudRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureCloudRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureCloudRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureCloudRequest(val *AzureCloudRequest) *NullableAzureCloudRequest {
	return &NullableAzureCloudRequest{value: val, isSet: true}
}

func (v NullableAzureCloudRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureCloudRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


