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

// AddScheduleRequestReminderReminderWhen struct for AddScheduleRequestReminderReminderWhen
type AddScheduleRequestReminderReminderWhen struct {
	Value *int32 `json:"value,omitempty"`
	Unit *string `json:"unit,omitempty"`
}

// NewAddScheduleRequestReminderReminderWhen instantiates a new AddScheduleRequestReminderReminderWhen object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddScheduleRequestReminderReminderWhen() *AddScheduleRequestReminderReminderWhen {
	this := AddScheduleRequestReminderReminderWhen{}
	return &this
}

// NewAddScheduleRequestReminderReminderWhenWithDefaults instantiates a new AddScheduleRequestReminderReminderWhen object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddScheduleRequestReminderReminderWhenWithDefaults() *AddScheduleRequestReminderReminderWhen {
	this := AddScheduleRequestReminderReminderWhen{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *AddScheduleRequestReminderReminderWhen) GetValue() int32 {
	if o == nil || o.Value == nil {
		var ret int32
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddScheduleRequestReminderReminderWhen) GetValueOk() (*int32, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *AddScheduleRequestReminderReminderWhen) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given int32 and assigns it to the Value field.
func (o *AddScheduleRequestReminderReminderWhen) SetValue(v int32) {
	o.Value = &v
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *AddScheduleRequestReminderReminderWhen) GetUnit() string {
	if o == nil || o.Unit == nil {
		var ret string
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddScheduleRequestReminderReminderWhen) GetUnitOk() (*string, bool) {
	if o == nil || o.Unit == nil {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *AddScheduleRequestReminderReminderWhen) HasUnit() bool {
	if o != nil && o.Unit != nil {
		return true
	}

	return false
}

// SetUnit gets a reference to the given string and assigns it to the Unit field.
func (o *AddScheduleRequestReminderReminderWhen) SetUnit(v string) {
	o.Unit = &v
}

func (o AddScheduleRequestReminderReminderWhen) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.Unit != nil {
		toSerialize["unit"] = o.Unit
	}
	return json.Marshal(toSerialize)
}

type NullableAddScheduleRequestReminderReminderWhen struct {
	value *AddScheduleRequestReminderReminderWhen
	isSet bool
}

func (v NullableAddScheduleRequestReminderReminderWhen) Get() *AddScheduleRequestReminderReminderWhen {
	return v.value
}

func (v *NullableAddScheduleRequestReminderReminderWhen) Set(val *AddScheduleRequestReminderReminderWhen) {
	v.value = val
	v.isSet = true
}

func (v NullableAddScheduleRequestReminderReminderWhen) IsSet() bool {
	return v.isSet
}

func (v *NullableAddScheduleRequestReminderReminderWhen) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddScheduleRequestReminderReminderWhen(val *AddScheduleRequestReminderReminderWhen) *NullableAddScheduleRequestReminderReminderWhen {
	return &NullableAddScheduleRequestReminderReminderWhen{value: val, isSet: true}
}

func (v NullableAddScheduleRequestReminderReminderWhen) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddScheduleRequestReminderReminderWhen) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

