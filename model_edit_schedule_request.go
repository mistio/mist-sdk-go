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

// EditScheduleRequest struct for EditScheduleRequest
type EditScheduleRequest struct {
	// The name of the schedule that is about to be edited
	Name *string `json:"name,omitempty"`
	// The description of the schedule that is about to be edited
	Description *string `json:"description,omitempty"`
	// Schedule status (enabled, disabled)
	Enabled *bool `json:"enabled,omitempty"`
	Selectors []Selector `json:"selectors,omitempty"`
	Actions []Action `json:"actions,omitempty"`
	When *When `json:"when,omitempty"`
	// The date after that schedule expires. The format should be ΥΥΥΥ-ΜΜ-DD HH:MM:SS
	Expires *string `json:"expires,omitempty"`
	Reminder *AddScheduleRequestReminder `json:"reminder,omitempty"`
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

// GetSelectors returns the Selectors field value if set, zero value otherwise.
func (o *EditScheduleRequest) GetSelectors() []Selector {
	if o == nil || o.Selectors == nil {
		var ret []Selector
		return ret
	}
	return o.Selectors
}

// GetSelectorsOk returns a tuple with the Selectors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditScheduleRequest) GetSelectorsOk() ([]Selector, bool) {
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
	o.Selectors = v
}

// GetActions returns the Actions field value if set, zero value otherwise.
func (o *EditScheduleRequest) GetActions() []Action {
	if o == nil || o.Actions == nil {
		var ret []Action
		return ret
	}
	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditScheduleRequest) GetActionsOk() ([]Action, bool) {
	if o == nil || o.Actions == nil {
		return nil, false
	}
	return o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *EditScheduleRequest) HasActions() bool {
	if o != nil && o.Actions != nil {
		return true
	}

	return false
}

// SetActions gets a reference to the given []Action and assigns it to the Actions field.
func (o *EditScheduleRequest) SetActions(v []Action) {
	o.Actions = v
}

// GetWhen returns the When field value if set, zero value otherwise.
func (o *EditScheduleRequest) GetWhen() When {
	if o == nil || o.When == nil {
		var ret When
		return ret
	}
	return *o.When
}

// GetWhenOk returns a tuple with the When field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditScheduleRequest) GetWhenOk() (*When, bool) {
	if o == nil || o.When == nil {
		return nil, false
	}
	return o.When, true
}

// HasWhen returns a boolean if a field has been set.
func (o *EditScheduleRequest) HasWhen() bool {
	if o != nil && o.When != nil {
		return true
	}

	return false
}

// SetWhen gets a reference to the given When and assigns it to the When field.
func (o *EditScheduleRequest) SetWhen(v When) {
	o.When = &v
}

// GetExpires returns the Expires field value if set, zero value otherwise.
func (o *EditScheduleRequest) GetExpires() string {
	if o == nil || o.Expires == nil {
		var ret string
		return ret
	}
	return *o.Expires
}

// GetExpiresOk returns a tuple with the Expires field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditScheduleRequest) GetExpiresOk() (*string, bool) {
	if o == nil || o.Expires == nil {
		return nil, false
	}
	return o.Expires, true
}

// HasExpires returns a boolean if a field has been set.
func (o *EditScheduleRequest) HasExpires() bool {
	if o != nil && o.Expires != nil {
		return true
	}

	return false
}

// SetExpires gets a reference to the given string and assigns it to the Expires field.
func (o *EditScheduleRequest) SetExpires(v string) {
	o.Expires = &v
}

// GetReminder returns the Reminder field value if set, zero value otherwise.
func (o *EditScheduleRequest) GetReminder() AddScheduleRequestReminder {
	if o == nil || o.Reminder == nil {
		var ret AddScheduleRequestReminder
		return ret
	}
	return *o.Reminder
}

// GetReminderOk returns a tuple with the Reminder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditScheduleRequest) GetReminderOk() (*AddScheduleRequestReminder, bool) {
	if o == nil || o.Reminder == nil {
		return nil, false
	}
	return o.Reminder, true
}

// HasReminder returns a boolean if a field has been set.
func (o *EditScheduleRequest) HasReminder() bool {
	if o != nil && o.Reminder != nil {
		return true
	}

	return false
}

// SetReminder gets a reference to the given AddScheduleRequestReminder and assigns it to the Reminder field.
func (o *EditScheduleRequest) SetReminder(v AddScheduleRequestReminder) {
	o.Reminder = &v
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
	if o.Selectors != nil {
		toSerialize["selectors"] = o.Selectors
	}
	if o.Actions != nil {
		toSerialize["actions"] = o.Actions
	}
	if o.When != nil {
		toSerialize["when"] = o.When
	}
	if o.Expires != nil {
		toSerialize["expires"] = o.Expires
	}
	if o.Reminder != nil {
		toSerialize["reminder"] = o.Reminder
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


