# OAuthConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientCredentials** | Pointer to [**OAuthClientCredentials**](OAuthClientCredentials.md) |  | [optional] 
**Scope** | **string** | The scope required to successfully access the API--for example, admin. | 
**SubjectAttributeName** | Pointer to **string** | The attribute you want to use from the OAuth access token as the subject for auditing purposes. At runtime, the attribute&#39;s value is used as the subject field in audit log entries for the Admin API. This field is ignored when the accessTokenValidator is configured. | [optional] 
**RoleMapping** | Pointer to [**RoleMappingConfiguration**](RoleMappingConfiguration.md) |  | [optional] 
**Enabled** | Pointer to **bool** | This field is true if OAuth authentication to the Administrative API is enabled. | [optional] 
**AccessTokenValidator** | Pointer to [**EmbeddableAccessTokenValidator**](EmbeddableAccessTokenValidator.md) |  | [optional] 
**ClientId** | Pointer to **string** | The client_id of the OAuth client used for validating OAuth access tokens. (DEPRECATED - to be removed in a future release; please use &#39;clientCredentials&#39; instead) | [optional] 
**ClientSecret** | Pointer to [**HiddenField**](HiddenField.md) |  | [optional] 

## Methods

### NewOAuthConfig

`func NewOAuthConfig(scope string, ) *OAuthConfig`

NewOAuthConfig instantiates a new OAuthConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuthConfigWithDefaults

`func NewOAuthConfigWithDefaults() *OAuthConfig`

NewOAuthConfigWithDefaults instantiates a new OAuthConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientCredentials

`func (o *OAuthConfig) GetClientCredentials() OAuthClientCredentials`

GetClientCredentials returns the ClientCredentials field if non-nil, zero value otherwise.

### GetClientCredentialsOk

`func (o *OAuthConfig) GetClientCredentialsOk() (*OAuthClientCredentials, bool)`

GetClientCredentialsOk returns a tuple with the ClientCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCredentials

`func (o *OAuthConfig) SetClientCredentials(v OAuthClientCredentials)`

SetClientCredentials sets ClientCredentials field to given value.

### HasClientCredentials

`func (o *OAuthConfig) HasClientCredentials() bool`

HasClientCredentials returns a boolean if a field has been set.

### GetScope

`func (o *OAuthConfig) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *OAuthConfig) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *OAuthConfig) SetScope(v string)`

SetScope sets Scope field to given value.


### GetSubjectAttributeName

`func (o *OAuthConfig) GetSubjectAttributeName() string`

GetSubjectAttributeName returns the SubjectAttributeName field if non-nil, zero value otherwise.

### GetSubjectAttributeNameOk

`func (o *OAuthConfig) GetSubjectAttributeNameOk() (*string, bool)`

GetSubjectAttributeNameOk returns a tuple with the SubjectAttributeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectAttributeName

`func (o *OAuthConfig) SetSubjectAttributeName(v string)`

SetSubjectAttributeName sets SubjectAttributeName field to given value.

### HasSubjectAttributeName

`func (o *OAuthConfig) HasSubjectAttributeName() bool`

HasSubjectAttributeName returns a boolean if a field has been set.

### GetRoleMapping

`func (o *OAuthConfig) GetRoleMapping() RoleMappingConfiguration`

GetRoleMapping returns the RoleMapping field if non-nil, zero value otherwise.

### GetRoleMappingOk

`func (o *OAuthConfig) GetRoleMappingOk() (*RoleMappingConfiguration, bool)`

GetRoleMappingOk returns a tuple with the RoleMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleMapping

`func (o *OAuthConfig) SetRoleMapping(v RoleMappingConfiguration)`

SetRoleMapping sets RoleMapping field to given value.

### HasRoleMapping

`func (o *OAuthConfig) HasRoleMapping() bool`

HasRoleMapping returns a boolean if a field has been set.

### GetEnabled

`func (o *OAuthConfig) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *OAuthConfig) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *OAuthConfig) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *OAuthConfig) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetAccessTokenValidator

`func (o *OAuthConfig) GetAccessTokenValidator() EmbeddableAccessTokenValidator`

GetAccessTokenValidator returns the AccessTokenValidator field if non-nil, zero value otherwise.

### GetAccessTokenValidatorOk

`func (o *OAuthConfig) GetAccessTokenValidatorOk() (*EmbeddableAccessTokenValidator, bool)`

GetAccessTokenValidatorOk returns a tuple with the AccessTokenValidator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenValidator

`func (o *OAuthConfig) SetAccessTokenValidator(v EmbeddableAccessTokenValidator)`

SetAccessTokenValidator sets AccessTokenValidator field to given value.

### HasAccessTokenValidator

`func (o *OAuthConfig) HasAccessTokenValidator() bool`

HasAccessTokenValidator returns a boolean if a field has been set.

### GetClientId

`func (o *OAuthConfig) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *OAuthConfig) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *OAuthConfig) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *OAuthConfig) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *OAuthConfig) GetClientSecret() HiddenField`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *OAuthConfig) GetClientSecretOk() (*HiddenField, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *OAuthConfig) SetClientSecret(v HiddenField)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *OAuthConfig) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


