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
	"time"
)

// When struct for When
type When struct {
	ScheduleType *string `json:"schedule_type,omitempty"`
	// When one_off schedule should run, e.g 2021-09-22T18:19:28Z
	Datetime time.Time `json:"datetime"`
	Minute string `json:"minute"`
	Hour string `json:"hour"`
	DayOfMonth string `json:"day_of_month"`
	MonthOfYear string `json:"month_of_year"`
	DayOfWeek string `json:"day_of_week"`
	// The datetime when schedule should start running, e.g 2021-09-22T18:19:28Z
	StartAfter *time.Time `json:"start_after,omitempty"`
	// The datetime when schedule should expire, e.g 2021-09-22T18:19:28Z
	Expires *time.Time `json:"expires,omitempty"`
	MaxRunCount *int32 `json:"max_run_count,omitempty"`
	Every int32 `json:"every"`
	Period string `json:"period"`
}

// NewWhen instantiates a new When object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWhen(datetime time.Time, minute string, hour string, dayOfMonth string, monthOfYear string, dayOfWeek string, every int32, period string) *When {
	this := When{}
	this.Datetime = datetime
	this.Minute = minute
	this.Hour = hour
	this.DayOfMonth = dayOfMonth
	this.MonthOfYear = monthOfYear
	this.DayOfWeek = dayOfWeek
	this.Every = every
	this.Period = period
	return &this
}

// NewWhenWithDefaults instantiates a new When object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWhenWithDefaults() *When {
	this := When{}
	return &this
}

// GetScheduleType returns the ScheduleType field value if set, zero value otherwise.
func (o *When) GetScheduleType() string {
	if o == nil || o.ScheduleType == nil {
		var ret string
		return ret
	}
	return *o.ScheduleType
}

// GetScheduleTypeOk returns a tuple with the ScheduleType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *When) GetScheduleTypeOk() (*string, bool) {
	if o == nil || o.ScheduleType == nil {
		return nil, false
	}
	return o.ScheduleType, true
}

// HasScheduleType returns a boolean if a field has been set.
func (o *When) HasScheduleType() bool {
	if o != nil && o.ScheduleType != nil {
		return true
	}

	return false
}

// SetScheduleType gets a reference to the given string and assigns it to the ScheduleType field.
func (o *When) SetScheduleType(v string) {
	o.ScheduleType = &v
}

// GetDatetime returns the Datetime field value
func (o *When) GetDatetime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *When) GetDatetimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *When) SetDatetime(v time.Time) {
	o.Datetime = v
}

// GetMinute returns the Minute field value
func (o *When) GetMinute() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Minute
}

// GetMinuteOk returns a tuple with the Minute field value
// and a boolean to check if the value has been set.
func (o *When) GetMinuteOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Minute, true
}

// SetMinute sets field value
func (o *When) SetMinute(v string) {
	o.Minute = v
}

// GetHour returns the Hour field value
func (o *When) GetHour() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hour
}

// GetHourOk returns a tuple with the Hour field value
// and a boolean to check if the value has been set.
func (o *When) GetHourOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hour, true
}

// SetHour sets field value
func (o *When) SetHour(v string) {
	o.Hour = v
}

// GetDayOfMonth returns the DayOfMonth field value
func (o *When) GetDayOfMonth() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DayOfMonth
}

// GetDayOfMonthOk returns a tuple with the DayOfMonth field value
// and a boolean to check if the value has been set.
func (o *When) GetDayOfMonthOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DayOfMonth, true
}

// SetDayOfMonth sets field value
func (o *When) SetDayOfMonth(v string) {
	o.DayOfMonth = v
}

// GetMonthOfYear returns the MonthOfYear field value
func (o *When) GetMonthOfYear() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MonthOfYear
}

// GetMonthOfYearOk returns a tuple with the MonthOfYear field value
// and a boolean to check if the value has been set.
func (o *When) GetMonthOfYearOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MonthOfYear, true
}

// SetMonthOfYear sets field value
func (o *When) SetMonthOfYear(v string) {
	o.MonthOfYear = v
}

// GetDayOfWeek returns the DayOfWeek field value
func (o *When) GetDayOfWeek() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DayOfWeek
}

// GetDayOfWeekOk returns a tuple with the DayOfWeek field value
// and a boolean to check if the value has been set.
func (o *When) GetDayOfWeekOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DayOfWeek, true
}

// SetDayOfWeek sets field value
func (o *When) SetDayOfWeek(v string) {
	o.DayOfWeek = v
}

// GetStartAfter returns the StartAfter field value if set, zero value otherwise.
func (o *When) GetStartAfter() time.Time {
	if o == nil || o.StartAfter == nil {
		var ret time.Time
		return ret
	}
	return *o.StartAfter
}

// GetStartAfterOk returns a tuple with the StartAfter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *When) GetStartAfterOk() (*time.Time, bool) {
	if o == nil || o.StartAfter == nil {
		return nil, false
	}
	return o.StartAfter, true
}

// HasStartAfter returns a boolean if a field has been set.
func (o *When) HasStartAfter() bool {
	if o != nil && o.StartAfter != nil {
		return true
	}

	return false
}

// SetStartAfter gets a reference to the given time.Time and assigns it to the StartAfter field.
func (o *When) SetStartAfter(v time.Time) {
	o.StartAfter = &v
}

// GetExpires returns the Expires field value if set, zero value otherwise.
func (o *When) GetExpires() time.Time {
	if o == nil || o.Expires == nil {
		var ret time.Time
		return ret
	}
	return *o.Expires
}

// GetExpiresOk returns a tuple with the Expires field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *When) GetExpiresOk() (*time.Time, bool) {
	if o == nil || o.Expires == nil {
		return nil, false
	}
	return o.Expires, true
}

// HasExpires returns a boolean if a field has been set.
func (o *When) HasExpires() bool {
	if o != nil && o.Expires != nil {
		return true
	}

	return false
}

// SetExpires gets a reference to the given time.Time and assigns it to the Expires field.
func (o *When) SetExpires(v time.Time) {
	o.Expires = &v
}

// GetMaxRunCount returns the MaxRunCount field value if set, zero value otherwise.
func (o *When) GetMaxRunCount() int32 {
	if o == nil || o.MaxRunCount == nil {
		var ret int32
		return ret
	}
	return *o.MaxRunCount
}

// GetMaxRunCountOk returns a tuple with the MaxRunCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *When) GetMaxRunCountOk() (*int32, bool) {
	if o == nil || o.MaxRunCount == nil {
		return nil, false
	}
	return o.MaxRunCount, true
}

// HasMaxRunCount returns a boolean if a field has been set.
func (o *When) HasMaxRunCount() bool {
	if o != nil && o.MaxRunCount != nil {
		return true
	}

	return false
}

// SetMaxRunCount gets a reference to the given int32 and assigns it to the MaxRunCount field.
func (o *When) SetMaxRunCount(v int32) {
	o.MaxRunCount = &v
}

// GetEvery returns the Every field value
func (o *When) GetEvery() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Every
}

// GetEveryOk returns a tuple with the Every field value
// and a boolean to check if the value has been set.
func (o *When) GetEveryOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Every, true
}

// SetEvery sets field value
func (o *When) SetEvery(v int32) {
	o.Every = v
}

// GetPeriod returns the Period field value
func (o *When) GetPeriod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Period
}

// GetPeriodOk returns a tuple with the Period field value
// and a boolean to check if the value has been set.
func (o *When) GetPeriodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Period, true
}

// SetPeriod sets field value
func (o *When) SetPeriod(v string) {
	o.Period = v
}

func (o When) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ScheduleType != nil {
		toSerialize["schedule_type"] = o.ScheduleType
	}
	if true {
		toSerialize["datetime"] = o.Datetime
	}
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
	if o.StartAfter != nil {
		toSerialize["start_after"] = o.StartAfter
	}
	if o.Expires != nil {
		toSerialize["expires"] = o.Expires
	}
	if o.MaxRunCount != nil {
		toSerialize["max_run_count"] = o.MaxRunCount
	}
	if true {
		toSerialize["every"] = o.Every
	}
	if true {
		toSerialize["period"] = o.Period
	}
	return json.Marshal(toSerialize)
}

type NullableWhen struct {
	value *When
	isSet bool
}

func (v NullableWhen) Get() *When {
	return v.value
}

func (v *NullableWhen) Set(val *When) {
	v.value = val
	v.isSet = true
}

func (v NullableWhen) IsSet() bool {
	return v.isSet
}

func (v *NullableWhen) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWhen(val *When) *NullableWhen {
	return &NullableWhen{value: val, isSet: true}
}

func (v NullableWhen) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWhen) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


