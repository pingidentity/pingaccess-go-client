# Engine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | When creating a new Engine, this is the ID for the Engine. If not specified, an ID will be automatically assigned. When updating an existing Engine, this field is ignored and the ID is determined by the path parameter. | [optional] 
**Name** | **string** | (sortable) The name of the engine. | 
**Description** | Pointer to **string** | (sortable) The description of the engine. | [optional] 
**ConfigReplicationEnabled** | Pointer to **bool** | (sortable) Set to true when configuration replication for this engine is enabled. False when configuration replication is suspended. | [optional] 
**Keys** | Pointer to [**[]PublicKey**](PublicKey.md) | An array of public keys associated with the engine. | [optional] 
**HttpProxyId** | Pointer to **int64** | The ID of the proxy to use for HTTP requests or zero if none. | [optional] 
**HttpsProxyId** | Pointer to **int64** | The ID of the proxy to use for HTTPS requests or zero if none. | [optional] 
**SelectedCertificateId** | Pointer to **int64** | The ID of the certificate the engine will use to contact PingAccess via SSL/TLS. | [optional] 
**CertificateHash** | Pointer to [**Hash**](Hash.md) |  | [optional] 

## Methods

### NewEngine

`func NewEngine(name string, ) *Engine`

NewEngine instantiates a new Engine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEngineWithDefaults

`func NewEngineWithDefaults() *Engine`

NewEngineWithDefaults instantiates a new Engine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Engine) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Engine) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Engine) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Engine) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Engine) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Engine) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Engine) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Engine) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Engine) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Engine) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Engine) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetConfigReplicationEnabled

`func (o *Engine) GetConfigReplicationEnabled() bool`

GetConfigReplicationEnabled returns the ConfigReplicationEnabled field if non-nil, zero value otherwise.

### GetConfigReplicationEnabledOk

`func (o *Engine) GetConfigReplicationEnabledOk() (*bool, bool)`

GetConfigReplicationEnabledOk returns a tuple with the ConfigReplicationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigReplicationEnabled

`func (o *Engine) SetConfigReplicationEnabled(v bool)`

SetConfigReplicationEnabled sets ConfigReplicationEnabled field to given value.

### HasConfigReplicationEnabled

`func (o *Engine) HasConfigReplicationEnabled() bool`

HasConfigReplicationEnabled returns a boolean if a field has been set.

### GetKeys

`func (o *Engine) GetKeys() []PublicKey`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *Engine) GetKeysOk() (*[]PublicKey, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *Engine) SetKeys(v []PublicKey)`

SetKeys sets Keys field to given value.

### HasKeys

`func (o *Engine) HasKeys() bool`

HasKeys returns a boolean if a field has been set.

### GetHttpProxyId

`func (o *Engine) GetHttpProxyId() int64`

GetHttpProxyId returns the HttpProxyId field if non-nil, zero value otherwise.

### GetHttpProxyIdOk

`func (o *Engine) GetHttpProxyIdOk() (*int64, bool)`

GetHttpProxyIdOk returns a tuple with the HttpProxyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpProxyId

`func (o *Engine) SetHttpProxyId(v int64)`

SetHttpProxyId sets HttpProxyId field to given value.

### HasHttpProxyId

`func (o *Engine) HasHttpProxyId() bool`

HasHttpProxyId returns a boolean if a field has been set.

### GetHttpsProxyId

`func (o *Engine) GetHttpsProxyId() int64`

GetHttpsProxyId returns the HttpsProxyId field if non-nil, zero value otherwise.

### GetHttpsProxyIdOk

`func (o *Engine) GetHttpsProxyIdOk() (*int64, bool)`

GetHttpsProxyIdOk returns a tuple with the HttpsProxyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpsProxyId

`func (o *Engine) SetHttpsProxyId(v int64)`

SetHttpsProxyId sets HttpsProxyId field to given value.

### HasHttpsProxyId

`func (o *Engine) HasHttpsProxyId() bool`

HasHttpsProxyId returns a boolean if a field has been set.

### GetSelectedCertificateId

`func (o *Engine) GetSelectedCertificateId() int64`

GetSelectedCertificateId returns the SelectedCertificateId field if non-nil, zero value otherwise.

### GetSelectedCertificateIdOk

`func (o *Engine) GetSelectedCertificateIdOk() (*int64, bool)`

GetSelectedCertificateIdOk returns a tuple with the SelectedCertificateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedCertificateId

`func (o *Engine) SetSelectedCertificateId(v int64)`

SetSelectedCertificateId sets SelectedCertificateId field to given value.

### HasSelectedCertificateId

`func (o *Engine) HasSelectedCertificateId() bool`

HasSelectedCertificateId returns a boolean if a field has been set.

### GetCertificateHash

`func (o *Engine) GetCertificateHash() Hash`

GetCertificateHash returns the CertificateHash field if non-nil, zero value otherwise.

### GetCertificateHashOk

`func (o *Engine) GetCertificateHashOk() (*Hash, bool)`

GetCertificateHashOk returns a tuple with the CertificateHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateHash

`func (o *Engine) SetCertificateHash(v Hash)`

SetCertificateHash sets CertificateHash field to given value.

### HasCertificateHash

`func (o *Engine) HasCertificateHash() bool`

HasCertificateHash returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


