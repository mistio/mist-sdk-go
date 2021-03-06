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

// ResponseMetadata struct for ResponseMetadata
type ResponseMetadata struct {
	// Total items matching the query
	Total *int32 `json:"total,omitempty"`
	// Number of items in response
	Returned *int32 `json:"returned,omitempty"`
	// Sort order of results
	Sort *string `json:"sort,omitempty"`
	// Index of first response item in total matching items
	Start *int32 `json:"start,omitempty"`
}

// NewResponseMetadata instantiates a new ResponseMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseMetadata() *ResponseMetadata {
	this := ResponseMetadata{}
	return &this
}

// NewResponseMetadataWithDefaults instantiates a new ResponseMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseMetadataWithDefaults() *ResponseMetadata {
	this := ResponseMetadata{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *ResponseMetadata) GetTotal() int32 {
	if o == nil || o.Total == nil {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseMetadata) GetTotalOk() (*int32, bool) {
	if o == nil || o.Total == nil {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *ResponseMetadata) HasTotal() bool {
	if o != nil && o.Total != nil {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *ResponseMetadata) SetTotal(v int32) {
	o.Total = &v
}

// GetReturned returns the Returned field value if set, zero value otherwise.
func (o *ResponseMetadata) GetReturned() int32 {
	if o == nil || o.Returned == nil {
		var ret int32
		return ret
	}
	return *o.Returned
}

// GetReturnedOk returns a tuple with the Returned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseMetadata) GetReturnedOk() (*int32, bool) {
	if o == nil || o.Returned == nil {
		return nil, false
	}
	return o.Returned, true
}

// HasReturned returns a boolean if a field has been set.
func (o *ResponseMetadata) HasReturned() bool {
	if o != nil && o.Returned != nil {
		return true
	}

	return false
}

// SetReturned gets a reference to the given int32 and assigns it to the Returned field.
func (o *ResponseMetadata) SetReturned(v int32) {
	o.Returned = &v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *ResponseMetadata) GetSort() string {
	if o == nil || o.Sort == nil {
		var ret string
		return ret
	}
	return *o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseMetadata) GetSortOk() (*string, bool) {
	if o == nil || o.Sort == nil {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *ResponseMetadata) HasSort() bool {
	if o != nil && o.Sort != nil {
		return true
	}

	return false
}

// SetSort gets a reference to the given string and assigns it to the Sort field.
func (o *ResponseMetadata) SetSort(v string) {
	o.Sort = &v
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *ResponseMetadata) GetStart() int32 {
	if o == nil || o.Start == nil {
		var ret int32
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseMetadata) GetStartOk() (*int32, bool) {
	if o == nil || o.Start == nil {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *ResponseMetadata) HasStart() bool {
	if o != nil && o.Start != nil {
		return true
	}

	return false
}

// SetStart gets a reference to the given int32 and assigns it to the Start field.
func (o *ResponseMetadata) SetStart(v int32) {
	o.Start = &v
}

func (o ResponseMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Total != nil {
		toSerialize["total"] = o.Total
	}
	if o.Returned != nil {
		toSerialize["returned"] = o.Returned
	}
	if o.Sort != nil {
		toSerialize["sort"] = o.Sort
	}
	if o.Start != nil {
		toSerialize["start"] = o.Start
	}
	return json.Marshal(toSerialize)
}

type NullableResponseMetadata struct {
	value *ResponseMetadata
	isSet bool
}

func (v NullableResponseMetadata) Get() *ResponseMetadata {
	return v.value
}

func (v *NullableResponseMetadata) Set(val *ResponseMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseMetadata(val *ResponseMetadata) *NullableResponseMetadata {
	return &NullableResponseMetadata{value: val, isSet: true}
}

func (v NullableResponseMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


