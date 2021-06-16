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

// Query struct for Query
type Query struct {
	// the metric's name, e.g. \"load.shortterm\"
	Target string `json:"target"`
	// the operator used to compare the computed value with the given threshold 
	Operator string `json:"operator"`
	// the value over/under which an alert will be raised
	Threshold string `json:"threshold"`
	// the function to be applied on the computed series. Must be one of: all, any, avg 
	Aggregation string `json:"aggregation"`
}

// NewQuery instantiates a new Query object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuery(target string, operator string, threshold string, aggregation string, ) *Query {
	this := Query{}
	this.Target = target
	this.Operator = operator
	this.Threshold = threshold
	this.Aggregation = aggregation
	return &this
}

// NewQueryWithDefaults instantiates a new Query object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryWithDefaults() *Query {
	this := Query{}
	return &this
}

// GetTarget returns the Target field value
func (o *Query) GetTarget() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Target
}

// GetTargetOk returns a tuple with the Target field value
// and a boolean to check if the value has been set.
func (o *Query) GetTargetOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Target, true
}

// SetTarget sets field value
func (o *Query) SetTarget(v string) {
	o.Target = v
}

// GetOperator returns the Operator field value
func (o *Query) GetOperator() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value
// and a boolean to check if the value has been set.
func (o *Query) GetOperatorOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Operator, true
}

// SetOperator sets field value
func (o *Query) SetOperator(v string) {
	o.Operator = v
}

// GetThreshold returns the Threshold field value
func (o *Query) GetThreshold() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Threshold
}

// GetThresholdOk returns a tuple with the Threshold field value
// and a boolean to check if the value has been set.
func (o *Query) GetThresholdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Threshold, true
}

// SetThreshold sets field value
func (o *Query) SetThreshold(v string) {
	o.Threshold = v
}

// GetAggregation returns the Aggregation field value
func (o *Query) GetAggregation() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Aggregation
}

// GetAggregationOk returns a tuple with the Aggregation field value
// and a boolean to check if the value has been set.
func (o *Query) GetAggregationOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Aggregation, true
}

// SetAggregation sets field value
func (o *Query) SetAggregation(v string) {
	o.Aggregation = v
}

func (o Query) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["target"] = o.Target
	}
	if true {
		toSerialize["operator"] = o.Operator
	}
	if true {
		toSerialize["threshold"] = o.Threshold
	}
	if true {
		toSerialize["aggregation"] = o.Aggregation
	}
	return json.Marshal(toSerialize)
}

type NullableQuery struct {
	value *Query
	isSet bool
}

func (v NullableQuery) Get() *Query {
	return v.value
}

func (v *NullableQuery) Set(val *Query) {
	v.value = val
	v.isSet = true
}

func (v NullableQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuery(val *Query) *NullableQuery {
	return &NullableQuery{value: val, isSet: true}
}

func (v NullableQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


