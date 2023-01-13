# VirtualHost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | When creating a new VirtualHost, this is the ID for the VirtualHost. If not specified, an ID will be automatically assigned. When updating an existing VirtualHost, this field is ignored and the ID is determined by the path parameter. | [optional] 
**Host** | **string** | (sortable) The host name for the Virtual Host. | 
**Port** | **int32** | (sortable) The integer port number for the Virtual Host. | 
**AgentResourceCacheTTL** | Pointer to **int32** | (sortable) Indicates the number of seconds the Agent can cache resources for this application. | [optional] 
**KeyPairId** | Pointer to **int32** | Key pair assigned to Virtual Host used by SNI, If no key pair is assigned to a virtual host, ENGINE HTTPS Listener key pair will be used. | [optional] 
**TrustedCertificateGroupId** | Pointer to **int32** | Trusted Certificate Group assigned to Virtual Host for client certificate authentication. | [optional] 

## Methods

### NewVirtualHost

`func NewVirtualHost(host string, port int32, ) *VirtualHost`

NewVirtualHost instantiates a new VirtualHost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualHostWithDefaults

`func NewVirtualHostWithDefaults() *VirtualHost`

NewVirtualHostWithDefaults instantiates a new VirtualHost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VirtualHost) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VirtualHost) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VirtualHost) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *VirtualHost) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHost

`func (o *VirtualHost) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *VirtualHost) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *VirtualHost) SetHost(v string)`

SetHost sets Host field to given value.


### GetPort

`func (o *VirtualHost) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *VirtualHost) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *VirtualHost) SetPort(v int32)`

SetPort sets Port field to given value.


### GetAgentResourceCacheTTL

`func (o *VirtualHost) GetAgentResourceCacheTTL() int32`

GetAgentResourceCacheTTL returns the AgentResourceCacheTTL field if non-nil, zero value otherwise.

### GetAgentResourceCacheTTLOk

`func (o *VirtualHost) GetAgentResourceCacheTTLOk() (*int32, bool)`

GetAgentResourceCacheTTLOk returns a tuple with the AgentResourceCacheTTL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentResourceCacheTTL

`func (o *VirtualHost) SetAgentResourceCacheTTL(v int32)`

SetAgentResourceCacheTTL sets AgentResourceCacheTTL field to given value.

### HasAgentResourceCacheTTL

`func (o *VirtualHost) HasAgentResourceCacheTTL() bool`

HasAgentResourceCacheTTL returns a boolean if a field has been set.

### GetKeyPairId

`func (o *VirtualHost) GetKeyPairId() int32`

GetKeyPairId returns the KeyPairId field if non-nil, zero value otherwise.

### GetKeyPairIdOk

`func (o *VirtualHost) GetKeyPairIdOk() (*int32, bool)`

GetKeyPairIdOk returns a tuple with the KeyPairId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyPairId

`func (o *VirtualHost) SetKeyPairId(v int32)`

SetKeyPairId sets KeyPairId field to given value.

### HasKeyPairId

`func (o *VirtualHost) HasKeyPairId() bool`

HasKeyPairId returns a boolean if a field has been set.

### GetTrustedCertificateGroupId

`func (o *VirtualHost) GetTrustedCertificateGroupId() int32`

GetTrustedCertificateGroupId returns the TrustedCertificateGroupId field if non-nil, zero value otherwise.

### GetTrustedCertificateGroupIdOk

`func (o *VirtualHost) GetTrustedCertificateGroupIdOk() (*int32, bool)`

GetTrustedCertificateGroupIdOk returns a tuple with the TrustedCertificateGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustedCertificateGroupId

`func (o *VirtualHost) SetTrustedCertificateGroupId(v int32)`

SetTrustedCertificateGroupId sets TrustedCertificateGroupId field to given value.

### HasTrustedCertificateGroupId

`func (o *VirtualHost) HasTrustedCertificateGroupId() bool`

HasTrustedCertificateGroupId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


