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
	"time"
)

// OneOffSchedule one_off schedule
type OneOffSchedule struct {
	ScheduleType string `json:"schedule_type"`
	// When one_off schedule should run, e.g 2021-09-22T18:19:28Z
	Datetime time.Time `json:"datetime"`
}

// NewOneOffSchedule instantiates a new OneOffSchedule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOneOffSchedule(scheduleType string, datetime time.Time, ) *OneOffSchedule {
	this := OneOffSchedule{}
	this.ScheduleType = scheduleType
	this.Datetime = datetime
	return &this
}

// NewOneOffScheduleWithDefaults instantiates a new OneOffSchedule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOneOffScheduleWithDefaults() *OneOffSchedule {
	this := OneOffSchedule{}
	return &this
}

// GetScheduleType returns the ScheduleType field value
func (o *OneOffSchedule) GetScheduleType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.ScheduleType
}

// GetScheduleTypeOk returns a tuple with the ScheduleType field value
// and a boolean to check if the value has been set.
func (o *OneOffSchedule) GetScheduleTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ScheduleType, true
}

// SetScheduleType sets field value
func (o *OneOffSchedule) SetScheduleType(v string) {
	o.ScheduleType = v
}

// GetDatetime returns the Datetime field value
func (o *OneOffSchedule) GetDatetime() time.Time {
	if o == nil  {
		var ret time.Time
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *OneOffSchedule) GetDatetimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *OneOffSchedule) SetDatetime(v time.Time) {
	o.Datetime = v
}

func (o OneOffSchedule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["schedule_type"] = o.ScheduleType
	}
	if true {
		toSerialize["datetime"] = o.Datetime
	}
	return json.Marshal(toSerialize)
}

type NullableOneOffSchedule struct {
	value *OneOffSchedule
	isSet bool
}

func (v NullableOneOffSchedule) Get() *OneOffSchedule {
	return v.value
}

func (v *NullableOneOffSchedule) Set(val *OneOffSchedule) {
	v.value = val
	v.isSet = true
}

func (v NullableOneOffSchedule) IsSet() bool {
	return v.isSet
}

func (v *NullableOneOffSchedule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOneOffSchedule(val *OneOffSchedule) *NullableOneOffSchedule {
	return &NullableOneOffSchedule{value: val, isSet: true}
}

func (v NullableOneOffSchedule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOneOffSchedule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


