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

// Cluster struct for Cluster
type Cluster struct {
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Cloud *string `json:"cloud,omitempty"`
	Provider *ClusterProviders `json:"provider,omitempty"`
	TotalNodes *int `json:"total_nodes,omitempty"`
	TotalCpus *string `json:"total_cpus,omitempty"`
	TotalMemory *string `json:"total_memory,omitempty"`
	Location *string `json:"location,omitempty"`
	Config *map[string]interface{} `json:"config,omitempty"`
	Tags *map[string]interface{} `json:"tags,omitempty"`
	CreatedBy *string `json:"created_by,omitempty"`
	OwnedBy *string `json:"owned_by,omitempty"`
	Extra *map[string]interface{} `json:"extra,omitempty"`
}

// NewCluster instantiates a new Cluster object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCluster() *Cluster {
	this := Cluster{}
	return &this
}

// NewClusterWithDefaults instantiates a new Cluster object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterWithDefaults() *Cluster {
	this := Cluster{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Cluster) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cluster) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Cluster) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Cluster) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Cluster) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cluster) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Cluster) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Cluster) SetName(v string) {
	o.Name = &v
}

// GetCloud returns the Cloud field value if set, zero value otherwise.
func (o *Cluster) GetCloud() string {
	if o == nil || o.Cloud == nil {
		var ret string
		return ret
	}
	return *o.Cloud
}

// GetCloudOk returns a tuple with the Cloud field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cluster) GetCloudOk() (*string, bool) {
	if o == nil || o.Cloud == nil {
		return nil, false
	}
	return o.Cloud, true
}

// HasCloud returns a boolean if a field has been set.
func (o *Cluster) HasCloud() bool {
	if o != nil && o.Cloud != nil {
		return true
	}

	return false
}

// SetCloud gets a reference to the given string and assigns it to the Cloud field.
func (o *Cluster) SetCloud(v string) {
	o.Cloud = &v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *Cluster) GetProvider() ClusterProviders {
	if o == nil || o.Provider == nil {
		var ret ClusterProviders
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cluster) GetProviderOk() (*ClusterProviders, bool) {
	if o == nil || o.Provider == nil {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *Cluster) HasProvider() bool {
	if o != nil && o.Provider != nil {
		return true
	}

	return false
}

// SetProvider gets a reference to the given ClusterProviders and assigns it to the Provider field.
func (o *Cluster) SetProvider(v ClusterProviders) {
	o.Provider = &v
}

// GetTotalNodes returns the TotalNodes field value if set, zero value otherwise.
func (o *Cluster) GetTotalNodes() int {
	if o == nil || o.TotalNodes == nil {
		var ret int
		return ret
	}
	return *o.TotalNodes
}

// GetTotalNodesOk returns a tuple with the TotalNodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cluster) GetTotalNodesOk() (*int, bool) {
	if o == nil || o.TotalNodes == nil {
		return nil, false
	}
	return o.TotalNodes, true
}

// HasTotalNodes returns a boolean if a field has been set.
func (o *Cluster) HasTotalNodes() bool {
	if o != nil && o.TotalNodes != nil {
		return true
	}

	return false
}

// SetTotalNodes gets a reference to the given int and assigns it to the TotalNodes field.
func (o *Cluster) SetTotalNodes(v int) {
	o.TotalNodes = &v
}

// GetTotalCpus returns the TotalCpus field value if set, zero value otherwise.
func (o *Cluster) GetTotalCpus() string {
	if o == nil || o.TotalCpus == nil {
		var ret string
		return ret
	}
	return *o.TotalCpus
}

// GetTotalCpusOk returns a tuple with the TotalCpus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cluster) GetTotalCpusOk() (*string, bool) {
	if o == nil || o.TotalCpus == nil {
		return nil, false
	}
	return o.TotalCpus, true
}

// HasTotalCpus returns a boolean if a field has been set.
func (o *Cluster) HasTotalCpus() bool {
	if o != nil && o.TotalCpus != nil {
		return true
	}

	return false
}

// SetTotalCpus gets a reference to the given string and assigns it to the TotalCpus field.
func (o *Cluster) SetTotalCpus(v string) {
	o.TotalCpus = &v
}

// GetTotalMemory returns the TotalMemory field value if set, zero value otherwise.
func (o *Cluster) GetTotalMemory() string {
	if o == nil || o.TotalMemory == nil {
		var ret string
		return ret
	}
	return *o.TotalMemory
}

// GetTotalMemoryOk returns a tuple with the TotalMemory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cluster) GetTotalMemoryOk() (*string, bool) {
	if o == nil || o.TotalMemory == nil {
		return nil, false
	}
	return o.TotalMemory, true
}

// HasTotalMemory returns a boolean if a field has been set.
func (o *Cluster) HasTotalMemory() bool {
	if o != nil && o.TotalMemory != nil {
		return true
	}

	return false
}

// SetTotalMemory gets a reference to the given string and assigns it to the TotalMemory field.
func (o *Cluster) SetTotalMemory(v string) {
	o.TotalMemory = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *Cluster) GetLocation() string {
	if o == nil || o.Location == nil {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cluster) GetLocationOk() (*string, bool) {
	if o == nil || o.Location == nil {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *Cluster) HasLocation() bool {
	if o != nil && o.Location != nil {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *Cluster) SetLocation(v string) {
	o.Location = &v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *Cluster) GetConfig() map[string]interface{} {
	if o == nil || o.Config == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cluster) GetConfigOk() (*map[string]interface{}, bool) {
	if o == nil || o.Config == nil {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *Cluster) HasConfig() bool {
	if o != nil && o.Config != nil {
		return true
	}

	return false
}

// SetConfig gets a reference to the given map[string]interface{} and assigns it to the Config field.
func (o *Cluster) SetConfig(v map[string]interface{}) {
	o.Config = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Cluster) GetTags() map[string]interface{} {
	if o == nil || o.Tags == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cluster) GetTagsOk() (*map[string]interface{}, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Cluster) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]interface{} and assigns it to the Tags field.
func (o *Cluster) SetTags(v map[string]interface{}) {
	o.Tags = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *Cluster) GetCreatedBy() string {
	if o == nil || o.CreatedBy == nil {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cluster) GetCreatedByOk() (*string, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *Cluster) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *Cluster) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetOwnedBy returns the OwnedBy field value if set, zero value otherwise.
func (o *Cluster) GetOwnedBy() string {
	if o == nil || o.OwnedBy == nil {
		var ret string
		return ret
	}
	return *o.OwnedBy
}

// GetOwnedByOk returns a tuple with the OwnedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cluster) GetOwnedByOk() (*string, bool) {
	if o == nil || o.OwnedBy == nil {
		return nil, false
	}
	return o.OwnedBy, true
}

// HasOwnedBy returns a boolean if a field has been set.
func (o *Cluster) HasOwnedBy() bool {
	if o != nil && o.OwnedBy != nil {
		return true
	}

	return false
}

// SetOwnedBy gets a reference to the given string and assigns it to the OwnedBy field.
func (o *Cluster) SetOwnedBy(v string) {
	o.OwnedBy = &v
}

// GetExtra returns the Extra field value if set, zero value otherwise.
func (o *Cluster) GetExtra() map[string]interface{} {
	if o == nil || o.Extra == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Extra
}

// GetExtraOk returns a tuple with the Extra field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cluster) GetExtraOk() (*map[string]interface{}, bool) {
	if o == nil || o.Extra == nil {
		return nil, false
	}
	return o.Extra, true
}

// HasExtra returns a boolean if a field has been set.
func (o *Cluster) HasExtra() bool {
	if o != nil && o.Extra != nil {
		return true
	}

	return false
}

// SetExtra gets a reference to the given map[string]interface{} and assigns it to the Extra field.
func (o *Cluster) SetExtra(v map[string]interface{}) {
	o.Extra = &v
}

func (o Cluster) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Cloud != nil {
		toSerialize["cloud"] = o.Cloud
	}
	if o.Provider != nil {
		toSerialize["provider"] = o.Provider
	}
	if o.TotalNodes != nil {
		toSerialize["total_nodes"] = o.TotalNodes
	}
	if o.TotalCpus != nil {
		toSerialize["total_cpus"] = o.TotalCpus
	}
	if o.TotalMemory != nil {
		toSerialize["total_memory"] = o.TotalMemory
	}
	if o.Location != nil {
		toSerialize["location"] = o.Location
	}
	if o.Config != nil {
		toSerialize["config"] = o.Config
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.CreatedBy != nil {
		toSerialize["created_by"] = o.CreatedBy
	}
	if o.OwnedBy != nil {
		toSerialize["owned_by"] = o.OwnedBy
	}
	if o.Extra != nil {
		toSerialize["extra"] = o.Extra
	}
	return json.Marshal(toSerialize)
}

type NullableCluster struct {
	value *Cluster
	isSet bool
}

func (v NullableCluster) Get() *Cluster {
	return v.value
}

func (v *NullableCluster) Set(val *Cluster) {
	v.value = val
	v.isSet = true
}

func (v NullableCluster) IsSet() bool {
	return v.isSet
}

func (v *NullableCluster) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCluster(val *Cluster) *NullableCluster {
	return &NullableCluster{value: val, isSet: true}
}

func (v NullableCluster) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCluster) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


