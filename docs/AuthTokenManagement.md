# AuthTokenManagement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeyRollEnabled** | Pointer to **bool** | This field is true if key rollover is enabled. When false, PingAccess will not rollover keys at the configured interval. | [optional] 
**KeyRollPeriodInHours** | Pointer to **int64** | The interval (in hours) at which PingAccess will roll the keys. Key rollover updates keys at regular intervals to ensure the security of signed auth tokens. | [optional] 
**Issuer** | Pointer to **string** | The issuer value to include in auth tokens. PingAccess inserts this value as the iss claim within the auth tokens. | [optional] 
**SigningAlgorithm** | Pointer to **string** | The signing algorithm used when creating signed auth tokens. | [optional] 

## Methods

### NewAuthTokenManagement

`func NewAuthTokenManagement() *AuthTokenManagement`

NewAuthTokenManagement instantiates a new AuthTokenManagement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthTokenManagementWithDefaults

`func NewAuthTokenManagementWithDefaults() *AuthTokenManagement`

NewAuthTokenManagementWithDefaults instantiates a new AuthTokenManagement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeyRollEnabled

`func (o *AuthTokenManagement) GetKeyRollEnabled() bool`

GetKeyRollEnabled returns the KeyRollEnabled field if non-nil, zero value otherwise.

### GetKeyRollEnabledOk

`func (o *AuthTokenManagement) GetKeyRollEnabledOk() (*bool, bool)`

GetKeyRollEnabledOk returns a tuple with the KeyRollEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyRollEnabled

`func (o *AuthTokenManagement) SetKeyRollEnabled(v bool)`

SetKeyRollEnabled sets KeyRollEnabled field to given value.

### HasKeyRollEnabled

`func (o *AuthTokenManagement) HasKeyRollEnabled() bool`

HasKeyRollEnabled returns a boolean if a field has been set.

### GetKeyRollPeriodInHours

`func (o *AuthTokenManagement) GetKeyRollPeriodInHours() int64`

GetKeyRollPeriodInHours returns the KeyRollPeriodInHours field if non-nil, zero value otherwise.

### GetKeyRollPeriodInHoursOk

`func (o *AuthTokenManagement) GetKeyRollPeriodInHoursOk() (*int64, bool)`

GetKeyRollPeriodInHoursOk returns a tuple with the KeyRollPeriodInHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyRollPeriodInHours

`func (o *AuthTokenManagement) SetKeyRollPeriodInHours(v int64)`

SetKeyRollPeriodInHours sets KeyRollPeriodInHours field to given value.

### HasKeyRollPeriodInHours

`func (o *AuthTokenManagement) HasKeyRollPeriodInHours() bool`

HasKeyRollPeriodInHours returns a boolean if a field has been set.

### GetIssuer

`func (o *AuthTokenManagement) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *AuthTokenManagement) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *AuthTokenManagement) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *AuthTokenManagement) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetSigningAlgorithm

`func (o *AuthTokenManagement) GetSigningAlgorithm() string`

GetSigningAlgorithm returns the SigningAlgorithm field if non-nil, zero value otherwise.

### GetSigningAlgorithmOk

`func (o *AuthTokenManagement) GetSigningAlgorithmOk() (*string, bool)`

GetSigningAlgorithmOk returns a tuple with the SigningAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningAlgorithm

`func (o *AuthTokenManagement) SetSigningAlgorithm(v string)`

SetSigningAlgorithm sets SigningAlgorithm field to given value.

### HasSigningAlgorithm

`func (o *AuthTokenManagement) HasSigningAlgorithm() bool`

HasSigningAlgorithm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


