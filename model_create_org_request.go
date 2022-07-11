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

// CreateOrgRequest struct for CreateOrgRequest
type CreateOrgRequest struct {
	Name *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Logo *string `json:"logo,omitempty"`
}

// NewCreateOrgRequest instantiates a new CreateOrgRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrgRequest() *CreateOrgRequest {
	this := CreateOrgRequest{}
	return &this
}

// NewCreateOrgRequestWithDefaults instantiates a new CreateOrgRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrgRequestWithDefaults() *CreateOrgRequest {
	this := CreateOrgRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateOrgRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrgRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateOrgRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateOrgRequest) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateOrgRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrgRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateOrgRequest) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateOrgRequest) SetDescription(v string) {
	o.Description = &v
}

// GetLogo returns the Logo field value if set, zero value otherwise.
func (o *CreateOrgRequest) GetLogo() string {
	if o == nil || o.Logo == nil {
		var ret string
		return ret
	}
	return *o.Logo
}

// GetLogoOk returns a tuple with the Logo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrgRequest) GetLogoOk() (*string, bool) {
	if o == nil || o.Logo == nil {
		return nil, false
	}
	return o.Logo, true
}

// HasLogo returns a boolean if a field has been set.
func (o *CreateOrgRequest) HasLogo() bool {
	if o != nil && o.Logo != nil {
		return true
	}

	return false
}

// SetLogo gets a reference to the given string and assigns it to the Logo field.
func (o *CreateOrgRequest) SetLogo(v string) {
	o.Logo = &v
}

func (o CreateOrgRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Logo != nil {
		toSerialize["logo"] = o.Logo
	}
	return json.Marshal(toSerialize)
}

type NullableCreateOrgRequest struct {
	value *CreateOrgRequest
	isSet bool
}

func (v NullableCreateOrgRequest) Get() *CreateOrgRequest {
	return v.value
}

func (v *NullableCreateOrgRequest) Set(val *CreateOrgRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrgRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrgRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrgRequest(val *CreateOrgRequest) *NullableCreateOrgRequest {
	return &NullableCreateOrgRequest{value: val, isSet: true}
}

func (v NullableCreateOrgRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrgRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


