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

// KubernetesCloudFeatures struct for KubernetesCloudFeatures
type KubernetesCloudFeatures struct {
	Container *bool `json:"container,omitempty"`
}

// NewKubernetesCloudFeatures instantiates a new KubernetesCloudFeatures object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesCloudFeatures() *KubernetesCloudFeatures {
	this := KubernetesCloudFeatures{}
	var container bool = false
	this.Container = &container
	return &this
}

// NewKubernetesCloudFeaturesWithDefaults instantiates a new KubernetesCloudFeatures object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesCloudFeaturesWithDefaults() *KubernetesCloudFeatures {
	this := KubernetesCloudFeatures{}
	var container bool = false
	this.Container = &container
	return &this
}

// GetContainer returns the Container field value if set, zero value otherwise.
func (o *KubernetesCloudFeatures) GetContainer() bool {
	if o == nil || o.Container == nil {
		var ret bool
		return ret
	}
	return *o.Container
}

// GetContainerOk returns a tuple with the Container field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesCloudFeatures) GetContainerOk() (*bool, bool) {
	if o == nil || o.Container == nil {
		return nil, false
	}
	return o.Container, true
}

// HasContainer returns a boolean if a field has been set.
func (o *KubernetesCloudFeatures) HasContainer() bool {
	if o != nil && o.Container != nil {
		return true
	}

	return false
}

// SetContainer gets a reference to the given bool and assigns it to the Container field.
func (o *KubernetesCloudFeatures) SetContainer(v bool) {
	o.Container = &v
}

func (o KubernetesCloudFeatures) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Container != nil {
		toSerialize["container"] = o.Container
	}
	return json.Marshal(toSerialize)
}

type NullableKubernetesCloudFeatures struct {
	value *KubernetesCloudFeatures
	isSet bool
}

func (v NullableKubernetesCloudFeatures) Get() *KubernetesCloudFeatures {
	return v.value
}

func (v *NullableKubernetesCloudFeatures) Set(val *KubernetesCloudFeatures) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesCloudFeatures) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesCloudFeatures) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesCloudFeatures(val *KubernetesCloudFeatures) *NullableKubernetesCloudFeatures {
	return &NullableKubernetesCloudFeatures{value: val, isSet: true}
}

func (v NullableKubernetesCloudFeatures) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesCloudFeatures) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


