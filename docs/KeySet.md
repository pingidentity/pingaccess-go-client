# KeySet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nonce** | **string** | The nonce used to derive the key used to encrypt the keySet value. | 
**KeySet** | **string** | The encrypted key set. This value is bound to the nonce field and cannot be decrypted without also submitting the nonce used to produce the value. | 

## Methods

### NewKeySet

`func NewKeySet(nonce string, keySet string, ) *KeySet`

NewKeySet instantiates a new KeySet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeySetWithDefaults

`func NewKeySetWithDefaults() *KeySet`

NewKeySetWithDefaults instantiates a new KeySet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNonce

`func (o *KeySet) GetNonce() string`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *KeySet) GetNonceOk() (*string, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *KeySet) SetNonce(v string)`

SetNonce sets Nonce field to given value.


### GetKeySet

`func (o *KeySet) GetKeySet() string`

GetKeySet returns the KeySet field if non-nil, zero value otherwise.

### GetKeySetOk

`func (o *KeySet) GetKeySetOk() (*string, bool)`

GetKeySetOk returns a tuple with the KeySet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeySet

`func (o *KeySet) SetKeySet(v string)`

SetKeySet sets KeySet field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


