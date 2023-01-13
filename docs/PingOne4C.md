# PingOne4C

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description of the PingOne for Customers OIDC provider. | [optional] 
**Issuer** | **string** | The issuer url of the PingOne for Customers OIDC provider. | 
**TrustedCertificateGroupId** | Pointer to **int32** | The group of certificates to use when authenticating to PingOne for Customers OIDC provider. | [optional] 
**UseProxy** | Pointer to **bool** | True if a proxy should be used for HTTPS requests. | [optional] 

## Methods

### NewPingOne4C

`func NewPingOne4C(issuer string, ) *PingOne4C`

NewPingOne4C instantiates a new PingOne4C object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPingOne4CWithDefaults

`func NewPingOne4CWithDefaults() *PingOne4C`

NewPingOne4CWithDefaults instantiates a new PingOne4C object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *PingOne4C) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PingOne4C) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PingOne4C) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PingOne4C) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIssuer

`func (o *PingOne4C) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *PingOne4C) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *PingOne4C) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.


### GetTrustedCertificateGroupId

`func (o *PingOne4C) GetTrustedCertificateGroupId() int32`

GetTrustedCertificateGroupId returns the TrustedCertificateGroupId field if non-nil, zero value otherwise.

### GetTrustedCertificateGroupIdOk

`func (o *PingOne4C) GetTrustedCertificateGroupIdOk() (*int32, bool)`

GetTrustedCertificateGroupIdOk returns a tuple with the TrustedCertificateGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustedCertificateGroupId

`func (o *PingOne4C) SetTrustedCertificateGroupId(v int32)`

SetTrustedCertificateGroupId sets TrustedCertificateGroupId field to given value.

### HasTrustedCertificateGroupId

`func (o *PingOne4C) HasTrustedCertificateGroupId() bool`

HasTrustedCertificateGroupId returns a boolean if a field has been set.

### GetUseProxy

`func (o *PingOne4C) GetUseProxy() bool`

GetUseProxy returns the UseProxy field if non-nil, zero value otherwise.

### GetUseProxyOk

`func (o *PingOne4C) GetUseProxyOk() (*bool, bool)`

GetUseProxyOk returns a tuple with the UseProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseProxy

`func (o *PingOne4C) SetUseProxy(v bool)`

SetUseProxy sets UseProxy field to given value.

### HasUseProxy

`func (o *PingOne4C) HasUseProxy() bool`

HasUseProxy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


