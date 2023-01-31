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

// CreateMachineRequest struct for CreateMachineRequest
type CreateMachineRequest struct {
	// Specify machine name
	Name string `json:"name"`
	Provider *SupportedProviders `json:"provider,omitempty"`
	// Specify cloud to provision on
	Cloud *string `json:"cloud,omitempty"`
	// Where to provision e.g. region, datacenter, rack
	Location NullableOneOfobjectstring `json:"location,omitempty"`
	// Machine sizing spec e.g. cpu/ram/flavor
	Size NullableOneOfobjectstring `json:"size,omitempty"`
	// Operating System image to boot from
	Image NullableOneOfobjectstring `json:"image"`
	// Specify network configuration parameters
	Net NullableAnyOfAlibabaNetAmazonNetAzureNetDockerNetEquinixMetalNetGoogleNetKVMNetLinodeNetLXDNetOpenstackNetVSphereNetVultrNet `json:"net,omitempty"`
	// Associate SSH key
	Key NullableOneOfobjectstring `json:"key,omitempty"`
	Disks *CreateMachineRequestDisks `json:"disks,omitempty"`
	// Configure of attached storage volumes, e.g. cloud disks
	Volumes NullableAnyOfarrayarrayarrayarrayarrayarrayarrayarrayarrayarray `json:"volumes,omitempty"`
	// Add DNS A Record that points machine's public IP to this Fully Qualified Domain Name. Zone needs to be managed by a configured Cloud DNS provider
	Fqdn *string `json:"fqdn,omitempty"`
	// Run this Cloud Init script on first boot
	Cloudinit *string `json:"cloudinit,omitempty"`
	// Run post deploy scripts over SSH
	Scripts []ScriptToRun `json:"scripts,omitempty"`
	// Configure scheduled actions for the provisioned machine
	Schedules []AddScheduleRequest `json:"schedules,omitempty"`
	// Assign tags to provisioned machine
	Tags map[string]interface{} `json:"tags,omitempty"`
	Expiration *Expiration `json:"expiration,omitempty"`
	// Configure additional parameters
	Extra NullableAnyOfAzureExtraDockerExtraEquinixMetalExtraGoogleExtraLinodeExtraLXDExtraVSphereExtraVultrExtra `json:"extra,omitempty"`
	// Enable monitoring of this machine
	Monitoring *bool `json:"monitoring,omitempty"`
	// Provision multiple machines of this type
	Quantity *float32 `json:"quantity,omitempty"`
	Template map[string]interface{} `json:"template,omitempty"`
	// Return provisioning plan and exit without executing it
	Dry *bool `json:"dry,omitempty"`
	// Save provisioning plan as template
	Save *bool `json:"save,omitempty"`
	// Criteria optimization, e.g \"cost\"
	Optimize *string `json:"optimize,omitempty"`
}

// NewCreateMachineRequest instantiates a new CreateMachineRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateMachineRequest(name string, image NullableOneOfobjectstring) *CreateMachineRequest {
	this := CreateMachineRequest{}
	this.Name = name
	this.Image = image
	return &this
}

// NewCreateMachineRequestWithDefaults instantiates a new CreateMachineRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateMachineRequestWithDefaults() *CreateMachineRequest {
	this := CreateMachineRequest{}
	return &this
}

// GetName returns the Name field value
func (o *CreateMachineRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateMachineRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateMachineRequest) SetName(v string) {
	o.Name = v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *CreateMachineRequest) GetProvider() SupportedProviders {
	if o == nil || o.Provider == nil {
		var ret SupportedProviders
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMachineRequest) GetProviderOk() (*SupportedProviders, bool) {
	if o == nil || o.Provider == nil {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *CreateMachineRequest) HasProvider() bool {
	if o != nil && o.Provider != nil {
		return true
	}

	return false
}

// SetProvider gets a reference to the given SupportedProviders and assigns it to the Provider field.
func (o *CreateMachineRequest) SetProvider(v SupportedProviders) {
	o.Provider = &v
}

// GetCloud returns the Cloud field value if set, zero value otherwise.
func (o *CreateMachineRequest) GetCloud() string {
	if o == nil || o.Cloud == nil {
		var ret string
		return ret
	}
	return *o.Cloud
}

// GetCloudOk returns a tuple with the Cloud field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMachineRequest) GetCloudOk() (*string, bool) {
	if o == nil || o.Cloud == nil {
		return nil, false
	}
	return o.Cloud, true
}

// HasCloud returns a boolean if a field has been set.
func (o *CreateMachineRequest) HasCloud() bool {
	if o != nil && o.Cloud != nil {
		return true
	}

	return false
}

// SetCloud gets a reference to the given string and assigns it to the Cloud field.
func (o *CreateMachineRequest) SetCloud(v string) {
	o.Cloud = &v
}

// GetLocation returns the Location field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateMachineRequest) GetLocation() OneOfobjectstring {
	if o == nil || o.Location.Get() == nil {
		var ret OneOfobjectstring
		return ret
	}
	return *o.Location.Get()
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateMachineRequest) GetLocationOk() (*OneOfobjectstring, bool) {
	if o == nil {
		return nil, false
	}
	return o.Location.Get(), o.Location.IsSet()
}

// HasLocation returns a boolean if a field has been set.
func (o *CreateMachineRequest) HasLocation() bool {
	if o != nil && o.Location.IsSet() {
		return true
	}

	return false
}

// SetLocation gets a reference to the given NullableOneOfobjectstring and assigns it to the Location field.
func (o *CreateMachineRequest) SetLocation(v OneOfobjectstring) {
	o.Location.Set(&v)
}
// SetLocationNil sets the value for Location to be an explicit nil
func (o *CreateMachineRequest) SetLocationNil() {
	o.Location.Set(nil)
}

// UnsetLocation ensures that no value is present for Location, not even an explicit nil
func (o *CreateMachineRequest) UnsetLocation() {
	o.Location.Unset()
}

// GetSize returns the Size field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateMachineRequest) GetSize() OneOfobjectstring {
	if o == nil || o.Size.Get() == nil {
		var ret OneOfobjectstring
		return ret
	}
	return *o.Size.Get()
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateMachineRequest) GetSizeOk() (*OneOfobjectstring, bool) {
	if o == nil {
		return nil, false
	}
	return o.Size.Get(), o.Size.IsSet()
}

// HasSize returns a boolean if a field has been set.
func (o *CreateMachineRequest) HasSize() bool {
	if o != nil && o.Size.IsSet() {
		return true
	}

	return false
}

// SetSize gets a reference to the given NullableOneOfobjectstring and assigns it to the Size field.
func (o *CreateMachineRequest) SetSize(v OneOfobjectstring) {
	o.Size.Set(&v)
}
// SetSizeNil sets the value for Size to be an explicit nil
func (o *CreateMachineRequest) SetSizeNil() {
	o.Size.Set(nil)
}

// UnsetSize ensures that no value is present for Size, not even an explicit nil
func (o *CreateMachineRequest) UnsetSize() {
	o.Size.Unset()
}

// GetImage returns the Image field value
// If the value is explicit nil, the zero value for OneOfobjectstring will be returned
func (o *CreateMachineRequest) GetImage() OneOfobjectstring {
	if o == nil || o.Image.Get() == nil {
		var ret OneOfobjectstring
		return ret
	}

	return *o.Image.Get()
}

// GetImageOk returns a tuple with the Image field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateMachineRequest) GetImageOk() (*OneOfobjectstring, bool) {
	if o == nil {
		return nil, false
	}
	return o.Image.Get(), o.Image.IsSet()
}

// SetImage sets field value
func (o *CreateMachineRequest) SetImage(v OneOfobjectstring) {
	o.Image.Set(&v)
}

// GetNet returns the Net field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateMachineRequest) GetNet() AnyOfAlibabaNetAmazonNetAzureNetDockerNetEquinixMetalNetGoogleNetKVMNetLinodeNetLXDNetOpenstackNetVSphereNetVultrNet {
	if o == nil || o.Net.Get() == nil {
		var ret AnyOfAlibabaNetAmazonNetAzureNetDockerNetEquinixMetalNetGoogleNetKVMNetLinodeNetLXDNetOpenstackNetVSphereNetVultrNet
		return ret
	}
	return *o.Net.Get()
}

// GetNetOk returns a tuple with the Net field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateMachineRequest) GetNetOk() (*AnyOfAlibabaNetAmazonNetAzureNetDockerNetEquinixMetalNetGoogleNetKVMNetLinodeNetLXDNetOpenstackNetVSphereNetVultrNet, bool) {
	if o == nil {
		return nil, false
	}
	return o.Net.Get(), o.Net.IsSet()
}

// HasNet returns a boolean if a field has been set.
func (o *CreateMachineRequest) HasNet() bool {
	if o != nil && o.Net.IsSet() {
		return true
	}

	return false
}

// SetNet gets a reference to the given NullableAnyOfAlibabaNetAmazonNetAzureNetDockerNetEquinixMetalNetGoogleNetKVMNetLinodeNetLXDNetOpenstackNetVSphereNetVultrNet and assigns it to the Net field.
func (o *CreateMachineRequest) SetNet(v AnyOfAlibabaNetAmazonNetAzureNetDockerNetEquinixMetalNetGoogleNetKVMNetLinodeNetLXDNetOpenstackNetVSphereNetVultrNet) {
	o.Net.Set(&v)
}
// SetNetNil sets the value for Net to be an explicit nil
func (o *CreateMachineRequest) SetNetNil() {
	o.Net.Set(nil)
}

// UnsetNet ensures that no value is present for Net, not even an explicit nil
func (o *CreateMachineRequest) UnsetNet() {
	o.Net.Unset()
}

// GetKey returns the Key field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateMachineRequest) GetKey() OneOfobjectstring {
	if o == nil || o.Key.Get() == nil {
		var ret OneOfobjectstring
		return ret
	}
	return *o.Key.Get()
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateMachineRequest) GetKeyOk() (*OneOfobjectstring, bool) {
	if o == nil {
		return nil, false
	}
	return o.Key.Get(), o.Key.IsSet()
}

// HasKey returns a boolean if a field has been set.
func (o *CreateMachineRequest) HasKey() bool {
	if o != nil && o.Key.IsSet() {
		return true
	}

	return false
}

// SetKey gets a reference to the given NullableOneOfobjectstring and assigns it to the Key field.
func (o *CreateMachineRequest) SetKey(v OneOfobjectstring) {
	o.Key.Set(&v)
}
// SetKeyNil sets the value for Key to be an explicit nil
func (o *CreateMachineRequest) SetKeyNil() {
	o.Key.Set(nil)
}

// UnsetKey ensures that no value is present for Key, not even an explicit nil
func (o *CreateMachineRequest) UnsetKey() {
	o.Key.Unset()
}

// GetDisks returns the Disks field value if set, zero value otherwise.
func (o *CreateMachineRequest) GetDisks() CreateMachineRequestDisks {
	if o == nil || o.Disks == nil {
		var ret CreateMachineRequestDisks
		return ret
	}
	return *o.Disks
}

// GetDisksOk returns a tuple with the Disks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMachineRequest) GetDisksOk() (*CreateMachineRequestDisks, bool) {
	if o == nil || o.Disks == nil {
		return nil, false
	}
	return o.Disks, true
}

// HasDisks returns a boolean if a field has been set.
func (o *CreateMachineRequest) HasDisks() bool {
	if o != nil && o.Disks != nil {
		return true
	}

	return false
}

// SetDisks gets a reference to the given CreateMachineRequestDisks and assigns it to the Disks field.
func (o *CreateMachineRequest) SetDisks(v CreateMachineRequestDisks) {
	o.Disks = &v
}

// GetVolumes returns the Volumes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateMachineRequest) GetVolumes() AnyOfarrayarrayarrayarrayarrayarrayarrayarrayarrayarray {
	if o == nil || o.Volumes.Get() == nil {
		var ret AnyOfarrayarrayarrayarrayarrayarrayarrayarrayarrayarray
		return ret
	}
	return *o.Volumes.Get()
}

// GetVolumesOk returns a tuple with the Volumes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateMachineRequest) GetVolumesOk() (*AnyOfarrayarrayarrayarrayarrayarrayarrayarrayarrayarray, bool) {
	if o == nil {
		return nil, false
	}
	return o.Volumes.Get(), o.Volumes.IsSet()
}

// HasVolumes returns a boolean if a field has been set.
func (o *CreateMachineRequest) HasVolumes() bool {
	if o != nil && o.Volumes.IsSet() {
		return true
	}

	return false
}

// SetVolumes gets a reference to the given NullableAnyOfarrayarrayarrayarrayarrayarrayarrayarrayarrayarray and assigns it to the Volumes field.
func (o *CreateMachineRequest) SetVolumes(v AnyOfarrayarrayarrayarrayarrayarrayarrayarrayarrayarray) {
	o.Volumes.Set(&v)
}
// SetVolumesNil sets the value for Volumes to be an explicit nil
func (o *CreateMachineRequest) SetVolumesNil() {
	o.Volumes.Set(nil)
}

// UnsetVolumes ensures that no value is present for Volumes, not even an explicit nil
func (o *CreateMachineRequest) UnsetVolumes() {
	o.Volumes.Unset()
}

// GetFqdn returns the Fqdn field value if set, zero value otherwise.
func (o *CreateMachineRequest) GetFqdn() string {
	if o == nil || o.Fqdn == nil {
		var ret string
		return ret
	}
	return *o.Fqdn
}

// GetFqdnOk returns a tuple with the Fqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMachineRequest) GetFqdnOk() (*string, bool) {
	if o == nil || o.Fqdn == nil {
		return nil, false
	}
	return o.Fqdn, true
}

// HasFqdn returns a boolean if a field has been set.
func (o *CreateMachineRequest) HasFqdn() bool {
	if o != nil && o.Fqdn != nil {
		return true
	}

	return false
}

// SetFqdn gets a reference to the given string and assigns it to the Fqdn field.
func (o *CreateMachineRequest) SetFqdn(v string) {
	o.Fqdn = &v
}

// GetCloudinit returns the Cloudinit field value if set, zero value otherwise.
func (o *CreateMachineRequest) GetCloudinit() string {
	if o == nil || o.Cloudinit == nil {
		var ret string
		return ret
	}
	return *o.Cloudinit
}

// GetCloudinitOk returns a tuple with the Cloudinit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMachineRequest) GetCloudinitOk() (*string, bool) {
	if o == nil || o.Cloudinit == nil {
		return nil, false
	}
	return o.Cloudinit, true
}

// HasCloudinit returns a boolean if a field has been set.
func (o *CreateMachineRequest) HasCloudinit() bool {
	if o != nil && o.Cloudinit != nil {
		return true
	}

	return false
}

// SetCloudinit gets a reference to the given string and assigns it to the Cloudinit field.
func (o *CreateMachineRequest) SetCloudinit(v string) {
	o.Cloudinit = &v
}

// GetScripts returns the Scripts field value if set, zero value otherwise.
func (o *CreateMachineRequest) GetScripts() []ScriptToRun {
	if o == nil || o.Scripts == nil {
		var ret []ScriptToRun
		return ret
	}
	return o.Scripts
}

// GetScriptsOk returns a tuple with the Scripts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMachineRequest) GetScriptsOk() ([]ScriptToRun, bool) {
	if o == nil || o.Scripts == nil {
		return nil, false
	}
	return o.Scripts, true
}

// HasScripts returns a boolean if a field has been set.
func (o *CreateMachineRequest) HasScripts() bool {
	if o != nil && o.Scripts != nil {
		return true
	}

	return false
}

// SetScripts gets a reference to the given []ScriptToRun and assigns it to the Scripts field.
func (o *CreateMachineRequest) SetScripts(v []ScriptToRun) {
	o.Scripts = v
}

// GetSchedules returns the Schedules field value if set, zero value otherwise.
func (o *CreateMachineRequest) GetSchedules() []AddScheduleRequest {
	if o == nil || o.Schedules == nil {
		var ret []AddScheduleRequest
		return ret
	}
	return o.Schedules
}

// GetSchedulesOk returns a tuple with the Schedules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMachineRequest) GetSchedulesOk() ([]AddScheduleRequest, bool) {
	if o == nil || o.Schedules == nil {
		return nil, false
	}
	return o.Schedules, true
}

// HasSchedules returns a boolean if a field has been set.
func (o *CreateMachineRequest) HasSchedules() bool {
	if o != nil && o.Schedules != nil {
		return true
	}

	return false
}

// SetSchedules gets a reference to the given []AddScheduleRequest and assigns it to the Schedules field.
func (o *CreateMachineRequest) SetSchedules(v []AddScheduleRequest) {
	o.Schedules = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *CreateMachineRequest) GetTags() map[string]interface{} {
	if o == nil || o.Tags == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMachineRequest) GetTagsOk() (map[string]interface{}, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *CreateMachineRequest) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]interface{} and assigns it to the Tags field.
func (o *CreateMachineRequest) SetTags(v map[string]interface{}) {
	o.Tags = v
}

// GetExpiration returns the Expiration field value if set, zero value otherwise.
func (o *CreateMachineRequest) GetExpiration() Expiration {
	if o == nil || o.Expiration == nil {
		var ret Expiration
		return ret
	}
	return *o.Expiration
}

// GetExpirationOk returns a tuple with the Expiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMachineRequest) GetExpirationOk() (*Expiration, bool) {
	if o == nil || o.Expiration == nil {
		return nil, false
	}
	return o.Expiration, true
}

// HasExpiration returns a boolean if a field has been set.
func (o *CreateMachineRequest) HasExpiration() bool {
	if o != nil && o.Expiration != nil {
		return true
	}

	return false
}

// SetExpiration gets a reference to the given Expiration and assigns it to the Expiration field.
func (o *CreateMachineRequest) SetExpiration(v Expiration) {
	o.Expiration = &v
}

// GetExtra returns the Extra field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateMachineRequest) GetExtra() AnyOfAzureExtraDockerExtraEquinixMetalExtraGoogleExtraLinodeExtraLXDExtraVSphereExtraVultrExtra {
	if o == nil || o.Extra.Get() == nil {
		var ret AnyOfAzureExtraDockerExtraEquinixMetalExtraGoogleExtraLinodeExtraLXDExtraVSphereExtraVultrExtra
		return ret
	}
	return *o.Extra.Get()
}

// GetExtraOk returns a tuple with the Extra field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateMachineRequest) GetExtraOk() (*AnyOfAzureExtraDockerExtraEquinixMetalExtraGoogleExtraLinodeExtraLXDExtraVSphereExtraVultrExtra, bool) {
	if o == nil {
		return nil, false
	}
	return o.Extra.Get(), o.Extra.IsSet()
}

// HasExtra returns a boolean if a field has been set.
func (o *CreateMachineRequest) HasExtra() bool {
	if o != nil && o.Extra.IsSet() {
		return true
	}

	return false
}

// SetExtra gets a reference to the given NullableAnyOfAzureExtraDockerExtraEquinixMetalExtraGoogleExtraLinodeExtraLXDExtraVSphereExtraVultrExtra and assigns it to the Extra field.
func (o *CreateMachineRequest) SetExtra(v AnyOfAzureExtraDockerExtraEquinixMetalExtraGoogleExtraLinodeExtraLXDExtraVSphereExtraVultrExtra) {
	o.Extra.Set(&v)
}
// SetExtraNil sets the value for Extra to be an explicit nil
func (o *CreateMachineRequest) SetExtraNil() {
	o.Extra.Set(nil)
}

// UnsetExtra ensures that no value is present for Extra, not even an explicit nil
func (o *CreateMachineRequest) UnsetExtra() {
	o.Extra.Unset()
}

// GetMonitoring returns the Monitoring field value if set, zero value otherwise.
func (o *CreateMachineRequest) GetMonitoring() bool {
	if o == nil || o.Monitoring == nil {
		var ret bool
		return ret
	}
	return *o.Monitoring
}

// GetMonitoringOk returns a tuple with the Monitoring field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMachineRequest) GetMonitoringOk() (*bool, bool) {
	if o == nil || o.Monitoring == nil {
		return nil, false
	}
	return o.Monitoring, true
}

// HasMonitoring returns a boolean if a field has been set.
func (o *CreateMachineRequest) HasMonitoring() bool {
	if o != nil && o.Monitoring != nil {
		return true
	}

	return false
}

// SetMonitoring gets a reference to the given bool and assigns it to the Monitoring field.
func (o *CreateMachineRequest) SetMonitoring(v bool) {
	o.Monitoring = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *CreateMachineRequest) GetQuantity() float32 {
	if o == nil || o.Quantity == nil {
		var ret float32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMachineRequest) GetQuantityOk() (*float32, bool) {
	if o == nil || o.Quantity == nil {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *CreateMachineRequest) HasQuantity() bool {
	if o != nil && o.Quantity != nil {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given float32 and assigns it to the Quantity field.
func (o *CreateMachineRequest) SetQuantity(v float32) {
	o.Quantity = &v
}

// GetTemplate returns the Template field value if set, zero value otherwise.
func (o *CreateMachineRequest) GetTemplate() map[string]interface{} {
	if o == nil || o.Template == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Template
}

// GetTemplateOk returns a tuple with the Template field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMachineRequest) GetTemplateOk() (map[string]interface{}, bool) {
	if o == nil || o.Template == nil {
		return nil, false
	}
	return o.Template, true
}

// HasTemplate returns a boolean if a field has been set.
func (o *CreateMachineRequest) HasTemplate() bool {
	if o != nil && o.Template != nil {
		return true
	}

	return false
}

// SetTemplate gets a reference to the given map[string]interface{} and assigns it to the Template field.
func (o *CreateMachineRequest) SetTemplate(v map[string]interface{}) {
	o.Template = v
}

// GetDry returns the Dry field value if set, zero value otherwise.
func (o *CreateMachineRequest) GetDry() bool {
	if o == nil || o.Dry == nil {
		var ret bool
		return ret
	}
	return *o.Dry
}

// GetDryOk returns a tuple with the Dry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMachineRequest) GetDryOk() (*bool, bool) {
	if o == nil || o.Dry == nil {
		return nil, false
	}
	return o.Dry, true
}

// HasDry returns a boolean if a field has been set.
func (o *CreateMachineRequest) HasDry() bool {
	if o != nil && o.Dry != nil {
		return true
	}

	return false
}

// SetDry gets a reference to the given bool and assigns it to the Dry field.
func (o *CreateMachineRequest) SetDry(v bool) {
	o.Dry = &v
}

// GetSave returns the Save field value if set, zero value otherwise.
func (o *CreateMachineRequest) GetSave() bool {
	if o == nil || o.Save == nil {
		var ret bool
		return ret
	}
	return *o.Save
}

// GetSaveOk returns a tuple with the Save field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMachineRequest) GetSaveOk() (*bool, bool) {
	if o == nil || o.Save == nil {
		return nil, false
	}
	return o.Save, true
}

// HasSave returns a boolean if a field has been set.
func (o *CreateMachineRequest) HasSave() bool {
	if o != nil && o.Save != nil {
		return true
	}

	return false
}

// SetSave gets a reference to the given bool and assigns it to the Save field.
func (o *CreateMachineRequest) SetSave(v bool) {
	o.Save = &v
}

// GetOptimize returns the Optimize field value if set, zero value otherwise.
func (o *CreateMachineRequest) GetOptimize() string {
	if o == nil || o.Optimize == nil {
		var ret string
		return ret
	}
	return *o.Optimize
}

// GetOptimizeOk returns a tuple with the Optimize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMachineRequest) GetOptimizeOk() (*string, bool) {
	if o == nil || o.Optimize == nil {
		return nil, false
	}
	return o.Optimize, true
}

// HasOptimize returns a boolean if a field has been set.
func (o *CreateMachineRequest) HasOptimize() bool {
	if o != nil && o.Optimize != nil {
		return true
	}

	return false
}

// SetOptimize gets a reference to the given string and assigns it to the Optimize field.
func (o *CreateMachineRequest) SetOptimize(v string) {
	o.Optimize = &v
}

func (o CreateMachineRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Provider != nil {
		toSerialize["provider"] = o.Provider
	}
	if o.Cloud != nil {
		toSerialize["cloud"] = o.Cloud
	}
	if o.Location.IsSet() {
		toSerialize["location"] = o.Location.Get()
	}
	if o.Size.IsSet() {
		toSerialize["size"] = o.Size.Get()
	}
	if true {
		toSerialize["image"] = o.Image.Get()
	}
	if o.Net.IsSet() {
		toSerialize["net"] = o.Net.Get()
	}
	if o.Key.IsSet() {
		toSerialize["key"] = o.Key.Get()
	}
	if o.Disks != nil {
		toSerialize["disks"] = o.Disks
	}
	if o.Volumes.IsSet() {
		toSerialize["volumes"] = o.Volumes.Get()
	}
	if o.Fqdn != nil {
		toSerialize["fqdn"] = o.Fqdn
	}
	if o.Cloudinit != nil {
		toSerialize["cloudinit"] = o.Cloudinit
	}
	if o.Scripts != nil {
		toSerialize["scripts"] = o.Scripts
	}
	if o.Schedules != nil {
		toSerialize["schedules"] = o.Schedules
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Expiration != nil {
		toSerialize["expiration"] = o.Expiration
	}
	if o.Extra.IsSet() {
		toSerialize["extra"] = o.Extra.Get()
	}
	if o.Monitoring != nil {
		toSerialize["monitoring"] = o.Monitoring
	}
	if o.Quantity != nil {
		toSerialize["quantity"] = o.Quantity
	}
	if o.Template != nil {
		toSerialize["template"] = o.Template
	}
	if o.Dry != nil {
		toSerialize["dry"] = o.Dry
	}
	if o.Save != nil {
		toSerialize["save"] = o.Save
	}
	if o.Optimize != nil {
		toSerialize["optimize"] = o.Optimize
	}
	return json.Marshal(toSerialize)
}

type NullableCreateMachineRequest struct {
	value *CreateMachineRequest
	isSet bool
}

func (v NullableCreateMachineRequest) Get() *CreateMachineRequest {
	return v.value
}

func (v *NullableCreateMachineRequest) Set(val *CreateMachineRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateMachineRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateMachineRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateMachineRequest(val *CreateMachineRequest) *NullableCreateMachineRequest {
	return &NullableCreateMachineRequest{value: val, isSet: true}
}

func (v NullableCreateMachineRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateMachineRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


