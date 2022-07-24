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

// CreateRecordRequest struct for CreateRecordRequest
type CreateRecordRequest struct {
	Name string `json:"name"`
	Cloud *string `json:"cloud,omitempty"`
	Value *string `json:"value,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewCreateRecordRequest instantiates a new CreateRecordRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateRecordRequest(name string) *CreateRecordRequest {
	this := CreateRecordRequest{}
	this.Name = name
	return &this
}

// NewCreateRecordRequestWithDefaults instantiates a new CreateRecordRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateRecordRequestWithDefaults() *CreateRecordRequest {
	this := CreateRecordRequest{}
	return &this
}

// GetName returns the Name field value
func (o *CreateRecordRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateRecordRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateRecordRequest) SetName(v string) {
	o.Name = v
}

// GetCloud returns the Cloud field value if set, zero value otherwise.
func (o *CreateRecordRequest) GetCloud() string {
	if o == nil || o.Cloud == nil {
		var ret string
		return ret
	}
	return *o.Cloud
}

// GetCloudOk returns a tuple with the Cloud field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRecordRequest) GetCloudOk() (*string, bool) {
	if o == nil || o.Cloud == nil {
		return nil, false
	}
	return o.Cloud, true
}

// HasCloud returns a boolean if a field has been set.
func (o *CreateRecordRequest) HasCloud() bool {
	if o != nil && o.Cloud != nil {
		return true
	}

	return false
}

// SetCloud gets a reference to the given string and assigns it to the Cloud field.
func (o *CreateRecordRequest) SetCloud(v string) {
	o.Cloud = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CreateRecordRequest) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRecordRequest) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CreateRecordRequest) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *CreateRecordRequest) SetValue(v string) {
	o.Value = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CreateRecordRequest) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRecordRequest) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CreateRecordRequest) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CreateRecordRequest) SetType(v string) {
	o.Type = &v
}

func (o CreateRecordRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Cloud != nil {
		toSerialize["cloud"] = o.Cloud
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableCreateRecordRequest struct {
	value *CreateRecordRequest
	isSet bool
}

func (v NullableCreateRecordRequest) Get() *CreateRecordRequest {
	return v.value
}

func (v *NullableCreateRecordRequest) Set(val *CreateRecordRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateRecordRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateRecordRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateRecordRequest(val *CreateRecordRequest) *NullableCreateRecordRequest {
	return &NullableCreateRecordRequest{value: val, isSet: true}
}

func (v NullableCreateRecordRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateRecordRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


