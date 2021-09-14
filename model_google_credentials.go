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

// GoogleCredentials struct for GoogleCredentials
type GoogleCredentials struct {
	// The Id of your GCP project
	ProjectId string `json:"projectId"`
	// Your GCP private key
	PrivateKey string `json:"privateKey"`
	// Your GCP client email
	Email Email `json:"email"`
}

// NewGoogleCredentials instantiates a new GoogleCredentials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGoogleCredentials(projectId string, privateKey string, email Email, ) *GoogleCredentials {
	this := GoogleCredentials{}
	this.ProjectId = projectId
	this.PrivateKey = privateKey
	this.Email = email
	return &this
}

// NewGoogleCredentialsWithDefaults instantiates a new GoogleCredentials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGoogleCredentialsWithDefaults() *GoogleCredentials {
	this := GoogleCredentials{}
	return &this
}

// GetProjectId returns the ProjectId field value
func (o *GoogleCredentials) GetProjectId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *GoogleCredentials) GetProjectIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *GoogleCredentials) SetProjectId(v string) {
	o.ProjectId = v
}

// GetPrivateKey returns the PrivateKey field value
func (o *GoogleCredentials) GetPrivateKey() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.PrivateKey
}

// GetPrivateKeyOk returns a tuple with the PrivateKey field value
// and a boolean to check if the value has been set.
func (o *GoogleCredentials) GetPrivateKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PrivateKey, true
}

// SetPrivateKey sets field value
func (o *GoogleCredentials) SetPrivateKey(v string) {
	o.PrivateKey = v
}

// GetEmail returns the Email field value
func (o *GoogleCredentials) GetEmail() Email {
	if o == nil  {
		var ret Email
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *GoogleCredentials) GetEmailOk() (*Email, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *GoogleCredentials) SetEmail(v Email) {
	o.Email = v
}

func (o GoogleCredentials) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["projectId"] = o.ProjectId
	}
	if true {
		toSerialize["privateKey"] = o.PrivateKey
	}
	if true {
		toSerialize["email"] = o.Email
	}
	return json.Marshal(toSerialize)
}

type NullableGoogleCredentials struct {
	value *GoogleCredentials
	isSet bool
}

func (v NullableGoogleCredentials) Get() *GoogleCredentials {
	return v.value
}

func (v *NullableGoogleCredentials) Set(val *GoogleCredentials) {
	v.value = val
	v.isSet = true
}

func (v NullableGoogleCredentials) IsSet() bool {
	return v.isSet
}

func (v *NullableGoogleCredentials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGoogleCredentials(val *GoogleCredentials) *NullableGoogleCredentials {
	return &NullableGoogleCredentials{value: val, isSet: true}
}

func (v NullableGoogleCredentials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGoogleCredentials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


