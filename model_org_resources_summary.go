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

// OrgResourcesSummary struct for OrgResourcesSummary
type OrgResourcesSummary struct {
	Clouds *ResourceCount `json:"clouds,omitempty"`
	Stacks *ResourceCount `json:"stacks,omitempty"`
	Clusters *ResourceCount `json:"clusters,omitempty"`
	Machines *ResourceCount `json:"machines,omitempty"`
	Volumes *ResourceCount `json:"volumes,omitempty"`
	Buckets *ResourceCount `json:"buckets,omitempty"`
	Networks *ResourceCount `json:"networks,omitempty"`
	Zones *ResourceCount `json:"zones,omitempty"`
	Images *ResourceCount `json:"images,omitempty"`
	Keys *ResourceCount `json:"keys,omitempty"`
	Scripts *ResourceCount `json:"scripts,omitempty"`
	Templates *ResourceCount `json:"templates,omitempty"`
	Tunnels *ResourceCount `json:"tunnels,omitempty"`
	Secrets *ResourceCount `json:"secrets,omitempty"`
	Schedules *ResourceCount `json:"schedules,omitempty"`
	Rules *ResourceCount `json:"rules,omitempty"`
	Teams *ResourceCount `json:"teams,omitempty"`
	Members *ResourceCount `json:"members,omitempty"`
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

// GetClouds returns the Clouds field value if set, zero value otherwise.
func (o *OrgResourcesSummary) GetClouds() ResourceCount {
	if o == nil || o.Clouds == nil {
		var ret ResourceCount
		return ret
	}
	return *o.Clouds
}

// GetCloudsOk returns a tuple with the Clouds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgResourcesSummary) GetCloudsOk() (*ResourceCount, bool) {
	if o == nil || o.Clouds == nil {
		return nil, false
	}
	return o.Clouds, true
}

// HasClouds returns a boolean if a field has been set.
func (o *OrgResourcesSummary) HasClouds() bool {
	if o != nil && o.Clouds != nil {
		return true
	}

	return false
}

// SetClouds gets a reference to the given ResourceCount and assigns it to the Clouds field.
func (o *OrgResourcesSummary) SetClouds(v ResourceCount) {
	o.Clouds = &v
}

// GetStacks returns the Stacks field value if set, zero value otherwise.
func (o *OrgResourcesSummary) GetStacks() ResourceCount {
	if o == nil || o.Stacks == nil {
		var ret ResourceCount
		return ret
	}
	return *o.Stacks
}

// GetStacksOk returns a tuple with the Stacks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgResourcesSummary) GetStacksOk() (*ResourceCount, bool) {
	if o == nil || o.Stacks == nil {
		return nil, false
	}
	return o.Stacks, true
}

// HasStacks returns a boolean if a field has been set.
func (o *OrgResourcesSummary) HasStacks() bool {
	if o != nil && o.Stacks != nil {
		return true
	}

	return false
}

// SetStacks gets a reference to the given ResourceCount and assigns it to the Stacks field.
func (o *OrgResourcesSummary) SetStacks(v ResourceCount) {
	o.Stacks = &v
}

// GetClusters returns the Clusters field value if set, zero value otherwise.
func (o *OrgResourcesSummary) GetClusters() ResourceCount {
	if o == nil || o.Clusters == nil {
		var ret ResourceCount
		return ret
	}
	return *o.Clusters
}

// GetClustersOk returns a tuple with the Clusters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgResourcesSummary) GetClustersOk() (*ResourceCount, bool) {
	if o == nil || o.Clusters == nil {
		return nil, false
	}
	return o.Clusters, true
}

// HasClusters returns a boolean if a field has been set.
func (o *OrgResourcesSummary) HasClusters() bool {
	if o != nil && o.Clusters != nil {
		return true
	}

	return false
}

// SetClusters gets a reference to the given ResourceCount and assigns it to the Clusters field.
func (o *OrgResourcesSummary) SetClusters(v ResourceCount) {
	o.Clusters = &v
}

// GetMachines returns the Machines field value if set, zero value otherwise.
func (o *OrgResourcesSummary) GetMachines() ResourceCount {
	if o == nil || o.Machines == nil {
		var ret ResourceCount
		return ret
	}
	return *o.Machines
}

// GetMachinesOk returns a tuple with the Machines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgResourcesSummary) GetMachinesOk() (*ResourceCount, bool) {
	if o == nil || o.Machines == nil {
		return nil, false
	}
	return o.Machines, true
}

// HasMachines returns a boolean if a field has been set.
func (o *OrgResourcesSummary) HasMachines() bool {
	if o != nil && o.Machines != nil {
		return true
	}

	return false
}

// SetMachines gets a reference to the given ResourceCount and assigns it to the Machines field.
func (o *OrgResourcesSummary) SetMachines(v ResourceCount) {
	o.Machines = &v
}

// GetVolumes returns the Volumes field value if set, zero value otherwise.
func (o *OrgResourcesSummary) GetVolumes() ResourceCount {
	if o == nil || o.Volumes == nil {
		var ret ResourceCount
		return ret
	}
	return *o.Volumes
}

// GetVolumesOk returns a tuple with the Volumes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgResourcesSummary) GetVolumesOk() (*ResourceCount, bool) {
	if o == nil || o.Volumes == nil {
		return nil, false
	}
	return o.Volumes, true
}

// HasVolumes returns a boolean if a field has been set.
func (o *OrgResourcesSummary) HasVolumes() bool {
	if o != nil && o.Volumes != nil {
		return true
	}

	return false
}

// SetVolumes gets a reference to the given ResourceCount and assigns it to the Volumes field.
func (o *OrgResourcesSummary) SetVolumes(v ResourceCount) {
	o.Volumes = &v
}

// GetBuckets returns the Buckets field value if set, zero value otherwise.
func (o *OrgResourcesSummary) GetBuckets() ResourceCount {
	if o == nil || o.Buckets == nil {
		var ret ResourceCount
		return ret
	}
	return *o.Buckets
}

// GetBucketsOk returns a tuple with the Buckets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgResourcesSummary) GetBucketsOk() (*ResourceCount, bool) {
	if o == nil || o.Buckets == nil {
		return nil, false
	}
	return o.Buckets, true
}

// HasBuckets returns a boolean if a field has been set.
func (o *OrgResourcesSummary) HasBuckets() bool {
	if o != nil && o.Buckets != nil {
		return true
	}

	return false
}

// SetBuckets gets a reference to the given ResourceCount and assigns it to the Buckets field.
func (o *OrgResourcesSummary) SetBuckets(v ResourceCount) {
	o.Buckets = &v
}

// GetNetworks returns the Networks field value if set, zero value otherwise.
func (o *OrgResourcesSummary) GetNetworks() ResourceCount {
	if o == nil || o.Networks == nil {
		var ret ResourceCount
		return ret
	}
	return *o.Networks
}

// GetNetworksOk returns a tuple with the Networks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgResourcesSummary) GetNetworksOk() (*ResourceCount, bool) {
	if o == nil || o.Networks == nil {
		return nil, false
	}
	return o.Networks, true
}

// HasNetworks returns a boolean if a field has been set.
func (o *OrgResourcesSummary) HasNetworks() bool {
	if o != nil && o.Networks != nil {
		return true
	}

	return false
}

// SetNetworks gets a reference to the given ResourceCount and assigns it to the Networks field.
func (o *OrgResourcesSummary) SetNetworks(v ResourceCount) {
	o.Networks = &v
}

// GetZones returns the Zones field value if set, zero value otherwise.
func (o *OrgResourcesSummary) GetZones() ResourceCount {
	if o == nil || o.Zones == nil {
		var ret ResourceCount
		return ret
	}
	return *o.Zones
}

// GetZonesOk returns a tuple with the Zones field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgResourcesSummary) GetZonesOk() (*ResourceCount, bool) {
	if o == nil || o.Zones == nil {
		return nil, false
	}
	return o.Zones, true
}

// HasZones returns a boolean if a field has been set.
func (o *OrgResourcesSummary) HasZones() bool {
	if o != nil && o.Zones != nil {
		return true
	}

	return false
}

// SetZones gets a reference to the given ResourceCount and assigns it to the Zones field.
func (o *OrgResourcesSummary) SetZones(v ResourceCount) {
	o.Zones = &v
}

// GetImages returns the Images field value if set, zero value otherwise.
func (o *OrgResourcesSummary) GetImages() ResourceCount {
	if o == nil || o.Images == nil {
		var ret ResourceCount
		return ret
	}
	return *o.Images
}

// GetImagesOk returns a tuple with the Images field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgResourcesSummary) GetImagesOk() (*ResourceCount, bool) {
	if o == nil || o.Images == nil {
		return nil, false
	}
	return o.Images, true
}

// HasImages returns a boolean if a field has been set.
func (o *OrgResourcesSummary) HasImages() bool {
	if o != nil && o.Images != nil {
		return true
	}

	return false
}

// SetImages gets a reference to the given ResourceCount and assigns it to the Images field.
func (o *OrgResourcesSummary) SetImages(v ResourceCount) {
	o.Images = &v
}

// GetKeys returns the Keys field value if set, zero value otherwise.
func (o *OrgResourcesSummary) GetKeys() ResourceCount {
	if o == nil || o.Keys == nil {
		var ret ResourceCount
		return ret
	}
	return *o.Keys
}

// GetKeysOk returns a tuple with the Keys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgResourcesSummary) GetKeysOk() (*ResourceCount, bool) {
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

// SetKeys gets a reference to the given ResourceCount and assigns it to the Keys field.
func (o *OrgResourcesSummary) SetKeys(v ResourceCount) {
	o.Keys = &v
}

// GetScripts returns the Scripts field value if set, zero value otherwise.
func (o *OrgResourcesSummary) GetScripts() ResourceCount {
	if o == nil || o.Scripts == nil {
		var ret ResourceCount
		return ret
	}
	return *o.Scripts
}

// GetScriptsOk returns a tuple with the Scripts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgResourcesSummary) GetScriptsOk() (*ResourceCount, bool) {
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

// SetScripts gets a reference to the given ResourceCount and assigns it to the Scripts field.
func (o *OrgResourcesSummary) SetScripts(v ResourceCount) {
	o.Scripts = &v
}

// GetTemplates returns the Templates field value if set, zero value otherwise.
func (o *OrgResourcesSummary) GetTemplates() ResourceCount {
	if o == nil || o.Templates == nil {
		var ret ResourceCount
		return ret
	}
	return *o.Templates
}

// GetTemplatesOk returns a tuple with the Templates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgResourcesSummary) GetTemplatesOk() (*ResourceCount, bool) {
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

// SetTemplates gets a reference to the given ResourceCount and assigns it to the Templates field.
func (o *OrgResourcesSummary) SetTemplates(v ResourceCount) {
	o.Templates = &v
}

// GetTunnels returns the Tunnels field value if set, zero value otherwise.
func (o *OrgResourcesSummary) GetTunnels() ResourceCount {
	if o == nil || o.Tunnels == nil {
		var ret ResourceCount
		return ret
	}
	return *o.Tunnels
}

// GetTunnelsOk returns a tuple with the Tunnels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgResourcesSummary) GetTunnelsOk() (*ResourceCount, bool) {
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

// SetTunnels gets a reference to the given ResourceCount and assigns it to the Tunnels field.
func (o *OrgResourcesSummary) SetTunnels(v ResourceCount) {
	o.Tunnels = &v
}

// GetSecrets returns the Secrets field value if set, zero value otherwise.
func (o *OrgResourcesSummary) GetSecrets() ResourceCount {
	if o == nil || o.Secrets == nil {
		var ret ResourceCount
		return ret
	}
	return *o.Secrets
}

// GetSecretsOk returns a tuple with the Secrets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgResourcesSummary) GetSecretsOk() (*ResourceCount, bool) {
	if o == nil || o.Secrets == nil {
		return nil, false
	}
	return o.Secrets, true
}

// HasSecrets returns a boolean if a field has been set.
func (o *OrgResourcesSummary) HasSecrets() bool {
	if o != nil && o.Secrets != nil {
		return true
	}

	return false
}

// SetSecrets gets a reference to the given ResourceCount and assigns it to the Secrets field.
func (o *OrgResourcesSummary) SetSecrets(v ResourceCount) {
	o.Secrets = &v
}

// GetSchedules returns the Schedules field value if set, zero value otherwise.
func (o *OrgResourcesSummary) GetSchedules() ResourceCount {
	if o == nil || o.Schedules == nil {
		var ret ResourceCount
		return ret
	}
	return *o.Schedules
}

// GetSchedulesOk returns a tuple with the Schedules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgResourcesSummary) GetSchedulesOk() (*ResourceCount, bool) {
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

// SetSchedules gets a reference to the given ResourceCount and assigns it to the Schedules field.
func (o *OrgResourcesSummary) SetSchedules(v ResourceCount) {
	o.Schedules = &v
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *OrgResourcesSummary) GetRules() ResourceCount {
	if o == nil || o.Rules == nil {
		var ret ResourceCount
		return ret
	}
	return *o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgResourcesSummary) GetRulesOk() (*ResourceCount, bool) {
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

// SetRules gets a reference to the given ResourceCount and assigns it to the Rules field.
func (o *OrgResourcesSummary) SetRules(v ResourceCount) {
	o.Rules = &v
}

// GetTeams returns the Teams field value if set, zero value otherwise.
func (o *OrgResourcesSummary) GetTeams() ResourceCount {
	if o == nil || o.Teams == nil {
		var ret ResourceCount
		return ret
	}
	return *o.Teams
}

// GetTeamsOk returns a tuple with the Teams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgResourcesSummary) GetTeamsOk() (*ResourceCount, bool) {
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

// SetTeams gets a reference to the given ResourceCount and assigns it to the Teams field.
func (o *OrgResourcesSummary) SetTeams(v ResourceCount) {
	o.Teams = &v
}

// GetMembers returns the Members field value if set, zero value otherwise.
func (o *OrgResourcesSummary) GetMembers() ResourceCount {
	if o == nil || o.Members == nil {
		var ret ResourceCount
		return ret
	}
	return *o.Members
}

// GetMembersOk returns a tuple with the Members field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgResourcesSummary) GetMembersOk() (*ResourceCount, bool) {
	if o == nil || o.Members == nil {
		return nil, false
	}
	return o.Members, true
}

// HasMembers returns a boolean if a field has been set.
func (o *OrgResourcesSummary) HasMembers() bool {
	if o != nil && o.Members != nil {
		return true
	}

	return false
}

// SetMembers gets a reference to the given ResourceCount and assigns it to the Members field.
func (o *OrgResourcesSummary) SetMembers(v ResourceCount) {
	o.Members = &v
}

func (o OrgResourcesSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Clouds != nil {
		toSerialize["clouds"] = o.Clouds
	}
	if o.Stacks != nil {
		toSerialize["stacks"] = o.Stacks
	}
	if o.Clusters != nil {
		toSerialize["clusters"] = o.Clusters
	}
	if o.Machines != nil {
		toSerialize["machines"] = o.Machines
	}
	if o.Volumes != nil {
		toSerialize["volumes"] = o.Volumes
	}
	if o.Buckets != nil {
		toSerialize["buckets"] = o.Buckets
	}
	if o.Networks != nil {
		toSerialize["networks"] = o.Networks
	}
	if o.Zones != nil {
		toSerialize["zones"] = o.Zones
	}
	if o.Images != nil {
		toSerialize["images"] = o.Images
	}
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
	if o.Secrets != nil {
		toSerialize["secrets"] = o.Secrets
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
	if o.Members != nil {
		toSerialize["members"] = o.Members
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


