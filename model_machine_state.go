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

func (v *MachineState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MachineState(value)
	for _, existing := range []MachineState{ "running", "starting", "stopping", "stopped", "pending", "suspended", "terminated", "error", "rebooting", "paused", "reconfiguring", "unknown", "updating", "migrating", "normal", "succeeded", "failed",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MachineState", value)
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

