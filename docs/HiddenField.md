# HiddenField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **string** | The plaintext value | [optional] 
**EncryptedValue** | Pointer to **string** | The ciphertext value | [optional] 

## Methods

### NewHiddenField

`func NewHiddenField() *HiddenField`

NewHiddenField instantiates a new HiddenField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHiddenFieldWithDefaults

`func NewHiddenFieldWithDefaults() *HiddenField`

NewHiddenFieldWithDefaults instantiates a new HiddenField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *HiddenField) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *HiddenField) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *HiddenField) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *HiddenField) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetEncryptedValue

`func (o *HiddenField) GetEncryptedValue() string`

GetEncryptedValue returns the EncryptedValue field if non-nil, zero value otherwise.

### GetEncryptedValueOk

`func (o *HiddenField) GetEncryptedValueOk() (*string, bool)`

GetEncryptedValueOk returns a tuple with the EncryptedValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedValue

`func (o *HiddenField) SetEncryptedValue(v string)`

SetEncryptedValue sets EncryptedValue field to given value.

### HasEncryptedValue

`func (o *HiddenField) HasEncryptedValue() bool`

HasEncryptedValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


