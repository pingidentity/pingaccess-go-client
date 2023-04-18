# AdminTokenProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description of the Admin Token Provider. | [optional] 
**Issuer** | **string** | The issuer url of the Admin Token Provider. | 
**TrustedCertificateGroupId** | Pointer to **int64** | The group of certificates to use when authenticating to the Admin Token Provider. | [optional] 
**UseProxy** | Pointer to **bool** | True if a proxy should be used for HTTPS requests. | [optional] 
**SslProtocols** | **[]string** | Array of SSL protocolsto use for HTTPS requests. Empty to use all available protocols. | 
**SslCiphers** | **[]string** | Array of SSL ciphers to use for HTTPS requests. Empty to use all available ciphers. | 

## Methods

### NewAdminTokenProvider

`func NewAdminTokenProvider(issuer string, sslProtocols []string, sslCiphers []string, ) *AdminTokenProvider`

NewAdminTokenProvider instantiates a new AdminTokenProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminTokenProviderWithDefaults

`func NewAdminTokenProviderWithDefaults() *AdminTokenProvider`

NewAdminTokenProviderWithDefaults instantiates a new AdminTokenProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *AdminTokenProvider) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AdminTokenProvider) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AdminTokenProvider) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AdminTokenProvider) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIssuer

`func (o *AdminTokenProvider) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *AdminTokenProvider) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *AdminTokenProvider) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.


### GetTrustedCertificateGroupId

`func (o *AdminTokenProvider) GetTrustedCertificateGroupId() int64`

GetTrustedCertificateGroupId returns the TrustedCertificateGroupId field if non-nil, zero value otherwise.

### GetTrustedCertificateGroupIdOk

`func (o *AdminTokenProvider) GetTrustedCertificateGroupIdOk() (*int64, bool)`

GetTrustedCertificateGroupIdOk returns a tuple with the TrustedCertificateGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustedCertificateGroupId

`func (o *AdminTokenProvider) SetTrustedCertificateGroupId(v int64)`

SetTrustedCertificateGroupId sets TrustedCertificateGroupId field to given value.

### HasTrustedCertificateGroupId

`func (o *AdminTokenProvider) HasTrustedCertificateGroupId() bool`

HasTrustedCertificateGroupId returns a boolean if a field has been set.

### GetUseProxy

`func (o *AdminTokenProvider) GetUseProxy() bool`

GetUseProxy returns the UseProxy field if non-nil, zero value otherwise.

### GetUseProxyOk

`func (o *AdminTokenProvider) GetUseProxyOk() (*bool, bool)`

GetUseProxyOk returns a tuple with the UseProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseProxy

`func (o *AdminTokenProvider) SetUseProxy(v bool)`

SetUseProxy sets UseProxy field to given value.

### HasUseProxy

`func (o *AdminTokenProvider) HasUseProxy() bool`

HasUseProxy returns a boolean if a field has been set.

### GetSslProtocols

`func (o *AdminTokenProvider) GetSslProtocols() []string`

GetSslProtocols returns the SslProtocols field if non-nil, zero value otherwise.

### GetSslProtocolsOk

`func (o *AdminTokenProvider) GetSslProtocolsOk() (*[]string, bool)`

GetSslProtocolsOk returns a tuple with the SslProtocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslProtocols

`func (o *AdminTokenProvider) SetSslProtocols(v []string)`

SetSslProtocols sets SslProtocols field to given value.


### GetSslCiphers

`func (o *AdminTokenProvider) GetSslCiphers() []string`

GetSslCiphers returns the SslCiphers field if non-nil, zero value otherwise.

### GetSslCiphersOk

`func (o *AdminTokenProvider) GetSslCiphersOk() (*[]string, bool)`

GetSslCiphersOk returns a tuple with the SslCiphers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCiphers

`func (o *AdminTokenProvider) SetSslCiphers(v []string)`

SetSslCiphers sets SslCiphers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


