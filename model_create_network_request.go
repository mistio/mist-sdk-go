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

// CreateNetworkRequest struct for CreateNetworkRequest
type CreateNetworkRequest struct {
	// Specify network name
	Name string `json:"name"`
	// Specify cloud to provision on
	Cloud string `json:"cloud"`
	// Assign tags to provisioned network
	Tags *map[string]interface{} `json:"tags,omitempty"`
	// Configure additional parameters
	Extra *map[string]interface{} `json:"extra,omitempty"`
	Template *map[string]interface{} `json:"template,omitempty"`
	// Return provisioning plan and exit without executing it
	Dry *bool `json:"dry,omitempty"`
	// Save provisioning plan as template
	Save *bool `json:"save,omitempty"`
}

// NewCreateNetworkRequest instantiates a new CreateNetworkRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworkRequest(name string, cloud string, ) *CreateNetworkRequest {
	this := CreateNetworkRequest{}
	this.Name = name
	this.Cloud = cloud
	return &this
}

// NewCreateNetworkRequestWithDefaults instantiates a new CreateNetworkRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworkRequestWithDefaults() *CreateNetworkRequest {
	this := CreateNetworkRequest{}
	return &this
}

// GetName returns the Name field value
func (o *CreateNetworkRequest) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateNetworkRequest) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateNetworkRequest) SetName(v string) {
	o.Name = v
}

// GetCloud returns the Cloud field value
func (o *CreateNetworkRequest) GetCloud() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Cloud
}

// GetCloudOk returns a tuple with the Cloud field value
// and a boolean to check if the value has been set.
func (o *CreateNetworkRequest) GetCloudOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Cloud, true
}

// SetCloud sets field value
func (o *CreateNetworkRequest) SetCloud(v string) {
	o.Cloud = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *CreateNetworkRequest) GetTags() map[string]interface{} {
	if o == nil || o.Tags == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkRequest) GetTagsOk() (*map[string]interface{}, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *CreateNetworkRequest) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]interface{} and assigns it to the Tags field.
func (o *CreateNetworkRequest) SetTags(v map[string]interface{}) {
	o.Tags = &v
}

// GetExtra returns the Extra field value if set, zero value otherwise.
func (o *CreateNetworkRequest) GetExtra() map[string]interface{} {
	if o == nil || o.Extra == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Extra
}

// GetExtraOk returns a tuple with the Extra field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkRequest) GetExtraOk() (*map[string]interface{}, bool) {
	if o == nil || o.Extra == nil {
		return nil, false
	}
	return o.Extra, true
}

// HasExtra returns a boolean if a field has been set.
func (o *CreateNetworkRequest) HasExtra() bool {
	if o != nil && o.Extra != nil {
		return true
	}

	return false
}

// SetExtra gets a reference to the given map[string]interface{} and assigns it to the Extra field.
func (o *CreateNetworkRequest) SetExtra(v map[string]interface{}) {
	o.Extra = &v
}

// GetTemplate returns the Template field value if set, zero value otherwise.
func (o *CreateNetworkRequest) GetTemplate() map[string]interface{} {
	if o == nil || o.Template == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Template
}

// GetTemplateOk returns a tuple with the Template field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkRequest) GetTemplateOk() (*map[string]interface{}, bool) {
	if o == nil || o.Template == nil {
		return nil, false
	}
	return o.Template, true
}

// HasTemplate returns a boolean if a field has been set.
func (o *CreateNetworkRequest) HasTemplate() bool {
	if o != nil && o.Template != nil {
		return true
	}

	return false
}

// SetTemplate gets a reference to the given map[string]interface{} and assigns it to the Template field.
func (o *CreateNetworkRequest) SetTemplate(v map[string]interface{}) {
	o.Template = &v
}

// GetDry returns the Dry field value if set, zero value otherwise.
func (o *CreateNetworkRequest) GetDry() bool {
	if o == nil || o.Dry == nil {
		var ret bool
		return ret
	}
	return *o.Dry
}

// GetDryOk returns a tuple with the Dry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkRequest) GetDryOk() (*bool, bool) {
	if o == nil || o.Dry == nil {
		return nil, false
	}
	return o.Dry, true
}

// HasDry returns a boolean if a field has been set.
func (o *CreateNetworkRequest) HasDry() bool {
	if o != nil && o.Dry != nil {
		return true
	}

	return false
}

// SetDry gets a reference to the given bool and assigns it to the Dry field.
func (o *CreateNetworkRequest) SetDry(v bool) {
	o.Dry = &v
}

// GetSave returns the Save field value if set, zero value otherwise.
func (o *CreateNetworkRequest) GetSave() bool {
	if o == nil || o.Save == nil {
		var ret bool
		return ret
	}
	return *o.Save
}

// GetSaveOk returns a tuple with the Save field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkRequest) GetSaveOk() (*bool, bool) {
	if o == nil || o.Save == nil {
		return nil, false
	}
	return o.Save, true
}

// HasSave returns a boolean if a field has been set.
func (o *CreateNetworkRequest) HasSave() bool {
	if o != nil && o.Save != nil {
		return true
	}

	return false
}

// SetSave gets a reference to the given bool and assigns it to the Save field.
func (o *CreateNetworkRequest) SetSave(v bool) {
	o.Save = &v
}

func (o CreateNetworkRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["cloud"] = o.Cloud
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Extra != nil {
		toSerialize["extra"] = o.Extra
	}
	if o.Template != nil {
		toSerialize["template"] = o.Template
	}
	if o.Dry != nil {
		toSerialize["dry"] = o.Dry
	}
	if o.Save != nil {
		toSerialize["save"] = o.Save
	}
	return json.Marshal(toSerialize)
}

type NullableCreateNetworkRequest struct {
	value *CreateNetworkRequest
	isSet bool
}

func (v NullableCreateNetworkRequest) Get() *CreateNetworkRequest {
	return v.value
}

func (v *NullableCreateNetworkRequest) Set(val *CreateNetworkRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetworkRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetworkRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetworkRequest(val *CreateNetworkRequest) *NullableCreateNetworkRequest {
	return &NullableCreateNetworkRequest{value: val, isSet: true}
}

func (v NullableCreateNetworkRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetworkRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


