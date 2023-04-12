# MasterKeys

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EncryptedValue** | Pointer to **[]string** | The encrypted master key. | [optional] 
**KeyId** | Pointer to **string** | The key identifier. | [optional] 

## Methods

### NewMasterKeys

`func NewMasterKeys() *MasterKeys`

NewMasterKeys instantiates a new MasterKeys object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMasterKeysWithDefaults

`func NewMasterKeysWithDefaults() *MasterKeys`

NewMasterKeysWithDefaults instantiates a new MasterKeys object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEncryptedValue

`func (o *MasterKeys) GetEncryptedValue() []string`

GetEncryptedValue returns the EncryptedValue field if non-nil, zero value otherwise.

### GetEncryptedValueOk

`func (o *MasterKeys) GetEncryptedValueOk() (*[]string, bool)`

GetEncryptedValueOk returns a tuple with the EncryptedValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedValue

`func (o *MasterKeys) SetEncryptedValue(v []string)`

SetEncryptedValue sets EncryptedValue field to given value.

### HasEncryptedValue

`func (o *MasterKeys) HasEncryptedValue() bool`

HasEncryptedValue returns a boolean if a field has been set.

### GetKeyId

`func (o *MasterKeys) GetKeyId() string`

GetKeyId returns the KeyId field if non-nil, zero value otherwise.

### GetKeyIdOk

`func (o *MasterKeys) GetKeyIdOk() (*string, bool)`

GetKeyIdOk returns a tuple with the KeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyId

`func (o *MasterKeys) SetKeyId(v string)`

SetKeyId sets KeyId field to given value.

### HasKeyId

`func (o *MasterKeys) HasKeyId() bool`

HasKeyId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


