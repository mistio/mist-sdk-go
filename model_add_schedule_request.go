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

// AddScheduleRequest struct for AddScheduleRequest
type AddScheduleRequest struct {
	// The name of the schedule
	Name string `json:"name"`
	// The description of the schedule
	Description *string `json:"description,omitempty"`
	// Schedule status (enabled, disabled)
	Enabled *bool `json:"enabled,omitempty"`
	Selectors []Selector `json:"selectors,omitempty"`
	Actions []Action `json:"actions"`
	When *When `json:"when,omitempty"`
	// The date after that schedule starts. The format should be ΥΥΥΥ-ΜΜ-DD HH:MM:SS
	StartAfter *string `json:"start_after,omitempty"`
	// The date after that schedule expires. The format should be ΥΥΥΥ-ΜΜ-DD HH:MM:SS
	Expires *string `json:"expires,omitempty"`
	Reminder *AddScheduleRequestReminder `json:"reminder,omitempty"`
	// Decides if the schedule runs immediately of not
	RunImmediately *bool `json:"run_immediately,omitempty"`
}

// NewAddScheduleRequest instantiates a new AddScheduleRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddScheduleRequest(name string, actions []Action) *AddScheduleRequest {
	this := AddScheduleRequest{}
	this.Name = name
	this.Actions = actions
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
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AddScheduleRequest) GetNameOk() (*string, bool) {
	if o == nil {
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

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *AddScheduleRequest) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddScheduleRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *AddScheduleRequest) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *AddScheduleRequest) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetSelectors returns the Selectors field value if set, zero value otherwise.
func (o *AddScheduleRequest) GetSelectors() []Selector {
	if o == nil || o.Selectors == nil {
		var ret []Selector
		return ret
	}
	return o.Selectors
}

// GetSelectorsOk returns a tuple with the Selectors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddScheduleRequest) GetSelectorsOk() ([]Selector, bool) {
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
	o.Selectors = v
}

// GetActions returns the Actions field value
func (o *AddScheduleRequest) GetActions() []Action {
	if o == nil {
		var ret []Action
		return ret
	}

	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value
// and a boolean to check if the value has been set.
func (o *AddScheduleRequest) GetActionsOk() ([]Action, bool) {
	if o == nil {
		return nil, false
	}
	return o.Actions, true
}

// SetActions sets field value
func (o *AddScheduleRequest) SetActions(v []Action) {
	o.Actions = v
}

// GetWhen returns the When field value if set, zero value otherwise.
func (o *AddScheduleRequest) GetWhen() When {
	if o == nil || o.When == nil {
		var ret When
		return ret
	}
	return *o.When
}

// GetWhenOk returns a tuple with the When field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddScheduleRequest) GetWhenOk() (*When, bool) {
	if o == nil || o.When == nil {
		return nil, false
	}
	return o.When, true
}

// HasWhen returns a boolean if a field has been set.
func (o *AddScheduleRequest) HasWhen() bool {
	if o != nil && o.When != nil {
		return true
	}

	return false
}

// SetWhen gets a reference to the given When and assigns it to the When field.
func (o *AddScheduleRequest) SetWhen(v When) {
	o.When = &v
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

// GetExpires returns the Expires field value if set, zero value otherwise.
func (o *AddScheduleRequest) GetExpires() string {
	if o == nil || o.Expires == nil {
		var ret string
		return ret
	}
	return *o.Expires
}

// GetExpiresOk returns a tuple with the Expires field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddScheduleRequest) GetExpiresOk() (*string, bool) {
	if o == nil || o.Expires == nil {
		return nil, false
	}
	return o.Expires, true
}

// HasExpires returns a boolean if a field has been set.
func (o *AddScheduleRequest) HasExpires() bool {
	if o != nil && o.Expires != nil {
		return true
	}

	return false
}

// SetExpires gets a reference to the given string and assigns it to the Expires field.
func (o *AddScheduleRequest) SetExpires(v string) {
	o.Expires = &v
}

// GetReminder returns the Reminder field value if set, zero value otherwise.
func (o *AddScheduleRequest) GetReminder() AddScheduleRequestReminder {
	if o == nil || o.Reminder == nil {
		var ret AddScheduleRequestReminder
		return ret
	}
	return *o.Reminder
}

// GetReminderOk returns a tuple with the Reminder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddScheduleRequest) GetReminderOk() (*AddScheduleRequestReminder, bool) {
	if o == nil || o.Reminder == nil {
		return nil, false
	}
	return o.Reminder, true
}

// HasReminder returns a boolean if a field has been set.
func (o *AddScheduleRequest) HasReminder() bool {
	if o != nil && o.Reminder != nil {
		return true
	}

	return false
}

// SetReminder gets a reference to the given AddScheduleRequestReminder and assigns it to the Reminder field.
func (o *AddScheduleRequest) SetReminder(v AddScheduleRequestReminder) {
	o.Reminder = &v
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
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Selectors != nil {
		toSerialize["selectors"] = o.Selectors
	}
	if true {
		toSerialize["actions"] = o.Actions
	}
	if o.When != nil {
		toSerialize["when"] = o.When
	}
	if o.StartAfter != nil {
		toSerialize["start_after"] = o.StartAfter
	}
	if o.Expires != nil {
		toSerialize["expires"] = o.Expires
	}
	if o.Reminder != nil {
		toSerialize["reminder"] = o.Reminder
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


