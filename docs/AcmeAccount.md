# AcmeAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | When creating a new AcmeAccount, this is the ID for the AcmeAccount. If not specified, an ID will be automatically assigned. When updating an existing AcmeAccount, this field is ignored and the ID is determined by the path parameter. | [optional] 
**Url** | Pointer to **string** | The URL the CA uses to reference the account. | [optional] 
**KeyAlgorithm** | Pointer to **string** | The key algorithm used to generate a key. | [optional] 
**PublicKey** | Pointer to [**PublicKey**](PublicKey.md) |  | [optional] 
**PrivateKey** | Pointer to [**HiddenField**](HiddenField.md) |  | [optional] 
**AcmeServerId** | Pointer to **string** | The associated ACME Server. | [optional] 

## Methods

### NewAcmeAccount

`func NewAcmeAccount() *AcmeAccount`

NewAcmeAccount instantiates a new AcmeAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAcmeAccountWithDefaults

`func NewAcmeAccountWithDefaults() *AcmeAccount`

NewAcmeAccountWithDefaults instantiates a new AcmeAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AcmeAccount) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AcmeAccount) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AcmeAccount) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AcmeAccount) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrl

`func (o *AcmeAccount) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AcmeAccount) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AcmeAccount) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AcmeAccount) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetKeyAlgorithm

`func (o *AcmeAccount) GetKeyAlgorithm() string`

GetKeyAlgorithm returns the KeyAlgorithm field if non-nil, zero value otherwise.

### GetKeyAlgorithmOk

`func (o *AcmeAccount) GetKeyAlgorithmOk() (*string, bool)`

GetKeyAlgorithmOk returns a tuple with the KeyAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyAlgorithm

`func (o *AcmeAccount) SetKeyAlgorithm(v string)`

SetKeyAlgorithm sets KeyAlgorithm field to given value.

### HasKeyAlgorithm

`func (o *AcmeAccount) HasKeyAlgorithm() bool`

HasKeyAlgorithm returns a boolean if a field has been set.

### GetPublicKey

`func (o *AcmeAccount) GetPublicKey() PublicKey`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *AcmeAccount) GetPublicKeyOk() (*PublicKey, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *AcmeAccount) SetPublicKey(v PublicKey)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *AcmeAccount) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### GetPrivateKey

`func (o *AcmeAccount) GetPrivateKey() HiddenField`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *AcmeAccount) GetPrivateKeyOk() (*HiddenField, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *AcmeAccount) SetPrivateKey(v HiddenField)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *AcmeAccount) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### GetAcmeServerId

`func (o *AcmeAccount) GetAcmeServerId() string`

GetAcmeServerId returns the AcmeServerId field if non-nil, zero value otherwise.

### GetAcmeServerIdOk

`func (o *AcmeAccount) GetAcmeServerIdOk() (*string, bool)`

GetAcmeServerIdOk returns a tuple with the AcmeServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcmeServerId

`func (o *AcmeAccount) SetAcmeServerId(v string)`

SetAcmeServerId sets AcmeServerId field to given value.

### HasAcmeServerId

`func (o *AcmeAccount) HasAcmeServerId() bool`

HasAcmeServerId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


