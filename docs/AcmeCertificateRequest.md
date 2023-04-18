# AcmeCertificateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | When creating a new AcmeCertificateRequest, this is the ID for the AcmeCertificateRequest. If not specified, an ID will be automatically assigned. When updating an existing AcmeCertificateRequest, this field is ignored and the ID is determined by the path parameter. | [optional] 
**KeyPairId** | **int64** | The ID of the Key Pair for which a signed certificate will be requested. | 
**AcmeCertStatus** | Pointer to [**AcmeCertStatus**](AcmeCertStatus.md) |  | [optional] 
**Url** | Pointer to **string** | The URL at the ACME server for the associated ACME order. | [optional] 
**AcmeServerId** | Pointer to **string** | The ID of the ACME Server to be used for the ACME protocol. This is read-only. | [optional] 
**AcmeAccountId** | Pointer to **string** | The ID of the ACME Account to be used for the ACME protocol. This is read-only. | [optional] 

## Methods

### NewAcmeCertificateRequest

`func NewAcmeCertificateRequest(keyPairId int64, ) *AcmeCertificateRequest`

NewAcmeCertificateRequest instantiates a new AcmeCertificateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAcmeCertificateRequestWithDefaults

`func NewAcmeCertificateRequestWithDefaults() *AcmeCertificateRequest`

NewAcmeCertificateRequestWithDefaults instantiates a new AcmeCertificateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AcmeCertificateRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AcmeCertificateRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AcmeCertificateRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AcmeCertificateRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKeyPairId

`func (o *AcmeCertificateRequest) GetKeyPairId() int64`

GetKeyPairId returns the KeyPairId field if non-nil, zero value otherwise.

### GetKeyPairIdOk

`func (o *AcmeCertificateRequest) GetKeyPairIdOk() (*int64, bool)`

GetKeyPairIdOk returns a tuple with the KeyPairId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyPairId

`func (o *AcmeCertificateRequest) SetKeyPairId(v int64)`

SetKeyPairId sets KeyPairId field to given value.


### GetAcmeCertStatus

`func (o *AcmeCertificateRequest) GetAcmeCertStatus() AcmeCertStatus`

GetAcmeCertStatus returns the AcmeCertStatus field if non-nil, zero value otherwise.

### GetAcmeCertStatusOk

`func (o *AcmeCertificateRequest) GetAcmeCertStatusOk() (*AcmeCertStatus, bool)`

GetAcmeCertStatusOk returns a tuple with the AcmeCertStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcmeCertStatus

`func (o *AcmeCertificateRequest) SetAcmeCertStatus(v AcmeCertStatus)`

SetAcmeCertStatus sets AcmeCertStatus field to given value.

### HasAcmeCertStatus

`func (o *AcmeCertificateRequest) HasAcmeCertStatus() bool`

HasAcmeCertStatus returns a boolean if a field has been set.

### GetUrl

`func (o *AcmeCertificateRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AcmeCertificateRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AcmeCertificateRequest) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AcmeCertificateRequest) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetAcmeServerId

`func (o *AcmeCertificateRequest) GetAcmeServerId() string`

GetAcmeServerId returns the AcmeServerId field if non-nil, zero value otherwise.

### GetAcmeServerIdOk

`func (o *AcmeCertificateRequest) GetAcmeServerIdOk() (*string, bool)`

GetAcmeServerIdOk returns a tuple with the AcmeServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcmeServerId

`func (o *AcmeCertificateRequest) SetAcmeServerId(v string)`

SetAcmeServerId sets AcmeServerId field to given value.

### HasAcmeServerId

`func (o *AcmeCertificateRequest) HasAcmeServerId() bool`

HasAcmeServerId returns a boolean if a field has been set.

### GetAcmeAccountId

`func (o *AcmeCertificateRequest) GetAcmeAccountId() string`

GetAcmeAccountId returns the AcmeAccountId field if non-nil, zero value otherwise.

### GetAcmeAccountIdOk

`func (o *AcmeCertificateRequest) GetAcmeAccountIdOk() (*string, bool)`

GetAcmeAccountIdOk returns a tuple with the AcmeAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcmeAccountId

`func (o *AcmeCertificateRequest) SetAcmeAccountId(v string)`

SetAcmeAccountId sets AcmeAccountId field to given value.

### HasAcmeAccountId

`func (o *AcmeCertificateRequest) HasAcmeAccountId() bool`

HasAcmeAccountId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


