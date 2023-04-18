# WebSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | When creating a new WebSession, this is the ID for the WebSession. If not specified, an ID will be automatically assigned. When updating an existing WebSession, this field is ignored and the ID is determined by the path parameter. | [optional] 
**Name** | **string** | (sortable) The name of the web session. | 
**RequestPreservationType** | Pointer to [**RequestPreservationType**](RequestPreservationType.md) |  | [optional] 
**EnablePushAuthorization** | Pointer to **bool** | Enabling will allow PA to push the payload of an OAuth 2.0 authorization request to the authorization server via a direct request and provide it with a request URI that is used as reference to the data in a subsequent call to the authorization endpoint. | [optional] 
**SessionTimeoutInMinutes** | Pointer to **int64** | (sortable) The length of time you want the PA Token to remain active. Once the PA Token expires, an authenticated user must re-authenticate. | [optional] 
**ValidateSessionIsAlive** | Pointer to **bool** | (sortable) Specify if PingAccess should validate sessions with the configured PingFederate instance during request processing. | [optional] 
**SendRequestedUrlToProvider** | Pointer to **bool** | (sortable) Specify if you want to send the requested URL as part of the authentication request to the OpenID Connect Provider. | [optional] 
**RefreshUserInfoClaimsInterval** | Pointer to **int64** | (sortable) Specify the maximum number of seconds to cache user attribute information when the Refresh User is enabled. | [optional] 
**ProvideAuthenticationFeedback** | Pointer to **bool** | (sortable) Specify if PingAccess should provide feedback to the authentication authority. For example, tell the authority that the user is being redirected because their session expired. | [optional] 
**CacheUserAttributes** | Pointer to **bool** | (sortable) Specify if PingAccess should cache user attribute information for use in policy decisions. When disabled, this data is encoded and stored in the session cookie. | [optional] 
**EnableRefreshUser** | Pointer to **bool** | (sortable) Specify if you want to have PingAccess periodically refresh user data from PingFederate for use in policy decisions. | [optional] 
**RequestProfile** | Pointer to **bool** | Specifies whether the default scopes (&#39;profile&#39;, &#39;email&#39;, &#39;address&#39;, and &#39;phone&#39;) should be specified in the access request. (DEPRECATED - to be removed in a future release; please use &#39;scopes&#39; instead) | [optional] 
**PromptParameter** | Pointer to **[]string** | Enter zero or more case sensitive string values. The values can be used by the Client to make sure that the End-User is still present for the current session or to bring attention to the request. | [optional] 
**SecureCookie** | Pointer to **bool** | (sortable) Specify whether the PingAccess Cookie must be sent using only HTTPS connections. | [optional] 
**HttpOnlyCookie** | Pointer to **bool** | (sortable) Enable the HttpOnly flag on cookies that contain the PA Token. | [optional] 
**IdleTimeoutInMinutes** | Pointer to **int64** | (sortable) The length of time you want the PingAccess Token to remain active when no activity is detected. | [optional] 
**Audience** | **string** | (sortable) Enter a unique identifier between 1 and 32 characters that defines who the PA Token is applicable to. | 
**CookieType** | Pointer to [**WebSessionCookieType**](WebSessionCookieType.md) |  | [optional] 
**SameSite** | Pointer to [**SameSiteType**](SameSiteType.md) |  | [optional] 
**Scopes** | Pointer to **[]string** | The list of scopes to be specified in the access request. If not specified, the default scopes (&#39;profile&#39;, &#39;email&#39;, &#39;address&#39;, and &#39;phone&#39;) will be used. The openid scope is implied and does not need to be specified in this list. | [optional] 
**OidcLoginType** | Pointer to [**OidcLoginType**](OidcLoginType.md) |  | [optional] 
**WebStorageType** | Pointer to [**WebStorageType**](WebStorageType.md) |  | [optional] 
**PkceChallengeType** | Pointer to [**PkceChallengeType**](PkceChallengeType.md) |  | [optional] 
**TimeoutGroovyScript** | Pointer to **string** | An optional Groovy script used to calculate override values for sessionTimeoutInMinutes and idleTimeoutInMinutes based on user attributes during token issuance. The script outputs a SessionTimeoutContext with an idle and session timeout. Values of 0 for this context will result in the existing values on the web session configuration being used, and negative values are not valid and will cause errors. | [optional] 
**ClientCredentials** | [**OAuthClientCredentials**](OAuthClientCredentials.md) |  | 
**CookieDomain** | Pointer to **string** | (sortable) The domain where the cookie is stored--for example, corp.yourcompany.com. | [optional] 
**FailOnUnsupportedPreservationContentType** | **bool** | (sortable) Specify if PingAccess should produce a 415 HTTP response when it receives an unauthenticated POST request with a content type unsupported by request preservation. The only content type supported by request preservation is application/x-www-form-urlencoded. When disabled, PingAccess will challenge an unauthenticated POST request using an unsupported content type with the same challenge response sent to an unauthenticated GET request. The default is false. | 
**PfsessionStateCacheInSeconds** | Pointer to **int64** | (sortable) Specify the number of seconds to cache PingFederate Session State information. | [optional] 

## Methods

### NewWebSession

`func NewWebSession(name string, audience string, clientCredentials OAuthClientCredentials, failOnUnsupportedPreservationContentType bool, ) *WebSession`

NewWebSession instantiates a new WebSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebSessionWithDefaults

`func NewWebSessionWithDefaults() *WebSession`

NewWebSessionWithDefaults instantiates a new WebSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebSession) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebSession) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebSession) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *WebSession) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *WebSession) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WebSession) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WebSession) SetName(v string)`

SetName sets Name field to given value.


### GetRequestPreservationType

`func (o *WebSession) GetRequestPreservationType() RequestPreservationType`

GetRequestPreservationType returns the RequestPreservationType field if non-nil, zero value otherwise.

### GetRequestPreservationTypeOk

`func (o *WebSession) GetRequestPreservationTypeOk() (*RequestPreservationType, bool)`

GetRequestPreservationTypeOk returns a tuple with the RequestPreservationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestPreservationType

`func (o *WebSession) SetRequestPreservationType(v RequestPreservationType)`

SetRequestPreservationType sets RequestPreservationType field to given value.

### HasRequestPreservationType

`func (o *WebSession) HasRequestPreservationType() bool`

HasRequestPreservationType returns a boolean if a field has been set.

### GetEnablePushAuthorization

`func (o *WebSession) GetEnablePushAuthorization() bool`

GetEnablePushAuthorization returns the EnablePushAuthorization field if non-nil, zero value otherwise.

### GetEnablePushAuthorizationOk

`func (o *WebSession) GetEnablePushAuthorizationOk() (*bool, bool)`

GetEnablePushAuthorizationOk returns a tuple with the EnablePushAuthorization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePushAuthorization

`func (o *WebSession) SetEnablePushAuthorization(v bool)`

SetEnablePushAuthorization sets EnablePushAuthorization field to given value.

### HasEnablePushAuthorization

`func (o *WebSession) HasEnablePushAuthorization() bool`

HasEnablePushAuthorization returns a boolean if a field has been set.

### GetSessionTimeoutInMinutes

`func (o *WebSession) GetSessionTimeoutInMinutes() int64`

GetSessionTimeoutInMinutes returns the SessionTimeoutInMinutes field if non-nil, zero value otherwise.

### GetSessionTimeoutInMinutesOk

`func (o *WebSession) GetSessionTimeoutInMinutesOk() (*int64, bool)`

GetSessionTimeoutInMinutesOk returns a tuple with the SessionTimeoutInMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionTimeoutInMinutes

`func (o *WebSession) SetSessionTimeoutInMinutes(v int64)`

SetSessionTimeoutInMinutes sets SessionTimeoutInMinutes field to given value.

### HasSessionTimeoutInMinutes

`func (o *WebSession) HasSessionTimeoutInMinutes() bool`

HasSessionTimeoutInMinutes returns a boolean if a field has been set.

### GetValidateSessionIsAlive

`func (o *WebSession) GetValidateSessionIsAlive() bool`

GetValidateSessionIsAlive returns the ValidateSessionIsAlive field if non-nil, zero value otherwise.

### GetValidateSessionIsAliveOk

`func (o *WebSession) GetValidateSessionIsAliveOk() (*bool, bool)`

GetValidateSessionIsAliveOk returns a tuple with the ValidateSessionIsAlive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidateSessionIsAlive

`func (o *WebSession) SetValidateSessionIsAlive(v bool)`

SetValidateSessionIsAlive sets ValidateSessionIsAlive field to given value.

### HasValidateSessionIsAlive

`func (o *WebSession) HasValidateSessionIsAlive() bool`

HasValidateSessionIsAlive returns a boolean if a field has been set.

### GetSendRequestedUrlToProvider

`func (o *WebSession) GetSendRequestedUrlToProvider() bool`

GetSendRequestedUrlToProvider returns the SendRequestedUrlToProvider field if non-nil, zero value otherwise.

### GetSendRequestedUrlToProviderOk

`func (o *WebSession) GetSendRequestedUrlToProviderOk() (*bool, bool)`

GetSendRequestedUrlToProviderOk returns a tuple with the SendRequestedUrlToProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendRequestedUrlToProvider

`func (o *WebSession) SetSendRequestedUrlToProvider(v bool)`

SetSendRequestedUrlToProvider sets SendRequestedUrlToProvider field to given value.

### HasSendRequestedUrlToProvider

`func (o *WebSession) HasSendRequestedUrlToProvider() bool`

HasSendRequestedUrlToProvider returns a boolean if a field has been set.

### GetRefreshUserInfoClaimsInterval

`func (o *WebSession) GetRefreshUserInfoClaimsInterval() int64`

GetRefreshUserInfoClaimsInterval returns the RefreshUserInfoClaimsInterval field if non-nil, zero value otherwise.

### GetRefreshUserInfoClaimsIntervalOk

`func (o *WebSession) GetRefreshUserInfoClaimsIntervalOk() (*int64, bool)`

GetRefreshUserInfoClaimsIntervalOk returns a tuple with the RefreshUserInfoClaimsInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshUserInfoClaimsInterval

`func (o *WebSession) SetRefreshUserInfoClaimsInterval(v int64)`

SetRefreshUserInfoClaimsInterval sets RefreshUserInfoClaimsInterval field to given value.

### HasRefreshUserInfoClaimsInterval

`func (o *WebSession) HasRefreshUserInfoClaimsInterval() bool`

HasRefreshUserInfoClaimsInterval returns a boolean if a field has been set.

### GetProvideAuthenticationFeedback

`func (o *WebSession) GetProvideAuthenticationFeedback() bool`

GetProvideAuthenticationFeedback returns the ProvideAuthenticationFeedback field if non-nil, zero value otherwise.

### GetProvideAuthenticationFeedbackOk

`func (o *WebSession) GetProvideAuthenticationFeedbackOk() (*bool, bool)`

GetProvideAuthenticationFeedbackOk returns a tuple with the ProvideAuthenticationFeedback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvideAuthenticationFeedback

`func (o *WebSession) SetProvideAuthenticationFeedback(v bool)`

SetProvideAuthenticationFeedback sets ProvideAuthenticationFeedback field to given value.

### HasProvideAuthenticationFeedback

`func (o *WebSession) HasProvideAuthenticationFeedback() bool`

HasProvideAuthenticationFeedback returns a boolean if a field has been set.

### GetCacheUserAttributes

`func (o *WebSession) GetCacheUserAttributes() bool`

GetCacheUserAttributes returns the CacheUserAttributes field if non-nil, zero value otherwise.

### GetCacheUserAttributesOk

`func (o *WebSession) GetCacheUserAttributesOk() (*bool, bool)`

GetCacheUserAttributesOk returns a tuple with the CacheUserAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheUserAttributes

`func (o *WebSession) SetCacheUserAttributes(v bool)`

SetCacheUserAttributes sets CacheUserAttributes field to given value.

### HasCacheUserAttributes

`func (o *WebSession) HasCacheUserAttributes() bool`

HasCacheUserAttributes returns a boolean if a field has been set.

### GetEnableRefreshUser

`func (o *WebSession) GetEnableRefreshUser() bool`

GetEnableRefreshUser returns the EnableRefreshUser field if non-nil, zero value otherwise.

### GetEnableRefreshUserOk

`func (o *WebSession) GetEnableRefreshUserOk() (*bool, bool)`

GetEnableRefreshUserOk returns a tuple with the EnableRefreshUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableRefreshUser

`func (o *WebSession) SetEnableRefreshUser(v bool)`

SetEnableRefreshUser sets EnableRefreshUser field to given value.

### HasEnableRefreshUser

`func (o *WebSession) HasEnableRefreshUser() bool`

HasEnableRefreshUser returns a boolean if a field has been set.

### GetRequestProfile

`func (o *WebSession) GetRequestProfile() bool`

GetRequestProfile returns the RequestProfile field if non-nil, zero value otherwise.

### GetRequestProfileOk

`func (o *WebSession) GetRequestProfileOk() (*bool, bool)`

GetRequestProfileOk returns a tuple with the RequestProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestProfile

`func (o *WebSession) SetRequestProfile(v bool)`

SetRequestProfile sets RequestProfile field to given value.

### HasRequestProfile

`func (o *WebSession) HasRequestProfile() bool`

HasRequestProfile returns a boolean if a field has been set.

### GetPromptParameter

`func (o *WebSession) GetPromptParameter() []string`

GetPromptParameter returns the PromptParameter field if non-nil, zero value otherwise.

### GetPromptParameterOk

`func (o *WebSession) GetPromptParameterOk() (*[]string, bool)`

GetPromptParameterOk returns a tuple with the PromptParameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptParameter

`func (o *WebSession) SetPromptParameter(v []string)`

SetPromptParameter sets PromptParameter field to given value.

### HasPromptParameter

`func (o *WebSession) HasPromptParameter() bool`

HasPromptParameter returns a boolean if a field has been set.

### GetSecureCookie

`func (o *WebSession) GetSecureCookie() bool`

GetSecureCookie returns the SecureCookie field if non-nil, zero value otherwise.

### GetSecureCookieOk

`func (o *WebSession) GetSecureCookieOk() (*bool, bool)`

GetSecureCookieOk returns a tuple with the SecureCookie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureCookie

`func (o *WebSession) SetSecureCookie(v bool)`

SetSecureCookie sets SecureCookie field to given value.

### HasSecureCookie

`func (o *WebSession) HasSecureCookie() bool`

HasSecureCookie returns a boolean if a field has been set.

### GetHttpOnlyCookie

`func (o *WebSession) GetHttpOnlyCookie() bool`

GetHttpOnlyCookie returns the HttpOnlyCookie field if non-nil, zero value otherwise.

### GetHttpOnlyCookieOk

`func (o *WebSession) GetHttpOnlyCookieOk() (*bool, bool)`

GetHttpOnlyCookieOk returns a tuple with the HttpOnlyCookie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpOnlyCookie

`func (o *WebSession) SetHttpOnlyCookie(v bool)`

SetHttpOnlyCookie sets HttpOnlyCookie field to given value.

### HasHttpOnlyCookie

`func (o *WebSession) HasHttpOnlyCookie() bool`

HasHttpOnlyCookie returns a boolean if a field has been set.

### GetIdleTimeoutInMinutes

`func (o *WebSession) GetIdleTimeoutInMinutes() int64`

GetIdleTimeoutInMinutes returns the IdleTimeoutInMinutes field if non-nil, zero value otherwise.

### GetIdleTimeoutInMinutesOk

`func (o *WebSession) GetIdleTimeoutInMinutesOk() (*int64, bool)`

GetIdleTimeoutInMinutesOk returns a tuple with the IdleTimeoutInMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleTimeoutInMinutes

`func (o *WebSession) SetIdleTimeoutInMinutes(v int64)`

SetIdleTimeoutInMinutes sets IdleTimeoutInMinutes field to given value.

### HasIdleTimeoutInMinutes

`func (o *WebSession) HasIdleTimeoutInMinutes() bool`

HasIdleTimeoutInMinutes returns a boolean if a field has been set.

### GetAudience

`func (o *WebSession) GetAudience() string`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *WebSession) GetAudienceOk() (*string, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *WebSession) SetAudience(v string)`

SetAudience sets Audience field to given value.


### GetCookieType

`func (o *WebSession) GetCookieType() WebSessionCookieType`

GetCookieType returns the CookieType field if non-nil, zero value otherwise.

### GetCookieTypeOk

`func (o *WebSession) GetCookieTypeOk() (*WebSessionCookieType, bool)`

GetCookieTypeOk returns a tuple with the CookieType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCookieType

`func (o *WebSession) SetCookieType(v WebSessionCookieType)`

SetCookieType sets CookieType field to given value.

### HasCookieType

`func (o *WebSession) HasCookieType() bool`

HasCookieType returns a boolean if a field has been set.

### GetSameSite

`func (o *WebSession) GetSameSite() SameSiteType`

GetSameSite returns the SameSite field if non-nil, zero value otherwise.

### GetSameSiteOk

`func (o *WebSession) GetSameSiteOk() (*SameSiteType, bool)`

GetSameSiteOk returns a tuple with the SameSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSameSite

`func (o *WebSession) SetSameSite(v SameSiteType)`

SetSameSite sets SameSite field to given value.

### HasSameSite

`func (o *WebSession) HasSameSite() bool`

HasSameSite returns a boolean if a field has been set.

### GetScopes

`func (o *WebSession) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *WebSession) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *WebSession) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *WebSession) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetOidcLoginType

`func (o *WebSession) GetOidcLoginType() OidcLoginType`

GetOidcLoginType returns the OidcLoginType field if non-nil, zero value otherwise.

### GetOidcLoginTypeOk

`func (o *WebSession) GetOidcLoginTypeOk() (*OidcLoginType, bool)`

GetOidcLoginTypeOk returns a tuple with the OidcLoginType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcLoginType

`func (o *WebSession) SetOidcLoginType(v OidcLoginType)`

SetOidcLoginType sets OidcLoginType field to given value.

### HasOidcLoginType

`func (o *WebSession) HasOidcLoginType() bool`

HasOidcLoginType returns a boolean if a field has been set.

### GetWebStorageType

`func (o *WebSession) GetWebStorageType() WebStorageType`

GetWebStorageType returns the WebStorageType field if non-nil, zero value otherwise.

### GetWebStorageTypeOk

`func (o *WebSession) GetWebStorageTypeOk() (*WebStorageType, bool)`

GetWebStorageTypeOk returns a tuple with the WebStorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebStorageType

`func (o *WebSession) SetWebStorageType(v WebStorageType)`

SetWebStorageType sets WebStorageType field to given value.

### HasWebStorageType

`func (o *WebSession) HasWebStorageType() bool`

HasWebStorageType returns a boolean if a field has been set.

### GetPkceChallengeType

`func (o *WebSession) GetPkceChallengeType() PkceChallengeType`

GetPkceChallengeType returns the PkceChallengeType field if non-nil, zero value otherwise.

### GetPkceChallengeTypeOk

`func (o *WebSession) GetPkceChallengeTypeOk() (*PkceChallengeType, bool)`

GetPkceChallengeTypeOk returns a tuple with the PkceChallengeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkceChallengeType

`func (o *WebSession) SetPkceChallengeType(v PkceChallengeType)`

SetPkceChallengeType sets PkceChallengeType field to given value.

### HasPkceChallengeType

`func (o *WebSession) HasPkceChallengeType() bool`

HasPkceChallengeType returns a boolean if a field has been set.

### GetTimeoutGroovyScript

`func (o *WebSession) GetTimeoutGroovyScript() string`

GetTimeoutGroovyScript returns the TimeoutGroovyScript field if non-nil, zero value otherwise.

### GetTimeoutGroovyScriptOk

`func (o *WebSession) GetTimeoutGroovyScriptOk() (*string, bool)`

GetTimeoutGroovyScriptOk returns a tuple with the TimeoutGroovyScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutGroovyScript

`func (o *WebSession) SetTimeoutGroovyScript(v string)`

SetTimeoutGroovyScript sets TimeoutGroovyScript field to given value.

### HasTimeoutGroovyScript

`func (o *WebSession) HasTimeoutGroovyScript() bool`

HasTimeoutGroovyScript returns a boolean if a field has been set.

### GetClientCredentials

`func (o *WebSession) GetClientCredentials() OAuthClientCredentials`

GetClientCredentials returns the ClientCredentials field if non-nil, zero value otherwise.

### GetClientCredentialsOk

`func (o *WebSession) GetClientCredentialsOk() (*OAuthClientCredentials, bool)`

GetClientCredentialsOk returns a tuple with the ClientCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCredentials

`func (o *WebSession) SetClientCredentials(v OAuthClientCredentials)`

SetClientCredentials sets ClientCredentials field to given value.


### GetCookieDomain

`func (o *WebSession) GetCookieDomain() string`

GetCookieDomain returns the CookieDomain field if non-nil, zero value otherwise.

### GetCookieDomainOk

`func (o *WebSession) GetCookieDomainOk() (*string, bool)`

GetCookieDomainOk returns a tuple with the CookieDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCookieDomain

`func (o *WebSession) SetCookieDomain(v string)`

SetCookieDomain sets CookieDomain field to given value.

### HasCookieDomain

`func (o *WebSession) HasCookieDomain() bool`

HasCookieDomain returns a boolean if a field has been set.

### GetFailOnUnsupportedPreservationContentType

`func (o *WebSession) GetFailOnUnsupportedPreservationContentType() bool`

GetFailOnUnsupportedPreservationContentType returns the FailOnUnsupportedPreservationContentType field if non-nil, zero value otherwise.

### GetFailOnUnsupportedPreservationContentTypeOk

`func (o *WebSession) GetFailOnUnsupportedPreservationContentTypeOk() (*bool, bool)`

GetFailOnUnsupportedPreservationContentTypeOk returns a tuple with the FailOnUnsupportedPreservationContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailOnUnsupportedPreservationContentType

`func (o *WebSession) SetFailOnUnsupportedPreservationContentType(v bool)`

SetFailOnUnsupportedPreservationContentType sets FailOnUnsupportedPreservationContentType field to given value.


### GetPfsessionStateCacheInSeconds

`func (o *WebSession) GetPfsessionStateCacheInSeconds() int64`

GetPfsessionStateCacheInSeconds returns the PfsessionStateCacheInSeconds field if non-nil, zero value otherwise.

### GetPfsessionStateCacheInSecondsOk

`func (o *WebSession) GetPfsessionStateCacheInSecondsOk() (*int64, bool)`

GetPfsessionStateCacheInSecondsOk returns a tuple with the PfsessionStateCacheInSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPfsessionStateCacheInSeconds

`func (o *WebSession) SetPfsessionStateCacheInSeconds(v int64)`

SetPfsessionStateCacheInSeconds sets PfsessionStateCacheInSeconds field to given value.

### HasPfsessionStateCacheInSeconds

`func (o *WebSession) HasPfsessionStateCacheInSeconds() bool`

HasPfsessionStateCacheInSeconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


