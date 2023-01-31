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

// CronSchedule crontab schedule
type CronSchedule struct {
	Minute string `json:"minute"`
	Hour string `json:"hour"`
	DayOfMonth string `json:"day_of_month"`
	MonthOfYear string `json:"month_of_year"`
	DayOfWeek string `json:"day_of_week"`
	MaxRunCount *int32 `json:"max_run_count,omitempty"`
}

// NewCronSchedule instantiates a new CronSchedule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCronSchedule(minute string, hour string, dayOfMonth string, monthOfYear string, dayOfWeek string) *CronSchedule {
	this := CronSchedule{}
	this.Minute = minute
	this.Hour = hour
	this.DayOfMonth = dayOfMonth
	this.MonthOfYear = monthOfYear
	this.DayOfWeek = dayOfWeek
	return &this
}

// NewCronScheduleWithDefaults instantiates a new CronSchedule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCronScheduleWithDefaults() *CronSchedule {
	this := CronSchedule{}
	return &this
}

// GetMinute returns the Minute field value
func (o *CronSchedule) GetMinute() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Minute
}

// GetMinuteOk returns a tuple with the Minute field value
// and a boolean to check if the value has been set.
func (o *CronSchedule) GetMinuteOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Minute, true
}

// SetMinute sets field value
func (o *CronSchedule) SetMinute(v string) {
	o.Minute = v
}

// GetHour returns the Hour field value
func (o *CronSchedule) GetHour() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hour
}

// GetHourOk returns a tuple with the Hour field value
// and a boolean to check if the value has been set.
func (o *CronSchedule) GetHourOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hour, true
}

// SetHour sets field value
func (o *CronSchedule) SetHour(v string) {
	o.Hour = v
}

// GetDayOfMonth returns the DayOfMonth field value
func (o *CronSchedule) GetDayOfMonth() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DayOfMonth
}

// GetDayOfMonthOk returns a tuple with the DayOfMonth field value
// and a boolean to check if the value has been set.
func (o *CronSchedule) GetDayOfMonthOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DayOfMonth, true
}

// SetDayOfMonth sets field value
func (o *CronSchedule) SetDayOfMonth(v string) {
	o.DayOfMonth = v
}

// GetMonthOfYear returns the MonthOfYear field value
func (o *CronSchedule) GetMonthOfYear() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MonthOfYear
}

// GetMonthOfYearOk returns a tuple with the MonthOfYear field value
// and a boolean to check if the value has been set.
func (o *CronSchedule) GetMonthOfYearOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MonthOfYear, true
}

// SetMonthOfYear sets field value
func (o *CronSchedule) SetMonthOfYear(v string) {
	o.MonthOfYear = v
}

// GetDayOfWeek returns the DayOfWeek field value
func (o *CronSchedule) GetDayOfWeek() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DayOfWeek
}

// GetDayOfWeekOk returns a tuple with the DayOfWeek field value
// and a boolean to check if the value has been set.
func (o *CronSchedule) GetDayOfWeekOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DayOfWeek, true
}

// SetDayOfWeek sets field value
func (o *CronSchedule) SetDayOfWeek(v string) {
	o.DayOfWeek = v
}

// GetMaxRunCount returns the MaxRunCount field value if set, zero value otherwise.
func (o *CronSchedule) GetMaxRunCount() int32 {
	if o == nil || o.MaxRunCount == nil {
		var ret int32
		return ret
	}
	return *o.MaxRunCount
}

// GetMaxRunCountOk returns a tuple with the MaxRunCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CronSchedule) GetMaxRunCountOk() (*int32, bool) {
	if o == nil || o.MaxRunCount == nil {
		return nil, false
	}
	return o.MaxRunCount, true
}

// HasMaxRunCount returns a boolean if a field has been set.
func (o *CronSchedule) HasMaxRunCount() bool {
	if o != nil && o.MaxRunCount != nil {
		return true
	}

	return false
}

// SetMaxRunCount gets a reference to the given int32 and assigns it to the MaxRunCount field.
func (o *CronSchedule) SetMaxRunCount(v int32) {
	o.MaxRunCount = &v
}

func (o CronSchedule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["minute"] = o.Minute
	}
	if true {
		toSerialize["hour"] = o.Hour
	}
	if true {
		toSerialize["day_of_month"] = o.DayOfMonth
	}
	if true {
		toSerialize["month_of_year"] = o.MonthOfYear
	}
	if true {
		toSerialize["day_of_week"] = o.DayOfWeek
	}
	if o.MaxRunCount != nil {
		toSerialize["max_run_count"] = o.MaxRunCount
	}
	return json.Marshal(toSerialize)
}

type NullableCronSchedule struct {
	value *CronSchedule
	isSet bool
}

func (v NullableCronSchedule) Get() *CronSchedule {
	return v.value
}

func (v *NullableCronSchedule) Set(val *CronSchedule) {
	v.value = val
	v.isSet = true
}

func (v NullableCronSchedule) IsSet() bool {
	return v.isSet
}

func (v *NullableCronSchedule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCronSchedule(val *CronSchedule) *NullableCronSchedule {
	return &NullableCronSchedule{value: val, isSet: true}
}

func (v NullableCronSchedule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCronSchedule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


