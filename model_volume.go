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

// Volume struct for Volume
type Volume struct {
	Cloud *string `json:"cloud,omitempty"`
	CreatedBy *string `json:"created_by,omitempty"`
	Default *bool `json:"default,omitempty"`
	ExternalId *string `json:"external_id,omitempty"`
	Id *string `json:"id,omitempty"`
	Location *string `json:"location,omitempty"`
	Name *string `json:"name,omitempty"`
	OwnedBy *string `json:"owned_by,omitempty"`
	Size *int32 `json:"size,omitempty"`
	Tags *map[string]interface{} `json:"tags,omitempty"`
}

// NewVolume instantiates a new Volume object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVolume() *Volume {
	this := Volume{}
	return &this
}

// NewVolumeWithDefaults instantiates a new Volume object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVolumeWithDefaults() *Volume {
	this := Volume{}
	return &this
}

// GetCloud returns the Cloud field value if set, zero value otherwise.
func (o *Volume) GetCloud() string {
	if o == nil || o.Cloud == nil {
		var ret string
		return ret
	}
	return *o.Cloud
}

// GetCloudOk returns a tuple with the Cloud field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Volume) GetCloudOk() (*string, bool) {
	if o == nil || o.Cloud == nil {
		return nil, false
	}
	return o.Cloud, true
}

// HasCloud returns a boolean if a field has been set.
func (o *Volume) HasCloud() bool {
	if o != nil && o.Cloud != nil {
		return true
	}

	return false
}

// SetCloud gets a reference to the given string and assigns it to the Cloud field.
func (o *Volume) SetCloud(v string) {
	o.Cloud = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *Volume) GetCreatedBy() string {
	if o == nil || o.CreatedBy == nil {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Volume) GetCreatedByOk() (*string, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *Volume) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *Volume) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *Volume) GetDefault() bool {
	if o == nil || o.Default == nil {
		var ret bool
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Volume) GetDefaultOk() (*bool, bool) {
	if o == nil || o.Default == nil {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *Volume) HasDefault() bool {
	if o != nil && o.Default != nil {
		return true
	}

	return false
}

// SetDefault gets a reference to the given bool and assigns it to the Default field.
func (o *Volume) SetDefault(v bool) {
	o.Default = &v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *Volume) GetExternalId() string {
	if o == nil || o.ExternalId == nil {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Volume) GetExternalIdOk() (*string, bool) {
	if o == nil || o.ExternalId == nil {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *Volume) HasExternalId() bool {
	if o != nil && o.ExternalId != nil {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *Volume) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Volume) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Volume) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Volume) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Volume) SetId(v string) {
	o.Id = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *Volume) GetLocation() string {
	if o == nil || o.Location == nil {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Volume) GetLocationOk() (*string, bool) {
	if o == nil || o.Location == nil {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *Volume) HasLocation() bool {
	if o != nil && o.Location != nil {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *Volume) SetLocation(v string) {
	o.Location = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Volume) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Volume) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Volume) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Volume) SetName(v string) {
	o.Name = &v
}

// GetOwnedBy returns the OwnedBy field value if set, zero value otherwise.
func (o *Volume) GetOwnedBy() string {
	if o == nil || o.OwnedBy == nil {
		var ret string
		return ret
	}
	return *o.OwnedBy
}

// GetOwnedByOk returns a tuple with the OwnedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Volume) GetOwnedByOk() (*string, bool) {
	if o == nil || o.OwnedBy == nil {
		return nil, false
	}
	return o.OwnedBy, true
}

// HasOwnedBy returns a boolean if a field has been set.
func (o *Volume) HasOwnedBy() bool {
	if o != nil && o.OwnedBy != nil {
		return true
	}

	return false
}

// SetOwnedBy gets a reference to the given string and assigns it to the OwnedBy field.
func (o *Volume) SetOwnedBy(v string) {
	o.OwnedBy = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *Volume) GetSize() int32 {
	if o == nil || o.Size == nil {
		var ret int32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Volume) GetSizeOk() (*int32, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *Volume) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given int32 and assigns it to the Size field.
func (o *Volume) SetSize(v int32) {
	o.Size = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Volume) GetTags() map[string]interface{} {
	if o == nil || o.Tags == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Volume) GetTagsOk() (*map[string]interface{}, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Volume) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]interface{} and assigns it to the Tags field.
func (o *Volume) SetTags(v map[string]interface{}) {
	o.Tags = &v
}

func (o Volume) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Cloud != nil {
		toSerialize["cloud"] = o.Cloud
	}
	if o.CreatedBy != nil {
		toSerialize["created_by"] = o.CreatedBy
	}
	if o.Default != nil {
		toSerialize["default"] = o.Default
	}
	if o.ExternalId != nil {
		toSerialize["external_id"] = o.ExternalId
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Location != nil {
		toSerialize["location"] = o.Location
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.OwnedBy != nil {
		toSerialize["owned_by"] = o.OwnedBy
	}
	if o.Size != nil {
		toSerialize["size"] = o.Size
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	return json.Marshal(toSerialize)
}

type NullableVolume struct {
	value *Volume
	isSet bool
}

func (v NullableVolume) Get() *Volume {
	return v.value
}

func (v *NullableVolume) Set(val *Volume) {
	v.value = val
	v.isSet = true
}

func (v NullableVolume) IsSet() bool {
	return v.isSet
}

func (v *NullableVolume) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVolume(val *Volume) *NullableVolume {
	return &NullableVolume{value: val, isSet: true}
}

func (v NullableVolume) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVolume) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


