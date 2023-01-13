# PublicKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Jwk** | **map[string]interface{}** | The JSON of the public key. | 
**Created** | Pointer to **time.Time** | A timestamp when the key was created. | [optional] 

## Methods

### NewPublicKey

`func NewPublicKey(jwk map[string]interface{}, ) *PublicKey`

NewPublicKey instantiates a new PublicKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicKeyWithDefaults

`func NewPublicKeyWithDefaults() *PublicKey`

NewPublicKeyWithDefaults instantiates a new PublicKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJwk

`func (o *PublicKey) GetJwk() map[string]interface{}`

GetJwk returns the Jwk field if non-nil, zero value otherwise.

### GetJwkOk

`func (o *PublicKey) GetJwkOk() (*map[string]interface{}, bool)`

GetJwkOk returns a tuple with the Jwk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwk

`func (o *PublicKey) SetJwk(v map[string]interface{})`

SetJwk sets Jwk field to given value.


### GetCreated

`func (o *PublicKey) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *PublicKey) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *PublicKey) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *PublicKey) HasCreated() bool`

HasCreated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


