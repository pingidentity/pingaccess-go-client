# PingFederateMetadataRuntime

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description of the PingFederate Runtime token provider. | [optional] 
**Issuer** | **string** | The issuer url of the PingFederate token provider. | 
**TrustedCertificateGroupId** | Pointer to **int32** | The group of certificates to use when authenticating to PingFederate. | [optional] 
**UseProxy** | Pointer to **bool** | Set to true if a proxy should be used for HTTP or HTTPS requests. | [optional] 
**UseSlo** | Pointer to **bool** | Set to true if OIDC single log out should be used on the /pa/oidc/logout on the engines. | [optional] 
**StsTokenExchangeEndpoint** | Pointer to **string** | The url of the PingFederate STS token-to-token exchange endpoint that is used for token mediation. Specify if it is being served from a different host/port than the issuer is. Otherwise, it is assumed to be {{issuer}}/pf/sts.wst. | [optional] 
**SkipHostnameVerification** | Pointer to **bool** | Set to true if HTTP communications to PingFederate should not perform hostname verification of the certificate. | [optional] 

## Methods

### NewPingFederateMetadataRuntime

`func NewPingFederateMetadataRuntime(issuer string, ) *PingFederateMetadataRuntime`

NewPingFederateMetadataRuntime instantiates a new PingFederateMetadataRuntime object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPingFederateMetadataRuntimeWithDefaults

`func NewPingFederateMetadataRuntimeWithDefaults() *PingFederateMetadataRuntime`

NewPingFederateMetadataRuntimeWithDefaults instantiates a new PingFederateMetadataRuntime object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *PingFederateMetadataRuntime) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PingFederateMetadataRuntime) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PingFederateMetadataRuntime) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PingFederateMetadataRuntime) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIssuer

`func (o *PingFederateMetadataRuntime) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *PingFederateMetadataRuntime) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *PingFederateMetadataRuntime) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.


### GetTrustedCertificateGroupId

`func (o *PingFederateMetadataRuntime) GetTrustedCertificateGroupId() int32`

GetTrustedCertificateGroupId returns the TrustedCertificateGroupId field if non-nil, zero value otherwise.

### GetTrustedCertificateGroupIdOk

`func (o *PingFederateMetadataRuntime) GetTrustedCertificateGroupIdOk() (*int32, bool)`

GetTrustedCertificateGroupIdOk returns a tuple with the TrustedCertificateGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustedCertificateGroupId

`func (o *PingFederateMetadataRuntime) SetTrustedCertificateGroupId(v int32)`

SetTrustedCertificateGroupId sets TrustedCertificateGroupId field to given value.

### HasTrustedCertificateGroupId

`func (o *PingFederateMetadataRuntime) HasTrustedCertificateGroupId() bool`

HasTrustedCertificateGroupId returns a boolean if a field has been set.

### GetUseProxy

`func (o *PingFederateMetadataRuntime) GetUseProxy() bool`

GetUseProxy returns the UseProxy field if non-nil, zero value otherwise.

### GetUseProxyOk

`func (o *PingFederateMetadataRuntime) GetUseProxyOk() (*bool, bool)`

GetUseProxyOk returns a tuple with the UseProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseProxy

`func (o *PingFederateMetadataRuntime) SetUseProxy(v bool)`

SetUseProxy sets UseProxy field to given value.

### HasUseProxy

`func (o *PingFederateMetadataRuntime) HasUseProxy() bool`

HasUseProxy returns a boolean if a field has been set.

### GetUseSlo

`func (o *PingFederateMetadataRuntime) GetUseSlo() bool`

GetUseSlo returns the UseSlo field if non-nil, zero value otherwise.

### GetUseSloOk

`func (o *PingFederateMetadataRuntime) GetUseSloOk() (*bool, bool)`

GetUseSloOk returns a tuple with the UseSlo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSlo

`func (o *PingFederateMetadataRuntime) SetUseSlo(v bool)`

SetUseSlo sets UseSlo field to given value.

### HasUseSlo

`func (o *PingFederateMetadataRuntime) HasUseSlo() bool`

HasUseSlo returns a boolean if a field has been set.

### GetStsTokenExchangeEndpoint

`func (o *PingFederateMetadataRuntime) GetStsTokenExchangeEndpoint() string`

GetStsTokenExchangeEndpoint returns the StsTokenExchangeEndpoint field if non-nil, zero value otherwise.

### GetStsTokenExchangeEndpointOk

`func (o *PingFederateMetadataRuntime) GetStsTokenExchangeEndpointOk() (*string, bool)`

GetStsTokenExchangeEndpointOk returns a tuple with the StsTokenExchangeEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStsTokenExchangeEndpoint

`func (o *PingFederateMetadataRuntime) SetStsTokenExchangeEndpoint(v string)`

SetStsTokenExchangeEndpoint sets StsTokenExchangeEndpoint field to given value.

### HasStsTokenExchangeEndpoint

`func (o *PingFederateMetadataRuntime) HasStsTokenExchangeEndpoint() bool`

HasStsTokenExchangeEndpoint returns a boolean if a field has been set.

### GetSkipHostnameVerification

`func (o *PingFederateMetadataRuntime) GetSkipHostnameVerification() bool`

GetSkipHostnameVerification returns the SkipHostnameVerification field if non-nil, zero value otherwise.

### GetSkipHostnameVerificationOk

`func (o *PingFederateMetadataRuntime) GetSkipHostnameVerificationOk() (*bool, bool)`

GetSkipHostnameVerificationOk returns a tuple with the SkipHostnameVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipHostnameVerification

`func (o *PingFederateMetadataRuntime) SetSkipHostnameVerification(v bool)`

SetSkipHostnameVerification sets SkipHostnameVerification field to given value.

### HasSkipHostnameVerification

`func (o *PingFederateMetadataRuntime) HasSkipHostnameVerification() bool`

HasSkipHostnameVerification returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


