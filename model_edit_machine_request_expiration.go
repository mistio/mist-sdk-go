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

// EditMachineRequestExpiration struct for EditMachineRequestExpiration
type EditMachineRequestExpiration struct {
	// format should be ΥΥΥΥ-ΜΜ-DD HH:MM:SS
	Date string `json:"date"`
	Action string `json:"action"`
	// seconds before the expiration date to be notified
	Notify *int32 `json:"notify,omitempty"`
}

// NewEditMachineRequestExpiration instantiates a new EditMachineRequestExpiration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEditMachineRequestExpiration(date string, action string) *EditMachineRequestExpiration {
	this := EditMachineRequestExpiration{}
	this.Date = date
	this.Action = action
	return &this
}

// NewEditMachineRequestExpirationWithDefaults instantiates a new EditMachineRequestExpiration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEditMachineRequestExpirationWithDefaults() *EditMachineRequestExpiration {
	this := EditMachineRequestExpiration{}
	return &this
}

// GetDate returns the Date field value
func (o *EditMachineRequestExpiration) GetDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Date
}

// GetDateOk returns a tuple with the Date field value
// and a boolean to check if the value has been set.
func (o *EditMachineRequestExpiration) GetDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Date, true
}

// SetDate sets field value
func (o *EditMachineRequestExpiration) SetDate(v string) {
	o.Date = v
}

// GetAction returns the Action field value
func (o *EditMachineRequestExpiration) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *EditMachineRequestExpiration) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *EditMachineRequestExpiration) SetAction(v string) {
	o.Action = v
}

// GetNotify returns the Notify field value if set, zero value otherwise.
func (o *EditMachineRequestExpiration) GetNotify() int32 {
	if o == nil || o.Notify == nil {
		var ret int32
		return ret
	}
	return *o.Notify
}

// GetNotifyOk returns a tuple with the Notify field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditMachineRequestExpiration) GetNotifyOk() (*int32, bool) {
	if o == nil || o.Notify == nil {
		return nil, false
	}
	return o.Notify, true
}

// HasNotify returns a boolean if a field has been set.
func (o *EditMachineRequestExpiration) HasNotify() bool {
	if o != nil && o.Notify != nil {
		return true
	}

	return false
}

// SetNotify gets a reference to the given int32 and assigns it to the Notify field.
func (o *EditMachineRequestExpiration) SetNotify(v int32) {
	o.Notify = &v
}

func (o EditMachineRequestExpiration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["date"] = o.Date
	}
	if true {
		toSerialize["action"] = o.Action
	}
	if o.Notify != nil {
		toSerialize["notify"] = o.Notify
	}
	return json.Marshal(toSerialize)
}

type NullableEditMachineRequestExpiration struct {
	value *EditMachineRequestExpiration
	isSet bool
}

func (v NullableEditMachineRequestExpiration) Get() *EditMachineRequestExpiration {
	return v.value
}

func (v *NullableEditMachineRequestExpiration) Set(val *EditMachineRequestExpiration) {
	v.value = val
	v.isSet = true
}

func (v NullableEditMachineRequestExpiration) IsSet() bool {
	return v.isSet
}

func (v *NullableEditMachineRequestExpiration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEditMachineRequestExpiration(val *EditMachineRequestExpiration) *NullableEditMachineRequestExpiration {
	return &NullableEditMachineRequestExpiration{value: val, isSet: true}
}

func (v NullableEditMachineRequestExpiration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEditMachineRequestExpiration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


