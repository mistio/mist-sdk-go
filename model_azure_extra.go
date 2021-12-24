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

// AzureExtra struct for AzureExtra
type AzureExtra struct {
	// The machine password. Only used on Microsoft images
	Password *string `json:"password,omitempty"`
	// A new or existing resource group. If not provided a `mist` resource group will be used.
	ResourceGroup *string `json:"resource_group,omitempty"`
	// Specifies the storage account type for the OS disk. Defaults to `StandardSSD_LRS`
	StorageAccountType *string `json:"storage_account_type,omitempty"`
	// The machine username. Defaults to azureuser
	User *string `json:"user,omitempty"`
}

// NewAzureExtra instantiates a new AzureExtra object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureExtra() *AzureExtra {
	this := AzureExtra{}
	return &this
}

// NewAzureExtraWithDefaults instantiates a new AzureExtra object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureExtraWithDefaults() *AzureExtra {
	this := AzureExtra{}
	return &this
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *AzureExtra) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureExtra) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *AzureExtra) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *AzureExtra) SetPassword(v string) {
	o.Password = &v
}

// GetResourceGroup returns the ResourceGroup field value if set, zero value otherwise.
func (o *AzureExtra) GetResourceGroup() string {
	if o == nil || o.ResourceGroup == nil {
		var ret string
		return ret
	}
	return *o.ResourceGroup
}

// GetResourceGroupOk returns a tuple with the ResourceGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureExtra) GetResourceGroupOk() (*string, bool) {
	if o == nil || o.ResourceGroup == nil {
		return nil, false
	}
	return o.ResourceGroup, true
}

// HasResourceGroup returns a boolean if a field has been set.
func (o *AzureExtra) HasResourceGroup() bool {
	if o != nil && o.ResourceGroup != nil {
		return true
	}

	return false
}

// SetResourceGroup gets a reference to the given string and assigns it to the ResourceGroup field.
func (o *AzureExtra) SetResourceGroup(v string) {
	o.ResourceGroup = &v
}

// GetStorageAccountType returns the StorageAccountType field value if set, zero value otherwise.
func (o *AzureExtra) GetStorageAccountType() string {
	if o == nil || o.StorageAccountType == nil {
		var ret string
		return ret
	}
	return *o.StorageAccountType
}

// GetStorageAccountTypeOk returns a tuple with the StorageAccountType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureExtra) GetStorageAccountTypeOk() (*string, bool) {
	if o == nil || o.StorageAccountType == nil {
		return nil, false
	}
	return o.StorageAccountType, true
}

// HasStorageAccountType returns a boolean if a field has been set.
func (o *AzureExtra) HasStorageAccountType() bool {
	if o != nil && o.StorageAccountType != nil {
		return true
	}

	return false
}

// SetStorageAccountType gets a reference to the given string and assigns it to the StorageAccountType field.
func (o *AzureExtra) SetStorageAccountType(v string) {
	o.StorageAccountType = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *AzureExtra) GetUser() string {
	if o == nil || o.User == nil {
		var ret string
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureExtra) GetUserOk() (*string, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *AzureExtra) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given string and assigns it to the User field.
func (o *AzureExtra) SetUser(v string) {
	o.User = &v
}

func (o AzureExtra) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	if o.ResourceGroup != nil {
		toSerialize["resource_group"] = o.ResourceGroup
	}
	if o.StorageAccountType != nil {
		toSerialize["storage_account_type"] = o.StorageAccountType
	}
	if o.User != nil {
		toSerialize["user"] = o.User
	}
	return json.Marshal(toSerialize)
}

type NullableAzureExtra struct {
	value *AzureExtra
	isSet bool
}

func (v NullableAzureExtra) Get() *AzureExtra {
	return v.value
}

func (v *NullableAzureExtra) Set(val *AzureExtra) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureExtra) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureExtra) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureExtra(val *AzureExtra) *NullableAzureExtra {
	return &NullableAzureExtra{value: val, isSet: true}
}

func (v NullableAzureExtra) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureExtra) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


