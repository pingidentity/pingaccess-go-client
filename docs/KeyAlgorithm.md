# KeyAlgorithm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The algorithm name. | 
**KeySizes** | **[]int64** | The list of available key sizes. | 
**DefaultKeySize** | **int64** | The default key size value to use. | 
**DefaultSignatureAlgorithm** | **string** | The default signature algorithm to use. | 
**SignatureAlgorithms** | **[]string** | The list of available signature algorithms. | 

## Methods

### NewKeyAlgorithm

`func NewKeyAlgorithm(name string, keySizes []int64, defaultKeySize int64, defaultSignatureAlgorithm string, signatureAlgorithms []string, ) *KeyAlgorithm`

NewKeyAlgorithm instantiates a new KeyAlgorithm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyAlgorithmWithDefaults

`func NewKeyAlgorithmWithDefaults() *KeyAlgorithm`

NewKeyAlgorithmWithDefaults instantiates a new KeyAlgorithm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *KeyAlgorithm) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KeyAlgorithm) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KeyAlgorithm) SetName(v string)`

SetName sets Name field to given value.


### GetKeySizes

`func (o *KeyAlgorithm) GetKeySizes() []int64`

GetKeySizes returns the KeySizes field if non-nil, zero value otherwise.

### GetKeySizesOk

`func (o *KeyAlgorithm) GetKeySizesOk() (*[]int64, bool)`

GetKeySizesOk returns a tuple with the KeySizes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeySizes

`func (o *KeyAlgorithm) SetKeySizes(v []int64)`

SetKeySizes sets KeySizes field to given value.


### GetDefaultKeySize

`func (o *KeyAlgorithm) GetDefaultKeySize() int64`

GetDefaultKeySize returns the DefaultKeySize field if non-nil, zero value otherwise.

### GetDefaultKeySizeOk

`func (o *KeyAlgorithm) GetDefaultKeySizeOk() (*int64, bool)`

GetDefaultKeySizeOk returns a tuple with the DefaultKeySize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultKeySize

`func (o *KeyAlgorithm) SetDefaultKeySize(v int64)`

SetDefaultKeySize sets DefaultKeySize field to given value.


### GetDefaultSignatureAlgorithm

`func (o *KeyAlgorithm) GetDefaultSignatureAlgorithm() string`

GetDefaultSignatureAlgorithm returns the DefaultSignatureAlgorithm field if non-nil, zero value otherwise.

### GetDefaultSignatureAlgorithmOk

`func (o *KeyAlgorithm) GetDefaultSignatureAlgorithmOk() (*string, bool)`

GetDefaultSignatureAlgorithmOk returns a tuple with the DefaultSignatureAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSignatureAlgorithm

`func (o *KeyAlgorithm) SetDefaultSignatureAlgorithm(v string)`

SetDefaultSignatureAlgorithm sets DefaultSignatureAlgorithm field to given value.


### GetSignatureAlgorithms

`func (o *KeyAlgorithm) GetSignatureAlgorithms() []string`

GetSignatureAlgorithms returns the SignatureAlgorithms field if non-nil, zero value otherwise.

### GetSignatureAlgorithmsOk

`func (o *KeyAlgorithm) GetSignatureAlgorithmsOk() (*[]string, bool)`

GetSignatureAlgorithmsOk returns a tuple with the SignatureAlgorithms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureAlgorithms

`func (o *KeyAlgorithm) SetSignatureAlgorithms(v []string)`

SetSignatureAlgorithms sets SignatureAlgorithms field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


