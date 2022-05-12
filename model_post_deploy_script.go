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

// PostDeployScript struct for PostDeployScript
type PostDeployScript struct {
	// Name or ID of the script to run
	Script string `json:"script"`
	Params *string `json:"params,omitempty"`
}

// NewPostDeployScript instantiates a new PostDeployScript object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostDeployScript(script string) *PostDeployScript {
	this := PostDeployScript{}
	this.Script = script
	return &this
}

// NewPostDeployScriptWithDefaults instantiates a new PostDeployScript object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostDeployScriptWithDefaults() *PostDeployScript {
	this := PostDeployScript{}
	return &this
}

// GetScript returns the Script field value
func (o *PostDeployScript) GetScript() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Script
}

// GetScriptOk returns a tuple with the Script field value
// and a boolean to check if the value has been set.
func (o *PostDeployScript) GetScriptOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Script, true
}

// SetScript sets field value
func (o *PostDeployScript) SetScript(v string) {
	o.Script = v
}

// GetParams returns the Params field value if set, zero value otherwise.
func (o *PostDeployScript) GetParams() string {
	if o == nil || o.Params == nil {
		var ret string
		return ret
	}
	return *o.Params
}

// GetParamsOk returns a tuple with the Params field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostDeployScript) GetParamsOk() (*string, bool) {
	if o == nil || o.Params == nil {
		return nil, false
	}
	return o.Params, true
}

// HasParams returns a boolean if a field has been set.
func (o *PostDeployScript) HasParams() bool {
	if o != nil && o.Params != nil {
		return true
	}

	return false
}

// SetParams gets a reference to the given string and assigns it to the Params field.
func (o *PostDeployScript) SetParams(v string) {
	o.Params = &v
}

func (o PostDeployScript) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["script"] = o.Script
	}
	if o.Params != nil {
		toSerialize["params"] = o.Params
	}
	return json.Marshal(toSerialize)
}

type NullablePostDeployScript struct {
	value *PostDeployScript
	isSet bool
}

func (v NullablePostDeployScript) Get() *PostDeployScript {
	return v.value
}

func (v *NullablePostDeployScript) Set(val *PostDeployScript) {
	v.value = val
	v.isSet = true
}

func (v NullablePostDeployScript) IsSet() bool {
	return v.isSet
}

func (v *NullablePostDeployScript) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostDeployScript(val *PostDeployScript) *NullablePostDeployScript {
	return &NullablePostDeployScript{value: val, isSet: true}
}

func (v NullablePostDeployScript) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostDeployScript) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


