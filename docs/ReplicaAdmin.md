# ReplicaAdmin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID of the replica admin. This field is read-only. | [optional] 
**Name** | **string** | The name of the replica admin. | 
**Description** | Pointer to **string** | The description of the replica admin. | [optional] 
**ConfigReplicationEnabled** | Pointer to **bool** | Set to true when configuration replication for the replica admin is enabled. False when configuration replication is suspended. | [optional] 
**Keys** | Pointer to [**[]PublicKey**](PublicKey.md) | An array of public keys associated with the replica admin. | [optional] 
**HttpProxyId** | Pointer to **int32** | The ID of the proxy to use for HTTP requests or zero if none. | [optional] 
**HttpsProxyId** | Pointer to **int32** | The ID of the proxy to use for HTTPS requests or zero if none. | [optional] 
**HostPort** | **string** | The host and port of the replica admin. | 
**SelectedCertificateId** | Pointer to **int32** | The ID of the certificate the replica admin will use to contact PingAccess via SSL/TLS. | [optional] 
**CertificateHash** | Pointer to [**Hash**](Hash.md) |  | [optional] 

## Methods

### NewReplicaAdmin

`func NewReplicaAdmin(name string, hostPort string, ) *ReplicaAdmin`

NewReplicaAdmin instantiates a new ReplicaAdmin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReplicaAdminWithDefaults

`func NewReplicaAdminWithDefaults() *ReplicaAdmin`

NewReplicaAdminWithDefaults instantiates a new ReplicaAdmin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ReplicaAdmin) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReplicaAdmin) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReplicaAdmin) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ReplicaAdmin) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ReplicaAdmin) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ReplicaAdmin) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ReplicaAdmin) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ReplicaAdmin) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ReplicaAdmin) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ReplicaAdmin) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ReplicaAdmin) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetConfigReplicationEnabled

`func (o *ReplicaAdmin) GetConfigReplicationEnabled() bool`

GetConfigReplicationEnabled returns the ConfigReplicationEnabled field if non-nil, zero value otherwise.

### GetConfigReplicationEnabledOk

`func (o *ReplicaAdmin) GetConfigReplicationEnabledOk() (*bool, bool)`

GetConfigReplicationEnabledOk returns a tuple with the ConfigReplicationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigReplicationEnabled

`func (o *ReplicaAdmin) SetConfigReplicationEnabled(v bool)`

SetConfigReplicationEnabled sets ConfigReplicationEnabled field to given value.

### HasConfigReplicationEnabled

`func (o *ReplicaAdmin) HasConfigReplicationEnabled() bool`

HasConfigReplicationEnabled returns a boolean if a field has been set.

### GetKeys

`func (o *ReplicaAdmin) GetKeys() []PublicKey`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *ReplicaAdmin) GetKeysOk() (*[]PublicKey, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *ReplicaAdmin) SetKeys(v []PublicKey)`

SetKeys sets Keys field to given value.

### HasKeys

`func (o *ReplicaAdmin) HasKeys() bool`

HasKeys returns a boolean if a field has been set.

### GetHttpProxyId

`func (o *ReplicaAdmin) GetHttpProxyId() int32`

GetHttpProxyId returns the HttpProxyId field if non-nil, zero value otherwise.

### GetHttpProxyIdOk

`func (o *ReplicaAdmin) GetHttpProxyIdOk() (*int32, bool)`

GetHttpProxyIdOk returns a tuple with the HttpProxyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpProxyId

`func (o *ReplicaAdmin) SetHttpProxyId(v int32)`

SetHttpProxyId sets HttpProxyId field to given value.

### HasHttpProxyId

`func (o *ReplicaAdmin) HasHttpProxyId() bool`

HasHttpProxyId returns a boolean if a field has been set.

### GetHttpsProxyId

`func (o *ReplicaAdmin) GetHttpsProxyId() int32`

GetHttpsProxyId returns the HttpsProxyId field if non-nil, zero value otherwise.

### GetHttpsProxyIdOk

`func (o *ReplicaAdmin) GetHttpsProxyIdOk() (*int32, bool)`

GetHttpsProxyIdOk returns a tuple with the HttpsProxyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpsProxyId

`func (o *ReplicaAdmin) SetHttpsProxyId(v int32)`

SetHttpsProxyId sets HttpsProxyId field to given value.

### HasHttpsProxyId

`func (o *ReplicaAdmin) HasHttpsProxyId() bool`

HasHttpsProxyId returns a boolean if a field has been set.

### GetHostPort

`func (o *ReplicaAdmin) GetHostPort() string`

GetHostPort returns the HostPort field if non-nil, zero value otherwise.

### GetHostPortOk

`func (o *ReplicaAdmin) GetHostPortOk() (*string, bool)`

GetHostPortOk returns a tuple with the HostPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostPort

`func (o *ReplicaAdmin) SetHostPort(v string)`

SetHostPort sets HostPort field to given value.


### GetSelectedCertificateId

`func (o *ReplicaAdmin) GetSelectedCertificateId() int32`

GetSelectedCertificateId returns the SelectedCertificateId field if non-nil, zero value otherwise.

### GetSelectedCertificateIdOk

`func (o *ReplicaAdmin) GetSelectedCertificateIdOk() (*int32, bool)`

GetSelectedCertificateIdOk returns a tuple with the SelectedCertificateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedCertificateId

`func (o *ReplicaAdmin) SetSelectedCertificateId(v int32)`

SetSelectedCertificateId sets SelectedCertificateId field to given value.

### HasSelectedCertificateId

`func (o *ReplicaAdmin) HasSelectedCertificateId() bool`

HasSelectedCertificateId returns a boolean if a field has been set.

### GetCertificateHash

`func (o *ReplicaAdmin) GetCertificateHash() Hash`

GetCertificateHash returns the CertificateHash field if non-nil, zero value otherwise.

### GetCertificateHashOk

`func (o *ReplicaAdmin) GetCertificateHashOk() (*Hash, bool)`

GetCertificateHashOk returns a tuple with the CertificateHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateHash

`func (o *ReplicaAdmin) SetCertificateHash(v Hash)`

SetCertificateHash sets CertificateHash field to given value.

### HasCertificateHash

`func (o *ReplicaAdmin) HasCertificateHash() bool`

HasCertificateHash returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


