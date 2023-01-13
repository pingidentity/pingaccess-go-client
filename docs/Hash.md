# Hash

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Algorithm** | Pointer to [**HashAlgorithm**](HashAlgorithm.md) |  | [optional] 
**HexValue** | Pointer to **string** |  | [optional] 

## Methods

### NewHash

`func NewHash() *Hash`

NewHash instantiates a new Hash object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHashWithDefaults

`func NewHashWithDefaults() *Hash`

NewHashWithDefaults instantiates a new Hash object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlgorithm

`func (o *Hash) GetAlgorithm() HashAlgorithm`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *Hash) GetAlgorithmOk() (*HashAlgorithm, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *Hash) SetAlgorithm(v HashAlgorithm)`

SetAlgorithm sets Algorithm field to given value.

### HasAlgorithm

`func (o *Hash) HasAlgorithm() bool`

HasAlgorithm returns a boolean if a field has been set.

### GetHexValue

`func (o *Hash) GetHexValue() string`

GetHexValue returns the HexValue field if non-nil, zero value otherwise.

### GetHexValueOk

`func (o *Hash) GetHexValueOk() (*string, bool)`

GetHexValueOk returns a tuple with the HexValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHexValue

`func (o *Hash) SetHexValue(v string)`

SetHexValue sets HexValue field to given value.

### HasHexValue

`func (o *Hash) HasHexValue() bool`

HasHexValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


