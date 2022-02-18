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

// OrgResourcesSummary struct for OrgResourcesSummary
type OrgResourcesSummary struct {
	Keys *OrgResourcesSummaryKeys `json:"keys,omitempty"`
	Scripts *OrgResourcesSummaryKeys `json:"scripts,omitempty"`
	Templates *OrgResourcesSummaryKeys `json:"templates,omitempty"`
	Tunnels *OrgResourcesSummaryKeys `json:"tunnels,omitempty"`
	Schedules *OrgResourcesSummaryKeys `json:"schedules,omitempty"`
	Rules *OrgResourcesSummaryKeys `json:"rules,omitempty"`
	Teams *OrgResourcesSummaryKeys `json:"teams,omitempty"`
}

// NewOrgResourcesSummary instantiates a new OrgResourcesSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrgResourcesSummary() *OrgResourcesSummary {
	this := OrgResourcesSummary{}
	return &this
}

// NewOrgResourcesSummaryWithDefaults instantiates a new OrgResourcesSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrgResourcesSummaryWithDefaults() *OrgResourcesSummary {
	this := OrgResourcesSummary{}
	return &this
}

// GetKeys returns the Keys field value if set, zero value otherwise.
func (o *OrgResourcesSummary) GetKeys() OrgResourcesSummaryKeys {
	if o == nil || o.Keys == nil {
		var ret OrgResourcesSummaryKeys
		return ret
	}
	return *o.Keys
}

// GetKeysOk returns a tuple with the Keys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgResourcesSummary) GetKeysOk() (*OrgResourcesSummaryKeys, bool) {
	if o == nil || o.Keys == nil {
		return nil, false
	}
	return o.Keys, true
}

// HasKeys returns a boolean if a field has been set.
func (o *OrgResourcesSummary) HasKeys() bool {
	if o != nil && o.Keys != nil {
		return true
	}

	return false
}

// SetKeys gets a reference to the given OrgResourcesSummaryKeys and assigns it to the Keys field.
func (o *OrgResourcesSummary) SetKeys(v OrgResourcesSummaryKeys) {
	o.Keys = &v
}

// GetScripts returns the Scripts field value if set, zero value otherwise.
func (o *OrgResourcesSummary) GetScripts() OrgResourcesSummaryKeys {
	if o == nil || o.Scripts == nil {
		var ret OrgResourcesSummaryKeys
		return ret
	}
	return *o.Scripts
}

// GetScriptsOk returns a tuple with the Scripts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgResourcesSummary) GetScriptsOk() (*OrgResourcesSummaryKeys, bool) {
	if o == nil || o.Scripts == nil {
		return nil, false
	}
	return o.Scripts, true
}

// HasScripts returns a boolean if a field has been set.
func (o *OrgResourcesSummary) HasScripts() bool {
	if o != nil && o.Scripts != nil {
		return true
	}

	return false
}

// SetScripts gets a reference to the given OrgResourcesSummaryKeys and assigns it to the Scripts field.
func (o *OrgResourcesSummary) SetScripts(v OrgResourcesSummaryKeys) {
	o.Scripts = &v
}

// GetTemplates returns the Templates field value if set, zero value otherwise.
func (o *OrgResourcesSummary) GetTemplates() OrgResourcesSummaryKeys {
	if o == nil || o.Templates == nil {
		var ret OrgResourcesSummaryKeys
		return ret
	}
	return *o.Templates
}

// GetTemplatesOk returns a tuple with the Templates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgResourcesSummary) GetTemplatesOk() (*OrgResourcesSummaryKeys, bool) {
	if o == nil || o.Templates == nil {
		return nil, false
	}
	return o.Templates, true
}

// HasTemplates returns a boolean if a field has been set.
func (o *OrgResourcesSummary) HasTemplates() bool {
	if o != nil && o.Templates != nil {
		return true
	}

	return false
}

// SetTemplates gets a reference to the given OrgResourcesSummaryKeys and assigns it to the Templates field.
func (o *OrgResourcesSummary) SetTemplates(v OrgResourcesSummaryKeys) {
	o.Templates = &v
}

// GetTunnels returns the Tunnels field value if set, zero value otherwise.
func (o *OrgResourcesSummary) GetTunnels() OrgResourcesSummaryKeys {
	if o == nil || o.Tunnels == nil {
		var ret OrgResourcesSummaryKeys
		return ret
	}
	return *o.Tunnels
}

// GetTunnelsOk returns a tuple with the Tunnels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgResourcesSummary) GetTunnelsOk() (*OrgResourcesSummaryKeys, bool) {
	if o == nil || o.Tunnels == nil {
		return nil, false
	}
	return o.Tunnels, true
}

// HasTunnels returns a boolean if a field has been set.
func (o *OrgResourcesSummary) HasTunnels() bool {
	if o != nil && o.Tunnels != nil {
		return true
	}

	return false
}

// SetTunnels gets a reference to the given OrgResourcesSummaryKeys and assigns it to the Tunnels field.
func (o *OrgResourcesSummary) SetTunnels(v OrgResourcesSummaryKeys) {
	o.Tunnels = &v
}

// GetSchedules returns the Schedules field value if set, zero value otherwise.
func (o *OrgResourcesSummary) GetSchedules() OrgResourcesSummaryKeys {
	if o == nil || o.Schedules == nil {
		var ret OrgResourcesSummaryKeys
		return ret
	}
	return *o.Schedules
}

// GetSchedulesOk returns a tuple with the Schedules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgResourcesSummary) GetSchedulesOk() (*OrgResourcesSummaryKeys, bool) {
	if o == nil || o.Schedules == nil {
		return nil, false
	}
	return o.Schedules, true
}

// HasSchedules returns a boolean if a field has been set.
func (o *OrgResourcesSummary) HasSchedules() bool {
	if o != nil && o.Schedules != nil {
		return true
	}

	return false
}

// SetSchedules gets a reference to the given OrgResourcesSummaryKeys and assigns it to the Schedules field.
func (o *OrgResourcesSummary) SetSchedules(v OrgResourcesSummaryKeys) {
	o.Schedules = &v
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *OrgResourcesSummary) GetRules() OrgResourcesSummaryKeys {
	if o == nil || o.Rules == nil {
		var ret OrgResourcesSummaryKeys
		return ret
	}
	return *o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgResourcesSummary) GetRulesOk() (*OrgResourcesSummaryKeys, bool) {
	if o == nil || o.Rules == nil {
		return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *OrgResourcesSummary) HasRules() bool {
	if o != nil && o.Rules != nil {
		return true
	}

	return false
}

// SetRules gets a reference to the given OrgResourcesSummaryKeys and assigns it to the Rules field.
func (o *OrgResourcesSummary) SetRules(v OrgResourcesSummaryKeys) {
	o.Rules = &v
}

// GetTeams returns the Teams field value if set, zero value otherwise.
func (o *OrgResourcesSummary) GetTeams() OrgResourcesSummaryKeys {
	if o == nil || o.Teams == nil {
		var ret OrgResourcesSummaryKeys
		return ret
	}
	return *o.Teams
}

// GetTeamsOk returns a tuple with the Teams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgResourcesSummary) GetTeamsOk() (*OrgResourcesSummaryKeys, bool) {
	if o == nil || o.Teams == nil {
		return nil, false
	}
	return o.Teams, true
}

// HasTeams returns a boolean if a field has been set.
func (o *OrgResourcesSummary) HasTeams() bool {
	if o != nil && o.Teams != nil {
		return true
	}

	return false
}

// SetTeams gets a reference to the given OrgResourcesSummaryKeys and assigns it to the Teams field.
func (o *OrgResourcesSummary) SetTeams(v OrgResourcesSummaryKeys) {
	o.Teams = &v
}

func (o OrgResourcesSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Keys != nil {
		toSerialize["keys"] = o.Keys
	}
	if o.Scripts != nil {
		toSerialize["scripts"] = o.Scripts
	}
	if o.Templates != nil {
		toSerialize["templates"] = o.Templates
	}
	if o.Tunnels != nil {
		toSerialize["tunnels"] = o.Tunnels
	}
	if o.Schedules != nil {
		toSerialize["schedules"] = o.Schedules
	}
	if o.Rules != nil {
		toSerialize["rules"] = o.Rules
	}
	if o.Teams != nil {
		toSerialize["teams"] = o.Teams
	}
	return json.Marshal(toSerialize)
}

type NullableOrgResourcesSummary struct {
	value *OrgResourcesSummary
	isSet bool
}

func (v NullableOrgResourcesSummary) Get() *OrgResourcesSummary {
	return v.value
}

func (v *NullableOrgResourcesSummary) Set(val *OrgResourcesSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableOrgResourcesSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableOrgResourcesSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrgResourcesSummary(val *OrgResourcesSummary) *NullableOrgResourcesSummary {
	return &NullableOrgResourcesSummary{value: val, isSet: true}
}

func (v NullableOrgResourcesSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrgResourcesSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

