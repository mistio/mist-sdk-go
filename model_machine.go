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

// Machine struct for Machine
type Machine struct {
	Id *string `json:"id,omitempty"`
	ExternalId *string `json:"external_id,omitempty"`
	Name *string `json:"name,omitempty"`
	Cloud *string `json:"cloud,omitempty"`
	Tags *map[string]interface{} `json:"tags,omitempty"`
	CreatedBy *string `json:"created_by,omitempty"`
	OwnedBy *string `json:"owned_by,omitempty"`
	Default *bool `json:"default,omitempty"`
	State *MachineState `json:"state,omitempty"`
	Actions *map[string]interface{} `json:"actions,omitempty"`
	Cluster *string `json:"cluster,omitempty"`
	Cores *int32 `json:"cores,omitempty"`
	Cost *map[string]interface{} `json:"cost,omitempty"`
	Created *string `json:"created,omitempty"`
	Expiration *string `json:"expiration,omitempty"`
	Extra *map[string]interface{} `json:"extra,omitempty"`
	Hostname *string `json:"hostname,omitempty"`
	Image *string `json:"image,omitempty"`
	KeyAssociations *[]KeyMachineAssociation `json:"key_associations,omitempty"`
	LastSeen *string `json:"last_seen,omitempty"`
	Location *string `json:"location,omitempty"`
	MissingSince *string `json:"missing_since,omitempty"`
	Monitoring *string `json:"monitoring,omitempty"`
	Network *string `json:"network,omitempty"`
	OsType *string `json:"os_type,omitempty"`
	Parent *string `json:"parent,omitempty"`
	Ports *map[string]interface{} `json:"ports,omitempty"`
	PrivateIps *[]string `json:"private_ips,omitempty"`
	Probe *map[string]interface{} `json:"probe,omitempty"`
	PublicIps *[]string `json:"public_ips,omitempty"`
	Size *string `json:"size,omitempty"`
	Subnet *string `json:"subnet,omitempty"`
	Type *string `json:"type,omitempty"`
	UnreachableSince *string `json:"unreachable_since,omitempty"`
}

// NewMachine instantiates a new Machine object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMachine() *Machine {
	this := Machine{}
	return &this
}

// NewMachineWithDefaults instantiates a new Machine object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMachineWithDefaults() *Machine {
	this := Machine{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Machine) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Machine) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Machine) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Machine) SetId(v string) {
	o.Id = &v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *Machine) GetExternalId() string {
	if o == nil || o.ExternalId == nil {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Machine) GetExternalIdOk() (*string, bool) {
	if o == nil || o.ExternalId == nil {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *Machine) HasExternalId() bool {
	if o != nil && o.ExternalId != nil {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *Machine) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Machine) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Machine) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Machine) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Machine) SetName(v string) {
	o.Name = &v
}

// GetCloud returns the Cloud field value if set, zero value otherwise.
func (o *Machine) GetCloud() string {
	if o == nil || o.Cloud == nil {
		var ret string
		return ret
	}
	return *o.Cloud
}

// GetCloudOk returns a tuple with the Cloud field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Machine) GetCloudOk() (*string, bool) {
	if o == nil || o.Cloud == nil {
		return nil, false
	}
	return o.Cloud, true
}

// HasCloud returns a boolean if a field has been set.
func (o *Machine) HasCloud() bool {
	if o != nil && o.Cloud != nil {
		return true
	}

	return false
}

// SetCloud gets a reference to the given string and assigns it to the Cloud field.
func (o *Machine) SetCloud(v string) {
	o.Cloud = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Machine) GetTags() map[string]interface{} {
	if o == nil || o.Tags == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Machine) GetTagsOk() (*map[string]interface{}, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Machine) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]interface{} and assigns it to the Tags field.
func (o *Machine) SetTags(v map[string]interface{}) {
	o.Tags = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *Machine) GetCreatedBy() string {
	if o == nil || o.CreatedBy == nil {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Machine) GetCreatedByOk() (*string, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *Machine) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *Machine) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetOwnedBy returns the OwnedBy field value if set, zero value otherwise.
func (o *Machine) GetOwnedBy() string {
	if o == nil || o.OwnedBy == nil {
		var ret string
		return ret
	}
	return *o.OwnedBy
}

// GetOwnedByOk returns a tuple with the OwnedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Machine) GetOwnedByOk() (*string, bool) {
	if o == nil || o.OwnedBy == nil {
		return nil, false
	}
	return o.OwnedBy, true
}

// HasOwnedBy returns a boolean if a field has been set.
func (o *Machine) HasOwnedBy() bool {
	if o != nil && o.OwnedBy != nil {
		return true
	}

	return false
}

// SetOwnedBy gets a reference to the given string and assigns it to the OwnedBy field.
func (o *Machine) SetOwnedBy(v string) {
	o.OwnedBy = &v
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *Machine) GetDefault() bool {
	if o == nil || o.Default == nil {
		var ret bool
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Machine) GetDefaultOk() (*bool, bool) {
	if o == nil || o.Default == nil {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *Machine) HasDefault() bool {
	if o != nil && o.Default != nil {
		return true
	}

	return false
}

// SetDefault gets a reference to the given bool and assigns it to the Default field.
func (o *Machine) SetDefault(v bool) {
	o.Default = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *Machine) GetState() MachineState {
	if o == nil || o.State == nil {
		var ret MachineState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Machine) GetStateOk() (*MachineState, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *Machine) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given MachineState and assigns it to the State field.
func (o *Machine) SetState(v MachineState) {
	o.State = &v
}

// GetActions returns the Actions field value if set, zero value otherwise.
func (o *Machine) GetActions() map[string]interface{} {
	if o == nil || o.Actions == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Machine) GetActionsOk() (*map[string]interface{}, bool) {
	if o == nil || o.Actions == nil {
		return nil, false
	}
	return o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *Machine) HasActions() bool {
	if o != nil && o.Actions != nil {
		return true
	}

	return false
}

// SetActions gets a reference to the given map[string]interface{} and assigns it to the Actions field.
func (o *Machine) SetActions(v map[string]interface{}) {
	o.Actions = &v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *Machine) GetCluster() string {
	if o == nil || o.Cluster == nil {
		var ret string
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Machine) GetClusterOk() (*string, bool) {
	if o == nil || o.Cluster == nil {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *Machine) HasCluster() bool {
	if o != nil && o.Cluster != nil {
		return true
	}

	return false
}

// SetCluster gets a reference to the given string and assigns it to the Cluster field.
func (o *Machine) SetCluster(v string) {
	o.Cluster = &v
}

// GetCores returns the Cores field value if set, zero value otherwise.
func (o *Machine) GetCores() int32 {
	if o == nil || o.Cores == nil {
		var ret int32
		return ret
	}
	return *o.Cores
}

// GetCoresOk returns a tuple with the Cores field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Machine) GetCoresOk() (*int32, bool) {
	if o == nil || o.Cores == nil {
		return nil, false
	}
	return o.Cores, true
}

// HasCores returns a boolean if a field has been set.
func (o *Machine) HasCores() bool {
	if o != nil && o.Cores != nil {
		return true
	}

	return false
}

// SetCores gets a reference to the given int32 and assigns it to the Cores field.
func (o *Machine) SetCores(v int32) {
	o.Cores = &v
}

// GetCost returns the Cost field value if set, zero value otherwise.
func (o *Machine) GetCost() map[string]interface{} {
	if o == nil || o.Cost == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Cost
}

// GetCostOk returns a tuple with the Cost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Machine) GetCostOk() (*map[string]interface{}, bool) {
	if o == nil || o.Cost == nil {
		return nil, false
	}
	return o.Cost, true
}

// HasCost returns a boolean if a field has been set.
func (o *Machine) HasCost() bool {
	if o != nil && o.Cost != nil {
		return true
	}

	return false
}

// SetCost gets a reference to the given map[string]interface{} and assigns it to the Cost field.
func (o *Machine) SetCost(v map[string]interface{}) {
	o.Cost = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *Machine) GetCreated() string {
	if o == nil || o.Created == nil {
		var ret string
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Machine) GetCreatedOk() (*string, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *Machine) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given string and assigns it to the Created field.
func (o *Machine) SetCreated(v string) {
	o.Created = &v
}

// GetExpiration returns the Expiration field value if set, zero value otherwise.
func (o *Machine) GetExpiration() string {
	if o == nil || o.Expiration == nil {
		var ret string
		return ret
	}
	return *o.Expiration
}

// GetExpirationOk returns a tuple with the Expiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Machine) GetExpirationOk() (*string, bool) {
	if o == nil || o.Expiration == nil {
		return nil, false
	}
	return o.Expiration, true
}

// HasExpiration returns a boolean if a field has been set.
func (o *Machine) HasExpiration() bool {
	if o != nil && o.Expiration != nil {
		return true
	}

	return false
}

// SetExpiration gets a reference to the given string and assigns it to the Expiration field.
func (o *Machine) SetExpiration(v string) {
	o.Expiration = &v
}

// GetExtra returns the Extra field value if set, zero value otherwise.
func (o *Machine) GetExtra() map[string]interface{} {
	if o == nil || o.Extra == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Extra
}

// GetExtraOk returns a tuple with the Extra field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Machine) GetExtraOk() (*map[string]interface{}, bool) {
	if o == nil || o.Extra == nil {
		return nil, false
	}
	return o.Extra, true
}

// HasExtra returns a boolean if a field has been set.
func (o *Machine) HasExtra() bool {
	if o != nil && o.Extra != nil {
		return true
	}

	return false
}

// SetExtra gets a reference to the given map[string]interface{} and assigns it to the Extra field.
func (o *Machine) SetExtra(v map[string]interface{}) {
	o.Extra = &v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *Machine) GetHostname() string {
	if o == nil || o.Hostname == nil {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Machine) GetHostnameOk() (*string, bool) {
	if o == nil || o.Hostname == nil {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *Machine) HasHostname() bool {
	if o != nil && o.Hostname != nil {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *Machine) SetHostname(v string) {
	o.Hostname = &v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *Machine) GetImage() string {
	if o == nil || o.Image == nil {
		var ret string
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Machine) GetImageOk() (*string, bool) {
	if o == nil || o.Image == nil {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *Machine) HasImage() bool {
	if o != nil && o.Image != nil {
		return true
	}

	return false
}

// SetImage gets a reference to the given string and assigns it to the Image field.
func (o *Machine) SetImage(v string) {
	o.Image = &v
}

// GetKeyAssociations returns the KeyAssociations field value if set, zero value otherwise.
func (o *Machine) GetKeyAssociations() []KeyMachineAssociation {
	if o == nil || o.KeyAssociations == nil {
		var ret []KeyMachineAssociation
		return ret
	}
	return *o.KeyAssociations
}

// GetKeyAssociationsOk returns a tuple with the KeyAssociations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Machine) GetKeyAssociationsOk() (*[]KeyMachineAssociation, bool) {
	if o == nil || o.KeyAssociations == nil {
		return nil, false
	}
	return o.KeyAssociations, true
}

// HasKeyAssociations returns a boolean if a field has been set.
func (o *Machine) HasKeyAssociations() bool {
	if o != nil && o.KeyAssociations != nil {
		return true
	}

	return false
}

// SetKeyAssociations gets a reference to the given []KeyMachineAssociation and assigns it to the KeyAssociations field.
func (o *Machine) SetKeyAssociations(v []KeyMachineAssociation) {
	o.KeyAssociations = &v
}

// GetLastSeen returns the LastSeen field value if set, zero value otherwise.
func (o *Machine) GetLastSeen() string {
	if o == nil || o.LastSeen == nil {
		var ret string
		return ret
	}
	return *o.LastSeen
}

// GetLastSeenOk returns a tuple with the LastSeen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Machine) GetLastSeenOk() (*string, bool) {
	if o == nil || o.LastSeen == nil {
		return nil, false
	}
	return o.LastSeen, true
}

// HasLastSeen returns a boolean if a field has been set.
func (o *Machine) HasLastSeen() bool {
	if o != nil && o.LastSeen != nil {
		return true
	}

	return false
}

// SetLastSeen gets a reference to the given string and assigns it to the LastSeen field.
func (o *Machine) SetLastSeen(v string) {
	o.LastSeen = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *Machine) GetLocation() string {
	if o == nil || o.Location == nil {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Machine) GetLocationOk() (*string, bool) {
	if o == nil || o.Location == nil {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *Machine) HasLocation() bool {
	if o != nil && o.Location != nil {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *Machine) SetLocation(v string) {
	o.Location = &v
}

// GetMissingSince returns the MissingSince field value if set, zero value otherwise.
func (o *Machine) GetMissingSince() string {
	if o == nil || o.MissingSince == nil {
		var ret string
		return ret
	}
	return *o.MissingSince
}

// GetMissingSinceOk returns a tuple with the MissingSince field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Machine) GetMissingSinceOk() (*string, bool) {
	if o == nil || o.MissingSince == nil {
		return nil, false
	}
	return o.MissingSince, true
}

// HasMissingSince returns a boolean if a field has been set.
func (o *Machine) HasMissingSince() bool {
	if o != nil && o.MissingSince != nil {
		return true
	}

	return false
}

// SetMissingSince gets a reference to the given string and assigns it to the MissingSince field.
func (o *Machine) SetMissingSince(v string) {
	o.MissingSince = &v
}

// GetMonitoring returns the Monitoring field value if set, zero value otherwise.
func (o *Machine) GetMonitoring() string {
	if o == nil || o.Monitoring == nil {
		var ret string
		return ret
	}
	return *o.Monitoring
}

// GetMonitoringOk returns a tuple with the Monitoring field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Machine) GetMonitoringOk() (*string, bool) {
	if o == nil || o.Monitoring == nil {
		return nil, false
	}
	return o.Monitoring, true
}

// HasMonitoring returns a boolean if a field has been set.
func (o *Machine) HasMonitoring() bool {
	if o != nil && o.Monitoring != nil {
		return true
	}

	return false
}

// SetMonitoring gets a reference to the given string and assigns it to the Monitoring field.
func (o *Machine) SetMonitoring(v string) {
	o.Monitoring = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *Machine) GetNetwork() string {
	if o == nil || o.Network == nil {
		var ret string
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Machine) GetNetworkOk() (*string, bool) {
	if o == nil || o.Network == nil {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *Machine) HasNetwork() bool {
	if o != nil && o.Network != nil {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given string and assigns it to the Network field.
func (o *Machine) SetNetwork(v string) {
	o.Network = &v
}

// GetOsType returns the OsType field value if set, zero value otherwise.
func (o *Machine) GetOsType() string {
	if o == nil || o.OsType == nil {
		var ret string
		return ret
	}
	return *o.OsType
}

// GetOsTypeOk returns a tuple with the OsType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Machine) GetOsTypeOk() (*string, bool) {
	if o == nil || o.OsType == nil {
		return nil, false
	}
	return o.OsType, true
}

// HasOsType returns a boolean if a field has been set.
func (o *Machine) HasOsType() bool {
	if o != nil && o.OsType != nil {
		return true
	}

	return false
}

// SetOsType gets a reference to the given string and assigns it to the OsType field.
func (o *Machine) SetOsType(v string) {
	o.OsType = &v
}

// GetParent returns the Parent field value if set, zero value otherwise.
func (o *Machine) GetParent() string {
	if o == nil || o.Parent == nil {
		var ret string
		return ret
	}
	return *o.Parent
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Machine) GetParentOk() (*string, bool) {
	if o == nil || o.Parent == nil {
		return nil, false
	}
	return o.Parent, true
}

// HasParent returns a boolean if a field has been set.
func (o *Machine) HasParent() bool {
	if o != nil && o.Parent != nil {
		return true
	}

	return false
}

// SetParent gets a reference to the given string and assigns it to the Parent field.
func (o *Machine) SetParent(v string) {
	o.Parent = &v
}

// GetPorts returns the Ports field value if set, zero value otherwise.
func (o *Machine) GetPorts() map[string]interface{} {
	if o == nil || o.Ports == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Ports
}

// GetPortsOk returns a tuple with the Ports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Machine) GetPortsOk() (*map[string]interface{}, bool) {
	if o == nil || o.Ports == nil {
		return nil, false
	}
	return o.Ports, true
}

// HasPorts returns a boolean if a field has been set.
func (o *Machine) HasPorts() bool {
	if o != nil && o.Ports != nil {
		return true
	}

	return false
}

// SetPorts gets a reference to the given map[string]interface{} and assigns it to the Ports field.
func (o *Machine) SetPorts(v map[string]interface{}) {
	o.Ports = &v
}

// GetPrivateIps returns the PrivateIps field value if set, zero value otherwise.
func (o *Machine) GetPrivateIps() []string {
	if o == nil || o.PrivateIps == nil {
		var ret []string
		return ret
	}
	return *o.PrivateIps
}

// GetPrivateIpsOk returns a tuple with the PrivateIps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Machine) GetPrivateIpsOk() (*[]string, bool) {
	if o == nil || o.PrivateIps == nil {
		return nil, false
	}
	return o.PrivateIps, true
}

// HasPrivateIps returns a boolean if a field has been set.
func (o *Machine) HasPrivateIps() bool {
	if o != nil && o.PrivateIps != nil {
		return true
	}

	return false
}

// SetPrivateIps gets a reference to the given []string and assigns it to the PrivateIps field.
func (o *Machine) SetPrivateIps(v []string) {
	o.PrivateIps = &v
}

// GetProbe returns the Probe field value if set, zero value otherwise.
func (o *Machine) GetProbe() map[string]interface{} {
	if o == nil || o.Probe == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Probe
}

// GetProbeOk returns a tuple with the Probe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Machine) GetProbeOk() (*map[string]interface{}, bool) {
	if o == nil || o.Probe == nil {
		return nil, false
	}
	return o.Probe, true
}

// HasProbe returns a boolean if a field has been set.
func (o *Machine) HasProbe() bool {
	if o != nil && o.Probe != nil {
		return true
	}

	return false
}

// SetProbe gets a reference to the given map[string]interface{} and assigns it to the Probe field.
func (o *Machine) SetProbe(v map[string]interface{}) {
	o.Probe = &v
}

// GetPublicIps returns the PublicIps field value if set, zero value otherwise.
func (o *Machine) GetPublicIps() []string {
	if o == nil || o.PublicIps == nil {
		var ret []string
		return ret
	}
	return *o.PublicIps
}

// GetPublicIpsOk returns a tuple with the PublicIps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Machine) GetPublicIpsOk() (*[]string, bool) {
	if o == nil || o.PublicIps == nil {
		return nil, false
	}
	return o.PublicIps, true
}

// HasPublicIps returns a boolean if a field has been set.
func (o *Machine) HasPublicIps() bool {
	if o != nil && o.PublicIps != nil {
		return true
	}

	return false
}

// SetPublicIps gets a reference to the given []string and assigns it to the PublicIps field.
func (o *Machine) SetPublicIps(v []string) {
	o.PublicIps = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *Machine) GetSize() string {
	if o == nil || o.Size == nil {
		var ret string
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Machine) GetSizeOk() (*string, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *Machine) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given string and assigns it to the Size field.
func (o *Machine) SetSize(v string) {
	o.Size = &v
}

// GetSubnet returns the Subnet field value if set, zero value otherwise.
func (o *Machine) GetSubnet() string {
	if o == nil || o.Subnet == nil {
		var ret string
		return ret
	}
	return *o.Subnet
}

// GetSubnetOk returns a tuple with the Subnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Machine) GetSubnetOk() (*string, bool) {
	if o == nil || o.Subnet == nil {
		return nil, false
	}
	return o.Subnet, true
}

// HasSubnet returns a boolean if a field has been set.
func (o *Machine) HasSubnet() bool {
	if o != nil && o.Subnet != nil {
		return true
	}

	return false
}

// SetSubnet gets a reference to the given string and assigns it to the Subnet field.
func (o *Machine) SetSubnet(v string) {
	o.Subnet = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Machine) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Machine) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Machine) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Machine) SetType(v string) {
	o.Type = &v
}

// GetUnreachableSince returns the UnreachableSince field value if set, zero value otherwise.
func (o *Machine) GetUnreachableSince() string {
	if o == nil || o.UnreachableSince == nil {
		var ret string
		return ret
	}
	return *o.UnreachableSince
}

// GetUnreachableSinceOk returns a tuple with the UnreachableSince field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Machine) GetUnreachableSinceOk() (*string, bool) {
	if o == nil || o.UnreachableSince == nil {
		return nil, false
	}
	return o.UnreachableSince, true
}

// HasUnreachableSince returns a boolean if a field has been set.
func (o *Machine) HasUnreachableSince() bool {
	if o != nil && o.UnreachableSince != nil {
		return true
	}

	return false
}

// SetUnreachableSince gets a reference to the given string and assigns it to the UnreachableSince field.
func (o *Machine) SetUnreachableSince(v string) {
	o.UnreachableSince = &v
}

func (o Machine) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.ExternalId != nil {
		toSerialize["external_id"] = o.ExternalId
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Cloud != nil {
		toSerialize["cloud"] = o.Cloud
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.CreatedBy != nil {
		toSerialize["created_by"] = o.CreatedBy
	}
	if o.OwnedBy != nil {
		toSerialize["owned_by"] = o.OwnedBy
	}
	if o.Default != nil {
		toSerialize["default"] = o.Default
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.Actions != nil {
		toSerialize["actions"] = o.Actions
	}
	if o.Cluster != nil {
		toSerialize["cluster"] = o.Cluster
	}
	if o.Cores != nil {
		toSerialize["cores"] = o.Cores
	}
	if o.Cost != nil {
		toSerialize["cost"] = o.Cost
	}
	if o.Created != nil {
		toSerialize["created"] = o.Created
	}
	if o.Expiration != nil {
		toSerialize["expiration"] = o.Expiration
	}
	if o.Extra != nil {
		toSerialize["extra"] = o.Extra
	}
	if o.Hostname != nil {
		toSerialize["hostname"] = o.Hostname
	}
	if o.Image != nil {
		toSerialize["image"] = o.Image
	}
	if o.KeyAssociations != nil {
		toSerialize["key_associations"] = o.KeyAssociations
	}
	if o.LastSeen != nil {
		toSerialize["last_seen"] = o.LastSeen
	}
	if o.Location != nil {
		toSerialize["location"] = o.Location
	}
	if o.MissingSince != nil {
		toSerialize["missing_since"] = o.MissingSince
	}
	if o.Monitoring != nil {
		toSerialize["monitoring"] = o.Monitoring
	}
	if o.Network != nil {
		toSerialize["network"] = o.Network
	}
	if o.OsType != nil {
		toSerialize["os_type"] = o.OsType
	}
	if o.Parent != nil {
		toSerialize["parent"] = o.Parent
	}
	if o.Ports != nil {
		toSerialize["ports"] = o.Ports
	}
	if o.PrivateIps != nil {
		toSerialize["private_ips"] = o.PrivateIps
	}
	if o.Probe != nil {
		toSerialize["probe"] = o.Probe
	}
	if o.PublicIps != nil {
		toSerialize["public_ips"] = o.PublicIps
	}
	if o.Size != nil {
		toSerialize["size"] = o.Size
	}
	if o.Subnet != nil {
		toSerialize["subnet"] = o.Subnet
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.UnreachableSince != nil {
		toSerialize["unreachable_since"] = o.UnreachableSince
	}
	return json.Marshal(toSerialize)
}

type NullableMachine struct {
	value *Machine
	isSet bool
}

func (v NullableMachine) Get() *Machine {
	return v.value
}

func (v *NullableMachine) Set(val *Machine) {
	v.value = val
	v.isSet = true
}

func (v NullableMachine) IsSet() bool {
	return v.isSet
}

func (v *NullableMachine) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMachine(val *Machine) *NullableMachine {
	return &NullableMachine{value: val, isSet: true}
}

func (v NullableMachine) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMachine) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


