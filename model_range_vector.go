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

// RangeVector struct for RangeVector
type RangeVector struct {
	Metric *map[string]interface{} `json:"metric,omitempty"`
	Value *[][]DatapointsValuesItem `json:"value,omitempty"`
}

// NewRangeVector instantiates a new RangeVector object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRangeVector() *RangeVector {
	this := RangeVector{}
	return &this
}

// NewRangeVectorWithDefaults instantiates a new RangeVector object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRangeVectorWithDefaults() *RangeVector {
	this := RangeVector{}
	return &this
}

// GetMetric returns the Metric field value if set, zero value otherwise.
func (o *RangeVector) GetMetric() map[string]interface{} {
	if o == nil || o.Metric == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Metric
}

// GetMetricOk returns a tuple with the Metric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RangeVector) GetMetricOk() (*map[string]interface{}, bool) {
	if o == nil || o.Metric == nil {
		return nil, false
	}
	return o.Metric, true
}

// HasMetric returns a boolean if a field has been set.
func (o *RangeVector) HasMetric() bool {
	if o != nil && o.Metric != nil {
		return true
	}

	return false
}

// SetMetric gets a reference to the given map[string]interface{} and assigns it to the Metric field.
func (o *RangeVector) SetMetric(v map[string]interface{}) {
	o.Metric = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *RangeVector) GetValue() [][]DatapointsValuesItem {
	if o == nil || o.Value == nil {
		var ret [][]DatapointsValuesItem
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RangeVector) GetValueOk() (*[][]DatapointsValuesItem, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *RangeVector) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given [][]DatapointsValuesItem and assigns it to the Value field.
func (o *RangeVector) SetValue(v [][]DatapointsValuesItem) {
	o.Value = &v
}

func (o RangeVector) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Metric != nil {
		toSerialize["metric"] = o.Metric
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableRangeVector struct {
	value *RangeVector
	isSet bool
}

func (v NullableRangeVector) Get() *RangeVector {
	return v.value
}

func (v *NullableRangeVector) Set(val *RangeVector) {
	v.value = val
	v.isSet = true
}

func (v NullableRangeVector) IsSet() bool {
	return v.isSet
}

func (v *NullableRangeVector) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRangeVector(val *RangeVector) *NullableRangeVector {
	return &NullableRangeVector{value: val, isSet: true}
}

func (v NullableRangeVector) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRangeVector) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

