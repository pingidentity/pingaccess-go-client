# PingFederateAccessToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientCredentials** | Pointer to [**OAuthClientCredentials**](OAuthClientCredentials.md) |  | [optional] 
**CacheTokens** | Pointer to **bool** | Enable to retain token details for subsequent requests. | [optional] 
**TokenTimeToLiveSeconds** | Pointer to **int64** | Defines the number of seconds to cache the access token. -1 means no limit. This value should be less than the PingFederate Token Lifetime. | [optional] 
**SubjectAttributeName** | **string** | The attribute you want to use from the OAuth access token as the subject for auditing purposes. | 
**SendAudience** | Pointer to **bool** | Enable to send the URI the user requested as the &#39;aud&#39; OAuth parameter for PingAccess to use to select an Access Token Manager. | [optional] 
**UseTokenIntrospection** | Pointer to **bool** | Specify if token introspection is enabled. | [optional] 
**AccessValidatorId** | **int64** | The Access Validator Id. This field is read-only. | 
**Name** | Pointer to **string** | The unique Access Validator name. This field is read-only. | [optional] 
**ClientSecret** | Pointer to [**HiddenField**](HiddenField.md) |  | [optional] 
**ClientId** | Pointer to **string** | The Client ID which PingAccess should use when requesting PingFederate to validate access tokens. The client must have Access Token Validation grant type allowed. (DEPRECATED - to be removed in a future release; please use &#39;clientCredentials&#39; instead) | [optional] 

## Methods

### NewPingFederateAccessToken

`func NewPingFederateAccessToken(subjectAttributeName string, accessValidatorId int64, ) *PingFederateAccessToken`

NewPingFederateAccessToken instantiates a new PingFederateAccessToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPingFederateAccessTokenWithDefaults

`func NewPingFederateAccessTokenWithDefaults() *PingFederateAccessToken`

NewPingFederateAccessTokenWithDefaults instantiates a new PingFederateAccessToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientCredentials

`func (o *PingFederateAccessToken) GetClientCredentials() OAuthClientCredentials`

GetClientCredentials returns the ClientCredentials field if non-nil, zero value otherwise.

### GetClientCredentialsOk

`func (o *PingFederateAccessToken) GetClientCredentialsOk() (*OAuthClientCredentials, bool)`

GetClientCredentialsOk returns a tuple with the ClientCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCredentials

`func (o *PingFederateAccessToken) SetClientCredentials(v OAuthClientCredentials)`

SetClientCredentials sets ClientCredentials field to given value.

### HasClientCredentials

`func (o *PingFederateAccessToken) HasClientCredentials() bool`

HasClientCredentials returns a boolean if a field has been set.

### GetCacheTokens

`func (o *PingFederateAccessToken) GetCacheTokens() bool`

GetCacheTokens returns the CacheTokens field if non-nil, zero value otherwise.

### GetCacheTokensOk

`func (o *PingFederateAccessToken) GetCacheTokensOk() (*bool, bool)`

GetCacheTokensOk returns a tuple with the CacheTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheTokens

`func (o *PingFederateAccessToken) SetCacheTokens(v bool)`

SetCacheTokens sets CacheTokens field to given value.

### HasCacheTokens

`func (o *PingFederateAccessToken) HasCacheTokens() bool`

HasCacheTokens returns a boolean if a field has been set.

### GetTokenTimeToLiveSeconds

`func (o *PingFederateAccessToken) GetTokenTimeToLiveSeconds() int64`

GetTokenTimeToLiveSeconds returns the TokenTimeToLiveSeconds field if non-nil, zero value otherwise.

### GetTokenTimeToLiveSecondsOk

`func (o *PingFederateAccessToken) GetTokenTimeToLiveSecondsOk() (*int64, bool)`

GetTokenTimeToLiveSecondsOk returns a tuple with the TokenTimeToLiveSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenTimeToLiveSeconds

`func (o *PingFederateAccessToken) SetTokenTimeToLiveSeconds(v int64)`

SetTokenTimeToLiveSeconds sets TokenTimeToLiveSeconds field to given value.

### HasTokenTimeToLiveSeconds

`func (o *PingFederateAccessToken) HasTokenTimeToLiveSeconds() bool`

HasTokenTimeToLiveSeconds returns a boolean if a field has been set.

### GetSubjectAttributeName

`func (o *PingFederateAccessToken) GetSubjectAttributeName() string`

GetSubjectAttributeName returns the SubjectAttributeName field if non-nil, zero value otherwise.

### GetSubjectAttributeNameOk

`func (o *PingFederateAccessToken) GetSubjectAttributeNameOk() (*string, bool)`

GetSubjectAttributeNameOk returns a tuple with the SubjectAttributeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectAttributeName

`func (o *PingFederateAccessToken) SetSubjectAttributeName(v string)`

SetSubjectAttributeName sets SubjectAttributeName field to given value.


### GetSendAudience

`func (o *PingFederateAccessToken) GetSendAudience() bool`

GetSendAudience returns the SendAudience field if non-nil, zero value otherwise.

### GetSendAudienceOk

`func (o *PingFederateAccessToken) GetSendAudienceOk() (*bool, bool)`

GetSendAudienceOk returns a tuple with the SendAudience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendAudience

`func (o *PingFederateAccessToken) SetSendAudience(v bool)`

SetSendAudience sets SendAudience field to given value.

### HasSendAudience

`func (o *PingFederateAccessToken) HasSendAudience() bool`

HasSendAudience returns a boolean if a field has been set.

### GetUseTokenIntrospection

`func (o *PingFederateAccessToken) GetUseTokenIntrospection() bool`

GetUseTokenIntrospection returns the UseTokenIntrospection field if non-nil, zero value otherwise.

### GetUseTokenIntrospectionOk

`func (o *PingFederateAccessToken) GetUseTokenIntrospectionOk() (*bool, bool)`

GetUseTokenIntrospectionOk returns a tuple with the UseTokenIntrospection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTokenIntrospection

`func (o *PingFederateAccessToken) SetUseTokenIntrospection(v bool)`

SetUseTokenIntrospection sets UseTokenIntrospection field to given value.

### HasUseTokenIntrospection

`func (o *PingFederateAccessToken) HasUseTokenIntrospection() bool`

HasUseTokenIntrospection returns a boolean if a field has been set.

### GetAccessValidatorId

`func (o *PingFederateAccessToken) GetAccessValidatorId() int64`

GetAccessValidatorId returns the AccessValidatorId field if non-nil, zero value otherwise.

### GetAccessValidatorIdOk

`func (o *PingFederateAccessToken) GetAccessValidatorIdOk() (*int64, bool)`

GetAccessValidatorIdOk returns a tuple with the AccessValidatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessValidatorId

`func (o *PingFederateAccessToken) SetAccessValidatorId(v int64)`

SetAccessValidatorId sets AccessValidatorId field to given value.


### GetName

`func (o *PingFederateAccessToken) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PingFederateAccessToken) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PingFederateAccessToken) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PingFederateAccessToken) HasName() bool`

HasName returns a boolean if a field has been set.

### GetClientSecret

`func (o *PingFederateAccessToken) GetClientSecret() HiddenField`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *PingFederateAccessToken) GetClientSecretOk() (*HiddenField, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *PingFederateAccessToken) SetClientSecret(v HiddenField)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *PingFederateAccessToken) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetClientId

`func (o *PingFederateAccessToken) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *PingFederateAccessToken) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *PingFederateAccessToken) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *PingFederateAccessToken) HasClientId() bool`

HasClientId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


