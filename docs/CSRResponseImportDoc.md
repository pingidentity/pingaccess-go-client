# CSRResponseImportDoc

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HsmProviderId** | **int32** | The HSM Provider ID. The default value is 0 indicating an HSM is not used for this key pair. | 
**FileData** | **string** | The CSR response data. | 
**ChainCertificates** | **[]string** | A list of base64-encoded certificates to add to the key pair certificate chain. | 
**TrustedCertGroupId** | **int32** | The ID of the trusted certificate group associated with the CSR response. | 

## Methods

### NewCSRResponseImportDoc

`func NewCSRResponseImportDoc(hsmProviderId int32, fileData string, chainCertificates []string, trustedCertGroupId int32, ) *CSRResponseImportDoc`

NewCSRResponseImportDoc instantiates a new CSRResponseImportDoc object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCSRResponseImportDocWithDefaults

`func NewCSRResponseImportDocWithDefaults() *CSRResponseImportDoc`

NewCSRResponseImportDocWithDefaults instantiates a new CSRResponseImportDoc object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHsmProviderId

`func (o *CSRResponseImportDoc) GetHsmProviderId() int32`

GetHsmProviderId returns the HsmProviderId field if non-nil, zero value otherwise.

### GetHsmProviderIdOk

`func (o *CSRResponseImportDoc) GetHsmProviderIdOk() (*int32, bool)`

GetHsmProviderIdOk returns a tuple with the HsmProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHsmProviderId

`func (o *CSRResponseImportDoc) SetHsmProviderId(v int32)`

SetHsmProviderId sets HsmProviderId field to given value.


### GetFileData

`func (o *CSRResponseImportDoc) GetFileData() string`

GetFileData returns the FileData field if non-nil, zero value otherwise.

### GetFileDataOk

`func (o *CSRResponseImportDoc) GetFileDataOk() (*string, bool)`

GetFileDataOk returns a tuple with the FileData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileData

`func (o *CSRResponseImportDoc) SetFileData(v string)`

SetFileData sets FileData field to given value.


### GetChainCertificates

`func (o *CSRResponseImportDoc) GetChainCertificates() []string`

GetChainCertificates returns the ChainCertificates field if non-nil, zero value otherwise.

### GetChainCertificatesOk

`func (o *CSRResponseImportDoc) GetChainCertificatesOk() (*[]string, bool)`

GetChainCertificatesOk returns a tuple with the ChainCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainCertificates

`func (o *CSRResponseImportDoc) SetChainCertificates(v []string)`

SetChainCertificates sets ChainCertificates field to given value.


### GetTrustedCertGroupId

`func (o *CSRResponseImportDoc) GetTrustedCertGroupId() int32`

GetTrustedCertGroupId returns the TrustedCertGroupId field if non-nil, zero value otherwise.

### GetTrustedCertGroupIdOk

`func (o *CSRResponseImportDoc) GetTrustedCertGroupIdOk() (*int32, bool)`

GetTrustedCertGroupIdOk returns a tuple with the TrustedCertGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustedCertGroupId

`func (o *CSRResponseImportDoc) SetTrustedCertGroupId(v int32)`

SetTrustedCertGroupId sets TrustedCertGroupId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


