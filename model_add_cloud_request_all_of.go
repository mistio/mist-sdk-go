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

// AddCloudRequestAllOf struct for AddCloudRequestAllOf
type AddCloudRequestAllOf struct {
	// The name of the cloud to add
	Name string `json:"name"`
	// The provider of the cloud
	Provider string `json:"provider"`
}

// NewAddCloudRequestAllOf instantiates a new AddCloudRequestAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddCloudRequestAllOf(name string, provider string, ) *AddCloudRequestAllOf {
	this := AddCloudRequestAllOf{}
	this.Name = name
	this.Provider = provider
	return &this
}

// NewAddCloudRequestAllOfWithDefaults instantiates a new AddCloudRequestAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddCloudRequestAllOfWithDefaults() *AddCloudRequestAllOf {
	this := AddCloudRequestAllOf{}
	return &this
}

// GetName returns the Name field value
func (o *AddCloudRequestAllOf) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AddCloudRequestAllOf) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AddCloudRequestAllOf) SetName(v string) {
	o.Name = v
}

// GetProvider returns the Provider field value
func (o *AddCloudRequestAllOf) GetProvider() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *AddCloudRequestAllOf) GetProviderOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value
func (o *AddCloudRequestAllOf) SetProvider(v string) {
	o.Provider = v
}

func (o AddCloudRequestAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["provider"] = o.Provider
	}
	return json.Marshal(toSerialize)
}

type NullableAddCloudRequestAllOf struct {
	value *AddCloudRequestAllOf
	isSet bool
}

func (v NullableAddCloudRequestAllOf) Get() *AddCloudRequestAllOf {
	return v.value
}

func (v *NullableAddCloudRequestAllOf) Set(val *AddCloudRequestAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAddCloudRequestAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAddCloudRequestAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddCloudRequestAllOf(val *AddCloudRequestAllOf) *NullableAddCloudRequestAllOf {
	return &NullableAddCloudRequestAllOf{value: val, isSet: true}
}

func (v NullableAddCloudRequestAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddCloudRequestAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


