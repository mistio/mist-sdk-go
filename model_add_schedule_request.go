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

// AddScheduleRequest struct for AddScheduleRequest
type AddScheduleRequest struct {
	Name string `json:"name"`
	Description *string `json:"description,omitempty"`
	TaskEnabled *bool `json:"task_enabled,omitempty"`
	Action string `json:"action"`
	Params *string `json:"params,omitempty"`
	ResourceType *string `json:"resource_type,omitempty"`
	Selectors *[]Selector `json:"selectors,omitempty"`
	ScheduleType *string `json:"schedule_type,omitempty"`
	// format should be ΥΥΥΥ-ΜΜ-DD HH:MM:SS
	ScheduleEntry *string `json:"schedule_entry,omitempty"`
	// format should be ΥΥΥΥ-ΜΜ-DD HH:MM:SS
	StartAfter *string `json:"start_after,omitempty"`
	RunImmediately *bool `json:"run_immediately,omitempty"`
}

// NewAddScheduleRequest instantiates a new AddScheduleRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddScheduleRequest(name string, action string, ) *AddScheduleRequest {
	this := AddScheduleRequest{}
	this.Name = name
	this.Action = action
	return &this
}

// NewAddScheduleRequestWithDefaults instantiates a new AddScheduleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddScheduleRequestWithDefaults() *AddScheduleRequest {
	this := AddScheduleRequest{}
	return &this
}

// GetName returns the Name field value
func (o *AddScheduleRequest) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AddScheduleRequest) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AddScheduleRequest) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddScheduleRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddScheduleRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddScheduleRequest) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddScheduleRequest) SetDescription(v string) {
	o.Description = &v
}

// GetTaskEnabled returns the TaskEnabled field value if set, zero value otherwise.
func (o *AddScheduleRequest) GetTaskEnabled() bool {
	if o == nil || o.TaskEnabled == nil {
		var ret bool
		return ret
	}
	return *o.TaskEnabled
}

// GetTaskEnabledOk returns a tuple with the TaskEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddScheduleRequest) GetTaskEnabledOk() (*bool, bool) {
	if o == nil || o.TaskEnabled == nil {
		return nil, false
	}
	return o.TaskEnabled, true
}

// HasTaskEnabled returns a boolean if a field has been set.
func (o *AddScheduleRequest) HasTaskEnabled() bool {
	if o != nil && o.TaskEnabled != nil {
		return true
	}

	return false
}

// SetTaskEnabled gets a reference to the given bool and assigns it to the TaskEnabled field.
func (o *AddScheduleRequest) SetTaskEnabled(v bool) {
	o.TaskEnabled = &v
}

// GetAction returns the Action field value
func (o *AddScheduleRequest) GetAction() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *AddScheduleRequest) GetActionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *AddScheduleRequest) SetAction(v string) {
	o.Action = v
}

// GetParams returns the Params field value if set, zero value otherwise.
func (o *AddScheduleRequest) GetParams() string {
	if o == nil || o.Params == nil {
		var ret string
		return ret
	}
	return *o.Params
}

// GetParamsOk returns a tuple with the Params field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddScheduleRequest) GetParamsOk() (*string, bool) {
	if o == nil || o.Params == nil {
		return nil, false
	}
	return o.Params, true
}

// HasParams returns a boolean if a field has been set.
func (o *AddScheduleRequest) HasParams() bool {
	if o != nil && o.Params != nil {
		return true
	}

	return false
}

// SetParams gets a reference to the given string and assigns it to the Params field.
func (o *AddScheduleRequest) SetParams(v string) {
	o.Params = &v
}

// GetResourceType returns the ResourceType field value if set, zero value otherwise.
func (o *AddScheduleRequest) GetResourceType() string {
	if o == nil || o.ResourceType == nil {
		var ret string
		return ret
	}
	return *o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddScheduleRequest) GetResourceTypeOk() (*string, bool) {
	if o == nil || o.ResourceType == nil {
		return nil, false
	}
	return o.ResourceType, true
}

// HasResourceType returns a boolean if a field has been set.
func (o *AddScheduleRequest) HasResourceType() bool {
	if o != nil && o.ResourceType != nil {
		return true
	}

	return false
}

// SetResourceType gets a reference to the given string and assigns it to the ResourceType field.
func (o *AddScheduleRequest) SetResourceType(v string) {
	o.ResourceType = &v
}

// GetSelectors returns the Selectors field value if set, zero value otherwise.
func (o *AddScheduleRequest) GetSelectors() []Selector {
	if o == nil || o.Selectors == nil {
		var ret []Selector
		return ret
	}
	return *o.Selectors
}

// GetSelectorsOk returns a tuple with the Selectors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddScheduleRequest) GetSelectorsOk() (*[]Selector, bool) {
	if o == nil || o.Selectors == nil {
		return nil, false
	}
	return o.Selectors, true
}

// HasSelectors returns a boolean if a field has been set.
func (o *AddScheduleRequest) HasSelectors() bool {
	if o != nil && o.Selectors != nil {
		return true
	}

	return false
}

// SetSelectors gets a reference to the given []Selector and assigns it to the Selectors field.
func (o *AddScheduleRequest) SetSelectors(v []Selector) {
	o.Selectors = &v
}

// GetScheduleType returns the ScheduleType field value if set, zero value otherwise.
func (o *AddScheduleRequest) GetScheduleType() string {
	if o == nil || o.ScheduleType == nil {
		var ret string
		return ret
	}
	return *o.ScheduleType
}

// GetScheduleTypeOk returns a tuple with the ScheduleType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddScheduleRequest) GetScheduleTypeOk() (*string, bool) {
	if o == nil || o.ScheduleType == nil {
		return nil, false
	}
	return o.ScheduleType, true
}

// HasScheduleType returns a boolean if a field has been set.
func (o *AddScheduleRequest) HasScheduleType() bool {
	if o != nil && o.ScheduleType != nil {
		return true
	}

	return false
}

// SetScheduleType gets a reference to the given string and assigns it to the ScheduleType field.
func (o *AddScheduleRequest) SetScheduleType(v string) {
	o.ScheduleType = &v
}

// GetScheduleEntry returns the ScheduleEntry field value if set, zero value otherwise.
func (o *AddScheduleRequest) GetScheduleEntry() string {
	if o == nil || o.ScheduleEntry == nil {
		var ret string
		return ret
	}
	return *o.ScheduleEntry
}

// GetScheduleEntryOk returns a tuple with the ScheduleEntry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddScheduleRequest) GetScheduleEntryOk() (*string, bool) {
	if o == nil || o.ScheduleEntry == nil {
		return nil, false
	}
	return o.ScheduleEntry, true
}

// HasScheduleEntry returns a boolean if a field has been set.
func (o *AddScheduleRequest) HasScheduleEntry() bool {
	if o != nil && o.ScheduleEntry != nil {
		return true
	}

	return false
}

// SetScheduleEntry gets a reference to the given string and assigns it to the ScheduleEntry field.
func (o *AddScheduleRequest) SetScheduleEntry(v string) {
	o.ScheduleEntry = &v
}

// GetStartAfter returns the StartAfter field value if set, zero value otherwise.
func (o *AddScheduleRequest) GetStartAfter() string {
	if o == nil || o.StartAfter == nil {
		var ret string
		return ret
	}
	return *o.StartAfter
}

// GetStartAfterOk returns a tuple with the StartAfter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddScheduleRequest) GetStartAfterOk() (*string, bool) {
	if o == nil || o.StartAfter == nil {
		return nil, false
	}
	return o.StartAfter, true
}

// HasStartAfter returns a boolean if a field has been set.
func (o *AddScheduleRequest) HasStartAfter() bool {
	if o != nil && o.StartAfter != nil {
		return true
	}

	return false
}

// SetStartAfter gets a reference to the given string and assigns it to the StartAfter field.
func (o *AddScheduleRequest) SetStartAfter(v string) {
	o.StartAfter = &v
}

// GetRunImmediately returns the RunImmediately field value if set, zero value otherwise.
func (o *AddScheduleRequest) GetRunImmediately() bool {
	if o == nil || o.RunImmediately == nil {
		var ret bool
		return ret
	}
	return *o.RunImmediately
}

// GetRunImmediatelyOk returns a tuple with the RunImmediately field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddScheduleRequest) GetRunImmediatelyOk() (*bool, bool) {
	if o == nil || o.RunImmediately == nil {
		return nil, false
	}
	return o.RunImmediately, true
}

// HasRunImmediately returns a boolean if a field has been set.
func (o *AddScheduleRequest) HasRunImmediately() bool {
	if o != nil && o.RunImmediately != nil {
		return true
	}

	return false
}

// SetRunImmediately gets a reference to the given bool and assigns it to the RunImmediately field.
func (o *AddScheduleRequest) SetRunImmediately(v bool) {
	o.RunImmediately = &v
}

func (o AddScheduleRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.TaskEnabled != nil {
		toSerialize["task_enabled"] = o.TaskEnabled
	}
	if true {
		toSerialize["action"] = o.Action
	}
	if o.Params != nil {
		toSerialize["params"] = o.Params
	}
	if o.ResourceType != nil {
		toSerialize["resource_type"] = o.ResourceType
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
	if o.RunImmediately != nil {
		toSerialize["run_immediately"] = o.RunImmediately
	}
	return json.Marshal(toSerialize)
}

type NullableAddScheduleRequest struct {
	value *AddScheduleRequest
	isSet bool
}

func (v NullableAddScheduleRequest) Get() *AddScheduleRequest {
	return v.value
}

func (v *NullableAddScheduleRequest) Set(val *AddScheduleRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddScheduleRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddScheduleRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddScheduleRequest(val *AddScheduleRequest) *NullableAddScheduleRequest {
	return &NullableAddScheduleRequest{value: val, isSet: true}
}

func (v NullableAddScheduleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddScheduleRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


