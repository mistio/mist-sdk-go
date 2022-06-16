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

// ScriptToRun struct for ScriptToRun
type ScriptToRun struct {
	ScriptType *string `json:"script_type,omitempty"`
	// Command that is about to run
	Command string `json:"command"`
	// Name or ID of the script to run
	Script string `json:"script"`
	Params *string `json:"params,omitempty"`
}

// NewScriptToRun instantiates a new ScriptToRun object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScriptToRun(command string, script string) *ScriptToRun {
	this := ScriptToRun{}
	this.Command = command
	this.Script = script
	return &this
}

// NewScriptToRunWithDefaults instantiates a new ScriptToRun object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScriptToRunWithDefaults() *ScriptToRun {
	this := ScriptToRun{}
	return &this
}

// GetScriptType returns the ScriptType field value if set, zero value otherwise.
func (o *ScriptToRun) GetScriptType() string {
	if o == nil || o.ScriptType == nil {
		var ret string
		return ret
	}
	return *o.ScriptType
}

// GetScriptTypeOk returns a tuple with the ScriptType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScriptToRun) GetScriptTypeOk() (*string, bool) {
	if o == nil || o.ScriptType == nil {
		return nil, false
	}
	return o.ScriptType, true
}

// HasScriptType returns a boolean if a field has been set.
func (o *ScriptToRun) HasScriptType() bool {
	if o != nil && o.ScriptType != nil {
		return true
	}

	return false
}

// SetScriptType gets a reference to the given string and assigns it to the ScriptType field.
func (o *ScriptToRun) SetScriptType(v string) {
	o.ScriptType = &v
}

// GetCommand returns the Command field value
func (o *ScriptToRun) GetCommand() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Command
}

// GetCommandOk returns a tuple with the Command field value
// and a boolean to check if the value has been set.
func (o *ScriptToRun) GetCommandOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Command, true
}

// SetCommand sets field value
func (o *ScriptToRun) SetCommand(v string) {
	o.Command = v
}

// GetScript returns the Script field value
func (o *ScriptToRun) GetScript() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Script
}

// GetScriptOk returns a tuple with the Script field value
// and a boolean to check if the value has been set.
func (o *ScriptToRun) GetScriptOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Script, true
}

// SetScript sets field value
func (o *ScriptToRun) SetScript(v string) {
	o.Script = v
}

// GetParams returns the Params field value if set, zero value otherwise.
func (o *ScriptToRun) GetParams() string {
	if o == nil || o.Params == nil {
		var ret string
		return ret
	}
	return *o.Params
}

// GetParamsOk returns a tuple with the Params field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScriptToRun) GetParamsOk() (*string, bool) {
	if o == nil || o.Params == nil {
		return nil, false
	}
	return o.Params, true
}

// HasParams returns a boolean if a field has been set.
func (o *ScriptToRun) HasParams() bool {
	if o != nil && o.Params != nil {
		return true
	}

	return false
}

// SetParams gets a reference to the given string and assigns it to the Params field.
func (o *ScriptToRun) SetParams(v string) {
	o.Params = &v
}

func (o ScriptToRun) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ScriptType != nil {
		toSerialize["script_type"] = o.ScriptType
	}
	if true {
		toSerialize["command"] = o.Command
	}
	if true {
		toSerialize["script"] = o.Script
	}
	if o.Params != nil {
		toSerialize["params"] = o.Params
	}
	return json.Marshal(toSerialize)
}

type NullableScriptToRun struct {
	value *ScriptToRun
	isSet bool
}

func (v NullableScriptToRun) Get() *ScriptToRun {
	return v.value
}

func (v *NullableScriptToRun) Set(val *ScriptToRun) {
	v.value = val
	v.isSet = true
}

func (v NullableScriptToRun) IsSet() bool {
	return v.isSet
}

func (v *NullableScriptToRun) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScriptToRun(val *ScriptToRun) *NullableScriptToRun {
	return &NullableScriptToRun{value: val, isSet: true}
}

func (v NullableScriptToRun) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScriptToRun) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


