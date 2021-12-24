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

// Member struct for Member
type Member struct {
	Email *string `json:"email,omitempty"`
	FirstName *string `json:"first_name,omitempty"`
	Id *string `json:"id,omitempty"`
	LastLogin *string `json:"last_login,omitempty"`
	LastName *string `json:"last_name,omitempty"`
	Pending *bool `json:"pending,omitempty"`
	RegistrationDate *string `json:"registration_date,omitempty"`
	Username *string `json:"username,omitempty"`
}

// NewMember instantiates a new Member object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMember() *Member {
	this := Member{}
	return &this
}

// NewMemberWithDefaults instantiates a new Member object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMemberWithDefaults() *Member {
	this := Member{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *Member) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Member) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *Member) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *Member) SetEmail(v string) {
	o.Email = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *Member) GetFirstName() string {
	if o == nil || o.FirstName == nil {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Member) GetFirstNameOk() (*string, bool) {
	if o == nil || o.FirstName == nil {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *Member) HasFirstName() bool {
	if o != nil && o.FirstName != nil {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *Member) SetFirstName(v string) {
	o.FirstName = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Member) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Member) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Member) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Member) SetId(v string) {
	o.Id = &v
}

// GetLastLogin returns the LastLogin field value if set, zero value otherwise.
func (o *Member) GetLastLogin() string {
	if o == nil || o.LastLogin == nil {
		var ret string
		return ret
	}
	return *o.LastLogin
}

// GetLastLoginOk returns a tuple with the LastLogin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Member) GetLastLoginOk() (*string, bool) {
	if o == nil || o.LastLogin == nil {
		return nil, false
	}
	return o.LastLogin, true
}

// HasLastLogin returns a boolean if a field has been set.
func (o *Member) HasLastLogin() bool {
	if o != nil && o.LastLogin != nil {
		return true
	}

	return false
}

// SetLastLogin gets a reference to the given string and assigns it to the LastLogin field.
func (o *Member) SetLastLogin(v string) {
	o.LastLogin = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *Member) GetLastName() string {
	if o == nil || o.LastName == nil {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Member) GetLastNameOk() (*string, bool) {
	if o == nil || o.LastName == nil {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *Member) HasLastName() bool {
	if o != nil && o.LastName != nil {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *Member) SetLastName(v string) {
	o.LastName = &v
}

// GetPending returns the Pending field value if set, zero value otherwise.
func (o *Member) GetPending() bool {
	if o == nil || o.Pending == nil {
		var ret bool
		return ret
	}
	return *o.Pending
}

// GetPendingOk returns a tuple with the Pending field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Member) GetPendingOk() (*bool, bool) {
	if o == nil || o.Pending == nil {
		return nil, false
	}
	return o.Pending, true
}

// HasPending returns a boolean if a field has been set.
func (o *Member) HasPending() bool {
	if o != nil && o.Pending != nil {
		return true
	}

	return false
}

// SetPending gets a reference to the given bool and assigns it to the Pending field.
func (o *Member) SetPending(v bool) {
	o.Pending = &v
}

// GetRegistrationDate returns the RegistrationDate field value if set, zero value otherwise.
func (o *Member) GetRegistrationDate() string {
	if o == nil || o.RegistrationDate == nil {
		var ret string
		return ret
	}
	return *o.RegistrationDate
}

// GetRegistrationDateOk returns a tuple with the RegistrationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Member) GetRegistrationDateOk() (*string, bool) {
	if o == nil || o.RegistrationDate == nil {
		return nil, false
	}
	return o.RegistrationDate, true
}

// HasRegistrationDate returns a boolean if a field has been set.
func (o *Member) HasRegistrationDate() bool {
	if o != nil && o.RegistrationDate != nil {
		return true
	}

	return false
}

// SetRegistrationDate gets a reference to the given string and assigns it to the RegistrationDate field.
func (o *Member) SetRegistrationDate(v string) {
	o.RegistrationDate = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *Member) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Member) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *Member) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *Member) SetUsername(v string) {
	o.Username = &v
}

func (o Member) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.FirstName != nil {
		toSerialize["first_name"] = o.FirstName
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.LastLogin != nil {
		toSerialize["last_login"] = o.LastLogin
	}
	if o.LastName != nil {
		toSerialize["last_name"] = o.LastName
	}
	if o.Pending != nil {
		toSerialize["pending"] = o.Pending
	}
	if o.RegistrationDate != nil {
		toSerialize["registration_date"] = o.RegistrationDate
	}
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	return json.Marshal(toSerialize)
}

type NullableMember struct {
	value *Member
	isSet bool
}

func (v NullableMember) Get() *Member {
	return v.value
}

func (v *NullableMember) Set(val *Member) {
	v.value = val
	v.isSet = true
}

func (v NullableMember) IsSet() bool {
	return v.isSet
}

func (v *NullableMember) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMember(val *Member) *NullableMember {
	return &NullableMember{value: val, isSet: true}
}

func (v NullableMember) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMember) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


