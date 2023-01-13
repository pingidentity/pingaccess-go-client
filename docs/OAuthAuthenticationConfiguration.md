# OAuthAuthenticationConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfiguredAuthorizationServerType** | [**ConfiguredAuthorizationServerType**](ConfiguredAuthorizationServerType.md) |  | 
**ClientCredentials** | [**OAuthClientCredentials**](OAuthClientCredentials.md) |  | 
**Scopes** | Pointer to **[]string** | The required scopes of validated ATs authorized to call the PingFederate administrative API. Scopes can be input as an array of case-sensitive strings. | [optional] 

## Methods

### NewOAuthAuthenticationConfiguration

`func NewOAuthAuthenticationConfiguration(configuredAuthorizationServerType ConfiguredAuthorizationServerType, clientCredentials OAuthClientCredentials, ) *OAuthAuthenticationConfiguration`

NewOAuthAuthenticationConfiguration instantiates a new OAuthAuthenticationConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuthAuthenticationConfigurationWithDefaults

`func NewOAuthAuthenticationConfigurationWithDefaults() *OAuthAuthenticationConfiguration`

NewOAuthAuthenticationConfigurationWithDefaults instantiates a new OAuthAuthenticationConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfiguredAuthorizationServerType

`func (o *OAuthAuthenticationConfiguration) GetConfiguredAuthorizationServerType() ConfiguredAuthorizationServerType`

GetConfiguredAuthorizationServerType returns the ConfiguredAuthorizationServerType field if non-nil, zero value otherwise.

### GetConfiguredAuthorizationServerTypeOk

`func (o *OAuthAuthenticationConfiguration) GetConfiguredAuthorizationServerTypeOk() (*ConfiguredAuthorizationServerType, bool)`

GetConfiguredAuthorizationServerTypeOk returns a tuple with the ConfiguredAuthorizationServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguredAuthorizationServerType

`func (o *OAuthAuthenticationConfiguration) SetConfiguredAuthorizationServerType(v ConfiguredAuthorizationServerType)`

SetConfiguredAuthorizationServerType sets ConfiguredAuthorizationServerType field to given value.


### GetClientCredentials

`func (o *OAuthAuthenticationConfiguration) GetClientCredentials() OAuthClientCredentials`

GetClientCredentials returns the ClientCredentials field if non-nil, zero value otherwise.

### GetClientCredentialsOk

`func (o *OAuthAuthenticationConfiguration) GetClientCredentialsOk() (*OAuthClientCredentials, bool)`

GetClientCredentialsOk returns a tuple with the ClientCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCredentials

`func (o *OAuthAuthenticationConfiguration) SetClientCredentials(v OAuthClientCredentials)`

SetClientCredentials sets ClientCredentials field to given value.


### GetScopes

`func (o *OAuthAuthenticationConfiguration) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *OAuthAuthenticationConfiguration) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *OAuthAuthenticationConfiguration) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *OAuthAuthenticationConfiguration) HasScopes() bool`

HasScopes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


