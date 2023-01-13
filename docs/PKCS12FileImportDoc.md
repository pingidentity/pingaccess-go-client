# PKCS12FileImportDoc

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HsmProviderId** | **int32** | The HSM Provider ID. The default value is 0 indicating an HSM is not used for this key pair. | 
**Password** | [**HiddenField**](HiddenField.md) |  | 
**Alias** | **string** | A unique alias name to identify the key pair. Special characters and spaces are allowed. | 
**ChainCertificates** | **[]string** | An array of chain certificates. | 
**FileData** | **string** | Base-64 encoded PKCS12 or PEM file data. In BCFIPS mode, only PEM with PBES2 and AES or Triple DES encryption is accepted and 128-bit salt is required. | 

## Methods

### NewPKCS12FileImportDoc

`func NewPKCS12FileImportDoc(hsmProviderId int32, password HiddenField, alias string, chainCertificates []string, fileData string, ) *PKCS12FileImportDoc`

NewPKCS12FileImportDoc instantiates a new PKCS12FileImportDoc object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPKCS12FileImportDocWithDefaults

`func NewPKCS12FileImportDocWithDefaults() *PKCS12FileImportDoc`

NewPKCS12FileImportDocWithDefaults instantiates a new PKCS12FileImportDoc object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHsmProviderId

`func (o *PKCS12FileImportDoc) GetHsmProviderId() int32`

GetHsmProviderId returns the HsmProviderId field if non-nil, zero value otherwise.

### GetHsmProviderIdOk

`func (o *PKCS12FileImportDoc) GetHsmProviderIdOk() (*int32, bool)`

GetHsmProviderIdOk returns a tuple with the HsmProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHsmProviderId

`func (o *PKCS12FileImportDoc) SetHsmProviderId(v int32)`

SetHsmProviderId sets HsmProviderId field to given value.


### GetPassword

`func (o *PKCS12FileImportDoc) GetPassword() HiddenField`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *PKCS12FileImportDoc) GetPasswordOk() (*HiddenField, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *PKCS12FileImportDoc) SetPassword(v HiddenField)`

SetPassword sets Password field to given value.


### GetAlias

`func (o *PKCS12FileImportDoc) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *PKCS12FileImportDoc) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *PKCS12FileImportDoc) SetAlias(v string)`

SetAlias sets Alias field to given value.


### GetChainCertificates

`func (o *PKCS12FileImportDoc) GetChainCertificates() []string`

GetChainCertificates returns the ChainCertificates field if non-nil, zero value otherwise.

### GetChainCertificatesOk

`func (o *PKCS12FileImportDoc) GetChainCertificatesOk() (*[]string, bool)`

GetChainCertificatesOk returns a tuple with the ChainCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainCertificates

`func (o *PKCS12FileImportDoc) SetChainCertificates(v []string)`

SetChainCertificates sets ChainCertificates field to given value.


### GetFileData

`func (o *PKCS12FileImportDoc) GetFileData() string`

GetFileData returns the FileData field if non-nil, zero value otherwise.

### GetFileDataOk

`func (o *PKCS12FileImportDoc) GetFileDataOk() (*string, bool)`

GetFileDataOk returns a tuple with the FileData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileData

`func (o *PKCS12FileImportDoc) SetFileData(v string)`

SetFileData sets FileData field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


