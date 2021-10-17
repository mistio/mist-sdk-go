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

// RunScriptRequest struct for RunScriptRequest
type RunScriptRequest struct {
	Machine string `json:"machine"`
	Params *string `json:"params,omitempty"`
	Su *string `json:"su,omitempty"`
	Env *string `json:"env,omitempty"`
	JobId *string `json:"job_id,omitempty"`
}

// NewRunScriptRequest instantiates a new RunScriptRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRunScriptRequest(machine string, ) *RunScriptRequest {
	this := RunScriptRequest{}
	this.Machine = machine
	return &this
}

// NewRunScriptRequestWithDefaults instantiates a new RunScriptRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRunScriptRequestWithDefaults() *RunScriptRequest {
	this := RunScriptRequest{}
	return &this
}

// GetMachine returns the Machine field value
func (o *RunScriptRequest) GetMachine() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Machine
}

// GetMachineOk returns a tuple with the Machine field value
// and a boolean to check if the value has been set.
func (o *RunScriptRequest) GetMachineOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Machine, true
}

// SetMachine sets field value
func (o *RunScriptRequest) SetMachine(v string) {
	o.Machine = v
}

// GetParams returns the Params field value if set, zero value otherwise.
func (o *RunScriptRequest) GetParams() string {
	if o == nil || o.Params == nil {
		var ret string
		return ret
	}
	return *o.Params
}

// GetParamsOk returns a tuple with the Params field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunScriptRequest) GetParamsOk() (*string, bool) {
	if o == nil || o.Params == nil {
		return nil, false
	}
	return o.Params, true
}

// HasParams returns a boolean if a field has been set.
func (o *RunScriptRequest) HasParams() bool {
	if o != nil && o.Params != nil {
		return true
	}

	return false
}

// SetParams gets a reference to the given string and assigns it to the Params field.
func (o *RunScriptRequest) SetParams(v string) {
	o.Params = &v
}

// GetSu returns the Su field value if set, zero value otherwise.
func (o *RunScriptRequest) GetSu() string {
	if o == nil || o.Su == nil {
		var ret string
		return ret
	}
	return *o.Su
}

// GetSuOk returns a tuple with the Su field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunScriptRequest) GetSuOk() (*string, bool) {
	if o == nil || o.Su == nil {
		return nil, false
	}
	return o.Su, true
}

// HasSu returns a boolean if a field has been set.
func (o *RunScriptRequest) HasSu() bool {
	if o != nil && o.Su != nil {
		return true
	}

	return false
}

// SetSu gets a reference to the given string and assigns it to the Su field.
func (o *RunScriptRequest) SetSu(v string) {
	o.Su = &v
}

// GetEnv returns the Env field value if set, zero value otherwise.
func (o *RunScriptRequest) GetEnv() string {
	if o == nil || o.Env == nil {
		var ret string
		return ret
	}
	return *o.Env
}

// GetEnvOk returns a tuple with the Env field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunScriptRequest) GetEnvOk() (*string, bool) {
	if o == nil || o.Env == nil {
		return nil, false
	}
	return o.Env, true
}

// HasEnv returns a boolean if a field has been set.
func (o *RunScriptRequest) HasEnv() bool {
	if o != nil && o.Env != nil {
		return true
	}

	return false
}

// SetEnv gets a reference to the given string and assigns it to the Env field.
func (o *RunScriptRequest) SetEnv(v string) {
	o.Env = &v
}

// GetJobId returns the JobId field value if set, zero value otherwise.
func (o *RunScriptRequest) GetJobId() string {
	if o == nil || o.JobId == nil {
		var ret string
		return ret
	}
	return *o.JobId
}

// GetJobIdOk returns a tuple with the JobId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunScriptRequest) GetJobIdOk() (*string, bool) {
	if o == nil || o.JobId == nil {
		return nil, false
	}
	return o.JobId, true
}

// HasJobId returns a boolean if a field has been set.
func (o *RunScriptRequest) HasJobId() bool {
	if o != nil && o.JobId != nil {
		return true
	}

	return false
}

// SetJobId gets a reference to the given string and assigns it to the JobId field.
func (o *RunScriptRequest) SetJobId(v string) {
	o.JobId = &v
}

func (o RunScriptRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["machine"] = o.Machine
	}
	if o.Params != nil {
		toSerialize["params"] = o.Params
	}
	if o.Su != nil {
		toSerialize["su"] = o.Su
	}
	if o.Env != nil {
		toSerialize["env"] = o.Env
	}
	if o.JobId != nil {
		toSerialize["job_id"] = o.JobId
	}
	return json.Marshal(toSerialize)
}

type NullableRunScriptRequest struct {
	value *RunScriptRequest
	isSet bool
}

func (v NullableRunScriptRequest) Get() *RunScriptRequest {
	return v.value
}

func (v *NullableRunScriptRequest) Set(val *RunScriptRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRunScriptRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRunScriptRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRunScriptRequest(val *RunScriptRequest) *NullableRunScriptRequest {
	return &NullableRunScriptRequest{value: val, isSet: true}
}

func (v NullableRunScriptRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRunScriptRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


