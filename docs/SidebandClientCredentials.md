# SidebandClientCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** | The created date of the credentials as the number of milliseconds since January 1, 1970, 00:00:00 GMT. | [optional] 
**CredentialsType** | Pointer to [**SidebandClientCredentialsType**](SidebandClientCredentialsType.md) |  | [optional] 
**SharedSecretConfig** | [**SidebandSharedSecretConfig**](SidebandSharedSecretConfig.md) |  | 

## Methods

### NewSidebandClientCredentials

`func NewSidebandClientCredentials(sharedSecretConfig SidebandSharedSecretConfig, ) *SidebandClientCredentials`

NewSidebandClientCredentials instantiates a new SidebandClientCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSidebandClientCredentialsWithDefaults

`func NewSidebandClientCredentialsWithDefaults() *SidebandClientCredentials`

NewSidebandClientCredentialsWithDefaults instantiates a new SidebandClientCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *SidebandClientCredentials) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *SidebandClientCredentials) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *SidebandClientCredentials) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *SidebandClientCredentials) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCredentialsType

`func (o *SidebandClientCredentials) GetCredentialsType() SidebandClientCredentialsType`

GetCredentialsType returns the CredentialsType field if non-nil, zero value otherwise.

### GetCredentialsTypeOk

`func (o *SidebandClientCredentials) GetCredentialsTypeOk() (*SidebandClientCredentialsType, bool)`

GetCredentialsTypeOk returns a tuple with the CredentialsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialsType

`func (o *SidebandClientCredentials) SetCredentialsType(v SidebandClientCredentialsType)`

SetCredentialsType sets CredentialsType field to given value.

### HasCredentialsType

`func (o *SidebandClientCredentials) HasCredentialsType() bool`

HasCredentialsType returns a boolean if a field has been set.

### GetSharedSecretConfig

`func (o *SidebandClientCredentials) GetSharedSecretConfig() SidebandSharedSecretConfig`

GetSharedSecretConfig returns the SharedSecretConfig field if non-nil, zero value otherwise.

### GetSharedSecretConfigOk

`func (o *SidebandClientCredentials) GetSharedSecretConfigOk() (*SidebandSharedSecretConfig, bool)`

GetSharedSecretConfigOk returns a tuple with the SharedSecretConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedSecretConfig

`func (o *SidebandClientCredentials) SetSharedSecretConfig(v SidebandSharedSecretConfig)`

SetSharedSecretConfig sets SharedSecretConfig field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


