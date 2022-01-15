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

// Selector struct for Selector
type Selector struct {
	// one of \"machines\" or \"tags\"
	Type string `json:"type"`
	// a list of UUIDs in case type is \"machines\"
	Ids *[]string `json:"ids,omitempty"`
	// a list of tags in case type is \"tags\"
	Include *[]string `json:"include,omitempty"`
}

// NewSelector instantiates a new Selector object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSelector(type_ string, ) *Selector {
	this := Selector{}
	this.Type = type_
	return &this
}

// NewSelectorWithDefaults instantiates a new Selector object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSelectorWithDefaults() *Selector {
	this := Selector{}
	return &this
}

// GetType returns the Type field value
func (o *Selector) GetType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Selector) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Selector) SetType(v string) {
	o.Type = v
}

// GetIds returns the Ids field value if set, zero value otherwise.
func (o *Selector) GetIds() []string {
	if o == nil || o.Ids == nil {
		var ret []string
		return ret
	}
	return *o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Selector) GetIdsOk() (*[]string, bool) {
	if o == nil || o.Ids == nil {
		return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *Selector) HasIds() bool {
	if o != nil && o.Ids != nil {
		return true
	}

	return false
}

// SetIds gets a reference to the given []string and assigns it to the Ids field.
func (o *Selector) SetIds(v []string) {
	o.Ids = &v
}

// GetInclude returns the Include field value if set, zero value otherwise.
func (o *Selector) GetInclude() []string {
	if o == nil || o.Include == nil {
		var ret []string
		return ret
	}
	return *o.Include
}

// GetIncludeOk returns a tuple with the Include field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Selector) GetIncludeOk() (*[]string, bool) {
	if o == nil || o.Include == nil {
		return nil, false
	}
	return o.Include, true
}

// HasInclude returns a boolean if a field has been set.
func (o *Selector) HasInclude() bool {
	if o != nil && o.Include != nil {
		return true
	}

	return false
}

// SetInclude gets a reference to the given []string and assigns it to the Include field.
func (o *Selector) SetInclude(v []string) {
	o.Include = &v
}

func (o Selector) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.Ids != nil {
		toSerialize["ids"] = o.Ids
	}
	if o.Include != nil {
		toSerialize["include"] = o.Include
	}
	return json.Marshal(toSerialize)
}

type NullableSelector struct {
	value *Selector
	isSet bool
}

func (v NullableSelector) Get() *Selector {
	return v.value
}

func (v *NullableSelector) Set(val *Selector) {
	v.value = val
	v.isSet = true
}

func (v NullableSelector) IsSet() bool {
	return v.isSet
}

func (v *NullableSelector) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSelector(val *Selector) *NullableSelector {
	return &NullableSelector{value: val, isSet: true}
}

func (v NullableSelector) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSelector) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


