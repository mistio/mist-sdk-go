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

// CreateClusterRequest struct for CreateClusterRequest
type CreateClusterRequest struct {
	// The cluster's name
	Name string `json:"name"`
	// The cloud the cluster belongs to
	Cloud string `json:"cloud"`
	Provider string `json:"provider"`
	// The Amazon Resource Name (ARN) of the IAM role that provides permissions for the Kubernetes control plane to make calls to AWS API operations on your behalf
	RoleArn string `json:"role_arn"`
	// Name or ID of the network to be associated with the cluster. If not given the default network will be selected
	Network *string `json:"network,omitempty"`
	// IDs of the subnets to be associated with the cluster. At least 2 subnets in different availability zones are required, if not given the default subnets will be used
	Subnets *[]string `json:"subnets,omitempty"`
	// The security groups associated with the cross-account elastic network interfaces that are used to allow communication between your nodes and the Kubernetes control plane
	SecurityGroups *[]string `json:"security_groups,omitempty"`
	// The initial number of nodes to provision for the nodegroup. Defaults to 2
	DesiredNodes *float32 `json:"desired_nodes,omitempty"`
	// The Amazon Resource Name (ARN) of the IAM role to associate with the node group
	NodegroupRoleArn *string `json:"nodegroup_role_arn,omitempty"`
	// Name or ID of size to use for the nodes. If not provided, the t3.medium size will be used
	NodegroupSize *string `json:"nodegroup_size,omitempty"`
	// The disk size for the nodegroup. Defaults to 20 GBs
	NodegroupDiskSize *float32 `json:"nodegroup_disk_size,omitempty"`
	// The name of the location to create the cluster in
	Location string `json:"location"`
}

// NewCreateClusterRequest instantiates a new CreateClusterRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateClusterRequest(name string, cloud string, provider string, roleArn string, location string, ) *CreateClusterRequest {
	this := CreateClusterRequest{}
	this.Name = name
	this.Cloud = cloud
	this.Provider = provider
	this.RoleArn = roleArn
	this.Location = location
	return &this
}

// NewCreateClusterRequestWithDefaults instantiates a new CreateClusterRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateClusterRequestWithDefaults() *CreateClusterRequest {
	this := CreateClusterRequest{}
	return &this
}

// GetName returns the Name field value
func (o *CreateClusterRequest) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateClusterRequest) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateClusterRequest) SetName(v string) {
	o.Name = v
}

// GetCloud returns the Cloud field value
func (o *CreateClusterRequest) GetCloud() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Cloud
}

// GetCloudOk returns a tuple with the Cloud field value
// and a boolean to check if the value has been set.
func (o *CreateClusterRequest) GetCloudOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Cloud, true
}

// SetCloud sets field value
func (o *CreateClusterRequest) SetCloud(v string) {
	o.Cloud = v
}

// GetProvider returns the Provider field value
func (o *CreateClusterRequest) GetProvider() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *CreateClusterRequest) GetProviderOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value
func (o *CreateClusterRequest) SetProvider(v string) {
	o.Provider = v
}

// GetRoleArn returns the RoleArn field value
func (o *CreateClusterRequest) GetRoleArn() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.RoleArn
}

// GetRoleArnOk returns a tuple with the RoleArn field value
// and a boolean to check if the value has been set.
func (o *CreateClusterRequest) GetRoleArnOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RoleArn, true
}

// SetRoleArn sets field value
func (o *CreateClusterRequest) SetRoleArn(v string) {
	o.RoleArn = v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *CreateClusterRequest) GetNetwork() string {
	if o == nil || o.Network == nil {
		var ret string
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateClusterRequest) GetNetworkOk() (*string, bool) {
	if o == nil || o.Network == nil {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *CreateClusterRequest) HasNetwork() bool {
	if o != nil && o.Network != nil {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given string and assigns it to the Network field.
func (o *CreateClusterRequest) SetNetwork(v string) {
	o.Network = &v
}

// GetSubnets returns the Subnets field value if set, zero value otherwise.
func (o *CreateClusterRequest) GetSubnets() []string {
	if o == nil || o.Subnets == nil {
		var ret []string
		return ret
	}
	return *o.Subnets
}

// GetSubnetsOk returns a tuple with the Subnets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateClusterRequest) GetSubnetsOk() (*[]string, bool) {
	if o == nil || o.Subnets == nil {
		return nil, false
	}
	return o.Subnets, true
}

// HasSubnets returns a boolean if a field has been set.
func (o *CreateClusterRequest) HasSubnets() bool {
	if o != nil && o.Subnets != nil {
		return true
	}

	return false
}

// SetSubnets gets a reference to the given []string and assigns it to the Subnets field.
func (o *CreateClusterRequest) SetSubnets(v []string) {
	o.Subnets = &v
}

// GetSecurityGroups returns the SecurityGroups field value if set, zero value otherwise.
func (o *CreateClusterRequest) GetSecurityGroups() []string {
	if o == nil || o.SecurityGroups == nil {
		var ret []string
		return ret
	}
	return *o.SecurityGroups
}

// GetSecurityGroupsOk returns a tuple with the SecurityGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateClusterRequest) GetSecurityGroupsOk() (*[]string, bool) {
	if o == nil || o.SecurityGroups == nil {
		return nil, false
	}
	return o.SecurityGroups, true
}

// HasSecurityGroups returns a boolean if a field has been set.
func (o *CreateClusterRequest) HasSecurityGroups() bool {
	if o != nil && o.SecurityGroups != nil {
		return true
	}

	return false
}

// SetSecurityGroups gets a reference to the given []string and assigns it to the SecurityGroups field.
func (o *CreateClusterRequest) SetSecurityGroups(v []string) {
	o.SecurityGroups = &v
}

// GetDesiredNodes returns the DesiredNodes field value if set, zero value otherwise.
func (o *CreateClusterRequest) GetDesiredNodes() float32 {
	if o == nil || o.DesiredNodes == nil {
		var ret float32
		return ret
	}
	return *o.DesiredNodes
}

// GetDesiredNodesOk returns a tuple with the DesiredNodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateClusterRequest) GetDesiredNodesOk() (*float32, bool) {
	if o == nil || o.DesiredNodes == nil {
		return nil, false
	}
	return o.DesiredNodes, true
}

// HasDesiredNodes returns a boolean if a field has been set.
func (o *CreateClusterRequest) HasDesiredNodes() bool {
	if o != nil && o.DesiredNodes != nil {
		return true
	}

	return false
}

// SetDesiredNodes gets a reference to the given float32 and assigns it to the DesiredNodes field.
func (o *CreateClusterRequest) SetDesiredNodes(v float32) {
	o.DesiredNodes = &v
}

// GetNodegroupRoleArn returns the NodegroupRoleArn field value if set, zero value otherwise.
func (o *CreateClusterRequest) GetNodegroupRoleArn() string {
	if o == nil || o.NodegroupRoleArn == nil {
		var ret string
		return ret
	}
	return *o.NodegroupRoleArn
}

// GetNodegroupRoleArnOk returns a tuple with the NodegroupRoleArn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateClusterRequest) GetNodegroupRoleArnOk() (*string, bool) {
	if o == nil || o.NodegroupRoleArn == nil {
		return nil, false
	}
	return o.NodegroupRoleArn, true
}

// HasNodegroupRoleArn returns a boolean if a field has been set.
func (o *CreateClusterRequest) HasNodegroupRoleArn() bool {
	if o != nil && o.NodegroupRoleArn != nil {
		return true
	}

	return false
}

// SetNodegroupRoleArn gets a reference to the given string and assigns it to the NodegroupRoleArn field.
func (o *CreateClusterRequest) SetNodegroupRoleArn(v string) {
	o.NodegroupRoleArn = &v
}

// GetNodegroupSize returns the NodegroupSize field value if set, zero value otherwise.
func (o *CreateClusterRequest) GetNodegroupSize() string {
	if o == nil || o.NodegroupSize == nil {
		var ret string
		return ret
	}
	return *o.NodegroupSize
}

// GetNodegroupSizeOk returns a tuple with the NodegroupSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateClusterRequest) GetNodegroupSizeOk() (*string, bool) {
	if o == nil || o.NodegroupSize == nil {
		return nil, false
	}
	return o.NodegroupSize, true
}

// HasNodegroupSize returns a boolean if a field has been set.
func (o *CreateClusterRequest) HasNodegroupSize() bool {
	if o != nil && o.NodegroupSize != nil {
		return true
	}

	return false
}

// SetNodegroupSize gets a reference to the given string and assigns it to the NodegroupSize field.
func (o *CreateClusterRequest) SetNodegroupSize(v string) {
	o.NodegroupSize = &v
}

// GetNodegroupDiskSize returns the NodegroupDiskSize field value if set, zero value otherwise.
func (o *CreateClusterRequest) GetNodegroupDiskSize() float32 {
	if o == nil || o.NodegroupDiskSize == nil {
		var ret float32
		return ret
	}
	return *o.NodegroupDiskSize
}

// GetNodegroupDiskSizeOk returns a tuple with the NodegroupDiskSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateClusterRequest) GetNodegroupDiskSizeOk() (*float32, bool) {
	if o == nil || o.NodegroupDiskSize == nil {
		return nil, false
	}
	return o.NodegroupDiskSize, true
}

// HasNodegroupDiskSize returns a boolean if a field has been set.
func (o *CreateClusterRequest) HasNodegroupDiskSize() bool {
	if o != nil && o.NodegroupDiskSize != nil {
		return true
	}

	return false
}

// SetNodegroupDiskSize gets a reference to the given float32 and assigns it to the NodegroupDiskSize field.
func (o *CreateClusterRequest) SetNodegroupDiskSize(v float32) {
	o.NodegroupDiskSize = &v
}

// GetLocation returns the Location field value
func (o *CreateClusterRequest) GetLocation() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Location
}

// GetLocationOk returns a tuple with the Location field value
// and a boolean to check if the value has been set.
func (o *CreateClusterRequest) GetLocationOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Location, true
}

// SetLocation sets field value
func (o *CreateClusterRequest) SetLocation(v string) {
	o.Location = v
}

func (o CreateClusterRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["cloud"] = o.Cloud
	}
	if true {
		toSerialize["provider"] = o.Provider
	}
	if true {
		toSerialize["role_arn"] = o.RoleArn
	}
	if o.Network != nil {
		toSerialize["network"] = o.Network
	}
	if o.Subnets != nil {
		toSerialize["subnets"] = o.Subnets
	}
	if o.SecurityGroups != nil {
		toSerialize["security_groups"] = o.SecurityGroups
	}
	if o.DesiredNodes != nil {
		toSerialize["desired_nodes"] = o.DesiredNodes
	}
	if o.NodegroupRoleArn != nil {
		toSerialize["nodegroup_role_arn"] = o.NodegroupRoleArn
	}
	if o.NodegroupSize != nil {
		toSerialize["nodegroup_size"] = o.NodegroupSize
	}
	if o.NodegroupDiskSize != nil {
		toSerialize["nodegroup_disk_size"] = o.NodegroupDiskSize
	}
	if true {
		toSerialize["location"] = o.Location
	}
	return json.Marshal(toSerialize)
}

type NullableCreateClusterRequest struct {
	value *CreateClusterRequest
	isSet bool
}

func (v NullableCreateClusterRequest) Get() *CreateClusterRequest {
	return v.value
}

func (v *NullableCreateClusterRequest) Set(val *CreateClusterRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateClusterRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateClusterRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateClusterRequest(val *CreateClusterRequest) *NullableCreateClusterRequest {
	return &NullableCreateClusterRequest{value: val, isSet: true}
}

func (v NullableCreateClusterRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateClusterRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


