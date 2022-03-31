# Go API client for mist_sdk

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 2.0.0
- Package version: 0.0.1
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import sw "./mist_sdk"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identifield by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```
ctx := context.WithValue(context.Background(), sw.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), sw.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://mist.io*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*CloudsApi* | [**AddCloud**](docs/CloudsApi.md#addcloud) | **Post** /api/v2/clouds | Add cloud
*CloudsApi* | [**EditCloud**](docs/CloudsApi.md#editcloud) | **Put** /api/v2/clouds/{cloud} | Edit cloud
*CloudsApi* | [**GetCloud**](docs/CloudsApi.md#getcloud) | **Get** /api/v2/clouds/{cloud} | Get cloud
*CloudsApi* | [**ListClouds**](docs/CloudsApi.md#listclouds) | **Get** /api/v2/clouds | List clouds
*CloudsApi* | [**RemoveCloud**](docs/CloudsApi.md#removecloud) | **Delete** /api/v2/clouds/{cloud} | Remove cloud
*ClustersApi* | [**CreateCluster**](docs/ClustersApi.md#createcluster) | **Post** /api/v2/clusters | Create cluster
*ClustersApi* | [**DestroyCluster**](docs/ClustersApi.md#destroycluster) | **Delete** /api/v2/clusters/{cluster} | Destroy cluster
*ClustersApi* | [**GetCluster**](docs/ClustersApi.md#getcluster) | **Get** /api/v2/clusters/{cluster} | Get cluster
*ClustersApi* | [**ListClusters**](docs/ClustersApi.md#listclusters) | **Get** /api/v2/clusters | List clusters
*DatapointsApi* | [**GetDatapoints**](docs/DatapointsApi.md#getdatapoints) | **Get** /api/v2/datapoints | Get datapoints
*ImagesApi* | [**GetImage**](docs/ImagesApi.md#getimage) | **Get** /api/v2/images/{image} | Get image
*ImagesApi* | [**ListImages**](docs/ImagesApi.md#listimages) | **Get** /api/v2/images | List images
*JobsApi* | [**GetJob**](docs/JobsApi.md#getjob) | **Get** /api/v2/jobs/{job_id} | Get job
*KeysApi* | [**AddKey**](docs/KeysApi.md#addkey) | **Post** /api/v2/keys | Add key
*KeysApi* | [**DeleteKey**](docs/KeysApi.md#deletekey) | **Delete** /api/v2/keys/{key} | Delete key
*KeysApi* | [**EditKey**](docs/KeysApi.md#editkey) | **Put** /api/v2/keys/{key} | Edit key
*KeysApi* | [**GetKey**](docs/KeysApi.md#getkey) | **Get** /api/v2/keys/{key} | Get key
*KeysApi* | [**ListKeys**](docs/KeysApi.md#listkeys) | **Get** /api/v2/keys | List keys
*LocationsApi* | [**GetLocation**](docs/LocationsApi.md#getlocation) | **Get** /api/v2/locations/{location} | Get location
*LocationsApi* | [**ListLocations**](docs/LocationsApi.md#listlocations) | **Get** /api/v2/locations | List locations
*MachinesApi* | [**AssociateKey**](docs/MachinesApi.md#associatekey) | **Put** /api/v2/machines/{machine}/actions/associate-key | Associate a key with a machine
*MachinesApi* | [**CloneMachine**](docs/MachinesApi.md#clonemachine) | **Post** /api/v2/machines/{machine}/actions/clone | Clone machine
*MachinesApi* | [**Console**](docs/MachinesApi.md#console) | **Post** /api/v2/machines/{machine}/actions/console | Open console
*MachinesApi* | [**CreateMachine**](docs/MachinesApi.md#createmachine) | **Post** /api/v2/machines | Create machine
*MachinesApi* | [**DestroyMachine**](docs/MachinesApi.md#destroymachine) | **Post** /api/v2/machines/{machine}/actions/destroy | Destroy machine
*MachinesApi* | [**DisassociateKey**](docs/MachinesApi.md#disassociatekey) | **Delete** /api/v2/machines/{machine}/actions/disassociate-key | Disassociate a key from a machine
*MachinesApi* | [**EditMachine**](docs/MachinesApi.md#editmachine) | **Put** /api/v2/machines/{machine} | Edit machine
*MachinesApi* | [**GetMachine**](docs/MachinesApi.md#getmachine) | **Get** /api/v2/machines/{machine} | Get machine
*MachinesApi* | [**ListMachines**](docs/MachinesApi.md#listmachines) | **Get** /api/v2/machines | List machines
*MachinesApi* | [**RebootMachine**](docs/MachinesApi.md#rebootmachine) | **Post** /api/v2/machines/{machine}/actions/reboot | Reboot machine
*MachinesApi* | [**RenameMachine**](docs/MachinesApi.md#renamemachine) | **Post** /api/v2/machines/{machine}/actions/rename | Rename machine
*MachinesApi* | [**ResizeMachine**](docs/MachinesApi.md#resizemachine) | **Post** /api/v2/machines/{machine}/actions/resize | Resize machine
*MachinesApi* | [**ResumeMachine**](docs/MachinesApi.md#resumemachine) | **Post** /api/v2/machines/{machine}/actions/resume | Resume machine
*MachinesApi* | [**Ssh**](docs/MachinesApi.md#ssh) | **Post** /api/v2/machines/{machine}/actions/ssh | Open secure shell
*MachinesApi* | [**StartMachine**](docs/MachinesApi.md#startmachine) | **Post** /api/v2/machines/{machine}/actions/start | Start machine
*MachinesApi* | [**StopMachine**](docs/MachinesApi.md#stopmachine) | **Post** /api/v2/machines/{machine}/actions/stop | Stop machine
*MachinesApi* | [**SuspendMachine**](docs/MachinesApi.md#suspendmachine) | **Post** /api/v2/machines/{machine}/actions/suspend | Suspend machine
*MachinesApi* | [**UndefineMachine**](docs/MachinesApi.md#undefinemachine) | **Post** /api/v2/machines/{machine}/actions/undefine | Undefine machine
*NetworksApi* | [**CreateNetwork**](docs/NetworksApi.md#createnetwork) | **Post** /api/v2/networks | Create network
*NetworksApi* | [**DeleteNetwork**](docs/NetworksApi.md#deletenetwork) | **Delete** /api/v2/networks/{network} | Delete network
*NetworksApi* | [**EditNetwork**](docs/NetworksApi.md#editnetwork) | **Put** /api/v2/networks/{network} | Edit network
*NetworksApi* | [**GetNetwork**](docs/NetworksApi.md#getnetwork) | **Get** /api/v2/networks/{network} | Get network
*NetworksApi* | [**ListNetworks**](docs/NetworksApi.md#listnetworks) | **Get** /api/v2/networks | List networks
*OrgsApi* | [**GetMember**](docs/OrgsApi.md#getmember) | **Get** /api/v2/orgs/{org}/members/{member} | Get Org
*OrgsApi* | [**GetOrg**](docs/OrgsApi.md#getorg) | **Get** /api/v2/orgs/{org} | Get Org
*OrgsApi* | [**ListOrgMembers**](docs/OrgsApi.md#listorgmembers) | **Get** /api/v2/orgs/{org}/members | List org members
*OrgsApi* | [**ListOrgTeams**](docs/OrgsApi.md#listorgteams) | **Get** /api/v2/orgs/{org}/teams | List org teams
*OrgsApi* | [**ListOrgs**](docs/OrgsApi.md#listorgs) | **Get** /api/v2/orgs | List orgs
*RulesApi* | [**AddRule**](docs/RulesApi.md#addrule) | **Post** /api/v2/rules | Add rule
*RulesApi* | [**DeleteRule**](docs/RulesApi.md#deleterule) | **Delete** /api/v2/rules/{rule} | Delete rule
*RulesApi* | [**EditRule**](docs/RulesApi.md#editrule) | **Post** /api/v2/rules/{rule} | Update rule
*RulesApi* | [**GetRule**](docs/RulesApi.md#getrule) | **Get** /api/v2/rules/{rule} | Get rule
*RulesApi* | [**ListRules**](docs/RulesApi.md#listrules) | **Get** /api/v2/rules | List rules
*RulesApi* | [**RenameRule**](docs/RulesApi.md#renamerule) | **Patch** /api/v2/rules/{rule} | Rename rule
*RulesApi* | [**ToggleRule**](docs/RulesApi.md#togglerule) | **Put** /api/v2/rules/{rule} | Toggle rule
*ScriptsApi* | [**AddScript**](docs/ScriptsApi.md#addscript) | **Post** /api/v2/scripts | Add script
*ScriptsApi* | [**DeleteScript**](docs/ScriptsApi.md#deletescript) | **Delete** /api/v2/scripts/{script} | Delete script
*ScriptsApi* | [**DownloadScript**](docs/ScriptsApi.md#downloadscript) | **Get** /api/v2/scripts/{script}/file | Download script
*ScriptsApi* | [**EditScript**](docs/ScriptsApi.md#editscript) | **Put** /api/v2/scripts/{script} | Edit script
*ScriptsApi* | [**GenerateScriptUrl**](docs/ScriptsApi.md#generatescripturl) | **Get** /api/v2/scripts/{script}/url | Generate script url
*ScriptsApi* | [**GetScript**](docs/ScriptsApi.md#getscript) | **Get** /api/v2/scripts/{script} | Get script
*ScriptsApi* | [**ListScripts**](docs/ScriptsApi.md#listscripts) | **Get** /api/v2/scripts | List scripts
*ScriptsApi* | [**RunScript**](docs/ScriptsApi.md#runscript) | **Post** /api/v2/scripts/{script} | Run script
*SizesApi* | [**GetSize**](docs/SizesApi.md#getsize) | **Get** /api/v2/sizes/{size} | Get size
*SizesApi* | [**ListSizes**](docs/SizesApi.md#listsizes) | **Get** /api/v2/sizes | List sizes
*SnapshotsApi* | [**CreateSnapshot**](docs/SnapshotsApi.md#createsnapshot) | **Post** /api/v2/machines/{machine}/snapshots | Create snapshot
*SnapshotsApi* | [**ListSnapshots**](docs/SnapshotsApi.md#listsnapshots) | **Get** /api/v2/machines/{machine}/snapshots | List machine snapshots
*SnapshotsApi* | [**RemoveSnapshot**](docs/SnapshotsApi.md#removesnapshot) | **Delete** /api/v2/machines/{machine}/snapshots/{snapshot} | Remove snapshot
*SnapshotsApi* | [**RevertToSnapshot**](docs/SnapshotsApi.md#reverttosnapshot) | **Post** /api/v2/machines/{machine}/snapshots/{snapshot} | Revert to snapshot
*UsersApi* | [**ListUsers**](docs/UsersApi.md#listusers) | **Get** /api/v2/users | List users
*VolumesApi* | [**CreateVolume**](docs/VolumesApi.md#createvolume) | **Post** /api/v2/volumes | Create volume
*VolumesApi* | [**DeleteVolume**](docs/VolumesApi.md#deletevolume) | **Delete** /api/v2/volumes/{volume} | Delete volume
*VolumesApi* | [**EditVolume**](docs/VolumesApi.md#editvolume) | **Put** /api/v2/volumes/{volume} | Edit volume
*VolumesApi* | [**GetVolume**](docs/VolumesApi.md#getvolume) | **Get** /api/v2/volumes/{volume} | Get volume
*VolumesApi* | [**ListVolumes**](docs/VolumesApi.md#listvolumes) | **Get** /api/v2/volumes | List volumes
*ZonesApi* | [**CreateZone**](docs/ZonesApi.md#createzone) | **Post** /api/v2/zones | Create zone
*ZonesApi* | [**DeleteZone**](docs/ZonesApi.md#deletezone) | **Delete** /api/v2/zones/{zone} | Delete zone
*ZonesApi* | [**EditZone**](docs/ZonesApi.md#editzone) | **Put** /api/v2/zones/{zone} | Edit zone
*ZonesApi* | [**GetZone**](docs/ZonesApi.md#getzone) | **Get** /api/v2/zones/{zone} | Get zone
*ZonesApi* | [**ListZones**](docs/ZonesApi.md#listzones) | **Get** /api/v2/zones | List zones


## Documentation For Models

 - [AddCloudRequest](docs/AddCloudRequest.md)
 - [AddCloudRequestAllOf](docs/AddCloudRequestAllOf.md)
 - [AddKeyRequest](docs/AddKeyRequest.md)
 - [AddKeyRequestAnyOf](docs/AddKeyRequestAnyOf.md)
 - [AddKeyRequestAnyOf1](docs/AddKeyRequestAnyOf1.md)
 - [AddKeyRequestAnyOf2](docs/AddKeyRequestAnyOf2.md)
 - [AddKeyResponse](docs/AddKeyResponse.md)
 - [AddRuleRequest](docs/AddRuleRequest.md)
 - [AddScriptRequest](docs/AddScriptRequest.md)
 - [AlibabaCloudRequest](docs/AlibabaCloudRequest.md)
 - [AlibabaCredentials](docs/AlibabaCredentials.md)
 - [AlibabaNet](docs/AlibabaNet.md)
 - [AmazonCloudRequest](docs/AmazonCloudRequest.md)
 - [AmazonCredentials](docs/AmazonCredentials.md)
 - [AmazonNet](docs/AmazonNet.md)
 - [AzureCloudRequest](docs/AzureCloudRequest.md)
 - [AzureCredentials](docs/AzureCredentials.md)
 - [AzureExtra](docs/AzureExtra.md)
 - [AzureNet](docs/AzureNet.md)
 - [Cloud](docs/Cloud.md)
 - [CloudFeatures](docs/CloudFeatures.md)
 - [CloudSigmaCloudRequest](docs/CloudSigmaCloudRequest.md)
 - [CloudSigmaCredentials](docs/CloudSigmaCredentials.md)
 - [Cluster](docs/Cluster.md)
 - [ClusterProviders](docs/ClusterProviders.md)
 - [CreateClusterRequest](docs/CreateClusterRequest.md)
 - [CreateClusterRequestAllOf](docs/CreateClusterRequestAllOf.md)
 - [CreateClusterRequestAllOfHelmCharts](docs/CreateClusterRequestAllOfHelmCharts.md)
 - [CreateClusterRequestAllOfNodepools](docs/CreateClusterRequestAllOfNodepools.md)
 - [CreateClusterResponse](docs/CreateClusterResponse.md)
 - [CreateMachineRequest](docs/CreateMachineRequest.md)
 - [CreateMachineRequestDisks](docs/CreateMachineRequestDisks.md)
 - [CreateMachineResponse](docs/CreateMachineResponse.md)
 - [CreateMachineResponseOneOf](docs/CreateMachineResponseOneOf.md)
 - [CreateMachineResponseOneOf1](docs/CreateMachineResponseOneOf1.md)
 - [CreateNetworkRequest](docs/CreateNetworkRequest.md)
 - [CreateNetworkResponse](docs/CreateNetworkResponse.md)
 - [CreateVolumeRequest](docs/CreateVolumeRequest.md)
 - [CreateVolumeResponse](docs/CreateVolumeResponse.md)
 - [CreateZoneRequest](docs/CreateZoneRequest.md)
 - [CreateZoneResponse](docs/CreateZoneResponse.md)
 - [CronSchedule](docs/CronSchedule.md)
 - [DataType](docs/DataType.md)
 - [Datapoints](docs/Datapoints.md)
 - [DatapointsData](docs/DatapointsData.md)
 - [DatapointsValuesItem](docs/DatapointsValuesItem.md)
 - [DestroyClusterResponse](docs/DestroyClusterResponse.md)
 - [DigitaloceanCloudRequest](docs/DigitaloceanCloudRequest.md)
 - [DigitaloceanCredentials](docs/DigitaloceanCredentials.md)
 - [DockerCloudRequest](docs/DockerCloudRequest.md)
 - [DockerCredentials](docs/DockerCredentials.md)
 - [DockerExtra](docs/DockerExtra.md)
 - [DockerNet](docs/DockerNet.md)
 - [EditCloudRequest](docs/EditCloudRequest.md)
 - [EditCloudRequestAnyOf](docs/EditCloudRequestAnyOf.md)
 - [EditMachineRequest](docs/EditMachineRequest.md)
 - [EditMachineRequestExpiration](docs/EditMachineRequestExpiration.md)
 - [EditRuleRequest](docs/EditRuleRequest.md)
 - [EquinixCloudRequest](docs/EquinixCloudRequest.md)
 - [EquinixCredentials](docs/EquinixCredentials.md)
 - [EquinixMetalExtra](docs/EquinixMetalExtra.md)
 - [EquinixMetalNet](docs/EquinixMetalNet.md)
 - [EquinixMetalNetIpAddresses](docs/EquinixMetalNetIpAddresses.md)
 - [Expiration](docs/Expiration.md)
 - [ExpirationNotify](docs/ExpirationNotify.md)
 - [Frequency](docs/Frequency.md)
 - [GetCloudResponse](docs/GetCloudResponse.md)
 - [GetClusterResponse](docs/GetClusterResponse.md)
 - [GetDatapointsResponse](docs/GetDatapointsResponse.md)
 - [GetImageResponse](docs/GetImageResponse.md)
 - [GetJobResponse](docs/GetJobResponse.md)
 - [GetKeyResponse](docs/GetKeyResponse.md)
 - [GetLocationResponse](docs/GetLocationResponse.md)
 - [GetMachineResponse](docs/GetMachineResponse.md)
 - [GetNetworkResponse](docs/GetNetworkResponse.md)
 - [GetOrgMemberResponse](docs/GetOrgMemberResponse.md)
 - [GetOrgResponse](docs/GetOrgResponse.md)
 - [GetRuleResponse](docs/GetRuleResponse.md)
 - [GetScriptResponse](docs/GetScriptResponse.md)
 - [GetSizeResponse](docs/GetSizeResponse.md)
 - [GetVolumeResponse](docs/GetVolumeResponse.md)
 - [GetZoneResponse](docs/GetZoneResponse.md)
 - [GoogleCloudRequest](docs/GoogleCloudRequest.md)
 - [GoogleCredentials](docs/GoogleCredentials.md)
 - [GoogleExtra](docs/GoogleExtra.md)
 - [GoogleExtraServiceAccount](docs/GoogleExtraServiceAccount.md)
 - [GoogleNet](docs/GoogleNet.md)
 - [IbmCloudRequest](docs/IbmCloudRequest.md)
 - [IbmCredentials](docs/IbmCredentials.md)
 - [Image](docs/Image.md)
 - [InlineResponse200](docs/InlineResponse200.md)
 - [InlineResponse2001](docs/InlineResponse2001.md)
 - [InlineScript](docs/InlineScript.md)
 - [InstantVector](docs/InstantVector.md)
 - [IntervalSchedule](docs/IntervalSchedule.md)
 - [Job](docs/Job.md)
 - [KVMNet](docs/KVMNet.md)
 - [KVMNetNetworks](docs/KVMNetNetworks.md)
 - [Key](docs/Key.md)
 - [KeyMachineAssociation](docs/KeyMachineAssociation.md)
 - [KeyMachineDisassociation](docs/KeyMachineDisassociation.md)
 - [KubernetesCloudRequest](docs/KubernetesCloudRequest.md)
 - [KubernetesCredentials](docs/KubernetesCredentials.md)
 - [KubevirtCloudRequest](docs/KubevirtCloudRequest.md)
 - [KvmCloudRequest](docs/KvmCloudRequest.md)
 - [LXDExtra](docs/LXDExtra.md)
 - [LXDNet](docs/LXDNet.md)
 - [LinodeCloudRequest](docs/LinodeCloudRequest.md)
 - [LinodeCredentials](docs/LinodeCredentials.md)
 - [LinodeExtra](docs/LinodeExtra.md)
 - [LinodeNet](docs/LinodeNet.md)
 - [ListCloudsResponse](docs/ListCloudsResponse.md)
 - [ListClustersResponse](docs/ListClustersResponse.md)
 - [ListImagesResponse](docs/ListImagesResponse.md)
 - [ListKeysResponse](docs/ListKeysResponse.md)
 - [ListLocationsResponse](docs/ListLocationsResponse.md)
 - [ListMachinesResponse](docs/ListMachinesResponse.md)
 - [ListNetworksResponse](docs/ListNetworksResponse.md)
 - [ListOrgMembersResponse](docs/ListOrgMembersResponse.md)
 - [ListOrgTeamsResponse](docs/ListOrgTeamsResponse.md)
 - [ListOrgsResponse](docs/ListOrgsResponse.md)
 - [ListRulesResponse](docs/ListRulesResponse.md)
 - [ListScriptsResponse](docs/ListScriptsResponse.md)
 - [ListSizesResponse](docs/ListSizesResponse.md)
 - [ListSnapshotsResponse](docs/ListSnapshotsResponse.md)
 - [ListUsersResponse](docs/ListUsersResponse.md)
 - [ListVolumesResponse](docs/ListVolumesResponse.md)
 - [ListZonesResponse](docs/ListZonesResponse.md)
 - [Location](docs/Location.md)
 - [Log](docs/Log.md)
 - [LxdCloudRequest](docs/LxdCloudRequest.md)
 - [LxdCredentials](docs/LxdCredentials.md)
 - [Machine](docs/Machine.md)
 - [MachineState](docs/MachineState.md)
 - [MaxihostCloudRequest](docs/MaxihostCloudRequest.md)
 - [MaxihostCredentials](docs/MaxihostCredentials.md)
 - [Member](docs/Member.md)
 - [Network](docs/Network.md)
 - [OnappCloudRequest](docs/OnappCloudRequest.md)
 - [OnappCredentials](docs/OnappCredentials.md)
 - [OneOffSchedule](docs/OneOffSchedule.md)
 - [OpenshiftCloudRequest](docs/OpenshiftCloudRequest.md)
 - [OpenshiftCredentials](docs/OpenshiftCredentials.md)
 - [OpenstackCloudRequest](docs/OpenstackCloudRequest.md)
 - [OpenstackCredentials](docs/OpenstackCredentials.md)
 - [OpenstackNet](docs/OpenstackNet.md)
 - [Org](docs/Org.md)
 - [OtherCloudRequest](docs/OtherCloudRequest.md)
 - [PostDeployScript](docs/PostDeployScript.md)
 - [Query](docs/Query.md)
 - [RackspaceCloudRequest](docs/RackspaceCloudRequest.md)
 - [RackspaceCredentials](docs/RackspaceCredentials.md)
 - [RangeVector](docs/RangeVector.md)
 - [ResponseMetadata](docs/ResponseMetadata.md)
 - [Rule](docs/Rule.md)
 - [RuleAction](docs/RuleAction.md)
 - [RunScriptRequest](docs/RunScriptRequest.md)
 - [RunScriptResponse](docs/RunScriptResponse.md)
 - [Script](docs/Script.md)
 - [Selector](docs/Selector.md)
 - [Size](docs/Size.md)
 - [SupportedProviders](docs/SupportedProviders.md)
 - [Team](docs/Team.md)
 - [TriggerAfter](docs/TriggerAfter.md)
 - [User](docs/User.md)
 - [VSphereExtra](docs/VSphereExtra.md)
 - [VSphereNet](docs/VSphereNet.md)
 - [Vector](docs/Vector.md)
 - [Volume](docs/Volume.md)
 - [VsphereCloudRequest](docs/VsphereCloudRequest.md)
 - [VsphereCredentials](docs/VsphereCredentials.md)
 - [VultrCloudRequest](docs/VultrCloudRequest.md)
 - [VultrCredentials](docs/VultrCredentials.md)
 - [VultrExtra](docs/VultrExtra.md)
 - [VultrNet](docs/VultrNet.md)
 - [Window](docs/Window.md)
 - [Zone](docs/Zone.md)


## Documentation For Authorization



### ApiKeyAuth

- **Type**: API key
- **API key parameter name**: Authorization
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: Authorization and passed in as the auth context for each request.


### CookieAuth

- **Type**: API key
- **API key parameter name**: session.id
- **Location**: 

Note, each API key must be added to a map of `map[string]APIKey` where the key is: session.id and passed in as the auth context for each request.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

api@mist.io

