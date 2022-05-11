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

// CreateNetworkResponse - struct for CreateNetworkResponse
type CreateNetworkResponse struct {
	CreateMachineResponseOneOf *CreateMachineResponseOneOf
	CreateMachineResponseOneOf1 *CreateMachineResponseOneOf1
}

// CreateMachineResponseOneOfAsCreateNetworkResponse is a convenience function that returns CreateMachineResponseOneOf wrapped in CreateNetworkResponse
func CreateMachineResponseOneOfAsCreateNetworkResponse(v *CreateMachineResponseOneOf) CreateNetworkResponse {
	return CreateNetworkResponse{
		CreateMachineResponseOneOf: v,
	}
}

// CreateMachineResponseOneOf1AsCreateNetworkResponse is a convenience function that returns CreateMachineResponseOneOf1 wrapped in CreateNetworkResponse
func CreateMachineResponseOneOf1AsCreateNetworkResponse(v *CreateMachineResponseOneOf1) CreateNetworkResponse {
	return CreateNetworkResponse{
		CreateMachineResponseOneOf1: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *CreateNetworkResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into CreateMachineResponseOneOf
	err = newStrictDecoder(data).Decode(&dst.CreateMachineResponseOneOf)
	if err == nil {
		jsonCreateMachineResponseOneOf, _ := json.Marshal(dst.CreateMachineResponseOneOf)
		if string(jsonCreateMachineResponseOneOf) == "{}" { // empty struct
			dst.CreateMachineResponseOneOf = nil
		} else {
			match++
		}
	} else {
		dst.CreateMachineResponseOneOf = nil
	}

	// try to unmarshal data into CreateMachineResponseOneOf1
	err = newStrictDecoder(data).Decode(&dst.CreateMachineResponseOneOf1)
	if err == nil {
		jsonCreateMachineResponseOneOf1, _ := json.Marshal(dst.CreateMachineResponseOneOf1)
		if string(jsonCreateMachineResponseOneOf1) == "{}" { // empty struct
			dst.CreateMachineResponseOneOf1 = nil
		} else {
			match++
		}
	} else {
		dst.CreateMachineResponseOneOf1 = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.CreateMachineResponseOneOf = nil
		dst.CreateMachineResponseOneOf1 = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(CreateNetworkResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(CreateNetworkResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CreateNetworkResponse) MarshalJSON() ([]byte, error) {
	if src.CreateMachineResponseOneOf != nil {
		return json.Marshal(&src.CreateMachineResponseOneOf)
	}

	if src.CreateMachineResponseOneOf1 != nil {
		return json.Marshal(&src.CreateMachineResponseOneOf1)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CreateNetworkResponse) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.CreateMachineResponseOneOf != nil {
		return obj.CreateMachineResponseOneOf
	}

	if obj.CreateMachineResponseOneOf1 != nil {
		return obj.CreateMachineResponseOneOf1
	}

	// all schemas are nil
	return nil
}

type NullableCreateNetworkResponse struct {
	value *CreateNetworkResponse
	isSet bool
}

func (v NullableCreateNetworkResponse) Get() *CreateNetworkResponse {
	return v.value
}

func (v *NullableCreateNetworkResponse) Set(val *CreateNetworkResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetworkResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetworkResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetworkResponse(val *CreateNetworkResponse) *NullableCreateNetworkResponse {
	return &NullableCreateNetworkResponse{value: val, isSet: true}
}

func (v NullableCreateNetworkResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetworkResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


