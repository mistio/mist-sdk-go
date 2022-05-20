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

// Command struct for Command
type Command struct {
	// the command to be executed 
	Command string `json:"command"`
}

// NewCommand instantiates a new Command object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommand(command string) *Command {
	this := Command{}
	this.Command = command
	return &this
}

// NewCommandWithDefaults instantiates a new Command object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommandWithDefaults() *Command {
	this := Command{}
	return &this
}

// GetCommand returns the Command field value
func (o *Command) GetCommand() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Command
}

// GetCommandOk returns a tuple with the Command field value
// and a boolean to check if the value has been set.
func (o *Command) GetCommandOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Command, true
}

// SetCommand sets field value
func (o *Command) SetCommand(v string) {
	o.Command = v
}

func (o Command) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["command"] = o.Command
	}
	return json.Marshal(toSerialize)
}

type NullableCommand struct {
	value *Command
	isSet bool
}

func (v NullableCommand) Get() *Command {
	return v.value
}

func (v *NullableCommand) Set(val *Command) {
	v.value = val
	v.isSet = true
}

func (v NullableCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommand(val *Command) *NullableCommand {
	return &NullableCommand{value: val, isSet: true}
}

func (v NullableCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


