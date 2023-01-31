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

// PatchOrganizationRequest struct for PatchOrganizationRequest
type PatchOrganizationRequest struct {
	// The organization's name
	Name *string `json:"name,omitempty"`
	Vault *PatchOrganizationRequestVault `json:"vault,omitempty"`
}

// NewPatchOrganizationRequest instantiates a new PatchOrganizationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchOrganizationRequest() *PatchOrganizationRequest {
	this := PatchOrganizationRequest{}
	return &this
}

// NewPatchOrganizationRequestWithDefaults instantiates a new PatchOrganizationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchOrganizationRequestWithDefaults() *PatchOrganizationRequest {
	this := PatchOrganizationRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchOrganizationRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchOrganizationRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchOrganizationRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchOrganizationRequest) SetName(v string) {
	o.Name = &v
}

// GetVault returns the Vault field value if set, zero value otherwise.
func (o *PatchOrganizationRequest) GetVault() PatchOrganizationRequestVault {
	if o == nil || o.Vault == nil {
		var ret PatchOrganizationRequestVault
		return ret
	}
	return *o.Vault
}

// GetVaultOk returns a tuple with the Vault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchOrganizationRequest) GetVaultOk() (*PatchOrganizationRequestVault, bool) {
	if o == nil || o.Vault == nil {
		return nil, false
	}
	return o.Vault, true
}

// HasVault returns a boolean if a field has been set.
func (o *PatchOrganizationRequest) HasVault() bool {
	if o != nil && o.Vault != nil {
		return true
	}

	return false
}

// SetVault gets a reference to the given PatchOrganizationRequestVault and assigns it to the Vault field.
func (o *PatchOrganizationRequest) SetVault(v PatchOrganizationRequestVault) {
	o.Vault = &v
}

func (o PatchOrganizationRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Vault != nil {
		toSerialize["vault"] = o.Vault
	}
	return json.Marshal(toSerialize)
}

type NullablePatchOrganizationRequest struct {
	value *PatchOrganizationRequest
	isSet bool
}

func (v NullablePatchOrganizationRequest) Get() *PatchOrganizationRequest {
	return v.value
}

func (v *NullablePatchOrganizationRequest) Set(val *PatchOrganizationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchOrganizationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchOrganizationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchOrganizationRequest(val *PatchOrganizationRequest) *NullablePatchOrganizationRequest {
	return &NullablePatchOrganizationRequest{value: val, isSet: true}
}

func (v NullablePatchOrganizationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchOrganizationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


