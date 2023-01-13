# OAuthClientCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | **string** | Specify the client ID. | 
**ClientSecret** | Pointer to [**HiddenField**](HiddenField.md) |  | [optional] 
**KeyPairId** | Pointer to **int32** | Specify the ID of a key pair to use for mutual TLS. | [optional] 
**CredentialsType** | Pointer to [**CredentialsType**](CredentialsType.md) |  | [optional] 

## Methods

### NewOAuthClientCredentials

`func NewOAuthClientCredentials(clientId string, ) *OAuthClientCredentials`

NewOAuthClientCredentials instantiates a new OAuthClientCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuthClientCredentialsWithDefaults

`func NewOAuthClientCredentialsWithDefaults() *OAuthClientCredentials`

NewOAuthClientCredentialsWithDefaults instantiates a new OAuthClientCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *OAuthClientCredentials) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *OAuthClientCredentials) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *OAuthClientCredentials) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientSecret

`func (o *OAuthClientCredentials) GetClientSecret() HiddenField`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *OAuthClientCredentials) GetClientSecretOk() (*HiddenField, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *OAuthClientCredentials) SetClientSecret(v HiddenField)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *OAuthClientCredentials) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetKeyPairId

`func (o *OAuthClientCredentials) GetKeyPairId() int32`

GetKeyPairId returns the KeyPairId field if non-nil, zero value otherwise.

### GetKeyPairIdOk

`func (o *OAuthClientCredentials) GetKeyPairIdOk() (*int32, bool)`

GetKeyPairIdOk returns a tuple with the KeyPairId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyPairId

`func (o *OAuthClientCredentials) SetKeyPairId(v int32)`

SetKeyPairId sets KeyPairId field to given value.

### HasKeyPairId

`func (o *OAuthClientCredentials) HasKeyPairId() bool`

HasKeyPairId returns a boolean if a field has been set.

### GetCredentialsType

`func (o *OAuthClientCredentials) GetCredentialsType() CredentialsType`

GetCredentialsType returns the CredentialsType field if non-nil, zero value otherwise.

### GetCredentialsTypeOk

`func (o *OAuthClientCredentials) GetCredentialsTypeOk() (*CredentialsType, bool)`

GetCredentialsTypeOk returns a tuple with the CredentialsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialsType

`func (o *OAuthClientCredentials) SetCredentialsType(v CredentialsType)`

SetCredentialsType sets CredentialsType field to given value.

### HasCredentialsType

`func (o *OAuthClientCredentials) HasCredentialsType() bool`

HasCredentialsType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


