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

// SupportedProviders the model 'SupportedProviders'
type SupportedProviders string

// List of SupportedProviders
const (
	AMAZON SupportedProviders = "amazon"
	AZURE SupportedProviders = "azure"
	GOOGLE SupportedProviders = "google"
	ALIBABA SupportedProviders = "alibaba"
	CLOUDSIGMA SupportedProviders = "cloudsigma"
	EQUINIX SupportedProviders = "equinix"
	IBM SupportedProviders = "ibm"
	DIGITALOCEAN SupportedProviders = "digitalocean"
	LINODE SupportedProviders = "linode"
	RACKSPACE SupportedProviders = "rackspace"
	MAXIHOST SupportedProviders = "maxihost"
	VULTR SupportedProviders = "vultr"
	OPENSTACK SupportedProviders = "openstack"
	ONAPP SupportedProviders = "onapp"
	VSPHERE SupportedProviders = "vsphere"
	VCLOUD SupportedProviders = "vcloud"
	KVM SupportedProviders = "kvm"
	LXD SupportedProviders = "lxd"
	DOCKER SupportedProviders = "docker"
	KUBEVIRT SupportedProviders = "kubevirt"
	KUBERNETES SupportedProviders = "kubernetes"
	OPENSHIFT SupportedProviders = "openshift"
	OTHER SupportedProviders = "other"
)

func (v *SupportedProviders) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SupportedProviders(value)
	for _, existing := range []SupportedProviders{ "amazon", "azure", "google", "alibaba", "cloudsigma", "equinix", "ibm", "digitalocean", "linode", "rackspace", "maxihost", "vultr", "openstack", "onapp", "vsphere", "vcloud", "kvm", "lxd", "docker", "kubevirt", "kubernetes", "openshift", "other",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SupportedProviders", value)
}

// Ptr returns reference to SupportedProviders value
func (v SupportedProviders) Ptr() *SupportedProviders {
	return &v
}

type NullableSupportedProviders struct {
	value *SupportedProviders
	isSet bool
}

func (v NullableSupportedProviders) Get() *SupportedProviders {
	return v.value
}

func (v *NullableSupportedProviders) Set(val *SupportedProviders) {
	v.value = val
	v.isSet = true
}

func (v NullableSupportedProviders) IsSet() bool {
	return v.isSet
}

func (v *NullableSupportedProviders) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSupportedProviders(val *SupportedProviders) *NullableSupportedProviders {
	return &NullableSupportedProviders{value: val, isSet: true}
}

func (v NullableSupportedProviders) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSupportedProviders) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

