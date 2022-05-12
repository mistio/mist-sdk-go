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
	"fmt"
)

// MachineState the model 'MachineState'
type MachineState string

// List of MachineState
const (
	RUNNING MachineState = "running"
	STARTING MachineState = "starting"
	STOPPING MachineState = "stopping"
	STOPPED MachineState = "stopped"
	PENDING MachineState = "pending"
	SUSPENDED MachineState = "suspended"
	TERMINATED MachineState = "terminated"
	ERROR MachineState = "error"
	REBOOTING MachineState = "rebooting"
	PAUSED MachineState = "paused"
	RECONFIGURING MachineState = "reconfiguring"
	UNKNOWN MachineState = "unknown"
	UPDATING MachineState = "updating"
	MIGRATING MachineState = "migrating"
	NORMAL MachineState = "normal"
	SUCCEEDED MachineState = "succeeded"
	FAILED MachineState = "failed"
)

// All allowed values of MachineState enum
var AllowedMachineStateEnumValues = []MachineState{
	"running",
	"starting",
	"stopping",
	"stopped",
	"pending",
	"suspended",
	"terminated",
	"error",
	"rebooting",
	"paused",
	"reconfiguring",
	"unknown",
	"updating",
	"migrating",
	"normal",
	"succeeded",
	"failed",
}

func (v *MachineState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MachineState(value)
	for _, existing := range AllowedMachineStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MachineState", value)
}

// NewMachineStateFromValue returns a pointer to a valid MachineState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMachineStateFromValue(v string) (*MachineState, error) {
	ev := MachineState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MachineState: valid values are %v", v, AllowedMachineStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MachineState) IsValid() bool {
	for _, existing := range AllowedMachineStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MachineState value
func (v MachineState) Ptr() *MachineState {
	return &v
}

type NullableMachineState struct {
	value *MachineState
	isSet bool
}

func (v NullableMachineState) Get() *MachineState {
	return v.value
}

func (v *NullableMachineState) Set(val *MachineState) {
	v.value = val
	v.isSet = true
}

func (v NullableMachineState) IsSet() bool {
	return v.isSet
}

func (v *NullableMachineState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMachineState(val *MachineState) *NullableMachineState {
	return &NullableMachineState{value: val, isSet: true}
}

func (v NullableMachineState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMachineState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

