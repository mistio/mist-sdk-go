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

// CreateMachineRequestDisks Configure local disks
type CreateMachineRequestDisks struct {
	// KVM, CloudSigma specific parameter
	DiskSize *int32 `json:"disk_size,omitempty"`
	// KVM specific parameter. Where the VM disk file will be created
	DiskPath *string `json:"disk_path,omitempty"`
}

// NewCreateMachineRequestDisks instantiates a new CreateMachineRequestDisks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateMachineRequestDisks() *CreateMachineRequestDisks {
	this := CreateMachineRequestDisks{}
	return &this
}

// NewCreateMachineRequestDisksWithDefaults instantiates a new CreateMachineRequestDisks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateMachineRequestDisksWithDefaults() *CreateMachineRequestDisks {
	this := CreateMachineRequestDisks{}
	return &this
}

// GetDiskSize returns the DiskSize field value if set, zero value otherwise.
func (o *CreateMachineRequestDisks) GetDiskSize() int32 {
	if o == nil || o.DiskSize == nil {
		var ret int32
		return ret
	}
	return *o.DiskSize
}

// GetDiskSizeOk returns a tuple with the DiskSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMachineRequestDisks) GetDiskSizeOk() (*int32, bool) {
	if o == nil || o.DiskSize == nil {
		return nil, false
	}
	return o.DiskSize, true
}

// HasDiskSize returns a boolean if a field has been set.
func (o *CreateMachineRequestDisks) HasDiskSize() bool {
	if o != nil && o.DiskSize != nil {
		return true
	}

	return false
}

// SetDiskSize gets a reference to the given int32 and assigns it to the DiskSize field.
func (o *CreateMachineRequestDisks) SetDiskSize(v int32) {
	o.DiskSize = &v
}

// GetDiskPath returns the DiskPath field value if set, zero value otherwise.
func (o *CreateMachineRequestDisks) GetDiskPath() string {
	if o == nil || o.DiskPath == nil {
		var ret string
		return ret
	}
	return *o.DiskPath
}

// GetDiskPathOk returns a tuple with the DiskPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMachineRequestDisks) GetDiskPathOk() (*string, bool) {
	if o == nil || o.DiskPath == nil {
		return nil, false
	}
	return o.DiskPath, true
}

// HasDiskPath returns a boolean if a field has been set.
func (o *CreateMachineRequestDisks) HasDiskPath() bool {
	if o != nil && o.DiskPath != nil {
		return true
	}

	return false
}

// SetDiskPath gets a reference to the given string and assigns it to the DiskPath field.
func (o *CreateMachineRequestDisks) SetDiskPath(v string) {
	o.DiskPath = &v
}

func (o CreateMachineRequestDisks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DiskSize != nil {
		toSerialize["disk_size"] = o.DiskSize
	}
	if o.DiskPath != nil {
		toSerialize["disk_path"] = o.DiskPath
	}
	return json.Marshal(toSerialize)
}

type NullableCreateMachineRequestDisks struct {
	value *CreateMachineRequestDisks
	isSet bool
}

func (v NullableCreateMachineRequestDisks) Get() *CreateMachineRequestDisks {
	return v.value
}

func (v *NullableCreateMachineRequestDisks) Set(val *CreateMachineRequestDisks) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateMachineRequestDisks) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateMachineRequestDisks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateMachineRequestDisks(val *CreateMachineRequestDisks) *NullableCreateMachineRequestDisks {
	return &NullableCreateMachineRequestDisks{value: val, isSet: true}
}

func (v NullableCreateMachineRequestDisks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateMachineRequestDisks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


