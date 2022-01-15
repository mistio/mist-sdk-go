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

// KeyMachineAssociation struct for KeyMachineAssociation
type KeyMachineAssociation struct {
	// Name or ID of the SSH key
	Key *string `json:"key,omitempty"`
	// Name or ID of the machine
	Machine *string `json:"machine,omitempty"`
	// Last used time
	LastUsed *string `json:"last_used,omitempty"`
	// SSH port
	Port *int32 `json:"port,omitempty"`
	// SSH user
	User *string `json:"user,omitempty"`
}

// NewKeyMachineAssociation instantiates a new KeyMachineAssociation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeyMachineAssociation() *KeyMachineAssociation {
	this := KeyMachineAssociation{}
	return &this
}

// NewKeyMachineAssociationWithDefaults instantiates a new KeyMachineAssociation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeyMachineAssociationWithDefaults() *KeyMachineAssociation {
	this := KeyMachineAssociation{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *KeyMachineAssociation) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyMachineAssociation) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *KeyMachineAssociation) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *KeyMachineAssociation) SetKey(v string) {
	o.Key = &v
}

// GetMachine returns the Machine field value if set, zero value otherwise.
func (o *KeyMachineAssociation) GetMachine() string {
	if o == nil || o.Machine == nil {
		var ret string
		return ret
	}
	return *o.Machine
}

// GetMachineOk returns a tuple with the Machine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyMachineAssociation) GetMachineOk() (*string, bool) {
	if o == nil || o.Machine == nil {
		return nil, false
	}
	return o.Machine, true
}

// HasMachine returns a boolean if a field has been set.
func (o *KeyMachineAssociation) HasMachine() bool {
	if o != nil && o.Machine != nil {
		return true
	}

	return false
}

// SetMachine gets a reference to the given string and assigns it to the Machine field.
func (o *KeyMachineAssociation) SetMachine(v string) {
	o.Machine = &v
}

// GetLastUsed returns the LastUsed field value if set, zero value otherwise.
func (o *KeyMachineAssociation) GetLastUsed() string {
	if o == nil || o.LastUsed == nil {
		var ret string
		return ret
	}
	return *o.LastUsed
}

// GetLastUsedOk returns a tuple with the LastUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyMachineAssociation) GetLastUsedOk() (*string, bool) {
	if o == nil || o.LastUsed == nil {
		return nil, false
	}
	return o.LastUsed, true
}

// HasLastUsed returns a boolean if a field has been set.
func (o *KeyMachineAssociation) HasLastUsed() bool {
	if o != nil && o.LastUsed != nil {
		return true
	}

	return false
}

// SetLastUsed gets a reference to the given string and assigns it to the LastUsed field.
func (o *KeyMachineAssociation) SetLastUsed(v string) {
	o.LastUsed = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *KeyMachineAssociation) GetPort() int32 {
	if o == nil || o.Port == nil {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyMachineAssociation) GetPortOk() (*int32, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *KeyMachineAssociation) HasPort() bool {
	if o != nil && o.Port != nil {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *KeyMachineAssociation) SetPort(v int32) {
	o.Port = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *KeyMachineAssociation) GetUser() string {
	if o == nil || o.User == nil {
		var ret string
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyMachineAssociation) GetUserOk() (*string, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *KeyMachineAssociation) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given string and assigns it to the User field.
func (o *KeyMachineAssociation) SetUser(v string) {
	o.User = &v
}

func (o KeyMachineAssociation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if o.Machine != nil {
		toSerialize["machine"] = o.Machine
	}
	if o.LastUsed != nil {
		toSerialize["last_used"] = o.LastUsed
	}
	if o.Port != nil {
		toSerialize["port"] = o.Port
	}
	if o.User != nil {
		toSerialize["user"] = o.User
	}
	return json.Marshal(toSerialize)
}

type NullableKeyMachineAssociation struct {
	value *KeyMachineAssociation
	isSet bool
}

func (v NullableKeyMachineAssociation) Get() *KeyMachineAssociation {
	return v.value
}

func (v *NullableKeyMachineAssociation) Set(val *KeyMachineAssociation) {
	v.value = val
	v.isSet = true
}

func (v NullableKeyMachineAssociation) IsSet() bool {
	return v.isSet
}

func (v *NullableKeyMachineAssociation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeyMachineAssociation(val *KeyMachineAssociation) *NullableKeyMachineAssociation {
	return &NullableKeyMachineAssociation{value: val, isSet: true}
}

func (v NullableKeyMachineAssociation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeyMachineAssociation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


