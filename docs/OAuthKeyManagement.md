# OAuthKeyManagement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeyRollEnabled** | Pointer to **bool** | This field is true if key rollover is enabled. When false, PingAccess will not rollover keys at the configured interval. | [optional] 
**KeyRollPeriodInHours** | Pointer to **int64** | The interval (in hours) at which PingAccess will roll the keys. Key rollover updates keys at regular intervals to ensure the security of encrypted OAuth access tokens and encrypted OIDC id_tokens. | [optional] 
**SigningAlgorithm** | Pointer to **string** | The signing algorithm used when creating tokens for private key JWT OAuth client authentication. When set to null or empty, the algorithm will be selected from the OpenID Provider metadata. | [optional] 

## Methods

### NewOAuthKeyManagement

`func NewOAuthKeyManagement() *OAuthKeyManagement`

NewOAuthKeyManagement instantiates a new OAuthKeyManagement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuthKeyManagementWithDefaults

`func NewOAuthKeyManagementWithDefaults() *OAuthKeyManagement`

NewOAuthKeyManagementWithDefaults instantiates a new OAuthKeyManagement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeyRollEnabled

`func (o *OAuthKeyManagement) GetKeyRollEnabled() bool`

GetKeyRollEnabled returns the KeyRollEnabled field if non-nil, zero value otherwise.

### GetKeyRollEnabledOk

`func (o *OAuthKeyManagement) GetKeyRollEnabledOk() (*bool, bool)`

GetKeyRollEnabledOk returns a tuple with the KeyRollEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyRollEnabled

`func (o *OAuthKeyManagement) SetKeyRollEnabled(v bool)`

SetKeyRollEnabled sets KeyRollEnabled field to given value.

### HasKeyRollEnabled

`func (o *OAuthKeyManagement) HasKeyRollEnabled() bool`

HasKeyRollEnabled returns a boolean if a field has been set.

### GetKeyRollPeriodInHours

`func (o *OAuthKeyManagement) GetKeyRollPeriodInHours() int64`

GetKeyRollPeriodInHours returns the KeyRollPeriodInHours field if non-nil, zero value otherwise.

### GetKeyRollPeriodInHoursOk

`func (o *OAuthKeyManagement) GetKeyRollPeriodInHoursOk() (*int64, bool)`

GetKeyRollPeriodInHoursOk returns a tuple with the KeyRollPeriodInHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyRollPeriodInHours

`func (o *OAuthKeyManagement) SetKeyRollPeriodInHours(v int64)`

SetKeyRollPeriodInHours sets KeyRollPeriodInHours field to given value.

### HasKeyRollPeriodInHours

`func (o *OAuthKeyManagement) HasKeyRollPeriodInHours() bool`

HasKeyRollPeriodInHours returns a boolean if a field has been set.

### GetSigningAlgorithm

`func (o *OAuthKeyManagement) GetSigningAlgorithm() string`

GetSigningAlgorithm returns the SigningAlgorithm field if non-nil, zero value otherwise.

### GetSigningAlgorithmOk

`func (o *OAuthKeyManagement) GetSigningAlgorithmOk() (*string, bool)`

GetSigningAlgorithmOk returns a tuple with the SigningAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningAlgorithm

`func (o *OAuthKeyManagement) SetSigningAlgorithm(v string)`

SetSigningAlgorithm sets SigningAlgorithm field to given value.

### HasSigningAlgorithm

`func (o *OAuthKeyManagement) HasSigningAlgorithm() bool`

HasSigningAlgorithm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


