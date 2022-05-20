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

// DatapointsData struct for DatapointsData
type DatapointsData struct {
	ResultType *string `json:"resultType,omitempty"`
	Result []Vector `json:"result,omitempty"`
}

// NewDatapointsData instantiates a new DatapointsData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatapointsData() *DatapointsData {
	this := DatapointsData{}
	return &this
}

// NewDatapointsDataWithDefaults instantiates a new DatapointsData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatapointsDataWithDefaults() *DatapointsData {
	this := DatapointsData{}
	return &this
}

// GetResultType returns the ResultType field value if set, zero value otherwise.
func (o *DatapointsData) GetResultType() string {
	if o == nil || o.ResultType == nil {
		var ret string
		return ret
	}
	return *o.ResultType
}

// GetResultTypeOk returns a tuple with the ResultType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatapointsData) GetResultTypeOk() (*string, bool) {
	if o == nil || o.ResultType == nil {
		return nil, false
	}
	return o.ResultType, true
}

// HasResultType returns a boolean if a field has been set.
func (o *DatapointsData) HasResultType() bool {
	if o != nil && o.ResultType != nil {
		return true
	}

	return false
}

// SetResultType gets a reference to the given string and assigns it to the ResultType field.
func (o *DatapointsData) SetResultType(v string) {
	o.ResultType = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *DatapointsData) GetResult() []Vector {
	if o == nil || o.Result == nil {
		var ret []Vector
		return ret
	}
	return o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatapointsData) GetResultOk() ([]Vector, bool) {
	if o == nil || o.Result == nil {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *DatapointsData) HasResult() bool {
	if o != nil && o.Result != nil {
		return true
	}

	return false
}

// SetResult gets a reference to the given []Vector and assigns it to the Result field.
func (o *DatapointsData) SetResult(v []Vector) {
	o.Result = v
}

func (o DatapointsData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ResultType != nil {
		toSerialize["resultType"] = o.ResultType
	}
	if o.Result != nil {
		toSerialize["result"] = o.Result
	}
	return json.Marshal(toSerialize)
}

type NullableDatapointsData struct {
	value *DatapointsData
	isSet bool
}

func (v NullableDatapointsData) Get() *DatapointsData {
	return v.value
}

func (v *NullableDatapointsData) Set(val *DatapointsData) {
	v.value = val
	v.isSet = true
}

func (v NullableDatapointsData) IsSet() bool {
	return v.isSet
}

func (v *NullableDatapointsData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatapointsData(val *DatapointsData) *NullableDatapointsData {
	return &NullableDatapointsData{value: val, isSet: true}
}

func (v NullableDatapointsData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatapointsData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


