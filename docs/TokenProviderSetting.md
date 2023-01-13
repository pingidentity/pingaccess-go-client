# TokenProviderSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UseThirdParty** | Pointer to **bool** | This field is true if third-party Token Provider is enabled, and false if PingFederate is enabled. (DEPRECATED - to be removed in a future release; please use &#39;type&#39; instead) | [optional] 
**Type** | Pointer to [**TokenProviderType**](TokenProviderType.md) |  | [optional] 

## Methods

### NewTokenProviderSetting

`func NewTokenProviderSetting() *TokenProviderSetting`

NewTokenProviderSetting instantiates a new TokenProviderSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenProviderSettingWithDefaults

`func NewTokenProviderSettingWithDefaults() *TokenProviderSetting`

NewTokenProviderSettingWithDefaults instantiates a new TokenProviderSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUseThirdParty

`func (o *TokenProviderSetting) GetUseThirdParty() bool`

GetUseThirdParty returns the UseThirdParty field if non-nil, zero value otherwise.

### GetUseThirdPartyOk

`func (o *TokenProviderSetting) GetUseThirdPartyOk() (*bool, bool)`

GetUseThirdPartyOk returns a tuple with the UseThirdParty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseThirdParty

`func (o *TokenProviderSetting) SetUseThirdParty(v bool)`

SetUseThirdParty sets UseThirdParty field to given value.

### HasUseThirdParty

`func (o *TokenProviderSetting) HasUseThirdParty() bool`

HasUseThirdParty returns a boolean if a field has been set.

### GetType

`func (o *TokenProviderSetting) GetType() TokenProviderType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TokenProviderSetting) GetTypeOk() (*TokenProviderType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TokenProviderSetting) SetType(v TokenProviderType)`

SetType sets Type field to given value.

### HasType

`func (o *TokenProviderSetting) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


