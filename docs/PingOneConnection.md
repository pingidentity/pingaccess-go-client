# PingOneConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | When creating a new PingOneConnection, this is the ID for the PingOneConnection. If not specified, an ID will be automatically assigned. When updating an existing PingOneConnection, this field is ignored and the ID is determined by the path parameter. | [optional] 
**Name** | **string** | (sortable) A unique name for the PingOne Connection | 
**Description** | Pointer to **string** | (sortable) The description of the PingOne Connection. | [optional] 
**Credential** | Pointer to [**HiddenField**](HiddenField.md) |  | [optional] 
**UseProxy** | Pointer to **bool** | (sortable) True if a proxy should be used for HTTPS requests to PingOne. | [optional] 
**TrustedCertificateGroupId** | Pointer to **int32** | The group of certificates to use when communicating to PingOne through this connection. | [optional] 
**SslProtocols** | **[]string** | Array of SSL protocols to use for HTTPS requests. Empty to use all available protocols. | 
**SslCiphers** | **[]string** | Array of SSL ciphers to use for HTTPS requests. Empty to use all available ciphers. | 

## Methods

### NewPingOneConnection

`func NewPingOneConnection(name string, sslProtocols []string, sslCiphers []string, ) *PingOneConnection`

NewPingOneConnection instantiates a new PingOneConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPingOneConnectionWithDefaults

`func NewPingOneConnectionWithDefaults() *PingOneConnection`

NewPingOneConnectionWithDefaults instantiates a new PingOneConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PingOneConnection) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PingOneConnection) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PingOneConnection) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PingOneConnection) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PingOneConnection) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PingOneConnection) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PingOneConnection) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *PingOneConnection) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PingOneConnection) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PingOneConnection) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PingOneConnection) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCredential

`func (o *PingOneConnection) GetCredential() HiddenField`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *PingOneConnection) GetCredentialOk() (*HiddenField, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *PingOneConnection) SetCredential(v HiddenField)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *PingOneConnection) HasCredential() bool`

HasCredential returns a boolean if a field has been set.

### GetUseProxy

`func (o *PingOneConnection) GetUseProxy() bool`

GetUseProxy returns the UseProxy field if non-nil, zero value otherwise.

### GetUseProxyOk

`func (o *PingOneConnection) GetUseProxyOk() (*bool, bool)`

GetUseProxyOk returns a tuple with the UseProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseProxy

`func (o *PingOneConnection) SetUseProxy(v bool)`

SetUseProxy sets UseProxy field to given value.

### HasUseProxy

`func (o *PingOneConnection) HasUseProxy() bool`

HasUseProxy returns a boolean if a field has been set.

### GetTrustedCertificateGroupId

`func (o *PingOneConnection) GetTrustedCertificateGroupId() int32`

GetTrustedCertificateGroupId returns the TrustedCertificateGroupId field if non-nil, zero value otherwise.

### GetTrustedCertificateGroupIdOk

`func (o *PingOneConnection) GetTrustedCertificateGroupIdOk() (*int32, bool)`

GetTrustedCertificateGroupIdOk returns a tuple with the TrustedCertificateGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustedCertificateGroupId

`func (o *PingOneConnection) SetTrustedCertificateGroupId(v int32)`

SetTrustedCertificateGroupId sets TrustedCertificateGroupId field to given value.

### HasTrustedCertificateGroupId

`func (o *PingOneConnection) HasTrustedCertificateGroupId() bool`

HasTrustedCertificateGroupId returns a boolean if a field has been set.

### GetSslProtocols

`func (o *PingOneConnection) GetSslProtocols() []string`

GetSslProtocols returns the SslProtocols field if non-nil, zero value otherwise.

### GetSslProtocolsOk

`func (o *PingOneConnection) GetSslProtocolsOk() (*[]string, bool)`

GetSslProtocolsOk returns a tuple with the SslProtocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslProtocols

`func (o *PingOneConnection) SetSslProtocols(v []string)`

SetSslProtocols sets SslProtocols field to given value.


### GetSslCiphers

`func (o *PingOneConnection) GetSslCiphers() []string`

GetSslCiphers returns the SslCiphers field if non-nil, zero value otherwise.

### GetSslCiphersOk

`func (o *PingOneConnection) GetSslCiphersOk() (*[]string, bool)`

GetSslCiphersOk returns a tuple with the SslCiphers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCiphers

`func (o *PingOneConnection) SetSslCiphers(v []string)`

SetSslCiphers sets SslCiphers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


