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

// Schedule struct for Schedule
type Schedule struct {
	// The id of the schedule
	Id *string `json:"id,omitempty"`
	// The name of the schedule
	Name *string `json:"name,omitempty"`
	// The description of the schedule
	Description *string `json:"description,omitempty"`
	// The tags related to the schedule
	Tags *map[string]interface{} `json:"tags,omitempty"`
	// Schedule status (enabled, disabled)
	Enabled *bool `json:"enabled,omitempty"`
	// The action that a schedule performs on a resource
	Action *string `json:"action,omitempty"`
	// The id of the script that schedule is about to run
	ScriptId *string `json:"script_id,omitempty"`
	// Schedule parameters
	Params *string `json:"params,omitempty"`
	Selectors *[]map[string]interface{} `json:"selectors,omitempty"`
	// The type of the schedule
	ScheduleType *string `json:"schedule_type,omitempty"`
	// The date that schedule starts. The format should be ΥΥΥΥ-ΜΜ-DD HH:MM:SS
	ScheduleEntry *string `json:"schedule_entry,omitempty"`
	// The date after that schedule starts. The format should be ΥΥΥΥ-ΜΜ-DD HH:MM:SS
	StartAfter *string `json:"start_after,omitempty"`
	// Decides if the schedule runs immediately of not
	RunImmediately *bool `json:"run_immediately,omitempty"`
	// The name of user that created the schedule
	CreatedBy *string `json:"created_by,omitempty"`
	// The name of user that owns the schedule
	OwnedBy *string `json:"owned_by,omitempty"`
}

// NewSchedule instantiates a new Schedule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchedule() *Schedule {
	this := Schedule{}
	return &this
}

// NewScheduleWithDefaults instantiates a new Schedule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScheduleWithDefaults() *Schedule {
	this := Schedule{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Schedule) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Schedule) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Schedule) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Schedule) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Schedule) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Schedule) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Schedule) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Schedule) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Schedule) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Schedule) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Schedule) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Schedule) SetDescription(v string) {
	o.Description = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Schedule) GetTags() map[string]interface{} {
	if o == nil || o.Tags == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Schedule) GetTagsOk() (*map[string]interface{}, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Schedule) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]interface{} and assigns it to the Tags field.
func (o *Schedule) SetTags(v map[string]interface{}) {
	o.Tags = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *Schedule) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Schedule) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *Schedule) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *Schedule) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *Schedule) GetAction() string {
	if o == nil || o.Action == nil {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Schedule) GetActionOk() (*string, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *Schedule) HasAction() bool {
	if o != nil && o.Action != nil {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *Schedule) SetAction(v string) {
	o.Action = &v
}

// GetScriptId returns the ScriptId field value if set, zero value otherwise.
func (o *Schedule) GetScriptId() string {
	if o == nil || o.ScriptId == nil {
		var ret string
		return ret
	}
	return *o.ScriptId
}

// GetScriptIdOk returns a tuple with the ScriptId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Schedule) GetScriptIdOk() (*string, bool) {
	if o == nil || o.ScriptId == nil {
		return nil, false
	}
	return o.ScriptId, true
}

// HasScriptId returns a boolean if a field has been set.
func (o *Schedule) HasScriptId() bool {
	if o != nil && o.ScriptId != nil {
		return true
	}

	return false
}

// SetScriptId gets a reference to the given string and assigns it to the ScriptId field.
func (o *Schedule) SetScriptId(v string) {
	o.ScriptId = &v
}

// GetParams returns the Params field value if set, zero value otherwise.
func (o *Schedule) GetParams() string {
	if o == nil || o.Params == nil {
		var ret string
		return ret
	}
	return *o.Params
}

// GetParamsOk returns a tuple with the Params field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Schedule) GetParamsOk() (*string, bool) {
	if o == nil || o.Params == nil {
		return nil, false
	}
	return o.Params, true
}

// HasParams returns a boolean if a field has been set.
func (o *Schedule) HasParams() bool {
	if o != nil && o.Params != nil {
		return true
	}

	return false
}

// SetParams gets a reference to the given string and assigns it to the Params field.
func (o *Schedule) SetParams(v string) {
	o.Params = &v
}

// GetSelectors returns the Selectors field value if set, zero value otherwise.
func (o *Schedule) GetSelectors() []map[string]interface{} {
	if o == nil || o.Selectors == nil {
		var ret []map[string]interface{}
		return ret
	}
	return *o.Selectors
}

// GetSelectorsOk returns a tuple with the Selectors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Schedule) GetSelectorsOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.Selectors == nil {
		return nil, false
	}
	return o.Selectors, true
}

// HasSelectors returns a boolean if a field has been set.
func (o *Schedule) HasSelectors() bool {
	if o != nil && o.Selectors != nil {
		return true
	}

	return false
}

// SetSelectors gets a reference to the given []map[string]interface{} and assigns it to the Selectors field.
func (o *Schedule) SetSelectors(v []map[string]interface{}) {
	o.Selectors = &v
}

// GetScheduleType returns the ScheduleType field value if set, zero value otherwise.
func (o *Schedule) GetScheduleType() string {
	if o == nil || o.ScheduleType == nil {
		var ret string
		return ret
	}
	return *o.ScheduleType
}

// GetScheduleTypeOk returns a tuple with the ScheduleType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Schedule) GetScheduleTypeOk() (*string, bool) {
	if o == nil || o.ScheduleType == nil {
		return nil, false
	}
	return o.ScheduleType, true
}

// HasScheduleType returns a boolean if a field has been set.
func (o *Schedule) HasScheduleType() bool {
	if o != nil && o.ScheduleType != nil {
		return true
	}

	return false
}

// SetScheduleType gets a reference to the given string and assigns it to the ScheduleType field.
func (o *Schedule) SetScheduleType(v string) {
	o.ScheduleType = &v
}

// GetScheduleEntry returns the ScheduleEntry field value if set, zero value otherwise.
func (o *Schedule) GetScheduleEntry() string {
	if o == nil || o.ScheduleEntry == nil {
		var ret string
		return ret
	}
	return *o.ScheduleEntry
}

// GetScheduleEntryOk returns a tuple with the ScheduleEntry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Schedule) GetScheduleEntryOk() (*string, bool) {
	if o == nil || o.ScheduleEntry == nil {
		return nil, false
	}
	return o.ScheduleEntry, true
}

// HasScheduleEntry returns a boolean if a field has been set.
func (o *Schedule) HasScheduleEntry() bool {
	if o != nil && o.ScheduleEntry != nil {
		return true
	}

	return false
}

// SetScheduleEntry gets a reference to the given string and assigns it to the ScheduleEntry field.
func (o *Schedule) SetScheduleEntry(v string) {
	o.ScheduleEntry = &v
}

// GetStartAfter returns the StartAfter field value if set, zero value otherwise.
func (o *Schedule) GetStartAfter() string {
	if o == nil || o.StartAfter == nil {
		var ret string
		return ret
	}
	return *o.StartAfter
}

// GetStartAfterOk returns a tuple with the StartAfter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Schedule) GetStartAfterOk() (*string, bool) {
	if o == nil || o.StartAfter == nil {
		return nil, false
	}
	return o.StartAfter, true
}

// HasStartAfter returns a boolean if a field has been set.
func (o *Schedule) HasStartAfter() bool {
	if o != nil && o.StartAfter != nil {
		return true
	}

	return false
}

// SetStartAfter gets a reference to the given string and assigns it to the StartAfter field.
func (o *Schedule) SetStartAfter(v string) {
	o.StartAfter = &v
}

// GetRunImmediately returns the RunImmediately field value if set, zero value otherwise.
func (o *Schedule) GetRunImmediately() bool {
	if o == nil || o.RunImmediately == nil {
		var ret bool
		return ret
	}
	return *o.RunImmediately
}

// GetRunImmediatelyOk returns a tuple with the RunImmediately field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Schedule) GetRunImmediatelyOk() (*bool, bool) {
	if o == nil || o.RunImmediately == nil {
		return nil, false
	}
	return o.RunImmediately, true
}

// HasRunImmediately returns a boolean if a field has been set.
func (o *Schedule) HasRunImmediately() bool {
	if o != nil && o.RunImmediately != nil {
		return true
	}

	return false
}

// SetRunImmediately gets a reference to the given bool and assigns it to the RunImmediately field.
func (o *Schedule) SetRunImmediately(v bool) {
	o.RunImmediately = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *Schedule) GetCreatedBy() string {
	if o == nil || o.CreatedBy == nil {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Schedule) GetCreatedByOk() (*string, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *Schedule) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *Schedule) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetOwnedBy returns the OwnedBy field value if set, zero value otherwise.
func (o *Schedule) GetOwnedBy() string {
	if o == nil || o.OwnedBy == nil {
		var ret string
		return ret
	}
	return *o.OwnedBy
}

// GetOwnedByOk returns a tuple with the OwnedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Schedule) GetOwnedByOk() (*string, bool) {
	if o == nil || o.OwnedBy == nil {
		return nil, false
	}
	return o.OwnedBy, true
}

// HasOwnedBy returns a boolean if a field has been set.
func (o *Schedule) HasOwnedBy() bool {
	if o != nil && o.OwnedBy != nil {
		return true
	}

	return false
}

// SetOwnedBy gets a reference to the given string and assigns it to the OwnedBy field.
func (o *Schedule) SetOwnedBy(v string) {
	o.OwnedBy = &v
}

func (o Schedule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
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
	if o.RunImmediately != nil {
		toSerialize["run_immediately"] = o.RunImmediately
	}
	if o.CreatedBy != nil {
		toSerialize["created_by"] = o.CreatedBy
	}
	if o.OwnedBy != nil {
		toSerialize["owned_by"] = o.OwnedBy
	}
	return json.Marshal(toSerialize)
}

type NullableSchedule struct {
	value *Schedule
	isSet bool
}

func (v NullableSchedule) Get() *Schedule {
	return v.value
}

func (v *NullableSchedule) Set(val *Schedule) {
	v.value = val
	v.isSet = true
}

func (v NullableSchedule) IsSet() bool {
	return v.isSet
}

func (v *NullableSchedule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchedule(val *Schedule) *NullableSchedule {
	return &NullableSchedule{value: val, isSet: true}
}

func (v NullableSchedule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchedule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


