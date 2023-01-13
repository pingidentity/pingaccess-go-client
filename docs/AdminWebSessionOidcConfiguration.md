# AdminWebSessionOidcConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RefreshUserInfoClaimsInterval** | Pointer to **int32** | Specify the maximum number of seconds to cache user attribute information when the Refresh User is enabled. | [optional] 
**SendRequestedUrlToProvider** | Pointer to **bool** | Specify if you want to send the requested URL as part of the authentication request to the OpenID Connect Provider. | [optional] 
**ValidateSessionIsAlive** | Pointer to **bool** | Specify if PingAccess should validate sessions with the configured PingFederate instance during request processing. | [optional] 
**ClientCredentials** | [**OAuthClientCredentials**](OAuthClientCredentials.md) |  | 
**OidcLoginType** | Pointer to [**OidcLoginType**](OidcLoginType.md) |  | [optional] 
**PkceChallengeType** | Pointer to [**PkceChallengeType**](PkceChallengeType.md) |  | [optional] 
**CacheUserAttributes** | Pointer to **bool** | Specify if PingAccess should cache user attribute information for use in policy decisions. When disabled, this data is encoded and stored in the session cookie. | [optional] 
**EnableRefreshUser** | Pointer to **bool** | Specify if you want to have PingAccess periodically refresh user data from PingFederate for use in policy decisions. | [optional] 
**Scopes** | Pointer to **[]string** | The list of scopes. The openid scope is implied and does not need to be specified in this list. | [optional] 
**PfsessionStateCacheInSeconds** | Pointer to **int32** | Specify the number of seconds to cache PingFederate Session State information. | [optional] 

## Methods

### NewAdminWebSessionOidcConfiguration

`func NewAdminWebSessionOidcConfiguration(clientCredentials OAuthClientCredentials, ) *AdminWebSessionOidcConfiguration`

NewAdminWebSessionOidcConfiguration instantiates a new AdminWebSessionOidcConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminWebSessionOidcConfigurationWithDefaults

`func NewAdminWebSessionOidcConfigurationWithDefaults() *AdminWebSessionOidcConfiguration`

NewAdminWebSessionOidcConfigurationWithDefaults instantiates a new AdminWebSessionOidcConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRefreshUserInfoClaimsInterval

`func (o *AdminWebSessionOidcConfiguration) GetRefreshUserInfoClaimsInterval() int32`

GetRefreshUserInfoClaimsInterval returns the RefreshUserInfoClaimsInterval field if non-nil, zero value otherwise.

### GetRefreshUserInfoClaimsIntervalOk

`func (o *AdminWebSessionOidcConfiguration) GetRefreshUserInfoClaimsIntervalOk() (*int32, bool)`

GetRefreshUserInfoClaimsIntervalOk returns a tuple with the RefreshUserInfoClaimsInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshUserInfoClaimsInterval

`func (o *AdminWebSessionOidcConfiguration) SetRefreshUserInfoClaimsInterval(v int32)`

SetRefreshUserInfoClaimsInterval sets RefreshUserInfoClaimsInterval field to given value.

### HasRefreshUserInfoClaimsInterval

`func (o *AdminWebSessionOidcConfiguration) HasRefreshUserInfoClaimsInterval() bool`

HasRefreshUserInfoClaimsInterval returns a boolean if a field has been set.

### GetSendRequestedUrlToProvider

`func (o *AdminWebSessionOidcConfiguration) GetSendRequestedUrlToProvider() bool`

GetSendRequestedUrlToProvider returns the SendRequestedUrlToProvider field if non-nil, zero value otherwise.

### GetSendRequestedUrlToProviderOk

`func (o *AdminWebSessionOidcConfiguration) GetSendRequestedUrlToProviderOk() (*bool, bool)`

GetSendRequestedUrlToProviderOk returns a tuple with the SendRequestedUrlToProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendRequestedUrlToProvider

`func (o *AdminWebSessionOidcConfiguration) SetSendRequestedUrlToProvider(v bool)`

SetSendRequestedUrlToProvider sets SendRequestedUrlToProvider field to given value.

### HasSendRequestedUrlToProvider

`func (o *AdminWebSessionOidcConfiguration) HasSendRequestedUrlToProvider() bool`

HasSendRequestedUrlToProvider returns a boolean if a field has been set.

### GetValidateSessionIsAlive

`func (o *AdminWebSessionOidcConfiguration) GetValidateSessionIsAlive() bool`

GetValidateSessionIsAlive returns the ValidateSessionIsAlive field if non-nil, zero value otherwise.

### GetValidateSessionIsAliveOk

`func (o *AdminWebSessionOidcConfiguration) GetValidateSessionIsAliveOk() (*bool, bool)`

GetValidateSessionIsAliveOk returns a tuple with the ValidateSessionIsAlive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidateSessionIsAlive

`func (o *AdminWebSessionOidcConfiguration) SetValidateSessionIsAlive(v bool)`

SetValidateSessionIsAlive sets ValidateSessionIsAlive field to given value.

### HasValidateSessionIsAlive

`func (o *AdminWebSessionOidcConfiguration) HasValidateSessionIsAlive() bool`

HasValidateSessionIsAlive returns a boolean if a field has been set.

### GetClientCredentials

`func (o *AdminWebSessionOidcConfiguration) GetClientCredentials() OAuthClientCredentials`

GetClientCredentials returns the ClientCredentials field if non-nil, zero value otherwise.

### GetClientCredentialsOk

`func (o *AdminWebSessionOidcConfiguration) GetClientCredentialsOk() (*OAuthClientCredentials, bool)`

GetClientCredentialsOk returns a tuple with the ClientCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCredentials

`func (o *AdminWebSessionOidcConfiguration) SetClientCredentials(v OAuthClientCredentials)`

SetClientCredentials sets ClientCredentials field to given value.


### GetOidcLoginType

`func (o *AdminWebSessionOidcConfiguration) GetOidcLoginType() OidcLoginType`

GetOidcLoginType returns the OidcLoginType field if non-nil, zero value otherwise.

### GetOidcLoginTypeOk

`func (o *AdminWebSessionOidcConfiguration) GetOidcLoginTypeOk() (*OidcLoginType, bool)`

GetOidcLoginTypeOk returns a tuple with the OidcLoginType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcLoginType

`func (o *AdminWebSessionOidcConfiguration) SetOidcLoginType(v OidcLoginType)`

SetOidcLoginType sets OidcLoginType field to given value.

### HasOidcLoginType

`func (o *AdminWebSessionOidcConfiguration) HasOidcLoginType() bool`

HasOidcLoginType returns a boolean if a field has been set.

### GetPkceChallengeType

`func (o *AdminWebSessionOidcConfiguration) GetPkceChallengeType() PkceChallengeType`

GetPkceChallengeType returns the PkceChallengeType field if non-nil, zero value otherwise.

### GetPkceChallengeTypeOk

`func (o *AdminWebSessionOidcConfiguration) GetPkceChallengeTypeOk() (*PkceChallengeType, bool)`

GetPkceChallengeTypeOk returns a tuple with the PkceChallengeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkceChallengeType

`func (o *AdminWebSessionOidcConfiguration) SetPkceChallengeType(v PkceChallengeType)`

SetPkceChallengeType sets PkceChallengeType field to given value.

### HasPkceChallengeType

`func (o *AdminWebSessionOidcConfiguration) HasPkceChallengeType() bool`

HasPkceChallengeType returns a boolean if a field has been set.

### GetCacheUserAttributes

`func (o *AdminWebSessionOidcConfiguration) GetCacheUserAttributes() bool`

GetCacheUserAttributes returns the CacheUserAttributes field if non-nil, zero value otherwise.

### GetCacheUserAttributesOk

`func (o *AdminWebSessionOidcConfiguration) GetCacheUserAttributesOk() (*bool, bool)`

GetCacheUserAttributesOk returns a tuple with the CacheUserAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheUserAttributes

`func (o *AdminWebSessionOidcConfiguration) SetCacheUserAttributes(v bool)`

SetCacheUserAttributes sets CacheUserAttributes field to given value.

### HasCacheUserAttributes

`func (o *AdminWebSessionOidcConfiguration) HasCacheUserAttributes() bool`

HasCacheUserAttributes returns a boolean if a field has been set.

### GetEnableRefreshUser

`func (o *AdminWebSessionOidcConfiguration) GetEnableRefreshUser() bool`

GetEnableRefreshUser returns the EnableRefreshUser field if non-nil, zero value otherwise.

### GetEnableRefreshUserOk

`func (o *AdminWebSessionOidcConfiguration) GetEnableRefreshUserOk() (*bool, bool)`

GetEnableRefreshUserOk returns a tuple with the EnableRefreshUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableRefreshUser

`func (o *AdminWebSessionOidcConfiguration) SetEnableRefreshUser(v bool)`

SetEnableRefreshUser sets EnableRefreshUser field to given value.

### HasEnableRefreshUser

`func (o *AdminWebSessionOidcConfiguration) HasEnableRefreshUser() bool`

HasEnableRefreshUser returns a boolean if a field has been set.

### GetScopes

`func (o *AdminWebSessionOidcConfiguration) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *AdminWebSessionOidcConfiguration) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *AdminWebSessionOidcConfiguration) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *AdminWebSessionOidcConfiguration) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetPfsessionStateCacheInSeconds

`func (o *AdminWebSessionOidcConfiguration) GetPfsessionStateCacheInSeconds() int32`

GetPfsessionStateCacheInSeconds returns the PfsessionStateCacheInSeconds field if non-nil, zero value otherwise.

### GetPfsessionStateCacheInSecondsOk

`func (o *AdminWebSessionOidcConfiguration) GetPfsessionStateCacheInSecondsOk() (*int32, bool)`

GetPfsessionStateCacheInSecondsOk returns a tuple with the PfsessionStateCacheInSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPfsessionStateCacheInSeconds

`func (o *AdminWebSessionOidcConfiguration) SetPfsessionStateCacheInSeconds(v int32)`

SetPfsessionStateCacheInSeconds sets PfsessionStateCacheInSeconds field to given value.

### HasPfsessionStateCacheInSeconds

`func (o *AdminWebSessionOidcConfiguration) HasPfsessionStateCacheInSeconds() bool`

HasPfsessionStateCacheInSeconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


