# SidebandSharedSecretConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SharedSecretHeaderName** | Pointer to **string** | The name of the HTTP header presented in the request by the sideband client. The default value is \&quot;CLIENT-TOKEN\&quot;. | [optional] 
**Secret** | Pointer to [**HiddenField**](HiddenField.md) |  | [optional] 

## Methods

### NewSidebandSharedSecretConfig

`func NewSidebandSharedSecretConfig() *SidebandSharedSecretConfig`

NewSidebandSharedSecretConfig instantiates a new SidebandSharedSecretConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSidebandSharedSecretConfigWithDefaults

`func NewSidebandSharedSecretConfigWithDefaults() *SidebandSharedSecretConfig`

NewSidebandSharedSecretConfigWithDefaults instantiates a new SidebandSharedSecretConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSharedSecretHeaderName

`func (o *SidebandSharedSecretConfig) GetSharedSecretHeaderName() string`

GetSharedSecretHeaderName returns the SharedSecretHeaderName field if non-nil, zero value otherwise.

### GetSharedSecretHeaderNameOk

`func (o *SidebandSharedSecretConfig) GetSharedSecretHeaderNameOk() (*string, bool)`

GetSharedSecretHeaderNameOk returns a tuple with the SharedSecretHeaderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedSecretHeaderName

`func (o *SidebandSharedSecretConfig) SetSharedSecretHeaderName(v string)`

SetSharedSecretHeaderName sets SharedSecretHeaderName field to given value.

### HasSharedSecretHeaderName

`func (o *SidebandSharedSecretConfig) HasSharedSecretHeaderName() bool`

HasSharedSecretHeaderName returns a boolean if a field has been set.

### GetSecret

`func (o *SidebandSharedSecretConfig) GetSecret() HiddenField`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *SidebandSharedSecretConfig) GetSecretOk() (*HiddenField, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *SidebandSharedSecretConfig) SetSecret(v HiddenField)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *SidebandSharedSecretConfig) HasSecret() bool`

HasSecret returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


