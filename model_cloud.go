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

// Cloud struct for Cloud
type Cloud struct {
	Id *string `json:"id,omitempty"`
	Title *string `json:"title,omitempty"`
	Provider *SupportedProviders `json:"provider,omitempty"`
	Config *map[string]interface{} `json:"config,omitempty"`
	Features *CloudFeatures `json:"features,omitempty"`
	Tags *map[string]interface{} `json:"tags,omitempty"`
	CreatedBy *string `json:"created_by,omitempty"`
	OwnedBy *string `json:"owned_by,omitempty"`
}

// NewCloud instantiates a new Cloud object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloud() *Cloud {
	this := Cloud{}
	return &this
}

// NewCloudWithDefaults instantiates a new Cloud object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudWithDefaults() *Cloud {
	this := Cloud{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Cloud) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cloud) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Cloud) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Cloud) SetId(v string) {
	o.Id = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *Cloud) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cloud) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *Cloud) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *Cloud) SetTitle(v string) {
	o.Title = &v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *Cloud) GetProvider() SupportedProviders {
	if o == nil || o.Provider == nil {
		var ret SupportedProviders
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cloud) GetProviderOk() (*SupportedProviders, bool) {
	if o == nil || o.Provider == nil {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *Cloud) HasProvider() bool {
	if o != nil && o.Provider != nil {
		return true
	}

	return false
}

// SetProvider gets a reference to the given SupportedProviders and assigns it to the Provider field.
func (o *Cloud) SetProvider(v SupportedProviders) {
	o.Provider = &v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *Cloud) GetConfig() map[string]interface{} {
	if o == nil || o.Config == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cloud) GetConfigOk() (*map[string]interface{}, bool) {
	if o == nil || o.Config == nil {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *Cloud) HasConfig() bool {
	if o != nil && o.Config != nil {
		return true
	}

	return false
}

// SetConfig gets a reference to the given map[string]interface{} and assigns it to the Config field.
func (o *Cloud) SetConfig(v map[string]interface{}) {
	o.Config = &v
}

// GetFeatures returns the Features field value if set, zero value otherwise.
func (o *Cloud) GetFeatures() CloudFeatures {
	if o == nil || o.Features == nil {
		var ret CloudFeatures
		return ret
	}
	return *o.Features
}

// GetFeaturesOk returns a tuple with the Features field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cloud) GetFeaturesOk() (*CloudFeatures, bool) {
	if o == nil || o.Features == nil {
		return nil, false
	}
	return o.Features, true
}

// HasFeatures returns a boolean if a field has been set.
func (o *Cloud) HasFeatures() bool {
	if o != nil && o.Features != nil {
		return true
	}

	return false
}

// SetFeatures gets a reference to the given CloudFeatures and assigns it to the Features field.
func (o *Cloud) SetFeatures(v CloudFeatures) {
	o.Features = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Cloud) GetTags() map[string]interface{} {
	if o == nil || o.Tags == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cloud) GetTagsOk() (*map[string]interface{}, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Cloud) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]interface{} and assigns it to the Tags field.
func (o *Cloud) SetTags(v map[string]interface{}) {
	o.Tags = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *Cloud) GetCreatedBy() string {
	if o == nil || o.CreatedBy == nil {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cloud) GetCreatedByOk() (*string, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *Cloud) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *Cloud) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetOwnedBy returns the OwnedBy field value if set, zero value otherwise.
func (o *Cloud) GetOwnedBy() string {
	if o == nil || o.OwnedBy == nil {
		var ret string
		return ret
	}
	return *o.OwnedBy
}

// GetOwnedByOk returns a tuple with the OwnedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cloud) GetOwnedByOk() (*string, bool) {
	if o == nil || o.OwnedBy == nil {
		return nil, false
	}
	return o.OwnedBy, true
}

// HasOwnedBy returns a boolean if a field has been set.
func (o *Cloud) HasOwnedBy() bool {
	if o != nil && o.OwnedBy != nil {
		return true
	}

	return false
}

// SetOwnedBy gets a reference to the given string and assigns it to the OwnedBy field.
func (o *Cloud) SetOwnedBy(v string) {
	o.OwnedBy = &v
}

func (o Cloud) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.Provider != nil {
		toSerialize["provider"] = o.Provider
	}
	if o.Config != nil {
		toSerialize["config"] = o.Config
	}
	if o.Features != nil {
		toSerialize["features"] = o.Features
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
	return json.Marshal(toSerialize)
}

type NullableCloud struct {
	value *Cloud
	isSet bool
}

func (v NullableCloud) Get() *Cloud {
	return v.value
}

func (v *NullableCloud) Set(val *Cloud) {
	v.value = val
	v.isSet = true
}

func (v NullableCloud) IsSet() bool {
	return v.isSet
}

func (v *NullableCloud) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloud(val *Cloud) *NullableCloud {
	return &NullableCloud{value: val, isSet: true}
}

func (v NullableCloud) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloud) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


