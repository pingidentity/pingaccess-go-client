# WebSessionManagement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeyRollEnabled** | Pointer to **bool** | This field is true if key rollover is enabled. When false, PingAccess will not rollover keys at the configured interval. | [optional] 
**KeyRollPeriodInHours** | Pointer to **int64** | The interval (in hours) at which PingAccess will roll the keys. Key rollover updates keys at regular intervals to ensure the security of signed and encrypted PA tokens. | [optional] 
**Issuer** | Pointer to **string** | The issuer value to include in the PA token. PingAccess inserts this value as the iss claim within the PA Token. | [optional] 
**SigningAlgorithm** | Pointer to **string** | The signing algorithm used when creating signed PA tokens and when verifying them from a user&#39;s browser. | [optional] 
**CookieName** | Pointer to **string** | The name for the browser cookie to contain the PA token. | [optional] 
**SessionStateCookieName** | Pointer to **string** | The name of the session state cookie. | [optional] 
**UpdateTokenWindowInSeconds** | Pointer to **int64** | The number of seconds before the idle timeout is updated in the PA token. | [optional] 
**EncryptionAlgorithm** | Pointer to **string** | The encryption algorithm used when creating encrypted PA tokens and when verifying them from a user&#39;s browser. | [optional] 
**NonceCookieTimeToLiveInMinutes** | Pointer to **int32** | The number of minutes that the nonce cookie is valid when multiple concurrent authentication requests are made. 0 indicates that the system default value should be used. | [optional] 

## Methods

### NewWebSessionManagement

`func NewWebSessionManagement() *WebSessionManagement`

NewWebSessionManagement instantiates a new WebSessionManagement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebSessionManagementWithDefaults

`func NewWebSessionManagementWithDefaults() *WebSessionManagement`

NewWebSessionManagementWithDefaults instantiates a new WebSessionManagement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeyRollEnabled

`func (o *WebSessionManagement) GetKeyRollEnabled() bool`

GetKeyRollEnabled returns the KeyRollEnabled field if non-nil, zero value otherwise.

### GetKeyRollEnabledOk

`func (o *WebSessionManagement) GetKeyRollEnabledOk() (*bool, bool)`

GetKeyRollEnabledOk returns a tuple with the KeyRollEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyRollEnabled

`func (o *WebSessionManagement) SetKeyRollEnabled(v bool)`

SetKeyRollEnabled sets KeyRollEnabled field to given value.

### HasKeyRollEnabled

`func (o *WebSessionManagement) HasKeyRollEnabled() bool`

HasKeyRollEnabled returns a boolean if a field has been set.

### GetKeyRollPeriodInHours

`func (o *WebSessionManagement) GetKeyRollPeriodInHours() int64`

GetKeyRollPeriodInHours returns the KeyRollPeriodInHours field if non-nil, zero value otherwise.

### GetKeyRollPeriodInHoursOk

`func (o *WebSessionManagement) GetKeyRollPeriodInHoursOk() (*int64, bool)`

GetKeyRollPeriodInHoursOk returns a tuple with the KeyRollPeriodInHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyRollPeriodInHours

`func (o *WebSessionManagement) SetKeyRollPeriodInHours(v int64)`

SetKeyRollPeriodInHours sets KeyRollPeriodInHours field to given value.

### HasKeyRollPeriodInHours

`func (o *WebSessionManagement) HasKeyRollPeriodInHours() bool`

HasKeyRollPeriodInHours returns a boolean if a field has been set.

### GetIssuer

`func (o *WebSessionManagement) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *WebSessionManagement) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *WebSessionManagement) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *WebSessionManagement) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetSigningAlgorithm

`func (o *WebSessionManagement) GetSigningAlgorithm() string`

GetSigningAlgorithm returns the SigningAlgorithm field if non-nil, zero value otherwise.

### GetSigningAlgorithmOk

`func (o *WebSessionManagement) GetSigningAlgorithmOk() (*string, bool)`

GetSigningAlgorithmOk returns a tuple with the SigningAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningAlgorithm

`func (o *WebSessionManagement) SetSigningAlgorithm(v string)`

SetSigningAlgorithm sets SigningAlgorithm field to given value.

### HasSigningAlgorithm

`func (o *WebSessionManagement) HasSigningAlgorithm() bool`

HasSigningAlgorithm returns a boolean if a field has been set.

### GetCookieName

`func (o *WebSessionManagement) GetCookieName() string`

GetCookieName returns the CookieName field if non-nil, zero value otherwise.

### GetCookieNameOk

`func (o *WebSessionManagement) GetCookieNameOk() (*string, bool)`

GetCookieNameOk returns a tuple with the CookieName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCookieName

`func (o *WebSessionManagement) SetCookieName(v string)`

SetCookieName sets CookieName field to given value.

### HasCookieName

`func (o *WebSessionManagement) HasCookieName() bool`

HasCookieName returns a boolean if a field has been set.

### GetSessionStateCookieName

`func (o *WebSessionManagement) GetSessionStateCookieName() string`

GetSessionStateCookieName returns the SessionStateCookieName field if non-nil, zero value otherwise.

### GetSessionStateCookieNameOk

`func (o *WebSessionManagement) GetSessionStateCookieNameOk() (*string, bool)`

GetSessionStateCookieNameOk returns a tuple with the SessionStateCookieName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionStateCookieName

`func (o *WebSessionManagement) SetSessionStateCookieName(v string)`

SetSessionStateCookieName sets SessionStateCookieName field to given value.

### HasSessionStateCookieName

`func (o *WebSessionManagement) HasSessionStateCookieName() bool`

HasSessionStateCookieName returns a boolean if a field has been set.

### GetUpdateTokenWindowInSeconds

`func (o *WebSessionManagement) GetUpdateTokenWindowInSeconds() int64`

GetUpdateTokenWindowInSeconds returns the UpdateTokenWindowInSeconds field if non-nil, zero value otherwise.

### GetUpdateTokenWindowInSecondsOk

`func (o *WebSessionManagement) GetUpdateTokenWindowInSecondsOk() (*int64, bool)`

GetUpdateTokenWindowInSecondsOk returns a tuple with the UpdateTokenWindowInSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTokenWindowInSeconds

`func (o *WebSessionManagement) SetUpdateTokenWindowInSeconds(v int64)`

SetUpdateTokenWindowInSeconds sets UpdateTokenWindowInSeconds field to given value.

### HasUpdateTokenWindowInSeconds

`func (o *WebSessionManagement) HasUpdateTokenWindowInSeconds() bool`

HasUpdateTokenWindowInSeconds returns a boolean if a field has been set.

### GetEncryptionAlgorithm

`func (o *WebSessionManagement) GetEncryptionAlgorithm() string`

GetEncryptionAlgorithm returns the EncryptionAlgorithm field if non-nil, zero value otherwise.

### GetEncryptionAlgorithmOk

`func (o *WebSessionManagement) GetEncryptionAlgorithmOk() (*string, bool)`

GetEncryptionAlgorithmOk returns a tuple with the EncryptionAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionAlgorithm

`func (o *WebSessionManagement) SetEncryptionAlgorithm(v string)`

SetEncryptionAlgorithm sets EncryptionAlgorithm field to given value.

### HasEncryptionAlgorithm

`func (o *WebSessionManagement) HasEncryptionAlgorithm() bool`

HasEncryptionAlgorithm returns a boolean if a field has been set.

### GetNonceCookieTimeToLiveInMinutes

`func (o *WebSessionManagement) GetNonceCookieTimeToLiveInMinutes() int32`

GetNonceCookieTimeToLiveInMinutes returns the NonceCookieTimeToLiveInMinutes field if non-nil, zero value otherwise.

### GetNonceCookieTimeToLiveInMinutesOk

`func (o *WebSessionManagement) GetNonceCookieTimeToLiveInMinutesOk() (*int32, bool)`

GetNonceCookieTimeToLiveInMinutesOk returns a tuple with the NonceCookieTimeToLiveInMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonceCookieTimeToLiveInMinutes

`func (o *WebSessionManagement) SetNonceCookieTimeToLiveInMinutes(v int32)`

SetNonceCookieTimeToLiveInMinutes sets NonceCookieTimeToLiveInMinutes field to given value.

### HasNonceCookieTimeToLiveInMinutes

`func (o *WebSessionManagement) HasNonceCookieTimeToLiveInMinutes() bool`

HasNonceCookieTimeToLiveInMinutes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


