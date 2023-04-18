# OIDCProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description of the third-party OpenID Connect provider. | [optional] 
**Issuer** | **string** | The issuer of the third-party OpenID Connect provider. | 
**TrustedCertificateGroupId** | Pointer to **int64** | The group of certificates to use when authenticating to third-party OpenID Connect provider. | [optional] 
**UseProxy** | Pointer to **bool** | True if a proxy should be used for HTTP or HTTPS requests. | [optional] 
**AuditLevel** | Pointer to [**AuditLevel**](AuditLevel.md) |  | [optional] 
**QueryParameters** | Pointer to [**[]QueryParameter**](QueryParameter.md) | The query parameters used by the authentication request to third-party OpenID Connect provider. | [optional] 
**UseSlo** | Pointer to **bool** | True if single log off (SLO) should be used. | [optional] 
**RequestSupportedScopesOnly** | Pointer to **bool** | Specifies whether the scopes in an access request should be limited to those advertised in the OIDC metadata. | [optional] 
**Plugin** | Pointer to [**OIDCProviderPlugin**](OIDCProviderPlugin.md) |  | [optional] 

## Methods

### NewOIDCProvider

`func NewOIDCProvider(issuer string, ) *OIDCProvider`

NewOIDCProvider instantiates a new OIDCProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOIDCProviderWithDefaults

`func NewOIDCProviderWithDefaults() *OIDCProvider`

NewOIDCProviderWithDefaults instantiates a new OIDCProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *OIDCProvider) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OIDCProvider) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OIDCProvider) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OIDCProvider) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIssuer

`func (o *OIDCProvider) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *OIDCProvider) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *OIDCProvider) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.


### GetTrustedCertificateGroupId

`func (o *OIDCProvider) GetTrustedCertificateGroupId() int64`

GetTrustedCertificateGroupId returns the TrustedCertificateGroupId field if non-nil, zero value otherwise.

### GetTrustedCertificateGroupIdOk

`func (o *OIDCProvider) GetTrustedCertificateGroupIdOk() (*int64, bool)`

GetTrustedCertificateGroupIdOk returns a tuple with the TrustedCertificateGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustedCertificateGroupId

`func (o *OIDCProvider) SetTrustedCertificateGroupId(v int64)`

SetTrustedCertificateGroupId sets TrustedCertificateGroupId field to given value.

### HasTrustedCertificateGroupId

`func (o *OIDCProvider) HasTrustedCertificateGroupId() bool`

HasTrustedCertificateGroupId returns a boolean if a field has been set.

### GetUseProxy

`func (o *OIDCProvider) GetUseProxy() bool`

GetUseProxy returns the UseProxy field if non-nil, zero value otherwise.

### GetUseProxyOk

`func (o *OIDCProvider) GetUseProxyOk() (*bool, bool)`

GetUseProxyOk returns a tuple with the UseProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseProxy

`func (o *OIDCProvider) SetUseProxy(v bool)`

SetUseProxy sets UseProxy field to given value.

### HasUseProxy

`func (o *OIDCProvider) HasUseProxy() bool`

HasUseProxy returns a boolean if a field has been set.

### GetAuditLevel

`func (o *OIDCProvider) GetAuditLevel() AuditLevel`

GetAuditLevel returns the AuditLevel field if non-nil, zero value otherwise.

### GetAuditLevelOk

`func (o *OIDCProvider) GetAuditLevelOk() (*AuditLevel, bool)`

GetAuditLevelOk returns a tuple with the AuditLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditLevel

`func (o *OIDCProvider) SetAuditLevel(v AuditLevel)`

SetAuditLevel sets AuditLevel field to given value.

### HasAuditLevel

`func (o *OIDCProvider) HasAuditLevel() bool`

HasAuditLevel returns a boolean if a field has been set.

### GetQueryParameters

`func (o *OIDCProvider) GetQueryParameters() []QueryParameter`

GetQueryParameters returns the QueryParameters field if non-nil, zero value otherwise.

### GetQueryParametersOk

`func (o *OIDCProvider) GetQueryParametersOk() (*[]QueryParameter, bool)`

GetQueryParametersOk returns a tuple with the QueryParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryParameters

`func (o *OIDCProvider) SetQueryParameters(v []QueryParameter)`

SetQueryParameters sets QueryParameters field to given value.

### HasQueryParameters

`func (o *OIDCProvider) HasQueryParameters() bool`

HasQueryParameters returns a boolean if a field has been set.

### GetUseSlo

`func (o *OIDCProvider) GetUseSlo() bool`

GetUseSlo returns the UseSlo field if non-nil, zero value otherwise.

### GetUseSloOk

`func (o *OIDCProvider) GetUseSloOk() (*bool, bool)`

GetUseSloOk returns a tuple with the UseSlo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSlo

`func (o *OIDCProvider) SetUseSlo(v bool)`

SetUseSlo sets UseSlo field to given value.

### HasUseSlo

`func (o *OIDCProvider) HasUseSlo() bool`

HasUseSlo returns a boolean if a field has been set.

### GetRequestSupportedScopesOnly

`func (o *OIDCProvider) GetRequestSupportedScopesOnly() bool`

GetRequestSupportedScopesOnly returns the RequestSupportedScopesOnly field if non-nil, zero value otherwise.

### GetRequestSupportedScopesOnlyOk

`func (o *OIDCProvider) GetRequestSupportedScopesOnlyOk() (*bool, bool)`

GetRequestSupportedScopesOnlyOk returns a tuple with the RequestSupportedScopesOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestSupportedScopesOnly

`func (o *OIDCProvider) SetRequestSupportedScopesOnly(v bool)`

SetRequestSupportedScopesOnly sets RequestSupportedScopesOnly field to given value.

### HasRequestSupportedScopesOnly

`func (o *OIDCProvider) HasRequestSupportedScopesOnly() bool`

HasRequestSupportedScopesOnly returns a boolean if a field has been set.

### GetPlugin

`func (o *OIDCProvider) GetPlugin() OIDCProviderPlugin`

GetPlugin returns the Plugin field if non-nil, zero value otherwise.

### GetPluginOk

`func (o *OIDCProvider) GetPluginOk() (*OIDCProviderPlugin, bool)`

GetPluginOk returns a tuple with the Plugin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlugin

`func (o *OIDCProvider) SetPlugin(v OIDCProviderPlugin)`

SetPlugin sets Plugin field to given value.

### HasPlugin

`func (o *OIDCProvider) HasPlugin() bool`

HasPlugin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


