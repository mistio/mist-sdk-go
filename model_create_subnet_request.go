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

// CreateSubnetRequest struct for CreateSubnetRequest
type CreateSubnetRequest struct {
	// Specify subnet name
	Name string `json:"name"`
	// Specify cloud to provision on
	Cloud string `json:"cloud"`
	// Assign tags to provisioned network
	Tags *map[string]interface{} `json:"tags,omitempty"`
	Cidr string `json:"cidr"`
	Provider string `json:"provider"`
	AvailabilityZone string `json:"availability_zone"`
	Region string `json:"region"`
	GatewayIp *string `json:"gateway_ip,omitempty"`
	IpVersion *string `json:"ip_version,omitempty"`
	EnableDhcp *bool `json:"enable_dhcp,omitempty"`
	AllocationPools *map[string]interface{} `json:"allocation_pools,omitempty"`
}

// NewCreateSubnetRequest instantiates a new CreateSubnetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSubnetRequest(name string, cloud string, cidr string, provider string, availabilityZone string, region string, ) *CreateSubnetRequest {
	this := CreateSubnetRequest{}
	this.Name = name
	this.Cloud = cloud
	this.Cidr = cidr
	this.Provider = provider
	this.AvailabilityZone = availabilityZone
	this.Region = region
	return &this
}

// NewCreateSubnetRequestWithDefaults instantiates a new CreateSubnetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSubnetRequestWithDefaults() *CreateSubnetRequest {
	this := CreateSubnetRequest{}
	return &this
}

// GetName returns the Name field value
func (o *CreateSubnetRequest) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateSubnetRequest) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateSubnetRequest) SetName(v string) {
	o.Name = v
}

// GetCloud returns the Cloud field value
func (o *CreateSubnetRequest) GetCloud() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Cloud
}

// GetCloudOk returns a tuple with the Cloud field value
// and a boolean to check if the value has been set.
func (o *CreateSubnetRequest) GetCloudOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Cloud, true
}

// SetCloud sets field value
func (o *CreateSubnetRequest) SetCloud(v string) {
	o.Cloud = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *CreateSubnetRequest) GetTags() map[string]interface{} {
	if o == nil || o.Tags == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSubnetRequest) GetTagsOk() (*map[string]interface{}, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *CreateSubnetRequest) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]interface{} and assigns it to the Tags field.
func (o *CreateSubnetRequest) SetTags(v map[string]interface{}) {
	o.Tags = &v
}

// GetCidr returns the Cidr field value
func (o *CreateSubnetRequest) GetCidr() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Cidr
}

// GetCidrOk returns a tuple with the Cidr field value
// and a boolean to check if the value has been set.
func (o *CreateSubnetRequest) GetCidrOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Cidr, true
}

// SetCidr sets field value
func (o *CreateSubnetRequest) SetCidr(v string) {
	o.Cidr = v
}

// GetProvider returns the Provider field value
func (o *CreateSubnetRequest) GetProvider() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *CreateSubnetRequest) GetProviderOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value
func (o *CreateSubnetRequest) SetProvider(v string) {
	o.Provider = v
}

// GetAvailabilityZone returns the AvailabilityZone field value
func (o *CreateSubnetRequest) GetAvailabilityZone() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.AvailabilityZone
}

// GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field value
// and a boolean to check if the value has been set.
func (o *CreateSubnetRequest) GetAvailabilityZoneOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AvailabilityZone, true
}

// SetAvailabilityZone sets field value
func (o *CreateSubnetRequest) SetAvailabilityZone(v string) {
	o.AvailabilityZone = v
}

// GetRegion returns the Region field value
func (o *CreateSubnetRequest) GetRegion() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *CreateSubnetRequest) GetRegionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value
func (o *CreateSubnetRequest) SetRegion(v string) {
	o.Region = v
}

// GetGatewayIp returns the GatewayIp field value if set, zero value otherwise.
func (o *CreateSubnetRequest) GetGatewayIp() string {
	if o == nil || o.GatewayIp == nil {
		var ret string
		return ret
	}
	return *o.GatewayIp
}

// GetGatewayIpOk returns a tuple with the GatewayIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSubnetRequest) GetGatewayIpOk() (*string, bool) {
	if o == nil || o.GatewayIp == nil {
		return nil, false
	}
	return o.GatewayIp, true
}

// HasGatewayIp returns a boolean if a field has been set.
func (o *CreateSubnetRequest) HasGatewayIp() bool {
	if o != nil && o.GatewayIp != nil {
		return true
	}

	return false
}

// SetGatewayIp gets a reference to the given string and assigns it to the GatewayIp field.
func (o *CreateSubnetRequest) SetGatewayIp(v string) {
	o.GatewayIp = &v
}

// GetIpVersion returns the IpVersion field value if set, zero value otherwise.
func (o *CreateSubnetRequest) GetIpVersion() string {
	if o == nil || o.IpVersion == nil {
		var ret string
		return ret
	}
	return *o.IpVersion
}

// GetIpVersionOk returns a tuple with the IpVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSubnetRequest) GetIpVersionOk() (*string, bool) {
	if o == nil || o.IpVersion == nil {
		return nil, false
	}
	return o.IpVersion, true
}

// HasIpVersion returns a boolean if a field has been set.
func (o *CreateSubnetRequest) HasIpVersion() bool {
	if o != nil && o.IpVersion != nil {
		return true
	}

	return false
}

// SetIpVersion gets a reference to the given string and assigns it to the IpVersion field.
func (o *CreateSubnetRequest) SetIpVersion(v string) {
	o.IpVersion = &v
}

// GetEnableDhcp returns the EnableDhcp field value if set, zero value otherwise.
func (o *CreateSubnetRequest) GetEnableDhcp() bool {
	if o == nil || o.EnableDhcp == nil {
		var ret bool
		return ret
	}
	return *o.EnableDhcp
}

// GetEnableDhcpOk returns a tuple with the EnableDhcp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSubnetRequest) GetEnableDhcpOk() (*bool, bool) {
	if o == nil || o.EnableDhcp == nil {
		return nil, false
	}
	return o.EnableDhcp, true
}

// HasEnableDhcp returns a boolean if a field has been set.
func (o *CreateSubnetRequest) HasEnableDhcp() bool {
	if o != nil && o.EnableDhcp != nil {
		return true
	}

	return false
}

// SetEnableDhcp gets a reference to the given bool and assigns it to the EnableDhcp field.
func (o *CreateSubnetRequest) SetEnableDhcp(v bool) {
	o.EnableDhcp = &v
}

// GetAllocationPools returns the AllocationPools field value if set, zero value otherwise.
func (o *CreateSubnetRequest) GetAllocationPools() map[string]interface{} {
	if o == nil || o.AllocationPools == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.AllocationPools
}

// GetAllocationPoolsOk returns a tuple with the AllocationPools field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSubnetRequest) GetAllocationPoolsOk() (*map[string]interface{}, bool) {
	if o == nil || o.AllocationPools == nil {
		return nil, false
	}
	return o.AllocationPools, true
}

// HasAllocationPools returns a boolean if a field has been set.
func (o *CreateSubnetRequest) HasAllocationPools() bool {
	if o != nil && o.AllocationPools != nil {
		return true
	}

	return false
}

// SetAllocationPools gets a reference to the given map[string]interface{} and assigns it to the AllocationPools field.
func (o *CreateSubnetRequest) SetAllocationPools(v map[string]interface{}) {
	o.AllocationPools = &v
}

func (o CreateSubnetRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["cloud"] = o.Cloud
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if true {
		toSerialize["cidr"] = o.Cidr
	}
	if true {
		toSerialize["provider"] = o.Provider
	}
	if true {
		toSerialize["availability_zone"] = o.AvailabilityZone
	}
	if true {
		toSerialize["region"] = o.Region
	}
	if o.GatewayIp != nil {
		toSerialize["gateway_ip"] = o.GatewayIp
	}
	if o.IpVersion != nil {
		toSerialize["ip_version"] = o.IpVersion
	}
	if o.EnableDhcp != nil {
		toSerialize["enable_dhcp"] = o.EnableDhcp
	}
	if o.AllocationPools != nil {
		toSerialize["allocation_pools"] = o.AllocationPools
	}
	return json.Marshal(toSerialize)
}

type NullableCreateSubnetRequest struct {
	value *CreateSubnetRequest
	isSet bool
}

func (v NullableCreateSubnetRequest) Get() *CreateSubnetRequest {
	return v.value
}

func (v *NullableCreateSubnetRequest) Set(val *CreateSubnetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSubnetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSubnetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSubnetRequest(val *CreateSubnetRequest) *NullableCreateSubnetRequest {
	return &NullableCreateSubnetRequest{value: val, isSet: true}
}

func (v NullableCreateSubnetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateSubnetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


