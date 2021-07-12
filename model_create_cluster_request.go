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

// CreateClusterRequest struct for CreateClusterRequest
type CreateClusterRequest struct {
	// The name of the cluster to create
	Title string `json:"title"`
	// The cloud the cluster belongs to
	Cloud string `json:"cloud"`
	Provider *ClusterProviders `json:"provider,omitempty"`
}

// NewCreateClusterRequest instantiates a new CreateClusterRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateClusterRequest(title string, cloud string, ) *CreateClusterRequest {
	this := CreateClusterRequest{}
	this.Title = title
	this.Cloud = cloud
	return &this
}

// NewCreateClusterRequestWithDefaults instantiates a new CreateClusterRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateClusterRequestWithDefaults() *CreateClusterRequest {
	this := CreateClusterRequest{}
	return &this
}

// GetTitle returns the Title field value
func (o *CreateClusterRequest) GetTitle() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *CreateClusterRequest) GetTitleOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *CreateClusterRequest) SetTitle(v string) {
	o.Title = v
}

// GetCloud returns the Cloud field value
func (o *CreateClusterRequest) GetCloud() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Cloud
}

// GetCloudOk returns a tuple with the Cloud field value
// and a boolean to check if the value has been set.
func (o *CreateClusterRequest) GetCloudOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Cloud, true
}

// SetCloud sets field value
func (o *CreateClusterRequest) SetCloud(v string) {
	o.Cloud = v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *CreateClusterRequest) GetProvider() ClusterProviders {
	if o == nil || o.Provider == nil {
		var ret ClusterProviders
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateClusterRequest) GetProviderOk() (*ClusterProviders, bool) {
	if o == nil || o.Provider == nil {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *CreateClusterRequest) HasProvider() bool {
	if o != nil && o.Provider != nil {
		return true
	}

	return false
}

// SetProvider gets a reference to the given ClusterProviders and assigns it to the Provider field.
func (o *CreateClusterRequest) SetProvider(v ClusterProviders) {
	o.Provider = &v
}

func (o CreateClusterRequest) MarshalJSON() ([]byte, error) {
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

type NullableCreateClusterRequest struct {
	value *CreateClusterRequest
	isSet bool
}

func (v NullableCreateClusterRequest) Get() *CreateClusterRequest {
	return v.value
}

func (v *NullableCreateClusterRequest) Set(val *CreateClusterRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateClusterRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateClusterRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateClusterRequest(val *CreateClusterRequest) *NullableCreateClusterRequest {
	return &NullableCreateClusterRequest{value: val, isSet: true}
}

func (v NullableCreateClusterRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateClusterRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


