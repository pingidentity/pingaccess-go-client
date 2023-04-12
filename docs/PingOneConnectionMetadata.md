# PingOneConnectionMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreationDate** | **time.Time** | The date this credential was created in PingOne. | 
**CredentialId** | **string** | The ID of the PingOne credential. | 
**PingOneConnectionId** | **string** | The ID of the PingOne connection as defined by PingOne. | 
**PingOneManagementApiEndpoint** | **string** | The PingOne management API endpoint. | 
**PingOneAuthenticationApiEndpoint** | **string** | The PingOne authentication API endpoint. | 
**CredentialStatus** | [**PingOneCredentialStatus**](PingOneCredentialStatus.md) |  | 
**Environments** | [**[]PingOneEnvironment**](PingOneEnvironment.md) | A list of environments available to the PingOne Connection. | 
**Id** | **string** | The ID of the PingOne Connection as defined by PingAccess. | 
**EnvironmentId** | Pointer to **string** | The ID of the environment of the PingOne credential. | [optional] 
**Region** | Pointer to **string** | The region of the PingOne connection. | [optional] 
**OrganizationName** | Pointer to **string** | The name of the organization associated with this PingOne connection.  | [optional] 

## Methods

### NewPingOneConnectionMetadata

`func NewPingOneConnectionMetadata(creationDate time.Time, credentialId string, pingOneConnectionId string, pingOneManagementApiEndpoint string, pingOneAuthenticationApiEndpoint string, credentialStatus PingOneCredentialStatus, environments []PingOneEnvironment, id string, ) *PingOneConnectionMetadata`

NewPingOneConnectionMetadata instantiates a new PingOneConnectionMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPingOneConnectionMetadataWithDefaults

`func NewPingOneConnectionMetadataWithDefaults() *PingOneConnectionMetadata`

NewPingOneConnectionMetadataWithDefaults instantiates a new PingOneConnectionMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreationDate

`func (o *PingOneConnectionMetadata) GetCreationDate() time.Time`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *PingOneConnectionMetadata) GetCreationDateOk() (*time.Time, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *PingOneConnectionMetadata) SetCreationDate(v time.Time)`

SetCreationDate sets CreationDate field to given value.


### GetCredentialId

`func (o *PingOneConnectionMetadata) GetCredentialId() string`

GetCredentialId returns the CredentialId field if non-nil, zero value otherwise.

### GetCredentialIdOk

`func (o *PingOneConnectionMetadata) GetCredentialIdOk() (*string, bool)`

GetCredentialIdOk returns a tuple with the CredentialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialId

`func (o *PingOneConnectionMetadata) SetCredentialId(v string)`

SetCredentialId sets CredentialId field to given value.


### GetPingOneConnectionId

`func (o *PingOneConnectionMetadata) GetPingOneConnectionId() string`

GetPingOneConnectionId returns the PingOneConnectionId field if non-nil, zero value otherwise.

### GetPingOneConnectionIdOk

`func (o *PingOneConnectionMetadata) GetPingOneConnectionIdOk() (*string, bool)`

GetPingOneConnectionIdOk returns a tuple with the PingOneConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingOneConnectionId

`func (o *PingOneConnectionMetadata) SetPingOneConnectionId(v string)`

SetPingOneConnectionId sets PingOneConnectionId field to given value.


### GetPingOneManagementApiEndpoint

`func (o *PingOneConnectionMetadata) GetPingOneManagementApiEndpoint() string`

GetPingOneManagementApiEndpoint returns the PingOneManagementApiEndpoint field if non-nil, zero value otherwise.

### GetPingOneManagementApiEndpointOk

`func (o *PingOneConnectionMetadata) GetPingOneManagementApiEndpointOk() (*string, bool)`

GetPingOneManagementApiEndpointOk returns a tuple with the PingOneManagementApiEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingOneManagementApiEndpoint

`func (o *PingOneConnectionMetadata) SetPingOneManagementApiEndpoint(v string)`

SetPingOneManagementApiEndpoint sets PingOneManagementApiEndpoint field to given value.


### GetPingOneAuthenticationApiEndpoint

`func (o *PingOneConnectionMetadata) GetPingOneAuthenticationApiEndpoint() string`

GetPingOneAuthenticationApiEndpoint returns the PingOneAuthenticationApiEndpoint field if non-nil, zero value otherwise.

### GetPingOneAuthenticationApiEndpointOk

`func (o *PingOneConnectionMetadata) GetPingOneAuthenticationApiEndpointOk() (*string, bool)`

GetPingOneAuthenticationApiEndpointOk returns a tuple with the PingOneAuthenticationApiEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingOneAuthenticationApiEndpoint

`func (o *PingOneConnectionMetadata) SetPingOneAuthenticationApiEndpoint(v string)`

SetPingOneAuthenticationApiEndpoint sets PingOneAuthenticationApiEndpoint field to given value.


### GetCredentialStatus

`func (o *PingOneConnectionMetadata) GetCredentialStatus() PingOneCredentialStatus`

GetCredentialStatus returns the CredentialStatus field if non-nil, zero value otherwise.

### GetCredentialStatusOk

`func (o *PingOneConnectionMetadata) GetCredentialStatusOk() (*PingOneCredentialStatus, bool)`

GetCredentialStatusOk returns a tuple with the CredentialStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialStatus

`func (o *PingOneConnectionMetadata) SetCredentialStatus(v PingOneCredentialStatus)`

SetCredentialStatus sets CredentialStatus field to given value.


### GetEnvironments

`func (o *PingOneConnectionMetadata) GetEnvironments() []PingOneEnvironment`

GetEnvironments returns the Environments field if non-nil, zero value otherwise.

### GetEnvironmentsOk

`func (o *PingOneConnectionMetadata) GetEnvironmentsOk() (*[]PingOneEnvironment, bool)`

GetEnvironmentsOk returns a tuple with the Environments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironments

`func (o *PingOneConnectionMetadata) SetEnvironments(v []PingOneEnvironment)`

SetEnvironments sets Environments field to given value.


### GetId

`func (o *PingOneConnectionMetadata) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PingOneConnectionMetadata) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PingOneConnectionMetadata) SetId(v string)`

SetId sets Id field to given value.


### GetEnvironmentId

`func (o *PingOneConnectionMetadata) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *PingOneConnectionMetadata) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *PingOneConnectionMetadata) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *PingOneConnectionMetadata) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.

### GetRegion

`func (o *PingOneConnectionMetadata) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *PingOneConnectionMetadata) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *PingOneConnectionMetadata) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *PingOneConnectionMetadata) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetOrganizationName

`func (o *PingOneConnectionMetadata) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *PingOneConnectionMetadata) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *PingOneConnectionMetadata) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.

### HasOrganizationName

`func (o *PingOneConnectionMetadata) HasOrganizationName() bool`

HasOrganizationName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


