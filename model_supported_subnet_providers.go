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
	"fmt"
)

// SupportedSubnetProviders the model 'SupportedSubnetProviders'
type SupportedSubnetProviders string

// List of SupportedSubnetProviders
const (
	AMAZON SupportedSubnetProviders = "amazon"
	AZURE SupportedSubnetProviders = "azure"
	GOOGLE SupportedSubnetProviders = "google"
	ALIBABA SupportedSubnetProviders = "alibaba"
	OPENSTACK SupportedSubnetProviders = "openstack"
	VEXXHOST SupportedSubnetProviders = "vexxhost"
)

func (v *SupportedSubnetProviders) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SupportedSubnetProviders(value)
	for _, existing := range []SupportedSubnetProviders{ "amazon", "azure", "google", "alibaba", "openstack", "vexxhost",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SupportedSubnetProviders", value)
}

// Ptr returns reference to SupportedSubnetProviders value
func (v SupportedSubnetProviders) Ptr() *SupportedSubnetProviders {
	return &v
}

type NullableSupportedSubnetProviders struct {
	value *SupportedSubnetProviders
	isSet bool
}

func (v NullableSupportedSubnetProviders) Get() *SupportedSubnetProviders {
	return v.value
}

func (v *NullableSupportedSubnetProviders) Set(val *SupportedSubnetProviders) {
	v.value = val
	v.isSet = true
}

func (v NullableSupportedSubnetProviders) IsSet() bool {
	return v.isSet
}

func (v *NullableSupportedSubnetProviders) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSupportedSubnetProviders(val *SupportedSubnetProviders) *NullableSupportedSubnetProviders {
	return &NullableSupportedSubnetProviders{value: val, isSet: true}
}

func (v NullableSupportedSubnetProviders) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSupportedSubnetProviders) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

