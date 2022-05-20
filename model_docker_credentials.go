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

// DockerCredentials struct for DockerCredentials
type DockerCredentials struct {
	// Your Docker host
	Host string `json:"host"`
	// Your Docker port
	Port string `json:"port"`
	// Your Docker username
	Username *string `json:"username,omitempty"`
	// Your Docker password
	Password *string `json:"password,omitempty"`
	// Your TLS auth key
	TlsKey *string `json:"tlsKey,omitempty"`
	// Your TLS auth certificate
	TlsCert *string `json:"tlsCert,omitempty"`
	// Your TLS CA certifcate
	TlsCaCert *string `json:"tlsCaCert,omitempty"`
	// Show all containers, including stopped
	ShowAll *bool `json:"showAll,omitempty"`
}

// NewDockerCredentials instantiates a new DockerCredentials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDockerCredentials(host string, port string) *DockerCredentials {
	this := DockerCredentials{}
	this.Host = host
	this.Port = port
	return &this
}

// NewDockerCredentialsWithDefaults instantiates a new DockerCredentials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDockerCredentialsWithDefaults() *DockerCredentials {
	this := DockerCredentials{}
	return &this
}

// GetHost returns the Host field value
func (o *DockerCredentials) GetHost() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Host
}

// GetHostOk returns a tuple with the Host field value
// and a boolean to check if the value has been set.
func (o *DockerCredentials) GetHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Host, true
}

// SetHost sets field value
func (o *DockerCredentials) SetHost(v string) {
	o.Host = v
}

// GetPort returns the Port field value
func (o *DockerCredentials) GetPort() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *DockerCredentials) GetPortOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *DockerCredentials) SetPort(v string) {
	o.Port = v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *DockerCredentials) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DockerCredentials) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *DockerCredentials) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *DockerCredentials) SetUsername(v string) {
	o.Username = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *DockerCredentials) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DockerCredentials) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *DockerCredentials) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *DockerCredentials) SetPassword(v string) {
	o.Password = &v
}

// GetTlsKey returns the TlsKey field value if set, zero value otherwise.
func (o *DockerCredentials) GetTlsKey() string {
	if o == nil || o.TlsKey == nil {
		var ret string
		return ret
	}
	return *o.TlsKey
}

// GetTlsKeyOk returns a tuple with the TlsKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DockerCredentials) GetTlsKeyOk() (*string, bool) {
	if o == nil || o.TlsKey == nil {
		return nil, false
	}
	return o.TlsKey, true
}

// HasTlsKey returns a boolean if a field has been set.
func (o *DockerCredentials) HasTlsKey() bool {
	if o != nil && o.TlsKey != nil {
		return true
	}

	return false
}

// SetTlsKey gets a reference to the given string and assigns it to the TlsKey field.
func (o *DockerCredentials) SetTlsKey(v string) {
	o.TlsKey = &v
}

// GetTlsCert returns the TlsCert field value if set, zero value otherwise.
func (o *DockerCredentials) GetTlsCert() string {
	if o == nil || o.TlsCert == nil {
		var ret string
		return ret
	}
	return *o.TlsCert
}

// GetTlsCertOk returns a tuple with the TlsCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DockerCredentials) GetTlsCertOk() (*string, bool) {
	if o == nil || o.TlsCert == nil {
		return nil, false
	}
	return o.TlsCert, true
}

// HasTlsCert returns a boolean if a field has been set.
func (o *DockerCredentials) HasTlsCert() bool {
	if o != nil && o.TlsCert != nil {
		return true
	}

	return false
}

// SetTlsCert gets a reference to the given string and assigns it to the TlsCert field.
func (o *DockerCredentials) SetTlsCert(v string) {
	o.TlsCert = &v
}

// GetTlsCaCert returns the TlsCaCert field value if set, zero value otherwise.
func (o *DockerCredentials) GetTlsCaCert() string {
	if o == nil || o.TlsCaCert == nil {
		var ret string
		return ret
	}
	return *o.TlsCaCert
}

// GetTlsCaCertOk returns a tuple with the TlsCaCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DockerCredentials) GetTlsCaCertOk() (*string, bool) {
	if o == nil || o.TlsCaCert == nil {
		return nil, false
	}
	return o.TlsCaCert, true
}

// HasTlsCaCert returns a boolean if a field has been set.
func (o *DockerCredentials) HasTlsCaCert() bool {
	if o != nil && o.TlsCaCert != nil {
		return true
	}

	return false
}

// SetTlsCaCert gets a reference to the given string and assigns it to the TlsCaCert field.
func (o *DockerCredentials) SetTlsCaCert(v string) {
	o.TlsCaCert = &v
}

// GetShowAll returns the ShowAll field value if set, zero value otherwise.
func (o *DockerCredentials) GetShowAll() bool {
	if o == nil || o.ShowAll == nil {
		var ret bool
		return ret
	}
	return *o.ShowAll
}

// GetShowAllOk returns a tuple with the ShowAll field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DockerCredentials) GetShowAllOk() (*bool, bool) {
	if o == nil || o.ShowAll == nil {
		return nil, false
	}
	return o.ShowAll, true
}

// HasShowAll returns a boolean if a field has been set.
func (o *DockerCredentials) HasShowAll() bool {
	if o != nil && o.ShowAll != nil {
		return true
	}

	return false
}

// SetShowAll gets a reference to the given bool and assigns it to the ShowAll field.
func (o *DockerCredentials) SetShowAll(v bool) {
	o.ShowAll = &v
}

func (o DockerCredentials) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["host"] = o.Host
	}
	if true {
		toSerialize["port"] = o.Port
	}
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	if o.TlsKey != nil {
		toSerialize["tlsKey"] = o.TlsKey
	}
	if o.TlsCert != nil {
		toSerialize["tlsCert"] = o.TlsCert
	}
	if o.TlsCaCert != nil {
		toSerialize["tlsCaCert"] = o.TlsCaCert
	}
	if o.ShowAll != nil {
		toSerialize["showAll"] = o.ShowAll
	}
	return json.Marshal(toSerialize)
}

type NullableDockerCredentials struct {
	value *DockerCredentials
	isSet bool
}

func (v NullableDockerCredentials) Get() *DockerCredentials {
	return v.value
}

func (v *NullableDockerCredentials) Set(val *DockerCredentials) {
	v.value = val
	v.isSet = true
}

func (v NullableDockerCredentials) IsSet() bool {
	return v.isSet
}

func (v *NullableDockerCredentials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDockerCredentials(val *DockerCredentials) *NullableDockerCredentials {
	return &NullableDockerCredentials{value: val, isSet: true}
}

func (v NullableDockerCredentials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDockerCredentials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


