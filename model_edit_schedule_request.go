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

// EditScheduleRequest struct for EditScheduleRequest
type EditScheduleRequest struct {
	// The name of the schedule that is about to be edited
	Name *string `json:"name,omitempty"`
	// The description of the schedule that is about to be edited
	Description *string `json:"description,omitempty"`
	// Schedule status (enabled, disabled)
	Enabled *bool `json:"enabled,omitempty"`
	// Edit the action that a schedule performs on a resource
	Action *string `json:"action,omitempty"`
	// The id of the script that schedule is about to run
	ScriptId *string `json:"script_id,omitempty"`
	// Edit schedule parameters
	Params *string `json:"params,omitempty"`
	Selectors *[]Selector `json:"selectors,omitempty"`
	// Edit the type of the schedule
	ScheduleType *string `json:"schedule_type,omitempty"`
	// In case of One Off schedule type the date string that schedule runs (The format should be ΥΥΥΥ-ΜΜ-DD HH:MM:SS). In case of Interval and Crontab schedule types a JSON string with need time unit values. For Interval schedule type interval integer value and period string value needed. For Crontab schedule type minute, hour, day_of_week, day_of_month and month_of_year string values needed.
	ScheduleEntry *string `json:"schedule_entry,omitempty"`
	// Edit the date after that schedule starts. The format should be ΥΥΥΥ-ΜΜ-DD HH:MM:SS
	StartAfter *string `json:"start_after,omitempty"`
}

// NewEditScheduleRequest instantiates a new EditScheduleRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEditScheduleRequest() *EditScheduleRequest {
	this := EditScheduleRequest{}
	return &this
}

// NewEditScheduleRequestWithDefaults instantiates a new EditScheduleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEditScheduleRequestWithDefaults() *EditScheduleRequest {
	this := EditScheduleRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EditScheduleRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditScheduleRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EditScheduleRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *EditScheduleRequest) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EditScheduleRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditScheduleRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EditScheduleRequest) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *EditScheduleRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *EditScheduleRequest) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditScheduleRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *EditScheduleRequest) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *EditScheduleRequest) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *EditScheduleRequest) GetAction() string {
	if o == nil || o.Action == nil {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditScheduleRequest) GetActionOk() (*string, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *EditScheduleRequest) HasAction() bool {
	if o != nil && o.Action != nil {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *EditScheduleRequest) SetAction(v string) {
	o.Action = &v
}

// GetScriptId returns the ScriptId field value if set, zero value otherwise.
func (o *EditScheduleRequest) GetScriptId() string {
	if o == nil || o.ScriptId == nil {
		var ret string
		return ret
	}
	return *o.ScriptId
}

// GetScriptIdOk returns a tuple with the ScriptId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditScheduleRequest) GetScriptIdOk() (*string, bool) {
	if o == nil || o.ScriptId == nil {
		return nil, false
	}
	return o.ScriptId, true
}

// HasScriptId returns a boolean if a field has been set.
func (o *EditScheduleRequest) HasScriptId() bool {
	if o != nil && o.ScriptId != nil {
		return true
	}

	return false
}

// SetScriptId gets a reference to the given string and assigns it to the ScriptId field.
func (o *EditScheduleRequest) SetScriptId(v string) {
	o.ScriptId = &v
}

// GetParams returns the Params field value if set, zero value otherwise.
func (o *EditScheduleRequest) GetParams() string {
	if o == nil || o.Params == nil {
		var ret string
		return ret
	}
	return *o.Params
}

// GetParamsOk returns a tuple with the Params field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditScheduleRequest) GetParamsOk() (*string, bool) {
	if o == nil || o.Params == nil {
		return nil, false
	}
	return o.Params, true
}

// HasParams returns a boolean if a field has been set.
func (o *EditScheduleRequest) HasParams() bool {
	if o != nil && o.Params != nil {
		return true
	}

	return false
}

// SetParams gets a reference to the given string and assigns it to the Params field.
func (o *EditScheduleRequest) SetParams(v string) {
	o.Params = &v
}

// GetSelectors returns the Selectors field value if set, zero value otherwise.
func (o *EditScheduleRequest) GetSelectors() []Selector {
	if o == nil || o.Selectors == nil {
		var ret []Selector
		return ret
	}
	return *o.Selectors
}

// GetSelectorsOk returns a tuple with the Selectors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditScheduleRequest) GetSelectorsOk() (*[]Selector, bool) {
	if o == nil || o.Selectors == nil {
		return nil, false
	}
	return o.Selectors, true
}

// HasSelectors returns a boolean if a field has been set.
func (o *EditScheduleRequest) HasSelectors() bool {
	if o != nil && o.Selectors != nil {
		return true
	}

	return false
}

// SetSelectors gets a reference to the given []Selector and assigns it to the Selectors field.
func (o *EditScheduleRequest) SetSelectors(v []Selector) {
	o.Selectors = &v
}

// GetScheduleType returns the ScheduleType field value if set, zero value otherwise.
func (o *EditScheduleRequest) GetScheduleType() string {
	if o == nil || o.ScheduleType == nil {
		var ret string
		return ret
	}
	return *o.ScheduleType
}

// GetScheduleTypeOk returns a tuple with the ScheduleType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditScheduleRequest) GetScheduleTypeOk() (*string, bool) {
	if o == nil || o.ScheduleType == nil {
		return nil, false
	}
	return o.ScheduleType, true
}

// HasScheduleType returns a boolean if a field has been set.
func (o *EditScheduleRequest) HasScheduleType() bool {
	if o != nil && o.ScheduleType != nil {
		return true
	}

	return false
}

// SetScheduleType gets a reference to the given string and assigns it to the ScheduleType field.
func (o *EditScheduleRequest) SetScheduleType(v string) {
	o.ScheduleType = &v
}

// GetScheduleEntry returns the ScheduleEntry field value if set, zero value otherwise.
func (o *EditScheduleRequest) GetScheduleEntry() string {
	if o == nil || o.ScheduleEntry == nil {
		var ret string
		return ret
	}
	return *o.ScheduleEntry
}

// GetScheduleEntryOk returns a tuple with the ScheduleEntry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditScheduleRequest) GetScheduleEntryOk() (*string, bool) {
	if o == nil || o.ScheduleEntry == nil {
		return nil, false
	}
	return o.ScheduleEntry, true
}

// HasScheduleEntry returns a boolean if a field has been set.
func (o *EditScheduleRequest) HasScheduleEntry() bool {
	if o != nil && o.ScheduleEntry != nil {
		return true
	}

	return false
}

// SetScheduleEntry gets a reference to the given string and assigns it to the ScheduleEntry field.
func (o *EditScheduleRequest) SetScheduleEntry(v string) {
	o.ScheduleEntry = &v
}

// GetStartAfter returns the StartAfter field value if set, zero value otherwise.
func (o *EditScheduleRequest) GetStartAfter() string {
	if o == nil || o.StartAfter == nil {
		var ret string
		return ret
	}
	return *o.StartAfter
}

// GetStartAfterOk returns a tuple with the StartAfter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditScheduleRequest) GetStartAfterOk() (*string, bool) {
	if o == nil || o.StartAfter == nil {
		return nil, false
	}
	return o.StartAfter, true
}

// HasStartAfter returns a boolean if a field has been set.
func (o *EditScheduleRequest) HasStartAfter() bool {
	if o != nil && o.StartAfter != nil {
		return true
	}

	return false
}

// SetStartAfter gets a reference to the given string and assigns it to the StartAfter field.
func (o *EditScheduleRequest) SetStartAfter(v string) {
	o.StartAfter = &v
}

func (o EditScheduleRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Action != nil {
		toSerialize["action"] = o.Action
	}
	if o.ScriptId != nil {
		toSerialize["script_id"] = o.ScriptId
	}
	if o.Params != nil {
		toSerialize["params"] = o.Params
	}
	if o.Selectors != nil {
		toSerialize["selectors"] = o.Selectors
	}
	if o.ScheduleType != nil {
		toSerialize["schedule_type"] = o.ScheduleType
	}
	if o.ScheduleEntry != nil {
		toSerialize["schedule_entry"] = o.ScheduleEntry
	}
	if o.StartAfter != nil {
		toSerialize["start_after"] = o.StartAfter
	}
	return json.Marshal(toSerialize)
}

type NullableEditScheduleRequest struct {
	value *EditScheduleRequest
	isSet bool
}

func (v NullableEditScheduleRequest) Get() *EditScheduleRequest {
	return v.value
}

func (v *NullableEditScheduleRequest) Set(val *EditScheduleRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEditScheduleRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEditScheduleRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEditScheduleRequest(val *EditScheduleRequest) *NullableEditScheduleRequest {
	return &NullableEditScheduleRequest{value: val, isSet: true}
}

func (v NullableEditScheduleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEditScheduleRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

