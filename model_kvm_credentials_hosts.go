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

// KvmCredentialsHosts struct for KvmCredentialsHosts
type KvmCredentialsHosts struct {
	Host string `json:"host"`
	Alias *string `json:"alias,omitempty"`
	Key string `json:"key"`
	User *string `json:"user,omitempty"`
	Port *string `json:"port,omitempty"`
}

// NewKvmCredentialsHosts instantiates a new KvmCredentialsHosts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKvmCredentialsHosts(host string, key string) *KvmCredentialsHosts {
	this := KvmCredentialsHosts{}
	this.Host = host
	this.Key = key
	var port string = "22"
	this.Port = &port
	return &this
}

// NewKvmCredentialsHostsWithDefaults instantiates a new KvmCredentialsHosts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKvmCredentialsHostsWithDefaults() *KvmCredentialsHosts {
	this := KvmCredentialsHosts{}
	var port string = "22"
	this.Port = &port
	return &this
}

// GetHost returns the Host field value
func (o *KvmCredentialsHosts) GetHost() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Host
}

// GetHostOk returns a tuple with the Host field value
// and a boolean to check if the value has been set.
func (o *KvmCredentialsHosts) GetHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Host, true
}

// SetHost sets field value
func (o *KvmCredentialsHosts) SetHost(v string) {
	o.Host = v
}

// GetAlias returns the Alias field value if set, zero value otherwise.
func (o *KvmCredentialsHosts) GetAlias() string {
	if o == nil || o.Alias == nil {
		var ret string
		return ret
	}
	return *o.Alias
}

// GetAliasOk returns a tuple with the Alias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KvmCredentialsHosts) GetAliasOk() (*string, bool) {
	if o == nil || o.Alias == nil {
		return nil, false
	}
	return o.Alias, true
}

// HasAlias returns a boolean if a field has been set.
func (o *KvmCredentialsHosts) HasAlias() bool {
	if o != nil && o.Alias != nil {
		return true
	}

	return false
}

// SetAlias gets a reference to the given string and assigns it to the Alias field.
func (o *KvmCredentialsHosts) SetAlias(v string) {
	o.Alias = &v
}

// GetKey returns the Key field value
func (o *KvmCredentialsHosts) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *KvmCredentialsHosts) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *KvmCredentialsHosts) SetKey(v string) {
	o.Key = v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *KvmCredentialsHosts) GetUser() string {
	if o == nil || o.User == nil {
		var ret string
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KvmCredentialsHosts) GetUserOk() (*string, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *KvmCredentialsHosts) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given string and assigns it to the User field.
func (o *KvmCredentialsHosts) SetUser(v string) {
	o.User = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *KvmCredentialsHosts) GetPort() string {
	if o == nil || o.Port == nil {
		var ret string
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KvmCredentialsHosts) GetPortOk() (*string, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *KvmCredentialsHosts) HasPort() bool {
	if o != nil && o.Port != nil {
		return true
	}

	return false
}

// SetPort gets a reference to the given string and assigns it to the Port field.
func (o *KvmCredentialsHosts) SetPort(v string) {
	o.Port = &v
}

func (o KvmCredentialsHosts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["host"] = o.Host
	}
	if o.Alias != nil {
		toSerialize["alias"] = o.Alias
	}
	if true {
		toSerialize["key"] = o.Key
	}
	if o.User != nil {
		toSerialize["user"] = o.User
	}
	if o.Port != nil {
		toSerialize["port"] = o.Port
	}
	return json.Marshal(toSerialize)
}

type NullableKvmCredentialsHosts struct {
	value *KvmCredentialsHosts
	isSet bool
}

func (v NullableKvmCredentialsHosts) Get() *KvmCredentialsHosts {
	return v.value
}

func (v *NullableKvmCredentialsHosts) Set(val *KvmCredentialsHosts) {
	v.value = val
	v.isSet = true
}

func (v NullableKvmCredentialsHosts) IsSet() bool {
	return v.isSet
}

func (v *NullableKvmCredentialsHosts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKvmCredentialsHosts(val *KvmCredentialsHosts) *NullableKvmCredentialsHosts {
	return &NullableKvmCredentialsHosts{value: val, isSet: true}
}

func (v NullableKvmCredentialsHosts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKvmCredentialsHosts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


