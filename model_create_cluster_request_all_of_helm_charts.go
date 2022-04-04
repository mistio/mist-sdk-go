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

// CreateClusterRequestAllOfHelmCharts struct for CreateClusterRequestAllOfHelmCharts
type CreateClusterRequestAllOfHelmCharts struct {
	// The url of the Helm repository
	RepoUrl string `json:"repo_url"`
	ChartName string `json:"chart_name"`
	ReleaseName string `json:"release_name"`
	// The namespace to install the release on
	Namespace *string `json:"namespace,omitempty"`
	// The contents of a Helm values.yaml file
	Values *string `json:"values,omitempty"`
	// A version constraint for the chart
	Version *string `json:"version,omitempty"`
}

// NewCreateClusterRequestAllOfHelmCharts instantiates a new CreateClusterRequestAllOfHelmCharts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateClusterRequestAllOfHelmCharts(repoUrl string, chartName string, releaseName string, ) *CreateClusterRequestAllOfHelmCharts {
	this := CreateClusterRequestAllOfHelmCharts{}
	this.RepoUrl = repoUrl
	this.ChartName = chartName
	this.ReleaseName = releaseName
	return &this
}

// NewCreateClusterRequestAllOfHelmChartsWithDefaults instantiates a new CreateClusterRequestAllOfHelmCharts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateClusterRequestAllOfHelmChartsWithDefaults() *CreateClusterRequestAllOfHelmCharts {
	this := CreateClusterRequestAllOfHelmCharts{}
	return &this
}

// GetRepoUrl returns the RepoUrl field value
func (o *CreateClusterRequestAllOfHelmCharts) GetRepoUrl() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.RepoUrl
}

// GetRepoUrlOk returns a tuple with the RepoUrl field value
// and a boolean to check if the value has been set.
func (o *CreateClusterRequestAllOfHelmCharts) GetRepoUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RepoUrl, true
}

// SetRepoUrl sets field value
func (o *CreateClusterRequestAllOfHelmCharts) SetRepoUrl(v string) {
	o.RepoUrl = v
}

// GetChartName returns the ChartName field value
func (o *CreateClusterRequestAllOfHelmCharts) GetChartName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.ChartName
}

// GetChartNameOk returns a tuple with the ChartName field value
// and a boolean to check if the value has been set.
func (o *CreateClusterRequestAllOfHelmCharts) GetChartNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ChartName, true
}

// SetChartName sets field value
func (o *CreateClusterRequestAllOfHelmCharts) SetChartName(v string) {
	o.ChartName = v
}

// GetReleaseName returns the ReleaseName field value
func (o *CreateClusterRequestAllOfHelmCharts) GetReleaseName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.ReleaseName
}

// GetReleaseNameOk returns a tuple with the ReleaseName field value
// and a boolean to check if the value has been set.
func (o *CreateClusterRequestAllOfHelmCharts) GetReleaseNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ReleaseName, true
}

// SetReleaseName sets field value
func (o *CreateClusterRequestAllOfHelmCharts) SetReleaseName(v string) {
	o.ReleaseName = v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *CreateClusterRequestAllOfHelmCharts) GetNamespace() string {
	if o == nil || o.Namespace == nil {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateClusterRequestAllOfHelmCharts) GetNamespaceOk() (*string, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *CreateClusterRequestAllOfHelmCharts) HasNamespace() bool {
	if o != nil && o.Namespace != nil {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *CreateClusterRequestAllOfHelmCharts) SetNamespace(v string) {
	o.Namespace = &v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *CreateClusterRequestAllOfHelmCharts) GetValues() string {
	if o == nil || o.Values == nil {
		var ret string
		return ret
	}
	return *o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateClusterRequestAllOfHelmCharts) GetValuesOk() (*string, bool) {
	if o == nil || o.Values == nil {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *CreateClusterRequestAllOfHelmCharts) HasValues() bool {
	if o != nil && o.Values != nil {
		return true
	}

	return false
}

// SetValues gets a reference to the given string and assigns it to the Values field.
func (o *CreateClusterRequestAllOfHelmCharts) SetValues(v string) {
	o.Values = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *CreateClusterRequestAllOfHelmCharts) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateClusterRequestAllOfHelmCharts) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *CreateClusterRequestAllOfHelmCharts) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *CreateClusterRequestAllOfHelmCharts) SetVersion(v string) {
	o.Version = &v
}

func (o CreateClusterRequestAllOfHelmCharts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["repo_url"] = o.RepoUrl
	}
	if true {
		toSerialize["chart_name"] = o.ChartName
	}
	if true {
		toSerialize["release_name"] = o.ReleaseName
	}
	if o.Namespace != nil {
		toSerialize["namespace"] = o.Namespace
	}
	if o.Values != nil {
		toSerialize["values"] = o.Values
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableCreateClusterRequestAllOfHelmCharts struct {
	value *CreateClusterRequestAllOfHelmCharts
	isSet bool
}

func (v NullableCreateClusterRequestAllOfHelmCharts) Get() *CreateClusterRequestAllOfHelmCharts {
	return v.value
}

func (v *NullableCreateClusterRequestAllOfHelmCharts) Set(val *CreateClusterRequestAllOfHelmCharts) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateClusterRequestAllOfHelmCharts) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateClusterRequestAllOfHelmCharts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateClusterRequestAllOfHelmCharts(val *CreateClusterRequestAllOfHelmCharts) *NullableCreateClusterRequestAllOfHelmCharts {
	return &NullableCreateClusterRequestAllOfHelmCharts{value: val, isSet: true}
}

func (v NullableCreateClusterRequestAllOfHelmCharts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateClusterRequestAllOfHelmCharts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


