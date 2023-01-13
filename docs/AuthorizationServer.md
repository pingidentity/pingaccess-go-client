# AuthorizationServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description of the third-party OAuth 2.0 Authorization Server. | [optional] 
**Targets** | **[]string** | One or more server hostname:port pairs used to access third-party OAuth 2.0 Authorization Server. | 
**AuditLevel** | Pointer to [**AuditLevel**](AuditLevel.md) |  | [optional] 
**Secure** | Pointer to **bool** | Enable if third-party OAuth 2.0 Authorization Server is expecting HTTPS connections. | [optional] 
**TrustedCertificateGroupId** | **int32** | The group of certificates to use when authenticating to third-party OAuth 2.0 Authorization Server. | 
**ClientCredentials** | Pointer to [**OAuthClientCredentials**](OAuthClientCredentials.md) |  | [optional] 
**CacheTokens** | Pointer to **bool** | Enable to retain token details for subsequent requests. | [optional] 
**TokenTimeToLiveSeconds** | Pointer to **int64** | Defines the number of seconds to cache the access token. -1 means no limit. This value should be less than the PingFederate Token Lifetime. | [optional] 
**SubjectAttributeName** | **string** | The attribute you want to use from the OAuth access token as the subject for auditing purposes. | 
**IntrospectionEndpoint** | **string** | The third-party OAuth 2.0 Authorization Server&#39;s token introspection endpoint. | 
**TokenEndpoint** | Pointer to **string** | The third-party OAuth 2.0 Authorization Server&#39;s token endpoint. | [optional] 
**SendAudience** | Pointer to **bool** | Enable to send the URI the user requested as the &#39;aud&#39; OAuth parameter for PingAccess to the OAuth 2.0 Authorization server. | [optional] 
**UseProxy** | Pointer to **bool** | True if a proxy should be used for HTTP or HTTPS requests. | [optional] 

## Methods

### NewAuthorizationServer

`func NewAuthorizationServer(targets []string, trustedCertificateGroupId int32, subjectAttributeName string, introspectionEndpoint string, ) *AuthorizationServer`

NewAuthorizationServer instantiates a new AuthorizationServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizationServerWithDefaults

`func NewAuthorizationServerWithDefaults() *AuthorizationServer`

NewAuthorizationServerWithDefaults instantiates a new AuthorizationServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *AuthorizationServer) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AuthorizationServer) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AuthorizationServer) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AuthorizationServer) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTargets

`func (o *AuthorizationServer) GetTargets() []string`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *AuthorizationServer) GetTargetsOk() (*[]string, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *AuthorizationServer) SetTargets(v []string)`

SetTargets sets Targets field to given value.


### GetAuditLevel

`func (o *AuthorizationServer) GetAuditLevel() AuditLevel`

GetAuditLevel returns the AuditLevel field if non-nil, zero value otherwise.

### GetAuditLevelOk

`func (o *AuthorizationServer) GetAuditLevelOk() (*AuditLevel, bool)`

GetAuditLevelOk returns a tuple with the AuditLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditLevel

`func (o *AuthorizationServer) SetAuditLevel(v AuditLevel)`

SetAuditLevel sets AuditLevel field to given value.

### HasAuditLevel

`func (o *AuthorizationServer) HasAuditLevel() bool`

HasAuditLevel returns a boolean if a field has been set.

### GetSecure

`func (o *AuthorizationServer) GetSecure() bool`

GetSecure returns the Secure field if non-nil, zero value otherwise.

### GetSecureOk

`func (o *AuthorizationServer) GetSecureOk() (*bool, bool)`

GetSecureOk returns a tuple with the Secure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecure

`func (o *AuthorizationServer) SetSecure(v bool)`

SetSecure sets Secure field to given value.

### HasSecure

`func (o *AuthorizationServer) HasSecure() bool`

HasSecure returns a boolean if a field has been set.

### GetTrustedCertificateGroupId

`func (o *AuthorizationServer) GetTrustedCertificateGroupId() int32`

GetTrustedCertificateGroupId returns the TrustedCertificateGroupId field if non-nil, zero value otherwise.

### GetTrustedCertificateGroupIdOk

`func (o *AuthorizationServer) GetTrustedCertificateGroupIdOk() (*int32, bool)`

GetTrustedCertificateGroupIdOk returns a tuple with the TrustedCertificateGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustedCertificateGroupId

`func (o *AuthorizationServer) SetTrustedCertificateGroupId(v int32)`

SetTrustedCertificateGroupId sets TrustedCertificateGroupId field to given value.


### GetClientCredentials

`func (o *AuthorizationServer) GetClientCredentials() OAuthClientCredentials`

GetClientCredentials returns the ClientCredentials field if non-nil, zero value otherwise.

### GetClientCredentialsOk

`func (o *AuthorizationServer) GetClientCredentialsOk() (*OAuthClientCredentials, bool)`

GetClientCredentialsOk returns a tuple with the ClientCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCredentials

`func (o *AuthorizationServer) SetClientCredentials(v OAuthClientCredentials)`

SetClientCredentials sets ClientCredentials field to given value.

### HasClientCredentials

`func (o *AuthorizationServer) HasClientCredentials() bool`

HasClientCredentials returns a boolean if a field has been set.

### GetCacheTokens

`func (o *AuthorizationServer) GetCacheTokens() bool`

GetCacheTokens returns the CacheTokens field if non-nil, zero value otherwise.

### GetCacheTokensOk

`func (o *AuthorizationServer) GetCacheTokensOk() (*bool, bool)`

GetCacheTokensOk returns a tuple with the CacheTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheTokens

`func (o *AuthorizationServer) SetCacheTokens(v bool)`

SetCacheTokens sets CacheTokens field to given value.

### HasCacheTokens

`func (o *AuthorizationServer) HasCacheTokens() bool`

HasCacheTokens returns a boolean if a field has been set.

### GetTokenTimeToLiveSeconds

`func (o *AuthorizationServer) GetTokenTimeToLiveSeconds() int64`

GetTokenTimeToLiveSeconds returns the TokenTimeToLiveSeconds field if non-nil, zero value otherwise.

### GetTokenTimeToLiveSecondsOk

`func (o *AuthorizationServer) GetTokenTimeToLiveSecondsOk() (*int64, bool)`

GetTokenTimeToLiveSecondsOk returns a tuple with the TokenTimeToLiveSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenTimeToLiveSeconds

`func (o *AuthorizationServer) SetTokenTimeToLiveSeconds(v int64)`

SetTokenTimeToLiveSeconds sets TokenTimeToLiveSeconds field to given value.

### HasTokenTimeToLiveSeconds

`func (o *AuthorizationServer) HasTokenTimeToLiveSeconds() bool`

HasTokenTimeToLiveSeconds returns a boolean if a field has been set.

### GetSubjectAttributeName

`func (o *AuthorizationServer) GetSubjectAttributeName() string`

GetSubjectAttributeName returns the SubjectAttributeName field if non-nil, zero value otherwise.

### GetSubjectAttributeNameOk

`func (o *AuthorizationServer) GetSubjectAttributeNameOk() (*string, bool)`

GetSubjectAttributeNameOk returns a tuple with the SubjectAttributeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectAttributeName

`func (o *AuthorizationServer) SetSubjectAttributeName(v string)`

SetSubjectAttributeName sets SubjectAttributeName field to given value.


### GetIntrospectionEndpoint

`func (o *AuthorizationServer) GetIntrospectionEndpoint() string`

GetIntrospectionEndpoint returns the IntrospectionEndpoint field if non-nil, zero value otherwise.

### GetIntrospectionEndpointOk

`func (o *AuthorizationServer) GetIntrospectionEndpointOk() (*string, bool)`

GetIntrospectionEndpointOk returns a tuple with the IntrospectionEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntrospectionEndpoint

`func (o *AuthorizationServer) SetIntrospectionEndpoint(v string)`

SetIntrospectionEndpoint sets IntrospectionEndpoint field to given value.


### GetTokenEndpoint

`func (o *AuthorizationServer) GetTokenEndpoint() string`

GetTokenEndpoint returns the TokenEndpoint field if non-nil, zero value otherwise.

### GetTokenEndpointOk

`func (o *AuthorizationServer) GetTokenEndpointOk() (*string, bool)`

GetTokenEndpointOk returns a tuple with the TokenEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenEndpoint

`func (o *AuthorizationServer) SetTokenEndpoint(v string)`

SetTokenEndpoint sets TokenEndpoint field to given value.

### HasTokenEndpoint

`func (o *AuthorizationServer) HasTokenEndpoint() bool`

HasTokenEndpoint returns a boolean if a field has been set.

### GetSendAudience

`func (o *AuthorizationServer) GetSendAudience() bool`

GetSendAudience returns the SendAudience field if non-nil, zero value otherwise.

### GetSendAudienceOk

`func (o *AuthorizationServer) GetSendAudienceOk() (*bool, bool)`

GetSendAudienceOk returns a tuple with the SendAudience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendAudience

`func (o *AuthorizationServer) SetSendAudience(v bool)`

SetSendAudience sets SendAudience field to given value.

### HasSendAudience

`func (o *AuthorizationServer) HasSendAudience() bool`

HasSendAudience returns a boolean if a field has been set.

### GetUseProxy

`func (o *AuthorizationServer) GetUseProxy() bool`

GetUseProxy returns the UseProxy field if non-nil, zero value otherwise.

### GetUseProxyOk

`func (o *AuthorizationServer) GetUseProxyOk() (*bool, bool)`

GetUseProxyOk returns a tuple with the UseProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseProxy

`func (o *AuthorizationServer) SetUseProxy(v bool)`

SetUseProxy sets UseProxy field to given value.

### HasUseProxy

`func (o *AuthorizationServer) HasUseProxy() bool`

HasUseProxy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


