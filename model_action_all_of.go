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

// ActionAllOf struct for ActionAllOf
type ActionAllOf struct {
	// the action's type: notification, resource_action, run_script 
	ActionType string `json:"action_type"`
}

// NewActionAllOf instantiates a new ActionAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActionAllOf(actionType string) *ActionAllOf {
	this := ActionAllOf{}
	this.ActionType = actionType
	return &this
}

// NewActionAllOfWithDefaults instantiates a new ActionAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActionAllOfWithDefaults() *ActionAllOf {
	this := ActionAllOf{}
	return &this
}

// GetActionType returns the ActionType field value
func (o *ActionAllOf) GetActionType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ActionType
}

// GetActionTypeOk returns a tuple with the ActionType field value
// and a boolean to check if the value has been set.
func (o *ActionAllOf) GetActionTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActionType, true
}

// SetActionType sets field value
func (o *ActionAllOf) SetActionType(v string) {
	o.ActionType = v
}

func (o ActionAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["action_type"] = o.ActionType
	}
	return json.Marshal(toSerialize)
}

type NullableActionAllOf struct {
	value *ActionAllOf
	isSet bool
}

func (v NullableActionAllOf) Get() *ActionAllOf {
	return v.value
}

func (v *NullableActionAllOf) Set(val *ActionAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableActionAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableActionAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActionAllOf(val *ActionAllOf) *NullableActionAllOf {
	return &NullableActionAllOf{value: val, isSet: true}
}

func (v NullableActionAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActionAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


