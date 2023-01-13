# RevocationChecking

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CrlChecking** | Pointer to **bool** |  | [optional] 
**Ocsp** | Pointer to **bool** |  | [optional] 
**DenyRevocationStatusUnknown** | Pointer to **bool** |  | [optional] 
**SupportDisorderedChain** | Pointer to **bool** |  | [optional] 
**SkipTrustAnchors** | Pointer to **bool** |  | [optional] 

## Methods

### NewRevocationChecking

`func NewRevocationChecking() *RevocationChecking`

NewRevocationChecking instantiates a new RevocationChecking object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRevocationCheckingWithDefaults

`func NewRevocationCheckingWithDefaults() *RevocationChecking`

NewRevocationCheckingWithDefaults instantiates a new RevocationChecking object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCrlChecking

`func (o *RevocationChecking) GetCrlChecking() bool`

GetCrlChecking returns the CrlChecking field if non-nil, zero value otherwise.

### GetCrlCheckingOk

`func (o *RevocationChecking) GetCrlCheckingOk() (*bool, bool)`

GetCrlCheckingOk returns a tuple with the CrlChecking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrlChecking

`func (o *RevocationChecking) SetCrlChecking(v bool)`

SetCrlChecking sets CrlChecking field to given value.

### HasCrlChecking

`func (o *RevocationChecking) HasCrlChecking() bool`

HasCrlChecking returns a boolean if a field has been set.

### GetOcsp

`func (o *RevocationChecking) GetOcsp() bool`

GetOcsp returns the Ocsp field if non-nil, zero value otherwise.

### GetOcspOk

`func (o *RevocationChecking) GetOcspOk() (*bool, bool)`

GetOcspOk returns a tuple with the Ocsp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOcsp

`func (o *RevocationChecking) SetOcsp(v bool)`

SetOcsp sets Ocsp field to given value.

### HasOcsp

`func (o *RevocationChecking) HasOcsp() bool`

HasOcsp returns a boolean if a field has been set.

### GetDenyRevocationStatusUnknown

`func (o *RevocationChecking) GetDenyRevocationStatusUnknown() bool`

GetDenyRevocationStatusUnknown returns the DenyRevocationStatusUnknown field if non-nil, zero value otherwise.

### GetDenyRevocationStatusUnknownOk

`func (o *RevocationChecking) GetDenyRevocationStatusUnknownOk() (*bool, bool)`

GetDenyRevocationStatusUnknownOk returns a tuple with the DenyRevocationStatusUnknown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenyRevocationStatusUnknown

`func (o *RevocationChecking) SetDenyRevocationStatusUnknown(v bool)`

SetDenyRevocationStatusUnknown sets DenyRevocationStatusUnknown field to given value.

### HasDenyRevocationStatusUnknown

`func (o *RevocationChecking) HasDenyRevocationStatusUnknown() bool`

HasDenyRevocationStatusUnknown returns a boolean if a field has been set.

### GetSupportDisorderedChain

`func (o *RevocationChecking) GetSupportDisorderedChain() bool`

GetSupportDisorderedChain returns the SupportDisorderedChain field if non-nil, zero value otherwise.

### GetSupportDisorderedChainOk

`func (o *RevocationChecking) GetSupportDisorderedChainOk() (*bool, bool)`

GetSupportDisorderedChainOk returns a tuple with the SupportDisorderedChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportDisorderedChain

`func (o *RevocationChecking) SetSupportDisorderedChain(v bool)`

SetSupportDisorderedChain sets SupportDisorderedChain field to given value.

### HasSupportDisorderedChain

`func (o *RevocationChecking) HasSupportDisorderedChain() bool`

HasSupportDisorderedChain returns a boolean if a field has been set.

### GetSkipTrustAnchors

`func (o *RevocationChecking) GetSkipTrustAnchors() bool`

GetSkipTrustAnchors returns the SkipTrustAnchors field if non-nil, zero value otherwise.

### GetSkipTrustAnchorsOk

`func (o *RevocationChecking) GetSkipTrustAnchorsOk() (*bool, bool)`

GetSkipTrustAnchorsOk returns a tuple with the SkipTrustAnchors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipTrustAnchors

`func (o *RevocationChecking) SetSkipTrustAnchors(v bool)`

SetSkipTrustAnchors sets SkipTrustAnchors field to given value.

### HasSkipTrustAnchors

`func (o *RevocationChecking) HasSkipTrustAnchors() bool`

HasSkipTrustAnchors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


