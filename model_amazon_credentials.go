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

// AmazonCredentials struct for AmazonCredentials
type AmazonCredentials struct {
	// Your AWS API key
	Apikey string `json:"apikey"`
	// Your AWS API secret
	Apisecret string `json:"apisecret"`
	// Your AWS region
	Region string `json:"region"`
}

// NewAmazonCredentials instantiates a new AmazonCredentials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAmazonCredentials(apikey string, apisecret string, region string, ) *AmazonCredentials {
	this := AmazonCredentials{}
	this.Apikey = apikey
	this.Apisecret = apisecret
	this.Region = region
	return &this
}

// NewAmazonCredentialsWithDefaults instantiates a new AmazonCredentials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAmazonCredentialsWithDefaults() *AmazonCredentials {
	this := AmazonCredentials{}
	return &this
}

// GetApikey returns the Apikey field value
func (o *AmazonCredentials) GetApikey() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Apikey
}

// GetApikeyOk returns a tuple with the Apikey field value
// and a boolean to check if the value has been set.
func (o *AmazonCredentials) GetApikeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Apikey, true
}

// SetApikey sets field value
func (o *AmazonCredentials) SetApikey(v string) {
	o.Apikey = v
}

// GetApisecret returns the Apisecret field value
func (o *AmazonCredentials) GetApisecret() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Apisecret
}

// GetApisecretOk returns a tuple with the Apisecret field value
// and a boolean to check if the value has been set.
func (o *AmazonCredentials) GetApisecretOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Apisecret, true
}

// SetApisecret sets field value
func (o *AmazonCredentials) SetApisecret(v string) {
	o.Apisecret = v
}

// GetRegion returns the Region field value
func (o *AmazonCredentials) GetRegion() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *AmazonCredentials) GetRegionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value
func (o *AmazonCredentials) SetRegion(v string) {
	o.Region = v
}

func (o AmazonCredentials) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["apikey"] = o.Apikey
	}
	if true {
		toSerialize["apisecret"] = o.Apisecret
	}
	if true {
		toSerialize["region"] = o.Region
	}
	return json.Marshal(toSerialize)
}

type NullableAmazonCredentials struct {
	value *AmazonCredentials
	isSet bool
}

func (v NullableAmazonCredentials) Get() *AmazonCredentials {
	return v.value
}

func (v *NullableAmazonCredentials) Set(val *AmazonCredentials) {
	v.value = val
	v.isSet = true
}

func (v NullableAmazonCredentials) IsSet() bool {
	return v.isSet
}

func (v *NullableAmazonCredentials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmazonCredentials(val *AmazonCredentials) *NullableAmazonCredentials {
	return &NullableAmazonCredentials{value: val, isSet: true}
}

func (v NullableAmazonCredentials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAmazonCredentials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


