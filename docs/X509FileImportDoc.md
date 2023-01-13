# X509FileImportDoc

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alias** | **string** | The alias for the certificate. | 
**FileData** | **string** | The base64-encoded data of the certificate. | 

## Methods

### NewX509FileImportDoc

`func NewX509FileImportDoc(alias string, fileData string, ) *X509FileImportDoc`

NewX509FileImportDoc instantiates a new X509FileImportDoc object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewX509FileImportDocWithDefaults

`func NewX509FileImportDocWithDefaults() *X509FileImportDoc`

NewX509FileImportDocWithDefaults instantiates a new X509FileImportDoc object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlias

`func (o *X509FileImportDoc) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *X509FileImportDoc) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *X509FileImportDoc) SetAlias(v string)`

SetAlias sets Alias field to given value.


### GetFileData

`func (o *X509FileImportDoc) GetFileData() string`

GetFileData returns the FileData field if non-nil, zero value otherwise.

### GetFileDataOk

`func (o *X509FileImportDoc) GetFileDataOk() (*string, bool)`

GetFileDataOk returns a tuple with the FileData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileData

`func (o *X509FileImportDoc) SetFileData(v string)`

SetFileData sets FileData field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


