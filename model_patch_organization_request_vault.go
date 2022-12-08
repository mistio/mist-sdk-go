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

// PatchOrganizationRequestVault struct for PatchOrganizationRequestVault
type PatchOrganizationRequestVault struct {
	// The Vault's url
	Address *string `json:"address,omitempty"`
	// Vault secrets engine path
	SecretsEnginePath *string `json:"secrets_engine_path,omitempty"`
	// The Vault token that will be used to authenticate against the new Vault. Either token or both role_id and secret_id must be specified
	Token *string `json:"token,omitempty"`
	// The Vault RoleID to use for approle authentication. Either token or both role_id and secret_id must be specified
	RoleId *string `json:"role_id,omitempty"`
	// The Vault SecretID to use for approle authentication. Either token or both role_id and secret_id must be specified
	SecretId *string `json:"secret_id,omitempty"`
}

// NewPatchOrganizationRequestVault instantiates a new PatchOrganizationRequestVault object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchOrganizationRequestVault() *PatchOrganizationRequestVault {
	this := PatchOrganizationRequestVault{}
	return &this
}

// NewPatchOrganizationRequestVaultWithDefaults instantiates a new PatchOrganizationRequestVault object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchOrganizationRequestVaultWithDefaults() *PatchOrganizationRequestVault {
	this := PatchOrganizationRequestVault{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *PatchOrganizationRequestVault) GetAddress() string {
	if o == nil || o.Address == nil {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchOrganizationRequestVault) GetAddressOk() (*string, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *PatchOrganizationRequestVault) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *PatchOrganizationRequestVault) SetAddress(v string) {
	o.Address = &v
}

// GetSecretsEnginePath returns the SecretsEnginePath field value if set, zero value otherwise.
func (o *PatchOrganizationRequestVault) GetSecretsEnginePath() string {
	if o == nil || o.SecretsEnginePath == nil {
		var ret string
		return ret
	}
	return *o.SecretsEnginePath
}

// GetSecretsEnginePathOk returns a tuple with the SecretsEnginePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchOrganizationRequestVault) GetSecretsEnginePathOk() (*string, bool) {
	if o == nil || o.SecretsEnginePath == nil {
		return nil, false
	}
	return o.SecretsEnginePath, true
}

// HasSecretsEnginePath returns a boolean if a field has been set.
func (o *PatchOrganizationRequestVault) HasSecretsEnginePath() bool {
	if o != nil && o.SecretsEnginePath != nil {
		return true
	}

	return false
}

// SetSecretsEnginePath gets a reference to the given string and assigns it to the SecretsEnginePath field.
func (o *PatchOrganizationRequestVault) SetSecretsEnginePath(v string) {
	o.SecretsEnginePath = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *PatchOrganizationRequestVault) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchOrganizationRequestVault) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *PatchOrganizationRequestVault) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *PatchOrganizationRequestVault) SetToken(v string) {
	o.Token = &v
}

// GetRoleId returns the RoleId field value if set, zero value otherwise.
func (o *PatchOrganizationRequestVault) GetRoleId() string {
	if o == nil || o.RoleId == nil {
		var ret string
		return ret
	}
	return *o.RoleId
}

// GetRoleIdOk returns a tuple with the RoleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchOrganizationRequestVault) GetRoleIdOk() (*string, bool) {
	if o == nil || o.RoleId == nil {
		return nil, false
	}
	return o.RoleId, true
}

// HasRoleId returns a boolean if a field has been set.
func (o *PatchOrganizationRequestVault) HasRoleId() bool {
	if o != nil && o.RoleId != nil {
		return true
	}

	return false
}

// SetRoleId gets a reference to the given string and assigns it to the RoleId field.
func (o *PatchOrganizationRequestVault) SetRoleId(v string) {
	o.RoleId = &v
}

// GetSecretId returns the SecretId field value if set, zero value otherwise.
func (o *PatchOrganizationRequestVault) GetSecretId() string {
	if o == nil || o.SecretId == nil {
		var ret string
		return ret
	}
	return *o.SecretId
}

// GetSecretIdOk returns a tuple with the SecretId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchOrganizationRequestVault) GetSecretIdOk() (*string, bool) {
	if o == nil || o.SecretId == nil {
		return nil, false
	}
	return o.SecretId, true
}

// HasSecretId returns a boolean if a field has been set.
func (o *PatchOrganizationRequestVault) HasSecretId() bool {
	if o != nil && o.SecretId != nil {
		return true
	}

	return false
}

// SetSecretId gets a reference to the given string and assigns it to the SecretId field.
func (o *PatchOrganizationRequestVault) SetSecretId(v string) {
	o.SecretId = &v
}

func (o PatchOrganizationRequestVault) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Address != nil {
		toSerialize["address"] = o.Address
	}
	if o.SecretsEnginePath != nil {
		toSerialize["secrets_engine_path"] = o.SecretsEnginePath
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.RoleId != nil {
		toSerialize["role_id"] = o.RoleId
	}
	if o.SecretId != nil {
		toSerialize["secret_id"] = o.SecretId
	}
	return json.Marshal(toSerialize)
}

type NullablePatchOrganizationRequestVault struct {
	value *PatchOrganizationRequestVault
	isSet bool
}

func (v NullablePatchOrganizationRequestVault) Get() *PatchOrganizationRequestVault {
	return v.value
}

func (v *NullablePatchOrganizationRequestVault) Set(val *PatchOrganizationRequestVault) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchOrganizationRequestVault) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchOrganizationRequestVault) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchOrganizationRequestVault(val *PatchOrganizationRequestVault) *NullablePatchOrganizationRequestVault {
	return &NullablePatchOrganizationRequestVault{value: val, isSet: true}
}

func (v NullablePatchOrganizationRequestVault) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchOrganizationRequestVault) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

