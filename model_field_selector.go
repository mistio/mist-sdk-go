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

// FieldSelector struct for FieldSelector
type FieldSelector struct {
	// field type
	Type *string `json:"type,omitempty"`
	// the name of the field
	Field *string `json:"field,omitempty"`
	// the value of the field
	Value *map[string]interface{} `json:"value,omitempty"`
	// one of equal (eq) and regular expression (regex) operators type
	Operator *string `json:"operator,omitempty"`
}

// NewFieldSelector instantiates a new FieldSelector object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFieldSelector() *FieldSelector {
	this := FieldSelector{}
	return &this
}

// NewFieldSelectorWithDefaults instantiates a new FieldSelector object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFieldSelectorWithDefaults() *FieldSelector {
	this := FieldSelector{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *FieldSelector) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldSelector) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *FieldSelector) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *FieldSelector) SetType(v string) {
	o.Type = &v
}

// GetField returns the Field field value if set, zero value otherwise.
func (o *FieldSelector) GetField() string {
	if o == nil || o.Field == nil {
		var ret string
		return ret
	}
	return *o.Field
}

// GetFieldOk returns a tuple with the Field field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldSelector) GetFieldOk() (*string, bool) {
	if o == nil || o.Field == nil {
		return nil, false
	}
	return o.Field, true
}

// HasField returns a boolean if a field has been set.
func (o *FieldSelector) HasField() bool {
	if o != nil && o.Field != nil {
		return true
	}

	return false
}

// SetField gets a reference to the given string and assigns it to the Field field.
func (o *FieldSelector) SetField(v string) {
	o.Field = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *FieldSelector) GetValue() map[string]interface{} {
	if o == nil || o.Value == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldSelector) GetValueOk() (*map[string]interface{}, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *FieldSelector) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given map[string]interface{} and assigns it to the Value field.
func (o *FieldSelector) SetValue(v map[string]interface{}) {
	o.Value = &v
}

// GetOperator returns the Operator field value if set, zero value otherwise.
func (o *FieldSelector) GetOperator() string {
	if o == nil || o.Operator == nil {
		var ret string
		return ret
	}
	return *o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldSelector) GetOperatorOk() (*string, bool) {
	if o == nil || o.Operator == nil {
		return nil, false
	}
	return o.Operator, true
}

// HasOperator returns a boolean if a field has been set.
func (o *FieldSelector) HasOperator() bool {
	if o != nil && o.Operator != nil {
		return true
	}

	return false
}

// SetOperator gets a reference to the given string and assigns it to the Operator field.
func (o *FieldSelector) SetOperator(v string) {
	o.Operator = &v
}

func (o FieldSelector) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Field != nil {
		toSerialize["field"] = o.Field
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.Operator != nil {
		toSerialize["operator"] = o.Operator
	}
	return json.Marshal(toSerialize)
}

type NullableFieldSelector struct {
	value *FieldSelector
	isSet bool
}

func (v NullableFieldSelector) Get() *FieldSelector {
	return v.value
}

func (v *NullableFieldSelector) Set(val *FieldSelector) {
	v.value = val
	v.isSet = true
}

func (v NullableFieldSelector) IsSet() bool {
	return v.isSet
}

func (v *NullableFieldSelector) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFieldSelector(val *FieldSelector) *NullableFieldSelector {
	return &NullableFieldSelector{value: val, isSet: true}
}

func (v NullableFieldSelector) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFieldSelector) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


