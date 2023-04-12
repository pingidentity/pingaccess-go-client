# AdminBasicWebSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SessionPollIntervalInSeconds** | **int32** | The interval between UI polling for session validity. | 
**ExpirationWarningInMinutes** | **int32** | The time to have the UI display a warning before the session expires. | 
**SessionTimeoutInMinutes** | **int32** | The length of time you want the PA Admin Token to remain active. Once the PA Admin Token expires, an authenticated user must re-authenticate. | 
**IdleTimeoutInMinutes** | **int32** | The length of time you want the PingAccess Admin Token to remain active when no activity is detected. | 
**Audience** | **string** | Enter a unique identifier between 1 and 32 characters that defines who the PA Admin Token is applicable to. | 
**CookieType** | [**WebSessionCookieType**](WebSessionCookieType.md) |  | 
**TimeoutGroovyScript** | **string** |  | 
**CookieDomain** | Pointer to **string** | The domain where the cookie is stored--for example, corp.yourcompany.com. | [optional] 

## Methods

### NewAdminBasicWebSession

`func NewAdminBasicWebSession(sessionPollIntervalInSeconds int32, expirationWarningInMinutes int32, sessionTimeoutInMinutes int32, idleTimeoutInMinutes int32, audience string, cookieType WebSessionCookieType, timeoutGroovyScript string, ) *AdminBasicWebSession`

NewAdminBasicWebSession instantiates a new AdminBasicWebSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminBasicWebSessionWithDefaults

`func NewAdminBasicWebSessionWithDefaults() *AdminBasicWebSession`

NewAdminBasicWebSessionWithDefaults instantiates a new AdminBasicWebSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSessionPollIntervalInSeconds

`func (o *AdminBasicWebSession) GetSessionPollIntervalInSeconds() int32`

GetSessionPollIntervalInSeconds returns the SessionPollIntervalInSeconds field if non-nil, zero value otherwise.

### GetSessionPollIntervalInSecondsOk

`func (o *AdminBasicWebSession) GetSessionPollIntervalInSecondsOk() (*int32, bool)`

GetSessionPollIntervalInSecondsOk returns a tuple with the SessionPollIntervalInSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionPollIntervalInSeconds

`func (o *AdminBasicWebSession) SetSessionPollIntervalInSeconds(v int32)`

SetSessionPollIntervalInSeconds sets SessionPollIntervalInSeconds field to given value.


### GetExpirationWarningInMinutes

`func (o *AdminBasicWebSession) GetExpirationWarningInMinutes() int32`

GetExpirationWarningInMinutes returns the ExpirationWarningInMinutes field if non-nil, zero value otherwise.

### GetExpirationWarningInMinutesOk

`func (o *AdminBasicWebSession) GetExpirationWarningInMinutesOk() (*int32, bool)`

GetExpirationWarningInMinutesOk returns a tuple with the ExpirationWarningInMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationWarningInMinutes

`func (o *AdminBasicWebSession) SetExpirationWarningInMinutes(v int32)`

SetExpirationWarningInMinutes sets ExpirationWarningInMinutes field to given value.


### GetSessionTimeoutInMinutes

`func (o *AdminBasicWebSession) GetSessionTimeoutInMinutes() int32`

GetSessionTimeoutInMinutes returns the SessionTimeoutInMinutes field if non-nil, zero value otherwise.

### GetSessionTimeoutInMinutesOk

`func (o *AdminBasicWebSession) GetSessionTimeoutInMinutesOk() (*int32, bool)`

GetSessionTimeoutInMinutesOk returns a tuple with the SessionTimeoutInMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionTimeoutInMinutes

`func (o *AdminBasicWebSession) SetSessionTimeoutInMinutes(v int32)`

SetSessionTimeoutInMinutes sets SessionTimeoutInMinutes field to given value.


### GetIdleTimeoutInMinutes

`func (o *AdminBasicWebSession) GetIdleTimeoutInMinutes() int32`

GetIdleTimeoutInMinutes returns the IdleTimeoutInMinutes field if non-nil, zero value otherwise.

### GetIdleTimeoutInMinutesOk

`func (o *AdminBasicWebSession) GetIdleTimeoutInMinutesOk() (*int32, bool)`

GetIdleTimeoutInMinutesOk returns a tuple with the IdleTimeoutInMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleTimeoutInMinutes

`func (o *AdminBasicWebSession) SetIdleTimeoutInMinutes(v int32)`

SetIdleTimeoutInMinutes sets IdleTimeoutInMinutes field to given value.


### GetAudience

`func (o *AdminBasicWebSession) GetAudience() string`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *AdminBasicWebSession) GetAudienceOk() (*string, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *AdminBasicWebSession) SetAudience(v string)`

SetAudience sets Audience field to given value.


### GetCookieType

`func (o *AdminBasicWebSession) GetCookieType() WebSessionCookieType`

GetCookieType returns the CookieType field if non-nil, zero value otherwise.

### GetCookieTypeOk

`func (o *AdminBasicWebSession) GetCookieTypeOk() (*WebSessionCookieType, bool)`

GetCookieTypeOk returns a tuple with the CookieType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCookieType

`func (o *AdminBasicWebSession) SetCookieType(v WebSessionCookieType)`

SetCookieType sets CookieType field to given value.


### GetTimeoutGroovyScript

`func (o *AdminBasicWebSession) GetTimeoutGroovyScript() string`

GetTimeoutGroovyScript returns the TimeoutGroovyScript field if non-nil, zero value otherwise.

### GetTimeoutGroovyScriptOk

`func (o *AdminBasicWebSession) GetTimeoutGroovyScriptOk() (*string, bool)`

GetTimeoutGroovyScriptOk returns a tuple with the TimeoutGroovyScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutGroovyScript

`func (o *AdminBasicWebSession) SetTimeoutGroovyScript(v string)`

SetTimeoutGroovyScript sets TimeoutGroovyScript field to given value.


### GetCookieDomain

`func (o *AdminBasicWebSession) GetCookieDomain() string`

GetCookieDomain returns the CookieDomain field if non-nil, zero value otherwise.

### GetCookieDomainOk

`func (o *AdminBasicWebSession) GetCookieDomainOk() (*string, bool)`

GetCookieDomainOk returns a tuple with the CookieDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCookieDomain

`func (o *AdminBasicWebSession) SetCookieDomain(v string)`

SetCookieDomain sets CookieDomain field to given value.

### HasCookieDomain

`func (o *AdminBasicWebSession) HasCookieDomain() bool`

HasCookieDomain returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


