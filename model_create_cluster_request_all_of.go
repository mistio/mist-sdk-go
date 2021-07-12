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

// CreateClusterRequestAllOf struct for CreateClusterRequestAllOf
type CreateClusterRequestAllOf struct {
	// The name of the cluster to create
	Title string `json:"title"`
	// The cloud the cluster belongs to
	Cloud string `json:"cloud"`
	Provider *ClusterProviders `json:"provider,omitempty"`
}

// NewCreateClusterRequestAllOf instantiates a new CreateClusterRequestAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateClusterRequestAllOf(title string, cloud string, ) *CreateClusterRequestAllOf {
	this := CreateClusterRequestAllOf{}
	this.Title = title
	this.Cloud = cloud
	return &this
}

// NewCreateClusterRequestAllOfWithDefaults instantiates a new CreateClusterRequestAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateClusterRequestAllOfWithDefaults() *CreateClusterRequestAllOf {
	this := CreateClusterRequestAllOf{}
	return &this
}

// GetTitle returns the Title field value
func (o *CreateClusterRequestAllOf) GetTitle() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *CreateClusterRequestAllOf) GetTitleOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *CreateClusterRequestAllOf) SetTitle(v string) {
	o.Title = v
}

// GetCloud returns the Cloud field value
func (o *CreateClusterRequestAllOf) GetCloud() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Cloud
}

// GetCloudOk returns a tuple with the Cloud field value
// and a boolean to check if the value has been set.
func (o *CreateClusterRequestAllOf) GetCloudOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Cloud, true
}

// SetCloud sets field value
func (o *CreateClusterRequestAllOf) SetCloud(v string) {
	o.Cloud = v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *CreateClusterRequestAllOf) GetProvider() ClusterProviders {
	if o == nil || o.Provider == nil {
		var ret ClusterProviders
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateClusterRequestAllOf) GetProviderOk() (*ClusterProviders, bool) {
	if o == nil || o.Provider == nil {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *CreateClusterRequestAllOf) HasProvider() bool {
	if o != nil && o.Provider != nil {
		return true
	}

	return false
}

// SetProvider gets a reference to the given ClusterProviders and assigns it to the Provider field.
func (o *CreateClusterRequestAllOf) SetProvider(v ClusterProviders) {
	o.Provider = &v
}

func (o CreateClusterRequestAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["title"] = o.Title
	}
	if true {
		toSerialize["cloud"] = o.Cloud
	}
	if o.Provider != nil {
		toSerialize["provider"] = o.Provider
	}
	return json.Marshal(toSerialize)
}

type NullableCreateClusterRequestAllOf struct {
	value *CreateClusterRequestAllOf
	isSet bool
}

func (v NullableCreateClusterRequestAllOf) Get() *CreateClusterRequestAllOf {
	return v.value
}

func (v *NullableCreateClusterRequestAllOf) Set(val *CreateClusterRequestAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateClusterRequestAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateClusterRequestAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateClusterRequestAllOf(val *CreateClusterRequestAllOf) *NullableCreateClusterRequestAllOf {
	return &NullableCreateClusterRequestAllOf{value: val, isSet: true}
}

func (v NullableCreateClusterRequestAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateClusterRequestAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


