# JsonWebKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**K** | Pointer to **string** | The key. | [optional] 
**Kid** | Pointer to **string** | The key id. | [optional] 
**KeyOps** | Pointer to **string** | The list of key operations. | [optional] 
**Kty** | Pointer to **string** | The key type. | [optional] 
**Alg** | Pointer to **string** | The algorithm name. | [optional] 

## Methods

### NewJsonWebKey

`func NewJsonWebKey() *JsonWebKey`

NewJsonWebKey instantiates a new JsonWebKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJsonWebKeyWithDefaults

`func NewJsonWebKeyWithDefaults() *JsonWebKey`

NewJsonWebKeyWithDefaults instantiates a new JsonWebKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetK

`func (o *JsonWebKey) GetK() string`

GetK returns the K field if non-nil, zero value otherwise.

### GetKOk

`func (o *JsonWebKey) GetKOk() (*string, bool)`

GetKOk returns a tuple with the K field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK

`func (o *JsonWebKey) SetK(v string)`

SetK sets K field to given value.

### HasK

`func (o *JsonWebKey) HasK() bool`

HasK returns a boolean if a field has been set.

### GetKid

`func (o *JsonWebKey) GetKid() string`

GetKid returns the Kid field if non-nil, zero value otherwise.

### GetKidOk

`func (o *JsonWebKey) GetKidOk() (*string, bool)`

GetKidOk returns a tuple with the Kid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKid

`func (o *JsonWebKey) SetKid(v string)`

SetKid sets Kid field to given value.

### HasKid

`func (o *JsonWebKey) HasKid() bool`

HasKid returns a boolean if a field has been set.

### GetKeyOps

`func (o *JsonWebKey) GetKeyOps() string`

GetKeyOps returns the KeyOps field if non-nil, zero value otherwise.

### GetKeyOpsOk

`func (o *JsonWebKey) GetKeyOpsOk() (*string, bool)`

GetKeyOpsOk returns a tuple with the KeyOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyOps

`func (o *JsonWebKey) SetKeyOps(v string)`

SetKeyOps sets KeyOps field to given value.

### HasKeyOps

`func (o *JsonWebKey) HasKeyOps() bool`

HasKeyOps returns a boolean if a field has been set.

### GetKty

`func (o *JsonWebKey) GetKty() string`

GetKty returns the Kty field if non-nil, zero value otherwise.

### GetKtyOk

`func (o *JsonWebKey) GetKtyOk() (*string, bool)`

GetKtyOk returns a tuple with the Kty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKty

`func (o *JsonWebKey) SetKty(v string)`

SetKty sets Kty field to given value.

### HasKty

`func (o *JsonWebKey) HasKty() bool`

HasKty returns a boolean if a field has been set.

### GetAlg

`func (o *JsonWebKey) GetAlg() string`

GetAlg returns the Alg field if non-nil, zero value otherwise.

### GetAlgOk

`func (o *JsonWebKey) GetAlgOk() (*string, bool)`

GetAlgOk returns a tuple with the Alg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlg

`func (o *JsonWebKey) SetAlg(v string)`

SetAlg sets Alg field to given value.

### HasAlg

`func (o *JsonWebKey) HasAlg() bool`

HasAlg returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


