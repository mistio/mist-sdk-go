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

// ResizeAction struct for ResizeAction
type ResizeAction struct {
	ActionType *string `json:"action_type,omitempty"`
	// the params of the resize action to be executed 
	Params string `json:"params"`
}

// NewResizeAction instantiates a new ResizeAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResizeAction(params string) *ResizeAction {
	this := ResizeAction{}
	this.Params = params
	return &this
}

// NewResizeActionWithDefaults instantiates a new ResizeAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResizeActionWithDefaults() *ResizeAction {
	this := ResizeAction{}
	return &this
}

// GetActionType returns the ActionType field value if set, zero value otherwise.
func (o *ResizeAction) GetActionType() string {
	if o == nil || o.ActionType == nil {
		var ret string
		return ret
	}
	return *o.ActionType
}

// GetActionTypeOk returns a tuple with the ActionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResizeAction) GetActionTypeOk() (*string, bool) {
	if o == nil || o.ActionType == nil {
		return nil, false
	}
	return o.ActionType, true
}

// HasActionType returns a boolean if a field has been set.
func (o *ResizeAction) HasActionType() bool {
	if o != nil && o.ActionType != nil {
		return true
	}

	return false
}

// SetActionType gets a reference to the given string and assigns it to the ActionType field.
func (o *ResizeAction) SetActionType(v string) {
	o.ActionType = &v
}

// GetParams returns the Params field value
func (o *ResizeAction) GetParams() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Params
}

// GetParamsOk returns a tuple with the Params field value
// and a boolean to check if the value has been set.
func (o *ResizeAction) GetParamsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Params, true
}

// SetParams sets field value
func (o *ResizeAction) SetParams(v string) {
	o.Params = v
}

func (o ResizeAction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ActionType != nil {
		toSerialize["action_type"] = o.ActionType
	}
	if true {
		toSerialize["params"] = o.Params
	}
	return json.Marshal(toSerialize)
}

type NullableResizeAction struct {
	value *ResizeAction
	isSet bool
}

func (v NullableResizeAction) Get() *ResizeAction {
	return v.value
}

func (v *NullableResizeAction) Set(val *ResizeAction) {
	v.value = val
	v.isSet = true
}

func (v NullableResizeAction) IsSet() bool {
	return v.isSet
}

func (v *NullableResizeAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResizeAction(val *ResizeAction) *NullableResizeAction {
	return &NullableResizeAction{value: val, isSet: true}
}

func (v NullableResizeAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResizeAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


