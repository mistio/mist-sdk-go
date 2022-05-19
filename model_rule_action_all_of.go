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

// RuleActionAllOf struct for RuleActionAllOf
type RuleActionAllOf struct {
	// the action's type: notification, machine_action, command 
	Type string `json:"type"`
}

// NewRuleActionAllOf instantiates a new RuleActionAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRuleActionAllOf(type_ string, ) *RuleActionAllOf {
	this := RuleActionAllOf{}
	this.Type = type_
	return &this
}

// NewRuleActionAllOfWithDefaults instantiates a new RuleActionAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRuleActionAllOfWithDefaults() *RuleActionAllOf {
	this := RuleActionAllOf{}
	return &this
}

// GetType returns the Type field value
func (o *RuleActionAllOf) GetType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RuleActionAllOf) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RuleActionAllOf) SetType(v string) {
	o.Type = v
}

func (o RuleActionAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableRuleActionAllOf struct {
	value *RuleActionAllOf
	isSet bool
}

func (v NullableRuleActionAllOf) Get() *RuleActionAllOf {
	return v.value
}

func (v *NullableRuleActionAllOf) Set(val *RuleActionAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableRuleActionAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableRuleActionAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRuleActionAllOf(val *RuleActionAllOf) *NullableRuleActionAllOf {
	return &NullableRuleActionAllOf{value: val, isSet: true}
}

func (v NullableRuleActionAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRuleActionAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


