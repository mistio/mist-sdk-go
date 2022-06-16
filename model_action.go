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

// Action struct for Action
type Action struct {
	// the action's type: notification, resource_action, run_script, webhook 
	ActionType string `json:"action_type"`
	// a list of user to be notified, denoted by their UUIDs 
	Users []string `json:"users,omitempty"`
	// a list of teams, denoted by their UUIDs, whose users will be notified 
	Teams []string `json:"teams,omitempty"`
	// a list of e-mails to send a notification to 
	Emails []string `json:"emails,omitempty"`
	// the query string parameters of the HTTP request
	Params string `json:"params"`
	ScriptType string `json:"script_type"`
	// Command that is about to run
	Command string `json:"command"`
	// Name or ID of the script to run
	Script string `json:"script"`
	// the HTTP method to be executed by the webhook
	Method string `json:"method"`
	// the URL of the endpoint that is called by the webhook
	Url string `json:"url"`
	// the body of the HTTP request
	Data *string `json:"data,omitempty"`
	// the JSON body of the HTTP request
	Json *string `json:"json,omitempty"`
	// the HTTP headers of the request
	Headers *string `json:"headers,omitempty"`
}

// NewAction instantiates a new Action object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAction(actionType string, params string, scriptType string, command string, script string, method string, url string) *Action {
	this := Action{}
	this.ActionType = actionType
	this.Params = params
	this.ScriptType = scriptType
	this.Command = command
	this.Script = script
	this.Method = method
	this.Url = url
	return &this
}

// NewActionWithDefaults instantiates a new Action object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActionWithDefaults() *Action {
	this := Action{}
	return &this
}

// GetActionType returns the ActionType field value
func (o *Action) GetActionType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ActionType
}

// GetActionTypeOk returns a tuple with the ActionType field value
// and a boolean to check if the value has been set.
func (o *Action) GetActionTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActionType, true
}

// SetActionType sets field value
func (o *Action) SetActionType(v string) {
	o.ActionType = v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *Action) GetUsers() []string {
	if o == nil || o.Users == nil {
		var ret []string
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Action) GetUsersOk() ([]string, bool) {
	if o == nil || o.Users == nil {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *Action) HasUsers() bool {
	if o != nil && o.Users != nil {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []string and assigns it to the Users field.
func (o *Action) SetUsers(v []string) {
	o.Users = v
}

// GetTeams returns the Teams field value if set, zero value otherwise.
func (o *Action) GetTeams() []string {
	if o == nil || o.Teams == nil {
		var ret []string
		return ret
	}
	return o.Teams
}

// GetTeamsOk returns a tuple with the Teams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Action) GetTeamsOk() ([]string, bool) {
	if o == nil || o.Teams == nil {
		return nil, false
	}
	return o.Teams, true
}

// HasTeams returns a boolean if a field has been set.
func (o *Action) HasTeams() bool {
	if o != nil && o.Teams != nil {
		return true
	}

	return false
}

// SetTeams gets a reference to the given []string and assigns it to the Teams field.
func (o *Action) SetTeams(v []string) {
	o.Teams = v
}

// GetEmails returns the Emails field value if set, zero value otherwise.
func (o *Action) GetEmails() []string {
	if o == nil || o.Emails == nil {
		var ret []string
		return ret
	}
	return o.Emails
}

// GetEmailsOk returns a tuple with the Emails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Action) GetEmailsOk() ([]string, bool) {
	if o == nil || o.Emails == nil {
		return nil, false
	}
	return o.Emails, true
}

// HasEmails returns a boolean if a field has been set.
func (o *Action) HasEmails() bool {
	if o != nil && o.Emails != nil {
		return true
	}

	return false
}

// SetEmails gets a reference to the given []string and assigns it to the Emails field.
func (o *Action) SetEmails(v []string) {
	o.Emails = v
}

// GetParams returns the Params field value
func (o *Action) GetParams() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Params
}

// GetParamsOk returns a tuple with the Params field value
// and a boolean to check if the value has been set.
func (o *Action) GetParamsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Params, true
}

// SetParams sets field value
func (o *Action) SetParams(v string) {
	o.Params = v
}

// GetScriptType returns the ScriptType field value
func (o *Action) GetScriptType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ScriptType
}

// GetScriptTypeOk returns a tuple with the ScriptType field value
// and a boolean to check if the value has been set.
func (o *Action) GetScriptTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScriptType, true
}

// SetScriptType sets field value
func (o *Action) SetScriptType(v string) {
	o.ScriptType = v
}

// GetCommand returns the Command field value
func (o *Action) GetCommand() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Command
}

// GetCommandOk returns a tuple with the Command field value
// and a boolean to check if the value has been set.
func (o *Action) GetCommandOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Command, true
}

// SetCommand sets field value
func (o *Action) SetCommand(v string) {
	o.Command = v
}

// GetScript returns the Script field value
func (o *Action) GetScript() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Script
}

// GetScriptOk returns a tuple with the Script field value
// and a boolean to check if the value has been set.
func (o *Action) GetScriptOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Script, true
}

// SetScript sets field value
func (o *Action) SetScript(v string) {
	o.Script = v
}

// GetMethod returns the Method field value
func (o *Action) GetMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Method
}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
func (o *Action) GetMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Method, true
}

// SetMethod sets field value
func (o *Action) SetMethod(v string) {
	o.Method = v
}

// GetUrl returns the Url field value
func (o *Action) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *Action) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *Action) SetUrl(v string) {
	o.Url = v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *Action) GetData() string {
	if o == nil || o.Data == nil {
		var ret string
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Action) GetDataOk() (*string, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *Action) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given string and assigns it to the Data field.
func (o *Action) SetData(v string) {
	o.Data = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *Action) GetJson() string {
	if o == nil || o.Json == nil {
		var ret string
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Action) GetJsonOk() (*string, bool) {
	if o == nil || o.Json == nil {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *Action) HasJson() bool {
	if o != nil && o.Json != nil {
		return true
	}

	return false
}

// SetJson gets a reference to the given string and assigns it to the Json field.
func (o *Action) SetJson(v string) {
	o.Json = &v
}

// GetHeaders returns the Headers field value if set, zero value otherwise.
func (o *Action) GetHeaders() string {
	if o == nil || o.Headers == nil {
		var ret string
		return ret
	}
	return *o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Action) GetHeadersOk() (*string, bool) {
	if o == nil || o.Headers == nil {
		return nil, false
	}
	return o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *Action) HasHeaders() bool {
	if o != nil && o.Headers != nil {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given string and assigns it to the Headers field.
func (o *Action) SetHeaders(v string) {
	o.Headers = &v
}

func (o Action) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["action_type"] = o.ActionType
	}
	if o.Users != nil {
		toSerialize["users"] = o.Users
	}
	if o.Teams != nil {
		toSerialize["teams"] = o.Teams
	}
	if o.Emails != nil {
		toSerialize["emails"] = o.Emails
	}
	if true {
		toSerialize["params"] = o.Params
	}
	if true {
		toSerialize["script_type"] = o.ScriptType
	}
	if true {
		toSerialize["command"] = o.Command
	}
	if true {
		toSerialize["script"] = o.Script
	}
	if true {
		toSerialize["method"] = o.Method
	}
	if true {
		toSerialize["url"] = o.Url
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Json != nil {
		toSerialize["json"] = o.Json
	}
	if o.Headers != nil {
		toSerialize["headers"] = o.Headers
	}
	return json.Marshal(toSerialize)
}

type NullableAction struct {
	value *Action
	isSet bool
}

func (v NullableAction) Get() *Action {
	return v.value
}

func (v *NullableAction) Set(val *Action) {
	v.value = val
	v.isSet = true
}

func (v NullableAction) IsSet() bool {
	return v.isSet
}

func (v *NullableAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAction(val *Action) *NullableAction {
	return &NullableAction{value: val, isSet: true}
}

func (v NullableAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


