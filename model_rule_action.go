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

// RuleAction struct for RuleAction
type RuleAction struct {
	// the action to be performed. Required by machine_action type 
	Action *string `json:"action,omitempty"`
	// the command to be executed. Required by the command type 
	Command *string `json:"command,omitempty"`
	// a list of e-mails to send a notification to. Can be used by a notification action (optional) 
	Emails *[]string `json:"emails,omitempty"`
	// a list of teams, denoted by their UUIDs, whose users will be notified. Can be used by a notification action (optional) 
	Teams *[]string `json:"teams,omitempty"`
	// the action's type: notification, machine_action, command 
	Type string `json:"type"`
	// a list of user to be notified, denoted by their UUIDs. Can be used by a notification action (optional) 
	Users *[]string `json:"users,omitempty"`
}

// NewRuleAction instantiates a new RuleAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRuleAction(type_ string, ) *RuleAction {
	this := RuleAction{}
	this.Type = type_
	return &this
}

// NewRuleActionWithDefaults instantiates a new RuleAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRuleActionWithDefaults() *RuleAction {
	this := RuleAction{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *RuleAction) GetAction() string {
	if o == nil || o.Action == nil {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleAction) GetActionOk() (*string, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *RuleAction) HasAction() bool {
	if o != nil && o.Action != nil {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *RuleAction) SetAction(v string) {
	o.Action = &v
}

// GetCommand returns the Command field value if set, zero value otherwise.
func (o *RuleAction) GetCommand() string {
	if o == nil || o.Command == nil {
		var ret string
		return ret
	}
	return *o.Command
}

// GetCommandOk returns a tuple with the Command field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleAction) GetCommandOk() (*string, bool) {
	if o == nil || o.Command == nil {
		return nil, false
	}
	return o.Command, true
}

// HasCommand returns a boolean if a field has been set.
func (o *RuleAction) HasCommand() bool {
	if o != nil && o.Command != nil {
		return true
	}

	return false
}

// SetCommand gets a reference to the given string and assigns it to the Command field.
func (o *RuleAction) SetCommand(v string) {
	o.Command = &v
}

// GetEmails returns the Emails field value if set, zero value otherwise.
func (o *RuleAction) GetEmails() []string {
	if o == nil || o.Emails == nil {
		var ret []string
		return ret
	}
	return *o.Emails
}

// GetEmailsOk returns a tuple with the Emails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleAction) GetEmailsOk() (*[]string, bool) {
	if o == nil || o.Emails == nil {
		return nil, false
	}
	return o.Emails, true
}

// HasEmails returns a boolean if a field has been set.
func (o *RuleAction) HasEmails() bool {
	if o != nil && o.Emails != nil {
		return true
	}

	return false
}

// SetEmails gets a reference to the given []string and assigns it to the Emails field.
func (o *RuleAction) SetEmails(v []string) {
	o.Emails = &v
}

// GetTeams returns the Teams field value if set, zero value otherwise.
func (o *RuleAction) GetTeams() []string {
	if o == nil || o.Teams == nil {
		var ret []string
		return ret
	}
	return *o.Teams
}

// GetTeamsOk returns a tuple with the Teams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleAction) GetTeamsOk() (*[]string, bool) {
	if o == nil || o.Teams == nil {
		return nil, false
	}
	return o.Teams, true
}

// HasTeams returns a boolean if a field has been set.
func (o *RuleAction) HasTeams() bool {
	if o != nil && o.Teams != nil {
		return true
	}

	return false
}

// SetTeams gets a reference to the given []string and assigns it to the Teams field.
func (o *RuleAction) SetTeams(v []string) {
	o.Teams = &v
}

// GetType returns the Type field value
func (o *RuleAction) GetType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RuleAction) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RuleAction) SetType(v string) {
	o.Type = v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *RuleAction) GetUsers() []string {
	if o == nil || o.Users == nil {
		var ret []string
		return ret
	}
	return *o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleAction) GetUsersOk() (*[]string, bool) {
	if o == nil || o.Users == nil {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *RuleAction) HasUsers() bool {
	if o != nil && o.Users != nil {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []string and assigns it to the Users field.
func (o *RuleAction) SetUsers(v []string) {
	o.Users = &v
}

func (o RuleAction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Action != nil {
		toSerialize["action"] = o.Action
	}
	if o.Command != nil {
		toSerialize["command"] = o.Command
	}
	if o.Emails != nil {
		toSerialize["emails"] = o.Emails
	}
	if o.Teams != nil {
		toSerialize["teams"] = o.Teams
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.Users != nil {
		toSerialize["users"] = o.Users
	}
	return json.Marshal(toSerialize)
}

type NullableRuleAction struct {
	value *RuleAction
	isSet bool
}

func (v NullableRuleAction) Get() *RuleAction {
	return v.value
}

func (v *NullableRuleAction) Set(val *RuleAction) {
	v.value = val
	v.isSet = true
}

func (v NullableRuleAction) IsSet() bool {
	return v.isSet
}

func (v *NullableRuleAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRuleAction(val *RuleAction) *NullableRuleAction {
	return &NullableRuleAction{value: val, isSet: true}
}

func (v NullableRuleAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRuleAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


