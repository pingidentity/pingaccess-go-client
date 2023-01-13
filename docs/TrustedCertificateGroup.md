# TrustedCertificateGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | When creating a new TrustedCertificateGroup, this is the ID for the TrustedCertificateGroup. If not specified, an ID will be automatically assigned. When updating an existing TrustedCertificateGroup, this field is ignored and the ID is determined by the path parameter. | [optional] 
**Name** | **string** | (sortable) The name of the trusted certificate group. | 
**UseJavaTrustStore** | Pointer to **bool** | (sortable) This field is true if the certificates in the group should also include all certificates in the Java Trust Store. | [optional] 
**SystemGroup** | Pointer to **bool** | (sortable) This field is read-only and indicates the trusted certificate group cannot be modified. | [optional] 
**IgnoreAllCertificateErrors** | Pointer to **bool** | (sortable) This field is read-only and is only set to true for the Trust Any certificate group. | [optional] 
**SkipCertificateDateCheck** | Pointer to **bool** | (sortable) This field is true if certificates that have expired or are not yet valid but have passed the other certificate checks should be trusted. | [optional] 
**RevocationChecking** | Pointer to [**RevocationChecking**](RevocationChecking.md) |  | [optional] 
**CertIds** | Pointer to **[]int32** | The IDs of the certificates that are in the trusted certificate group. | [optional] 

## Methods

### NewTrustedCertificateGroup

`func NewTrustedCertificateGroup(name string, ) *TrustedCertificateGroup`

NewTrustedCertificateGroup instantiates a new TrustedCertificateGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrustedCertificateGroupWithDefaults

`func NewTrustedCertificateGroupWithDefaults() *TrustedCertificateGroup`

NewTrustedCertificateGroupWithDefaults instantiates a new TrustedCertificateGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TrustedCertificateGroup) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TrustedCertificateGroup) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TrustedCertificateGroup) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *TrustedCertificateGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *TrustedCertificateGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TrustedCertificateGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TrustedCertificateGroup) SetName(v string)`

SetName sets Name field to given value.


### GetUseJavaTrustStore

`func (o *TrustedCertificateGroup) GetUseJavaTrustStore() bool`

GetUseJavaTrustStore returns the UseJavaTrustStore field if non-nil, zero value otherwise.

### GetUseJavaTrustStoreOk

`func (o *TrustedCertificateGroup) GetUseJavaTrustStoreOk() (*bool, bool)`

GetUseJavaTrustStoreOk returns a tuple with the UseJavaTrustStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseJavaTrustStore

`func (o *TrustedCertificateGroup) SetUseJavaTrustStore(v bool)`

SetUseJavaTrustStore sets UseJavaTrustStore field to given value.

### HasUseJavaTrustStore

`func (o *TrustedCertificateGroup) HasUseJavaTrustStore() bool`

HasUseJavaTrustStore returns a boolean if a field has been set.

### GetSystemGroup

`func (o *TrustedCertificateGroup) GetSystemGroup() bool`

GetSystemGroup returns the SystemGroup field if non-nil, zero value otherwise.

### GetSystemGroupOk

`func (o *TrustedCertificateGroup) GetSystemGroupOk() (*bool, bool)`

GetSystemGroupOk returns a tuple with the SystemGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemGroup

`func (o *TrustedCertificateGroup) SetSystemGroup(v bool)`

SetSystemGroup sets SystemGroup field to given value.

### HasSystemGroup

`func (o *TrustedCertificateGroup) HasSystemGroup() bool`

HasSystemGroup returns a boolean if a field has been set.

### GetIgnoreAllCertificateErrors

`func (o *TrustedCertificateGroup) GetIgnoreAllCertificateErrors() bool`

GetIgnoreAllCertificateErrors returns the IgnoreAllCertificateErrors field if non-nil, zero value otherwise.

### GetIgnoreAllCertificateErrorsOk

`func (o *TrustedCertificateGroup) GetIgnoreAllCertificateErrorsOk() (*bool, bool)`

GetIgnoreAllCertificateErrorsOk returns a tuple with the IgnoreAllCertificateErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreAllCertificateErrors

`func (o *TrustedCertificateGroup) SetIgnoreAllCertificateErrors(v bool)`

SetIgnoreAllCertificateErrors sets IgnoreAllCertificateErrors field to given value.

### HasIgnoreAllCertificateErrors

`func (o *TrustedCertificateGroup) HasIgnoreAllCertificateErrors() bool`

HasIgnoreAllCertificateErrors returns a boolean if a field has been set.

### GetSkipCertificateDateCheck

`func (o *TrustedCertificateGroup) GetSkipCertificateDateCheck() bool`

GetSkipCertificateDateCheck returns the SkipCertificateDateCheck field if non-nil, zero value otherwise.

### GetSkipCertificateDateCheckOk

`func (o *TrustedCertificateGroup) GetSkipCertificateDateCheckOk() (*bool, bool)`

GetSkipCertificateDateCheckOk returns a tuple with the SkipCertificateDateCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipCertificateDateCheck

`func (o *TrustedCertificateGroup) SetSkipCertificateDateCheck(v bool)`

SetSkipCertificateDateCheck sets SkipCertificateDateCheck field to given value.

### HasSkipCertificateDateCheck

`func (o *TrustedCertificateGroup) HasSkipCertificateDateCheck() bool`

HasSkipCertificateDateCheck returns a boolean if a field has been set.

### GetRevocationChecking

`func (o *TrustedCertificateGroup) GetRevocationChecking() RevocationChecking`

GetRevocationChecking returns the RevocationChecking field if non-nil, zero value otherwise.

### GetRevocationCheckingOk

`func (o *TrustedCertificateGroup) GetRevocationCheckingOk() (*RevocationChecking, bool)`

GetRevocationCheckingOk returns a tuple with the RevocationChecking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocationChecking

`func (o *TrustedCertificateGroup) SetRevocationChecking(v RevocationChecking)`

SetRevocationChecking sets RevocationChecking field to given value.

### HasRevocationChecking

`func (o *TrustedCertificateGroup) HasRevocationChecking() bool`

HasRevocationChecking returns a boolean if a field has been set.

### GetCertIds

`func (o *TrustedCertificateGroup) GetCertIds() []int32`

GetCertIds returns the CertIds field if non-nil, zero value otherwise.

### GetCertIdsOk

`func (o *TrustedCertificateGroup) GetCertIdsOk() (*[]int32, bool)`

GetCertIdsOk returns a tuple with the CertIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertIds

`func (o *TrustedCertificateGroup) SetCertIds(v []int32)`

SetCertIds sets CertIds field to given value.

### HasCertIds

`func (o *TrustedCertificateGroup) HasCertIds() bool`

HasCertIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


