# SharedSecret

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | When creating a new SharedSecret, this is the ID for the SharedSecret. If not specified, an ID will be automatically assigned. When updating an existing SharedSecret, this field is ignored and the ID is determined by the path parameter. | [optional] 
**Created** | Pointer to **time.Time** | (sortable) The created date of the secret as the number of milliseconds since January 1, 1970, 00:00:00 GMT. | [optional] 
**Secret** | [**HiddenField**](HiddenField.md) |  | 

## Methods

### NewSharedSecret

`func NewSharedSecret(secret HiddenField, ) *SharedSecret`

NewSharedSecret instantiates a new SharedSecret object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSharedSecretWithDefaults

`func NewSharedSecretWithDefaults() *SharedSecret`

NewSharedSecretWithDefaults instantiates a new SharedSecret object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SharedSecret) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SharedSecret) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SharedSecret) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SharedSecret) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreated

`func (o *SharedSecret) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *SharedSecret) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *SharedSecret) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *SharedSecret) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetSecret

`func (o *SharedSecret) GetSecret() HiddenField`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *SharedSecret) GetSecretOk() (*HiddenField, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *SharedSecret) SetSecret(v HiddenField)`

SetSecret sets Secret field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


