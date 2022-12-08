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

// StartAction struct for StartAction
type StartAction struct {
	ActionType *string `json:"action_type,omitempty"`
}

// NewStartAction instantiates a new StartAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStartAction() *StartAction {
	this := StartAction{}
	return &this
}

// NewStartActionWithDefaults instantiates a new StartAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStartActionWithDefaults() *StartAction {
	this := StartAction{}
	return &this
}

// GetActionType returns the ActionType field value if set, zero value otherwise.
func (o *StartAction) GetActionType() string {
	if o == nil || o.ActionType == nil {
		var ret string
		return ret
	}
	return *o.ActionType
}

// GetActionTypeOk returns a tuple with the ActionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartAction) GetActionTypeOk() (*string, bool) {
	if o == nil || o.ActionType == nil {
		return nil, false
	}
	return o.ActionType, true
}

// HasActionType returns a boolean if a field has been set.
func (o *StartAction) HasActionType() bool {
	if o != nil && o.ActionType != nil {
		return true
	}

	return false
}

// SetActionType gets a reference to the given string and assigns it to the ActionType field.
func (o *StartAction) SetActionType(v string) {
	o.ActionType = &v
}

func (o StartAction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ActionType != nil {
		toSerialize["action_type"] = o.ActionType
	}
	return json.Marshal(toSerialize)
}

type NullableStartAction struct {
	value *StartAction
	isSet bool
}

func (v NullableStartAction) Get() *StartAction {
	return v.value
}

func (v *NullableStartAction) Set(val *StartAction) {
	v.value = val
	v.isSet = true
}

func (v NullableStartAction) IsSet() bool {
	return v.isSet
}

func (v *NullableStartAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStartAction(val *StartAction) *NullableStartAction {
	return &NullableStartAction{value: val, isSet: true}
}

func (v NullableStartAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStartAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

